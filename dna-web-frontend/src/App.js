import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Home from './Pages/Home';
import Tes from './Pages/TesDNA';
import './App.css';
import 'bootstrap/dist/css/bootstrap.css';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/test" element={<Tes />} />
      </Routes>
    </Router>
  );
}

export default App;
