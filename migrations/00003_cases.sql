-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cases (
    caseId BIGSERIAL PRIMARY KEY,
    crimeSceneId BIGINT NOT NULL REFERENCES crime_scenes(crimeSceneId)
);

CREATE TABLE IF NOT EXISTS case_suspects (
    caseId BIGINT NOT NULL REFERENCES cases(caseId) ON DELETE CASCADE,
    suspectId BIGINT NOT NULL REFERENCES suspects(suspectId),
    PRIMARY KEY (caseId, suspectId)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cases;
DROP TABLE case_suspects;
-- +goose StatementEnd