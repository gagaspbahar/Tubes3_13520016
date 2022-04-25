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
            disease: "Sukses",
            similarity: "100%",
            result: "True",
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
                        <div class="d-flex flex-column justify-content-center align-items-center">
                            {dummyResults.length > 0 && (
                                <div className={styles.results}>
                                  <p className={styles.resultsCount}>
                                      Ditemukan {`${dummyResults.length}`} hasil!
                                  </p>
                                  {dummyResults.map((result, index) => {
                                      return (
                                    <table class="table table-striped">
                                      <thead>
                                        <tr>
                                          <th scope="col">#</th>
                                          <th scope="col">Date</th>
                                          <th scope="col">Patient</th>
                                          <th scope="col">Disease</th>
                                          <th scope="col">Similarity</th>
                                          <th scope="col">Result</th>
                                        </tr>
                                      </thead>
                                      <tbody>
                                        <tr>
                                          <th scope="row">{index+1}</th>
                                          <td>{result.date}</td>
                                          <td>{result.name}</td>
                                          <td>{result.disease}</td>
                                          <td>{result.similarity}</td>
                                          <td>{result.result}</td>
                                        </tr>
                                      </tbody>
                                    </table>                                         
                                  );
                                  })}
                                </div>
                            )}
                        </div>   
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