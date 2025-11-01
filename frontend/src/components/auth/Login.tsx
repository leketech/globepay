import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { login } from '../../store/authSlice';
import type { AppDispatch, RootState } from '../../store';

const Login: React.FC = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch<AppDispatch>();
  const { loading, error } = useSelector((state: RootState) => state.auth);
  
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    try {
      // Call login and unwrap the result to catch any errors
      const result = await dispatch(login({ email, password })).unwrap();
      console.log('Login successful, result:', result);
      
      // Navigate to dashboard
      console.log('Navigating to dashboard');
      navigate('/dashboard');
    } catch (err: any) {
      console.error('Login error caught:', err);
    }
  };

  return (
    <div className="relative flex min-h-screen w-full flex-col">
      <div className="flex h-full min-h-screen w-full flex-col lg:flex-row">
        {/* Left Pane (Branding) */}
        <div className="relative hidden w-full flex-col items-center justify-center bg-primary/10 p-10 dark:bg-primary/20 lg:flex lg:w-1/2">
          <div className="absolute inset-0 z-0">
            <img 
              alt="" 
              className="h-full w-full object-cover opacity-10" 
              src="https://lh3.googleusercontent.com/aida-public/AB6AXuDRdgdj1wI3eBHoQexTTW30mauLEjbJaovOUNpg96LNcwgNrJkxGUNqJ2zjpP73fabGYUzS3QQ3JopS8lJHiJXmQln4UYqOJ54U8TSgqlvpdg6tjpYVg6bPjUiac-mTXwYjv6eTzKr_xpAyMriOp98QTYBlKhlNUq_GcbzsXBjE5jF_hiUEOk8gD51ZZmetT8jPxrUm1pEknFKFcGyRcBYxiZ-UDZH1OcZgnD3Nsyz7kjDprEIIr7K6pokFpsZlBJigygA4IGHZAL0"
            />
          </div>
          <div className="relative z-10 flex max-w-md flex-col items-start gap-6 text-left">
            <div className="flex items-center gap-3">
              <div className="size-8 text-primary">
                <svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
                  <path d="M42.4379 44C42.4379 44 36.0744 33.9038 41.1692 24C46.8624 12.9336 42.2078 4 42.2078 4L7.01134 4C7.01134 4 11.6577 12.932 5.96912 23.9969C0.876273 33.9029 7.27094 44 7.27094 44L42.4379 44Z" fill="currentColor"></path>
                </svg>
              </div>
              <h1 className="text-3xl font-bold text-gray-800 dark:text-gray-100">Globepay</h1>
            </div>
            <p className="text-4xl font-black leading-tight tracking-tight text-gray-800 dark:text-gray-100">
              Securely send money, <br/> instantly.
            </p>
            <p className="text-base font-normal text-gray-600 dark:text-gray-300">
              Join millions of users who trust Globepay for fast and reliable international transfers.
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
                    <path d="M42.4379 44C42.4379 44 36.0744 33.9038 41.1692 24C46.8624 12.9336 42.2078 4 42.2078 4L7.01134 4C7.01134 4 11.6577 12.932 5.96912 23.9969C0.876273 33.9029 7.27094 44 7.27094 44L42.4379 44Z" fill="currentColor"></path>
                  </svg>
                </div>
                <h2 className="text-xl font-bold leading-tight tracking-[-0.015em] text-gray-800 dark:text-gray-100">Globepay</h2>
              </div>
            </header>
            
            <div className="flex w-full flex-col gap-8">
              <div className="flex flex-col gap-2">
                <p className="text-4xl font-black leading-tight tracking-[-0.033em] text-gray-900 dark:text-white">Welcome Back</p>
                <p className="text-base font-normal leading-normal text-gray-500 dark:text-gray-400">Log in to your Globepay account</p>
              </div>
              
              {error && (
                <div className="mb-4 bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded relative">
                  {error}
                </div>
              )}
              
              <form className="flex w-full flex-col gap-4" onSubmit={handleSubmit}>
                <label className="flex flex-1 flex-col">
                  <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">Email Address</p>
                  <input 
                    className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg border border-gray-300 bg-background-light p-[15px] text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary" 
                    placeholder="you@example.com" 
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                  />
                </label>
                
                <div className="flex flex-col">
                  <label className="flex flex-1 flex-col">
                    <p className="pb-2 text-base font-medium leading-normal text-gray-800 dark:text-gray-200">Password</p>
                    <div className="flex w-full flex-1 items-stretch">
                      <input 
                        className="form-input h-14 w-full min-w-0 flex-1 resize-none overflow-hidden rounded-l-lg border border-r-0 border-gray-300 bg-background-light p-[15px] pr-2 text-base font-normal leading-normal text-gray-900 placeholder:text-gray-400 focus:border-primary focus:outline-0 focus:ring-2 focus:ring-primary/20 dark:border-gray-600 dark:border-r-0 dark:bg-gray-800 dark:text-white dark:placeholder:text-gray-500 dark:focus:border-primary" 
                        placeholder="Enter your password" 
                        type={showPassword ? "text" : "password"}
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                      />
                      <div 
                        className="flex cursor-pointer items-center justify-center rounded-r-lg border border-l-0 border-gray-300 bg-background-light px-3 text-gray-400 dark:border-gray-600 dark:border-l-0 dark:bg-gray-800 dark:text-gray-500"
                        onClick={() => setShowPassword(!showPassword)}
                      >
                        <span className="material-symbols-outlined text-2xl">
                          {showPassword ? 'visibility_off' : 'visibility'}
                        </span>
                      </div>
                    </div>
                  </label>
                  <div className="flex w-full justify-end">
                    <a className="cursor-pointer pt-2 text-sm font-medium leading-normal text-primary hover:underline" href="#">Forgot Password?</a>
                  </div>
                </div>
                
                <button 
                  type="submit"
                  disabled={loading}
                  className="flex h-12 w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg bg-primary text-base font-bold text-white transition-colors hover:bg-primary/90 disabled:opacity-50"
                >
                  {loading ? 'Signing in...' : 'Login'}
                </button>
              </form>
              
              <div className="flex w-full flex-col items-center gap-4">
                <Link to="/" className="flex cursor-pointer items-center gap-2 text-sm font-medium text-gray-600 hover:text-primary dark:text-gray-300 dark:hover:text-primary">
                  <span className="material-symbols-outlined">arrow_back</span>
                  Back to Landing Page
                </Link>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                  Don't have an account? <Link to="/signup" className="font-medium text-primary hover:underline">Sign Up</Link>
                </p>
              </div>
            </div>
            
            <footer className="w-full text-center">
              <p className="text-xs text-gray-400 dark:text-gray-500">
                © 2024 Globepay. All Rights Reserved. <br/>
                <a className="underline hover:text-primary" href="#">Terms of Service</a> • <a className="underline hover:text-primary" href="#">Privacy Policy</a>
              </p>
            </footer>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;