<!-- rope-ladder:begin document 865826b57fd983e2508b080e747ce4e939861e03df0d10942e3c4b59cfda22f8 -->
# Implementation-language concepts

This is a guided reading path, not a dictionary. Read topics in order: each section explains the problem first, names the ideas it builds on, then walks through a small piece of the codebase.

<a id="topic-cpp-engine-polymorphism"></a>
## 1. C++ inheritance, virtual interfaces, and Ref ownership

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Objective-C++ Apple platform adapters](#topic-objective-cpp-apple-adapters).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Core abstractions such as resources, input events, and resource format loaders are C++ class hierarchies.

C++ classes use inheritance, virtual methods, Ref ownership, and static registries to implement core services.

### How the implementation applies it

#### ResourceFormatLoader::load

ResourceFormatLoader derives from RefCounted and declares virtual loading, recognition, dependency, and type-query operations that concrete format loaders override.

#### ResourceLoader::ThreadLoadTask

The loader keeps task, path, status, cache-mode, resource, dependency, and awaiting state in its nested threaded-load record.

#### InputEvent hierarchy

InputEvent derives from Resource, and InputEventFromWindow, InputEventWithModifiers, key, mouse, joypad, touch, gesture, MIDI, and action types extend that base.

### Walk through a small source example

```c
class InputEvent : public Resource {
```

Source: `core/input/input_event.h` — InputEvent

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource, InputEvent, GDExtension

Evidence:
- Code: `core/io/resource_loader.h:48` — ResourceFormatLoader
- Code: `core/io/resource_loader.h` — ResourceLoader::ThreadLoadTask
- Code: `core/input/input_event.h:52` — InputEvent
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C resource lifecycles and ownership](#topic-c-resource-lifecycles)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)

<a id="topic-cpp-class-registration"></a>
## 2. C++ runtime class registration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Godot's native API surface is assembled through ClassDB rather than by a separate hand-written runtime schema.

C++ templates and classes implement the runtime registry that records native class identity, inheritance, construction, methods, properties, signals, and constants.

### How the implementation applies it

#### Class registration

ClassDB::register_class initializes the type metadata, assigns a creator function, marks the class exposed, and records the current API category.

#### Method binding

ClassDB::bind_method builds a MethodBind from a C++ member function and stores it under an exposed method name.

#### Property and signal metadata

ClassDB::add_property and ClassDB::add_signal populate the reflection data later consumed by the engine API.

### Walk through a small source example

```c
class ClassDB {
	friend class Object;
	friend class GDType;

public:
	enum APIType {
		API_CORE,
		API_EDITOR,
		API_EXTENSION,
		API_EDITOR_EXTENSION,
		API_NONE
	};
```

Source: `core/object/class_db.h` — ClassDB (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource, RID

Evidence:
- Code: `core/object/class_db.h:62` — ClassDB
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C ABI bridging](#topic-c-abi-bridging)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [CMake native source indexing](#topic-cmake-native-source-index)
- [C++ enumerated export state](#topic-cplusplus-enumerated-export-state)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ binary data codecs](#topic-cpp-binary-data-codecs)
- [C++ byte-exact binary parsing](#topic-cpp-byte-exact-binary-parsing)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ mapping of C API types](#topic-cpp-c-type-mapping)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ const-reference accessors](#topic-cpp-const-reference-accessors)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ exception containment at ABI boundaries](#topic-cpp-exception-abi-boundaries)
- [C++ frame-time arithmetic](#topic-cpp-frame-scheduling)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ engine binding macros](#topic-cpp-godot-binding-macros)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ JNI Variant marshalling](#topic-cpp-jni-variant-marshalling)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ plugin specialization](#topic-cpp-plugin-specialization)
- [C++ polymorphic backend classes](#topic-cpp-polymorphic-backends)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ resources and polymorphic engine objects](#topic-cpp-resource-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ inheritance and reference-counted adapters](#topic-cpp-runtime-adapters)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ scope-bound locking](#topic-cpp-scope-locking)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ tagged storage and casts](#topic-cpp-tagged-storage)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ typed API records](#topic-cpp-typed-api-records)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ atomic synchronization](#topic-cxx-atomics)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ lambdas and callback objects](#topic-cxx-closures)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [JavaScript browser FFI](#topic-javascript-browser-ffi)
- [Objective-C++ Apple platform adapters](#topic-objective-cpp-apple-adapters)
- [Objective-C++ iOS adapters](#topic-objective-cpp-ios-adapters)
- [Python build source generation](#topic-python-build-source-generation)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)

<a id="topic-cpp-templates-traits"></a>
## 3. C++ templates and traits

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling), [C++ variadic binding](#topic-cpp-variadic-binding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Generic core containers and binding helpers are primarily header-defined template code.

The implementation uses C++ templates and trait specializations to adapt behavior to types, including zero construction, hashing, tuple recursion, and argument conversion.

### How the implementation applies it

#### Zero-construction trait

AABB specializes is_zero_constructible as true, providing a type-level property consumed by generic code.

#### Hash capability detection

Hash functions define has_hash_method and a HashMapHasherDefault implementation selected through template parameters.

#### Recursive tuple representation

Tuple recursively inherits Tuple<Rest...>, with TupleGet specializations for indexed access.

### Walk through a small source example

```c
struct is_zero_constructible<AABB> : std::true_type {};
```

Source: `core/math/aabb.h` — is_zero_constructible<AABB>

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

AABB

Evidence:
- Code: `core/math/aabb.h:515` — is_zero_constructible<AABB>
- Code: `core/templates/hashfuncs.h:164` — has_hash_method
- Code: `core/templates/tuple.h:33` — Tuple
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)

<a id="topic-python-build-code-generation"></a>
## 4. Python build code generation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is build-time support rather than the engine's C++ runtime.

The Python build helper defines generation and build-entry functions for virtual-method source generation.

### How the implementation applies it

#### Generation function

generate_version accepts argument-count and compatibility switches for emitted virtual-method variants.

#### Build entry function

run receives target, source, and environment parameters for build integration.

### Walk through a small source example

```python
def generate_version(argcount, const=False, returns=False, required=False, compat=False):
```

Source: `core/object/make_virtuals.py` — generate_version

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/object/make_virtuals.py:67` — generate_version
- Code: `core/object/make_virtuals.py:193` — run
- External (official, verified): [The Python Language Reference — Function definitions](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C# partial types and source generation](#topic-csharp-source-generation)
- [Python build-action functions](#topic-python-build-actions)
- [Python build hooks](#topic-python-build-hooks)
- [Python build scripts](#topic-python-build-scripts)
- [Python build source generation](#topic-python-build-source-generation)
- [Python SCons build hooks](#topic-python-scons-build-hooks)
- [Python SCons configuration functions](#topic-python-scons-configuration)
- [Python SCons module configuration](#topic-python-scons-module-configuration)

<a id="topic-cpp-object-hierarchies"></a>
## 5. C++ class inheritance

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ static casts across track types](#topic-cpp-runtime-casts).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The source uses base-class declarations to organize runtime, GUI, and asset types.

C++ class declarations define the repository's object families through public base classes, including Node-derived runtime objects, Resource-derived assets, and Viewport-derived windows.

### How the implementation applies it

#### Node root

Node is declared as an Object subclass, establishing the root of the scene runtime family.

#### Viewport specialization

Window publicly derives from Viewport, so window behavior is represented as a viewport specialization.

#### Resource specialization

PackedScene publicly derives from Resource, placing serialized scenes in the reusable asset family.

### Walk through a small source example

```c
class Node : public Object {
```

Source: `scene/main/node.h` — class Node

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node, PackedScene

Evidence:
- Code: `scene/main/node.h:54` — class Node : public Object
- Code: `scene/main/window.h:42` — class Window : public Viewport
- Code: `scene/resources/packed_scene.h:248` — class PackedScene : public Resource
- External (primary, verified): [Working Draft, Programming Languages — C++: Derived classes](https://eel.is/c++draft/class.derived), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Java Android host APIs](#topic-java-android-host-api)

<a id="topic-cpp-class-inheritance"></a>
## 6. C++ class inheritance for backend contracts

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code uses public derived classes to specialize a common engine-facing interface.

C++ class inheritance expresses shared renderer and server contracts and allows concrete dummy, extension, wrapper, clustered, and mobile implementations.

### How the implementation applies it

#### Physics backend substitution

PhysicsServer3DDummy, PhysicsServer3DExtension, and PhysicsServer3DWrapMT derive from PhysicsServer3D.

#### Renderer specialization

RenderForwardClustered and RenderForwardMobile derive from RendererSceneRenderRD.

#### Text and XR extension surfaces

TextServerExtension derives from TextServer and XRInterfaceExtension derives from XRInterface.

### Walk through a small source example

```c
class PhysicsServer3DDummy : public PhysicsServer3D {
```

Source: `servers/physics_3d/physics_server_3d_dummy.h` — PhysicsServer3DDummy

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `servers/physics_3d/physics_server_3d_dummy.h:131` — PhysicsServer3DDummy
- Code: `servers/rendering/renderer_rd/forward_clustered/render_forward_clustered.h:60` — RenderForwardClustered
- Code: `servers/rendering/renderer_rd/forward_mobile/render_forward_mobile.h:42` — RenderForwardMobile
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4928.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Java Android host APIs](#topic-java-android-host-api)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ class inheritance](#topic-cpp-object-hierarchies)

<a id="topic-cpp-const-borrowing"></a>
## 7. C++ const references and pointers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is a pervasive access pattern in rendering, scene, GUI, and resource code.

C++ const references and pointers expose existing repository objects and buffers without copying them, such as shader parameter names, vectors, arrays, and geometry data.

### How the implementation applies it

#### Name lookup return pointer

CanvasItem returns a pointer to an existing StringName remapping entry.

#### Connection list borrow

GraphEdit returns its connection collection by const reference.

#### Buffer access

Animation and geometry code obtains typed pointers from packed container storage through ptr().

### Walk through a small source example

```cpp
const StringName *CanvasItem::_instance_shader_parameter_get_remap(const StringName &p_name) const {
```

Source: `scene/main/canvas_item.cpp` — CanvasItem::_instance_shader_parameter_get_remap

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Animation, GraphEditConnection

Evidence:
- Code: `scene/main/canvas_item.cpp:672` — CanvasItem::_instance_shader_parameter_get_remap
- Code: `scene/gui/graph_edit.cpp:338` — GraphEdit::get_connections
- Code: `scene/resources/animation.cpp` — typed ptr() accesses for animation tracks and compression
- External (primary, verified): [Working Draft, Programming Languages — C++: References](https://eel.is/c++draft/dcl.ref), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ const-reference accessors](#topic-cpp-const-reference-accessors)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-cpp-interface-polymorphism"></a>
## 8. C++ interface polymorphism

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The public extension classes mirror native polymorphic interfaces and distinguish required callbacks from optional ones.

C++ base classes and virtual callbacks define extension contracts whose subclasses provide physics, rendering, stream, and text implementations.

### How the implementation applies it

#### RenderingDevice inheritance

RenderingDevice derives from RenderingDeviceCommons and uses Godot class metadata macros to participate in the native object model.

#### Physics virtual callbacks

PhysicsServer2DExtension declares virtual required operations such as _area_add_shape for external physics implementations.

#### Rendering buffer configuration

RenderSceneBuffersExtension exposes _configure so an extension handles viewport reconfiguration.

### Walk through a small source example

```c
class RenderingDevice : public RenderingDeviceCommons {
	GDCLASS(RenderingDevice, Object)

	_THREAD_SAFE_CLASS_

private:
	Thread::ID render_thread_id;

public:
	typedef int64_t DrawListID;
	typedef int64_t ComputeListID;
	typedef int64_t RaytracingListID;
```

Source: `servers/rendering/rendering_device.h` — RenderingDevice (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RenderSceneBuffersConfiguration, RID

Evidence:
- Code: `servers/rendering/rendering_device.h:67` — RenderingDevice
- Code: `doc/classes/PhysicsServer2DExtension.xml` — PhysicsServer2DExtension._area_add_shape
- Code: `doc/classes/RenderSceneBuffersExtension.xml` — RenderSceneBuffersExtension._configure
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ virtual interfaces](#topic-cpp-runtime-polymorphism)
- [C++ virtual dispatch](#topic-cpp-virtual-dispatch)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Java Android host APIs](#topic-java-android-host-api)
- [Python build code generation](#topic-python-build-code-generation)

<a id="topic-cpp-native-classes"></a>
## 9. C++ native class hierarchy

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ template-based method binding](#topic-cpp-template-binding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The engine’s native implementation expresses its core object, scene, resource, extension, and scripting types as C++ classes, then exposes selected APIs through ClassDB.

Native components use C++ classes, inheritance, and virtual functions to implement engine types such as Resource, Node, and GDScript.

### How the implementation applies it

#### Resource inheritance

Resource derives from RefCounted, and PackedScene derives from Resource, establishing the resource specialization chain.

#### Virtual subsystem APIs

Mesh declares virtual surface and material operations so concrete mesh types can supply their implementations.

#### Native registration

ClassDB registration stores creation functions and exposure metadata for registered concrete classes.

### Walk through a small source example

```c
class PackedScene : public Resource {
	GDCLASS(PackedScene, Resource);
```

Source: `scene/resources/packed_scene.h` — PackedScene

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource, Node, PackedScene, GDScript

Evidence:
- Code: `scene/resources/packed_scene.h:36` — PackedScene
- Code: `core/object/class_db.h:570` — ClassDB::register_class
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C ABI bridging](#topic-c-abi-bridging)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [CMake native source indexing](#topic-cmake-native-source-index)
- [C++ enumerated export state](#topic-cplusplus-enumerated-export-state)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ binary data codecs](#topic-cpp-binary-data-codecs)
- [C++ byte-exact binary parsing](#topic-cpp-byte-exact-binary-parsing)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ mapping of C API types](#topic-cpp-c-type-mapping)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ const-reference accessors](#topic-cpp-const-reference-accessors)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ exception containment at ABI boundaries](#topic-cpp-exception-abi-boundaries)
- [C++ frame-time arithmetic](#topic-cpp-frame-scheduling)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ engine binding macros](#topic-cpp-godot-binding-macros)
- [C++ JNI Variant marshalling](#topic-cpp-jni-variant-marshalling)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ plugin specialization](#topic-cpp-plugin-specialization)
- [C++ polymorphic backend classes](#topic-cpp-polymorphic-backends)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ resources and polymorphic engine objects](#topic-cpp-resource-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ inheritance and reference-counted adapters](#topic-cpp-runtime-adapters)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ scope-bound locking](#topic-cpp-scope-locking)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ tagged storage and casts](#topic-cpp-tagged-storage)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ typed API records](#topic-cpp-typed-api-records)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ atomic synchronization](#topic-cxx-atomics)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ lambdas and callback objects](#topic-cxx-closures)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [Java Android host APIs](#topic-java-android-host-api)
- [JavaScript browser FFI](#topic-javascript-browser-ffi)
- [Objective-C++ Apple platform adapters](#topic-objective-cpp-apple-adapters)
- [Objective-C++ iOS adapters](#topic-objective-cpp-ios-adapters)
- [Python build source generation](#topic-python-build-source-generation)
- [C++ virtual interfaces](#topic-cpp-runtime-polymorphism)
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [GDScript query-object API use](#topic-gdscript-query-objects)

<a id="topic-cpp-templates-and-views"></a>
## 10. C++ templates and non-owning views

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The rendering code combines engine containers with pointer-and-length style views for recorded GPU work.

Template-based containers and views expose typed data without copying it, including VectorView access to command and resource arrays.

### How the implementation applies it

#### Typed view access

VectorView stores a templated pointer and returns a const reference from indexed access.

#### Typed conversion helpers

VariantConverterStd140 specializations and conversion helpers translate variant values and arrays into GPU uniform layouts.

### Walk through a small source example

```c
template <typename T>
class VectorView {
	const T *_ptr = nullptr;
	uint32_t _size = 0;

public:
	const T &operator[](uint32_t p_index) const {
		DEV_ASSERT(p_index < _size);
		return _ptr[p_index];
	}

	_ALWAYS_INLINE_ const T *ptr() const { return _ptr; }
	_ALWAYS_INLINE_ uint32_t size() const { return _size; }
```

Source: `servers/rendering/rendering_device_commons.h` — VectorView (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RecordedCommand

Evidence:
- Code: `servers/rendering/rendering_device_commons.h:40` — VectorView
- Code: `servers/rendering/storage/variant_converters.h:41` — VariantConverterStd140
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4928.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ templates](#topic-cpp-templates)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ variadic binding](#topic-cpp-variadic-binding)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)

<a id="topic-cpp-generic-containers"></a>
## 11. C++ templates and typed containers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Godot-specific Vector and Ref types are used as C++ template applications throughout this partition.

C++ template applications express typed containers and handles, including Vector collections of Ref-managed connections and parameter-pack-sized argument arrays.

### How the implementation applies it

#### Typed graph connection collection

GraphEdit returns a Vector whose element type is Ref<GraphEdit::Connection>.

#### Variadic call storage

Node header code sizes an array from sizeof...(p_args), showing a parameter-pack-based helper.

#### Typed scene-state stack

Property utilities pass a Vector<SceneState::PackState> through their comparison and packing logic.

### Walk through a small source example

```cpp
const Vector<Ref<GraphEdit::Connection>> &GraphEdit::get_connections() const {
```

Source: `scene/gui/graph_edit.cpp` — GraphEdit::get_connections

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GraphEditConnection, SceneState

Evidence:
- Code: `scene/gui/graph_edit.cpp:338` — GraphEdit::get_connections
- Code: `scene/main/node.h` — sizeof...(p_args) argument-pointer arrays
- Code: `scene/property_utils.cpp` — Vector<SceneState::PackState> states_stack
- External (primary, verified): [Working Draft, Programming Languages — C++: Templates](https://eel.is/c++draft/temp), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ variadic binding](#topic-cpp-variadic-binding)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)

<a id="topic-glsl-compute-workgroups"></a>
## 12. GLSL compute workgroups and synchronization

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied shaders use compute passes for velocity, luminance, resolve, roughness, hierarchy, skinning, and postprocessing work.

Compute shaders use declared resource interfaces, workgroup identifiers, bounds checks, shared arrays, and memory barriers for image-processing operations.

### How the implementation applies it

#### Per-pixel velocity dispatch

motion_vectors_store.glsl rejects out-of-bounds invocations, derives a pixel coordinate from gl_GlobalInvocationID, and writes velocity to an image.

#### Workgroup reduction

luminance_reduce.glsl allocates shared temporary luminance storage and calls groupMemoryBarrier before reduction work.

### Walk through a small source example

```text
void main() {
	uint t = gl_LocalInvocationID.y * BLOCK_SIZE + gl_LocalInvocationID.x;
	ivec2 pos = ivec2(gl_GlobalInvocationID.xy);

	if (any(lessThan(pos, params.source_size))) {
#ifdef READ_TEXTURE
		vec3 v = texelFetch(source_texture, pos, 0).rgb;
		tmp_data[t] = max(v.r, max(v.g, v.b));
#else
		tmp_data[t] = imageLoad(source_luminance, pos).r;
#endif
	} else {
		tmp_data[t] = 0.0;
```

Source: `servers/rendering/renderer_rd/shaders/effects/luminance_reduce.glsl` — main (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `servers/rendering/renderer_rd/shaders/effects/motion_vectors_store.glsl:20` — main
- Code: `servers/rendering/renderer_rd/shaders/effects/luminance_reduce.glsl:15` — main
- Code: `servers/rendering/renderer_rd/shaders/effects/resolve.glsl:35` — main
- External (official, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.html), accessed 2026-07-15

<a id="topic-python-build-scripts"></a>
## 13. Python build scripts

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The inventory contains Python builders beside SCsub files that import them.

Python build scripts import builder modules and define builder functions used by SCons-facing scene-theme scripts.

### How the implementation applies it

#### Builder-module import

The icon SCsub script imports default_theme_icons_builders.

#### Builder function definition

The default theme builder module defines make_fonts_header.

### Walk through a small source example

```text
import default_theme_builders
```

Source: `scene/theme/SCsub` — module import

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/theme/SCsub:6` — import default_theme_builders
- Code: `scene/theme/default_theme_builders.py:8` — make_fonts_header
- Code: `scene/theme/icons/default_theme_icons_builders.py:9` — make_default_theme_icons_action
- External (official, verified): [The Python Language Reference: The import system](https://docs.python.org/3/reference/import.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ standard-library module import](#topic-cpp-module-import)
- [Python build-action functions](#topic-python-build-actions)
- [Python build code generation](#topic-python-build-code-generation)
- [Python build hooks](#topic-python-build-hooks)
- [Python build source generation](#topic-python-build-source-generation)
- [Python SCons build hooks](#topic-python-scons-build-hooks)
- [Python SCons configuration functions](#topic-python-scons-configuration)
- [Python SCons module configuration](#topic-python-scons-module-configuration)

<a id="topic-csharp-assembly-load-contexts"></a>
## 14. C# AssemblyLoadContext plugin loading

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

PluginLoadContext selects shared assemblies from the main context and resolves other dependencies through AssemblyDependencyResolver.

The editor plugin loader creates AssemblyLoadContext instances to resolve dependencies and support unloading.

### How the implementation applies it

#### Context construction

PluginLoadContext receives the plugin path, shared-assembly collection, main load context, and collectible flag, then creates AssemblyDependencyResolver.

#### Plugin entry loading

GodotPlugins.Main creates a PluginLoadContext, keeps a WeakReference to it, loads the requested assembly, and locates the GodotSharpEditor entry point.

### Walk through a small source example

```csharp
(var projectAssembly, _projectLoadContext) = LoadPlugin(assemblyPath, isCollectible: _editorHint);

                string loadedAssemblyPath = _projectLoadContext.AssemblyLoadedPath ?? assemblyPath;
                *outLoadedAssemblyPath = Marshaling.ConvertStringToNative(loadedAssemblyPath);

                ScriptManagerBridge.LookupScriptsInAssembly(projectAssembly);

                return godot_bool.True;
            }
            catch (Exception e)
            {
                Console.Error.WriteLine(e);
                return godot_bool.False;
```

Source: `modules/mono/glue/GodotSharp/GodotPlugins/Main.cs` — LoadPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/PluginLoadContext.cs:9` — PluginLoadContext
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/Main.cs:149` — LoadPlugin
- External (official, unverified (source anchor not found)): [About AssemblyLoadContext - .NET](https://learn.microsoft.com/en-us/dotnet/core/dependency-loading/understanding-assemblyloadcontext), accessed 2026-07-15

<a id="topic-csharp-attributes-reflection"></a>
## 15. C# attributes and reflection

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C# partial types and source generation](#topic-csharp-source-generation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Godot declares attributes for scripts, exported members, signals, RPCs, tools, and serialization, then bridge code inspects attributes at runtime.

Attributes declare engine-facing metadata and reflection reads it to connect types, members, and script information.

### How the implementation applies it

#### Attribute declarations

ExportAttribute constrains use to fields and properties and carries editor hint data; RpcAttribute supplies networking mode, call-local, transfer mode, and channel metadata.

#### Runtime discovery

ScriptManagerBridge reads ScriptPathAttribute, AssemblyHasScriptsAttribute, IconAttribute, RpcAttribute, and generated static metadata through reflection.

#### Analyzer inspection

GlobalClassAnalyzer uses Roslyn symbols and attributes to enforce global-class constraints during compilation.

### Walk through a small source example

```csharp
[AttributeUsage(AttributeTargets.Field | AttributeTargets.Property)]
    public sealed class ExportAttribute : Attribute
    {
        /// <summary>
        /// Optional hint that determines how the property should be handled by the editor.
        /// </summary>
        public PropertyHint Hint { get; }

        /// <summary>
        /// Optional string that can contain additional metadata for the <see cref="Hint"/>.
        /// </summary>
        public string HintString { get; }

        /// <summary>
```

Source: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/ExportAttribute.cs` — ExportAttribute (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ScriptPathAttribute, AssemblyHasScriptsAttribute

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/ExportAttribute.cs:9` — ExportAttribute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/RpcAttribute.cs:12` — RpcAttribute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ScriptManagerBridge.cs:20` — ScriptManagerBridge
- External (official, unverified (source anchor not found)): [Attributes and reflection - C#](https://learn.microsoft.com/en-us/dotnet/csharp/advanced-topics/reflection-and-attributes/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)

<a id="topic-csharp-partial-source-generation"></a>
## 16. C# partial types and source-generator tests

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Partial declarations let generated files add members to a type declared in the sample or test source.

The .NET SDK test projects define partial C# types and verify generators that emit Godot-facing signal, property, method, serialization, and script-path code.

### How the implementation applies it

#### Signal-bearing partial type

The EventSignals sample declares a partial GodotObject class with a Signal-attributed delegate.

#### Generator verification

ScriptPropertiesGeneratorTests invokes a CSharpSourceGeneratorVerifier for exported fields and properties.

#### Generated serialization implementation

The checked-in generated serialization files override SaveGodotObjectData and RestoreGodotObjectData.

### Walk through a small source example

```csharp
public partial class EventSignals : GodotObject
{
    [Signal]
    public delegate void MySignalEventHandler(string str, int num);
}
```

Source: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Sample/EventSignals.cs` — EventSignals

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CSharpScript

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Sample/EventSignals.cs:3` — EventSignals
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/ScriptPropertiesGeneratorTests.cs:6` — ScriptPropertiesGeneratorTests
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/TestData/GeneratedSources/ScriptBoilerplate_ScriptSerialization.generated.cs` — SaveGodotObjectData and RestoreGodotObjectData
- External (official, verified): [Partial type - C# reference](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/partial-type), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C# partial types and source generation](#topic-csharp-source-generation)
- [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)
- [Properties and accessors](#topic-properties-and-accessors)
- [Signals and await](#topic-signals-and-await)

<a id="topic-cpp-binary-data-codecs"></a>
## 17. C++ binary data codecs

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The accessor implementation separates byte transport from conversion into numeric and vector values.

glTF decoding reads typed C++ containers of raw bytes through buffer-view offsets and strides, while encoding writes byte arrays into newly allocated buffer views.

### How the implementation applies it

#### Buffer-view byte extraction

GLTFBufferView obtains bytes from the indexed buffer and calculates an end offset from byte_offset and byte_length.

#### Accessor byte decoding

GLTFAccessor loads a PackedByteArray from a buffer view before interpreting component-sized values.

#### Encoded buffer-view creation

Accessor encoding sends generated bytes to GLTFBufferView::write_new_buffer_view_into_state.

### Walk through a small source example

```cpp
const PackedByteArray raw_bytes = raw_buffer_view->load_buffer_view_data(p_gltf_state);
```

Source: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor::_decode_as_numbers

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GLTFBufferView, GLTFAccessor

Evidence:
- Code: `modules/gltf/structures/gltf_buffer_view.cpp:39` — GLTFBufferView::load_buffer_view_data
- Code: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor byte decode and buffer-view write paths
- External (primary, verified): [C++ Working Draft: Fundamental types](https://eel.is/c++draft/basic.types), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-const-reference-accessors"></a>
## 18. C++ const-reference accessors

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Metadata and API accessors make read-only intent explicit in both parameter and return types.

C++ member functions expose lookup results through const pointers and accept immutable String references to avoid copying lookup inputs.

### How the implementation applies it

#### Profile metadata lookup

The two-argument get_io_path takes both path strings by const reference and returns a const IOPath pointer.

#### Wrapper registry access

get_registered_extension_wrappers returns a const Vector reference, giving callers read-only access to the registered wrapper collection.

### Walk through a small source example

```cpp
const OpenXRInteractionProfileMetadata::IOPath *OpenXRInteractionProfileMetadata::get_io_path(const String &p_interaction_profile, const String &p_io_path) const {
```

Source: `modules/openxr/action_map/openxr_interaction_profile_metadata.cpp` — OpenXRInteractionProfileMetadata::get_io_path

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OpenXRInteractionProfile

Evidence:
- Code: `modules/openxr/action_map/openxr_interaction_profile_metadata.cpp:210` — OpenXRInteractionProfileMetadata::get_io_path
- Code: `modules/openxr/openxr_api.cpp:106` — OpenXRAPI::get_registered_extension_wrappers
- External (primary, verified): [Working Draft, Programming Languages — C++](https://eel.is/c++draft/dcl.fct), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-inheritance-interfaces"></a>
## 19. C++ inheritance interfaces

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The partition consistently uses derived classes to plug implementations into engine-facing abstractions.

C++ classes specialize engine and OpenXR base interfaces through public inheritance.

### How the implementation applies it

#### Spatial extension specialization

OpenXRSpatialEntityExtension derives from OpenXRExtensionWrapper, placing spatial behavior behind the wrapper base type.

#### Image-loader specialization

ImageLoaderSVG and ImageLoaderTGA derive from ImageFormatLoader, expressing the image-format adapter contract in their class declarations.

#### Text-server specialization

TextServerAdvanced and TextServerFallback each derive from TextServerExtension, providing alternate implementations behind one engine-facing type.

### Walk through a small source example

```c
class OpenXRSpatialEntityExtension : public OpenXRExtensionWrapper {
```

Source: `modules/openxr/extensions/spatial_entities/openxr_spatial_entity_extension.h` — OpenXRSpatialEntityExtension declaration

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OpenXRStructureBase, ShapedTextDataAdvanced

Evidence:
- Code: `modules/openxr/extensions/spatial_entities/openxr_spatial_entity_extension.h:42` — class OpenXRSpatialEntityExtension : public OpenXRExtensionWrapper
- Code: `modules/svg/image_loader_svg.h:35` — class ImageLoaderSVG : public ImageFormatLoader
- Code: `modules/text_server_adv/text_server_adv.h:100` — class TextServerAdvanced : public TextServerExtension
- External (primary, verified): [Working Draft, Programming Languages — C++](https://eel.is/c++draft/class.derived), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Java Android host APIs](#topic-java-android-host-api)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ class inheritance](#topic-cpp-object-hierarchies)

<a id="topic-cpp-polymorphic-backends"></a>
## 20. C++ polymorphic backend classes

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The repository uses public inheritance across audio, filesystem, networking, rendering-context, rendering-device, and editor plugin implementations.

C++ backend classes inherit engine interfaces so a rendering or platform service can be selected through a common base type.

### How the implementation applies it

#### Rendering context specialization

RenderingContextDriverVulkan derives from RenderingContextDriver, identifying a concrete context implementation.

#### Rendering device specialization

RenderingDeviceDriverVulkan derives from RenderingDeviceDriver, and the Direct3D 12 and Metal drivers use the same abstraction pattern.

#### Platform service specialization

AudioDriverALSA derives from AudioDriver, while Unix file and socket implementations derive their corresponding engine interfaces.

### Walk through a small source example

```c
class AudioDriverALSA : public AudioDriver {
	Thread thread;
	Mutex mutex;

	snd_pcm_t *pcm_handle = nullptr;

	String output_device_name = "Default";
	String new_output_device = "Default";

	Vector<int32_t> samples_in;
	Vector<int16_t> samples_out;

	Error init_output_device();
```

Source: `drivers/alsa/audio_driver_alsa.h` — class AudioDriverALSA : public AudioDriver (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RenderingDeviceDriverVulkan::BufferInfo

Evidence:
- Code: `drivers/vulkan/rendering_context_driver_vulkan.h:46` — class RenderingContextDriverVulkan : public RenderingContextDriver
- Code: `drivers/alsa/audio_driver_alsa.h:43` — class AudioDriverALSA : public AudioDriver
- External (primary, unverified (source anchor not found)): [C++ Working Draft: Virtual Functions](https://eel.is/c++draft/class.virtual), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ virtual dispatch](#topic-cpp-virtual-dispatch)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)

<a id="topic-cpp-resource-and-polymorphism"></a>
## 21. C++ resources and polymorphic engine objects

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The implementation derives persisted configuration from Resource and runtime services from engine base classes such as MultiplayerAPI and NavigationServer2D or NavigationServer3D.

C++ engine modules define polymorphic Resource and server objects that own long-lived configuration and runtime state.

### How the implementation applies it

#### Persisted replication resource

SceneReplicationConfig derives from Resource, declares a save type and resource extension, and maintains property lists and replication modes.

#### Runtime service inheritance

SceneMultiplayer derives from MultiplayerAPI, while the 2D and 3D servers derive from their respective NavigationServer base classes.

#### Specialized resource hierarchies

Noise, OggPacketSequence, OpenXRActionMap, and related types use Resource inheritance for engine-facing configuration and content.

### Walk through a small source example

```c
class SceneReplicationConfig : public Resource {
	GDCLASS(SceneReplicationConfig, Resource);
	OBJ_SAVE_TYPE(SceneReplicationConfig);
```

Source: `modules/multiplayer/scene_replication_config.h` — SceneReplicationConfig

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SceneReplicationConfig, OggPacketSequence, OpenXRActionMap

Evidence:
- Code: `modules/multiplayer/scene_replication_config.h:36` — SceneReplicationConfig
- Code: `modules/multiplayer/scene_multiplayer.h:64` — SceneMultiplayer
- Code: `modules/navigation_2d/2d/godot_navigation_server_2d.h:59` — GodotNavigationServer2D
- Code: `modules/navigation_3d/3d/godot_navigation_server_3d.h:59` — GodotNavigationServer3D
- External (primary, verified): [C++ working draft — Object model](https://eel.is/c++draft/intro.object), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ standard-library module import](#topic-cpp-module-import)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-scope-locking"></a>
## 22. C++ scope-bound locking

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The repeated local MutexLock declarations make the synchronization boundary visible at write sites.

The Jolt contact listener creates scope-bound C++ lock objects before mutating shared contact, manifold, and debug-contact collections.

### How the implementation applies it

#### Contact write protection

JoltContactListener3D creates MutexLock write_lock(write_mutex) before handling contact data.

#### Shared debug-contact updates

The listener uses the same lock pattern around debug-contact capacity and append operations.

### Walk through a small source example

```cpp
const MutexLock write_lock(write_mutex);
```

Source: `modules/jolt_physics/spaces/jolt_contact_listener_3d.cpp` — JoltContactListener3D contact write paths

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JoltSpace3D

Evidence:
- Code: `modules/jolt_physics/spaces/jolt_contact_listener_3d.cpp:195` — MutexLock write_lock
- External (primary, verified): [C++ Working Draft: Mutex requirements](https://eel.is/c++draft/thread.mutex.requirements), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)

<a id="topic-cpp-templates-and-const-access"></a>
## 23. C++ templates and const access

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code uses Ref<T>, Vector<T>, LocalVector<T>, TypedArray<T>, and const references or pointers throughout module and scene implementations.

C++ templates and const-qualified values provide typed access to engine handles and collection views.

### How the implementation applies it

#### Typed resource handle

VisualShader language integration stores a typed Ref<VisualShader> handle from its shader input.

#### Read-only varying access

VisualShader exposes varyings through a pointer-to-const lookup API.

#### Read-only navigation path access

NavigationAgent3D returns its current path as a const Vector<Vector3> reference.

### Walk through a small source example

```cpp
const Ref<VisualShader> shader = p_shader;
```

Source: `modules/visual_shader/editor/visual_shader_language_plugin.cpp` — VisualShaderLanguagePlugin

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VisualShaderGraph, NavigationAgent3D

Evidence:
- Code: `modules/visual_shader/editor/visual_shader_language_plugin.cpp:35` — VisualShaderLanguagePlugin
- Code: `modules/visual_shader/visual_shader.h` — VisualShader::get_varying_by_index
- Code: `scene/3d/navigation/navigation_agent_3d.h` — NavigationAgent3D::get_current_navigation_path
- External (primary, verified): [Working Draft, Standard for Programming Language C++](https://eel.is/c++draft/temp), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ variadic binding](#topic-cpp-variadic-binding)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)

<a id="topic-cpp-templates-and-specialization"></a>
## 24. C++ templates and specialization

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The reusable interpolation structure is specialized where quaternion behavior needs its own implementation.

The glTF importer uses C++ class templates for interpolation and supplies a quaternion specialization; accessor decoding also instantiates numeric decoding for concrete element types.

### How the implementation applies it

#### Quaternion interpolation specialization

SceneFormatImporterGLTFInterpolate has a Quaternion specialization in the glTF document implementation.

#### Typed numeric decoding

GLTFAccessor decodes numeric representations through _decode_as_numbers<double>.

### Walk through a small source example

```cpp
struct SceneFormatImporterGLTFInterpolate<Quaternion> {
```

Source: `modules/gltf/gltf_document.cpp` — SceneFormatImporterGLTFInterpolate<Quaternion>

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GLTFAnimation, GLTFAccessor

Evidence:
- Code: `modules/gltf/gltf_document.cpp:5004` — SceneFormatImporterGLTFInterpolate<Quaternion>
- Code: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor::_decode_as_numbers<double>
- External (primary, verified): [C++ Working Draft: Templates](https://eel.is/c++draft/temp), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [Python SCons build hooks](#topic-python-scons-build-hooks)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ templates and specialization](#topic-cxx-templates)

<a id="topic-callables-and-lambdas"></a>
## 25. Callables and lambdas

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Signals and await](#topic-signals-and-await).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The tests use typed parameters and returns, direct `call()`, captures, binding, and callable introspection.

Callable values can reference functions, constructors, built-ins, and lambdas; lambda bodies capture their enclosing context.

### How the implementation applies it

#### Typed lambda invocation

Lambdas with zero, one, and typed return values are stored and invoked through `call()`.

#### Captured state

A lambda reads a surrounding local value and a member value when called later.

### Walk through a small source example

```text
func test():
	var lambda_0 := func() -> void:
		print(0)
	lambda_0.call()
```

Source: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd` — test()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd:1` — test()
- Code: `modules/gdscript/tests/scripts/runtime/features/lambda_use_self.gd` — lambda1, lambda2, nested, lambda3, and lambda4
- Code: `modules/gdscript/tests/scripts/runtime/features/ctor_as_callable.gd` — inner_ctor and native_ctor
- External (official, verified): [Callable class reference](https://docs.godotengine.org/en/stable/classes/class_callable.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C stateful struct APIs](#topic-c-stateful-struct-apis)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ lambdas and callback objects](#topic-cxx-closures)
- [Lambdas and standard algorithms](#topic-cxx-lambdas-standard-algorithms)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)

<a id="topic-classes-inheritance-and-lookup"></a>
## 26. Classes, inheritance, and lookup

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle), [Properties and accessors](#topic-properties-and-accessors).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Fixtures test nested and external resolution separately from runtime inheritance behavior.

Classes supply nested types, inheritance, `super`, and lexical member lookup; preloaded classes also participate in lookup.

### How the implementation applies it

#### Inner construction

Nested classes construct both themselves and sibling inner classes through their class names.

#### Inherited nested lookup

A nested class inherits from an outer class and resolves a constant through the inherited base.

### Walk through a small source example

```text
class A:
	func a():
		return A.new()
	func b():
		return B.new()
```

Source: `modules/gdscript/tests/scripts/analyzer/features/inner_class_access_from_inside.gd` — A.a() and A.b()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/inner_class_access_from_inside.gd` — A and B
- Code: `modules/gdscript/tests/scripts/analyzer/features/lookup_class.gd:40` — Y.B.check()
- Code: `modules/gdscript/tests/scripts/parser/features/super.gd:12` — SayAnotherThing
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [Scoped enums and bit flags](#topic-cxx-enums-bitflags)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [Java Android host APIs](#topic-java-android-host-api)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)

<a id="topic-gdscript-typed-collections"></a>
## 27. GDScript typed collections and signature compatibility

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The analyzer test corpus specifies accepted and rejected typed containers, function variance, and override signatures.

This extends type declarations with typed Array and Dictionary element constraints, plus parameter and return compatibility checks for inherited functions.

### How the implementation applies it

#### Typed array argument

The analyzer test declares a function parameter as Array[int].

#### Typed dictionary diagnostics

Tests expect errors when dictionary key or value types conflict with Dictionary[int, int].

#### Override compatibility

Feature tests cover contravariant parameter and covariant return cases.

### Walk through a small source example

```text
func expect_typed(typed: Array[int]):
```

Source: `modules/gdscript/tests/scripts/analyzer/errors/typed_array.gd` — expect_typed

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDScriptParser::Node

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/errors/typed_array.gd:1` — expect_typed
- Code: `modules/gdscript/tests/scripts/analyzer/features/function_param_type_contravariance.gd` — function parameter compatibility cases
- External (official, unverified (source anchor not found)): [Static typing in GDScript](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/static_typing.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [Typed arrays and dictionaries](#topic-typed-collections)

<a id="topic-glsl-compute-shaders"></a>
## 28. GLSL compute shaders

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These shaders are converted into C++ headers by the module build.

Betsy shader sources declare compute workgroup sizes, bind sampled and storage images, fetch texels by invocation ID, and store generated output.

### How the implementation applies it

#### Workgroup declaration

alpha_stitch.glsl sets an 8 by 8 by 1 local workgroup.

#### Resource bindings

The shader binds two unsigned sampler inputs and one rgba32ui output image.

#### Invocation-indexed writes

main fetches source texels using gl_GlobalInvocationID.xy and stores a combined block.

### Walk through a small source example

```text
void main() {
	uvec2 rgbBlock = texelFetch(srcRGB, ivec2(gl_GlobalInvocationID.xy), 0).xy;
	uvec2 alphaBlock = texelFetch(srcAlpha, ivec2(gl_GlobalInvocationID.xy), 0).xy;

	imageStore(dstTexture, ivec2(gl_GlobalInvocationID.xy), uvec4(rgbBlock.xy, alphaBlock.xy));
}
```

Source: `modules/betsy/alpha_stitch.glsl` — main

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Image

Evidence:
- Code: `modules/betsy/alpha_stitch.glsl:16` — main
- Code: `modules/betsy/SCsub:10` — GLSL_HEADER
- External (official, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-iteration-protocols"></a>
## 29. Iteration protocols

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Typed and `Variant` paths are exercised for built-in and user-defined iteration.

`for` loops cover ranges, collections, and custom objects that provide iterator hook functions.

### How the implementation applies it

#### Custom iterator hooks

Custom iterator classes declare `_iter_init`, `_iter_next`, and `_iter_get` methods.

#### Typed loop sources

Fixtures iterate over typed arrays, dictionaries, ranges, strings, and object iterators.

### Walk through a small source example

```text
func _iter_get(_count) -> StringName:
```

Source: `modules/gdscript/tests/scripts/runtime/features/for_loop_iterator_types.gd` — Iterator._iter_get()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/runtime/features/for_loop_iterator_object.gd` — ExponentialIterator._iter_init(), _iter_next(), and _iter_get()
- Code: `modules/gdscript/tests/scripts/runtime/features/for_loop_iterator_types.gd` — Iterator._iter_get()
- Code: `modules/gdscript/tests/scripts/parser/features/for_range.gd:1` — test()
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ member-function overloading](#topic-cpp-overloading)

<a id="topic-javascript-browser-ffi"></a>
## 30. JavaScript browser FFI

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The bridges use runtime allocation and parsing helpers, ID mapping, browser event callbacks, and Web API object methods.

JavaScript bridge libraries marshal browser objects, strings, buffers, and callbacks between web APIs and C++ platform code.

### How the implementation applies it

#### RTC data-channel creation

The WebRTC bridge parses a configuration object and invokes createDataChannel on the browser peer connection.

#### WebSocket event callbacks

The WebSocket bridge retrieves runtime functions and binds open, message, error, and close callbacks.

#### WebXR session and input bridge

The WebXR bridge parses session options, allocates callback strings, and maintains browser-side frame and input-source state.

### Walk through a small source example

```javascript
const channel = ref.createDataChannel(label, config);
```

Source: `modules/webrtc/library_godot_webrtc.js` — GodotRTCPeerConnection

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebRTCPeerConnection, WebRTCDataChannel, WebXRInterfaceJS

Evidence:
- Code: `modules/webrtc/library_godot_webrtc.js:232` — GodotRTCPeerConnection
- Code: `modules/websocket/library_godot_websocket.js:31` — GodotWebSocket
- Code: `modules/webxr/native/library_godot_webxr.js:31` — GodotWebXR
- External (primary, verified): [ECMAScript Language Specification](https://tc39.es/ecma262/multipage/ecmascript-language-functions-and-classes.html#sec-arrow-function-definitions), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [JavaScript browser runtime libraries](#topic-javascript-web-runtime)

<a id="topic-pattern-matching"></a>
## 31. Pattern matching and guards

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Parser and runtime fixtures separately cover structural matching syntax and guarded execution.

`match` supports literals, arrays, dictionaries, bindings, multiple patterns, wildcards, and guarded branches.

### How the implementation applies it

#### Structural patterns

Parser fixtures exercise array and dictionary patterns, including nested patterns and rest matching.

#### Guarded bindings

Runtime fixtures test `when` guards after bindings and verify guard-side-effect behavior.

### Walk through a small source example

```text
var a_bind when b == 0:
```

Source: `modules/gdscript/tests/scripts/runtime/features/match_with_pattern_guards.gd` — guarded match branch

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/parser/features/match_dictionary.gd:47` — test()
- Code: `modules/gdscript/tests/scripts/runtime/features/match_with_pattern_guards.gd:3` — test()
- Code: `modules/gdscript/tests/scripts/parser/errors/match_multiple_variable_binds_in_branch.out` — expected parser error
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

<a id="topic-python-scons-configuration"></a>
## 32. Python SCons configuration functions

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Module configuration is executable build logic rather than declarative project metadata alone.

Python build files define can_build and configure functions and call SCons environment methods to select module sources and generated artifacts.

### How the implementation applies it

#### Conditional module availability

The 2D physics module's can_build function reads disable_physics_2d from the build environment.

#### Source-file registration

SCsub scripts call add_source_files on the SCons environment, and the lightmapper script requests generated GLSL headers.

### Walk through a small source example

```python
def can_build(env, platform):
    return not env["disable_physics_2d"]
```

Source: `modules/godot_physics_2d/config.py` — can_build

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/godot_physics_2d/config.py:1` — can_build
- Code: `modules/lightmapper_rd/SCsub` — GLSL_HEADER and add_source_files
- External (official, verified): [Python Language Reference: Function definitions](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Python build-action functions](#topic-python-build-actions)
- [Python build code generation](#topic-python-build-code-generation)
- [Python build hooks](#topic-python-build-hooks)
- [Python build scripts](#topic-python-build-scripts)
- [Python SCons build hooks](#topic-python-scons-build-hooks)
- [Python build source generation](#topic-python-build-source-generation)
- [Python SCons module configuration](#topic-python-scons-module-configuration)

<a id="topic-python-scons-module-configuration"></a>
## 33. Python SCons module configuration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These scripts are the build-time composition layer for the supplied modules.

Python SCons scripts declare module buildability, dependencies, cloned environments, and source-file collection.

### How the implementation applies it

#### Module eligibility hook

The SVG config defines can_build and declares jpg and webp module dependencies before returning True.

#### Source collection

OpenXR and scene SCsub scripts create a module object list, add matching C++ sources, and append it to env.modules_sources.

#### Third-party source selection

The Theora and TinyEXR scripts conditionally assemble bundled third-party sources and add them to build environments.

### Walk through a small source example

```python
def can_build(env, platform):
    env.module_add_dependencies("svg", ["jpg", "webp"], True)
    return True
```

Source: `modules/svg/config.py` — can_build

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/svg/config.py:1` — can_build
- Code: `modules/openxr/extensions/SCsub` — module_obj and env.modules_sources
- Code: `modules/theora/SCsub:15` — thirdparty_sources
- External (official, verified): [Python 3 Language Reference](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Python build-action functions](#topic-python-build-actions)
- [Python build code generation](#topic-python-build-code-generation)
- [Python build hooks](#topic-python-build-hooks)
- [Python build scripts](#topic-python-build-scripts)
- [Python SCons build hooks](#topic-python-scons-build-hooks)
- [Python SCons configuration functions](#topic-python-scons-configuration)

<a id="topic-python-build-hooks"></a>
## 34. Python build hooks

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Top-level and module-local scripts define functions such as can_build and configure, while SCsub files add sources and web JavaScript libraries conditionally.

Python configuration scripts define module build hooks and configure build environments.

### How the implementation applies it

#### Module availability hook

The WebP module's can_build hook returns True.

#### Module configuration hook

The WebP module defines a configure hook even though its body is pass.

#### Platform-specific source configuration

WebRTC and WebXR SCsub scripts add JavaScript libraries when the platform is web.

### Walk through a small source example

```python
def can_build(env, platform):
    return True


def configure(env):
    pass
```

Source: `modules/webp/config.py` — can_build and configure

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/webp/config.py:1` — can_build
- Code: `modules/webrtc/SCsub` — web platform JavaScript library configuration
- Code: `modules/webxr/SCsub` — web platform JavaScript library configuration
- External (official, verified): [The Python Language Reference: Compound statements](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Python build-action functions](#topic-python-build-actions)
- [Python build code generation](#topic-python-build-code-generation)
- [Python build scripts](#topic-python-build-scripts)
- [Python build source generation](#topic-python-build-source-generation)
- [Python SCons build hooks](#topic-python-scons-build-hooks)
- [Python SCons configuration functions](#topic-python-scons-configuration)
- [Python SCons module configuration](#topic-python-scons-module-configuration)

<a id="topic-types-inference-and-conversions"></a>
## 35. Types, inference, and conversions

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Typed arrays and dictionaries](#topic-typed-collections), [Warnings and selective suppression](#topic-warnings-and-suppression).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the foundational type-system topic for the fixture corpus.

GDScript fixtures contrast explicit type hints, `:=` inference, `Variant` boundaries, null assignment, and conversion at typed assignments and returns.

### How the implementation applies it

#### Typed return conversion

Functions declared to return `float` accept integer literals, arguments, and locals in the conversion cases.

#### Weak local inference

A local declared with `=` can later receive values with different types, while a typed destination receives the resulting value.

### Walk through a small source example

```text
func convert_literal_int_to_float() -> float: return 76
func convert_arg_int_to_float(arg: int) -> float: return arg
```

Source: `modules/gdscript/tests/scripts/analyzer/features/return_conversions.gd` — convert_literal_int_to_float()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/return_conversions.gd:1` — convert_literal_int_to_float()
- Code: `modules/gdscript/tests/scripts/analyzer/features/local_inference_is_weak.gd:4` — test()
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [GDScript query-object API use](#topic-gdscript-query-objects)
- [C++ JNI Variant marshalling](#topic-cpp-jni-variant-marshalling)
- [C++ tagged storage and casts](#topic-cpp-tagged-storage)
- [C++ variadic binding](#topic-cpp-variadic-binding)

<a id="topic-warnings-and-suppression"></a>
## 36. Warnings and selective suppression

**Scope:** First-party

**Builds on:** [Types, inference, and conversions](#topic-types-inference-and-conversions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Warning text is itself captured as a tested output contract.

The analyzer reports unsafe, unused, shadowing, inference, conversion, and coroutine cases, while `@warning_ignore` selectively suppresses specified warnings.

### How the implementation applies it

#### Unsafe argument suppression

A local suppression precedes a call whose argument may have an unsafe static type.

#### Targeted warning fixtures

Separate fixtures capture unsafe casts, unused identifiers, confusable names, and await diagnostics.

### Walk through a small source example

```text
@warning_ignore("unsafe_call_argument")
		print(check(arg))
```

Source: `modules/gdscript/tests/scripts/analyzer/features/null_initializer.gd` — check_arg()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Expected Result Fixture, Test Script Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/null_initializer.gd` — check_arg()
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/unsafe_call_argument.out:2` — UNSAFE_CALL_ARGUMENT
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/confusable_capture_reassignment.out:2` — CONFUSABLE_CAPTURE_REASSIGNMENT
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Signals and await](#topic-signals-and-await)

<a id="topic-cpp-classes-and-inheritance"></a>
## 37. C++ classes and inheritance

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters), [C++ plugin specialization](#topic-cpp-plugin-specialization), [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime), [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Class declarations establish the editor subsystem types and their engine-facing base behavior.

The editor source declares engine services as C++ classes derived from Object, Resource, Control, and other engine base classes.

### How the implementation applies it

#### Persisted settings resource

EditorSettings derives from Resource, placing editor settings in the engine's resource hierarchy.

#### Interactive viewport control

Node3DEditorViewport derives from Control, making the 3D editor viewport a GUI control type.

#### Editor-facing object interface

EditorVCSInterface derives from Object and groups version-control data structures with its interface.

### Walk through a small source example

```c
class EditorSettings : public Resource {
```

Source: `editor/settings/editor_settings.h` — EditorSettings

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorSettings, VCS Diff File

Evidence:
- Code: `editor/settings/editor_settings.h:40` — EditorSettings : public Resource
- Code: `editor/scene/3d/node_3d_editor_viewport.h:146` — Node3DEditorViewport : public Control
- Code: `editor/version_control/editor_vcs_interface.h:37` — EditorVCSInterface : public Object
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-classes-inheritance"></a>
## 38. C++ classes and inheritance

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ class templates and specialization](#topic-cpp-class-templates), [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers), [C++ engine binding macros](#topic-cpp-godot-binding-macros), [C++ virtual interfaces](#topic-cpp-runtime-polymorphism), [C++ scoped lock guards](#topic-cpp-scoped-locks).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Class inheritance provides the common contracts on which the editor’s services and UI controls are built.

The editor is structured from C++ classes that derive from engine base types such as Node, Container, ScrollContainer, ResourceImporter, and RefCounted.

### How the implementation applies it

#### Editor plug-in base

EditorPlugin derives from Node, making plug-ins engine nodes with editor-specific hooks.

#### Importer base

ResourceImporterScene derives from ResourceImporter, separating scene-specific importing from the generic resource-import contract.

#### Inspector UI bases

EditorProperty and EditorInspector derive from GUI container classes, so property editors compose into the editor UI hierarchy.

### Walk through a small source example

```c
class EditorPlugin : public Node {
```

Source: `editor/plugins/editor_plugin.h` — EditorPlugin

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/plugins/editor_plugin.h:59` — EditorPlugin
- Code: `editor/import/3d/resource_importer_scene.h:155` — ResourceImporterScene
- Code: `editor/inspector/editor_inspector.h:70` — EditorProperty
- External (primary, verified): [C++ Working Draft — Derived classes](https://eel.is/c++draft/class.derived), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Java Android host APIs](#topic-java-android-host-api)
- [C++ inheritance and reference-counted adapters](#topic-cpp-runtime-adapters)

<a id="topic-cplusplus-enumerated-export-state"></a>
## 39. C++ enumerated export state

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Named integral categories keep export configuration and reporting states explicit at API boundaries.

C++ enumerations classify export messages, export filters, and file or script export modes.

### How the implementation applies it

#### Message severity

EditorExportPlatform::ExportMessageType classifies information, warnings, and errors emitted during export.

#### Preset resource policy

EditorExportPreset::ExportFilter and FileExportMode encode file-selection and per-file customization policies.

#### Script representation

EditorExportPreset::ScriptExportMode encodes text and binary script export choices.

### Walk through a small source example

```c
enum ExportMessageType {
```

Source: `editor/export/editor_export_platform.h` — EditorExportPlatform::ExportMessageType

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorExportPreset, EditorExportPlatform::ExportMessage

Evidence:
- Code: `editor/export/editor_export_platform.h:372` — EditorExportPlatform::ExportMessageType
- Code: `editor/export/editor_export_preset.h:230` — EditorExportPreset::ExportFilter
- Code: `editor/export/editor_export_preset.h:232` — EditorExportPreset::ScriptExportMode
- External (primary, unverified (source anchor not found)): [Draft C++ Standard — Enumerations](https://eel.is/c++draft/dcl.enum), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Scoped enums and bit flags](#topic-cxx-enums-bitflags)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cplusplus-export-callbacks"></a>
## 40. C++ export callbacks

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

EditorExportPlatform carries callback types so common packing code can delegate destination-specific work.

C++ function-pointer and callable declarations connect generic packaging code to save, removal, shared-object, and script callback implementations.

### How the implementation applies it

#### Save callback type

EditorExportSaveFunction defines the function-pointer boundary used by platform packaging operations.

#### Script callback data

ScriptCallbackData groups callback data used when export processing invokes script-provided behavior.

#### Packaging entry points

The platform's pack and ZIP operations accept callback-oriented save paths instead of hard-coding every output destination.

### Walk through a small source example

```c
struct ScriptCallbackData {
```

Source: `editor/export/editor_export_platform.h` — EditorExportPlatform::ScriptCallbackData

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorExportPlatform

Evidence:
- Code: `editor/export/editor_export_platform.h:58` — EditorExportSaveFunction
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::ScriptCallbackData
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::save_pack
- External (primary, verified): [Draft C++ Standard — Function declarators](https://eel.is/c++draft/dcl.fct), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C function-pointer callbacks](#topic-c-function-pointer-callbacks)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [Callables and lambdas](#topic-callables-and-lambdas)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)
- [C function-pointer API declarations](#topic-c-function-pointer-apis)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ lambdas and callback objects](#topic-cxx-closures)

<a id="topic-cpp-runtime-adapters"></a>
## 41. C++ inheritance and reference-counted adapters

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ JNI Variant marshalling](#topic-cpp-jni-variant-marshalling).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Android wrapper types and display/OS classes use C++ classes and inheritance as the native adaptation mechanism.

C++ defines reference-counted Java adapter classes and platform subclasses that specialize common engine interfaces.

### How the implementation applies it

#### JavaClass

`JavaClass` is declared as a `RefCounted` class, establishing a reference-counted native wrapper type.

#### JavaObject

`JavaObject` is another `RefCounted` wrapper class in the same Android API header.

#### DisplayServerAndroid

`DisplayServerAndroid` derives from `DisplayServer`, placing Android display behavior behind the engine display-server interface.

### Walk through a small source example

```c
class JavaClass : public RefCounted {
```

Source: `platform/android/api/java_class_wrapper.h` — JavaClass

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/api/java_class_wrapper.h:47` — JavaClass
- Code: `platform/android/api/java_class_wrapper.h:44` — JavaObject
- Code: `platform/android/display_server_android.h:43` — DisplayServerAndroid
- External (primary, unverified (source anchor not found)): [Draft C++ Standard: Contents](https://eel.is/c++draft/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)

<a id="topic-cplusplus-polymorphic-platform-adapters"></a>
## 42. C++ polymorphic platform adapters

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The common export-platform base is extended by platform-specific exporters, including the Apple embedded exporter.

C++ class inheritance and virtual functions define the target-platform adapter contract.

### How the implementation applies it

#### Common target contract

EditorExportPlatform declares the target-export interface that concrete platform exporters implement.

#### Apple specialization

EditorExportPlatformAppleEmbedded derives from EditorExportPlatform to provide Apple embedded behavior.

### Walk through a small source example

```c
class EditorExportPlatformAppleEmbedded : public EditorExportPlatform {
```

Source: `editor/export/editor_export_platform_apple_embedded.h` — EditorExportPlatformAppleEmbedded

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorExportPlatform

Evidence:
- Code: `editor/export/editor_export_platform.h:51` — EditorExportPlatform
- Code: `editor/export/editor_export_platform_apple_embedded.h:60` — EditorExportPlatformAppleEmbedded
- External (primary, unverified (source anchor not found)): [Draft C++ Standard — Virtual functions](https://eel.is/c++draft/class.virtual), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ virtual interfaces](#topic-cpp-runtime-polymorphism)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)

<a id="topic-javascript-web-runtime"></a>
## 43. JavaScript browser runtime libraries

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Web platform places browser-facing implementation in JavaScript files while C++ classes call their exported bridge functions.

JavaScript libraries implement engine startup, preloading, runtime heap access, browser display and input callbacks, audio worklets, fetch, MIDI, and JavaScript object proxies.

### How the implementation applies it

#### Engine

The engine library creates an Engine closure and uses Preloader and configuration objects to initialize the Web build.

#### GodotRuntime

Runtime helpers allocate strings and access typed heap views used to exchange WebAssembly memory with JavaScript.

#### GodotJSWrapper

The object bridge maps proxied JavaScript values and callbacks to IDs and conversion callbacks.

#### GodotProcessor

The audio worklet implementation consumes and produces audio through a RingBuffer and AudioWorkletProcessor.

### Walk through a small source example

```javascript
const Engine = (function () {
```

Source: `platform/web/js/engine/engine.js` — Engine

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JavaScriptObjectImpl

Evidence:
- Code: `platform/web/js/engine/engine.js:2` — Engine
- Code: `platform/web/js/libs/library_godot_runtime.js:31` — GodotRuntime
- Code: `platform/web/js/libs/library_godot_javascript_singleton.js:31` — GodotJSWrapper
- Code: `platform/web/js/libs/audio.worklet.js:96` — GodotProcessor
- External (primary, verified): [ECMAScript Language Specification](https://tc39.es/ecma262/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [JavaScript browser FFI](#topic-javascript-browser-ffi)

<a id="topic-python-scons-build-hooks"></a>
## 44. Python SCons build hooks

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The platform build layer uses Python functions called by SCons environments.

Python build scripts define platform detection, option and flag setup, configuration hooks, bundle generation, template assembly, and Emscripten helper functions.

### How the implementation applies it

#### platform.web.detect.configure

The Web platform exposes a configure hook alongside platform name, build capability, option, flag, and library-emitter functions.

#### create_engine_file

Web Emscripten helpers create engine output and collect JavaScript libraries, pre/post scripts, and externs.

#### generate_bundle

macOS builder code supplies bundle-generation and debug-app construction functions.

### Walk through a small source example

```python
def generate_bundle(target, source, env):
```

Source: `platform/macos/platform_macos_builders.py` — generate_bundle

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/web/detect.py:105` — configure
- Code: `platform/web/emscripten_helpers.py:26` — create_engine_file
- Code: `platform/macos/platform_macos_builders.py:11` — generate_bundle
- External (official, unverified (source anchor not found)): [The Python Language Reference](https://docs.python.org/3/reference/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ templates](#topic-cpp-templates)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates and specialization](#topic-cxx-templates)
- [Python build-action functions](#topic-python-build-actions)
- [Python build code generation](#topic-python-build-code-generation)
- [Python build hooks](#topic-python-build-hooks)
- [Python build scripts](#topic-python-build-scripts)
- [GDScript query-object API use](#topic-gdscript-query-objects)
- [Python SCons configuration functions](#topic-python-scons-configuration)
- [Python SCons module configuration](#topic-python-scons-module-configuration)

<a id="topic-c-abi-bridging"></a>
## 45. C ABI bridging

**Scope:** Vendored: openxr

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Calling-convention macros are selected per platform and are used in both exported functions and function-pointer declarations.

OpenXR headers use macros and `extern "C"` guards so C++ callers expose C-compatible API declarations and function-pointer types.

### How the implementation applies it

#### XRAPI_CALL

The platform header selects a calling convention, including `__stdcall` on Windows.

#### XRAPI_PTR

The same calling-convention selection is applied to API function-pointer types.

#### extern "C" guard

The header opens C linkage when included by C++ code.

### Walk through a small source example

```c
#ifdef __cplusplus
extern "C" {
#endif
```

Source: `thirdparty/openxr/include/openxr/openxr_platform_defines.h` — C linkage guard

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

XrInstanceCreateInfo

Evidence:
- Code: `thirdparty/openxr/include/openxr/openxr_platform_defines.h:21` — XRAPI_CALL
- Code: `thirdparty/openxr/include/openxr/openxr_platform_defines.h:23` — XRAPI_PTR
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)

<a id="topic-c-abi-manual-lifetime"></a>
## 46. C ABI structures and manual lifetime

**Scope:** Vendored: spirv-reflect

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The header is usable from C and wraps its declarations in extern "C" when compiled as C++.

The SPIR-V reflection interface exchanges C structs and pointers through create, enumerate, query, change, and explicit destroy functions.

### How the implementation applies it

#### Caller-provided module record

spvReflectCreateShaderModule receives SPIR-V bytes and a caller-provided SpvReflectShaderModule pointer.

#### Explicit retirement

spvReflectDestroyShaderModule accepts a module pointer and is the paired lifetime operation.

#### Pointer-array enumeration

Enumeration functions accept a count pointer and an output array of pointers owned by the reflection module.

### Walk through a small source example

```c
SpvReflectResult spvReflectCreateShaderModule(
  size_t                   size,
  const void*              p_code,
  SpvReflectShaderModule*  p_module
);
```

Source: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectCreateShaderModule

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SpvReflectShaderModule, SpvReflectDescriptorSet, SpvReflectDescriptorBinding

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectCreateShaderModule, spvReflectDestroyShaderModule
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:723` — spvReflectEnumerateDescriptorBindings
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — extern "C" linkage guard
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI bridging](#topic-c-abi-bridging)
- [C ABI structures and opaque state](#topic-c-abi-data-structures)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)
- [C ABI-versioned initialization](#topic-c-abi-versioned-initialization)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C-managed concurrent state](#topic-c-concurrent-state)
- [C dynamic-library function wrappers](#topic-c-dynamic-library-wrappers)
- [C explicit resource lifecycles](#topic-c-explicit-resource-lifecycles)
- [C function-pointer API declarations](#topic-c-function-pointer-apis)
- [C function-pointer callbacks](#topic-c-function-pointer-callbacks)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C macro-based binary decoding](#topic-c-macro-codecs)
- [C parser state machines](#topic-c-parser-state)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C preprocessor composition](#topic-c-preprocessor-composition)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C stateful streaming APIs](#topic-c-stateful-streaming-api)
- [C stateful struct APIs](#topic-c-stateful-struct-apis)
- [C structs and pointer-linked state](#topic-c-structs-pointers)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [CMake native source indexing](#topic-cmake-native-source-index)
- [C++ enumerated export state](#topic-cplusplus-enumerated-export-state)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ binary data codecs](#topic-cpp-binary-data-codecs)
- [C++ byte-exact binary parsing](#topic-cpp-byte-exact-binary-parsing)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ mapping of C API types](#topic-cpp-c-type-mapping)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ const-reference accessors](#topic-cpp-const-reference-accessors)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ exception containment at ABI boundaries](#topic-cpp-exception-abi-boundaries)
- [C++ frame-time arithmetic](#topic-cpp-frame-scheduling)
- [C-compatible function-pointer interfaces](#topic-cpp-function-pointer-interfaces)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ engine binding macros](#topic-cpp-godot-binding-macros)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ JNI Variant marshalling](#topic-cpp-jni-variant-marshalling)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ plugin specialization](#topic-cpp-plugin-specialization)
- [C++ polymorphic backend classes](#topic-cpp-polymorphic-backends)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ resources and polymorphic engine objects](#topic-cpp-resource-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ inheritance and reference-counted adapters](#topic-cpp-runtime-adapters)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ scope-bound locking](#topic-cpp-scope-locking)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ static ABI contracts](#topic-cpp-static-abi-contracts)
- [C++ tagged storage and casts](#topic-cpp-tagged-storage)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ typed API records](#topic-cpp-typed-api-records)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C# partial types and source-generator tests](#topic-csharp-partial-source-generation)
- [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop)
- [C++ atomic synchronization](#topic-cxx-atomics)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ lambdas and callback objects](#topic-cxx-closures)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [JavaScript browser FFI](#topic-javascript-browser-ffi)
- [Objective-C++ Apple platform adapters](#topic-objective-cpp-apple-adapters)
- [Objective-C++ iOS adapters](#topic-objective-cpp-ios-adapters)
- [Python build source generation](#topic-python-build-source-generation)
- [C# attributes and reflection](#topic-csharp-attributes-reflection)
- [GDScript query-object API use](#topic-gdscript-query-objects)

<a id="topic-c-abi-data-structures"></a>
## 47. C ABI structures and opaque state

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Vulkan descriptors expose typed fields, while zlib hides its internal state behind a pointer.

C declarations use tagged structures, pointer members, and opaque implementation pointers to express caller-owned descriptors and library-managed state across API boundaries.

### How the implementation applies it

#### Caller-owned Vulkan descriptor

VkWaylandSurfaceCreateInfoKHR carries structure type, extension pointer, flags, display, and surface fields into a creation call.

#### Opaque compression state

zlib.h declares struct internal_state and stores its pointer in the public stream state rather than exposing the implementation layout.

#### Forward-declared implementation contexts

Wslay and Zstandard headers declare context structures used through API-level pointers.

### Walk through a small source example

```c
typedef struct VkWaylandSurfaceCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkWaylandSurfaceCreateFlagsKHR    flags;
    struct wl_display*                display;
    struct wl_surface*                surface;
} VkWaylandSurfaceCreateInfoKHR;
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h` — VkWaylandSurfaceCreateInfoKHR

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VkWaylandSurfaceCreateInfoKHR, z_stream

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h:27` — VkWaylandSurfaceCreateInfoKHR
- Code: `thirdparty/zlib/zlib.h:100` — struct internal_state FAR *state
- Code: `thirdparty/wslay/wslay/wslay.h` — struct wslay_frame_context and struct wslay_event_context
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)

<a id="topic-c-abi-versioned-initialization"></a>
## 48. C ABI-versioned initialization

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This protects the public structure contract against caller and library version mismatch.

Public structures are initialized through inline wrappers that pass a library ABI version to internal initialization functions before callers use the fields.

### How the implementation applies it

#### WebPConfigInit

The inline wrapper supplies WEBP_ENCODER_ABI_VERSION to WebPConfigInitInternal.

#### WebPPictureInit

Picture initialization similarly delegates to a version-checked internal entry point.

### Walk through a small source example

```c
static WEBP_INLINE int WebPConfigInit(WebPConfig* config) {
  return WebPConfigInitInternal(config, WEBP_PRESET_DEFAULT, 75.f,
                                WEBP_ENCODER_ABI_VERSION);
}
```

Source: `thirdparty/libwebp/src/webp/encode.h` — WebPConfigInit

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPConfig, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — WEBP_ENCODER_ABI_VERSION, WebPConfigInit, WebPPictureInit
- Code: `thirdparty/libwebp/src/enc/config_enc.c:27` — WebPConfigInitInternal
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)
- [C++ static ABI contracts](#topic-cpp-static-abi-contracts)
- [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop)

<a id="topic-c-aggregate-callback-modules"></a>
## 49. C aggregate state and callback modules

**Scope:** Vendored: libjpeg-turbo

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation uses public context structures plus private extension structures instead of language-level objects.

C aggregate types and typed function pointers define stateful processing modules, letting JPEG upsamplers retain shared buffers and dispatch a per-component routine.

### How the implementation applies it

#### jdsample.h: upsample1_ptr

The typedef fixes the signature for a per-component upsampling routine, including the decompression context, component metadata, and input/output sample arrays.

#### jdsample.h: my_upsampler

The private upsampler aggregates color-row buffers, an array of routine pointers, row counters, and cached expansion factors.

#### jpeglib.h: jpeg_decompress_struct

The public decompression context stores pointers to the selected decompression modules, including inverse DCT, upsampling, color conversion, and quantization.

### Walk through a small source example

```c
typedef void (*upsample1_ptr) (j_decompress_ptr cinfo,
                               jpeg_component_info *compptr,
                               _JSAMPARRAY input_data,
                               _JSAMPARRAY *output_data_ptr);
```

Source: `thirdparty/libjpeg-turbo/src/jdsample.h` — upsample1_ptr

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JPEG Decompression Context, JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jdsample.h:18` — upsample1_ptr
- Code: `thirdparty/libjpeg-turbo/src/jdsample.h:51` — my_upsampler
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming Languages: C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)
- [C explicit resource lifecycles](#topic-c-explicit-resource-lifecycles)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C stateful struct APIs](#topic-c-stateful-struct-apis)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C++ standard-library module import](#topic-cpp-module-import)
- [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop)

<a id="topic-c-integer-bitwise-packing"></a>
## 50. C fixed-width integers and bitwise packing

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code uses uint8_t through uint64_t to state intended storage widths in data and bitstream paths.

Fixed-width integer values hold packed ARGB samples, codec fields, chunk tags, and arithmetic intermediates; shifts, masks, and bitwise operators extract or assemble channel values.

### How the implementation applies it

#### lossless.c channel extraction

ARGB words are shifted and masked to form component values used by lossless transforms.

#### format_constants.h MKFOURCC

A four-character container tag is assembled with shifted byte values.

### Walk through a small source example

```c
VP8LTransformColorInverseFunc VP8LTransformColorInverse;
VP8LTransformColorInverseFunc VP8LTransformColorInverse_SSE;

VP8LConvertFunc VP8LConvertBGRAToRGB;
VP8LConvertFunc VP8LConvertBGRAToRGB_SSE;
VP8LConvertFunc VP8LConvertBGRAToRGBA;
VP8LConvertFunc VP8LConvertBGRAToRGBA_SSE;
VP8LConvertFunc VP8LConvertBGRAToRGBA4444;
VP8LConvertFunc VP8LConvertBGRAToRGB565;
VP8LConvertFunc VP8LConvertBGRAToBGR;

VP8LMapARGBFunc VP8LMapColor32b;
VP8LMapAlphaFunc VP8LMapColor8b;
```

Source: `thirdparty/libwebp/src/dsp/lossless.c` — VP8LTransformColor (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPData, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/dsp/lossless.c` — VP8LTransformColor
- Code: `thirdparty/libwebp/src/webp/format_constants.h:18` — MKFOURCC
- Code: `thirdparty/libwebp/src/webp/types.h` — uint8_t, uint16_t, uint32_t, uint64_t definitions
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scoped enums and bit flags](#topic-cxx-enums-bitflags)

<a id="topic-c-function-pointer-apis"></a>
## 51. C function-pointer API declarations

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The generated Vulkan headers provide both PFN typedefs and optionally visible direct prototypes.

C function-pointer declarations expose addressable Vulkan entry points and callback-based frame/event integration without requiring a concrete implementation type.

### How the implementation applies it

#### Dispatchable creation entry point

PFN_vkCreateWaylandSurfaceKHR specifies the complete call signature for the extension function.

#### Conditional direct declaration

The same header declares vkCreateWaylandSurfaceKHR only when VK_NO_PROTOTYPES and VK_ONLY_EXPORTED_PROTOTYPES do not suppress it.

### Walk through a small source example

```c
typedef VkResult (VKAPI_PTR *PFN_vkCreateWaylandSurfaceKHR)(VkInstance instance, const VkWaylandSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h` — PFN_vkCreateWaylandSurfaceKHR

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VkWaylandSurfaceCreateInfoKHR, VkSurfaceKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h:35` — PFN_vkCreateWaylandSurfaceKHR
- Code: `thirdparty/wslay/wslay/wslay.h` — wslay_frame_callbacks and wslay_event_callbacks
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)
- [C function-pointer callbacks](#topic-c-function-pointer-callbacks)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ lambdas and callback objects](#topic-cxx-closures)

<a id="topic-c-function-pointer-callbacks"></a>
## 52. C function-pointer callbacks

**Scope:** Vendored: enet

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The ENet callback record is initialized with standard allocation functions and can be replaced through its initialization API.

C function-pointer fields let ENet obtain allocator, deallocator, and out-of-memory handlers through a cross-cutting callbacks record.

### How the implementation applies it

#### Callback record definition

ENetCallbacks stores three callable function-pointer fields for allocation, deallocation, and allocation failure handling.

#### Runtime callback installation

enet_initialize_with_callbacks validates supplied allocation callbacks and assigns them to the process-level callbacks record.

### Walk through a small source example

```c
typedef struct _ENetCallbacks
{
    void * (ENET_CALLBACK * malloc) (size_t size);
    void (ENET_CALLBACK * free) (void * memory);
    void (ENET_CALLBACK * no_memory) (void);
} ENetCallbacks;
```

Source: `thirdparty/enet/enet/callbacks.h` — ENetCallbacks

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ENetCallbacks

Evidence:
- Code: `thirdparty/enet/enet/callbacks.h:15` — ENetCallbacks
- Code: `thirdparty/enet/callbacks.c:11` — enet_initialize_with_callbacks
- External (primary, unverified (source anchor not found)): [N1570 — Committee Draft, Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C function-pointer API declarations](#topic-c-function-pointer-apis)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ pluggable allocation](#topic-cpp-pluggable-allocation)

<a id="topic-c-macro-codecs"></a>
## 53. C macro-based binary decoding

**Scope:** Vendored: miniupnpc

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The MiniUPnP decoder offers pointer, callback-reader, and bounds-checked variants of the same 7-bit encoding loop.

C macros implement variable-length decoding by shifting accumulated values and testing continuation bits.

### How the implementation applies it

#### DECODELENGTH

The macro advances a byte pointer while the high continuation bit is set.

#### DECODELENGTH_READ

The callback form obtains each input byte through a supplied reader macro.

#### DECODELENGTH_CHECKLIMIT

The bounded form checks the input limit before consuming data.

### Walk through a small source example

```c
n = (n << 7) | (c & 0x07f);
```

Source: `thirdparty/miniupnpc/src/codelength.h` — DECODELENGTH_READ

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/miniupnpc/src/codelength.h:17` — DECODELENGTH
- Code: `thirdparty/miniupnpc/src/codelength.h:37` — DECODELENGTH_CHECKLIMIT
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ static ABI contracts](#topic-cpp-static-abi-contracts)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)

<a id="topic-c-parser-state"></a>
## 54. C parser state machines

**Scope:** Vendored: miniupnpc

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The XML parser state stores the start, end, and current input positions for the parsing routines.

C structs retain XML cursor state and parser result heads while miniUPnPc parsing functions build protocol records.

### How the implementation applies it

#### xmlparser

The parser record tracks the source boundaries and current character pointer.

#### NameValueParserData

Name/value parsing maintains a linked result head.

#### PortMappingParserData

Port-list parsing maintains a linked mapping head.

### Walk through a small source example

```c
/*! \brief data structure for parsing */
struct NameValueParserData {
	/*! \brief name/value linked list */
	struct NameValue * l_head;
	/*! \brief current element name */
	char curelt[64];
	/*! \brief port listing array */
	char * portListing;
	/*! \brief port listing array length */
	int portListingLength;
	/*! \brief flag indicating the current element is  */
	int topelt;
	/*! \brief top element character data */
	const char * cdata;
```

Source: `thirdparty/miniupnpc/include/miniupnpc/upnpreplyparse.h` — NameValueParserData (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

NameValueParserData, PortMappingParserData

Evidence:
- Code: `thirdparty/miniupnpc/src/minixml.h:17` — struct xmlparser
- Code: `thirdparty/miniupnpc/include/miniupnpc/upnpreplyparse.h:25` — NameValueParserData
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)

<a id="topic-c-pointers-arrays-and-strides"></a>
## 55. C pointers, arrays, and strides

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C resource lifecycles and ownership](#topic-c-resource-lifecycles).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Most image operations accept borrowed pointer ranges and compute row addresses instead of owning per-pixel objects.

C aggregates are reached through pointers and stepped across image memory with width- and stride-based pointer arithmetic for rows and planes.

### How the implementation applies it

#### VP8IteratorImport

The iterator derives luma and chroma source addresses from plane bases, macroblock coordinates, and distinct strides.

#### WebPDecodeYUV

The decoding API returns luma and chroma pointers plus separate luma and UV strides.

### Walk through a small source example

```c
void VP8IteratorImport(VP8EncIterator* const it, uint8_t* const tmp_32) {
  const VP8Encoder* const enc = it->enc;
  const int x = it->x, y = it->y;
  const WebPPicture* const pic = enc->pic;
  const uint8_t* const ysrc = pic->y + (y * pic->y_stride  + x) * 16;
  const uint8_t* const usrc = pic->u + (y * pic->uv_stride + x) * 8;
  const uint8_t* const vsrc = pic->v + (y * pic->uv_stride + x) * 8;
  const int w = MinSize(pic->width - x * 16, 16);
  const int h = MinSize(pic->height - y * 16, 16);
  const int uv_w = (w + 1) >> 1;
  const int uv_h = (h + 1) >> 1;

  ImportBlock(ysrc, pic->y_stride,  it->yuv_in + Y_OFF_ENC, w, h, 16);
```

Source: `thirdparty/libwebp/src/enc/iterator_enc.c` — VP8IteratorImport (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPPicture, WebPRescaler, WebPDecBuffer

Evidence:
- Code: `thirdparty/libwebp/src/enc/iterator_enc.c:136` — VP8IteratorImport
- Code: `thirdparty/libwebp/src/webp/decode.h:91` — WebPDecodeYUV
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI structures and opaque state](#topic-c-abi-data-structures)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C stateful struct APIs](#topic-c-stateful-struct-apis)
- [C structs and pointer-linked state](#topic-c-structs-pointers)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ const-reference accessors](#topic-cpp-const-reference-accessors)
- [C++ pluggable allocation](#topic-cpp-pluggable-allocation)
- [C++ typed API records](#topic-cpp-typed-api-records)
- [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop)

<a id="topic-c-preprocessor-composition"></a>
## 56. C preprocessor composition

**Scope:** Vendored: freetype

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

FreeType uses this source-level composition pattern for both its base and autofit module entry translation units.

C preprocessing defines FT_MAKE_OPTION_SINGLE_OBJECT and then includes component implementation files to form a single-object library unit.

### How the implementation applies it

#### Autofit aggregation

autofit.c defines FT_MAKE_OPTION_SINGLE_OBJECT before including the autofit implementation files.

#### Base aggregation

ftbase.c uses the same macro before including the base implementation files.

### Walk through a small source example

```c
#define FT_MAKE_OPTION_SINGLE_OBJECT

#include "ft-hb.c"
#include "ft-hb-ft.c"
#include "afadjust.c"
```

Source: `thirdparty/freetype/src/autofit/autofit.c` — FT_MAKE_OPTION_SINGLE_OBJECT

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/freetype/src/autofit/autofit.c:19` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/base/ftbase.c:19` — FT_MAKE_OPTION_SINGLE_OBJECT
- External (primary, unverified (source anchor not found)): [N1570 — Committee Draft, Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-c-preprocessor-configuration"></a>
## 57. C preprocessor feature and platform configuration

**Scope:** Vendored: spirv-reflect

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The inspected code also uses feature macros to compile TinyEXR's header implementation and platform branches in Volk.

Configuration headers use macros and conditional inclusion to enable ThorVG capabilities, select platform-dependent threading, and choose system or bundled SPIR-V headers.

### How the implementation applies it

#### ThorVG feature switches

config.h enables software raster support, SVG loading, PNG loading, and conditionally thread support.

#### Header-source selection

spirv_reflect.h chooses a system SPIR-V header when SPIRV_REFLECT_USE_SYSTEM_SPIRV_H is defined.

#### Single-header implementation switch

tinyexr.cc defines TINYEXR_IMPLEMENTATION before including tinyexr.h.

### Walk through a small source example

```c
#define THORVG_SW_RASTER_SUPPORT
#define THORVG_SVG_LOADER_SUPPORT
#define THORVG_PNG_LOADER_SUPPORT
#ifndef WEB_ENABLED
#define THORVG_THREAD_SUPPORT
#endif
```

Source: `thirdparty/thorvg/inc/config.h` — THORVG_* feature macros

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/thorvg/inc/config.h` — THORVG_SW_RASTER_SUPPORT, THORVG_SVG_LOADER_SUPPORT, THORVG_PNG_LOADER_SUPPORT, THORVG_THREAD_SUPPORT
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:34` — SPIRV_REFLECT_USE_SYSTEM_SPIRV_H
- Code: `thirdparty/tinyexr/tinyexr.cc:9` — TINYEXR_IMPLEMENTATION
- Code: `thirdparty/volk/volk.c` — _WIN32 and __APPLE__ preprocessor branches
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ static ABI contracts](#topic-cpp-static-abi-contracts)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-c-platform-selection"></a>
## 58. C preprocessor platform and precision selection

**Scope:** Vendored: libjpeg-turbo

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The same source tree supports multiple sample precisions, compilers, and CPU feature sets through macro-controlled compilation.

C conditional compilation selects precision-dependent names, optional compiler intrinsics, and architecture-specific implementations before the codec modules are compiled.

### How the implementation applies it

#### jpeg_nbits.h: USE_CLZ_INTRINSIC

Compiler and architecture macros select either a count-leading-zero intrinsic path or a lookup-table path for bit counts.

#### jsamplecomp.h: BITS_IN_JSAMPLE

Precision selection maps generic sample types and API/internal names to their 16-bit or alternate precision-specific forms.

#### libpng/intel/intel_init.c: png_init_filter_functions_sse2

When the SSE2 implementation is enabled, the initializer assigns specialized filter functions to the PNG read-filter table according to bytes per pixel.

### Walk through a small source example

```c
#ifdef USE_CLZ_INTRINSIC
#if defined(_MSC_VER) && !defined(__clang__)
#define JPEG_NBITS_NONZERO(x)  (32 - _CountLeadingZeros(x))
#else
#define JPEG_NBITS_NONZERO(x)  (32 - __builtin_clz(x))
#endif
#define JPEG_NBITS(x)          (x ? JPEG_NBITS_NONZERO(x) : 0)
```

Source: `thirdparty/libjpeg-turbo/src/jpeg_nbits.h` — JPEG_NBITS

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jpeg_nbits.h:11` — USE_CLZ_INTRINSIC
- Code: `thirdparty/libjpeg-turbo/src/jsamplecomp.h:13` — BITS_IN_JSAMPLE
- Code: `thirdparty/libpng/intel/intel_init.c:19` — png_init_filter_functions_sse2
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming Languages: C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ static ABI contracts](#topic-cpp-static-abi-contracts)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [Godot shader source composition](#topic-godot-shader-language)
- [C++ standard-library module import](#topic-cpp-module-import)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)

<a id="topic-c-preprocessor-platform-selection"></a>
## 59. C preprocessor platform selection

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SDL's source files and headers select platform code with feature macros and use extern C guards around C-facing declarations.

C preprocessor conditions select platform implementations and preserve a C linkage boundary for C++ compilation units.

### How the implementation applies it

#### SDL_build_config.h

The build configuration header centralizes feature and platform definitions used by SDL sources.

#### SDL_xinput.c

The source uses an __cplusplus guard to request C linkage when a C++ unit includes the C API declarations.

### Walk through a small source example

```c
#ifdef __cplusplus
extern "C" {
#endif
```

Source: `thirdparty/sdl/core/windows/SDL_xinput.c` — C linkage guard

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/sdl/include/build_config/SDL_build_config.h:22` — SDL_build_config_h_
- Code: `thirdparty/sdl/core/windows/SDL_xinput.c` — C linkage guard
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor composition](#topic-c-preprocessor-composition)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [Godot shader source composition](#topic-godot-shader-language)
- [C preprocessor portability](#topic-c-preprocessor-portability)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)

<a id="topic-c-preprocessor-portability"></a>
## 60. C preprocessor portability

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the portability layer that lets scalar and architecture-specific sources coexist.

Conditional compilation and macros select target-dependent code, inline helpers, endianness behavior, and instruction-family implementations at compile time.

### How the implementation applies it

#### endian_inl_utils.h HToLE32

WORDS_BIGENDIAN selects byte-swapping macros or identity macros for host-to-little-endian conversion.

#### mips_macro.h ADD_SUB_HALVES

A macro packages repeated target assembly text while preserving operand substitution.

### Walk through a small source example

```c
#if defined(WORDS_BIGENDIAN)
#define HToLE32 BSwap32
#define HToLE16 BSwap16
#else
#define HToLE32(x) (x)
#define HToLE16(x) (x)
#endif
```

Source: `thirdparty/libwebp/src/utils/endian_inl_utils.h` — HToLE32 and HToLE16

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/libwebp/src/utils/endian_inl_utils.h` — WORDS_BIGENDIAN, HToLE32, HToLE16
- Code: `thirdparty/libwebp/src/dsp/mips_macro.h:26` — ADD_SUB_HALVES
- Code: `thirdparty/libwebp/src/webp/types.h` — WEBP_INLINE conditional definition
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-c-stateful-streaming-api"></a>
## 61. C stateful streaming APIs

**Scope:** Vendored: brotli

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Brotli exposes decoder state through C interfaces while storing its operational fields in an internal structure.

C-facing code represents decoder progress as pointer-rich state and accepts caller-supplied allocation callbacks.

### How the implementation applies it

#### Allocator callback default

BrotliDefaultAllocFunc accepts an opaque caller context and allocation size, then returns malloc-managed memory.

#### Stateful output retrieval

BrotliDecoderTakeOutput accepts a BrotliDecoderState pointer so output is obtained from maintained decoding state.

#### Internal state layout

BrotliDecoderStateStruct includes byte-pointer fields used by decoding internals.

### Walk through a small source example

```c
void* BrotliDefaultAllocFunc(void* opaque, size_t size) {
  BROTLI_UNUSED(opaque);
  return malloc(size);
}
```

Source: `thirdparty/brotli/common/platform.c` — BrotliDefaultAllocFunc

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BrotliDecoderState

Evidence:
- Code: `thirdparty/brotli/common/platform.c:10` — BrotliDefaultAllocFunc
- Code: `thirdparty/brotli/dec/decode.c:2882` — BrotliDecoderTakeOutput
- Code: `thirdparty/brotli/dec/state.h:250` — BrotliDecoderStateStruct
- External (official, verified): [ISO/IEC 9899:2018 — Information technology — Programming languages — C](https://www.iso.org/standard/74528.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C structures and pointer handles](#topic-c-structures-and-pointers)

<a id="topic-c-stateful-struct-apis"></a>
## 62. C stateful struct APIs

**Scope:** Vendored: mbedtls

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

PSA, Mbed TLS TLS, and message-processing code all use this idiom.

C APIs expose mutable state through struct pointers, initialize it explicitly, and communicate failures through status-returning functions and caller-owned buffers.

### How the implementation applies it

#### psa_reset_key_attributes

The function resets a caller-provided key-attribute object with memset.

#### mbedtls_ssl_session_save

The session serializer accepts a session pointer plus an output buffer, capacity, and output-length pointer.

#### mbedtls_mps_reader

The message-processing stack defines a reader structure whose implementation separates memory loads and stores from other operations.

### Walk through a small source example

```c
void psa_reset_key_attributes(psa_key_attributes_t *attributes)
{
    memset(attributes, 0, sizeof(*attributes));
}
```

Source: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_client.c` — psa_reset_key_attributes

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

psa_key_attributes_t, mbedtls_ssl_session

Evidence:
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_client.c:17` — psa_reset_key_attributes
- Code: `thirdparty/mbedtls/library/ssl_tls.c` — mbedtls_ssl_session_save; mbedtls_ssl_session_load
- Code: `thirdparty/mbedtls/library/mps_reader.c` — mbedtls_mps_reader implementation
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [Callables and lambdas](#topic-callables-and-lambdas)
- [Node paths and indexed access](#topic-nodepaths-and-indexed-access)
- [C structs and pointer-linked state](#topic-c-structs-pointers)
- [C structures and pointer handles](#topic-c-structures-and-pointers)

<a id="topic-c-structs-pointers"></a>
## 63. C structs and pointer-linked state

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C-managed concurrent state](#topic-c-concurrent-state), [C function-pointer tables](#topic-c-function-pointer-tables).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SDL uses named and anonymous structs together with pointers for queues, device records, property storage, and platform handles.

C structs and pointers represent explicit runtime state, ownership links, and opaque-handle implementation data.

### How the implementation applies it

#### SDL_EventEntry in SDL_events.c

Each queued event embeds an SDL_Event, points at temporary payload memory, and links to previous and next queue entries.

#### SDL_EventQ in SDL_events.c

The queue state groups its lock, active flag, atomic count, head, tail, and reusable-entry list in one anonymous C struct.

#### hid_device_info in hidapi.h

The HID enumeration API returns linked device-information records through a next pointer.

### Walk through a small source example

```c
typedef struct SDL_EventEntry
{
    SDL_Event event;
    SDL_TemporaryMemory *memory;
    struct SDL_EventEntry *prev;
    struct SDL_EventEntry *next;
} SDL_EventEntry;
```

Source: `thirdparty/sdl/events/SDL_events.c` — SDL_EventEntry

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_EventEntry, SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/events/SDL_events.c:127` — SDL_EventEntry
- Code: `thirdparty/sdl/hidapi/hidapi/hidapi.h:152` — hid_device_info
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C resource lifecycles and ownership](#topic-c-resource-lifecycles)
- [C explicit resource lifecycles](#topic-c-explicit-resource-lifecycles)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C stateful struct APIs](#topic-c-stateful-struct-apis)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)

<a id="topic-c-structures-and-pointers"></a>
## 64. C structures and pointer handles

**Scope:** Vendored: freetype

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C explicit resource lifecycles](#topic-c-explicit-resource-lifecycles), [C function-pointer interfaces](#topic-c-function-pointer-interfaces).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

FreeType records commonly embed a base record and keep state behind pointer typedefs.

C structures and pointers package mutable subsystem state and expose opaque handle types.

### How the implementation applies it

#### Bzip2 stream state

FT_BZip2FileRec stores source and embedding streams, memory, bzlib state, buffers, cursors, and reset status in one record.

#### PFR face extension

PFR_FaceRec embeds FT_FaceRec as root and adds parsed PFR header, logical-font, and physical-font data.

#### SVG renderer extension

SVG_RendererRec embeds FT_RendererRec as root and adds hooks plus hook-owned state.

### Walk through a small source example

```c
typedef struct PFR_FaceRec_*  PFR_Face;

  typedef struct PFR_SizeRec_*  PFR_Size;

  typedef struct PFR_SlotRec_*  PFR_Slot;


  typedef struct  PFR_FaceRec_
  {
    FT_FaceRec      root;
    PFR_HeaderRec   header;
    PFR_LogFontRec  log_font;
    PFR_PhyFontRec  phy_font;
```

Source: `thirdparty/freetype/src/pfr/pfrobjs.h` — PFR_FaceRec_ (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FT_BZip2FileRec, PFR_FaceRec, SVG_RendererRec

Evidence:
- Code: `thirdparty/freetype/src/pfr/pfrobjs.h:27` — PFR_FaceRec_
- Code: `thirdparty/freetype/src/svg/svgtypes.h:27` — SVG_RendererRec_
- External (primary, unverified (source anchor not found)): [N1570: Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and opaque state](#topic-c-abi-data-structures)
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C ABI bridging](#topic-c-abi-bridging)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)
- [C dynamic-library function wrappers](#topic-c-dynamic-library-wrappers)
- [C function-pointer API declarations](#topic-c-function-pointer-apis)
- [C function-pointer callbacks](#topic-c-function-pointer-callbacks)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C resource lifecycles and ownership](#topic-c-resource-lifecycles)
- [C stateful streaming APIs](#topic-c-stateful-streaming-api)
- [C stateful struct APIs](#topic-c-stateful-struct-apis)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ byte-exact binary parsing](#topic-cpp-byte-exact-binary-parsing)
- [C-compatible function-pointer interfaces](#topic-cpp-function-pointer-interfaces)
- [C++ object-representation casts](#topic-cpp-object-representation-casts)
- [C++ static casts across track types](#topic-cpp-runtime-casts)
- [C++ variadic binding](#topic-cpp-variadic-binding)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)

<a id="topic-cxx-atomics"></a>
## 65. C++ atomic synchronization

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation wraps std::atomic and uses it directly in barrier and allocator-adjacent synchronization paths.

C++ atomic counters and compare-exchange loops coordinate active barriers and conditional minimum or maximum updates across workers.

### How the implementation applies it

#### Atomic wrapper

embree::atomic derives from std::atomic<T> and defines copy construction and assignment through load and store.

#### Compare-exchange update

_atomic_min and _atomic_max repeatedly load, compare, and compare_exchange_strong until the requested bound is installed or already satisfied.

#### Active barrier

BarrierActive increments an atomic counter and spins until it equals the required thread count.

### Walk through a small source example

```c
const T b = bref.load();
    while (true) {
      T a = aref.load();
      if (a <= b) break;
      if (aref.compare_exchange_strong(a,b)) break;
    }
```

Source: `thirdparty/embree/common/sys/atomic.h` — _atomic_min

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

atomic, BarrierActive, BarrierActiveAutoReset

Evidence:
- Code: `thirdparty/embree/common/sys/atomic.h:39` — _atomic_min
- Code: `thirdparty/embree/common/sys/atomic.h:6` — atomic
- Code: `thirdparty/embree/common/sys/barrier.h:38` — BarrierActive
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C-managed concurrent state](#topic-c-concurrent-state)

<a id="topic-cxx-class-hierarchy"></a>
## 66. C++ class hierarchies and reference bases

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code combines reusable base classes with specialized derived classes rather than representing all runtime objects as unrelated records.

C++ class inheritance connects shared reference ownership to acceleration, scene, geometry, builder, and scheduler abstractions.

### How the implementation applies it

#### Reference-count base

RefCount underlies retained objects; AccelData, Builder, and other abstractions inherit from it.

#### Acceleration hierarchy

AccelData is the base for acceleration data, while Accel and AccelN organize intersector and aggregate-acceleration behavior.

#### Geometry hierarchy

Geometry is a common base, and TriangleMesh, GridMesh, Instance, and other concrete forms derive from it.

### Walk through a small source example

```c
class AccelData : public RefCount
```

Source: `thirdparty/embree/kernels/common/accel.h` — AccelData

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RefCount, AccelData, Geometry, Builder

Evidence:
- Code: `thirdparty/embree/common/sys/ref.h:15` — RefCount
- Code: `thirdparty/embree/kernels/common/accel.h:16` — AccelData
- Code: `thirdparty/embree/kernels/common/geometry.h:16` — Geometry
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C resource lifecycles and ownership](#topic-c-resource-lifecycles)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [Scoped enums and bit flags](#topic-cxx-enums-bitflags)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [Java Android host APIs](#topic-java-android-host-api)
- [Properties and accessors](#topic-properties-and-accessors)

<a id="topic-cpp-class-state-and-composition"></a>
## 67. C++ class state and composition

**Scope:** Vendored: gamepadmotionhelpers

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ numeric value operations](#topic-cpp-numeric-value-operations).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The GamepadMotion helper is implemented in a C++ header containing class and struct definitions plus their behavior.

C++ classes group motion settings, calibration, vector, quaternion, and motion-update behavior behind named types.

### How the implementation applies it

#### Motion domain types

GamepadMotion.hpp declares GyroCalibration, Quat, Vec, SensorMinMaxWindow, AutoCalibration, Motion, GamepadMotionSettings, and GamepadMotion.

#### Stateful motion API

GamepadMotion is the top-level stateful type, while Motion and AutoCalibration separate orientation update and calibration responsibilities.

### Walk through a small source example

```cpp
class GamepadMotion
```

Source: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — GamepadMotion

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp:13` — GamepadMotionSettings
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp:130` — Motion
- External (primary, unverified (source anchor not found)): [N4950: Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [Scoped enums and bit flags](#topic-cxx-enums-bitflags)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [Java Android host APIs](#topic-java-android-host-api)
- [Properties and accessors](#topic-properties-and-accessors)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)

<a id="topic-cpp-classes"></a>
## 68. C++ classes and inheritance

**Scope:** Vendored: icu4c

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ member-function overloading](#topic-cpp-overloading), [C++ virtual dispatch](#topic-cpp-virtual-dispatch).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Class declarations organize most ICU and Jolt stateful services; HarfBuzz more often uses structs.

C++ classes define reusable state and behavior, and the implementation derives UnicodeString from Replaceable and many ICU services from UObject or UMemory.

### How the implementation applies it

#### UnicodeString

UnicodeString is declared as a derived class, making Replaceable operations part of its type interface.

#### LanguageBreakEngine

LanguageBreakEngine is a UObject-derived base class used by multiple break-engine implementations.

### Walk through a small source example

```c
class U_COMMON_API UnicodeString : public Replaceable
```

Source: `thirdparty/icu4c/common/unicode/unistr.h` — UnicodeString

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Locale, AABBTreeBuilder

Evidence:
- Code: `thirdparty/icu4c/common/unicode/unistr.h:63` — UnicodeString
- Code: `thirdparty/icu4c/common/brkeng.h:28` — LanguageBreakEngine
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [Scoped enums and bit flags](#topic-cxx-enums-bitflags)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [Java Android host APIs](#topic-java-android-host-api)
- [Properties and accessors](#topic-properties-and-accessors)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cxx-object-model"></a>
## 69. C++ classes, inheritance, and polymorphic interfaces

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership), [Macro-based RTTI registration](#topic-cxx-reflection-macros).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation favors abstract base types with specialized derived classes.

C++ classes and inheritance define extensible interfaces for shapes, constraints, collision queries, job systems, renderers, and body access.

### How the implementation applies it

#### Shape hierarchy

Shape derives from reference-counting and non-copyable bases, while ConvexShape, CompoundShape, mesh, terrain, and decorated types provide specialized collision behavior.

#### Read and write body locks

BodyLockRead and BodyLockWrite specialize a common lock base with const and mutable Body access.

### Walk through a small source example

```c
class JPH_EXPORT Shape : public RefTarget<Shape>, public NonCopyable
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h` — Shape

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shape, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h:36` — Shape
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/BodyLock.h` — BodyLockBase, BodyLockRead, BodyLockWrite
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [Java Android host APIs](#topic-java-android-host-api)

<a id="topic-cpp-exception-abi-boundaries"></a>
## 70. C++ exception containment at ABI boundaries

**Scope:** Vendored: openxr

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The macros are attached to exported loader functions so exceptions do not escape the loader ABI.

OpenXR loader macros use C++ exceptions internally but convert `std::bad_alloc` and other `std::exception` failures into OpenXR result codes.

### How the implementation applies it

#### XRLOADER_ABI_TRY

The macro expands to a C++ `try` block unless exception handling is disabled.

#### XRLOADER_ABI_CATCH_BAD_ALLOC_OOM

Allocation failure is logged and mapped to `XR_ERROR_OUT_OF_MEMORY`.

#### XRLOADER_ABI_CATCH_FALLBACK

Other standard exceptions are logged and mapped to `XR_ERROR_RUNTIME_FAILURE`.

### Walk through a small source example

```cpp
#define XRLOADER_ABI_TRY try
```

Source: `thirdparty/openxr/src/loader/exception_handling.hpp` — XRLOADER_ABI_TRY

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

LoaderInstance

Evidence:
- Code: `thirdparty/openxr/src/loader/exception_handling.hpp:18` — XRLOADER_ABI_CATCH_BAD_ALLOC_OOM
- Code: `thirdparty/openxr/src/loader/exception_handling.hpp:19` — XRLOADER_ABI_CATCH_FALLBACK
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-classes-and-ref-handles"></a>
## 71. C++ inheritance and engine reference handles

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied C++ code combines class hierarchies with engine-specific ownership wrappers in tests and mocks.

C++ test infrastructure uses public inheritance for test doubles and stores engine objects in Ref<T> handles created with memnew.

### How the implementation applies it

#### Display-server test double

DisplayServerMock publicly derives from DisplayServerHeadless, making the mock usable where the headless display-server interface is expected.

#### Signal observation object

SignalWatcher publicly derives from Object, placing the watcher in the engine object hierarchy.

#### Animation resource handle

The animation tests bind a memnew(Animation) result to Ref<Animation> before adding tracks.

### Walk through a small source example

```c
class DisplayServerMock : public DisplayServerHeadless {
```

Source: `tests/display_server_mock.h` — DisplayServerMock declaration

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Animation, AnimationTrack

Evidence:
- Code: `tests/display_server_mock.h:37` — DisplayServerMock
- Code: `tests/signal_watcher.h:41` — SignalWatcher
- Code: `tests/scene/test_animation.cpp` — Ref<Animation> allocation
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2024/n4981.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C ABI bridging](#topic-c-abi-bridging)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [CMake native source indexing](#topic-cmake-native-source-index)
- [C++ enumerated export state](#topic-cplusplus-enumerated-export-state)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ binary data codecs](#topic-cpp-binary-data-codecs)
- [C++ byte-exact binary parsing](#topic-cpp-byte-exact-binary-parsing)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ mapping of C API types](#topic-cpp-c-type-mapping)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ const-reference accessors](#topic-cpp-const-reference-accessors)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ exception containment at ABI boundaries](#topic-cpp-exception-abi-boundaries)
- [C++ frame-time arithmetic](#topic-cpp-frame-scheduling)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ engine binding macros](#topic-cpp-godot-binding-macros)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ JNI Variant marshalling](#topic-cpp-jni-variant-marshalling)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ plugin specialization](#topic-cpp-plugin-specialization)
- [C++ polymorphic backend classes](#topic-cpp-polymorphic-backends)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ resources and polymorphic engine objects](#topic-cpp-resource-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ inheritance and reference-counted adapters](#topic-cpp-runtime-adapters)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ scope-bound locking](#topic-cpp-scope-locking)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ tagged storage and casts](#topic-cpp-tagged-storage)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ typed API records](#topic-cpp-typed-api-records)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ atomic synchronization](#topic-cxx-atomics)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ lambdas and callback objects](#topic-cxx-closures)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [JavaScript browser FFI](#topic-javascript-browser-ffi)
- [Objective-C++ Apple platform adapters](#topic-objective-cpp-apple-adapters)
- [Objective-C++ iOS adapters](#topic-objective-cpp-ios-adapters)
- [Python build source generation](#topic-python-build-source-generation)

<a id="topic-cpp-native-wrappers"></a>
## 72. C++ inline native wrappers

**Scope:** Vendored: metal-cpp

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership), [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Metal-cpp declares lightweight C++ classes and forwards calls through typed `Object::sendMessage` invocations.

C++ wrapper classes use inheritance and inline member functions to expose native framework methods.

### How the implementation applies it

#### MTL4::Archive inheritance

`MTL4::Archive` derives from `NS::Referencing<Archive>`, establishing the wrapper's framework-object base.

#### MTL4::Archive::newBinaryFunction

The inline method binds a typed return value, selector, descriptor argument, and error out-parameter in one forwarding call.

#### MTL4::BinaryFunction::name

A similarly inline wrapper returns an `NS::String*` by sending the `name` selector.

### Walk through a small source example

```cpp
_MTL_INLINE MTL4::BinaryFunction* MTL4::Archive::newBinaryFunction(const MTL4::BinaryFunctionDescriptor* descriptor, NS::Error** error)
{
    return Object::sendMessage<MTL4::BinaryFunction*>(this, _MTL_PRIVATE_SEL(newBinaryFunctionWithDescriptor_error_), descriptor, error);
}
```

Source: `thirdparty/metal-cpp/Metal/MTL4Archive.hpp` — MTL4::Archive::newBinaryFunction

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MTL4::Archive, MTL4::BinaryFunction

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTL4Archive.hpp:60` — MTL4::Archive
- Code: `thirdparty/metal-cpp/Metal/MTL4BinaryFunction.hpp:47` — MTL4::BinaryFunction::name
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ mapping of C API types](#topic-cpp-c-type-mapping)
- [C++ static ABI contracts](#topic-cpp-static-abi-contracts)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)

<a id="topic-cpp-module-interface"></a>
## 73. C++ module interface partition

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The file is a C++ module interface artifact for the Vulkan video API namespace.

`vulkan_video.cppm` declares the exported `vulkan_hpp:video` module partition and an exported video namespace.

### How the implementation applies it

#### Partition declaration

`export module vulkan_hpp:video;` names an exported partition of the `vulkan_hpp` module.

#### Namespace export

The file exports the namespace formed by the Vulkan-Hpp namespace macros.

### Walk through a small source example

```text
export module vulkan_hpp:video;

export namespace VULKAN_HPP_NAMESPACE::VULKAN_HPP_VIDEO_NAMESPACE
{

  //=================
  //=== CONSTANTs ===
  //=================

#if defined( VULKAN_VIDEO_CODEC_H264STD_H_ )
  //=== vulkan_video_codec_h264std ===
  using VULKAN_HPP_NAMESPACE::VULKAN_HPP_VIDEO_NAMESPACE::H264CpbCntListSize;
  using VULKAN_HPP_NAMESPACE::VULKAN_HPP_VIDEO_NAMESPACE::H264MaxChromaPlanes;
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_video.cppm` — export module vulkan_hpp:video (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VideoSessionCreateInfoKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_video.cppm:24` — export module vulkan_hpp:video
- Code: `thirdparty/vulkan/include/vulkan/vulkan_video.cppm:26` — export namespace VULKAN_HPP_NAMESPACE::VULKAN_HPP_VIDEO_NAMESPACE
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

<a id="topic-cpp-object-representation-casts"></a>
## 74. C++ object-representation casts

**Scope:** Vendored: icu4c

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These operations are concentrated in code that consumes serialized binary layouts.

Pointer reinterpretation reads packed data through typed views, particularly when ICU validates or loads binary resource and Unicode data.

### How the implementation applies it

#### SpoofImpl data loading

Spoof implementation code derives a UDataInfo pointer by viewing bytes at an offset.

#### UCPTrie opening

UCPTrie code views serialized trie storage through typed pointers.

### Walk through a small source example

```cpp
const UDataInfo *pInfo = (const UDataInfo *)((const char *)inData+4);
```

Source: `thirdparty/icu4c/i18n/uspoof_impl.cpp` — SpoofImpl data loader

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UResourceBundle, UCPTrie

Evidence:
- Code: `thirdparty/icu4c/i18n/uspoof_impl.cpp` — SpoofImpl data loader
- Code: `thirdparty/icu4c/common/ucptrie.cpp:22` — ucptrie_openFromBinary
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C structures and pointer handles](#topic-c-structures-and-pointers)

<a id="topic-cpp-ownership-and-polymorphism"></a>
## 75. C++ ownership and polymorphic trees

**Scope:** Vendored: manifold

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the ownership structure behind lazily composed CSG values.

C++ class templates support CSG nodes through shared ownership, enable_shared_from_this, derived leaf and operation classes, and the Manifold implementation boundary.

### How the implementation applies it

#### CsgNode

CsgNode inherits enable_shared_from_this, allowing a shared node to recover shared ownership of itself.

#### CsgLeafNode and CsgOpNode

Leaf and operation classes derive from CsgNode to represent alternative tree-node roles.

#### Manifold::Impl

The public Manifold API declares an internal Impl type whose detailed mesh state is defined in the implementation header.

### Walk through a small source example

```c
class CsgNode : public std::enable_shared_from_this<CsgNode> {
```

Source: `thirdparty/manifold/src/csg_tree.h` — class CsgNode

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Manifold

Evidence:
- Code: `thirdparty/manifold/src/csg_tree.h` — class CsgNode; class CsgLeafNode; class CsgOpNode
- Code: `thirdparty/manifold/include/manifold/manifold.h` — Manifold::Impl declaration
- Code: `thirdparty/manifold/src/impl.h:27` — struct Manifold::Impl
- External (primary, unverified (source anchor not found)): [C++ Working Draft N4950](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C resource lifecycles and ownership](#topic-c-resource-lifecycles)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)

<a id="topic-cpp-pluggable-allocation"></a>
## 76. C++ pluggable allocation

**Scope:** Vendored: meshoptimizer

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The hook is explicit rather than relying on a global replacement of operator new.

meshoptimizer stores allocation and deallocation function pointers in static allocator storage so callers can replace the allocation policy.

### How the implementation applies it

#### meshopt_Allocator::storage

storage returns the static allocation-function storage used by the allocator wrapper.

#### meshopt_setAllocator

The setter asserts both callbacks exist and stores them in the shared allocation storage.

#### meshopt_Allocator use sites

Mesh processing functions instantiate meshopt_Allocator before allocating temporary working memory.

### Walk through a small source example

```cpp
void meshopt_setAllocator(void* (MESHOPTIMIZER_ALLOC_CALLCONV* allocate)(size_t), void (MESHOPTIMIZER_ALLOC_CALLCONV* deallocate)(void*))
```

Source: `thirdparty/meshoptimizer/allocator.cpp` — meshopt_setAllocator

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

meshopt_Meshlet

Evidence:
- Code: `thirdparty/meshoptimizer/allocator.cpp` — meshopt_Allocator::storage; meshopt_setAllocator
- Code: `thirdparty/meshoptimizer/vfetchoptimizer.cpp:33` — meshopt_Allocator allocator
- External (primary, unverified (source anchor not found)): [C++ Working Draft N4950](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C function-pointer callbacks](#topic-c-function-pointer-callbacks)
- [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop)

<a id="topic-cpp-preprocessor-configuration"></a>
## 77. C++ preprocessor configuration

**Scope:** Vendored: harfbuzz

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ static ABI contracts](#topic-cpp-static-abi-contracts).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

HarfBuzz uses feature macros such as HB_NO_COLOR and include guards throughout this partition.

Preprocessor conditions select optional HarfBuzz subsystems and header inclusion boundaries before C++ compilation.

### How the implementation applies it

#### _hb_subset_table_color

The color-table dispatcher is compiled only when HB_NO_COLOR is not defined.

#### hb-raster.hh

The raster header uses a conventional include guard.

### Walk through a small source example

```cpp
#ifndef HB_NO_COLOR
  switch (tag)
```

Source: `thirdparty/harfbuzz/src/hb-subset-table-color.cc` — _hb_subset_table_color

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-table-color.cc:8` — _hb_subset_table_color
- Code: `thirdparty/harfbuzz/src/hb-raster.hh:27` — HB_RASTER_HH
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor composition](#topic-c-preprocessor-composition)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C preprocessor portability](#topic-c-preprocessor-portability)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [Godot shader source composition](#topic-godot-shader-language)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)

<a id="topic-cxx-conditional-compilation"></a>
## 78. C++ preprocessor configuration

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The configuration is compile-time: the resulting build exposes only code paths enabled by its macro definitions.

C++ preprocessing selects tasking backends, supported geometry features, ISA namespaces, platform headers, and fallback shims before compilation.

### How the implementation applies it

#### ISA selection

sysinfo.h maps AVX-512, AVX2, AVX, and SSE feature macros to an isa namespace and ISA constant.

#### Feature gates

config.h defines or omits EMBREE_GEOMETRY_* and related macros, then maps them to IF_ENABLED_* expansion macros.

#### Task backend selection

taskscheduler.h includes an internal, TBB, or PPL scheduler header according to TASKING_* macros.

### Walk through a small source example

```c
#if defined (__AVX512VL__)
#  define isa avx512
#  define ISA AVX512
#  define ISA_STR "AVX512"
#elif defined (__AVX2__)
#  define isa avx2
#  define ISA AVX2
```

Source: `thirdparty/embree/common/sys/sysinfo.h` — ISA selection macros

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TaskScheduler, Geometry, vint8

Evidence:
- Code: `thirdparty/embree/common/sys/sysinfo.h` — ISA selection macros
- Code: `thirdparty/embree/kernels/config.h:29` — IF_ENABLED_TRIS
- Code: `thirdparty/embree/common/tasking/taskscheduler.h:6` — TASKING_INTERNAL
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ static ABI contracts](#topic-cpp-static-abi-contracts)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-cpp-single-header-implementation"></a>
## 79. C++ single-header implementation selection

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The .cpp file is intentionally minimal because VMA bodies are emitted from the included header.

VMA uses preprocessor configuration to place its header implementation in one C++ translation unit while selecting a non-MSVC debug macro before inclusion.

### How the implementation applies it

#### Implementation emission

VMA_IMPLEMENTATION is defined before vk_mem_alloc.h is included.

#### Build-mode bridge

When DEBUG_ENABLED is set outside MSVC, the translation unit defines _DEBUG.

#### Godot Vulkan dependency

The patched VMA header includes drivers/vulkan/godot_vulkan.h instead of directly including the upstream Vulkan header.

### Walk through a small source example

```cpp
#define VMA_IMPLEMENTATION
#ifdef DEBUG_ENABLED
#ifndef _MSC_VER
#define _DEBUG
#endif
#endif
#include "vk_mem_alloc.h"
```

Source: `thirdparty/vulkan/vk_mem_alloc.cpp` — translation-unit implementation configuration

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VmaAllocatorCreateInfo

Evidence:
- Code: `thirdparty/vulkan/vk_mem_alloc.cpp:1` — VMA_IMPLEMENTATION
- Code: `thirdparty/vulkan/patches/0002-VMA-godot-vulkan.patch` — vk_mem_alloc.h include replacement
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://eel.is/c++draft/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C macro-based binary decoding](#topic-c-macro-codecs)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor composition](#topic-c-preprocessor-composition)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C preprocessor portability](#topic-c-preprocessor-portability)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)

<a id="topic-cpp-module-import"></a>
## 80. C++ standard-library module import

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is a conditional module-aware inclusion path in the generated C++ headers.

The Vulkan-Hpp façade contains an import std declaration and other headers test VULKAN_HPP_CXX_MODULE before falling back to textual includes.

### How the implementation applies it

#### vulkan.hpp import std

The façade's inventory includes a standard-library module import declaration.

#### VULKAN_HPP_CXX_MODULE branch

The extension-inspection header avoids its normal includes when the Vulkan-Hpp C++ module macro is defined.

### Walk through a small source example

```cpp
import std;
```

Source: `thirdparty/vulkan/include/vulkan/vulkan.hpp` — import std

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan.hpp:41` — import std
- Code: `thirdparty/vulkan/include/vulkan/vulkan_extension_inspection.hpp:11` — VULKAN_HPP_CXX_MODULE
- External (primary, verified): [C++ Working Draft: Import declaration](https://eel.is/c++draft/module.import), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C++ resources and polymorphic engine objects](#topic-cpp-resource-and-polymorphism)
- [Python build scripts](#topic-python-build-scripts)

<a id="topic-cpp-static-generated-data"></a>
## 81. C++ static generated data

**Scope:** Vendored: harfbuzz

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The repository contains generated HarfBuzz UCD tables and generated ICU normalization, bidi, case, and character-property tables.

Static const arrays and values embed generated Unicode and normalization data directly in translation units.

### How the implementation applies it

#### norm2_nfc_data_formatVersion

Normalization data records format and data versions as static constants.

#### _hb_ucd_sc_map

HarfBuzz embeds a generated script map as a static array.

### Walk through a small source example

```c
static const UVersionInfo norm2_nfc_data_formatVersion={5,0,0,0};
static const UVersionInfo norm2_nfc_data_dataVersion={0x11,0,0,0};
```

Source: `thirdparty/icu4c/common/norm2_nfc_data.h` — norm2_nfc_data_formatVersion

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UCPTrie

Evidence:
- Code: `thirdparty/icu4c/common/norm2_nfc_data.h:12` — norm2_nfc_data_formatVersion
- Code: `thirdparty/harfbuzz/src/hb-ucd-table.hh:17` — _hb_ucd_sc_map
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [Lambdas and standard algorithms](#topic-cxx-lambdas-standard-algorithms)

<a id="topic-cpp-templates"></a>
## 82. C++ templates

**Scope:** Vendored: harfbuzz

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code uses template parameters for both general utility code and statically selected font-table implementations.

Templates parameterize data structures and algorithms, including string code-point appenders and typed OpenType table subset calls.

### How the implementation applies it

#### appendCodePoint

A StringClass template selects a code-unit representation while a boolean template parameter controls validation.

#### _hb_subset_table

Table-tag cases instantiate _hb_subset_table with a concrete OpenType table type.

### Walk through a small source example

```c
template<typename StringClass, bool validate>
inline StringClass &appendCodePoint(StringClass &s, uint32_t c) {
```

Source: `thirdparty/icu4c/common/unicode/utfstring.h` — utfstring::prv::appendCodePoint

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UCPTrie, hb_subset_plan_t

Evidence:
- Code: `thirdparty/icu4c/common/unicode/utfstring.h` — utfstring::prv::appendCodePoint
- Code: `thirdparty/harfbuzz/src/hb-subset-table-color.cc:8` — _hb_subset_table_color
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [Python SCons build hooks](#topic-python-scons-build-hooks)

<a id="topic-cxx-templates"></a>
## 83. C++ templates and specialization

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ SIMD wrapper specialization](#topic-cxx-wide-simd).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation uses type and non-type template parameters instead of a single runtime-polymorphic kernel for all widths and primitive forms.

C++ templates parameterize BVH branching, packet width, vector width, primitive layout, and builder behavior at compile time.

### How the implementation applies it

#### BVH branching factor

BVHN<N> and its builders use N as a compile-time branching factor, including builder settings that assign branchingFactor from N.

#### Width-specific API layouts

RTCRayNt, RTCHitNt, and RTCRayHitNt provide packet layouts parameterized by compile-time N.

#### Specialized builder dispatch

The two-level builder header declares specializations such as MortonBuilder<4,TriangleMesh,Triangle4> and SAHBuilder variants for concrete mesh and primitive combinations.

### Walk through a small source example

```cpp
settings.branchingFactor = N;
      settings.maxDepth = BVH::maxBuildDepthLeaf;
      return BVHBuilderBinnedSAH::build<NodeRef>
```

Source: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BVHN, RTCRayNt, MortonBuilder

Evidence:
- Code: `thirdparty/embree/kernels/bvh/bvh.h:42` — BVHN
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:229` — RTCRayNt
- Code: `thirdparty/embree/kernels/bvh/bvh_builder_twolevel_internal.h:67` — MortonBuilder
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [Python SCons build hooks](#topic-python-scons-build-hooks)

<a id="topic-cpp-templates-for-binary-data"></a>
## 84. C++ templates for binary data operations

**Scope:** Vendored: graphite

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The libraries use templates to retain typed interfaces while parsing binary font tables and copying fixed-size byte sequences.

C++ templates specialize reusable byte operations for a compile-time size and typed OpenType offsets.

### How the implementation applies it

#### unaligned_copy<S>

The LZ4 helper takes the copy size as a non-type template parameter and invokes memcpy for that fixed size.

#### OffsetTo<Type>

HarfBuzz's OpenType layer uses templated offsets to represent references from binary-table locations to typed objects.

### Walk through a small source example

```c
template<int S>
inline
void unaligned_copy(void * d, void const * s) {
  ::memcpy(d, s, S);
}
```

Source: `thirdparty/graphite/src/inc/Compression.h` — unaligned_copy

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_blob_t, hb_face_t

Evidence:
- Code: `thirdparty/graphite/src/inc/Compression.h:34` — unaligned_copy
- Code: `thirdparty/harfbuzz/src/hb-open-type.hh` — OffsetTo, UnsizedArrayOf, ArrayOf, and CFFIndex
- External (primary, unverified (source anchor not found)): [C++ working draft: Templates](https://eel.is/c++draft/temp), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ variadic binding](#topic-cpp-variadic-binding)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)

<a id="topic-cpp-typed-api-records"></a>
## 85. C++ typed API records

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ mapping of C API types](#topic-cpp-c-type-mapping), [C++ flag stringification](#topic-cpp-flag-stringification).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The C++ binding exposes API data as named records instead of raw untyped parameter lists.

Generated C++ structs use typed pointers and brace defaults to represent Vulkan creation inputs and optional state records.

### How the implementation applies it

#### Pipeline state composition

`GraphicsPipelineCreateInfo` has typed pointers to shader-stage and fixed-function state records.

#### Default-initialized pointers

Pointer members in the generated records are initialized with `{}`, representing an unset pointer value in the declaration.

#### Presentation arrays

`PresentInfoKHR` carries typed pointers for wait semaphores, swapchains, and image indices.

### Concrete structures to recognize

GraphicsPipelineCreateInfo, PresentInfoKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:56895` — struct GraphicsPipelineCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:135085` — struct PresentInfoKHR
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-abstraction-and-polymorphism"></a>
## 86. C++ types, encapsulation, and inheritance

**Scope:** Vendored: basis_universal

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C++ byte-exact binary parsing](#topic-cpp-byte-exact-binary-parsing), [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics), [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The partition uses classes as stable subsystem boundaries and exposes operations through member functions.

This C++ code models decoders, geometry engines, and allocators as named types with inheritance and encapsulated state.

### How the implementation applies it

#### Stream specialization

jpeg_decoder_file_stream and jpeg_decoder_mem_stream derive from jpeg_decoder_stream, separating file and memory input implementations behind a common decoder-stream type.

#### Geometry engine specialization

Clipper64 and ClipperD derive from ClipperBase while retaining integer and scaled floating-point result construction.

#### Encapsulated transcoder state

basisu_transcoder stores its low-level decoder members privately and exposes accessor methods for them.

### Walk through a small source example

```cpp
class png_memory_file : public png_file
```

Source: `thirdparty/basis_universal/encoder/pvpngreader.cpp` — png_memory_file

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Ktx2TranscoderState, PolyPathD

Evidence:
- Code: `thirdparty/basis_universal/encoder/pvpngreader.cpp:107` — png_memory_file
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h:459` — Clipper64
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:1` — basisu_transcoder
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [Scoped enums and bit flags](#topic-cxx-enums-bitflags)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Macro-based RTTI registration](#topic-cxx-reflection-macros)
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [Java Android host APIs](#topic-java-android-host-api)
- [Properties and accessors](#topic-properties-and-accessors)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)

<a id="topic-cpp-function-pointer-interfaces"></a>
## 87. C-compatible function-pointer interfaces

**Scope:** Vendored: icu4c

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is a data-driven dispatch mechanism used alongside C++ class polymorphism.

C-compatible function-pointer interfaces let converter implementations select operations through a shared implementation record.

### How the implementation applies it

#### UConverterSharedData::impl

Shared converter data points to an implementation record documented as a vtable-style group of function pointers.

#### UDataSwapper

The swapper structure contains function pointers for binary-data transformation operations.

### Walk through a small source example

```c
const UConverterImpl *impl;     /* vtable-style struct of mostly function pointers */
```

Source: `thirdparty/icu4c/common/ucnv_bld.h` — UConverterSharedData::impl

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UConverter, UConverterSharedData

Evidence:
- Code: `thirdparty/icu4c/common/ucnv_bld.h` — UConverterSharedData::impl
- Code: `thirdparty/icu4c/common/udataswp.h:33` — UDataSwapper
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)

<a id="topic-c-compatible-header-guards"></a>
## 88. C-compatible guarded headers

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the boundary mechanism for the platform extension's C-facing declarations.

The VI declaration file uses preprocessor guards and conditionally enters `extern "C"` when compiled as C++, so one declaration set can be included by C and C++ translation units.

### How the implementation applies it

#### Outer inclusion guard

`VULKAN_VI_H_` prevents the header body from being expanded more than once in one translation unit.

#### Conditional C++ linkage

The `__cplusplus` branch surrounds declarations with `extern "C"` only for C++ compilation.

#### Prototype availability switch

`VK_NO_PROTOTYPES` and `VK_ONLY_EXPORTED_PROTOTYPES` control whether the direct function prototype is declared.

### Walk through a small source example

```c
#ifndef VULKAN_VI_H_
#define VULKAN_VI_H_ 1
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VULKAN_VI_H_

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VkViSurfaceCreateInfoNN

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:1` — VULKAN_VI_H_
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:17` — extern "C"
- External (primary, unverified (source anchor not found)): [N1570 Committee Draft — Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor composition](#topic-c-preprocessor-composition)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C preprocessor portability](#topic-c-preprocessor-portability)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)

<a id="topic-cxx-lambdas-standard-algorithms"></a>
## 89. Lambdas and standard algorithms

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation uses lambdas for ordered curve lookup and for constructing static sampled geometry.

Lambdas provide inline comparison and initialization behavior when standard algorithms or static data construction need local callable logic.

### How the implementation applies it

#### Curve segment lookup

LinearCurve::GetValue passes a lambda comparator to std::lower_bound to locate the first point at or above an input coordinate.

#### Static sampled geometry

Vec3::sUnitSphere is initialized by a lambda that constructs a StaticArray of sample vectors.

### Walk through a small source example

```cpp
Points::const_iterator i2 = std::lower_bound(mPoints.begin(), mPoints.end(), inX, [](const Point &inPoint, float inValue) { return inPoint.mX < inValue; });
```

Source: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp` — LinearCurve::GetValue

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

LinearCurve

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp:25` — LinearCurve::GetValue
- Code: `thirdparty/jolt_physics/Jolt/Math/Vec3.cpp:44` — Vec3::sUnitSphere
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [Callables and lambdas](#topic-callables-and-lambdas)
- [C++ static generated data](#topic-cpp-static-generated-data)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)

<a id="topic-cxx-preprocessor-configuration"></a>
## 90. Preprocessor-selected platform configuration

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [SIMD wrappers and alignment](#topic-cxx-simd-alignment).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Jolt centralizes many feature flags in Core.h and applies branches in platform-specific implementation headers.

Preprocessor configuration selects processor architecture, platform APIs, floating-point precision, instruction-set support, diagnostics, and optional subsystems.

### How the implementation applies it

#### Build configuration string

GetConfigurationString conditionally emits architecture, word size, SIMD, and feature information.

#### Precision-selected real types

Real.h selects float-based or double-based Real, RVec3, and RMat44 aliases using JPH_DOUBLE_PRECISION.

### Walk through a small source example

```c
#ifdef JPH_DOUBLE_PRECISION

// Define real to double
using Real = double;
using Real3 = Double3;
using RVec3 = DVec3;
```

Source: `thirdparty/jolt_physics/Jolt/Math/Real.h` — JPH_DOUBLE_PRECISION branch

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/ConfigurationString.h:10` — GetConfigurationString
- Code: `thirdparty/jolt_physics/Jolt/Math/Real.h` — Real, RVec3, RMat44 aliases
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor composition](#topic-c-preprocessor-composition)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C preprocessor portability](#topic-c-preprocessor-portability)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-cxx-enums-bitflags"></a>
## 91. Scoped enums and bit flags

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code uses compact underlying integer types for many externally meaningful simulation modes.

Scoped enum class values model motion, body, collision, update-error, and configuration states, while selected enums provide bitwise combinations.

### How the implementation applies it

#### Motion state

EMotionType distinguishes Static, Kinematic, and Dynamic behavior.

#### Degrees of freedom

EAllowedDOFs encodes translational and rotational permissions as bit values and defines combinations such as Plane2D.

### Walk through a small source example

```c
enum class EMotionType : uint8
{
	Static,						///< Non movable
	Kinematic,					///< Movable using velocities only, does not respond to forces
	Dynamic,					///< Responds to forces as a normal physics object
};
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionType.h` — EMotionType

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body, ConstraintSettings

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionType.h:10` — EMotionType
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/AllowedDOFs.h` — EAllowedDOFs and operator |
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C fixed-width integers and bitwise packing](#topic-c-integer-bitwise-packing)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ enumerated export state](#topic-cplusplus-enumerated-export-state)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)

<a id="topic-cpp-copy-on-write-containers"></a>
## 92. C++ copy-on-write container storage

**Scope:** First-party

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Capacity growth, exact allocation, reallocation, and copy-to-new-buffer operations are explicitly represented in CowData.

The container layer uses C++ container classes to separate CowData buffer allocation and copying from Vector-style value interfaces.

### How the implementation applies it

#### Buffer growth

CowData names growth and next-capacity calculations before allocation or reallocation.

#### Copy-to-new-buffer path

CowData has a copy-to-new-buffer operation for mutations that require separate storage.

#### Vector façade

Vector exposes data, iterators, and size operations over its underlying container storage.

### Walk through a small source example

```c
const CowData prev_data;
```

Source: `core/templates/cowdata.h` — CowData::copy_from

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Array, Dictionary

Evidence:
- Code: `core/templates/cowdata.h:60` — CowData
- Code: `core/templates/vector.h:34` — Vector
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [SIMD wrappers and alignment](#topic-cxx-simd-alignment)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-scoped-locks"></a>
## 93. C++ scoped lock guards

**Scope:** First-party

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Synchronization wrappers are declared in the OS layer and reused by object and reflection code.

The implementation uses a scoped C++ MutexLock object and class-specific lock helpers around shared runtime state.

### How the implementation applies it

#### Mutex lock wrapper

MutexLock is a nodiscard class in the OS synchronization header.

#### ClassDB lock types

ClassDB declares Locker and Lock helper classes for its registry-related work.

#### Safe binary mutex specialization

MutexLock is specialized for SafeBinaryMutex tags.

### Walk through a small source example

```c
class [[nodiscard]] MutexLock {
```

Source: `core/os/mutex.h` — MutexLock

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/os/mutex.h:47` — MutexLock
- Code: `core/object/class_db.h` — ClassDB::Lock
- Code: `core/os/safe_binary_mutex.h:51` — MutexLock<SafeBinaryMutex<Tag>>
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ scope-bound locking](#topic-cpp-scope-locking)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)

<a id="topic-cpp-tagged-storage"></a>
## 94. C++ tagged storage and casts

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Variant separates the active type tag from storage-handling code for each supported runtime type.

The implementation represents runtime Variant values with a C++ type tag and accesses payloads through explicit casts and type-specific accessors.

### How the implementation applies it

#### Active type tag

Variant initializes its type member to NIL, establishing an explicit runtime type discriminator.

#### Stored transform access

Variant implementation code accesses stored Transform3D data through a type-specific internal storage member.

#### Type-specific accessors

VariantInternalAccessor has specializations for value, object, reference, packed-array, and typed-container forms.

### Walk through a small source example

```c
Type type = NIL;
```

Source: `core/variant/variant.h` — Variant::type

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant, Transform3D

Evidence:
- Code: `core/variant/variant.h` — Variant::type
- Code: `core/variant/variant.cpp:1942` — Variant::operator Transform3D
- Code: `core/variant/variant_internal.h:577` — VariantInternalAccessor
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [Types, inference, and conversions](#topic-types-inference-and-conversions)

<a id="topic-cpp-variadic-binding"></a>
## 95. C++ variadic binding

**Scope:** First-party

**Builds on:** [C++ templates and traits](#topic-cpp-templates-traits).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This lets native method and signal calls share typed, compile-time-generated argument marshalling.

The binding code uses variadic templates and parameter packs to move argument lists into Variant pointer arrays.

### How the implementation applies it

#### Method-definition names

MethodDefinition materializes a parameter pack into a null-terminated argument-name array.

#### Object call arguments

Object templates allocate Variant argument and pointer arrays sized with sizeof...(p_args).

#### Callable argument arrays

Callable templates likewise form Variant pointer arrays before dispatch.

### Walk through a small source example

```c
const char *args[sizeof...(p_args) + 1] = { p_args..., nullptr }; // +1 makes sure zero sized arrays are also supported.
```

Source: `core/object/class_db.h` — MethodDefinition

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant, Object

Evidence:
- Code: `core/object/class_db.h:66` — MethodDefinition
- Code: `core/object/object.h` — Object::call
- Code: `core/variant/callable.h` — Callable::call
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [Types, inference, and conversions](#topic-types-inference-and-conversions)

<a id="topic-cpp-runtime-casts"></a>
## 96. C++ static casts across track types

**Scope:** First-party

**Builds on:** [C++ class inheritance](#topic-cpp-object-hierarchies).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Animation evaluation repeats this pattern for several concrete track classes.

The implementation uses static_cast to interpret a base Track pointer as a derived Track object after a track kind has been selected.

### How the implementation applies it

#### Value track selection

Animation code casts a selected Track pointer to ValueTrack before accessing value-track keys.

#### Other typed tracks

The same evaluation area casts Track pointers to BezierTrack, AudioTrack, and AnimationTrack variants.

### Walk through a small source example

```cpp
const ValueTrack *vt = static_cast<const ValueTrack *>(tracks[track]);
```

Source: `scene/resources/animation.cpp` — Animation track evaluation

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Animation

Evidence:
- Code: `scene/resources/animation.cpp:674` — static_cast<const ValueTrack *>(tracks[track])
- Code: `scene/resources/animation.cpp` — static_cast selections for BezierTrack, AudioTrack, and AnimationTrack
- External (primary, verified): [Working Draft, Programming Languages — C++: Static cast](https://eel.is/c++draft/expr.static.cast), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C structures and pointer handles](#topic-c-structures-and-pointers)

<a id="topic-annotations-static-state-and-lifecycle"></a>
## 97. Annotations, static state, and lifecycle

**Scope:** First-party

**Builds on:** [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The suite covers annotation validity, tool-mode requirements, static initialization, and ready-time values.

Annotations modify class declarations and members, while static state and `@onready` participate in initialization and lifecycle behavior.

### How the implementation applies it

#### Ready-time member

An inner `Node` subclass declares an `@onready` member even when the outer script extends `RefCounted`.

#### Static initialization

Static-variable fixtures observe static initialization and state across loaded scripts and nested classes.

### Walk through a small source example

```text
class Inner extends Node:
	@onready var okay = 0
```

Source: `modules/gdscript/tests/scripts/analyzer/features/onready_on_inner_class_with_non_node_outer.gd` — Inner

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/onready_on_inner_class_with_non_node_outer.gd:6` — Inner
- Code: `modules/gdscript/tests/scripts/runtime/features/static_variables.gd` — Inner and InnerInner
- Code: `modules/gdscript/tests/scripts/parser/errors/export_tool_button_requires_tool_mode.out` — expected analyzer error
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ pluggable allocation](#topic-cpp-pluggable-allocation)
- [C++ static casts across track types](#topic-cpp-runtime-casts)
- [C++ static generated data](#topic-cpp-static-generated-data)
- [Lambdas and standard algorithms](#topic-cxx-lambdas-standard-algorithms)
- [Properties and accessors](#topic-properties-and-accessors)
- [Typed arrays and dictionaries](#topic-typed-collections)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)

<a id="topic-csharp-source-generation"></a>
## 98. C# partial types and source generation

**Scope:** First-party

**Builds on:** [C# attributes and reflection](#topic-csharp-attributes-reflection).

**This prepares you for:** [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Godot's analyzers and generators inspect C# syntax and semantic symbols, report invalid declarations, and add generated source to the compilation.

Compilation produces attribute-bearing partial declarations for script paths and assembly script types.

### How the implementation applies it

#### Script path generation

ScriptPathAttributeGenerator.Execute discovers Godot script classes, rejects non-partial classes, emits ScriptPathAttribute declarations, and emits an AssemblyHasScriptsAttribute source file.

#### Member metadata generation

ScriptPropertiesGenerator, ScriptMethodsGenerator, and ScriptSignalsGenerator generate bridge-facing metadata from compatible members and delegates.

#### Unmanaged callback generation

UnmanagedCallbacksGenerator examines partial method definitions and emits callback trampolines and matching function-structure members.

### Walk through a small source example

```csharp
[Generator]
public class UnmanagedCallbacksGenerator : ISourceGenerator
{
    public void Initialize(GeneratorInitializationContext context)
    {
        context.RegisterForPostInitialization(ctx => { GenerateAttribute(ctx); });
    }

    public void Execute(GeneratorExecutionContext context)
    {
        INamedTypeSymbol[] unmanagedCallbacksClasses = context
            .Compilation.SyntaxTrees
            .SelectMany(tree =>
                tree.GetRoot().DescendantNodes()
```

Source: `modules/mono/glue/GodotSharp/Godot.SourceGenerators.Internal/UnmanagedCallbacksGenerator.cs` — UnmanagedCallbacksGenerator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ScriptPathAttribute, AssemblyHasScriptsAttribute, ManagedCallbacks

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPathAttributeGenerator.cs` — ScriptPathAttributeGenerator.Execute
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/Common.cs` — Common.ClassPartialModifierRule
- Code: `modules/mono/glue/GodotSharp/Godot.SourceGenerators.Internal/UnmanagedCallbacksGenerator.cs:11` — UnmanagedCallbacksGenerator
- External (official, unverified (source anchor not found)): [Partial type (C# Reference)](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/partial-type), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C# partial types and source-generator tests](#topic-csharp-partial-source-generation)
- [Python build code generation](#topic-python-build-code-generation)
- [Python build source generation](#topic-python-build-source-generation)

<a id="topic-properties-and-accessors"></a>
## 99. Properties and accessors

**Scope:** First-party

**Builds on:** [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Tests cover direct assignment, backing variables, getter/setter syntax, and compound access chains.

Class properties route reads and writes through inline or named accessors, including static and chained access.

### How the implementation applies it

#### Named accessor binding

A property delegates its getter and setter to separately declared methods.

#### Inline backing property

An inline property returns and updates a separate backing field.

### Walk through a small source example

```text
var prop:
	get = get_prop, set = set_prop

func get_prop():
	return _prop
```

Source: `modules/gdscript/tests/scripts/analyzer/features/property_functions.gd` — prop and get_prop()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/property_functions.gd:2` — prop
- Code: `modules/gdscript/tests/scripts/analyzer/features/property_inline.gd:24` — prop4
- Code: `modules/gdscript/tests/scripts/runtime/features/setter_chain_shared_types.gd` — prop1, prop2, prop3, and prop4
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C# partial types and source-generator tests](#topic-csharp-partial-source-generation)

<a id="topic-signals-and-await"></a>
## 100. Signals and await

**Scope:** First-party

**Builds on:** [Callables and lambdas](#topic-callables-and-lambdas).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Fixtures distinguish meaningful awaits from redundant ones and cover signal payload shapes.

Signal values and callable functions support waiting for emissions or coroutine completion with `await`.

### How the implementation applies it

#### Signal identity

Signals declared at outer, nested, and inherited classes are compared for identity.

#### Await payload shape

Awaiting signals with zero, one, and two parameters produces distinct returned values in the test output.

### Walk through a small source example

```text
func await_one_parameter():
	var result = await one_parameter
```

Source: `modules/gdscript/tests/scripts/runtime/features/await_signal_with_parameters.gd` — await_one_parameter()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/runtime/features/await_signal_with_parameters.gd:9` — await_one_parameter()
- Code: `modules/gdscript/tests/scripts/analyzer/features/lookup_signal.gd:3` — get_signal()
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/redundant_await.out:2` — REDUNDANT_AWAIT
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)
- [C# partial types and source-generator tests](#topic-csharp-partial-source-generation)
- [Warnings and selective suppression](#topic-warnings-and-suppression)

<a id="topic-typed-collections"></a>
## 101. Typed arrays and dictionaries

**Scope:** First-party

**Builds on:** [Types, inference, and conversions](#topic-types-inference-and-conversions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Collection shape is tested both during analysis and during runtime assignment or argument passing.

Typed arrays and dictionaries apply static typing to element and key/value shapes; the fixture suite exercises declarations, inference, casts, defaults, transfers, and invalid runtime assignments.

### How the implementation applies it

#### Custom-class array element

`Array[Inner]` stores an inner-class instance and returns it through an `Inner`-typed local.

#### Typed-dictionary failure paths

Runtime error cases distinguish untyped dictionaries from dictionaries whose key and value types are constrained.

### Walk through a small source example

```text
class Inner:
	var prop = "Inner"


var array: Array[Inner] = [Inner.new()]
```

Source: `modules/gdscript/tests/scripts/analyzer/features/typed_array_with_custom_class.gd` — Inner and array

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/typed_array_with_custom_class.gd:5` — array
- Code: `modules/gdscript/tests/scripts/analyzer/features/typed_dictionary_with_custom_class.gd:4` — dict
- Code: `modules/gdscript/tests/scripts/runtime/errors/typed_dictionary.out` — TypedDictionary.Key and TypedDictionary.Value errors
- External (official, unverified (source anchor not found)): [Array class reference](https://docs.godotengine.org/en/stable/classes/class_array.html), accessed 2026-07-15
- External (official, unverified (source anchor not found)): [Dictionary class reference](https://docs.godotengine.org/en/stable/classes/class_dictionary.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Annotations, static state, and lifecycle](#topic-annotations-static-state-and-lifecycle)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [GDScript typed collections and signature compatibility](#topic-gdscript-typed-collections)

<a id="topic-c-dynamic-library-wrappers"></a>
## 102. C dynamic-library function wrappers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The generated wrapper headers declare replacement function-pointer symbols and identify the wrapped shared-library sonames.

Generated C function-pointer wrappers dynamically route DBus, Fontconfig, Speech Dispatcher, and Wayland/libdecor calls for Linux/BSD desktop integration.

### How the implementation applies it

#### dbus-so_wrap.c

The DBus wrapper declares function pointers matching DBus API signatures.

#### fontconfig-so_wrap.h

The Fontconfig wrapper header identifies itself as generated and redirects API names through wrapper symbols.

#### wayland-client-core-so_wrap.c

The Wayland wrapper declares function pointers for core Wayland client calls such as display connection and event-queue creation.

### Walk through a small source example

```c
const char* (*dbus_address_entry_get_value_dylibloader_wrapper_dbus)( DBusAddressEntry*,const char*);
```

Source: `platform/linuxbsd/dbus-so_wrap.c` — dbus_address_entry_get_value_dylibloader_wrapper_dbus

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/linuxbsd/dbus-so_wrap.c:495` — dbus_address_entry_get_value_dylibloader_wrapper_dbus
- Code: `platform/linuxbsd/fontconfig-so_wrap.h:1` — DYLIBLOAD_WRAPPER_FONTCONFIG
- Code: `platform/linuxbsd/wayland/dynwrappers/wayland-client-core-so_wrap.c:152` — wl_display_connect_dylibloader_wrapper_wayland_client
- External (official, unverified (source anchor not found)): [ISO/IEC 9899:2024 - Programming languages — C](https://www.iso.org/standard/82075.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C structures and pointer handles](#topic-c-structures-and-pointers)

<a id="topic-cpp-jni-variant-marshalling"></a>
## 103. C++ JNI Variant marshalling

**Scope:** First-party

**Builds on:** [C++ inheritance and reference-counted adapters](#topic-cpp-runtime-adapters).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The native Android wrapper identifies Java values through JNI calls and produces Godot `Variant` representations.

C++ JNI Variant marshalling builds on RefCounted adapter classes to convert Java primitive wrappers, strings, arrays, objects, and callables into engine Variant values.

### How the implementation applies it

#### JavaClassWrapper::_jobject_to_variant

The conversion implementation calls Java wrapper methods such as `Integer.intValue()` through JNI and assigns the result to a Variant.

#### jcallable_to_callable

The conversion source identifies a helper that turns a Java callable into an engine callable representation.

#### callable_jni

The callable JNI source reconstructs Variant argument pointers before executing a native callable.

### Walk through a small source example

```cpp
var = env->CallIntMethod(obj, JavaClassWrapper::singleton->Integer_integerValue);
```

Source: `platform/android/java_class_wrapper.cpp` — JavaClassWrapper::_jobject_to_variant

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Callable

Evidence:
- Code: `platform/android/java_class_wrapper.cpp` — JavaClassWrapper::_jobject_to_variant
- Code: `platform/android/java_class_wrapper.cpp:1071` — jcallable_to_callable
- Code: `platform/android/variant/callable_jni.cpp:2` — callable_jni
- External (primary, unverified (source anchor not found)): [Draft C++ Standard: Contents](https://eel.is/c++draft/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [Types, inference, and conversions](#topic-types-inference-and-conversions)

<a id="topic-cpp-godot-binding-macros"></a>
## 104. C++ engine binding macros

**Scope:** First-party

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These macros are a Godot-specific layer over C++ declarations used throughout editor extension points.

Godot's C++ classes declare engine binding and script-overridable callbacks through macros such as GDCLASS and GDVIRTUAL.

### How the implementation applies it

#### Class declaration marker

EditorPlugin invokes GDCLASS immediately inside its class declaration.

#### Script-visible importer hooks

EditorSceneFormatImporter declares required virtual extension and import callbacks with GDVIRTUAL macros.

#### Inspector parser hooks

EditorInspectorPlugin declares script-overridable handling and parsing callbacks with GDVIRTUAL macros.

### Walk through a small source example

```c
GDCLASS(EditorPlugin, Node);
```

Source: `editor/plugins/editor_plugin.h` — EditorPlugin

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/plugins/editor_plugin.h:59` — EditorPlugin
- Code: `editor/import/3d/resource_importer_scene.h:49` — EditorSceneFormatImporter
- Code: `editor/inspector/editor_inspector.h:341` — EditorInspectorPlugin
- External (primary, verified): [C++ Working Draft — Preprocessing directives](https://eel.is/c++draft/cpp), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C function-pointer tables](#topic-c-function-pointer-tables)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-plugin-specialization"></a>
## 105. C++ plugin specialization

**Scope:** First-party

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-and-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The pattern is repeated across scene, script, shader, resource, and version-control integrations.

C++ classes specialize EditorPlugin for feature-specific integrations; this requires C++ classes as the base abstraction mechanism.

### How the implementation applies it

#### Scene plugin specialization

MeshLibraryEditorPlugin derives from EditorPlugin for mesh-library editing.

#### Script plugin specialization

EditorScriptPlugin derives from EditorPlugin for editor-script integration.

#### Shader plugin specialization

ShaderEditorPlugin derives from EditorPlugin for shader authoring integration.

### Walk through a small source example

```c
class MeshLibraryEditorPlugin : public EditorPlugin {
```

Source: `editor/scene/3d/mesh_library_editor_plugin.h` — MeshLibraryEditorPlugin

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/3d/mesh_library_editor_plugin.h:126` — MeshLibraryEditorPlugin : public EditorPlugin
- Code: `editor/script/editor_script_plugin.h:35` — EditorScriptPlugin : public EditorPlugin
- Code: `editor/shader/shader_editor_plugin.h:48` — ShaderEditorPlugin : public EditorPlugin
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-runtime-polymorphism"></a>
## 106. C++ virtual interfaces

**Scope:** First-party

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the main extension idiom in the inspected editor code.

Godot uses polymorphic base interfaces so importers, inspector plug-ins, preview generators, and editor plug-ins can supply type-specific behavior through virtual methods.

### How the implementation applies it

#### Inspector eligibility

EditorInspectorPlugin exposes can_handle and parsing hooks so an implementation can select and build controls for an edited Object.

#### Scene-format operation

EditorSceneFormatImporter declares virtual extension discovery, scene import, option collection, and visibility methods.

#### Editor lifecycle hooks

EditorPlugin declares virtual methods for editing objects, reacting to UI input, saving state, building, and running scenes.

### Walk through a small source example

```c
virtual bool can_handle(Object *p_object);
```

Source: `editor/inspector/editor_inspector.h` — EditorInspectorPlugin::can_handle

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/inspector/editor_inspector.h:341` — EditorInspectorPlugin
- Code: `editor/import/3d/resource_importer_scene.h:49` — EditorSceneFormatImporter
- Code: `editor/plugins/editor_plugin.h:59` — EditorPlugin
- External (primary, verified): [C++ Working Draft — Virtual functions](https://eel.is/c++draft/class.virtual), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-objective-cpp-apple-adapters"></a>
## 107. Objective-C++ Apple platform adapters

**Scope:** First-party

**Builds on:** [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

macOS and visionOS implementation files use .mm sources where C++ engine classes meet Apple platform APIs.

Objective-C++ implementation files combine C++ platform adapters with Cocoa-style objects, display coordinates, events, menus, views, and text-to-speech calls.

### How the implementation applies it

#### DisplayServerMacOS::_get_hdr_output

The macOS display server implements HDR-output lookup in an Objective-C++ source file.

#### KeyMappingMacOS

macOS key mapping accesses keyboard-layout data and maps keycodes, key symbols, and key locations.

#### NativeMenuMacOS

The macOS native-menu adapter stores MenuData and interacts with NSMenu and NSMenuItem objects.

### Walk through a small source example

```text
const NSPoint mouse_pos = [NSEvent mouseLocation];
```

Source: `platform/macos/display_server_macos.mm` — DisplayServerMacOS::mouse_get_position

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/macos/display_server_macos.mm:1381` — DisplayServerMacOS::mouse_get_position
- Code: `platform/macos/key_mapping_macos.mm:52` — KeyMappingMacOS
- Code: `platform/macos/native_menu_macos.h:40` — NativeMenuMacOS
- External (official, unverified (source anchor not found)): [About Objective-C](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html), accessed 2026-07-15
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [Objective-C++ iOS adapters](#topic-objective-cpp-ios-adapters)

<a id="topic-objective-cpp-ios-adapters"></a>
## 108. Objective-C++ iOS adapters

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The iOS SCons build compiles `.mm` sources for device metrics, display, view, main, and OS behavior.

Objective-C++ iOS adapters implement an embedded platform layer for the iOS embedded adapter.

### How the implementation applies it

#### platform/ios/SCsub

The iOS library source list includes multiple `.mm` files.

#### DisplayServerIOS

The iOS display-server class is declared as a final subclass of `DisplayServerAppleEmbedded`.

#### display_server_ios.mm

The implementation source contains a `utsname` system-information value.

### Walk through a small source example

```text
struct utsname systemInfo;
```

Source: `platform/ios/display_server_ios.mm` — systemInfo

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/ios/SCsub:10` — ios_lib
- Code: `platform/ios/display_server_ios.h:35` — DisplayServerIOS
- Code: `platform/ios/display_server_ios.mm:68` — systemInfo
- External (official, verified): [About Objective-C](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [Objective-C++ Apple platform adapters](#topic-objective-cpp-apple-adapters)

<a id="topic-python-build-source-generation"></a>
## 109. Python build source generation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The functions produce text and byte-oriented C++ output rather than editor runtime behavior.

Python build helpers perform build-time source generation by transforming inputs into generated C++ source and invoking external translation tooling when available.

### How the implementation applies it

#### Documentation path table

doc_data_class_path_builder reads and orders paths before emitting a C++ class-path table.

#### Exporter registration

register_exporters_builder emits includes and registration code from platform input.

#### Compressed documentation

make_doc_header compresses documentation data and emits constexpr C++ data.

#### Translation compilation

make_translations uses temporary files and subprocess execution for msgfmt before emitting translation tables.

### Walk through a small source example

```python
def make_translations(target, source, env):
```

Source: `editor/editor_builders.py` — make_translations

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorTranslationList

Evidence:
- Code: `editor/editor_builders.py:12` — doc_data_class_path_builder
- Code: `editor/editor_builders.py:32` — register_exporters_builder
- Code: `editor/editor_builders.py:55` — make_doc_header
- Code: `editor/editor_builders.py:71` — make_translations
- External (official, unverified (source anchor not found)): [subprocess — Subprocess management — Python 3.14.6 documentation](https://docs.python.org/3/library/subprocess.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C# partial types and source generation](#topic-csharp-source-generation)
- [Python build-action functions](#topic-python-build-actions)
- [Python build code generation](#topic-python-build-code-generation)
- [Python build hooks](#topic-python-build-hooks)
- [Python build scripts](#topic-python-build-scripts)
- [Python SCons configuration functions](#topic-python-scons-configuration)

<a id="topic-python-build-actions"></a>
## 110. Python build-action functions

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Python is used for editor build support, not the editor’s primary runtime implementation.

Python defines the icon-build action that receives target, source, and environment arguments, while SCsub scripts invoke Python-based build-environment methods.

### How the implementation applies it

#### Build action definition

make_editor_icons_action is a Python function with target, source, and env parameters.

#### Build-script calls

SCsub scripts call env.add_source_files to contribute C++ files to env.editor_sources.

### Walk through a small source example

```python
def make_editor_icons_action(target, source, env):
```

Source: `editor/icons/editor_icons_builders.py` — make_editor_icons_action

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/icons/editor_icons_builders.py:9` — make_editor_icons_action
- Code: `editor/import/SCsub:6` — env.add_source_files
- External (official, verified): [Python documentation — Defining Functions](https://docs.python.org/3/tutorial/controlflow.html#defining-functions), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Python build code generation](#topic-python-build-code-generation)
- [Python build hooks](#topic-python-build-hooks)
- [Python build scripts](#topic-python-build-scripts)
- [Python build source generation](#topic-python-build-source-generation)
- [Python SCons build hooks](#topic-python-scons-build-hooks)
- [Python SCons configuration functions](#topic-python-scons-configuration)
- [Python SCons module configuration](#topic-python-scons-module-configuration)

<a id="topic-c-abi-record-and-dispatch"></a>
## 111. C ABI records and dispatch signatures

**Scope:** Vendored: vulkan

**Builds on:** [C-compatible guarded headers](#topic-c-compatible-header-guards).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The record is supplied by a caller; the dispatch signature names instance, creation input, allocation callbacks, and an output surface pointer.

`VkViSurfaceCreateInfoNN` and `PFN_vkCreateViSurfaceNN` express ABI-shaped call parameters through a tagged C record and a typed creation-function pointer.

### How the implementation applies it

#### Tagged input record

The record contains `sType`, `pNext`, flags, and the platform window pointer.

#### Dispatchable function type

The `PFN_` typedef gives the creation call a reusable typed function-pointer signature.

### Walk through a small source example

```c
typedef VkResult (VKAPI_PTR *PFN_vkCreateViSurfaceNN)(VkInstance instance, const VkViSurfaceCreateInfoNN* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — PFN_vkCreateViSurfaceNN

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VkViSurfaceCreateInfoNN, VkSurfaceKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:27` — VkViSurfaceCreateInfoNN
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:34` — PFN_vkCreateViSurfaceNN
- External (primary, unverified (source anchor not found)): [N1570 Committee Draft — Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI-versioned initialization](#topic-c-abi-versioned-initialization)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C function-pointer API declarations](#topic-c-function-pointer-apis)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C-compatible function-pointer interfaces](#topic-cpp-function-pointer-interfaces)

<a id="topic-c-explicit-resource-lifecycles"></a>
## 112. C explicit resource lifecycles

**Scope:** Vendored: freetype

**Builds on:** [C structures and pointer handles](#topic-c-structures-and-pointers).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The Bzip2 adapter names initialization, reset, finalization, I/O, and close functions separately and wires them into FT_Stream.

C functions establish, reset, and release C struct state explicitly through paired lifecycle operations.

### How the implementation applies it

#### Initialize then publish

FT_Stream_OpenBzip2 allocates a wrapper, initializes it, and stores it in stream->descriptor.pointer only after success.

#### Backward-seek reset

ft_bzip2_file_io resets decompression when a requested position precedes the current output position or reset is flagged.

#### Close owns wrapper cleanup

ft_bzip2_stream_close calls ft_bzip2_file_done, frees the wrapper, and clears descriptor.pointer.

### Walk through a small source example

```c
static void
  ft_bzip2_file_done( FT_BZip2File  zip )
  {
    bz_stream*  bzstream = &zip->bzstream;


    BZ2_bzDecompressEnd( bzstream );

    /* clear the rest */
    bzstream->bzalloc   = NULL;
    bzstream->bzfree    = NULL;
    bzstream->opaque    = NULL;
    bzstream->next_in   = NULL;
    bzstream->next_out  = NULL;
```

Source: `thirdparty/freetype/src/bzip2/ftbzip2.c` — ft_bzip2_file_done (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:195` — ft_bzip2_file_done
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:439` — ft_bzip2_stream_close
- External (primary, unverified (source anchor not found)): [N1570: Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C structs and pointer-linked state](#topic-c-structs-pointers)

<a id="topic-c-function-pointer-interfaces"></a>
## 113. C function-pointer interfaces

**Scope:** Vendored: freetype

**Builds on:** [C structures and pointer handles](#topic-c-structures-and-pointers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation assigns allocation and I/O functions into external library or stream interface records.

C function pointers attach behavior to struct-and-pointer handles without requiring a language-level object system.

### How the implementation applies it

#### Bzip2 allocator hooks

ft_bzip2_file_init assigns ft_bzip2_alloc and ft_bzip2_free to bz_stream callbacks.

#### Stream callbacks

FT_Stream_OpenBzip2 installs ft_bzip2_stream_io and ft_bzip2_stream_close in the target stream.

#### Cache extension callbacks

The cache callback header declares new, free, weight, and comparison functions for image, small-bitmap, and generic cache nodes.

### Walk through a small source example

```c
FT_LOCAL( FT_Error )
  ftc_inode_new( FTC_Node   *pinode,
                 FT_Pointer  gquery,
                 FTC_Cache   cache );

  FT_LOCAL( FT_Offset )
  ftc_inode_weight( FTC_Node   inode,
                    FTC_Cache  cache );


  FT_LOCAL( void )
  ftc_snode_free( FTC_Node   snode,
                  FTC_Cache  cache );
```

Source: `thirdparty/freetype/src/cache/ftccback.h` — ftc_inode_new (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/cache/ftccback.h:31` — ftc_inode_new
- External (primary, unverified (source anchor not found)): [N1570: Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C ABI records and dispatch signatures](#topic-c-abi-record-and-dispatch)
- [C function-pointer API declarations](#topic-c-function-pointer-apis)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C structs and pointer-linked state](#topic-c-structs-pointers)
- [C# unsafe interop and function pointers](#topic-csharp-unsafe-interop)

<a id="topic-c-function-pointer-tables"></a>
## 114. C function-pointer tables

**Scope:** Vendored: sdl

**Builds on:** [C structs and pointer-linked state](#topic-c-structs-pointers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SDL declares native-library symbols and device-driver operations as typed function pointers.

C function-pointer tables attach implementation callbacks to C struct state for dynamically selected system and device services.

### How the implementation applies it

#### SDL_UDEV_Symbols in SDL_udev.h

The UDev adapter stores function pointers for libudev operations such as action and device-node queries.

#### hidapi_backend in SDL_hidapi.c

The HID layer defines backend operations behind function pointers and selects a backend per device.

### Walk through a small source example

```c
struct hidapi_backend
{
    int (*hid_write)(void *device, const unsigned char *data, size_t length);
    int (*hid_read_timeout)(void *device, unsigned char *data, size_t length, int milliseconds);
    int (*hid_read)(void *device, unsigned char *data, size_t length);
    int (*hid_set_nonblocking)(void *device, int nonblock);
    int (*hid_send_feature_report)(void *device, const unsigned char *data, size_t length);
    int (*hid_get_feature_report)(void *device, unsigned char *data, size_t length);
    int (*hid_get_input_report)(void *device, unsigned char *data, size_t length);
    void (*hid_close)(void *device);
    int (*hid_get_manufacturer_string)(void *device, wchar_t *string, size_t maxlen);
    int (*hid_get_product_string)(void *device, wchar_t *string, size_t maxlen);
    int (*hid_get_serial_number_string)(void *device, wchar_t *string, size_t maxlen);
```

Source: `thirdparty/sdl/hidapi/SDL_hidapi.c` — hidapi_backend (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/core/linux/SDL_udev.h:56` — SDL_UDEV_Symbols
- Code: `thirdparty/sdl/hidapi/SDL_hidapi.c:923` — hidapi_backend
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C function-pointer callbacks](#topic-c-function-pointer-callbacks)
- [C stateful streaming APIs](#topic-c-stateful-streaming-api)
- [C++ engine binding macros](#topic-cpp-godot-binding-macros)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [JavaScript browser FFI](#topic-javascript-browser-ffi)
- [JavaScript browser runtime libraries](#topic-javascript-web-runtime)
- [C structures and pointer handles](#topic-c-structures-and-pointers)

<a id="topic-c-resource-lifecycles"></a>
## 115. C resource lifecycles and ownership

**Scope:** Vendored: libwebp

**Builds on:** [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Resource retirement is explicit; structure storage itself is generally caller-owned.

Pointer-bearing structures use explicit init, allocate, clear, free, copy, and view functions so the caller can distinguish owned image buffers from borrowed views.

### How the implementation applies it

#### WebPPictureAlloc and WebPPictureFree

The picture API allocates or releases pixel memory while retaining the caller's structure object.

#### WebPMemoryWriterInit and WebPMemoryWriterClear

The memory writer initializes and later releases its owned output allocation.

### Walk through a small source example

```c
int WebPPictureAllocARGB(WebPPicture* const picture) {
  void* memory;
  const int width = picture->width;
  const int height = picture->height;
  const uint64_t argb_size = (uint64_t)width * height;

  if (!WebPValidatePicture(picture)) return 0;

  WebPSafeFree(picture->memory_argb_);
  WebPPictureResetBufferARGB(picture);

  // allocate a new buffer.
  memory = WebPSafeMalloc(argb_size + WEBP_ALIGN_CST, sizeof(*picture->argb));
```

Source: `thirdparty/libwebp/src/enc/picture_enc.c` — WebPPictureAlloc (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPPicture, WebPMemoryWriter, WebPDecBuffer

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — WebPPictureAlloc, WebPPictureFree, WebPPictureView, WebPMemoryWriterClear
- Code: `thirdparty/libwebp/src/enc/picture_enc.c:167` — WebPPictureAlloc
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C structs and pointer-linked state](#topic-c-structs-pointers)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C structures and pointer handles](#topic-c-structures-and-pointers)

<a id="topic-cpp-basis-transcoding"></a>
## 116. C++ Basis transcoding behind a C-compatible API

**Scope:** Vendored: libktx

**Builds on:** [C aggregate state and callback modules](#topic-c-aggregate-callback-modules).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the material C++ portion of the supplied KTX implementation; most neighboring public and internal KTX APIs are C declarations.

The KTX Basis transcode implementation uses C++ references and standard-library state vectors, while its headers retain a C-compatible interface for the enclosing C API.

### How the implementation applies it

#### basis_transcode.cpp: ktxTexture2_transcodeEtc1s

The ETC1S path binds each image description as a C++ const reference before processing it.

#### basis_transcode.cpp: ktxTexture2_transcodeUastc

The UASTC path constructs a Basis Universal transcoder and maintains transcoder state in a std::vector.

#### basis_sgd.h: __cplusplus linkage guard

The shared global-data header surrounds its C declarations with an extern "C" guard when included from C++.

### Walk through a small source example

```cpp
const ktxBasisLzEtc1sImageDesc& imageDesc = imageDescs[image];
```

Source: `thirdparty/libktx/lib/basis_transcode.cpp` — ktxTexture2_transcodeEtc1s

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

KTX2 Texture, KTX2 Private Texture State

Evidence:
- Code: `thirdparty/libktx/lib/basis_transcode.cpp` — ktxTexture2_transcodeEtc1s
- Code: `thirdparty/libktx/lib/basis_sgd.h:24` — __cplusplus
- External (primary, unverified (source anchor not found)): [C++ Working Draft N4861](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/n4861.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI bridging](#topic-c-abi-bridging)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [CMake native source indexing](#topic-cmake-native-source-index)
- [C++ enumerated export state](#topic-cplusplus-enumerated-export-state)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ polymorphic platform adapters](#topic-cplusplus-polymorphic-platform-adapters)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ binary data codecs](#topic-cpp-binary-data-codecs)
- [C++ byte-exact binary parsing](#topic-cpp-byte-exact-binary-parsing)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)
- [C++ mapping of C API types](#topic-cpp-c-type-mapping)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ const-reference accessors](#topic-cpp-const-reference-accessors)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ exception containment at ABI boundaries](#topic-cpp-exception-abi-boundaries)
- [C++ frame-time arithmetic](#topic-cpp-frame-scheduling)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ engine binding macros](#topic-cpp-godot-binding-macros)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ JNI Variant marshalling](#topic-cpp-jni-variant-marshalling)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ plugin specialization](#topic-cpp-plugin-specialization)
- [C++ polymorphic backend classes](#topic-cpp-polymorphic-backends)
- [C++ polymorphism and ownership](#topic-cpp-polymorphic-ownership)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ resources and polymorphic engine objects](#topic-cpp-resource-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ inheritance and reference-counted adapters](#topic-cpp-runtime-adapters)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ scope-bound locking](#topic-cpp-scope-locking)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [C++ single-header implementation selection](#topic-cpp-single-header-implementation)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ tagged storage and casts](#topic-cpp-tagged-storage)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and generic containers](#topic-cpp-templates-and-generic-containers)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ typed API records](#topic-cpp-typed-api-records)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [C++ atomic synchronization](#topic-cxx-atomics)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ lambdas and callback objects](#topic-cxx-closures)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [JavaScript browser FFI](#topic-javascript-browser-ffi)
- [Objective-C++ Apple platform adapters](#topic-objective-cpp-apple-adapters)
- [Objective-C++ iOS adapters](#topic-objective-cpp-ios-adapters)
- [Python build source generation](#topic-python-build-source-generation)

<a id="topic-cpp-c-linkage-adapters"></a>
## 117. C++ C-linkage adapters

**Scope:** Vendored: graphite

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-and-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Graphite exposes C-callable functions that delegate to C++ objects such as CharInfo, Face, Segment, and Slot.

C++ classes are wrapped behind C linkage functions and derived opaque API types, so C callers use handles rather than C++ class declarations.

### How the implementation applies it

#### CharInfo C entry points

The gr_cinfo_* functions have C linkage, validate the handle with assert, and delegate to CharInfo methods.

#### Opaque derived handles

Types such as gr_face, gr_font, gr_segment, and gr_slot derive from their C++ implementation types and are the types accepted by the C API.

### Walk through a small source example

```cpp
extern "C"
{

unsigned int gr_cinfo_unicode_char(const gr_char_info* p/*not NULL*/)
{
    assert(p);
    return p->unicodeChar();
}
```

Source: `thirdparty/graphite/src/gr_char_info.cpp` — gr_cinfo_unicode_char

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Face, Font, Segment, Slot

Evidence:
- Code: `thirdparty/graphite/src/gr_char_info.cpp` — extern "C" and gr_cinfo_unicode_char
- Code: `thirdparty/graphite/src/inc/CharInfo.h:43` — gr_char_info
- Code: `thirdparty/graphite/src/inc/Face.h:202` — gr_face
- External (primary, unverified (source anchor not found)): [C++ working draft: Linkage specifications](https://eel.is/c++draft/dcl.link), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ linkage and opaque API handles](#topic-cxx-c-abi)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)

<a id="topic-cxx-wide-simd"></a>
## 118. C++ SIMD wrapper specialization

**Scope:** Vendored: embree

**Builds on:** [C++ templates and specialization](#topic-cxx-templates).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The same vector-oriented algorithms are instantiated against widths such as four, eight, and sixteen lanes, with platform-specific definitions behind the wrapper names.

C++ templates and wrapper types represent fixed-width integer, unsigned-integer, float, mask, and long-vector lanes while ISA-specific headers implement their operations.

### How the implementation applies it

#### ISA-specific vector classes

vint<4> is implemented in an SSE2 header, vint<8> in AVX and AVX2 headers, and vuint<16> in an AVX-512 header.

#### Lane algorithms

The integer wrappers compose shuffle, min or max, and select operations to sort or rearrange lane values.

#### Traversal use

BVH traversal uses vfloat<N>, vint<N>, and vbool<N> to calculate slab intersections and masks for multiple children or rays.

### Walk through a small source example

```c
template<>
  struct vint<4>
  {
    ALIGNED_STRUCT_(16);
    
    typedef vboolf4 Bool;
    typedef vint4   Int;
    typedef vfloat4 Float;

    enum  { size = 4 };             // number of SIMD elements
    union { __m128i v; int i[4]; }; // data

    ////////////////////////////////////////////////////////////////////////////////
    /// Constructors, Assignment & Cast Operators
```

Source: `thirdparty/embree/common/simd/vint4_sse2.h` — vint<4> (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

vint4, vint8, vuint16

Evidence:
- Code: `thirdparty/embree/common/simd/vint4_sse2.h:20` — vint<4>
- Code: `thirdparty/embree/common/simd/vint8_avx2.h:18` — vint<8>
- Code: `thirdparty/embree/common/simd/vuint16_avx512.h:18` — vuint<16>
- Code: `thirdparty/embree/kernels/bvh/node_intersector1.h:1294` — BVHNNodeIntersector1
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [SIMD wrappers and alignment](#topic-cxx-simd-alignment)

<a id="topic-cpp-simd-intrinsics"></a>
## 119. C++ SIMD wrappers and intrinsics

**Scope:** Vendored: cvtt

**Builds on:** [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Embree wraps x86 and ARM SIMD operations, while CVTT selects SSE2 support at preprocessing time.

C++ vector wrappers invoke architecture intrinsics to operate on multiple lane values.

### How the implementation applies it

#### Vector reciprocal operation

Embree's vfloat4 wrapper invokes _mm_rcp_ps for an SSE reciprocal estimate.

#### Platform-specific vector definitions

Embree chooses SIMD vector specializations through varying.h and architecture-specific headers.

#### SSE2 capability selection

CVTT defines CVTT_USE_SSE2 when compiler or target macros indicate SSE2 availability.

### Walk through a small source example

```c
const vfloat4 r = _mm_rcp_ps(a);
```

Source: `thirdparty/embree/common/simd/vfloat4_sse2.h` — embree::rcp

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/embree/common/simd/vfloat4_sse2.h` — embree::rcp
- Code: `thirdparty/embree/common/simd/varying.h` — embree::vtypes
- Code: `thirdparty/cvtt/ConvectionKernels_Config.h:6` — CVTT_USE_SSE2
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16
- External (official, unverified (source anchor not found)): [Intel Intrinsics Guide](https://www.intel.com/content/www/us/en/docs/intrinsics-guide/index.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)

<a id="topic-cpp-byte-exact-binary-parsing"></a>
## 120. C++ byte-exact binary parsing

**Scope:** Vendored: basis_universal

**Builds on:** [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The texture containers deliberately operate on byte-addressed buffers and packed integer fields.

This C++ code reads byte streams through packed fields, pointer casts, and bit readers; this is necessary to map serialized texture headers to in-memory values.

### How the implementation applies it

#### Header reinterpretation

Basis transcoding casts an input byte pointer to basis_file_header before obtaining offsets and slice descriptors.

#### Packed KTX2 fields

ktx2_header and ktx2_level_index use packed_uint fields for on-disk values, while pragma packing surrounds the KTX2 layout declarations.

#### Bitwise decoder cursor

bitwise_decoder maintains start, current, and end byte pointers while providing bit-reading operations.

### Walk through a small source example

```cpp
const basis_file_header* pHeader = reinterpret_cast<const basis_file_header*>(pData);
```

Source: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_transcoder::get_file_info

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BasisFileHeader, BasisSliceDescriptor, Ktx2Header, Ktx2LevelIndex

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp:12204` — basisu_transcoder::get_file_info
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:721` — ktx2_header
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder_internal.h:148` — bitwise_decoder
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-cpp-class-templates"></a>
## 121. C++ class templates and specialization

**Scope:** Vendored: embree

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Embree passes the lane-width parameter through vector types such as Vec3vf<K> and specializes subgrid intersectors for width four.

C++ class templates parameterize Embree vector and intersector code by lane width, and a fixed-width specialization selects a width-specific implementation; C++ classes are needed to explain class templates.

### How the implementation applies it

#### Lane-parameterized vector arguments

The packet intersector accepts Vec3vf<K> ray vectors, binding the vector width to template parameter K.

#### Fixed-width subgrid path

The SubGridQuadMIntersector1MoellerTrumbore<4,filter> declaration provides a width-four specialization path.

### Walk through a small source example

```c
struct SubGridQuadMIntersector1MoellerTrumbore<4,filter>
```

Source: `thirdparty/embree/kernels/geometry/subgrid_intersector_moeller.h` — SubGridQuadMIntersector1MoellerTrumbore<4,filter>

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMi, SubGrid

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h:247` — MoellerTrumboreIntersectorK
- Code: `thirdparty/embree/kernels/geometry/subgrid_intersector_moeller.h:102` — SubGridQuadMIntersector1MoellerTrumbore<4,filter>
- External (primary, verified): [C++ Working Draft — Templates](https://eel.is/c++draft/temp), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ variadic binding](#topic-cpp-variadic-binding)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-and-inheritance)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)

<a id="topic-cpp-resource-lifetime"></a>
## 122. C++ constructor and destructor resource lifetime

**Scope:** Vendored: graphite

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-and-inheritance).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The Font constructor establishes borrowed references and per-glyph cache storage; its destructor frees that storage.

A C++ class constructor and destructor control the Font's allocation and release of its cached advance array.

### How the implementation applies it

#### Font::Font initialization

The constructor stores the face reference, derives scale from units per em, selects font-operation callbacks, allocates advances, and fills them with an invalid sentinel.

#### Font::~Font cleanup

The destructor releases the dynamically allocated advances array.

### Walk through a small source example

```cpp
Font::Font(float ppm, const Face & f, const void * appFontHandle, const gr_font_ops * ops)
: m_appFontHandle(appFontHandle ? appFontHandle : this),
  m_face(f),
  m_scale(ppm / f.glyphs().unitsPerEm()),
  m_hinted(appFontHandle && ops && (ops->glyph_advance_x || ops->glyph_advance_y))
{
```

Source: `thirdparty/graphite/src/Font.cpp` — Font::Font

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Font

Evidence:
- Code: `thirdparty/graphite/src/Font.cpp:10` — Font::Font
- Code: `thirdparty/graphite/src/Font.cpp:32` — Font::~Font
- External (primary, unverified (source anchor not found)): [C++ working draft: Constructors](https://eel.is/c++draft/class.ctor), accessed 2026-07-15
- External (primary, unverified (source anchor not found)): [C++ working draft: Destructors](https://eel.is/c++draft/class.dtor), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [RAII, non-copyable resources, and intrusive references](#topic-cxx-raii-reference-ownership)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [GDScript typed collections and signature compatibility](#topic-gdscript-typed-collections)
- [Typed arrays and dictionaries](#topic-typed-collections)

<a id="topic-cpp-flag-stringification"></a>
## 123. C++ flag stringification

**Scope:** Vendored: vulkan

**Builds on:** [C++ typed API records](#topic-cpp-typed-api-records).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

It is a diagnostic representation helper for the generated C++ Vulkan binding.

The helper converts typed `FormatFeatureFlags` values into a string by testing individual flag bits and appending their names.

### How the implementation applies it

#### Bitwise flag test

Each conditional tests whether a named `FormatFeatureFlagBits` enumerator is present in the supplied flag value.

#### Incremental textual output

The function initializes a result string and appends a name for each matching flag.

#### Library feature selection

The header checks `__cpp_lib_format` and includes either `<format>` or `<sstream>`.

### Walk through a small source example

```cpp
if ( value & FormatFeatureFlagBits::eSampledImage )
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_to_string.hpp` — to_string( FormatFeatureFlags value )

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_to_string.hpp:41` — to_string( FormatFeatureFlags value )
- Code: `thirdparty/vulkan/include/vulkan/vulkan_to_string.hpp:25` — __cpp_lib_format
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

<a id="topic-cxx-c-abi"></a>
## 124. C++ linkage and opaque API handles

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Implementation classes remain internal while API calls exchange opaque handle types and plain structures such as RTCRayHit.

C++ exposes RTCDevice, RTCScene, RTCGeometry, and RTCBuffer as opaque pointer handles through a C-linkage public boundary.

### How the implementation applies it

#### Opaque handles

The public headers typedef incomplete struct pointers for device, scene, geometry, and buffer handles.

#### Linkage configuration

rtcore_config.h defines API linkage macros that choose extern C or extern C++ depending on API namespace configuration.

#### Plain query structures

rtcore_ray.h exposes public ray and hit structs independently of internal C++ RayK and HitK representations.

### Walk through a small source example

```c
typedef struct RTCDeviceTy* RTCDevice;
typedef struct RTCSceneTy* RTCScene;
```

Source: `thirdparty/embree/include/embree4/rtcore_device.h` — RTCDevice

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCDevice, RTCScene, RTCGeometry, RTCBuffer

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h:11` — RTCDevice
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h:36` — RTCBuffer
- Code: `thirdparty/embree/include/embree4/rtcore_config.h:38` — RTC_API_EXTERN_C
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C structures and pointer handles](#topic-c-structures-and-pointers)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C ABI bridging](#topic-c-abi-bridging)
- [C ABI structures and opaque state](#topic-c-abi-data-structures)
- [C-compatible guarded headers](#topic-c-compatible-header-guards)
- [C structs and pointer-linked state](#topic-c-structs-pointers)
- [C++ C-linkage adapters](#topic-cpp-c-linkage-adapters)

<a id="topic-cpp-c-type-mapping"></a>
## 125. C++ mapping of C API types

**Scope:** Vendored: vulkan

**Builds on:** [C++ typed API records](#topic-cpp-typed-api-records).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The generated mapping provides a type-level bridge between the C Vulkan ABI names and the C++ API layer.

`CppType<Vk...>` specializations associate C Vulkan type names with C++ wrapper types.

### How the implementation applies it

#### Wrapper-type association

A `CppType` specialization is emitted for each mapped Vulkan C record type.

#### Structure-type association

Many record mappings are accompanied by a `CppType<StructureType, ...>` specialization keyed by the Vulkan structure-type enumerator.

### Concrete structures to recognize

GraphicsPipelineCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:57235` — CppType<VkGraphicsPipelineCreateInfo>
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:57242` — CppType<StructureType, StructureType::eGraphicsPipelineCreateInfo>
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)

<a id="topic-cpp-overloading"></a>
## 126. C++ member-function overloading

**Scope:** Vendored: icu4c

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The overload set preserves one conceptual operation while accepting multiple input representations.

Overloaded member functions provide related operations for different argument forms, such as LocaleMatcher matching a single locale, an iterator, or a list string.

### How the implementation applies it

#### LocaleMatcher::getBestMatch

The class declares variants that take a Locale and a Locale iterator.

#### LocaleMatcher::getBestMatchForListString

A separate member accepts a StringPiece locale list.

### Walk through a small source example

```cpp
const Locale *LocaleMatcher::getBestMatch(const Locale &desiredLocale, UErrorCode &errorCode) const {
```

Source: `thirdparty/icu4c/common/localematcher.cpp` — LocaleMatcher::getBestMatch

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

LocaleMatcher, Locale

Evidence:
- Code: `thirdparty/icu4c/common/localematcher.cpp:608` — LocaleMatcher::getBestMatch
- Code: `thirdparty/icu4c/common/unicode/localematcher.h:28` — LocaleMatcher
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Iteration protocols](#topic-iteration-protocols)

<a id="topic-cpp-numeric-value-operations"></a>
## 127. C++ numeric value operations

**Scope:** Vendored: gamepadmotionhelpers

**Builds on:** [C++ class state and composition](#topic-cpp-class-state-and-composition).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The motion helper uses float-valued vectors and quaternions together with standard-library clamps and maxima.

C++ member-based value types calculate vector lengths, normalized directions, quaternion orientation, and time-scaled motion updates.

### How the implementation applies it

#### Angular integration inputs

Motion update constructs an axis vector from gyroscope samples, computes angular speed, and scales the angle by delta time.

#### Gravity correction

The implementation normalizes acceleration and gravity vectors, bounds dot products, and uses correction speeds from settings.

### Walk through a small source example

```cpp
struct Quat
	{
		float w;
		float x;
		float y;
		float z;

		Quat();
		Quat(float inW, float inX, float inY, float inZ);
		void Set(float inW, float inX, float inY, float inZ);
		Quat& operator*=(const Quat& rhs);
		friend Quat operator*(Quat lhs, const Quat& rhs);
		void Normalize();
```

Source: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — Quat (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp:27` — Quat
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp:45` — Vec
- External (primary, unverified (source anchor not found)): [N4950: Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ SIMD wrappers and intrinsics](#topic-cpp-simd-intrinsics)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [SIMD wrappers and alignment](#topic-cxx-simd-alignment)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)

<a id="topic-cpp-polymorphic-ownership"></a>
## 128. C++ polymorphism and ownership

**Scope:** Vendored: openxr

**Builds on:** [C++ inline native wrappers](#topic-cpp-native-wrappers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The OpenXR loader exposes abstract recorder types while factory functions return owning smart pointers.

C++ classes use public inheritance and `std::unique_ptr` to retain polymorphic loader services and dispatch state.

### How the implementation applies it

#### LoaderLogRecorder

The recorder hierarchy supplies an interface for alternative logging destinations.

#### MakeStdErrLoaderLogRecorder

The factory returns `std::unique_ptr<LoaderLogRecorder>`, making ownership transfer explicit in the function type.

#### LoaderInstance::DispatchTable

The loader instance stores its generated dispatch table behind a unique pointer.

### Walk through a small source example

```cpp
std::unique_ptr<LoaderLogRecorder> MakeStdErrLoaderLogRecorder(void* user_data);
```

Source: `thirdparty/openxr/src/loader/loader_logger_recorders.hpp` — MakeStdErrLoaderLogRecorder

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

LoaderInstance, XrGeneratedDispatchTableCore

Evidence:
- Code: `thirdparty/openxr/src/loader/loader_logger.hpp:64` — LoaderLogRecorder
- Code: `thirdparty/openxr/src/loader/loader_instance.hpp` — LoaderInstance::DispatchTable
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)

<a id="topic-cpp-runtime-symbol-loading"></a>
## 129. C++ runtime symbol loading

**Scope:** Vendored: metal-cpp

**Builds on:** [C++ inline native wrappers](#topic-cpp-native-wrappers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Metal, MetalFX, and QuartzCore private headers use this pattern to bind symbols at runtime.

C++ wrapper types resolve framework symbols dynamically through `dlsym`-based helper templates.

### How the implementation applies it

#### MTL::Private::LoadSymbol

The helper casts a `dlsym` result to the requested pointer type.

#### MTLFX::Private::LoadSymbol

The MetalFX private header repeats the dynamic symbol-loading mechanism for its namespace.

### Walk through a small source example

```cpp
const _Type* pAddress = static_cast<_Type*>(dlsym(RTLD_DEFAULT, pSymbol));
```

Source: `thirdparty/metal-cpp/Metal/MTLPrivate.hpp` — MTL::Private::LoadSymbol

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MTL4::Archive

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTLPrivate.hpp:86` — MTL::Private::LoadSymbol
- Code: `thirdparty/metal-cpp/MetalFX/MTLFXPrivate.hpp` — MTLFX::Private::LoadSymbol
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)

<a id="topic-cpp-static-abi-contracts"></a>
## 130. C++ static ABI contracts

**Scope:** Vendored: vulkan

**Builds on:** [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The static-assertion header checks wrapper size compatibility with C Vulkan handles and selected type traits.

The binding uses a macro-selected assertion facility to verify wrapper ABI layout and copy or move properties at compile time.

### How the implementation applies it

#### Wrapper-size contract

The `Instance` wrapper must have the same size as `VkInstance`.

#### Copy and move contracts

The header checks that `Instance` is copy constructible and nothrow move constructible.

#### Macro-selected assertion facility

`VULKAN_HPP_STATIC_ASSERT` defaults to `static_assert` unless callers configure it differently.

### Walk through a small source example

```cpp
VULKAN_HPP_STATIC_ASSERT( sizeof( VULKAN_HPP_NAMESPACE::Instance ) == sizeof( VkInstance ), "handle and wrapper have different size!" );
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_static_assertions.hpp` — VK_VERSION_1_0 Instance size assertion

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_static_assertions.hpp` — VK_VERSION_1_0 Instance size, copy-construction, and move-construction assertions
- Code: `thirdparty/vulkan/include/vulkan/vulkan_hpp_macros.hpp:82` — VULKAN_HPP_STATIC_ASSERT
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/n4861.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI-versioned initialization](#topic-c-abi-versioned-initialization)
- [C macro-based binary decoding](#topic-c-macro-codecs)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [C++ inline native wrappers](#topic-cpp-native-wrappers)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)

<a id="topic-cpp-templates-and-generic-containers"></a>
## 131. C++ templates and generic containers

**Scope:** Vendored: basis_universal

**Builds on:** [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Basis containers, Clipper points and paths, and Embree math types use type parameters rather than duplicate implementations per element type.

This C++ code parameterizes collection and math utilities by type, so their typed operations can be reused.

### How the implementation applies it

#### Basis vector container

basisu::vector is a templated collection type that derives relational operations from its element-specialized form.

#### Generic geometry points

Clipper's Point<T> supports coordinate representations selected by the template parameter.

#### Generic SIMD-adjacent math

Embree's Vec3<T> and BBox<T> operate on a supplied scalar or vector-like element type.

### Walk through a small source example

```c
class vector : public rel_ops< vector<T> >
```

Source: `thirdparty/basis_universal/transcoder/basisu_containers.h` — basisu::vector

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Ktx2LevelIndex, PolyPathD

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_containers.h:3296` — basisu::vector
- Code: `thirdparty/clipper2/include/clipper2/clipper.core.h` — Clipper2Lib::Point
- Code: `thirdparty/embree/common/math/vec3.h` — embree::Vec3
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [Python SCons build hooks](#topic-python-scons-build-hooks)

<a id="topic-cpp-virtual-dispatch"></a>
## 132. C++ virtual dispatch

**Scope:** Vendored: icu4c

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is used where ICU needs an interface with multiple concrete implementations.

Virtual dispatch uses derived classes to substitute behavior, as locale iterators override next and ICU break engines derive from a common engine type.

### How the implementation applies it

#### Locale::RangeIterator::next

The iterator provides its concrete next operation with the override specifier.

#### LanguageBreakEngine

The base engine type supports specialized dictionary, unhandled, wrapper, and LSTM break-engine classes.

### Walk through a small source example

```c
const Locale &next() override { return *it_++; }
```

Source: `thirdparty/icu4c/common/unicode/locid.h` — Locale::RangeIterator::next

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

LocaleMatcher, RuleBasedBreakIterator

Evidence:
- Code: `thirdparty/icu4c/common/unicode/locid.h` — Locale::RangeIterator::next
- Code: `thirdparty/icu4c/common/brkeng.h:28` — LanguageBreakEngine
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ polymorphic backend classes](#topic-cpp-polymorphic-backends)

<a id="topic-cpp-visitor-traversal"></a>
## 133. C++ visitor-style intermediate-tree traversal

**Scope:** Vendored: glslang

**Builds on:** [C++ classes and inheritance](#topic-cpp-classes-and-inheritance).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

glslang uses specialized TIntermTraverser subclasses for reflection, textual output, live-variable gathering, link validation, and limit checking.

C++ inheritance provides traversal specializations that analyze intermediate shader-tree nodes.

### How the implementation applies it

#### Reflection traversal

TReflectionTraverser is a TIntermTraverser subclass used by reflection.cpp to inspect intermediate declarations and types.

#### Other traversal passes

TOutputTraverser, TVarGatherTraverser, TIOTraverser, and TInductiveTraverser demonstrate separate intermediate-tree analyses and outputs.

### Walk through a small source example

```cpp
class TReflectionTraverser : public TIntermTraverser {
```

Source: `thirdparty/glslang/glslang/MachineIndependent/reflection.cpp` — TReflectionTraverser

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TObjectReflection

Evidence:
- Code: `thirdparty/glslang/glslang/MachineIndependent/reflection.cpp:78` — TReflectionTraverser
- Code: `thirdparty/glslang/glslang/MachineIndependent/intermOut.cpp:65` — TOutputTraverser
- Code: `thirdparty/glslang/glslang/MachineIndependent/iomapper.cpp:166` — TVarGatherTraverser
- External (primary, unverified (source anchor not found)): [C++ working draft: Classes](https://eel.is/c++draft/class), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ polymorphic backend classes](#topic-cpp-polymorphic-backends)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)
- [Godot shader source composition](#topic-godot-shader-language)

<a id="topic-c-concurrent-state"></a>
## 134. C-managed concurrent state

**Scope:** Vendored: sdl

**Builds on:** [C structs and pointer-linked state](#topic-c-structs-pointers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SDL wraps synchronization behind its own mutex and atomic APIs while its C code explicitly maintains the protected data.

C structs hold mutexes, atomic counts, and linked entries when SDL manages concurrent event and device state.

### How the implementation applies it

#### SDL_EventQ in SDL_events.c

The queue combines SDL_Mutex, SDL_AtomicInt, and linked queue pointers.

#### SDL_StartEventLoop in SDL_events.c

Event-loop startup creates and locks the queue mutex before activating the queue and event-watch list.

### Walk through a small source example

```c
SDL_EventEntry *free;
} SDL_EventQ = { NULL, false, { 0 }, 0, NULL, NULL, NULL };


static void SDL_CleanupTemporaryMemory(void *data)
{
    SDL_TemporaryMemoryState *state = (SDL_TemporaryMemoryState *)data;

    SDL_FreeTemporaryMemory();
    SDL_free(state);
}

static SDL_TemporaryMemoryState *SDL_GetTemporaryMemoryState(bool create)
{
```

Source: `thirdparty/sdl/events/SDL_events.c` — SDL_EventQ (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_EventEntry

Evidence:
- Code: `thirdparty/sdl/events/SDL_events.c:144` — SDL_EventQ
- Code: `thirdparty/sdl/events/SDL_events.c:946` — SDL_StartEventLoop
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ atomic synchronization](#topic-cxx-atomics)

<a id="topic-godot-shader-language"></a>
## 135. Godot shader source composition

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The repository distinguishes high-level .gdshader source from RenderingDevice-native GLSL and SPIR-V resources.

Godot's shader language writes Shader resource source and can compose reusable ShaderInclude fragments with a preprocessor include directive.

### How the implementation applies it

#### Shader source resource

Shader is the Resource class for a custom shader program implemented in the Godot shading language.

#### Include composition

ShaderInclude stores source text intended for a Shader preprocessor include.

#### RenderingDevice stage sources

RDShaderSource stores native stage source text for RenderingDevice instead of the high-level Shader resource.

### Walk through a small source example

```text
#include "res://shader_lib.gdshaderinc"
```

Source: `doc/classes/ShaderInclude.xml` — ShaderInclude description

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ShaderMaterial, RDPipelineSpecializationConstant

Evidence:
- Code: `doc/classes/Shader.xml:2` — Shader
- Code: `doc/classes/ShaderInclude.xml:2` — ShaderInclude
- Code: `doc/classes/RDShaderSource.xml` — RDShaderSource.set_stage_source
- External (official, unverified (source anchor not found)): [Shading language](https://docs.godotengine.org/en/stable/tutorials/shaders/shader_reference/shading_language.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor composition](#topic-c-preprocessor-composition)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [C preprocessor platform selection](#topic-c-preprocessor-platform-selection)
- [C preprocessor portability](#topic-c-preprocessor-portability)
- [C++ preprocessor configuration](#topic-cpp-preprocessor-configuration)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)
- [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration)
- [C++ const references and pointers](#topic-cpp-const-borrowing)
- [C++ visitor-style intermediate-tree traversal](#topic-cpp-visitor-traversal)
- [GLSL compute shaders](#topic-glsl-compute-shaders)

<a id="topic-cxx-reflection-macros"></a>
## 136. Macro-based RTTI registration

**Scope:** Vendored: jolt_physics

**Builds on:** [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model).

**This prepares you for:** [Stream-oriented binary serialization](#topic-cxx-stream-serialization).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The JPH_IMPLEMENT_SERIALIZABLE and JPH_ADD_ATTRIBUTE macro families are the repository's reflection convention.

Macro-based RTTI registration attaches class metadata and serializable member attributes to types that participate in dynamic lookup and object streaming.

### How the implementation applies it

#### Member attribute registration

LinearCurve registers its Point coordinates and points array through serializable attributes.

#### Runtime type lookup

Factory finds RTTI by name or hash and inspects registered primitive member types.

### Walk through a small source example

```cpp
JPH_IMPLEMENT_SERIALIZABLE_NON_VIRTUAL(LinearCurve::Point)
{
	JPH_ADD_ATTRIBUTE(Point, mX)
	JPH_ADD_ATTRIBUTE(Point, mY)
}
```

Source: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp` — LinearCurve::Point RTTI registration

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ConstraintSettings, PhysicsMaterialSimple

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp:14` — JPH_IMPLEMENT_SERIALIZABLE_NON_VIRTUAL
- Code: `thirdparty/jolt_physics/Jolt/Core/Factory.cpp:19` — Factory::Find
- Code: `thirdparty/jolt_physics/Jolt/Core/RTTI.h:14` — RTTI
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C macro-based binary decoding](#topic-c-macro-codecs)
- [C preprocessor platform and precision selection](#topic-c-platform-selection)
- [C preprocessor feature and platform configuration](#topic-c-preprocessor-configuration)
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C# attributes and reflection](#topic-csharp-attributes-reflection)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ preprocessor configuration](#topic-cxx-conditional-compilation)

<a id="topic-cxx-raii-reference-ownership"></a>
## 137. RAII, non-copyable resources, and intrusive references

**Scope:** Vendored: jolt_physics

**Builds on:** [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The library explicitly marks many resources NonCopyable and wraps longer-lived objects in intrusive reference types.

RAII-bound C++ objects manage locks and scopes, while Ref and RefTarget keep shared simulation objects alive without ordinary copying.

### How the implementation applies it

#### Mutex-scoped recording

DebugRendererRecorder creates a lock_guard before modifying its current frame or serialized command stream.

#### Job lifetime ownership

JobHandle privately derives from Ref<Job>, so a handle retains its Job while dependency and queue operations proceed.

### Walk through a small source example

```cpp
lock_guard lock(mMutex);
```

Source: `thirdparty/jolt_physics/Jolt/Renderer/DebugRendererRecorder.cpp` — DebugRendererRecorder::DrawLine

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shape, PhysicsScene

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/Reference.h` — RefTarget, Ref, RefConst
- Code: `thirdparty/jolt_physics/Jolt/Core/JobSystem.h:79` — JobHandle : private Ref<Job>
- Code: `thirdparty/jolt_physics/Jolt/Renderer/DebugRendererRecorder.cpp:13` — DebugRendererRecorder::DrawLine
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ inheritance, virtual interfaces, and Ref ownership](#topic-cpp-engine-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ constructor and destructor resource lifetime](#topic-cpp-resource-lifetime)
- [C++ scoped lock guards](#topic-cpp-scoped-locks)

<a id="topic-cxx-simd-alignment"></a>
## 138. SIMD wrappers and alignment

**Scope:** Vendored: jolt_physics

**Builds on:** [Preprocessor-selected platform configuration](#topic-cxx-preprocessor-configuration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Public math types hide platform intrinsics behind Vec3, Vec4, UVec4, matrices, and box helpers.

SIMD vector types use aligned storage and compile-time CPU branches to implement vector, matrix, bounding-box, and geometry operations across SSE, NEON, and RISC-V vector paths.

### How the implementation applies it

#### RISC-V vector gathers

Vec4 implementations scale per-lane offsets and gather four floating-point values through RVV intrinsics.

#### Four-box overlap testing

AABox4 splats a single box into vector lanes and computes axis separation tests against four boxes at once.

### Walk through a small source example

```text
const vfloat32m1_t gathered = __riscv_vluxei32_v_f32m1(inBase, scaled_offsets, 4);
```

Source: `thirdparty/jolt_physics/Jolt/Math/Vec4.inl` — RVV gather implementation

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Math/Vec4.inl` — RVV gather implementation
- Code: `thirdparty/jolt_physics/Jolt/Geometry/AABox4.h:13` — AABox4VsBox
- Code: `thirdparty/jolt_physics/Jolt/Math/Vec3.h:16` — alignas(JPH_VECTOR_ALIGNMENT) Vec3
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C++ copy-on-write container storage](#topic-cpp-copy-on-write-containers)
- [C++ numeric value operations](#topic-cpp-numeric-value-operations)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)

<a id="topic-cpp-specialized-marshalling"></a>
## 139. C++ specialized argument marshalling

**Scope:** First-party

**Builds on:** [C++ templates and traits](#topic-cpp-templates-traits).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The specializations select direct, converting, reference, vector, pointer, object, reference-counted, and typed-container paths.

The binding layer specializes PtrToArg through C++ templates to marshal bound method types and return values.

### How the implementation applies it

#### Scalar conversion specialization

The boolean specialization converts through uint8_t rather than passing a bool storage pointer directly.

#### Geometry argument choices

Method pointer conversion chooses direct or by-reference transport for individual geometry types.

#### Typed container support

PtrToArg has dedicated typed-array and typed-dictionary specializations.

### Walk through a small source example

```c
struct PtrToArg<bool> : Internal::PtrToArgConvert<bool, uint8_t> {};
```

Source: `core/variant/method_ptrcall.h` — PtrToArg<bool>

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant, PropertyInfo

Evidence:
- Code: `core/variant/method_ptrcall.h:148` — PtrToArg<bool>
- Code: `core/variant/method_ptrcall.h:183` — PtrToArg<Vector3>
- Code: `core/variant/type_info.h:59` — GetTypeInfo
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ template-based method binding](#topic-cpp-template-binding)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)

<a id="topic-cpp-template-binding"></a>
## 140. C++ template-based method binding

**Scope:** First-party

**Builds on:** [C++ native class hierarchy](#topic-cpp-native-classes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This keeps native method registration type-aware while producing the reflective MethodBind representation used by script-facing APIs.

ClassDB uses C++ templates, `if constexpr`, and member-function traits to generate binders for native methods.

### How the implementation applies it

#### Typed member binding

ClassDB::bind_method creates a MethodBind from a member-function pointer and detects raw Object pointer return types with if constexpr.

#### Variadic argument metadata

D_METHOD builds debug argument-name metadata from a variadic template parameter pack.

#### Build-gated registration

GDREGISTER_CLASS conditionally instantiates class registration through GD_IS_CLASS_ENABLED.

### Walk through a small source example

```c
#define GDREGISTER_CLASS(m_class) \
	if constexpr (GD_IS_CLASS_ENABLED(m_class)) { \
		::ClassDB::register_class<m_class>(); \
	}
```

Source: `core/object/class_db.h` — GDREGISTER_CLASS

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ClassInfo, Variant

Evidence:
- Code: `core/object/class_db.h` — ClassDB::bind_method
- Code: `core/object/class_db.h:79` — D_METHOD
- Code: `core/object/class_db.h:568` — GDREGISTER_CLASS
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ class templates and specialization](#topic-cpp-class-templates)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ templates and typed containers](#topic-cpp-generic-containers)
- [C++ ownership and polymorphic trees](#topic-cpp-ownership-and-polymorphism)
- [C++ runtime symbol loading](#topic-cpp-runtime-symbol-loading)
- [C++ specialized argument marshalling](#topic-cpp-specialized-marshalling)
- [C++ templates](#topic-cpp-templates)
- [C++ templates and const access](#topic-cpp-templates-and-const-access)
- [C++ templates and specialization](#topic-cpp-templates-and-specialization)
- [C++ templates for binary data operations](#topic-cpp-templates-for-binary-data)
- [C++ templates and traits](#topic-cpp-templates-traits)
- [C++ variadic binding](#topic-cpp-variadic-binding)
- [C++ templates and specialization](#topic-cxx-templates)
- [C++ SIMD wrapper specialization](#topic-cxx-wide-simd)
- [C++ templates and non-owning views](#topic-cpp-templates-and-views)

<a id="topic-csharp-unsafe-interop"></a>
## 141. C# unsafe interop and function pointers

**Scope:** First-party

**Builds on:** [C# partial types and source generation](#topic-csharp-source-generation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

GodotSharp validates unmanaged table size before reinterpreting its address, marshals engine values through native layouts, and uses generated callback tables for native-to-managed calls.

Unsafe C# stores pointers and function pointers in ABI structs and invokes generated partial native calls.

### How the implementation applies it

#### Native function table initialization

NativeFuncs.Initialize verifies the unmanaged callback-structure size and stores a pointer-backed UnmanagedCallbacks value.

#### Managed callback ABI

ManagedCallbacks is sequentially laid out and contains unmanaged function pointers for signals, delegates, script management, and marshaling.

#### Variant argument access

NativeVariantPtrArgs is an unsafe ref struct that stores a godot_variant double pointer and exposes indexed references without enabling unsafe blocks in game projects.

### Walk through a small source example

```csharp
[StructLayout(LayoutKind.Sequential)]
    public unsafe struct ManagedCallbacks
    {
        // @formatter:off
        public delegate* unmanaged<IntPtr, godot_variant**, int, godot_bool*, void> SignalAwaiter_SignalCallback;
        public delegate* unmanaged<IntPtr, void*, godot_variant**, int, godot_variant*, void> DelegateUtils_InvokeWithVariantArgs;
        public delegate* unmanaged<IntPtr, IntPtr, godot_bool> DelegateUtils_DelegateEquals;
        public delegate* unmanaged<IntPtr, int> DelegateUtils_DelegateHash;
        public delegate* unmanaged<IntPtr, godot_bool*, int> DelegateUtils_GetArgumentCount;
        public delegate* unmanaged<IntPtr, godot_array*, godot_bool> DelegateUtils_TrySerializeDelegateWithGCHandle;
        public delegate* unmanaged<godot_array*, IntPtr*, godot_bool> DelegateUtils_TryDeserializeDelegateWithGCHandle;
        public delegate* unmanaged<void> ScriptManagerBridge_FrameCallback;
        public delegate* unmanaged<godot_string_name*, IntPtr, IntPtr> ScriptManagerBridge_CreateManagedForGodotObjectBinding;
        public delegate* unmanaged<IntPtr, IntPtr, godot_variant**, int, godot_bool> ScriptManagerBridge_CreateManagedForGodotObjectScriptInstance;
```

Source: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs` — ManagedCallbacks (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ManagedCallbacks, Variant

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/NativeFuncs.cs` — NativeFuncs.Initialize
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs:8` — ManagedCallbacks
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/NativeVariantPtrArgs.cs:8` — NativeVariantPtrArgs
- External (official, unverified (source anchor not found)): [Unsafe code, pointers to data, and function pointers - C# reference](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/unsafe-code), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C ABI-versioned initialization](#topic-c-abi-versioned-initialization)
- [C pointers, arrays, and strides](#topic-c-pointers-arrays-and-strides)
- [C# partial types and source-generator tests](#topic-csharp-partial-source-generation)
- [C aggregate state and callback modules](#topic-c-aggregate-callback-modules)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [C++ pluggable allocation](#topic-cpp-pluggable-allocation)

<a id="topic-cmake-native-source-index"></a>
## 142. CMake native source indexing

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The CMake file explicitly labels itself non-functional and configures it for Android Studio editor support.

CMake source-index configuration exposes the native C/C++ source tree to Android Studio as part of Android Gradle assembly.

### How the implementation applies it

#### CMAKE_CXX_STANDARD

The native-source-index CMake file sets the requested C++ standard to 14.

#### file(GLOB_RECURSE SOURCES)

The file recursively collects C-family source files below the Godot root for the editor-support target.

#### add_executable

The source-index target is declared with the collected source and header lists.

### Walk through a small source example

```text
set(CMAKE_CXX_STANDARD 14)
```

Source: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt` — CMAKE_CXX_STANDARD

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt:5` — CMAKE_CXX_STANDARD
- Code: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt` — file(GLOB_RECURSE SOURCES)
- Code: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt:18` — add_executable
- External (official, unverified (source anchor not found)): [cmake-language(7)](https://cmake.org/cmake/help/latest/manual/cmake-language.7.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-groovy-gradle-build-logic"></a>
## 143. Groovy Gradle build logic

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Android Gradle root, app, library, and editor build files contain executable Groovy configuration.

Groovy Gradle scripts generate Android build tasks and select variants as part of Android Gradle assembly.

### How the implementation applies it

#### generateBuildTasks

The root Android build script declares a function with flavor, edition, and Android-distribution parameters.

#### getSconsTaskName

The same script declares a helper that names SCons tasks from flavor, build type, and ABI inputs.

#### app/config.gradle

The app configuration script begins by creating a Groovy map.

### Walk through a small source example

```text
def generateBuildTasks(String flavor = "template", String edition = "standard", String androidDistro = "android") {
```

Source: `platform/android/java/build.gradle` — generateBuildTasks

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/build.gradle:101` — generateBuildTasks
- Code: `platform/android/java/build.gradle:47` — getSconsTaskName
- Code: `platform/android/java/app/config.gradle:104` — map
- External (official, unverified (source anchor not found)): [The Apache Groovy programming language - Syntax](https://groovy-lang.org/syntax.html), accessed 2026-07-15

<a id="topic-gdscript-query-objects"></a>
## 144. GDScript query-object API use

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The documented API uses object construction and member access rather than an inline positional-argument ray query.

GDScript can configure a physics query object, pass it to a physics space, and receive the collision result through chained method calls.

### How the implementation applies it

#### Ray parameter construction

PhysicsRayQueryParameters2D.create constructs a query from origin and destination vectors.

#### Direct space query

PhysicsDirectSpaceState2D.intersect_ray accepts the parameter object and returns collision data.

#### Shape query configuration

PhysicsShapeQueryParameters2D groups shape, motion, mask, collision-kind, and exclusion settings before a direct-space query.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="PhysicsShapeQueryParameters2D" inherits="RefCounted" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Provides parameters for [PhysicsDirectSpaceState2D]'s methods.
	</brief_description>
	<description>
		By changing various properties of this object, such as the shape, you can configure the parameters for [PhysicsDirectSpaceState2D]'s methods.
	</description>
	<tutorials>
	</tutorials>
	<members>
		<member name="collide_with_areas" type="bool" setter="set_collide_with_areas" getter="is_collide_with_areas_enabled" default="false">
			If [code]true[/code], the query will take [Area2D]s into account.
		</member>
```

Source: `doc/classes/PhysicsShapeQueryParameters2D.xml` — PhysicsShapeQueryParameters2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D

Evidence:
- Code: `doc/classes/PhysicsRayQueryParameters2D.xml` — PhysicsRayQueryParameters2D usage example
- Code: `doc/classes/PhysicsDirectSpaceState2D.xml` — PhysicsDirectSpaceState2D.intersect_ray
- Code: `doc/classes/PhysicsShapeQueryParameters2D.xml:2` — PhysicsShapeQueryParameters2D
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [GDScript Android plugin integration](#topic-gdscript-plugin-integration)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)
- [Types, inference, and conversions](#topic-types-inference-and-conversions)
- [Python SCons build hooks](#topic-python-scons-build-hooks)

<a id="topic-cxx-stream-serialization"></a>
## 145. Stream-oriented binary serialization

**Scope:** Vendored: jolt_physics

**Builds on:** [Macro-based RTTI registration](#topic-cxx-reflection-macros).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Several types pair SaveBinaryState and RestoreBinaryState methods with their RTTI registrations.

StreamIn and StreamOut implement binary persistence using RTTI hashes for dynamically typed objects and explicit field writes for concrete state.

### How the implementation applies it

#### Polymorphic type identity

GroupFilter writes its RTTI hash before subtype-specific state is restored through StreamUtils.

#### Field-wise material persistence

PhysicsMaterialSimple writes and reads its debug name and debug color explicitly.

### Walk through a small source example

```cpp
void GroupFilter::SaveBinaryState(StreamOut &inStream) const
{
	inStream.Write(GetRTTI()->GetHash());
}
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/GroupFilter.cpp` — GroupFilter::SaveBinaryState

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ConstraintSettings, PhysicsMaterialSimple, Skeleton

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/StreamIn.h:12` — StreamIn
- Code: `thirdparty/jolt_physics/Jolt/Core/StreamOut.h:12` — StreamOut
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/GroupFilter.cpp` — GroupFilter::SaveBinaryState and sRestoreFromBinaryState
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

<a id="topic-cpp-frame-scheduling"></a>
## 146. C++ frame-time arithmetic

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

`MainTimerSync` contains the timing policy; `DeltaSmoother` is a separate timing-related class in the same subsystem.

C++ frame timing turns accumulated time and a physics tick rate into bounded physics-step counts for frame timing.

### How the implementation applies it

#### MainTimerSync::advance

The implementation computes a maximum possible step count with `std::floor`, accumulated time, the physics tick rate, and a jitter-fix value.

#### MainTimerSync

`MainTimerSync` is the class that owns the timing synchronization API.

#### DeltaSmoother

`DeltaSmoother` is declared beside `MainTimerSync`, separating delta smoothing from the timer-sync class.

### Walk through a small source example

```cpp
const int max_possible_steps = std::floor((time_accum)*p_physics_ticks_per_second + get_physics_jitter_fix());
```

Source: `main/main_timer_sync.cpp` — MainTimerSync::advance

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MainFrameTime

Evidence:
- Code: `main/main_timer_sync.h:43` — MainTimerSync
- Code: `main/main_timer_sync.h:44` — DeltaSmoother
- Code: `main/main_timer_sync.cpp:539` — MainTimerSync::advance
- External (primary, unverified (source anchor not found)): [Draft C++ Standard: Contents](https://eel.is/c++draft/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)

<a id="topic-gdscript-declarations"></a>
## 147. GDScript declarations and typing

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [GDScript signals and callbacks](#topic-gdscript-signals-callables).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

GDScript is implemented in a dedicated module that models compiled scripts, member information, analysis, and compilation.

GDScript source uses `extends`, `class_name`, typed variables, functions, and annotations; the engine represents each script as a GDScript Resource.

### How the implementation applies it

#### Script runtime representation

GDScript derives from Script and records validity, reload state, and member metadata in the native implementation.

#### Typed collection syntax

The class-reference examples use a typed Dictionary declaration with String keys and int values.

#### Parser implementation

The GDScript module contains a dedicated GDScriptParser implementation for source parsing.

### Walk through a small source example

```text
var typed_dict: Dictionary[String, int] = {
```

Source: `doc/classes/Dictionary.xml` — Dictionary typed dictionary code block

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDScript, Variant

Evidence:
- Code: `modules/gdscript/gdscript.h:57` — GDScript
- Code: `modules/gdscript/gdscript_parser.cpp:449` — GDScriptParser::parse
- Code: `doc/classes/Dictionary.xml` — Dictionary typed dictionary code block
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/latest/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [GDScript Android plugin integration](#topic-gdscript-plugin-integration)
- [GDScript query-object API use](#topic-gdscript-query-objects)
- [Types, inference, and conversions](#topic-types-inference-and-conversions)

<a id="topic-nodepaths-and-indexed-access"></a>
## 148. Node paths and indexed access

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Completion expectations vary with the source expression and with the node and script information available from a scene.

The fixtures use `$`, `%`, and indexed self access to resolve Node members inside a scene-aware test context.

### How the implementation applies it

#### Indexed self access

A `Node` name is read and written through `self["name"]`.

#### Scene-informed completion

Scene-backed completion cases distinguish native, scripted, uniquely named, broad, and incompatible node references.

### Walk through a small source example

```text
name = "Node"
	print(self["name"])
	self["name"] = "Changed"
```

Source: `modules/gdscript/tests/scripts/analyzer/features/subscript_self.gd` — test()

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Scene Fixture, Completion Test Configuration

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/subscript_self.gd:4` — test()
- Code: `modules/gdscript/tests/scripts/completion/get_node/literal_scene/dollar_unique.cfg` — [input] and [output]
- Code: `modules/gdscript/tests/scripts/completion/get_node/local_typehint_scene_incompatible/native_local_typehint_scene.cfg` — [output]
- External (official, unverified (source anchor not found)): [Node class reference](https://docs.godotengine.org/en/stable/classes/class_node.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C stateful struct APIs](#topic-c-stateful-struct-apis)

<a id="topic-java-android-host-api"></a>
## 149. Java Android host APIs

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Java classes supply the Android-facing half of rendering, input, plugins, and native-library integration.

Java uses class inheritance and interface implementation to attach GL/Vulkan views, plugin registration, native library calls, and input handling to the Android platform host.

### How the implementation applies it

#### GodotGLRenderView

The GL render view extends `GLSurfaceView` and implements the engine `GodotRenderView` interface.

#### GodotVulkanRenderView

The Vulkan render view is a separate Java class implementing `GodotRenderView` through the Vulkan surface-view path.

#### GodotPluginRegistry

The plugin registry is a Java class in the host library and imports Android package metadata APIs and concurrent collections.

#### GodotLib

The Java native-library class imports rendering, storage, native-bridge, TTS, network, and callable types.

### Walk through a small source example

```java
class GodotGLRenderView extends GLSurfaceView implements GodotRenderView {
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotGLRenderView.java` — GodotGLRenderView

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotGLRenderView.java:2` — GodotGLRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotVulkanRenderView.java:2` — GodotVulkanRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/plugin/GodotPluginRegistry.java:2` — GodotPluginRegistry
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotLib.java:2` — GodotLib
- External (official, verified): [The Java Language Specification, Chapter 8: Classes](https://docs.oracle.com/javase/specs/jls/se17/html/jls-8.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Classes, inheritance, and lookup](#topic-classes-inheritance-and-lookup)
- [C++ types, encapsulation, and inheritance](#topic-cpp-abstraction-and-polymorphism)
- [C++ class inheritance for backend contracts](#topic-cpp-class-inheritance)
- [C++ class state and composition](#topic-cpp-class-state-and-composition)
- [C++ classes and inheritance](#topic-cpp-classes)
- [C++ classes and inheritance](#topic-cpp-classes-inheritance)
- [C++ inheritance interfaces](#topic-cpp-inheritance-interfaces)
- [C++ interface polymorphism](#topic-cpp-interface-polymorphism)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C++ class inheritance](#topic-cpp-object-hierarchies)
- [C++ class hierarchies and reference bases](#topic-cxx-class-hierarchy)
- [C++ classes, inheritance, and polymorphic interfaces](#topic-cxx-object-model)

<a id="topic-kotlin-runtime-coordination"></a>
## 150. Kotlin runtime and service coordination

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Kotlin is used for the main `Godot` host, activities, services, remote fragments, storage handlers, and editor code.

Kotlin coordinates the Android platform host through nullable runtime state, Android host construction, and integer-coded remote-service messages.

### How the implementation applies it

#### Godot

The main Kotlin host is a class with a private constructor that receives an Android `Context`.

#### GodotService

The service source defines constants for initialization, start, stop, destroy, engine error, status updates, and restart requests.

#### RemoteGodotFragment

The remote fragment imports Android binding and Messenger types and refers to the `GodotService` engine-status and engine-error types.

### Walk through a small source example

```kotlin
class Godot private constructor(val context: Context) {
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt` — Godot

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Callable, GodotService.EngineStatus

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt:2` — Godot
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt:83` — MSG_INIT_ENGINE
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/RemoteGodotFragment.kt:2` — RemoteGodotFragment
- External (official, unverified (source anchor not found)): [Classes | Kotlin Documentation](https://kotlinlang.org/docs/classes.html), accessed 2026-07-15

<a id="topic-cxx-closures"></a>
## 151. C++ lambdas and callback objects

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This lets generic builder code operate on caller-selected node and leaf representations without hard-coding a single output format.

C++ lambdas and callable objects bind BVH builders to allocation, node creation, leaf creation, progress, and completion behavior.

### How the implementation applies it

#### Leaf-creation closure

BVHNBuilderVirtual creates a capturing lambda that forwards a primitive range and allocator to createLeaf.

#### Builder callback parameters

GeneralBVHBuilder::BuilderT accepts CreateAllocFunc, CreateNodeFunc, UpdateNodeFunc, CreateLeafFunc, and progress callback types.

#### Progress monitor closure

ProgressMonitorClosure stores a Closure for build-progress reporting.

### Walk through a small source example

```cpp
auto createLeafFunc = [&] (const PrimRef* prims, const range<size_t>& set, const Allocator& alloc) -> NodeRef {
        return createLeaf(prims,set,alloc);
      };
```

Source: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ProgressMonitorClosure, GeneralBVHBuilder, BVHNBuilderVirtual

Evidence:
- Code: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build
- Code: `thirdparty/embree/kernels/builders/bvh_builder_sah.h` — GeneralBVHBuilder::BuilderT
- Code: `thirdparty/embree/kernels/common/builder.h:36` — ProgressMonitorClosure
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C ABI structures and manual lifetime](#topic-c-abi-manual-lifetime)
- [Callables and lambdas](#topic-callables-and-lambdas)
- [C++ Basis transcoding behind a C-compatible API](#topic-cpp-basis-transcoding)
- [C++ runtime class registration](#topic-cpp-class-registration)
- [C++ inheritance and engine reference handles](#topic-cpp-classes-and-ref-handles)
- [C++ native class hierarchy](#topic-cpp-native-classes)
- [C function-pointer API declarations](#topic-c-function-pointer-apis)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [GDScript signals and callbacks](#topic-gdscript-signals-callables)

<a id="topic-gdscript-signals-callables"></a>
## 152. GDScript signals and callbacks

**Scope:** First-party

**Builds on:** [GDScript declarations and typing](#topic-gdscript-declarations).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the idiomatic script-facing bridge to Godot’s signal and input-dispatch systems.

GDScript connects a Signal to a Callable and uses callback functions such as `_input` in Object-derived nodes.

### How the implementation applies it

#### Signal connection syntax

The ConfirmationDialog reference connects the pressed signal to a callback with Signal.connect syntax.

#### Input callback convention

InputEvent documentation directs event handling to Node._input, which GDScript examples implement as a function callback.

#### Callable transport

Object and editor APIs accept Callable values for deferred or externally invoked behavior.

### Walk through a small source example

```text
get_cancel_button().pressed.connect(_on_canceled)
```

Source: `doc/classes/ConfirmationDialog.xml` — ConfirmationDialog description code block

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node, Variant

Evidence:
- Code: `doc/classes/ConfirmationDialog.xml` — ConfirmationDialog description code block
- Code: `doc/classes/InputEvent.xml:2` — InputEvent
- Code: `doc/classes/Object.xml` — Object.bind
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/latest/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [C function-pointer callbacks](#topic-c-function-pointer-callbacks)
- [C function-pointer interfaces](#topic-c-function-pointer-interfaces)
- [Callables and lambdas](#topic-callables-and-lambdas)
- [C++ export callbacks](#topic-cplusplus-export-callbacks)
- [C++ lambdas and callback objects](#topic-cxx-closures)
- [GDScript query-object API use](#topic-gdscript-query-objects)
- [C# partial types and source-generator tests](#topic-csharp-partial-source-generation)
- [Lambdas and standard algorithms](#topic-cxx-lambdas-standard-algorithms)
- [Signals and await](#topic-signals-and-await)

<a id="topic-gdscript-plugin-integration"></a>
## 153. GDScript Android plugin integration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The instrumented Godot project is a runnable scene with test scripts for Android plugins, storage, and Java class wrapping.

GDScript test code retrieves JNISingleton objects, calls JavaClassWrapper, and connects signals through Android plugin signals.

### How the implementation applies it

#### signal_tests.gd::_init

The signal test keeps a `JNISingleton` plugin reference.

#### signal_tests.gd::test_signal_connection

The test connects two named plugin signals to GDScript callback methods.

#### java_class_wrapper_tests.gd::test_interface_object_proxy

The test constructs a GDScript object and passes it to `JavaClassWrapper.create_proxy`.

### Walk through a small source example

```text
var _plugin: JNISingleton
```

Source: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd` — _plugin

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SignalInfo

Evidence:
- Code: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd:18` — test_signal_connection
- Code: `platform/android/java/app/src/instrumented/assets/test/javaclasswrapper/java_class_wrapper_tests.gd:25` — test_interface_object_proxy
- Code: `platform/android/java/app/src/instrumented/assets/main.gd:69` — _on_gd_script_toast_button_pressed
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript declarations and typing](#topic-gdscript-declarations)
- [GDScript query-object API use](#topic-gdscript-query-objects)
<!-- rope-ladder:end document -->
