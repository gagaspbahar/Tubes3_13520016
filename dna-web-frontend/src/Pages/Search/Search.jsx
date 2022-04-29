import React, { useState } from "react";
// import navbar from bootstrap
import { Form, Button } from "react-bootstrap";
import NavbarDNA from "../../components/Navbar/Navbar";
import styles from "./search.module.css";

const LOCALBACKEND = "http://localhost:8080";
const HEROKUBACKEND = "https://shrouded-mountain-85549.herokuapp.com/"

function Search() {
  const axios = require("axios");
  const axiosInstance = axios.create({
    baseURL: HEROKUBACKEND,
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Headers": "*",
    },
  });

  const [dummyResults, setDummyResults] = useState([]);
  const [query, setQuery] = useState("");

  const handleChange = (e) => {
    setQuery(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const data = {
      query: query,
    };

    axiosInstance
      .post("/history", data, {
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
      })
      .then((res) => {
        let result = res.data.records;
        setDummyResults(JSON.parse(result));
      })
      .catch((err) => console.log(err));
  };

  return (
    <>
      {/* show navbar from ../components/navbar */}
      <NavbarDNA />

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
                        <Form.Label>
                          Tuliskan Tanggal dan Nama Penyakit atau Tanggal saja
                          atau Nama Penyakit saja
                        </Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="contoh: 13 April 2002 HIV atau 13 April 2002 atau HIV"
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </div>
                  </div>
                  <div className={styles.search}>
                    <div class="col-md-12 text-center">
                      <Button variant="success" type="submit">
                        <strong>Search</strong>
                      </Button>
                    </div>
                  </div>
                  <div className={styles.resultContainer}>
                    <div class="d-flex flex-column justify-content-center align-items-center">
                      {dummyResults != null && dummyResults.length > 0 && (
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
                                    <th scope="row">{index + 1}</th>
                                    <td>{result.tanggal}</td>
                                    <td>{result.nama_pengguna}</td>
                                    <td>{result.nama_penyakit}</td>
                                    <td>{result.similarity}</td>
                                    <td>
                                      {result.status_tes ? "True" : "False"}
                                    </td>
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
