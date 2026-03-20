DROP FUNCTION IF EXISTS set_updated_at;
DROP TRIGGER IF EXISTS trigger_update_customers ON customers;
DROP TRIGGER IF EXISTS trigger_update_purchases ON purchases;
DROP TRIGGER IF EXISTS trigger_update_products ON products;
DROP TRIGGER IF EXISTS trigger_update_product_purchase ON product_purchase;
