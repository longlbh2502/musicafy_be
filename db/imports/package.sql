INSERT INTO packages (title, code, description, thumb) VALUES ('Free', 'FREE-TRIAL', 'The free music package is a service that allows users to access and enjoy millions of their favorite songs on a platform without any cost. With the free package, users can search for and listen to music through playlists, genres, or favorite artists. However, some features may be limited, such as the inability to download music to devices, lower audio quality, or interruptions by ads. Despite these limitations, it remains a great option for those who want to explore music without committing to a paid subscription.', 'https://f005.backblazeb2.com/file/Musicafy-b2/common/packages/premium-quality.png');

INSERT INTO package_price (title, code, recommend, package, price, description, duration) VALUES ('Trải nghiệm', 'MONTHLY', true, 'FREE-TRIAL', 0, 'Free', 2592000);

INSERT INTO packages (title, code, description, thumb) VALUES ('Plus', 'PLUS', 'Không quảng cáo: Người dùng có thể nghe nhạc mà không bị gián đoạn bởi quảng cáo.
Chất lượng âm thanh cao: Tận hưởng âm thanh với chất lượng tốt hơn so với gói miễn phí.
Chế độ nghe ngoại tuyến: Tải nhạc về thiết bị và nghe mà không cần kết nối internet.
Trải nghiệm mượt mà hơn: Không bị gián đoạn và có thể chuyển bài hát nhanh chóng.', 'https://via.placeholder.com/150');

INSERT INTO package_price (title, code, recommend, package, price, description, duration) VALUES ('1 tuần', 'WEEKLY', false, 'PLUS', 1, '1 week', 2592000);
INSERT INTO package_price (title, code, recommend, package, price, description, duration) VALUES ('1 tháng', 'MONTHLY', true, 'PLUS', 3, '1 month', 2592000);
INSERT INTO package_price (title, code, recommend, package, price, description, duration) VALUES ('1 năm', 'YEARLY', false, 'PLUS', 10, '1 year', 2592000);

INSERT INTO packages (title, code, description, thumb) VALUES ('PREMIUM', 'PREMIUM', 'Tất cả tính năng của gói Plus: Bao gồm không quảng cáo, chất lượng âm thanh cao và nghe ngoại tuyến.
Tính năng nâng cao: Có thể truy cập các tính năng đặc biệt như nghe nhạc với chất lượng cao nhất, tạo playlist không giới hạn, và điều khiển âm thanh nâng cao.
Dùng trên nhiều thiết bị: Tận hưởng trải nghiệm âm nhạc trên nhiều thiết bị khác nhau cùng lúc.
Tùy chỉnh cá nhân hóa cao: Được đề xuất các bài hát, album, và nghệ sĩ phù hợp với sở thích âm nhạc của người dùng.', 'https://via.placeholder.com/150');

INSERT INTO package_price (title, code, recommend, package, price, description, duration) VALUES ('1 tuần', 'WEEKLY', false, 'PREMIUM', 2, '1 week', 2592000);
INSERT INTO package_price (title, code, recommend, package, price, description, duration) VALUES ('1 tháng', 'MONTHLY', false, 'PREMIUM', 6, '1 month', 2592000);
INSERT INTO package_price (title, code, recommend, package, price, description, duration) VALUES ('1 năm', 'YEARLY', true, 'PREMIUM', 20, '1 year', 2592000);
