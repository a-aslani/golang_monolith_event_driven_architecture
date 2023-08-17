CREATE TABLE IF NOT EXISTS tournaments (
    id TEXT NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL ,
    gamer_1_id TEXT NOT NULL,
    gamer_2_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP NULL ,
    PRIMARY KEY (id),
    CONSTRAINT fk_tournament_gamer1 FOREIGN KEY (gamer_1_id) REFERENCES gamers(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_tournament_gamer2 FOREIGN KEY (gamer_2_id) REFERENCES gamers(id) ON UPDATE CASCADE ON DELETE CASCADE
);