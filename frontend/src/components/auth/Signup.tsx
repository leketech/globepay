import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';

const Signup: React.FC = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    firstName: '',
    lastName: '',
    phoneNumber: '',
    dateOfBirth: '',
    country: '',
  });
  const [confirmPassword, setConfirmPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();
  const { signup, loading } = useAuth();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log('Signup form submitted');

    // Check if passwords match
    if (formData.password !== confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    try {
      console.log('Dispatching signup action with:', formData);
      const resultAction = await signup(formData);
      console.log('Signup successful, result:', resultAction);

      // Add a small delay to ensure state is updated before navigation
      setTimeout(() => {
        navigate('/dashboard');
      }, 100);
    } catch (err) {
      console.error('Signup error:', err);
      // Handle different types of errors
      if (err && typeof err === 'object' && 'message' in err) {
        setError((err as { message: string }).message || 'Signup failed');
      } else if (typeof err === 'string') {
        setError(err);
      } else {
        setError('Signup failed. Please try again.');
      }
    }
  };

  // Password strength indicator
  const getPasswordStrength = (password: string) => {
    if (password.length === 0) return 0;
    if (password.length < 6) return 1;
    if (password.length < 10) return 2;
    if (/[A-Z]/.test(password) && /[0-9]/.test(password)) return 4;
    return 3;
  };

  const passwordStrength = getPasswordStrength(formData.password);
  const strengthText = ['None', 'Weak', 'Fair', 'Good', 'Strong'][passwordStrength];
  const strengthColor = ['gray', 'red', 'yellow', 'blue', 'green'][passwordStrength];

  return (
    <div className="relative flex min-h-screen w-full flex-col">
      <div className="flex h-full min-h-screen w-full flex-col lg:flex-row">
        {/* Left Pane (Branding) - Hidden on mobile */}
        <div className="relative hidden w-full flex-col items-center justify-center bg-primary/10 p-10 dark:bg-primary/20 lg:flex lg:w-1/2">
          <div className="absolute inset-0 z-0">
            <img
              alt=""
              className="h-full w-full object-cover opacity-10"
              src="https://lh3.googleusercontent.com/aida-public/AB6AXuDRdgdj1wI3eBHoQexTTW30mauLEjbJaovOUNpg96LNcwgNrJkxGUNqJ2zjpP73fabGYUzS3QQ3JopS8lJHiJXmQln4UYqOJ54U8TSgqlvpdg6tjpYVg6bPjUiac-mTXwYjv6eTzKr_xpAyMriOp98QTYBlKhlNUq_GcbzsXBjE5jF_hiUEOk8gD51ZZmetTz8jPxrUm1pEknFKFcGyRcBYxiZ-UDZH1OcZgnD3Nsyz7kjDprEIIr7K6pokFpsZlBJigygA4IGHZAL0"
            />
          </div>
          <div className="relative z-10 flex max-w-md flex-col items-start gap-6 text-left">
            <div className="flex items-center gap-3">
              <div className="size-8 text-primary">
                <svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
                  <path
                    d="M42.4379 44C42.4379 44 36.0744 33.9038 41.1692 24C46.8624 12.9336 42.2078 4 42.2078 4L7.01134 4C7.01134 4 11.6577 12.932 5.96912 23.9969C0.876273 33.9029 7.27094 44 7.27094 44L42.4379 44Z"
                    fill="currentColor"
                  ></path>
                </svg>
              </div>
              <h1 className="text-3xl font-bold text-gray-800 dark:text-gray-100">Globepay</h1>
            </div>
            <p className="text-4xl font-black leading-tight tracking-tight text-gray-800 dark:text-gray-100">
              Securely send money, <br /> instantly.
            </p>
            <p className="text-base font-normal text-gray-600 dark:text-gray-300">
              Join millions of users who trust Globepay for fast and reliable international
              transfers.
            </p>
          </div>
        </div>

        {/* Right Pane (Form) */}
        <div className="flex w-full flex-col items-center justify-center p-6 lg:w-1/2">
          <div className="flex w-full max-w-md flex-col items-start justify-center gap-8">
            <header className="flex w-full items-center justify-between">
              <div className="flex items-center gap-3 lg:hidden">
                <div className="size-6 text-primary">
                  <svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
                    <path
                      d="M42.4379 44C42.4379 44 36.0744 33.9038 41.1692 24C46.8624 12.9336 42.2078 4 42.2078 4L7.01134 4C7.01134 4 11.6577 12.932 5.96912 23.9969C0.876273 33.9029 7.27094 44 7.27094 44L42.4379 44Z"
                      fill="currentColor"
                    ></path>
                  </svg>
                </div>
                <h2 className="text-xl font-bold leading-tight tracking-[-0.015em] text-gray-800 dark:text-gray-100">
                  Globepay
                </h2>
              </div>
            </header>

            <div className="flex w-full flex-col gap-8">
              <div className="flex flex-col gap-2">
                <p className="text-4xl font-black leading-tight tracking-[-0.033em] text-gray-900 dark:text-white">
                  Create Your Globepay Account
                </p>
                <p className="text-base font-normal leading-normal text-gray-500 dark:text-gray-400">
                  Join Globepay today
                </p>
              </div>

              {error && (
                <div className="mb-4 bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded relative">
                  {error}
                </div>
              )}

              <form className="flex w-full flex-col gap-4" onSubmit={handleSubmit}>
                <div className="grid grid-cols-2 gap-4">
                  <label className="flex flex-1 flex-col">
                    <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                      First Name
                    </p>
                    <input
                      className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary"
                      placeholder="John"
                      name="firstName"
                      value={formData.firstName}
                      onChange={handleChange}
                      required
                    />
                  </label>

                  <label className="flex flex-1 flex-col">
                    <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                      Last Name
                    </p>
                    <input
                      className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary"
                      placeholder="Doe"
                      name="lastName"
                      value={formData.lastName}
                      onChange={handleChange}
                      required
                    />
                  </label>
                </div>

                <label className="flex flex-1 flex-col">
                  <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                    Email Address
                  </p>
                  <input
                    className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary"
                    placeholder="you@example.com"
                    type="email"
                    name="email"
                    value={formData.email}
                    onChange={handleChange}
                    required
                  />
                </label>

                <label className="flex flex-1 flex-col">
                  <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                    Password
                  </p>
                  <div className="relative">
                    <input
                      className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary pr-12"
                      placeholder="Enter your password"
                      type={showPassword ? 'text' : 'password'}
                      name="password"
                      value={formData.password}
                      onChange={handleChange}
                      required
                    />
                    <button
                      type="button"
                      className="absolute inset-y-0 right-0 pr-3 flex items-center"
                      onClick={() => setShowPassword(!showPassword)}
                    >
                      {showPassword ? (
                        <svg
                          className="h-5 w-5 text-gray-500"
                          fill="none"
                          viewBox="0 0 24 24"
                          stroke="currentColor"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                          />
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                          />
                        </svg>
                      ) : (
                        <svg
                          className="h-5 w-5 text-gray-500"
                          fill="none"
                          viewBox="0 0 24 24"
                          stroke="currentColor"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543 7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"
                          />
                        </svg>
                      )}
                    </button>
                  </div>
                </label>

                <div className="flex items-center gap-2 pt-1">
                  {[1, 2, 3, 4].map((index) => (
                    <div
                      key={index}
                      className={`h-1 flex-1 rounded-full ${
                        index <= passwordStrength
                          ? strengthColor === 'red'
                            ? 'bg-red-500'
                            : strengthColor === 'yellow'
                            ? 'bg-yellow-500'
                            : strengthColor === 'blue'
                            ? 'bg-blue-500'
                            : strengthColor === 'green'
                            ? 'bg-green-500'
                            : 'bg-gray-300'
                          : 'bg-border-light dark:bg-border-dark'
                      }`}
                    ></div>
                  ))}
                  <p
                    className={`text-xs font-medium ${
                      strengthColor === 'red'
                        ? 'text-red-500'
                        : strengthColor === 'yellow'
                        ? 'text-yellow-500'
                        : strengthColor === 'blue'
                        ? 'text-blue-500'
                        : strengthColor === 'green'
                        ? 'text-green-500'
                        : 'text-gray-500'
                    } pl-2`}
                  >
                    {strengthText}
                  </p>
                </div>

                <label className="flex flex-1 flex-col">
                  <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                    Confirm Password
                  </p>
                  <input
                    className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary"
                    placeholder="Confirm your password"
                    type="password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    required
                  />
                </label>

                <label className="flex flex-1 flex-col">
                  <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                    Phone Number (Optional)
                  </p>
                  <input
                    className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary"
                    placeholder="+1 (555) 123-4567"
                    type="tel"
                    name="phoneNumber"
                    value={formData.phoneNumber}
                    onChange={handleChange}
                  />
                </label>

                <div className="grid grid-cols-2 gap-4">
                  <label className="flex flex-1 flex-col">
                    <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                      Date of Birth (Optional)
                    </p>
                    <input
                      className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary"
                      type="date"
                      name="dateOfBirth"
                      value={formData.dateOfBirth}
                      onChange={handleChange}
                    />
                  </label>

                  <label className="flex flex-1 flex-col">
                    <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">
                      Country (Optional)
                    </p>
                    <select
                      className="form-select h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary"
                      name="country"
                      value={formData.country}
                      onChange={handleChange}
                    >
                      <option value="">Select your country</option>
                      <option value="US">United States</option>
                      <option value="GB">United Kingdom</option>
                      <option value="CA">Canada</option>
                      <option value="AU">Australia</option>
                      <option value="DE">Germany</option>
                      <option value="FR">France</option>
                    </select>
                  </label>
                </div>

                <div className="flex items-start gap-3 pt-2">
                  <input
                    className="form-checkbox mt-0.5 h-5 w-5 rounded border-border-light dark:border-border-dark bg-background-light dark:bg-subtle-dark text-primary focus:ring-primary/50"
                    id="terms-checkbox"
                    type="checkbox"
                    required
                  />
                  <label
                    className="text-sm text-text-light/80 dark:text-text-dark/80"
                    htmlFor="terms-checkbox"
                  >
                    By creating an account, you agree to our{' '}
                    <a className="font-semibold text-primary hover:underline" href="#">
                      Terms of Service
                    </a>{' '}
                    and{' '}
                    <a className="font-semibold text-primary hover:underline" href="#">
                      Privacy Policy
                    </a>
                    .
                  </label>
                </div>

                <button
                  type="submit"
                  disabled={loading}
                  className="flex h-12 w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg bg-primary text-base font-bold text-white transition-colors hover:bg-primary/90 disabled:opacity-50"
                >
                  {loading ? 'Creating account...' : 'Create Account'}
                </button>
              </form>

              <div className="flex w-full flex-col items-center gap-4">
                <Link
                  to="/"
                  className="flex cursor-pointer items-center gap-2 text-sm font-medium text-gray-600 hover:text-primary dark:text-gray-300 dark:hover:text-primary"
                >
                  <span className="material-symbols-outlined">arrow_back</span>
                  Back to Landing Page
                </Link>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                  Already have an account?{' '}
                  <Link to="/login" className="font-medium text-primary hover:underline">
                    Sign In
                  </Link>
                </p>
              </div>
            </div>

            <footer className="w-full text-center">
              <p className="text-xs text-gray-400 dark:text-gray-500">
                © 2024 Globepay. All Rights Reserved. <br />
                <a className="underline hover:text-primary" href="#">
                  Terms of Service
                </a>{' '}
                •{' '}
                <a className="underline hover:text-primary" href="#">
                  Privacy Policy
                </a>
              </p>
            </footer>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Signup;
