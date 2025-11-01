export const validateEmail = (email: string): boolean => {
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return re.test(email);
};

export const validatePassword = (password: string): boolean => {
  // At least 8 characters, one uppercase, one lowercase, one number
  const re = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d@$!%*?&]{8,}$/;
  return re.test(password);
};

export const validatePhoneNumber = (phone: string): boolean => {
  // Basic phone number validation
  const re = /^\+?[\d\s\-\(\)]{10,}$/;
  return re.test(phone);
};

export const validateAccountNumber = (accountNumber: string): boolean => {
  // Basic account number validation (alphanumeric, 8-20 characters)
  const re = /^[A-Za-z0-9]{8,20}$/;
  return re.test(accountNumber);
};

export const validateAmount = (amount: string): boolean => {
  const num = parseFloat(amount);
  return !isNaN(num) && num > 0;
};