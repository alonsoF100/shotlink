-- +goose Up
-- +goose StatementBegin
CREATE TABLE short_urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    original_url TEXT NOT NULL,
    short_code VARCHAR(15) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    click_count BIGINT DEFAULT 0,
    is_active BOOLEAN DEFAULT true
);

CREATE INDEX idx_short_urls_short_code ON short_urls(short_code);
CREATE INDEX idx_short_urls_original_url ON short_urls(original_url);
CREATE INDEX idx_short_urls_created_at ON short_urls(created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE short_urls;
-- +goose StatementEnd
