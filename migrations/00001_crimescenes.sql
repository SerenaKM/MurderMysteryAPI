-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS crime_scenes (
    crimeSceneId BIGSERIAL PRIMARY KEY,
    location TEXT NOT NULL,
    description TEXT NOT NULL, 
    keyObjects TEXT NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crime_scenes;
-- +goose StatementEnd