# **varC: Memory-Safe C with Rust Integration**

## **Project Aim**

`varC` is a modern pre-compiler project designed to **combine the performance and low-level control of C with the safety features of Rust**. Its primary goals are:

1. **Type Safety:**
   Enhance traditional C by adding ownership, borrowing, and safe memory management concepts inspired by Rust, reducing common bugs like use-after-free, dangling pointers, and data races.

2. **Seamless Integration with C:**
   varC code remains compatible with standard C code. You can integrate legacy C libraries and systems code while benefiting from enhanced type safety.

3. **Rust-like Ownership Model:**

   * `owner` → Unique ownership of memory.
   * `borrow` → Temporary non-mutating access.
   * `borrow_mut` → Temporary mutable access.
   * `move` → Transfer ownership.
     These concepts help prevent memory misuse at compile time.

4. **Extended C Syntax Support:**
   varC preserves all standard C99 features (types, pointers, structures, unions, enums, control flow, operators) and adds its own extensions for safer memory management.

5. **Intermediate Representation via AST:**
   The compiler tokenizes and parses varC code into a structured **Abstract Syntax Tree (AST)**, which allows future stages like optimization, code generation, or Rust-like safety checks.

---

## **Key Features**

* Fully **C99-compliant** syntax with additional safety extensions.
* **Memory safety** features inspired by Rust.
* Preprocessor support (`#define`, `#include`, etc.).
* Supports pointers, arrays, function pointers, and complex declarators.
* Lexer and parser produce **JSON AST**, enabling:

  * Debugging
  * AST-based tooling
  * Future code generation or transpilation to safe C or Rust.

---

## **Example varC Code**

```c
owner int* make_value(opt int seed) {
    owner int* p = (owner int*)malloc(sizeof(int));
    *p = seed ? seed : 123;
    return p;
}

void print_value(borrow int* p) {
    printf("%d\n", *p);
}
```

* `owner` ensures `p` is uniquely owned.
* `borrow` provides read-only temporary access without taking ownership.
* `opt` marks an optional argument (can be `NULL`).

---

## **Building the Compiler**

```bash
# Build the compiler
go build -o ./build/varC ./cmd/varC
```

This compiles the Go-based varC compiler into a binary called `varC` inside the `./build` directory.

---

## **Running the Compiler**

```bash
./build/varC --file ./examples/example.varc
```

* `--file` points to the varC source code file.
```

---

## **Future Goals**

* Optimization passes for **performance and safety**.
* IDE tooling support via **AST JSON**.

---

## **Conclusion**

varC is a **type-safe C language extension** with Rust-inspired memory safety. It allows developers to write C-style code **without sacrificing safety**, combining the best of both worlds: the **speed of C** and the **safety of Rust**.

This project is ideal for **systems programming, embedded development, and safe low-level code**.

---

If you want, I can also **write a visual diagram of how varC works internally**—lexer → parser → AST → memory safety checks—which will make this Markdown even stronger.

Do you want me to add that diagram?
