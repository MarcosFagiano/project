import { Link } from 'react-router-dom';

function LandingPage() {
  return (
    <div className="container text-center mt-5">
      <h1 className="mb-4">Bienvenido al Gimnasio</h1>
      <p>Gestiona tus actividades deportivas fácilmente.</p>
      <div className="mt-4">
        <Link to="/login" className="btn btn-primary me-2">Iniciar sesión</Link>
        <Link to="/register" className="btn btn-secondary">Registrarse</Link>
      </div>
    </div>
  );
}

export default LandingPage;
