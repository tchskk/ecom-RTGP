import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const API_URL = "http://localhost:8080/api";

type User = {
  id: number;
  email: string;
  role: string;
  createdAt: string;
};

export default function AdminDashboard() {
  const nav = useNavigate();
  const [users, setUsers] = useState<User[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      nav("/login");
      return;
    }

    (async () => {
      try {
        const res = await fetch(`${API_URL}/admin/users`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        if (res.status === 401 || res.status === 403) {
          setError("Not authorized as admin");
          return;
        }
        const data = await res.json();
        if (!res.ok) {
          setError(data.error || "Failed to load users");
          return;
        }
        // console.log(data);
        setUsers(data.users);
      } catch {
        setError("Network error");
      }
    })();
  }, [nav]);

  return (
    <div className="min-h-screen bg-slate-100 p-6">
      <div className="max-w-4xl mx-auto bg-white rounded-xl shadow-md p-6">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-xl font-semibold">Admin Dashboard</h1>
          <button
            onClick={() => {
              localStorage.removeItem("token");
              nav("/login");
            }}
            className="text-sm border px-3 py-1 rounded-md"
          >
            Logout
          </button>
        </div>
        {error && <p className="text-red-500 text-sm mb-3">{error}</p>}

        <table className="w-full text-sm border">
          <thead className="bg-slate-50">
            <tr>
              <th className="border px-3 py-2 text-left">ID</th>
              <th className="border px-3 py-2 text-left">Email</th>
              <th className="border px-3 py-2 text-left">Role</th>
              <th className="border px-3 py-2 text-left">Created</th>
            </tr>
          </thead>
          <tbody>
            {users.map(u => (
              <tr key={u.id}>
                <td className="border px-3 py-1">{u.id}</td>
                <td className="border px-3 py-1">{u.email}</td>
                <td className="border px-3 py-1">{u.role}</td>
                <td className="border px-3 py-1">
                  {new Date(u.createdAt).toLocaleString()}
                </td>
              </tr>
            ))}
            {users.length === 0 && !error && (
              <tr>
                <td className="border px-3 py-2 text-center" colSpan={4}>
                  No users yet.
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}