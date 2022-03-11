-- provided the field is named the same thing in all tables that use this, you can use a centralized function
CREATE FUNCTION update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
  BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
  END;
$$;
