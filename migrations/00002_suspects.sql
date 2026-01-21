-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS suspects (
    suspectId BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    occupation VARCHAR(255) NOT NULL,
    backstory TEXT NOT NULL,
    possibleMotive TEXT NOT NULL,
    relationshipToVictim VARCHAR(255) NOT NULL,
    alibi TEXT NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE suspects;
-- +goose StatementEnd