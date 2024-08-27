#include "cgo.h"

#include <map>
#include <vector>


#include <climits>

constexpr int clz(unsigned int v) {
    return __builtin_clz(v);
} 

constexpr int clz(unsigned long v) {
    return __builtin_clzl(v);
} 

constexpr int clz(unsigned long long v) {
    return __builtin_clzll(v);
} 

constexpr int firstbit(int v) {
    return (sizeof(v) << 3) - clz((unsigned int)v);
}

constexpr int firstbit(unsigned int v) {
    return (sizeof(v) << 3) - clz(v);
}

constexpr int firstbit(unsigned long v) {
    return (sizeof(v) << 3) - clz(v);
} 

constexpr int firstbit(unsigned long long v) {
    return (sizeof(v) << 3) - clz(v);
}

typedef int Comp1;
typedef long Comp2;
typedef unsigned int Comp3;

template <typename T>
struct arch_t {
    std::vector<T> comp;    
};

template <typename... Args>
struct archtype : arch_t<Args>... {
    template <typename T>
    arch_t<T>& a() { return *this; }
};

typedef archtype<Comp1, Comp2, Comp3> arch;

int id = 1;

struct e_id {
    int64_t id : firstbit(0xffffffff);
    int64_t e_id : firstbit(0xffff);
    int64_t flags : firstbit(0xf);
};

int a_id = 1;
struct entity_a {
    inline static arch e {};

    comp1_t *comp1;
};

entity_t* ecs_new_entity_a() {
    auto _id = e_id {
        id++,
        a_id++,
    };
    entity_a::e.a<Comp1>().comp.resize(_id.e_id);

    thread_local static entity_a ec_;
    ec_ = entity_a {
        &entity_a::e.a<Comp1>().comp[_id.e_id]
    };
    return (entity_t*)&ec_;
}
