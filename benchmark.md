# Benchmark processing

Commande à utiliser : 
```
./dataprocessing -location=path -modeProc=0 -modeComp= //Nuages des point monochromes sans couleurs
```
```
./dataprocessing -location=path -modeProc=1 -modeComp= //Sans couleurs et utiisation de float32 au lieu de float64.
```
```
./dataprocessing -location=path -modeProc=2 -modeComp= //With a brand new .ply reader.
```

## 1. Benchmark sur X86 et sur Rasberry Pi

### Benchmark du dataprocessing

| Système                        | Kernel           | Processeur              | Stockage                                | Paramètres                               | Mode                                         | Volume | Temps           | Vitesse      |
| ------------------------------ | ---------------- | ----------------------- | --------------------------------------- | --------------------------------------- | ------------------------------------------------------------- | ------ | --------------- | ------------ |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8                | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 0     | 290Mb  | 19.15 secondes  | 15.14 Mb/s |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8                | calculator.PlaneMonoConsecRANSAC32(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 1    | 303Mb  | 20.22 secondes  | 14.99 Mb/s | 
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8                | calculator.PlaneMonoConsecRANSAC32(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 2    | 303Mb  | 3.811 secondes  | 79.51 Mb/s |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | Micro SD Verbatim 32GB class 10 SDHC U1 | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) |  | 230Mb  |  1m17,694s |   2.9603Mb/s |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | Micro SD Verbatim 32GB class 10 SDHC U1 | calculator.calculator.PlaneMonoConsecRANSAC32(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | With a brand new .ply reader. | 230Mb  |  55.793 |  4.122 Mb/s |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | Micro SD Verbatim 32GB class 10 SDHC U1 | calculator.calculator.PlaneMonoConsecRANSAC32(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | With a brand new .ply reader. | 230Mb  |  1m06sec |  3.4848 Mb/s |

### Benchmark du datacompressing

| Paramètre | Fonctionnement |
| --- | --- |
| 0 | No compression |
| -2 | Huffman only |
| 1 | Maximal speed |
| 2...8  | Intermediate levels |
| 9 | Minimal size after compression (Theoretical) |

| Système                        | Kernel           | Processeur              | Stockage                 | Niveau de compression                  | Volume avant compression | Volume après compression | Volume compressé | Temps  | Vitesse (Volume compressé / temps) | Taux d'usage du processeur |
| ------------------------------ | ---------------- | ----------------------- | ------------------------ | -------------------------------------- | ------------------------ | ------------------------ | ---------------- | ------ | ---------------------------------- | -------------------------- |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 0 (Sans compression, copier et coller) | 303Mb                    | 303Mb                    | 0Mb              | 0.34s  | 0Mb/s (BP : 891.18Mb/s)            | 2%                         |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | -2 (Huffman only)                      | 303Mb                    | 222Mb                    | 81Mb             | 3.855s | 21.01Mb/s                          | 15%-17%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 1 (Vitesse optimale)                   | 303Mb                    | 129Mb                    | 174Mb            | 4.807s | 36.19Mb/s                          | 15%-17%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 2                                      | 303Mb                    | 134Mb                    | 169Mb            | 7.243s | 23.33Mb/s                          | 15%-18%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 3                                      | 303Mb                    | 118Mb                    | 185Mb            | 5.972s | 30.98Mb/s                          | 15%-18%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 4                                      | 303Mb                    | 117Mb                    | 186Mb            | 6.290s | 29.57Mb/s                          | 15%-18%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 5                                      | 303Mb                    | 127Mb                    | 176Mb            | 8.526s | 20.64Mb/s                          | 16%-19%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 6                                      | 303Mb                    | 127Mb                    | 176Mb            | 7.988s | 22.03Mb/s                          | 15%-18%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 7                                      | 303Mb                    | 127Mb                    | 176Mb            | 8.848s | 19.89Mb/s                          | 15%-17%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 8                                      | 303Mb                    | 127Mb                    | 176Mb            | 7.813s | 22.52Mb/s                          | 15%-18%                    |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8 | 9                                      | 303Mb                    | 127Mb                    | 176Mb            | 8.349s | 21.08Mb/s                          | 16%-19%                    |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 0 (Sans compression, copier et coller) | 207Mb                    | 207Mb                    | 0Mb              | 2.04s  | 0Mb/s (BP : 422  Mb/s)             | 45%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | -2 (Huffman only)                      | 207Mb                    | 150Mb                    | 57Mb             | 13.50s | 4,2Mb/s                            | 32%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 1 (Vitesse optimale)                   | 207Mb                    | 84Mb                     | 123Mb            | 16.06s | 7,6Mb/s                            | 32%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 2                                      | 207Mb                    | 86Mb                     | 121Mb            | 29.79s | 4,06Mb/s                           | 32%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 3                                      | 207Mb                    | 80Mb                     | 127Mb            | 28.05s | 4,5Mb/s                            | 31%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 4                                      | 207Mb                    | 80Mb                     | 127Mb            | 31.08s | 4,08Mb/s                           | 31%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 5                                      | 207Mb                    | 85Mb                     | 122Mb            | 39.18s | 3,11Mb/s                           | 32%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 6                                      | 207Mb                    | 85Mb                     | 122Mb            | 38.85s | 3,14Mb/s                           | 32%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 7                                      | 207Mb                    | 85Mb                     | 122Mb            | 37.54s | 3,24Mb/s                           | 31%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 8                                      | 207Mb                    | 85Mb                     | 122Mb            | 37.97s | 3,21Mb/s                           | 31%                        |
| Raspbian GNU/Linux 10 (buster) | 5.10.17-v7l+     | ARMv7 rev3 (v7l)        | SSD SATA 500gb 860 Evo   | 9                                      | 207Mb                    | 85Mb                     | 122Mb            | 40.45s | 3,02Mb/s                           | 31%                        |


### Benchmark du dataprocessing avec la compression des données intégrée

| Système                        | Kernel           | Processeur              | Stockage                                | Paramètres                               | Mode de traitement | Mode de compression                      | Volume | Temps           | Vitesse totale |
| ------------------------------ | ---------------- | ----------------------- | --------------------------------------- | --------------------------------------- | ------------------------------------------------------------- | ------ | ------ | --------------- | ------------ |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8                | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 1 | 1   | 303Mb  | 22.82 s  | 13.28 Mb/s |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8                | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 2 | 1   | 303Mb  | 4.15 s  | 73.01 Mb/s |
| Ubuntu 20.04.2 LTS             | 5.4.0-70-generic | i5-1035G7 CPU @ 1.20GHz | NVMe Micron CT1000P1SSD8                | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 1 | 3   | 303Mb  | 24.83 s  | 12.21 Mb/s |
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8    | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 2 | 1   | 303Mb  | 2431 s  |  |

## 2. Benchmark sur une machine virtuelle. Utilise un jeu de données <simple> à cause de la performance limitée

Jeu de données utilisé : https://groupeeiffage-my.sharepoint.com/:u:/g/personal/zhiyuan_liu_eiffage_com/EXJzQPP62H9PpFLQfXfcjdYBkuhC6PXHUuJ1rQRos2WwxQ?e=WCat9p

Ici on fait un nouveau benchmark sur la machine virtuelle qui émule un Rasberry Pi sur la machine x86, en utilisant les même codes par rapport ce qu'on utilise sur le vrai PI. Vu que la performance de la machine virtuelle est très limitée, on lance ce benchmark avec 10 fichiers .ply dont la structure est simple (i.e. les points sont sur la même surface) au lieu de 100 fichiers.

### Benchmark du dataprocessing

| Système                        | Kernel           | Processeur              | Stockage                                | Paramètres                               | Mode                                         | Volume | Temps           | Vitesse      |
| ------------------------------ | ---------------- | ----------------------- | --------------------------------------- | --------------------------------------- | ------------------------------------------------------------- | ------ | --------------- | ------------ |
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8     | calculator.PlaneMonoConsecRANSAC32(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 1    | 35 Mb  | 338 s  | 0.104 Mb/s | 
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8     | calculator.PlaneMonoConsecRANSAC32(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 2    | 35 Mb  | 87 s  | 0.402 Mb/s | 

### Benchmark du datacompressing (supression intégrée)

| Système                        | Kernel           | Processeur              | Stockage                 | Niveau de compression                  | Volume avant compression | Volume après compression | Volume compressé | Temps  | Vitesse (Volume compressé / temps) | Taux d'usage du processeur |
| ------------------------------ | ---------------- | ----------------------- | ------------------------ | -------------------------------------- | ------------------------ | ------------------------ | ---------------- | ------ | ---------------------------------- | -------------------------- |
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8    | 0  | 35Mb                    | 35Mb                    | 0Mb              | 1.88 s  | 0 Mb/s (BP : 18.62 Mb/s)            | 80 %                         |
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8    | 1  | 35Mb                    | 13Mb                    | 22Mb              | 13.13 s  | 1.68 Mb/s       | 100 %                         |
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8    | 3  | 35Mb                    | 12Mb                    | 23Mb              | 16.83 s  | 1.37 Mb/s       | 100 %                       |

### Benchmark du dataprocessing avec la compression des données intégrée

| Système                        | Kernel           | Processeur              | Stockage                                | Paramètres                               | Mode de traitement | Mode de compression                      | Volume | Temps           | Vitesse totale |
| ------------------------------ | ---------------- | ----------------------- | --------------------------------------- | --------------------------------------- | ------------------------------------------------------------- | ------ | ------ | --------------- | ------------ |
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8               | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 2 | 3   | 35Mb  | 105 s  | 0.33 Mb/s |

### Benchmark du dataprocessing avec la compression des données intégrée, nombre de goroutines lancés au même temps limité à 2

| Système                        | Kernel           | Processeur              | Stockage                                | Paramètres                               | Mode de traitement | Mode de compression                      | Volume | Temps           | Vitesse totale |
| ------------------------------ | ---------------- | ----------------------- | --------------------------------------- | --------------------------------------- | ------------------------------------------------------------- | ------ | ------ | --------------- | ------------ |
| Linux raspberrypi (VM) |  4.4.34+  | ARM1176 | Carte SD émulée sur NVMe Micron CT1000P1SSD8               | calculator.PlaneMonoConsecRANSAC(vlist, 0.02, 0.1, 500, 0.1, 6000, 2000, 500) | 2 | 3   | 35Mb  | 106.5 s  | 0.33 Mb/s |
