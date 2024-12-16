Your task is to implement an application that simulates an electronic sign. The sign should consist of a view with 6 * 36 pixels and should have a memory to be able to hold a sequence of views. The pixels can be either On or Off and input is done by specifying a sequence of the pixels to be On. The pixels are indicated by letter vertically and number horizontally, i.e. the pixel in the top left corner is called A0 and at the bottom right corner is called F35.

The interface the application should provide the following features:

* Enter a view as a sequence of pixels and save it in memory
* Print all views stored in memory
* Clear the memory

Exempel:

The following sequence:
A5A6A8A9A13A14A16A17A20A21A22A23A24A30B4B5B6B9B10B12B13B16B17B19B20B29B31C3C4C5C6C10C11C12C16C17C20C21C28C32D2D3D5D6D10D11D12DG561J1216D17D22D23D27D2GH1201245168D29D33E1E2E3E4E5E6E9E10E12E13E16E17E23E24E26E30E34F1F2F5F6F8F9F13F14F16F17F19F20F21F22F23F25F26F27F28F29F30F31F32F33F34F35

Should give the following result:
```
     ** **   ** **  *****     *    
    ***  ** **  ** **        * *   
   ****   ***   **  **      *   *  
  ** **   ***   **    **   ***   * 
 ******  ** **  **     ** *   *   *
**   ** **   ** ** ***** ***********
```

* You should also be able to enter a sequence of characters and numbers and get them represented as views. (A, B, C, 1, 2 and 3 is enough). These views should also be stored and displayed in the same way as the pixel sequence views above.

Exempel:

The following sequence:
ABC123

Should give the following result:
```
 ***  ****   ***    *    ***   *** 
*   * *   * *   *  **   *   * *   *
***** ****  *       *      *    ** 
*   * *   * *       *     *       *
*   * *   * *   *   *    *    *   *
*   * ****   ***   ***  *****  *** 
```