CREATE TRIGGER stories_updated_at_modtime BEFORE UPDATE ON stories FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
