import React, { useState, useRef } from "react";
import { Link } from "react-router-dom";
// import navbar from bootstrap
import { Navbar, Nav, NavDropdown, Form, FormControl, Button, Card, Badge, Label } from "react-bootstrap";
import NavbarDNA from "../../components/Navbar/Navbar";
import styles from "./tesDNA.module.css";
 

function Tes(){
  const textRef = useRef(null);
  const infoRef = useRef(null);

  const [isSubmitted, setIsSubmitted] = useState(false);

  const dummyData = {
    date: "22 April 2022",
    name: "Gagas Praharsa Bahar",
    disease: "HIV",
    similarity: "30%",
    result: "False"
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setIsSubmitted(true);
  }

  return (
    <>
    {/* show navbar from ../components/navbar */}
    <NavbarDNA/>

    {/* create body form input name and upload file sequence dna */}
    <div className={styles.body}>
      <div class="align-self-center">
        {/* title "Cek DNA Pasien" */}
        <div className={styles.container}>
          <div class="d-grid gap-3">
            <h1>Cek DNA Pasien</h1>
            <Form onSubmit={handleSubmit}>
              <div className={styles.form}>
              <div class="d-grid gap-3">
                <Form.Group controlId="formBasicNamaPasien">
                  <Form.Label>Nama Pasien</Form.Label>
                  <Form.Control type="text" placeholder="Nama Pasien" />
                </Form.Group>

                <Form.Group controlId="formBasicSequenceDNA">
                  <Form.Label>Upload Sequence DNA</Form.Label>
                  <Form.Control type="file" placeholder="Enter file" />
                </Form.Group>

                <Form.Group controlId="formBasicNamaPenyakit">
                  <Form.Label>Nama Penyakit</Form.Label>
                  <Form.Control type="text" placeholder="Nama Penyakit" />
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
        {isSubmitted && (
        <div className={styles.result}>
            <Card border ="success" style={{ width: '18rem' }}>
              <Card.Body>
                {/* card title in center */}
                  <h3 className={styles.resultHeading}>Hasil</h3>
                {/* card label and dummy */}
                <Card.Text>
                <div className={styles.resultCard}>
                  <div className={styles.resultFlex}>
                    <p className={styles.resultInfoL}>Date</p>
                    <p className={styles.resultInfo}>{dummyData.date}</p>
                  </div>
                  <div className={styles.resultFlex}>
                    <p className={styles.resultInfoL}>Patient</p>
                    <p className={styles.resultInfo}>{dummyData.name}</p>
                  </div>
                  <div className={styles.resultFlex}>
                    <p className={styles.resultInfoL}>Disease</p>
                    <p className={styles.resultInfo}>{dummyData.disease}</p>
                  </div>
                  <div className={styles.resultFlex}>
                    <p className={styles.resultInfoL}>Similarity</p>
                    <p className={styles.resultInfo}>{dummyData.similarity}</p>
                  </div>
                  <div className={styles.resultFlex}>
                    <p className={styles.resultInfoL}>Result</p>
                    <p className={styles.resultInfo}>{dummyData.result}</p>
                  </div>
                </div>
                </Card.Text>
              </Card.Body>
            </Card>
      
        </div>
        )}  
      </div>
    </div>
    </>

  );
}
  export default Tes;