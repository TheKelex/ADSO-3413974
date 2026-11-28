CREATE DATABASE IF NOT EXISTS sigres_p2
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_0900_ai_ci;
USE sigres_p2;

CREATE TABLE rol (
    id_rol INT UNSIGNED NOT NULL AUTO_INCREMENT,
    nombre VARCHAR(40) NOT NULL,
    descripcion VARCHAR(200) NOT NULL,
    PRIMARY KEY (id_rol),
    CONSTRAINT uq_rol_nombre UNIQUE (nombre)
) ENGINE=InnoDB;

CREATE TABLE permiso (
    id_permiso INT UNSIGNED NOT NULL AUTO_INCREMENT,
    codigo VARCHAR(60) NOT NULL,
    descripcion VARCHAR(200) NOT NULL,
    PRIMARY KEY (id_permiso),
    CONSTRAINT uq_permiso_codigo UNIQUE (codigo)
) ENGINE=InnoDB;

CREATE TABLE usuario (
    id_usuario INT UNSIGNED NOT NULL AUTO_INCREMENT,
    id_rol INT UNSIGNED NOT NULL,
    nombre_cifrado VARBINARY(512) NOT NULL,
    apellido_cifrado VARBINARY(512) NOT NULL,
    identificacion_cifrada VARBINARY(512) NOT NULL,
    identificacion_hmac BINARY(32) NOT NULL,
    nombre_usuario VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    telefono_cifrado VARBINARY(512) NULL,
    activo TINYINT(1) NOT NULL DEFAULT 1,
    PRIMARY KEY (id_usuario),
    CONSTRAINT uq_usuario_identificacion_hmac UNIQUE (identificacion_hmac),
    CONSTRAINT uq_usuario_nombre_usuario UNIQUE (nombre_usuario),
    CONSTRAINT chk_usuario_activo CHECK (activo IN (0, 1)),
    CONSTRAINT fk_usuario_rol FOREIGN KEY (id_rol)
        REFERENCES rol (id_rol)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE rol_permiso (
    id_rol INT UNSIGNED NOT NULL,
    id_permiso INT UNSIGNED NOT NULL,
    PRIMARY KEY (id_rol, id_permiso),
    CONSTRAINT fk_rol_permiso_rol FOREIGN KEY (id_rol)
        REFERENCES rol (id_rol)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso)
        REFERENCES permiso (id_permiso)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE mesa (
    id_mesa INT UNSIGNED NOT NULL AUTO_INCREMENT,
    numero SMALLINT UNSIGNED NOT NULL,
    capacidad TINYINT UNSIGNED NOT NULL,
    estado VARCHAR(20) NOT NULL,
    PRIMARY KEY (id_mesa),
    CONSTRAINT uq_mesa_numero UNIQUE (numero),
    CONSTRAINT chk_mesa_capacidad CHECK (capacidad > 0),
    CONSTRAINT chk_mesa_estado CHECK (estado IN ('Disponible', 'Ocupada', 'Reservada'))
) ENGINE=InnoDB;

CREATE TABLE categoria_menu (
    id_categoria INT UNSIGNED NOT NULL AUTO_INCREMENT,
    nombre VARCHAR(80) NOT NULL,
    descripcion VARCHAR(250) NULL,
    activa TINYINT(1) NOT NULL DEFAULT 1,
    PRIMARY KEY (id_categoria),
    CONSTRAINT uq_categoria_menu_nombre UNIQUE (nombre),
    CONSTRAINT chk_categoria_menu_activa CHECK (activa IN (0, 1))
) ENGINE=InnoDB;

CREATE TABLE producto_menu (
    id_producto INT UNSIGNED NOT NULL AUTO_INCREMENT,
    id_categoria INT UNSIGNED NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    descripcion VARCHAR(500) NOT NULL,
    precio DECIMAL(12,2) NOT NULL,
    url_imagen VARCHAR(500) NOT NULL,
    disponible TINYINT(1) NOT NULL DEFAULT 1,
    PRIMARY KEY (id_producto),
    CONSTRAINT uq_producto_menu_categoria_nombre UNIQUE (id_categoria, nombre),
    CONSTRAINT chk_producto_menu_precio CHECK (precio >= 0),
    CONSTRAINT chk_producto_menu_disponible CHECK (disponible IN (0, 1)),
    CONSTRAINT fk_producto_menu_categoria FOREIGN KEY (id_categoria)
        REFERENCES categoria_menu (id_categoria)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE pedido (
    id_pedido INT UNSIGNED NOT NULL AUTO_INCREMENT,
    id_usuario INT UNSIGNED NOT NULL,
    id_mesa INT UNSIGNED NOT NULL,
    fecha_hora DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    numero_comensales TINYINT UNSIGNED NOT NULL,
    estado VARCHAR(20) NOT NULL,
    PRIMARY KEY (id_pedido),
    CONSTRAINT chk_pedido_comensales CHECK (numero_comensales > 0),
    CONSTRAINT chk_pedido_estado CHECK (
        estado IN ('Pendiente', 'EnPreparacion', 'Listo', 'Entregado', 'Cancelado')
    ),
    CONSTRAINT fk_pedido_usuario FOREIGN KEY (id_usuario)
        REFERENCES usuario (id_usuario)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_pedido_mesa FOREIGN KEY (id_mesa)
        REFERENCES mesa (id_mesa)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE detalle_pedido (
    id_pedido INT UNSIGNED NOT NULL,
    numero_linea SMALLINT UNSIGNED NOT NULL,
    id_producto INT UNSIGNED NOT NULL,
    cantidad SMALLINT UNSIGNED NOT NULL,
    precio_unitario DECIMAL(12,2) NOT NULL,
    nota VARCHAR(500) NULL,
    PRIMARY KEY (id_pedido, numero_linea),
    CONSTRAINT chk_detalle_pedido_linea CHECK (numero_linea > 0),
    CONSTRAINT chk_detalle_pedido_cantidad CHECK (cantidad > 0),
    CONSTRAINT chk_detalle_pedido_precio CHECK (precio_unitario >= 0),
    CONSTRAINT fk_detalle_pedido_pedido FOREIGN KEY (id_pedido)
        REFERENCES pedido (id_pedido)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_detalle_pedido_producto FOREIGN KEY (id_producto)
        REFERENCES producto_menu (id_producto)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE factura (
    id_factura INT UNSIGNED NOT NULL AUTO_INCREMENT,
    numero VARCHAR(30) NOT NULL,
    id_pedido INT UNSIGNED NOT NULL,
    id_cajero INT UNSIGNED NOT NULL,
    fecha_hora DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    subtotal DECIMAL(12,2) NOT NULL,
    impuestos DECIMAL(12,2) NOT NULL,
    total DECIMAL(12,2) NOT NULL,
    estado VARCHAR(15) NOT NULL DEFAULT 'Pendiente',
    motivo_anulacion TEXT NULL,
    anulada_por_usuario_id INT UNSIGNED NULL,
    anulada_en DATETIME NULL,
    PRIMARY KEY (id_factura),
    CONSTRAINT uq_factura_numero UNIQUE (numero),
    CONSTRAINT uq_factura_pedido UNIQUE (id_pedido),
    CONSTRAINT chk_factura_subtotal CHECK (subtotal >= 0),
    CONSTRAINT chk_factura_impuestos CHECK (impuestos >= 0),
    CONSTRAINT chk_factura_total CHECK (total = subtotal + impuestos),
    CONSTRAINT chk_factura_estado CHECK (estado IN ('Pendiente', 'Pagada', 'Anulada')),
    CONSTRAINT chk_factura_anulacion CHECK (
        (estado = 'Anulada'
            AND motivo_anulacion IS NOT NULL
            AND CHAR_LENGTH(TRIM(motivo_anulacion)) > 0
            AND anulada_por_usuario_id IS NOT NULL
            AND anulada_en IS NOT NULL)
        OR
        (estado <> 'Anulada'
            AND motivo_anulacion IS NULL
            AND anulada_por_usuario_id IS NULL
            AND anulada_en IS NULL)
    ),
    CONSTRAINT fk_factura_pedido FOREIGN KEY (id_pedido)
        REFERENCES pedido (id_pedido)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_factura_cajero FOREIGN KEY (id_cajero)
        REFERENCES usuario (id_usuario)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_factura_anulador FOREIGN KEY (anulada_por_usuario_id)
        REFERENCES usuario (id_usuario)
        ON UPDATE RESTRICT ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE metodo_pago (
    id_metodo_pago INT UNSIGNED NOT NULL AUTO_INCREMENT,
    codigo VARCHAR(30) NOT NULL,
    nombre VARCHAR(60) NOT NULL,
    activo TINYINT(1) NOT NULL DEFAULT 1,
    PRIMARY KEY (id_metodo_pago),
    CONSTRAINT uq_metodo_pago_codigo UNIQUE (codigo),
    CONSTRAINT uq_metodo_pago_nombre UNIQUE (nombre),
    CONSTRAINT chk_metodo_pago_codigo CHECK (
        codigo IN ('EFECTIVO', 'TARJETA', 'TRANSFERENCIA', 'BILLETERA_DIGITAL')
    ),
    CONSTRAINT chk_metodo_pago_activo CHECK (activo IN (0, 1))
) ENGINE=InnoDB;

CREATE TABLE pago (
    id_pago INT UNSIGNED NOT NULL AUTO_INCREMENT,
    id_factura INT UNSIGNED NOT NULL,
    id_metodo_pago INT UNSIGNED NOT NULL,
    fecha_hora DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    monto DECIMAL(12,2) NOT NULL,
    propina DECIMAL(12,2) NOT NULL DEFAULT 0,
    PRIMARY KEY (id_pago),
    CONSTRAINT uq_pago_factura UNIQUE (id_factura),
    CONSTRAINT chk_pago_monto CHECK (monto > 0),
    CONSTRAINT chk_pago_propina CHECK (propina >= 0),
    CONSTRAINT fk_pago_factura FOREIGN KEY (id_factura)
        REFERENCES factura (id_factura)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_pago_metodo FOREIGN KEY (id_metodo_pago)
        REFERENCES metodo_pago (id_metodo_pago)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE insumo (
    id_insumo INT UNSIGNED NOT NULL AUTO_INCREMENT,
    nombre VARCHAR(100) NOT NULL,
    categoria VARCHAR(80) NOT NULL,
    existencia DECIMAL(12,3) NOT NULL,
    unidad VARCHAR(20) NOT NULL,
    minimo DECIMAL(12,3) NOT NULL,
    costo_unitario DECIMAL(12,2) NOT NULL,
    proveedor VARCHAR(120) NULL,
    PRIMARY KEY (id_insumo),
    CONSTRAINT uq_insumo_nombre UNIQUE (nombre),
    CONSTRAINT chk_insumo_existencia CHECK (existencia >= 0),
    CONSTRAINT chk_insumo_minimo CHECK (minimo >= 0),
    CONSTRAINT chk_insumo_costo CHECK (costo_unitario >= 0)
) ENGINE=InnoDB;

CREATE TABLE producto_insumo (
    id_producto INT UNSIGNED NOT NULL,
    id_insumo INT UNSIGNED NOT NULL,
    cantidad_por_producto DECIMAL(12,3) NOT NULL,
    PRIMARY KEY (id_producto, id_insumo),
    CONSTRAINT chk_producto_insumo_cantidad CHECK (cantidad_por_producto > 0),
    CONSTRAINT fk_producto_insumo_producto FOREIGN KEY (id_producto)
        REFERENCES producto_menu (id_producto)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_producto_insumo_insumo FOREIGN KEY (id_insumo)
        REFERENCES insumo (id_insumo)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE movimiento_inventario (
    id_movimiento INT UNSIGNED NOT NULL AUTO_INCREMENT,
    id_insumo INT UNSIGNED NOT NULL,
    id_usuario INT UNSIGNED NOT NULL,
    id_pedido INT UNSIGNED NULL,
    numero_linea SMALLINT UNSIGNED NULL,
    fecha_hora DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    tipo VARCHAR(20) NOT NULL,
    cantidad DECIMAL(12,3) NOT NULL,
    existencia_resultante DECIMAL(12,3) NOT NULL,
    observacion VARCHAR(250) NULL,
    PRIMARY KEY (id_movimiento),
    CONSTRAINT chk_movimiento_referencia CHECK (
        (id_pedido IS NULL AND numero_linea IS NULL)
        OR (id_pedido IS NOT NULL AND numero_linea IS NOT NULL)
    ),
    CONSTRAINT chk_movimiento_tipo CHECK (
        tipo IN ('Entrada', 'Salida', 'AjusteEntrada', 'AjusteSalida')
    ),
    CONSTRAINT chk_movimiento_cantidad CHECK (cantidad > 0),
    CONSTRAINT chk_movimiento_existencia CHECK (existencia_resultante >= 0),
    CONSTRAINT fk_movimiento_insumo FOREIGN KEY (id_insumo)
        REFERENCES insumo (id_insumo)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_movimiento_usuario FOREIGN KEY (id_usuario)
        REFERENCES usuario (id_usuario)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_movimiento_detalle FOREIGN KEY (id_pedido, numero_linea)
        REFERENCES detalle_pedido (id_pedido, numero_linea)
        ON UPDATE RESTRICT ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE TABLE alerta_inventario (
    id_alerta INT UNSIGNED NOT NULL AUTO_INCREMENT,
    id_insumo INT UNSIGNED NOT NULL,
    atendida_por_usuario_id INT UNSIGNED NULL,
    fecha_hora DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    existencia_actual DECIMAL(12,3) NOT NULL,
    minimo_alerta DECIMAL(12,3) NOT NULL,
    estado VARCHAR(10) NOT NULL DEFAULT 'Activa',
    atendida_en DATETIME NULL,
    PRIMARY KEY (id_alerta),
    CONSTRAINT chk_alerta_existencia CHECK (existencia_actual >= 0),
    CONSTRAINT chk_alerta_minimo CHECK (minimo_alerta >= 0),
    CONSTRAINT chk_alerta_estado CHECK (estado IN ('Activa', 'Atendida')),
    CONSTRAINT chk_alerta_atencion CHECK (
        (estado = 'Activa'
            AND atendida_por_usuario_id IS NULL
            AND atendida_en IS NULL)
        OR
        (estado = 'Atendida'
            AND atendida_por_usuario_id IS NOT NULL
            AND atendida_en IS NOT NULL)
    ),
    CONSTRAINT fk_alerta_insumo FOREIGN KEY (id_insumo)
        REFERENCES insumo (id_insumo)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_alerta_usuario FOREIGN KEY (atendida_por_usuario_id)
        REFERENCES usuario (id_usuario)
        ON UPDATE RESTRICT ON DELETE RESTRICT
) ENGINE=InnoDB;

DELIMITER //
CREATE TRIGGER trg_factura_paga_inmutable
BEFORE UPDATE ON factura
FOR EACH ROW
BEGIN
    IF OLD.estado = 'Pagada' THEN
        IF NOT (
            NEW.estado = 'Anulada'
            AND NEW.id_factura <=> OLD.id_factura
            AND NEW.numero <=> OLD.numero
            AND NEW.id_pedido <=> OLD.id_pedido
            AND NEW.id_cajero <=> OLD.id_cajero
            AND NEW.fecha_hora <=> OLD.fecha_hora
            AND NEW.subtotal <=> OLD.subtotal
            AND NEW.impuestos <=> OLD.impuestos
            AND NEW.total <=> OLD.total
            AND NEW.motivo_anulacion IS NOT NULL
            AND CHAR_LENGTH(TRIM(NEW.motivo_anulacion)) > 0
            AND NEW.anulada_por_usuario_id IS NOT NULL
            AND NEW.anulada_en IS NOT NULL
        ) THEN
            SIGNAL SQLSTATE '45000'
                SET MESSAGE_TEXT = 'La factura pagada es inmutable; solo se permite anularla con motivo y responsable';
        END IF;
    ELSEIF OLD.estado = 'Anulada' THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Una factura anulada no puede modificarse';
    END IF;
END//
DELIMITER ;

CREATE VIEW vw_reporte_ventas AS
SELECT
    DATE(f.fecha_hora) AS fecha_venta,
    p.id_usuario AS id_mesero,
    dp.id_producto,
    pm.nombre AS producto,
    SUM(dp.cantidad) AS cantidad_vendida,
    SUM(dp.cantidad * dp.precio_unitario) AS total_vendido
FROM factura AS f
INNER JOIN pedido AS p ON p.id_pedido = f.id_pedido
INNER JOIN detalle_pedido AS dp ON dp.id_pedido = p.id_pedido
INNER JOIN producto_menu AS pm ON pm.id_producto = dp.id_producto
WHERE f.estado = 'Pagada'
GROUP BY DATE(f.fecha_hora), p.id_usuario, dp.id_producto, pm.nombre;
