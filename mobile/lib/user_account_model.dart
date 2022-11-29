class Product {
  final int first_name;
  final String last_name;
  final String email;
  final int password;

  const Product({
    required this.first_name,
    required this.last_name,
    required this.email,
    required this.password,
  });

  factory Product.fromJson(Map json) {
    return Product(
      first_name: json['firstname'],
      last_name: json['lastname'],
      email: json['email'],
      password: json['password']
    );
  }
}