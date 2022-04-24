import React, { useState, useRef } from "react";
import { Link } from "react-router-dom";
// import navbar from bootstrap
import {Nav, NavDropdown, Form, FormControl, Button, Card } from "react-bootstrap";
import NavbarDNA from "../../components/Navbar/Navbar";
import styles from "./search.module.css"; 

function Search(){
    const [dummyResults, setDummyResults] = useState([]);

    const handleSubmit = (e) => {
      e.preventDefault();
      setDummyResults(oldDummyResults => {
        return [
          {
            date: "22 April 2022",
            name: "Gagas Praharsa Bahar",
            disease: "HIV",
            similarity: "30%",
            result: "False",
          },
          {
            date: "22 April 2022",
            name: "Gagas Praharsa Bahar",
            disease: "Epyrogenemie",
            similarity: "20%",
            result: "False",
          },
          {
            date: "22 April 2022",
            name: "Gagas Praharsa Bahar",
            disease: "Epylepsia",
            similarity: "15%",
            result: "False",
          },
          {
            date: "22 April 2022",
            name: "Gagas Praharsa Bahar",
            disease: "Meningitis",
            similarity: "5%",
            result: "False",
          },
        ];
      })
    }
  
  return (
    <>
    {/* show navbar from ../components/navbar */}
    <NavbarDNA/>


    {/* create body form input name and upload file sequence dna */}
    <div className={styles.body}>
      <div class="align-self-center">
        {/* title "Cari Hasil Tes" */}
        <div className={styles.container}>
          <div class="d-grid gap-3">
            <h1>Cari Hasil Tes</h1>
            {/* create a search bar by Date*/}
            <div className={styles.searchBar}>
                <Form onSubmit={handleSubmit}>
                    <div className={styles.form}>
                        <div class="d-grid gap-3">
                        <Form.Group controlId="formSearchBar">
                            <Form.Label>Tuliskan Tanggal dan Nama Penyakit atau Tanggal saja atau Nama Penyakit saja</Form.Label>
                            <Form.Control type="text" placeholder="contoh: 13 April 2002 HIV atau 13 April 2002 atau HIV" />
                        </Form.Group>
                        </div>
                    </div>
                    <div className={styles.search}>
                        <div class="col-md-12 text-center">
                        <Button variant="success" type="submit" >
                            <strong>Search</strong>
                        </Button>
                        </div>
                    </div>
                    <div className={styles.resultContainer}>
                        {dummyResults.length > 0 && (
                            <div className={styles.results}>
                            <p className={styles.resultsCount}>
                                Ditemukan {`${dummyResults.length}`} hasil!
                            </p>
                            <Card border ="success" style={{ width: '18rem' }}>
                                <Card.Body>
                        
                                    {/* card label and dummy */}
                                    <Card.Text>
                                    {dummyResults.map((result, idx) => {
                                        return (
                                            <div className={styles.resultCard}>
                                            <h3 className={styles.resultHeading}>
                                                Hasil Pencarian #{idx + 1}
                                            </h3>
                                            <div className={styles.resultFlex}>
                                                <p className={styles.resultInfoL}>Date</p>
                                                <p className={styles.resultInfo}>{result.date}</p>
                                            </div>
                                            <div className={styles.resultFlex}>
                                                <p className={styles.resultInfoL}>Patient</p>
                                                <p className={styles.resultInfo}>{result.name}</p>
                                            </div>
                                            <div className={styles.resultFlex}>
                                                <p className={styles.resultInfoL}>Disease</p>
                                                <p className={styles.resultInfo}>{result.disease}</p>
                                            </div>
                                            <div className={styles.resultFlex}>
                                                <p className={styles.resultInfoL}>Similarity</p>
                                                <p className={styles.resultInfo}>{result.similarity}</p>
                                            </div>
                                            <div className={styles.resultFlex}>
                                                <p className={styles.resultInfoL}>Result</p>
                                                <p className={styles.resultInfo}>{result.result}</p>
                                            </div>
                                            </div>
                                        );
                                        })}
                                    </Card.Text>
                                </Card.Body>
                            </Card> 
                            </div>
                        )}
                    </div>
                </Form>
            </div>           
          </div>
        </div>
      </div>
    </div>
    </>
  );
}
  export default Search;