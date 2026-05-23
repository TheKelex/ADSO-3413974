-- Seed parametrico y volumetrico para demo local del MVP.
-- No contiene codigos oficiales SENA. Todo identificador institucional queda marcado como Pendiente/Demo.

insert into training_programs (name, level) values
  ('Programa de formacion Pendiente Demo - Software', 'Tecnologo Pendiente'),
  ('Programa de formacion Pendiente Demo - Gestion', 'Tecnico Pendiente'),
  ('Programa de formacion Pendiente Demo - Redes', 'Tecnologo Pendiente')
on conflict (name) do nothing;

insert into instructors (name, instructor_type) values
  ('Instructor Demo Planta 1', 'planta'),
  ('Instructor Demo Planta 2', 'planta'),
  ('Instructor Demo Contratista 1', 'contratista'),
  ('Instructor Demo Contratista 2', 'contratista'),
  ('Instructor Demo Contratista 3', 'contratista')
on conflict (name) do nothing;

insert into training_environments (name) values
  ('Ambiente Demo 101 - Pendiente'),
  ('Ambiente Demo 202 - Pendiente'),
  ('Ambiente Demo Laboratorio - Pendiente'),
  ('Ambiente Demo Virtual - Pendiente')
on conflict (name) do nothing;

insert into fichas (code, program_id)
select 'Ficha Demo 001 - Pendiente', id from training_programs where name = 'Programa de formacion Pendiente Demo - Software'
on conflict (code) do nothing;

insert into fichas (code, program_id)
select 'Ficha Demo 002 - Pendiente', id from training_programs where name = 'Programa de formacion Pendiente Demo - Gestion'
on conflict (code) do nothing;

insert into fichas (code, program_id)
select 'Ficha Demo 003 - Pendiente', id from training_programs where name = 'Programa de formacion Pendiente Demo - Redes'
on conflict (code) do nothing;

insert into fichas (code, program_id)
select 'Ficha Demo 004 - Pendiente', id from training_programs where name = 'Programa de formacion Pendiente Demo - Software'
on conflict (code) do nothing;

