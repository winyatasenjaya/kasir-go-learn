-- Migration untuk menambahkan category_id ke tabel products

-- Tambahkan kolom category_id ke tabel products
ALTER TABLE products 
ADD COLUMN category_id INTEGER;

-- Tambahkan foreign key constraint
ALTER TABLE products 
ADD CONSTRAINT fk_products_category 
FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL;

-- Optional: Buat index untuk meningkatkan performa JOIN
CREATE INDEX idx_products_category_id ON products(category_id);
