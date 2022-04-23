import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import styles from "./navbar.module.css";
// import navbar from bootstrap
import { Navbar, Nav, NavDropdown, Form, FormControl, Button } from "react-bootstrap";
 

function NavbarDNA(){
    return (
      <div className={styles.navbar}>
        {/* create navbar variant dark */}
        <nav class="navbar navbar-expand-lg bg-success navbar-dark">
            <div class="container-fluid">
                <a class="navbar-brand ms-4" href="#">
                  <strong>CekKesehatan.com</strong>
                </a>
                <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
                {/* create navbar brand */}
                <li class="nav-item">
                    <a class="nav-link active" href="#"><strong>Tes DNA</strong></a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="#"><strong>Lihat Hasil</strong></a>
                </li>
                </ul>
            </div>
        </nav>
      </div>
    );
  }
    export default NavbarDNA;