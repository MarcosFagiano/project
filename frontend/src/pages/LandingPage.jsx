import React from 'react';
import LoginPage from "./LoginPage.jsx";
import Button from "bootstrap/js/src/button.js";
import {Link} from "react-router-dom";

function LandingPage() {
  return (
    <div className="d-flex justify-content-center align-items-center vh-100 bg-light text-dark">
      <h1>Bienvenido al Gimnasio</h1>
      <p>Gestioná tus actividades deportivas fácilmente.</p>
        <div className="mt-4">
        <Link to="/login" className="btn btn-primary me-2">Iniciar sesión</Link>
        <Link to="/register" className="btn btn-secondary me-2">Registrarse</Link>
        <Link to="/activities" className="btn btn-success">Ver actividades</Link>
      </div>
    </div>
  );
}

export default LandingPage;
