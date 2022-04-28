import React from "react";
import styles from "./navbar.module.css";
// import navbar from bootstrap
 

function NavbarDNA(){
    return (
      <div className={styles.navbar}>
        {/* create navbar variant dark */}
        <nav class="navbar navbar-expand-lg bg-success navbar-dark">
            <div class="container-fluid">
                <a class="navbar-brand ms-4" href="/">
                    {/* <img src={logo} width="50" height="50" class="d-inline-block align-top" alt=""/> */}
                  <strong>Gafira Hospital+</strong>
                </a>
                <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
                <li class="nav-item">
                    <a class="nav-link" href="/test"><strong>Tes DNA</strong></a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/search"><strong>Search</strong></a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/add"><strong>Tambahkan Penyakit</strong></a>
                </li>
                </ul>
            </div>
        </nav>
      </div>
    );
  }
    export default NavbarDNA;