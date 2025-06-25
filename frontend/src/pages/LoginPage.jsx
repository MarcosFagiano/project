import { useState } from 'react';
import axios from 'axios';
import '../styles/formStyles.css';

function LoginPage() {
  const [formData, setFormData] = useState({
    email: '',
    password: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/login', formData);
      const token = response.data.token;
      if (token) {
        localStorage.setItem('token', token);
        console.log('Token almacenado:', token);
      } else {
        console.error('No se recibió token');
      }
    } catch (error) {
      console.error('Error en login:', error);
    }
  };

  return (
    <div className="d-flex justify-content-center align-items-center vh-100 bg-light text-dark">
      <h2>Iniciar sesión</h2>
      <form onSubmit={handleSubmit}>
        <fieldset>
          <legend>Acceso</legend>

          <div className="mb-3">
            <label className="form-label">Email</label>
            <input type="email" name="email" className="form-control"
                   placeholder="usuario@ejemplo.com"
                   value={formData.email} onChange={handleChange} />
          </div>

          <div className="mb-3">
            <label className="form-label">Contraseña</label>
            <input type="password" name="password" className="form-control"
                   placeholder="••••••"
                   value={formData.password} onChange={handleChange} autoComplete="off" />
          </div>

          <button type="submit" className="btn btn-primary">Entrar</button>
        </fieldset>
      </form>
    </div>
  );
}

export default LoginPage;
