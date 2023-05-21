# Guía: Árboles Binario de Búsqueda Balanceados

## Ejercicios

1. En lápiz y papel dibujar cada paso.

   Indicar a cada paso si se realizan rotaciones.

   1. Crear un árbol AVL vacío e insertar: 10, 50, 30, 40, 9, 8, 11, 12, 13, 14, 15. Indicando a cada paso si se realizan rotaciones.

   2. Crear un árbol AVL vacío e insertar: 100, 29, 71, 82, 48, 39, 101, 22, 46, 17, 3, 20, 25, 10. Se elimina el 100 y a continuación el 29.

2. Dado el siguiente árbol AVL, borrar los nodos: 60, 55, 50, 40, 70, 20. Indicar los pasos y las rotaciones realizadas, en caso de que sean necesarias.

   ![Árbol AVL](imagenes/arbol-avl.drawio.svg)

3. Escribir un método en [árbol binario de búsqueda](binarytree/binarynode.go) con el nombre `EsAVL` que devuelva un valor booleano indicando si el árbol es AVL.

4. Implementar en [árbol AVL](avl/avltreenode.go) las operaciones de inserción, eliminación y búsqueda.

5. Escribir un iterador PreOrder.

   > Pista: Al crear el iterador apilar la raiz
   >
   > Al devolver el siguiente desapilar y devolver el elemento del tope. Si ese nodo tiene hijo izquierdo apilarlo y si tiene hijo derecho apilarlo

6. Escribir un iterador PosOrder.

7. Utilizando como estructura auxiliar un árbol AVL, inplementar el TAD TreeSet

8. Utilizando como estructura auxiliar un árbol AVL, inplementar el algoritmo de ordenamiento `AVLTreeSort`. Analizar la complejidad computacional. ¿Le parece un algoritmo eficiente? ¿Por qué?
