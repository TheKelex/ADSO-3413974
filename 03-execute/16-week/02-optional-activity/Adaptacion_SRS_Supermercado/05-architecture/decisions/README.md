# Architecture Decision Records (ADR)

## Convenciones

- Formato: `ADR-NNN-titulo-corto.md`
- Los ADRs **nunca se eliminan ni se mueven** — permanecen en `records/` con `status: DEPRECATED`
- Cuando una ADR queda obsoleta: cambiar su estado a `DEPRECATED` y agregar al inicio del archivo una línea como `> Reemplazada por: ADR-NNN-nueva.md`
- `99-archive/old-decisions/` se usa solo para documentos de decisión que existían antes de adoptar el formato ADR
- Numeración secuencial desde `ADR-001`

## Template

Usar el archivo [`_template-adr.md`](./_template-adr.md):

```bash
# Reemplazar NNN con el número siguiente y el título con kebab-case
cp 05-architecture/decisions/_template-adr.md \
   05-architecture/decisions/records/ADR-NNN-titulo-corto.md
```

## Registro

| ADR | Título | Estado |
|-----|--------|--------|
