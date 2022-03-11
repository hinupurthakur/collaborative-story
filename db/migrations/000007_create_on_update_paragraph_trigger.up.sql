CREATE TRIGGER paragraphs_updated_at_modtime BEFORE UPDATE ON paragraphs FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
