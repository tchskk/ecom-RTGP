import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
});

// Add token to requests if available
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  role: string;
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  first_name: string;
  last_name: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export const authAPI = {
  register: (data: RegisterRequest) =>
    axios.post<LoginResponse>(`${API_BASE_URL}/auth/register`, data),
  
  login: (data: LoginRequest) =>
    axios.post<LoginResponse>(`${API_BASE_URL}/auth/login`, data),
  
  getProfile: () =>
    api.get<User>('/profile'),
};

export const adminAPI = {
  getAllUsers: () =>
    api.get<User[]>('/admin/users'),
};

export default api;
