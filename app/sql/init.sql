CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nim BIGINT NOT NULL,
    mata_kuliah VARCHAR(255) NOT NULL,
    jurusan VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Membuat function untuk auto-update kolom updated_at
CREATE OR REPLACE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Membuat trigger untuk menjalankan function saat data diperbarui
CREATE TRIGGER trigger_set_timestamp
BEFORE UPDATE ON students
FOR EACH ROW
EXECUTE FUNCTION set_timestamp();
