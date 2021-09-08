# ReadMe

Ce programme analyse les surfaces dans un fichier .ply collecté par une caméra de profondeur ou un LiDAR d'INTEL.

## Les paramètres

Les paramètres se trouvent dans le fichier config.json

    InputAddr : L'adresse des fichiers de nuage de points, doit se terminer avec un "/"

    OutputAddr : L'adresse pour enregistrer les résultats de calcul, doit se terminer avec un "/"

    NumCore : Nombre de coeurs pour la parallélisation

    MaxGort : Nombre maximal de goroutines au même temps

    CompressLev : Niveau de compression pour donner à la fonction gzip. Utilise 0 pour ne pas compresser. Le niveau recommandé est 3, qui est optimal entre temps et taux de compression

Les paramètres pour la RANSAC : 

    MaxDistance : La distance maximale tolérée entre un point et un plan pour que le point soit considéré comme un inlierde ce plan. Ce paramètre est utilisé pour qualifier les inliersd’un plan. S’il est très grand, la qualité des plans ajustés peutbaisser. Dans le cas contraire, on risque d’augmenter le nombre d’itérations car il est plus difficile detrouver assez d’inliers. 

    MinScoreRANSAC : La proportion minimale d’inliersnécessaires parmi un batch pour qu’un plan soit considéré comme validé. C’est-à-dire si le nombre de points considérés comme inliersd’un plan est au-dessus de ce seuil, le plan sera validé. Avec le paramètre précédent, les deux formentun arbitrage entre la qualité des plans et le temps de calcul.

    MinVertexPlane : Le nombre minimal de points pour ajuster un plan.

    MaxAnglePlane : L’angle maximal toléré entre deux plans pour qu’ils soient considérés parallèles. Dans ce cas, si la distance est inférieure au paramètre ci-dessus, on fusionnera les deux plans. 

    MaxVertexQuit : Le nombre maximal des points qui n’appartiennent à aucun plan, pour continuer les itérations. S’il est inférieur au seuil donné, cela veut dire que le système est bien analysé et on pourra terminer le programme.

    MaxIteration : Le nombre maximal des itérations avant la fin du programme. Le programme s’arrête une fois que ceseuil estatteint. 

    NumBatch : Le nombre de points qui composent le sous-ensemble où on réalise le calcul

## Mode alternatif : SVD / Inverse de matrice et 32 / 64 bits

SVD / Inverse de matrice 

    La différence se trouvent en ajustement de plan avec un groupe de points. Le problème est ramené à résoudre un système superdéterminé par l'inverse de matrice ou par la décomposition de matrice. 

    Le benchmark montre que la méthode d'inverse (par défaut) est plus rapide.

    Pour utiliser la méthode de SVD : Installer la bibliothèque https://gogs.forclum.pro/mzalt/Implementation-SVD/src/master, puis remplacer la fonction PlaneMonoConsecRANSAC32 par PlaneMonoConsecRANSAC32SVD

32 / 64 bits

    La différence se trouve dans la taillé de données pendant les calculs. 

    Le benckmark montre que la version de 32 bits (par défaut) est plus rapide que celle de 64bits

    Pour utiliser la version de 64bits : Remplacer les fonction de 32bits et les variables en main.go par 64bits

## Configuration de l'environnement :

Pour installer le package Gonum : go get -u gonum.org/v1/gonum/mat

Pour exécuter le programme et obtenir un résultat : 


```sh
go build
```
```sh
./dataprocessing
```
