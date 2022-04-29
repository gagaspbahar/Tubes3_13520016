-- MariaDB dump 10.19  Distrib 10.6.5-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: tubes3_basdat
-- ------------------------------------------------------
-- Server version	10.6.5-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `tubes3_basdat`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `tubes3_basdat` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `tubes3_basdat`;

--
-- Table structure for table `data_uji`
--

DROP TABLE IF EXISTS `data_uji`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `data_uji` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `tanggal_tes` date NOT NULL,
  `nama_pengguna` varchar(255) CHARACTER SET utf8mb3 NOT NULL,
  `nama_penyakit` char(50) CHARACTER SET utf8mb3 NOT NULL,
  `similarity` float NOT NULL DEFAULT 0,
  `status_tes` tinyint(4) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK_data_uji_sequence_penyakit` (`nama_penyakit`),
  CONSTRAINT `FK_data_uji_sequence_penyakit` FOREIGN KEY (`nama_penyakit`) REFERENCES `sequence_penyakit` (`penyakit`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_uji`
--

LOCK TABLES `data_uji` WRITE;
/*!40000 ALTER TABLE `data_uji` DISABLE KEYS */;
INSERT INTO `data_uji` (`ID`, `tanggal_tes`, `nama_pengguna`, `nama_penyakit`, `similarity`, `status_tes`) VALUES (2,'2022-04-24','aira','kista',100,1),(3,'2022-04-24','thalca','malaria',0.315789,0),(4,'2022-04-28','Aira Thalca','kista',100,1),(5,'2022-04-28','Fikron','kista',100,1),(6,'2022-04-28','Fikron','kista',100,1),(7,'2022-04-28','Fikron','kista',100,1),(8,'2022-04-28','Aira Thalca','kista',100,1),(9,'2022-04-28','Fikron','kista',100,1),(10,'2022-04-28','Fikron','kista',100,1),(11,'2022-04-28','Fikron','kista',100,1),(12,'2022-04-28','Fikron','kista',100,1),(13,'2022-04-28','Fikron','kista',100,1),(14,'2022-04-28','Fikron','kista',100,1),(15,'2022-04-28','Fikron','kista',100,1),(16,'2022-04-28','Fikron','kista',100,1),(17,'2022-04-28','Fikron','kista',100,1),(18,'2022-04-28','Fikron','kista',100,1),(19,'2022-04-28','Fikron','kista',100,1),(20,'2022-04-28','Fikron','kista',100,1),(21,'2022-04-28','Fikron','kista',100,1),(22,'2022-04-28','Fikron','kista',100,1),(23,'2022-04-28','Fikron','kista',100,1),(24,'2022-04-29','Gagas','sakithati',100,1),(25,'2022-04-29','Gagas','stress',35,0);
/*!40000 ALTER TABLE `data_uji` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sequence_penyakit`
--

DROP TABLE IF EXISTS `sequence_penyakit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sequence_penyakit` (
  `penyakit` char(50) CHARACTER SET utf8mb3 NOT NULL,
  `sequence` varchar(1000) CHARACTER SET utf8mb3 DEFAULT NULL,
  PRIMARY KEY (`penyakit`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sequence_penyakit`
--

LOCK TABLES `sequence_penyakit` WRITE;
/*!40000 ALTER TABLE `sequence_penyakit` DISABLE KEYS */;
INSERT INTO `sequence_penyakit` (`penyakit`, `sequence`) VALUES ('HIV','CTAGCTGATCGATGAT'),('kista','CGATCAGCATGACGTC'),('malaria','ATGCTGACATACAGCATCG'),('sakithati','AGTCAGTCGTAGTCGATCCCGTAG'),('stress','AGCTATGCTAAGTCGATAGC');
/*!40000 ALTER TABLE `sequence_penyakit` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-29 14:29:13
