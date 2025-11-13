-- +goose Up
-- +goose StatementBegin
CREATE TABLE clicks (
    id BIGSERIAL PRIMARY KEY,
    short_code VARCHAR(10) NOT NULL,
    clicked_at TIMESTAMPTZ DEFAULT NOW(),
    user_agent TEXT,
    ip_address INET,
    referrer TEXT,
    country_code VARCHAR(2),
    device_type VARCHAR(20)
);

CREATE INDEX idx_clicks_short_code ON clicks(short_code);
CREATE INDEX idx_clicks_clicked_at ON clicks(clicked_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clicks;
-- +goose StatementEnd