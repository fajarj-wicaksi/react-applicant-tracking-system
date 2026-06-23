-- 1. Create a dummy tenant
INSERT INTO tenants (id, name, domain, created_at, updated_at) 
VALUES ('00000000-0000-0000-0000-000000000001', 'Default Tenant', 'default.talentflow.com', NOW(), NOW())
ON CONFLICT DO NOTHING;

-- 2. Create a dummy role
INSERT INTO roles (id, name, description, created_at, updated_at) 
VALUES ('11111111-1111-1111-1111-111111111111', 'Admin', 'Super Administrator', NOW(), NOW())
ON CONFLICT DO NOTHING;

-- 3. Create a dummy user
INSERT INTO users (id, tenant_id, role_id, email, password, first_name, last_name, is_active, created_at, updated_at)
VALUES (
    '22222222-2222-2222-2222-222222222222', 
    '00000000-0000-0000-0000-000000000001',
    '11111111-1111-1111-1111-111111111111', 
    'admin@talentflow.com', 
    '$2a$10$xyz123abc456def789ghijklmnoPQRSTUVWXYZ0123456789ABCD', 
    'Super', 
    'Admin', 
    true,
    NOW(),
    NOW()
)
ON CONFLICT DO NOTHING;
