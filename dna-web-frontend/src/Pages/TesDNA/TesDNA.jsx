import React, { useState } from "react";
// import navbar from bootstrap
import {
  Form,
  Button,
  Card,
} from "react-bootstrap";
import NavbarDNA from "../../components/Navbar/Navbar";
import styles from "./tesDNA.module.css";

const LOCALBACKEND = "http://localhost:8080";

function Tes() {
  const axios = require("axios");
  const axiosInstance = axios.create({
    baseURL: LOCALBACKEND,
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Headers": "*",
    },
  });


  const defaultData = {
    tanggal: "22 April 2022",
    nama_pengguna: "Gagas Praharsa Bahar",
    nama_penyakit: "HIV",
    similarity: "30%",
    status_tes: "False",
  };

  const [textName, setTextName] = useState("");
  const [namaPasien, setNamaPasien] = useState("");
  const [namaPenyakit, setNamaPenyakit] = useState("");
  const [metode, setMetode] = useState("KMP");
  const [dummyData, setDummyData] = useState(defaultData);

  const handleChangePasien = (e) => {
    setNamaPasien(e.target.value);
  };

  const handleRadioButton = (e) => {
    setMetode(e.target.value);
  };

  const handleChangePenyakit = (e) => {
    setNamaPenyakit(e.target.value);
  };

  const handleUploadText = function(ev) {
    // ev.preventDefault();
    setTextName(ev.target.files[0].name);
    const data = new FormData();
    data.append("file", ev.target.files[0]);
    const filename = ev.target.files[0].name;
    axiosInstance
      .post("/uploadUser/" + filename, data, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((res) => {
        console.log(res);
      })
      .catch((err) => console.log(err));
  };

  const [isSubmitted, setIsSubmitted] = useState(false);


  const handleSubmit = (e) => {
    e.preventDefault();
    setIsSubmitted(true);
    const seq = textName.slice(0, -4)
    const data = {
      nama: namaPasien,
      penyakit: namaPenyakit,
      metode: metode,
      sequenceDNA: seq,
    };
    let result = {};
    axiosInstance
      .post("/tes", data, {
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
      })
      .then((res) => {
        console.log(res.data);
        result = res.data;
        setDummyData(result);
      })
      .catch((err) => console.log(err));
      console.log(dummyData)
    };

  return (
    <>
      {/* show navbar from ../components/navbar */}
      <NavbarDNA />

      {/* create body form input name and upload file sequence dna */}
      <div className={styles.body}>
        <div class="align-self-center">
          {/* title "Cek DNA Pasien" */}
          <div className={styles.container}>
            <div class="d-grid gap-3">
              <h1>Tes DNA Pasien</h1>
              <Form onSubmit={handleSubmit}>
                <div className={styles.form}>
                  <div class="d-grid gap-3">
                    <Form.Group controlId="formBasicNamaPasien">
                      <Form.Label>Nama Pasien</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="contoh: Gagas Praharsa Bahar"
                        onChange={handleChangePasien}
                      />
                    </Form.Group>

                    <Form.Group controlId="formBasicSequenceDNA">
                      <Form.Label>Upload Sequence DNA</Form.Label>
                      <Form.Control
                        type="file"
                        placeholder="Enter file"
                        onChange={handleUploadText}
                      />
                    </Form.Group>

                    <Form.Group controlId="formBasicNamaPenyakit">
                      <Form.Label>Nama Penyakit</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="contoh: HIV"
                        onChange={handleChangePenyakit}
                      />
                    </Form.Group>
                  </div>
                </div>
                {/* create radio */}
                <div className={styles.radio}>
                  <div class="d-grid gap-3">
                    <div class="form-check">
                      <input
                        type="radio"
                        class="form-check-input"
                        id="radio1"
                        name="optradio"
                        value="KMP"
                        onClick={handleRadioButton}
                        checked
                      />
                      KMP
                      <label class="form-check-label" for="radio1"></label>
                    </div>
                    <div class="form-check">
                      <input
                        type="radio"
                        class="form-check-input"
                        id="radio2"
                        name="optradio"
                        value="BM"
                        onClick={handleRadioButton}
                      />
                      Boyer Moore
                      <label class="form-check-label" for="radio2"></label>
                    </div>
                  </div>
                </div>
                <div className={styles.submit}>
                  <div class="col-md-12 text-center">
                    <Button variant="success" type="submit">
                      <strong>Submit</strong>
                    </Button>
                  </div>
                </div>
              </Form>
            </div>
          </div>
          <div class="d-flex flex-column justify-content-center align-items-center">
            {isSubmitted && (
              <div className={styles.result}>
                <Card border="success" style={{ width: "18rem" }}>
                  <Card.Body>
                    {/* card title in center */}
                    <h3 className={styles.resultHeading}>Hasil</h3>
                    {/* card label and dummy */}
                    <Card.Text>
                      <div className={styles.resultCard}>
                        <div className={styles.resultFlex}>
                          <p className={styles.resultInfoL}>Date</p>
                          <p className={styles.resultInfo}>{dummyData.tanggal}</p>
                        </div>
                        <div className={styles.resultFlex}>
                          <p className={styles.resultInfoL}>Patient</p>
                          <p className={styles.resultInfo}>{dummyData.nama_pengguna}</p>
                        </div>
                        <div className={styles.resultFlex}>
                          <p className={styles.resultInfoL}>Disease</p>
                          <p className={styles.resultInfo}>
                            {dummyData.nama_penyakit}
                          </p>
                        </div>
                        <div className={styles.resultFlex}>
                          <p className={styles.resultInfoL}>Similarity</p>
                          <p className={styles.resultInfo}>
                            {dummyData.similarity}
                          </p>
                        </div>
                        <div className={styles.resultFlex}>
                          <p className={styles.resultInfoL}>Result</p>
                          <p className={styles.resultInfo}>
                            {dummyData.status_tes ? "True" : "False"}
                          </p>
                        </div>
                      </div>
                    </Card.Text>
                  </Card.Body>
                </Card>
              </div>
            )}
          </div>
        </div>
      </div>
    </>
  );
}
export default Tes;
