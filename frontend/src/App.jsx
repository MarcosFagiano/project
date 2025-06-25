import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import LandingPage from './pages/LandingPage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ActivitiesPage from './pages/ActivitiesPage';

function App() {
  return (
    <Router>
      <div className="d-flex flex-column min-vh-100">
        <header>
          <Navbar />
        </header>

        <main className="d-flex flex-column min-vw-100 text-center align-content-center">
          <Routes>
            <Route path="/" element={<LandingPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/activities" element={<ActivitiesPage />} />
          </Routes>
        </main>

        <footer className="bg-dark text-white text-center py-3">
          <Footer />
        </footer>
      </div>
    </Router>
  );
}

export default App;
