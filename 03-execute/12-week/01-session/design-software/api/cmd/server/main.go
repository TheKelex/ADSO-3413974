package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type api struct {
	db *pgxpool.Pool
}

type instructor struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	InstructorType string `json:"instructorType"`
}

type program struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
}

type ficha struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	ProgramID   *int64 `json:"programId,omitempty"`
	ProgramName string `json:"programName,omitempty"`
}

type environment struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type scheduleRequest struct {
	Date             string `json:"date"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	InstructorID     int64  `json:"instructorId"`
	FichaID          int64  `json:"fichaId"`
	EnvironmentID    int64  `json:"environmentId"`
	ProgramID        *int64 `json:"programId,omitempty"`
	Competence       string `json:"competence"`
	RAP              string `json:"rap"`
	LearningActivity string `json:"learningActivity"`
	Observation      string `json:"observation"`
}

type schedule struct {
	ID                  int64  `json:"id"`
	Date                string `json:"date"`
	StartTime           string `json:"startTime"`
	EndTime             string `json:"endTime"`
	InstructorID        int64  `json:"instructorId"`
	InstructorName      string `json:"instructorName"`
	InstructorType      string `json:"instructorType"`
	FichaID             int64  `json:"fichaId"`
	FichaCode           string `json:"fichaCode"`
	EnvironmentID       int64  `json:"environmentId"`
	EnvironmentName     string `json:"environmentName"`
	ProgramID           *int64 `json:"programId,omitempty"`
	ProgramName         string `json:"programName,omitempty"`
	ProgramLevel        string `json:"programLevel,omitempty"`
	Competence          string `json:"competence"`
	RAP                 string `json:"rap"`
	LearningActivity    string `json:"learningActivity"`
	Observation         string `json:"observation"`
	ConflictFingerprint string `json:"conflictFingerprint,omitempty"`
}

type conflict struct {
	Kind      string `json:"kind"`
	EntryID   int64  `json:"entryId"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type noProgrammingDayRequest struct {
	Date         string `json:"date"`
	FichaID      *int64 `json:"fichaId,omitempty"`
	InstructorID *int64 `json:"instructorId,omitempty"`
	Reason       string `json:"reason"`
}

type noProgrammingDay struct {
	ID             int64  `json:"id"`
	Date           string `json:"date"`
	FichaID        *int64 `json:"fichaId,omitempty"`
	FichaCode      string `json:"fichaCode,omitempty"`
	InstructorID   *int64 `json:"instructorId,omitempty"`
	InstructorName string `json:"instructorName,omitempty"`
	Reason         string `json:"reason"`
}

func main() {
	ctx := context.Background()
	databaseURL := getenv("DATABASE_URL", "postgres://sena:sena@localhost:5432/sena_horarios?sslmode=disable")
	port := getenv("PORT", "8080")

	pool, err := connect(ctx, databaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()

	app := &api{db: pool}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", app.health)
	mux.HandleFunc("GET /api/instructors", app.listInstructors)
	mux.HandleFunc("POST /api/instructors", app.createInstructor)
	mux.HandleFunc("GET /api/programs", app.listPrograms)
	mux.HandleFunc("POST /api/programs", app.createProgram)
	mux.HandleFunc("GET /api/fichas", app.listFichas)
	mux.HandleFunc("POST /api/fichas", app.createFicha)
	mux.HandleFunc("GET /api/environments", app.listEnvironments)
	mux.HandleFunc("POST /api/environments", app.createEnvironment)
	mux.HandleFunc("GET /api/schedules", app.listSchedules)
	mux.HandleFunc("POST /api/schedules", app.createSchedule)
	mux.HandleFunc("GET /api/no-programming-days", app.listNoProgrammingDays)
	mux.HandleFunc("POST /api/no-programming-days", app.createNoProgrammingDay)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           cors(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("SENA horarios API listening on :%s", port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	var lastErr error
	for i := 0; i < 20; i++ {
		pool, err := pgxpool.New(ctx, databaseURL)
		if err == nil {
			pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
			err = pool.Ping(pingCtx)
			cancel()
			if err == nil {
				return pool, nil
			}
			pool.Close()
		}
		lastErr = err
		time.Sleep(1 * time.Second)
	}
	return nil, lastErr
}

func (a *api) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := a.db.Ping(ctx); err != nil {
		writeError(w, http.StatusServiceUnavailable, "database_unavailable", err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (a *api) listInstructors(w http.ResponseWriter, r *http.Request) {
	rows, err := a.db.Query(r.Context(), `select id, name, instructor_type from instructors order by name`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}
	defer rows.Close()

	items := []instructor{}
	for rows.Next() {
		var item instructor
		if err := rows.Scan(&item.ID, &item.Name, &item.InstructorType); err != nil {
			writeError(w, http.StatusInternalServerError, "scan_failed", err.Error())
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func (a *api) createInstructor(w http.ResponseWriter, r *http.Request) {
	var input instructor
	if !decode(w, r, &input) {
		return
	}
	input.Name = clean(input.Name)
	input.InstructorType = strings.ToLower(clean(input.InstructorType))
	if input.Name == "" || (input.InstructorType != "planta" && input.InstructorType != "contratista") {
		writeError(w, http.StatusBadRequest, "invalid_instructor", "name and instructorType planta|contratista are required")
		return
	}
	err := a.db.QueryRow(r.Context(), `insert into instructors (name, instructor_type) values ($1, $2) returning id`, input.Name, input.InstructorType).Scan(&input.ID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "insert_failed", err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, input)
}

func (a *api) listPrograms(w http.ResponseWriter, r *http.Request) {
	rows, err := a.db.Query(r.Context(), `select id, name, coalesce(level, '') from training_programs order by name`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}
	defer rows.Close()
	items := []program{}
	for rows.Next() {
		var item program
		if err := rows.Scan(&item.ID, &item.Name, &item.Level); err != nil {
			writeError(w, http.StatusInternalServerError, "scan_failed", err.Error())
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func (a *api) createProgram(w http.ResponseWriter, r *http.Request) {
	var input program
	if !decode(w, r, &input) {
		return
	}
	input.Name = clean(input.Name)
	input.Level = clean(input.Level)
	if input.Name == "" {
		writeError(w, http.StatusBadRequest, "invalid_program", "name is required")
		return
	}
	err := a.db.QueryRow(r.Context(), `insert into training_programs (name, level) values ($1, $2) returning id`, input.Name, nullIfEmpty(input.Level)).Scan(&input.ID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "insert_failed", err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, input)
}

func (a *api) listFichas(w http.ResponseWriter, r *http.Request) {
	rows, err := a.db.Query(r.Context(), `
		select f.id, f.code, f.program_id, coalesce(p.name, '')
		from fichas f
		left join training_programs p on p.id = f.program_id
		order by f.code`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}
	defer rows.Close()
	items := []ficha{}
	for rows.Next() {
		var item ficha
		if err := rows.Scan(&item.ID, &item.Code, &item.ProgramID, &item.ProgramName); err != nil {
			writeError(w, http.StatusInternalServerError, "scan_failed", err.Error())
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func (a *api) createFicha(w http.ResponseWriter, r *http.Request) {
	var input ficha
	if !decode(w, r, &input) {
		return
	}
	input.Code = clean(input.Code)
	if input.Code == "" {
		writeError(w, http.StatusBadRequest, "invalid_ficha", "code is required")
		return
	}
	err := a.db.QueryRow(r.Context(), `insert into fichas (code, program_id) values ($1, $2) returning id`, input.Code, input.ProgramID).Scan(&input.ID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "insert_failed", err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, input)
}

func (a *api) listEnvironments(w http.ResponseWriter, r *http.Request) {
	rows, err := a.db.Query(r.Context(), `select id, name from training_environments order by name`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}
	defer rows.Close()
	items := []environment{}
	for rows.Next() {
		var item environment
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			writeError(w, http.StatusInternalServerError, "scan_failed", err.Error())
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func (a *api) createEnvironment(w http.ResponseWriter, r *http.Request) {
	var input environment
	if !decode(w, r, &input) {
		return
	}
	input.Name = clean(input.Name)
	if input.Name == "" {
		writeError(w, http.StatusBadRequest, "invalid_environment", "name is required")
		return
	}
	err := a.db.QueryRow(r.Context(), `insert into training_environments (name) values ($1) returning id`, input.Name).Scan(&input.ID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "insert_failed", err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, input)
}

func (a *api) listSchedules(w http.ResponseWriter, r *http.Request) {
	query := `
		select s.id, to_char(s.date, 'YYYY-MM-DD'), to_char(s.start_time, 'HH24:MI'), to_char(s.end_time, 'HH24:MI'),
			s.instructor_id, i.name, i.instructor_type,
			s.ficha_id, f.code,
			s.environment_id, e.name,
			s.program_id, coalesce(p.name, ''), coalesce(p.level, ''),
			coalesce(s.competence, ''), coalesce(s.rap, ''), coalesce(s.learning_activity, ''), coalesce(s.observation, '')
		from schedule_entries s
		join instructors i on i.id = s.instructor_id
		join fichas f on f.id = s.ficha_id
		join training_environments e on e.id = s.environment_id
		left join training_programs p on p.id = s.program_id
		order by s.date, s.start_time, f.code`
	rows, err := a.db.Query(r.Context(), query)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}
	defer rows.Close()

	items := []schedule{}
	for rows.Next() {
		var item schedule
		if err := rows.Scan(&item.ID, &item.Date, &item.StartTime, &item.EndTime, &item.InstructorID, &item.InstructorName, &item.InstructorType, &item.FichaID, &item.FichaCode, &item.EnvironmentID, &item.EnvironmentName, &item.ProgramID, &item.ProgramName, &item.ProgramLevel, &item.Competence, &item.RAP, &item.LearningActivity, &item.Observation); err != nil {
			writeError(w, http.StatusInternalServerError, "scan_failed", err.Error())
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func (a *api) createSchedule(w http.ResponseWriter, r *http.Request) {
	var input scheduleRequest
	if !decode(w, r, &input) {
		return
	}
	if err := validateSchedule(input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_schedule", err.Error())
		return
	}

	conflicts, err := a.findConflicts(r.Context(), input)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "conflict_check_failed", err.Error())
		return
	}
	if len(conflicts) > 0 {
		writeJSON(w, http.StatusConflict, map[string]any{
			"error":     "schedule_conflict",
			"message":   "La programacion cruza con instructor, ficha o ambiente ya asignado.",
			"conflicts": conflicts,
		})
		return
	}

	var id int64
	err = a.db.QueryRow(r.Context(), `
		insert into schedule_entries
			(date, start_time, end_time, instructor_id, ficha_id, environment_id, program_id, competence, rap, learning_activity, observation)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		returning id`,
		input.Date, input.StartTime, input.EndTime, input.InstructorID, input.FichaID, input.EnvironmentID, input.ProgramID,
		nullIfEmpty(input.Competence), nullIfEmpty(input.RAP), nullIfEmpty(input.LearningActivity), nullIfEmpty(input.Observation),
	).Scan(&id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "insert_failed", err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"id": id})
}

func (a *api) findConflicts(ctx context.Context, input scheduleRequest) ([]conflict, error) {
	rows, err := a.db.Query(ctx, `
		select 'instructor', id, to_char(date, 'YYYY-MM-DD'), to_char(start_time, 'HH24:MI'), to_char(end_time, 'HH24:MI')
		from schedule_entries
		where date = $1 and start_time < $3 and end_time > $2 and instructor_id = $4
		union all
		select 'ficha', id, to_char(date, 'YYYY-MM-DD'), to_char(start_time, 'HH24:MI'), to_char(end_time, 'HH24:MI')
		from schedule_entries
		where date = $1 and start_time < $3 and end_time > $2 and ficha_id = $5
		union all
		select 'ambiente', id, to_char(date, 'YYYY-MM-DD'), to_char(start_time, 'HH24:MI'), to_char(end_time, 'HH24:MI')
		from schedule_entries
		where date = $1 and start_time < $3 and end_time > $2 and environment_id = $6`,
		input.Date, input.StartTime, input.EndTime, input.InstructorID, input.FichaID, input.EnvironmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	conflicts := []conflict{}
	for rows.Next() {
		var item conflict
		if err := rows.Scan(&item.Kind, &item.EntryID, &item.Date, &item.StartTime, &item.EndTime); err != nil {
			return nil, err
		}
		conflicts = append(conflicts, item)
	}
	return conflicts, rows.Err()
}

func (a *api) listNoProgrammingDays(w http.ResponseWriter, r *http.Request) {
	rows, err := a.db.Query(r.Context(), `
		select n.id, to_char(n.date, 'YYYY-MM-DD'), n.ficha_id, coalesce(f.code, ''), n.instructor_id, coalesce(i.name, ''), n.reason
		from no_programming_days n
		left join fichas f on f.id = n.ficha_id
		left join instructors i on i.id = n.instructor_id
		order by n.date desc, n.id desc`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}
	defer rows.Close()
	items := []noProgrammingDay{}
	for rows.Next() {
		var item noProgrammingDay
		if err := rows.Scan(&item.ID, &item.Date, &item.FichaID, &item.FichaCode, &item.InstructorID, &item.InstructorName, &item.Reason); err != nil {
			writeError(w, http.StatusInternalServerError, "scan_failed", err.Error())
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func (a *api) createNoProgrammingDay(w http.ResponseWriter, r *http.Request) {
	var input noProgrammingDayRequest
	if !decode(w, r, &input) {
		return
	}
	if _, err := time.Parse("2006-01-02", input.Date); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_date", "date must use YYYY-MM-DD")
		return
	}
	input.Reason = clean(input.Reason)
	if input.Reason == "" {
		writeError(w, http.StatusBadRequest, "invalid_reason", "reason is required")
		return
	}
	var id int64
	err := a.db.QueryRow(r.Context(), `insert into no_programming_days (date, ficha_id, instructor_id, reason) values ($1, $2, $3, $4) returning id`, input.Date, input.FichaID, input.InstructorID, input.Reason).Scan(&id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "insert_failed", err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"id": id})
}

func validateSchedule(input scheduleRequest) error {
	if _, err := time.Parse("2006-01-02", input.Date); err != nil {
		return fmt.Errorf("date must use YYYY-MM-DD")
	}
	start, err := time.Parse("15:04", input.StartTime)
	if err != nil {
		return fmt.Errorf("startTime must use HH:MM")
	}
	end, err := time.Parse("15:04", input.EndTime)
	if err != nil {
		return fmt.Errorf("endTime must use HH:MM")
	}
	if !end.After(start) {
		return fmt.Errorf("endTime must be after startTime")
	}
	if input.InstructorID <= 0 || input.FichaID <= 0 || input.EnvironmentID <= 0 {
		return fmt.Errorf("instructorId, fichaId and environmentId are required")
	}
	return nil
}

func decode(w http.ResponseWriter, r *http.Request, dst any) bool {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", err.Error())
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("encode response failed: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, map[string]string{"error": code, "message": message})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", getenv("CORS_ORIGIN", "http://localhost:5173"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getenv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func clean(value string) string {
	return strings.TrimSpace(value)
}

func nullIfEmpty(value string) any {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	return value
}
