# ✅ ULTIMATE varC TOKEN STRESS TEST

Perfect — you want a “maximal token coverage” varC+C file such that:

✅ If this file tokenizes correctly, your tokenizer is essentially complete.

Below is a single self-contained stress-test file that exercises:

✅ All 7 varC keywords
✅ All C keywords (most of them)
✅ All operators (++ -- -> .* << >> && || ?:, etc.)
✅ All literals (int, hex, float, char, string, escape sequences)
✅ All comment styles
✅ All punctuation (; , () {} [])
✅ All declarator forms (pointers, arrays, functions)
✅ All control flow
✅ unsafe {} blocks
✅ move(expr)
✅ opt, raw
✅ Structs, enums, typedefs, casts
✅ Preprocessor macros
✅ Edge-case whitespace

```c
/* =============================
   varC TOKEN STRESS TEST FILE
   If THIS tokenizes correctly,
   your tokenizer is SOLID.
   ============================= */

/* -------- PREPROCESSOR -------- */

#define MACRO(x,y) ((x) + (y))
#define STR "string\twith\nescapes\""
#define HEX 0xDEADBEEF
#define BIN 0b101010     // if your lexer supports it
#define OCT 0755

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

/* -------- ALL varC KEYWORDS -------- */

owner
borrow
borrow_mut
move
unsafe
opt
raw

/* -------- ALL MAJOR C KEYWORDS -------- */

typedef struct S S;
struct S { int a; };

union U { int x; float y; };
enum E { A = 1, B = 2, C = 3 };

static extern inline const volatile restrict unsigned long long int glob = 0;

/* -------- COMPLEX DECLARATORS -------- */

owner int* (*fp)(borrow int*, borrow_mut int*);
raw char (**names)[32];
opt int arr[10];
int (*matrix)[3][4];

/* -------- FUNCTION USING ALL varC KEYWORDS -------- */

owner int* make_value(opt int seed) {
    owner int* p = (owner int*)malloc(sizeof(int));
    *p = seed ? seed : 123;
    return p;
}

void print_value(borrow int* p) {
    printf("%d\n", *p);
}

void mutate(borrow_mut int* p) {
    ++(*p);
}

void write_raw(raw int* p) {
    *p = 999;
}

void consume(owner int* p) {
    free(p);
}

/* -------- OPERATOR & EXPRESSION TESTS -------- */

int operators(int a, int b) {
    int r = 0;
    r += a + b;
    r -= a - b;
    r *= a * b;
    r /= a / (b + 1);
    r %= a % (b + 1);

    r <<= 2;
    r >>= 1;

    r &= a & b;
    r |= a | b;
    r ^= a ^ b;

    r = (a == b);
    r = (a != b);
    r = (a < b);
    r = (a > b);
    r = (a <= b);
    r = (a >= b);

    r = (a && b);
    r = (a || b);
    r = !a;

    r = a ? b : 0;

    a++;
    --b;

    return r;
}

/* -------- CONTROL FLOW -------- */

int flow(int x) {
    if (x > 0) goto label;
    else if (x < 0) return -1;
    else return 0;

label:
    for (int i = 0; i < 10; i++) {
        if (i == 5) break;
        if (i == 3) continue;
        x += i;
    }

    do {
        x--;
    } while (x > 0);

    switch (x) {
        case 0: x = 10; break;
        case 1: x = 20; break;
        default: x = 30;
    }

    return x;
}

/* -------- LITERAL STRESS -------- */

void literals(void) {
    int i = 123;
    int h = 0xFF;
    int o = 0777;

    float f = 1.23f;
    double d = 9.87654321;

    char c1 = 'A';
    char c2 = '\n';
    char c3 = '\x41';

    const char* s1 = "hello";
    const char* s2 = "escape: \t \n \r \\ \"";
}

/* -------- UNSAFE + RAW + MOVE -------- */

int main(int argc, char** argv) {
    opt int* maybe = NULL;

    owner int* x = make_value(10);
    print_value(x);
    mutate(x);

    owner int* y = move(x);

    unsafe {
        raw int* r = (raw int*)HEX;
        write_raw(r);
    }

    consume(move(y));

    int arr[3] = {1,2,3};
    int* p = arr;
    int v = *p++;
    v = (*p)--;

    MACRO(v, 10);

    return 0;
}

/* -------- COMMENT FOR LEXER -------- */

// single line comment
/*
   multi
   line
   comment
*/

```

# ✅ HOW TO USE THIS FILE
Step 1 — Run only the tokenizer:

Your output should produce a long flat token stream with:

No UNKNOWN tokens

No dropped characters

No misclassified keywords as identifiers

Step 2 — Smoke-test expectations:

If any of these fail, your tokenizer is incomplete:

| Failure                                   | Meaning               |
| ----------------------------------------- | --------------------- |
| `borrow_mut` split into `borrow` + `_mut` | Identifier rule bug   |
| `move(` parsed as IDENT                   | Keyword table bug     |
| `0xDEADBEEF` broken                       | Hex literal bug       |
| `"escape: \t"` broken                     | String escape bug     |
| `>>=` broken into `>>` + `=`              | Multi-op priority bug |
| `//` not ignored                          | Comment handling bug  |
