function RegisterPage() {
  return (
    <div className="container mt-5">
      <h2>Registrarse</h2>
      <form>
        <div className="mb-3">
          <label className="form-label">Nombre</label>
          <input type="text" className="form-control" placeholder="Nombre completo" />
        </div>
        <div className="mb-3">
          <label className="form-label">Email</label>
          <input type="email" className="form-control" placeholder="usuario@ejemplo.com" />
        </div>
        <div className="mb-3">
          <label className="form-label">Contraseña</label>
          <input type="password" className="form-control" placeholder="••••••" />
        </div>
        <button type="submit" className="btn btn-success">Registrarse</button>
      </form>
    </div>
  );
}

export default RegisterPage;
