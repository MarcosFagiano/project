import { useState } from 'react';
import axios from 'axios';
import '../styles/formStyles.css';

function RegisterPage() {
  const [formData, setFormData] = useState({
    given_name: '',
    family_name: '',
    phone_number: '',
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
      const response = await axios.post('http://localhost:8080/register', formData);
      sessionStorage.setItem('token', response.data.token);
      console.log('Registro exitoso:', response.data);
    } catch (error) {
      console.error('Error al registrar:', error);
    }
  };

  return (
    <div className="container mt-5">
      <h2>Registrarse</h2>
      <form onSubmit={handleSubmit}>
        <fieldset>
          <legend>Datos personales</legend>

          <div className="mb-3">
            <label className="form-label">Nombre</label>
            <input type="text" name="given_name" className="form-control"
                   placeholder="Nombre"
                   value={formData.given_name} onChange={handleChange} />
          </div>

          <div className="mb-3">
            <label className="form-label">Apellido</label>
            <input type="text" name="family_name" className="form-control"
                   placeholder="Apellido"
                   value={formData.family_name} onChange={handleChange} />
          </div>

          <div className="mb-3">
            <label className="form-label">Número de teléfono</label>
            <input type="text" name="phone_number" className="form-control"
                   placeholder="Número de teléfono"
                   value={formData.phone_number} onChange={handleChange} />
          </div>

          <div className="mb-3">
            <label className="form-label">Email</label>
            <input type="email" name="email" className="form-control"
                   placeholder="usuario@ejemplo.com"
                   value={formData.email} onChange={handleChange} />
            <small className="form-text text-muted">
              Nunca compartiremos tu email con nadie más.
            </small>
          </div>

          <div className="mb-3">
            <label className="form-label">Contraseña</label>
            <input type="password" name="password" className="form-control"
                   placeholder="••••••"
                   value={formData.password} onChange={handleChange} autoComplete="off" />
          </div>

          <button type="submit" className="btn btn-success">Registrarse</button>
        </fieldset>
      </form>
    </div>
  );
}

export default RegisterPage;
