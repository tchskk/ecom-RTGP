import React from 'react';
import { useAuth } from '../contexts/AuthContext';
import { Link } from 'react-router-dom';

const Dashboard: React.FC = () => {
  const { user, logout } = useAuth();

  return (
    <div className="min-h-screen bg-gray-100">
      <nav className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <h1 className="text-xl font-bold text-gray-800">RTGP Dashboard</h1>
            </div>
            <div className="flex items-center space-x-4">
              <span className="text-gray-700">
                {user?.first_name} {user?.last_name}
                {user?.role === 'admin' && (
                  <span className="ml-2 px-2 py-1 text-xs bg-purple-100 text-purple-800 rounded">
                    Admin
                  </span>
                )}
              </span>
              {user?.role === 'admin' && (
                <Link
                  to="/admin"
                  className="px-4 py-2 text-sm font-medium text-white bg-purple-600 rounded-md hover:bg-purple-700"
                >
                  Admin Panel
                </Link>
              )}
              <button
                onClick={logout}
                className="px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-md hover:bg-red-700"
              >
                Logout
              </button>
            </div>
          </div>
        </div>
      </nav>

      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <div className="bg-white rounded-lg shadow p-6">
            <h2 className="text-2xl font-bold text-gray-800 mb-4">
              Welcome, {user?.first_name}!
            </h2>
            <p className="text-gray-600 mb-4">
              You are logged in as <span className="font-semibold">{user?.email}</span>
            </p>
            
            <div className="mt-6 border-t pt-6">
              <h3 className="text-lg font-semibold text-gray-800 mb-3">Profile Information</h3>
              <dl className="space-y-2">
                <div className="flex">
                  <dt className="text-gray-600 w-32">Email:</dt>
                  <dd className="text-gray-800 font-medium">{user?.email}</dd>
                </div>
                <div className="flex">
                  <dt className="text-gray-600 w-32">Name:</dt>
                  <dd className="text-gray-800 font-medium">
                    {user?.first_name} {user?.last_name}
                  </dd>
                </div>
                <div className="flex">
                  <dt className="text-gray-600 w-32">Role:</dt>
                  <dd className="text-gray-800 font-medium capitalize">{user?.role}</dd>
                </div>
                <div className="flex">
                  <dt className="text-gray-600 w-32">Member Since:</dt>
                  <dd className="text-gray-800 font-medium">
                    {user?.created_at && new Date(user.created_at).toLocaleDateString()}
                  </dd>
                </div>
              </dl>
            </div>

            <div className="mt-6 border-t pt-6">
              <h3 className="text-lg font-semibold text-gray-800 mb-3">
                Coming Soon: E-commerce Features
              </h3>
              <p className="text-gray-600">
                This project will be extended with e-commerce functionalities including:
              </p>
              <ul className="mt-2 list-disc list-inside text-gray-600 space-y-1">
                <li>Product catalog</li>
                <li>Shopping cart</li>
                <li>Order management</li>
                <li>Payment processing</li>
                <li>And much more!</li>
              </ul>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;
