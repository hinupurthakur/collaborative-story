CREATE TRIGGER sentences_updated_at_modtime BEFORE UPDATE ON sentences FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
