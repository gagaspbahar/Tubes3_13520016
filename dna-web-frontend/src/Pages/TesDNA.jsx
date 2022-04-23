import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import styles from "./home.module.css";
// import navbar from bootstrap
import { Navbar, Nav, NavDropdown, Form, FormControl, Button } from "react-bootstrap";
import NavbarDNA from "../components/Navbar";
 

function Tes(){
  return (
    <>
    {/* show navbar from ../components/navbar */}
    <NavbarDNA/>
    {/* create body form input name and upload file sequence dna */}
    <div class="d-flex justify-content-center">
      <Form>
        <Form.Group controlId="formBasicEmail">
          <Form.Label>Name</Form.Label>
          <Form.Control type="text" placeholder="Enter name" />
        </Form.Group>

        <Form.Group controlId="formBasicPassword">
          <Form.Label>Upload Sequence DNA</Form.Label>
          <Form.Control type="file" placeholder="Enter file" />
        </Form.Group>
        <div class="col-md-12 text-center">
          <Button variant="primary" type="submit" >
            Submit
          </Button>
        </div>
      </Form>

    </div>
    </>



    
  );
}
  export default Tes;