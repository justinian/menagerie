package main

var databaseSchema = []string{`
	CREATE TABLE worlds (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE,
		iter INTEGER DEFAULT 0
	);`,

	`CREATE TABLE classes (
		id INTEGER PRIMARY KEY,
		class TEXT,
		name TEXT
	);`,

	`CREATE TABLE dinos (
		id INTEGER,
		list INTEGER,
		world INTEGER,
		class INTEGER,
		is_tamed BOOLEAN,
		name TEXT,
		level_wild INTEGER,
		level_tamed INTEGER,
		dino_id1 INTEGER,
		dino_id2 INTEGER,
		is_cryo BOOLEAN,
		parent_class INTEGER,
		parent_name TEXT,
		x FLOAT,
		y FLOAT,
		z FLOAT,

		color0 INTEGER,
		color1 INTEGER,
		color2 INTEGER,
		color3 INTEGER,
		color4 INTEGER,
		color5 INTEGER,

		health_current FLOAT,
		stamina_current FLOAT,
		torpor_current FLOAT,
		oxygen_current FLOAT,
		food_current FLOAT,
		weight_current FLOAT,
		melee_current FLOAT,
		speed_current FLOAT,

		health_wild INTEGER,
		stamina_wild INTEGER,
		torpor_wild INTEGER,
		oxygen_wild INTEGER,
		food_wild INTEGER,
		weight_wild INTEGER,
		melee_wild INTEGER,
		speed_wild INTEGER,

		health_tamed INTEGER,
		stamina_tamed INTEGER,
		torpor_tamed INTEGER,
		oxygen_tamed INTEGER,
		food_tamed INTEGER,
		weight_tamed INTEGER,
		melee_tamed INTEGER,
		speed_tamed INTEGER,

		level_total INTEGER AS (level_wild+level_tamed),

		health_total INTEGER AS (health_wild+health_tamed),
		stamina_total INTEGER AS (stamina_wild+stamina_tamed),
		torpor_total INTEGER AS (torpor_wild+torpor_tamed),
		oxygen_total INTEGER AS (oxygen_wild+oxygen_tamed),
		food_total INTEGER AS (food_wild+food_tamed),
		weight_total INTEGER AS (weight_wild+weight_tamed),
		melee_total INTEGER AS (melee_wild+melee_tamed),
		speed_total INTEGER AS (speed_wild+speed_tamed),

		PRIMARY KEY (id, list, world)
	);`,
}

const insertDino = `
	INSERT INTO dinos (

		id, list, world,
		class, name, is_tamed,
		level_wild, level_tamed,
		dino_id1, dino_id2,
		is_cryo, parent_class, parent_name,

		x, y, z,

		color0, color1, color2, color3, color4, color5,

		health_current, stamina_current, torpor_current, oxygen_current,
		food_current, weight_current, melee_current, speed_current,

		health_wild, stamina_wild, torpor_wild, oxygen_wild,
		food_wild, weight_wild, melee_wild, speed_wild,

		health_tamed, stamina_tamed, torpor_tamed, oxygen_tamed,
		food_tamed, weight_tamed, melee_tamed, speed_tamed
	)
	VALUES (
		?, ?, ?,
		?, ?, ?,
		?, ?,
		?, ?,
		?, ?, ?,

		?, ?, ?,

		?, ?, ?, ?, ?, ?,

		?, ?, ?, ?,
		?, ?, ?, ?,

		?, ?, ?, ?,
		?, ?, ?, ?,

		?, ?, ?, ?,
		?, ?, ?, ?
	)`