with seed_rows as (
  select * from (values
    ('2026-05-18'::date, '07:00'::time, '09:00'::time, 'Instructor Demo Planta 1', 'Ficha Demo 001 - Pendiente', 'Ambiente Demo 101 - Pendiente', 'Programa de formacion Pendiente Demo - Software', 'Competencia Pendiente Demo 1', 'RAP Pendiente Demo 1', 'Actividad de aprendizaje Pendiente Demo 1', 'Inicio de bloque formativo demo'),
    ('2026-05-18'::date, '09:00'::time, '11:00'::time, 'Instructor Demo Planta 1', 'Ficha Demo 002 - Pendiente', 'Ambiente Demo 202 - Pendiente', 'Programa de formacion Pendiente Demo - Gestion', 'Competencia Pendiente Demo 2', 'RAP Pendiente Demo 2', 'Actividad de aprendizaje Pendiente Demo 2', 'Cambio de ambiente pendiente de confirmar'),
    ('2026-05-19'::date, '07:00'::time, '10:00'::time, 'Instructor Demo Planta 1', 'Ficha Demo 001 - Pendiente', 'Ambiente Demo Laboratorio - Pendiente', 'Programa de formacion Pendiente Demo - Software', 'Competencia Pendiente Demo 1', 'RAP Pendiente Demo 3', 'Actividad de aprendizaje Pendiente Demo 3', 'Practica demo'),
    ('2026-05-20'::date, '13:00'::time, '16:00'::time, 'Instructor Demo Planta 1', 'Ficha Demo 004 - Pendiente', 'Ambiente Demo Virtual - Pendiente', 'Programa de formacion Pendiente Demo - Software', 'Competencia Pendiente Demo 4', 'RAP Pendiente Demo 4', 'Actividad de aprendizaje Pendiente Demo 4', 'Sesion virtual demo'),
    ('2026-05-18'::date, '07:00'::time, '09:00'::time, 'Instructor Demo Planta 2', 'Ficha Demo 003 - Pendiente', 'Ambiente Demo Laboratorio - Pendiente', 'Programa de formacion Pendiente Demo - Redes', 'Competencia Pendiente Demo 5', 'RAP Pendiente Demo 5', 'Actividad de aprendizaje Pendiente Demo 5', 'Sin novedad'),
    ('2026-05-19'::date, '10:00'::time, '12:00'::time, 'Instructor Demo Planta 2', 'Ficha Demo 002 - Pendiente', 'Ambiente Demo 101 - Pendiente', 'Programa de formacion Pendiente Demo - Gestion', 'Competencia Pendiente Demo 2', 'RAP Pendiente Demo 6', 'Actividad de aprendizaje Pendiente Demo 6', 'Observacion demo'),
    ('2026-05-21'::date, '08:00'::time, '11:00'::time, 'Instructor Demo Contratista 1', 'Ficha Demo 001 - Pendiente', 'Ambiente Demo 202 - Pendiente', 'Programa de formacion Pendiente Demo - Software', 'Competencia Pendiente Demo 7', 'RAP Pendiente Demo 7', 'Actividad de aprendizaje Pendiente Demo 7', 'Pendiente validar RAP'),
    ('2026-05-22'::date, '14:00'::time, '17:00'::time, 'Instructor Demo Contratista 1', 'Ficha Demo 003 - Pendiente', 'Ambiente Demo Virtual - Pendiente', 'Programa de formacion Pendiente Demo - Redes', 'Competencia Pendiente Demo 8', 'RAP Pendiente Demo 8', 'Actividad de aprendizaje Pendiente Demo 8', 'Trabajo remoto demo'),
    ('2026-05-18'::date, '11:00'::time, '13:00'::time, 'Instructor Demo Contratista 2', 'Ficha Demo 004 - Pendiente', 'Ambiente Demo Laboratorio - Pendiente', 'Programa de formacion Pendiente Demo - Software', 'Competencia Pendiente Demo 9', 'RAP Pendiente Demo 9', 'Actividad de aprendizaje Pendiente Demo 9', 'Sin novedad'),
    ('2026-05-20'::date, '07:00'::time, '09:00'::time, 'Instructor Demo Contratista 2', 'Ficha Demo 002 - Pendiente', 'Ambiente Demo 101 - Pendiente', 'Programa de formacion Pendiente Demo - Gestion', 'Competencia Pendiente Demo 10', 'RAP Pendiente Demo 10', 'Actividad de aprendizaje Pendiente Demo 10', 'Novedad pendiente'),
    ('2026-05-21'::date, '13:00'::time, '15:00'::time, 'Instructor Demo Contratista 3', 'Ficha Demo 003 - Pendiente', 'Ambiente Demo 202 - Pendiente', 'Programa de formacion Pendiente Demo - Redes', 'Competencia Pendiente Demo 11', 'RAP Pendiente Demo 11', 'Actividad de aprendizaje Pendiente Demo 11', 'Sin novedad'),
    ('2026-05-22'::date, '09:00'::time, '12:00'::time, 'Instructor Demo Contratista 3', 'Ficha Demo 004 - Pendiente', 'Ambiente Demo 101 - Pendiente', 'Programa de formacion Pendiente Demo - Software', 'Competencia Pendiente Demo 12', 'RAP Pendiente Demo 12', 'Actividad de aprendizaje Pendiente Demo 12', 'Cierre semanal demo')
  ) as rows(date, start_time, end_time, instructor_name, ficha_code, environment_name, program_name, competence, rap, learning_activity, observation)
)
insert into schedule_entries (
  date, start_time, end_time, instructor_id, ficha_id, environment_id, program_id,
  competence, rap, learning_activity, observation
)
select
  sr.date,
  sr.start_time,
  sr.end_time,
  i.id,
  f.id,
  e.id,
  p.id,
  sr.competence,
  sr.rap,
  sr.learning_activity,
  sr.observation
from seed_rows sr
join instructors i on i.name = sr.instructor_name
join fichas f on f.code = sr.ficha_code
join training_environments e on e.name = sr.environment_name
join training_programs p on p.name = sr.program_name
where not exists (
  select 1
  from schedule_entries s
  where s.date = sr.date
    and s.start_time = sr.start_time
    and s.end_time = sr.end_time
    and s.instructor_id = i.id
    and s.ficha_id = f.id
    and s.environment_id = e.id
);

insert into no_programming_days (date, ficha_id, reason)
select '2026-05-20'::date, f.id, 'Dia sin programacion demo por ajuste de planeacion'
from fichas f
where f.code = 'Ficha Demo 003 - Pendiente'
  and not exists (
    select 1 from no_programming_days n
    where n.date = '2026-05-20'::date
      and n.ficha_id = f.id
      and n.reason = 'Dia sin programacion demo por ajuste de planeacion'
  );

insert into no_programming_days (date, instructor_id, reason)
select '2026-05-22'::date, i.id, 'Disponibilidad de instructor pendiente de confirmar'
from instructors i
where i.name = 'Instructor Demo Planta 2'
  and not exists (
    select 1 from no_programming_days n
    where n.date = '2026-05-22'::date
      and n.instructor_id = i.id
      and n.reason = 'Disponibilidad de instructor pendiente de confirmar'
  );
