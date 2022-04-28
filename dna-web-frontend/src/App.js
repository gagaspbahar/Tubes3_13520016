import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./Pages/Home/Home";
import Tes from "./Pages/TesDNA/TesDNA";
import Search from "./Pages/Search/Search";
import "./App.css";
import "bootstrap/dist/css/bootstrap.css";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Tes />} />
          <Route path="/add" element={<Home />} />
          <Route path="/test" element={<Tes />} />
          <Route path="/search" element={<Search />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
