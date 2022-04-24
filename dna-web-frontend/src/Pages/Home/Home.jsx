import React, { useState, useRef } from "react";
import { Link } from "react-router-dom";
// import navbar from bootstrap
import {Nav, NavDropdown, Form, FormControl, Button } from "react-bootstrap";
import NavbarDNA from "../../components/Navbar/Navbar";
import styles from "./home.module.css"; 

function Home(){


  return (
    <>
    {/* show navbar from ../components/navbar */}
    <NavbarDNA/>


    {/* create body form input name and upload file sequence dna */}
    <div className={styles.body}>
      <div class="align-self-center">
        {/* title "Tambahkan Penyakit" */}
        <div className={styles.container}>
          <div class="d-grid gap-3">
            <h1>Tambahkan Penyakit</h1>
            <Form>
              <div className={styles.form}>
                <div class="d-grid gap-3">
                  <Form.Group controlId="formBasicNamaPenyakit">
                    <Form.Label>Nama Penyakit</Form.Label>
                    <Form.Control type="text" placeholder="contoh: HIV" />
                  </Form.Group>

                  <Form.Group controlId="formBasicSequenceDNA">
                    <Form.Label>Upload Sequence DNA</Form.Label>
                    <Form.Control type="file" placeholder="Enter file" />
                  </Form.Group>
                </div>
              </div>
              <div className={styles.submit}>
                <div class="col-md-12 text-center">
                  <Button variant="success" type="submit" >
                    <strong>Submit</strong>
                  </Button>
                </div>
              </div>
            </Form>
          </div>
        </div>
      </div>
    </div>
    </>

  );
}
  export default Home;