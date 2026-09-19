# PRD - Automotive Technical Support Management System

## Business context

An automotive service center receives vehicles with reported failures and today
tracks the whole repair cycle on paper notes and isolated spreadsheets. Work is
lost between the reception desk, the technician and the customer: nobody can say
where a vehicle is, who is working on it, what was actually repaired, or whether
a part is still under warranty. The MVP is a web system that digitizes the core
operating domain of the workshop: customer and vehicle check-in, technical
diagnostic, technician assignment, live status tracking of the service order,
recorded interventions, warranty control, and a unified clinical timeline per
vehicle.

Billing, tax invoicing and online payment stay out of the product: the workshop
charges at its physical cashier and the corporate ERP issues legal invoices.

## Users

- **Administrator (workshop manager)**: the operational owner. Signs in with
  username and password, registers customers and vehicles, creates service
  orders at check-in, assigns technicians, controls warranties and reviews the
  clinical timeline of any vehicle. The system ships with a seeded administrator
  (bootstrap) so the first deployment is usable immediately.
- **Technician (mechanic)**: signs in, sees the service orders assigned to them,
  records the technical diagnostic and registers each intervention performed on
  the vehicle, advancing the order status.
- **Customer**: consulted through the administrator in this MVP. Customer data
  and the vehicle timeline are stored so the workshop can inform the customer of
  the real state of their vehicle.

## MVP scope

Included:
- Authentication with username and password issuing a session token, and two
  roles: administrator and technician.
- Customer management: create and list customers with contact data.
- Vehicle management: register a vehicle against a customer with plate, VIN,
  brand, model and year. Plate and VIN are unique.
- Service order check-in: create a service order for a vehicle with the failure
  reported by the customer and the reception date.
- Technical diagnostic: a technician records the diagnostic of a service order
  with the finding and the components to repair.
- Technician assignment: the administrator assigns an available technician to a
  service order; a technician cannot hold more than one active order at a time.
- Status tracking: the service order moves through RECEIVED, IN_DIAGNOSIS,
  IN_REPAIR, READY and DELIVERED, and every transition is stored with its
  timestamp so the history is auditable.
- Interventions: the technician records each physical action executed on the
  vehicle, with description, labor hour count and the parts used.
- Warranty management: a warranty is issued over an intervention, with kind
  (LABOR or PART), coverage in months and an expiration date derived from the
  issue date; the system reports whether a warranty is valid or expired at the
  consulted date.
- Clinical timeline: a chronological view of everything that happened to a
  vehicle - orders, diagnostics, interventions and warranties - filterable by
  vehicle.

Excluded from this version: predictive maintenance analytics, the corporate ERP
inventory integration, the SMS/WhatsApp notification provider, online payments,
tax invoicing, supplier purchasing and corporate SSO. These are deliberate
deferrals, not omissions: their external dependencies are not available yet.

## Functional requirements

1. A user signs in with valid credentials and receives a session token; invalid
   credentials are rejected with a clear error and no token is issued.
2. The administrator creates a customer with name, document, phone and email,
   and lists the registered customers.
3. The administrator registers a vehicle for an existing customer with plate,
   VIN, brand, model and year; a duplicate plate or VIN is rejected.
4. The administrator creates a service order for a vehicle at check-in with the
   reported failure; the order starts in RECEIVED status.
5. The administrator assigns a technician to a service order; the assignment is
   rejected when the technician already holds an active service order.
6. The assigned technician records the diagnostic of the service order with the
   finding and the components to repair; recording the diagnostic moves the
   order to IN_DIAGNOSIS.
7. The assigned technician registers an intervention on a service order with
   description, labor hours and parts used; the first intervention moves the
   order to IN_REPAIR.
8. A service order advances to READY and then to DELIVERED; a transition that
   does not follow the defined lifecycle is rejected with an explicit error.
9. The administrator issues a warranty over an intervention with kind and
   coverage in months; the expiration date is computed from the issue date.
10. The system reports, for any warranty, whether it is valid or expired at the
    consulted date.
11. The administrator consults the clinical timeline of a vehicle and sees, in
    chronological order, its service orders, diagnostics, interventions and
    warranties.
12. The administrator sees a dashboard with the number of open service orders,
    the count of orders per status and the technicians currently occupied.

## Non-functional requirements

- **security**: every write route requires a valid token; a technician can only
  write on the service orders assigned to them; passwords are stored hashed,
  never in plain text.
- **localization**: all technical identifiers - tables, columns, endpoints,
  code symbols, files and documentation - are written in English and singular.
  Every string shown to the end user in the interface is written in Spanish,
  including labels, buttons, validation messages and error messages.
- **performance**: catalog and timeline read responses answer in under 500 ms
  for a workshop with up to 50,000 vehicle records per year.
- **reliability**: a service order status transition is atomic; a rejected
  transition leaves the order exactly as it was.
- **maintainability**: backend, frontend and database are separate projects with
  no shared build; the domain rules live in the backend, not in the interface.
- **usability**: the interface is operable on a tablet, which is the device the
  technicians use on the workshop floor.

## Acceptance criteria

- A user can sign in, register a customer and a vehicle, open a service order,
  assign a technician, record a diagnostic, register an intervention, issue a
  warranty and see all of it reflected in the vehicle clinical timeline.
- A warranty issued with a coverage of N months reports as valid before its
  expiration date and expired after it.
- An out-of-lifecycle status transition is rejected and the order keeps its
  previous status.
- The deployment starts with `docker compose` and includes backend, frontend and
  database with a seeded administrator.
- There is evidence of automated backend and frontend tests.

## Stack

- Backend: Go
- Frontend: React
- Database: MySQL
- Operation: Docker Compose
