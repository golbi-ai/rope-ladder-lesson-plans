<!-- rope-ladder:begin document a7d17a6d3937168e5faba7d984c1178fdf22c971475d61b1d476c202ebd21c30 -->
# Domain concepts

This is a guided reading path, not a dictionary. Read topics in order: each section explains the problem first, names the ideas it builds on, then walks through a small piece of the codebase.

<a id="topic-crypto-resources"></a>
## 1. Cryptographic resources and contexts

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The crypto area combines persisted key or certificate resources with operational contexts.

CryptoKey and X509Certificate are Resource types, while AES, hashing, HMAC, TLS options, and crypto operations are exposed through dedicated contexts and services.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CryptoKey : public Resource {
	GDCLASS(CryptoKey, Resource);

protected:
	static void _bind_methods();
	static CryptoKey *(*_create)(bool p_notify_postinitialize);

public:
	static CryptoKey *create(bool p_notify_postinitialize = true);
	virtual Error load(const String &p_path, bool p_public_only = false) = 0;
	virtual Error save(const String &p_path, bool p_public_only = false) = 0;
	virtual String save_to_string(bool p_public_only = false) = 0;
	virtual Error load_from_string(const String &p_string_key, bool p_public_only = false) = 0;
```

Source: `core/crypto/crypto.h` — CryptoKey (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CryptoKey, X509Certificate

Evidence:
- Code: `core/crypto/crypto.h:37` — CryptoKey
- Code: `core/crypto/crypto.h:53` — X509Certificate
- Code: `core/crypto/hashing_context.h:36` — HashingContext

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Cryptographic keys, hashing, and TLS support](#topic-cryptography)
- [Resources and asset lifecycle](#topic-resources)
- [TLS connection and session state](#topic-tls-connection-state)
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-debugger-transport"></a>
## 2. Debugger capture and transport

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Remote debugging is modeled as a service with message and capture state rather than as a platform adapter.

The debugger subsystem has engine captures and profilers, local and remote debugger implementations, and a socket-backed remote debugger peer.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EngineDebugger {
public:
	typedef void (*ProfilingToggle)(void *p_user, bool p_enable, const Array &p_opts);
	typedef void (*ProfilingTick)(void *p_user, double p_frame_time, double p_process_time, double p_physics_time, double p_physics_frame_time);
	typedef void (*ProfilingAdd)(void *p_user, const Array &p_arr);

	typedef Error (*CaptureFunc)(void *p_user, const String &p_msg, const Array &p_args, bool &r_captured);

	typedef Ref<RemoteDebuggerPeer> (*CreatePeerFunc)(const String &p_uri);

	class Profiler {
		friend class EngineDebugger;
```

Source: `core/debugger/engine_debugger.h` — EngineDebugger (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/debugger/engine_debugger.h:43` — EngineDebugger
- Code: `core/debugger/remote_debugger.h:39` — RemoteDebugger
- Code: `core/debugger/remote_debugger_peer.h:59` — RemoteDebuggerPeerTCP

<a id="topic-dynamic-variant"></a>
## 3. Dynamic runtime values

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Dynamic invocation and signals](#topic-dynamic-invocation-signals), [Expression parsing and evaluation](#topic-expression-evaluation), [Dynamic arrays and dictionaries](#topic-variant-containers), [Variant text parsing and writing](#topic-variant-text-serialization).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the central interchange value across the object, container, parser, and binding code.

Variant is the tagged runtime-value representation used for construction, conversion, operators, calls, indexing, and member access.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Variant {
public:
	// If this changes the table in variant_op must be updated
	enum Type {
		NIL,

		// atomic types
		BOOL,
		INT,
		FLOAT,
		STRING,

		// math types
```

Source: `core/variant/variant.h` — Variant (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant

Evidence:
- Code: `core/variant/variant.h:94` — Variant
- Code: `core/variant/variant_construct.h:97` — VariantConstructor
- Code: `core/variant/variant_op.h:39` — CommonEvaluate

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [Device runtime](#topic-device-runtime)
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Android JNI interoperation](#topic-android-jni-interop)
- [Managed-native interop](#topic-managed-native-interop)
- [Packets, HTTP, JSON, and JSON-RPC](#topic-network-data-exchange)
- [Project settings and feature overrides](#topic-project-settings)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)

<a id="topic-dynamic-values"></a>
## 4. Dynamic values and dictionaries

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Packets, HTTP, JSON, and JSON-RPC](#topic-network-data-exchange).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

It is the common value carrier used by reflection, scene serialization, scripting, and several networking APIs.

Variant is the cross-cutting tagged value representation for scalar, math, object, callable, signal, dictionary, array, and packed-array values.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static GDType &_get_gdtype_for_type(Variant::Type p_type);
	static void _register_variant_gdtypes();
	static void _unregister_variant_gdtypes();

	struct ObjData {
		ObjectID id;
		Object *obj = nullptr;

		void ref(const ObjData &p_from);
		void ref_pointer(Object *p_object);
		void ref_pointer(RefCounted *p_object);
		void unref();
```

Source: `core/variant/variant.h` — Variant::Type (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant, JSON-RPC document

Evidence:
- Code: `core/variant/variant.h:170` — Variant::Type
- Code: `doc/classes/Dictionary.xml:2` — Dictionary

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic invocation and signals](#topic-dynamic-invocation-signals)
- [Android JNI interoperation](#topic-android-jni-interop)
- [Text break iteration](#topic-break-iteration)
- [Dynamic runtime values](#topic-dynamic-variant)
- [Expression parsing and evaluation](#topic-expression-evaluation)
- [Generic container infrastructure](#topic-generic-containers)
- [Managed-native interop](#topic-managed-native-interop)
- [Project settings and feature overrides](#topic-project-settings)
- [Dynamic arrays and dictionaries](#topic-variant-containers)
- [Variant text parsing and writing](#topic-variant-text-serialization)
- [Engine object model](#topic-object-model)

<a id="topic-memory-allocation"></a>
## 5. Engine allocation helpers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Memory helpers define default allocators, typed allocators, allocation result wrappers, and global nil support.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class DefaultAllocator {
public:
	_FORCE_INLINE_ static void *alloc(size_t p_memory) { return Memory::alloc_static(p_memory, false); }
	_FORCE_INLINE_ static void free(void *p_ptr) { Memory::free_static(p_ptr, false); }
};

// Overload of `new` operator to use the `Memory::alloc_static()` function.
// The `DefaultAllocator` parameter is just a tag to select this overload.
// NOTE: do not inline `new` operators due to GCC+LTO compiler bug (see GH-119752).
void *operator new(size_t p_size, DefaultAllocator p_allocator);

// Overload of `new` operator to use a custom allocation function.
void *operator new(size_t p_size, void *(*p_allocfunc)(size_t p_size));
```

Source: `core/os/memory.h` — DefaultAllocator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/os/memory.h:101` — DefaultAllocator
- Code: `core/os/memory.h:300` — DefaultTypedAllocator
- Code: `core/os/memory.h:132` — memnew_result

<a id="topic-object-model"></a>
## 6. Engine object model

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Class metadata and reflection](#topic-reflection), [Scene graph and lifecycle](#topic-scene-tree).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the base model that the scene, resource, script, and editor APIs build upon.

Object-derived types provide the common runtime identity and behavior base, while RefCounted supplies reference-counted lifetime for resource-like values.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct ObjectGDExtension {
	GDExtension *library = nullptr;
	ObjectGDExtension *parent = nullptr;
	List<ObjectGDExtension *> children;
	StringName parent_class_name;
	StringName class_name;
	bool editor_class = false;
	bool reloadable = false;
	bool is_virtual = false;
	bool is_abstract = false;
	bool is_exposed = true;
#ifdef TOOLS_ENABLED
	bool is_runtime = false;
```

Source: `core/object/object.h` — Object (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node, Resource, ClassInfo

Evidence:
- Code: `core/object/object.h:132` — Object
- Code: `core/object/ref_counted.h:36` — RefCounted

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [Device runtime](#topic-device-runtime)
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Collision filtering](#topic-collision-filtering)
- [Deferred calls and worker tasks](#topic-deferred-execution)
- [Dynamic invocation and signals](#topic-dynamic-invocation-signals)
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Metal-cpp object bridge](#topic-metal-cpp-object-bridge)
- [2D and 3D navigation queries](#topic-navigation-queries)
- [ObjectDB snapshots](#topic-objectdb-snapshots)
- [2D and 3D physics queries](#topic-physics-server-queries)
- [3D physics query contracts](#topic-physics-space-queries)
- [Reflective property inspection](#topic-property-inspection)
- [Reflection metadata](#topic-reflection-metadata)
- [Regular-expression matching](#topic-regular-expression-results)
- [Resources and asset lifecycle](#topic-resources)
- [Main loop, OS, and time services](#topic-runtime-loop-time)
- [Script languages and instances](#topic-script-hosting)
- [Script resources and instances](#topic-scripting)
- [RTTI-based serialization](#topic-serialization)
- [Text shaping plans](#topic-text-shaping-plans)
- [Undo and redo actions](#topic-undo-redo)
- [UPnP device control](#topic-upnp-device-control)
- [Web JavaScript bridge](#topic-web-javascript-bridge)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)

<a id="topic-gdextension-libraries"></a>
## 7. GDExtension libraries

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The extension runtime exposes a generated interface while retaining library and class lifecycle state in the engine.

GDExtension is a Resource that holds a loader, registered extension classes, initialization state, and library open, initialize, deinitialize, and close operations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
class CallableCustomExtension : public CallableCustom {
	void *userdata;
	void *token;

	ObjectID object;

	GDExtensionCallableCustomCall call_func;
	GDExtensionCallableCustomIsValid is_valid_func;
	GDExtensionCallableCustomFree free_func;

	GDExtensionCallableCustomEqual equal_func;
	GDExtensionCallableCustomLessThan less_than_func;
```

Source: `core/extension/gdextension_interface.cpp` — CallableCustomExtension (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDExtension

Evidence:
- Code: `core/extension/gdextension.h:39` — GDExtension
- Code: `core/extension/gdextension_interface.cpp:46` — CallableCustomExtension
- Code: `core/extension/gdextension_library_loader.h:39` — GDExtensionLibraryLoader

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Extension API compatibility baselines](#topic-extension-api-compatibility)
- [Virtual implementation extensions](#topic-extensions)
- [Native extension loading](#topic-native-extensions)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-geometry-transforms"></a>
## 8. Geometry and transform values

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Mesh geometry algorithms](#topic-mesh-geometry-algorithms), [Spatial indexing and ray queries](#topic-spatial-indexing), [Transform interpolation](#topic-transform-interpolation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These are the coordinate and bounds values shared by geometric algorithms and dynamic argument conversion.

The math layer defines 2D and 3D vectors, rectangles, boxes, planes, bases, quaternions, projections, and transforms as reusable value types.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct [[nodiscard]] AABB {
	Vector3 position;
	Vector3 size;

	real_t get_volume() const;
	_FORCE_INLINE_ bool has_volume() const {
		return size.x > 0.0f && size.y > 0.0f && size.z > 0.0f;
	}

	_FORCE_INLINE_ bool has_surface() const {
		return size.x > 0.0f || size.y > 0.0f || size.z > 0.0f;
	}
```

Source: `core/math/aabb.h` — AABB (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

AABB, Transform3D

Evidence:
- Code: `core/math/aabb.h:38` — AABB
- Code: `core/math/transform_3d.h:38` — Transform3D
- Code: `core/math/projection.h:45` — Projection

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Picture planes and pixel ownership](#topic-input-picture-planes)

<a id="topic-input-events-actions"></a>
## 9. Input events and actions

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Platform display and window adaptation](#topic-platform-display-adaptation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation separates concrete device events from named action bindings.

InputEvent resource subclasses represent key, mouse, joypad, touch, gesture, MIDI, shortcut, and action input; InputMap associates named actions with InputEvent lists and deadzones.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class InputEvent : public Resource {
	GDCLASS(InputEvent, Resource);

	int device = 0;

protected:
	bool canceled = false;
	bool pressed = false;

	static void _bind_methods();

public:
	static constexpr int DEVICE_ID_EMULATION = -1;
```

Source: `core/input/input_event.h` — InputEvent (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

InputEvent, InputMap::Action

Evidence:
- Code: `core/input/input_event.h:52` — InputEvent
- Code: `core/input/input_event.h:151` — InputEventKey
- Code: `core/input/input_map.h` — InputMap::Action

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Input events and actions](#topic-input-actions)
- [Keyboard and MIDI parsing](#topic-input-midi)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-input-midi"></a>
## 10. Keyboard and MIDI parsing

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Keyboard utilities map key codes and names, while MIDIDriver::Parser interprets status and data bytes into MIDI messages.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void MIDIDriver::Parser::parse_fragment(uint8_t p_fragment) {
	switch (category(p_fragment)) {
		case MessageCategory::RealTime:
			// Real-Time messages are single byte messages that can
			// occur at any point and do not interrupt other messages.
			// We pass them straight through.
			MIDIDriver::send_event(device_index, p_fragment);
			break;

		case MessageCategory::SysExBegin:
			status_byte = p_fragment;
			skipping_sys_ex = true;
			break;
```

Source: `core/os/midi_driver.cpp` — MIDIDriver::Parser::parse (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/os/keyboard.h:343` — find_keycode_name
- Code: `core/os/midi_driver.h` — MIDIDriver::Parser
- Code: `core/os/midi_driver.cpp` — MIDIDriver::Parser::parse

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Apple embedded hosting](#topic-apple-embedded-hosting)
- [Browser runtime adapters](#topic-browser-runtime-adapters)
- [Input events and actions](#topic-input-events-actions)
- [MIDI input adapters](#topic-midi-input-adapters)

<a id="topic-runtime-loop-time"></a>
## 11. Main loop, OS, and time services

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

OS hosts platform runtime services, MainLoop defines the loop-facing object type, and Time exposes date and time behavior.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class MainLoop : public Object {
	GDCLASS(MainLoop, Object);

protected:
	static void _bind_methods();

	GDVIRTUAL0(_initialize)
	GDVIRTUAL1R(bool, _physics_process, double)
	GDVIRTUAL1R(bool, _process, double)
	GDVIRTUAL0(_finalize)

public:
	enum {
```

Source: `core/os/main_loop.h` — MainLoop (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/os/os.h:46` — OS
- Code: `core/os/main_loop.h:36` — MainLoop
- Code: `core/os/time.h:37` — Time

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Motion blur](#topic-motion-blur)
- [Engine object model](#topic-object-model)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Android Gradle assembly](#topic-android-gradle-assembly)
- [CPU-specialized DSP backends](#topic-cpu-specialized-dsp)
- [Motion-blur geometry](#topic-motion-blur-geometry)
- [Project catalog and selection](#topic-project-catalog)
- [Ray query state](#topic-ray-query)
- [Compile-time platform backends](#topic-sdl-platform-backends)
- [Wraparound network time](#topic-wraparound-network-time)

<a id="topic-network-io"></a>
## 12. Networking and transport I/O

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Protocol-facing abstractions are organized in core I/O, with Web-specific implementations for browser constraints.

The core exposes HTTP clients, stream peers, packet peers, TCP and UDP servers, DTLS and TLS peers, IP resolution, and Unix-domain socket support.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class HTTPClient : public RefCounted {
	GDCLASS(HTTPClient, RefCounted);

public:
	enum ResponseCode {
		// 1xx informational
		RESPONSE_CONTINUE = 100,
		RESPONSE_SWITCHING_PROTOCOLS = 101,
		RESPONSE_PROCESSING = 102,

		// 2xx successful
		RESPONSE_OK = 200,
		RESPONSE_CREATED = 201,
```

Source: `core/io/http_client.h` — HTTPClient (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

IPAddress

Evidence:
- Code: `core/io/http_client.h:37` — HTTPClient
- Code: `core/io/packet_peer.h:39` — PacketPeer
- Code: `core/io/stream_peer.h:38` — StreamPeer
- Code: `core/io/ip_address.h:35` — IPAddress

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Godot ENet socket adaptation](#topic-godot-enet-socket-adaptation)
- [WebRTC data channels](#topic-webrtc-data-channels)
- [Ray packets and streams](#topic-packet-queries)
- [TLS connection and session state](#topic-tls-connection-state)
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-object-identity-lifetime"></a>
## 13. Object identity and lifetime

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Dynamic invocation and signals](#topic-dynamic-invocation-signals), [Reflection metadata](#topic-reflection-metadata), [Script languages and instances](#topic-script-hosting).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Object supplies engine object identity and ObjectDB support, while RefCounted and Ref provide reference-managed object use.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct ObjectGDExtension {
	GDExtension *library = nullptr;
	ObjectGDExtension *parent = nullptr;
	List<ObjectGDExtension *> children;
	StringName parent_class_name;
	StringName class_name;
	bool editor_class = false;
	bool reloadable = false;
	bool is_virtual = false;
	bool is_abstract = false;
	bool is_exposed = true;
#ifdef TOOLS_ENABLED
	bool is_runtime = false;
```

Source: `core/object/object.h` — Object (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Object, ObjectID

Evidence:
- Code: `core/object/object.h:132` — Object
- Code: `core/object/object_id.h:41` — ObjectID
- Code: `core/object/ref_counted.h:36` — RefCounted

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [2D and 3D navigation queries](#topic-navigation-queries)
- [Engine object model](#topic-object-model)
- [2D and 3D physics queries](#topic-physics-server-queries)
- [3D physics query contracts](#topic-physics-space-queries)
- [Regular-expression matching](#topic-regular-expression-results)
- [Resources and asset lifecycle](#topic-resources)
- [UPnP device control](#topic-upnp-device-control)

<a id="topic-project-settings"></a>
## 14. Project settings and feature overrides

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the central configuration store for project-level state.

ProjectSettings stores named Variant values with persistence metadata, feature overrides, autoload definitions, resource paths, and a change version.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ProjectSettings : public Object {
	GDCLASS(ProjectSettings, Object);
	_THREAD_SAFE_CLASS_
	friend class TestProjectSettingsInternalsAccessor;

	bool is_changed = false;

	// Starting version from 1 ensures that all callers can reset their tested version to 0,
	// and will always detect the initial project settings as a "change".
	uint32_t _version = 1;

	// Track changed settings for get_changed_settings functionality
	HashSet<StringName> changed_settings;
```

Source: `core/config/project_settings.h` — ProjectSettings (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ProjectSettings, ProjectSettings::VariantContainer

Evidence:
- Code: `core/config/project_settings.h:40` — ProjectSettings
- Code: `core/config/project_settings.h` — ProjectSettings::VariantContainer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Dynamic runtime values](#topic-dynamic-variant)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Class metadata and reflection](#topic-reflection)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-random-generation"></a>
## 15. Pseudo-random generation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

RandomPCG supplies pseudo-random generation and RandomNumberGenerator exposes a reference-counted runtime wrapper.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class RandomNumberGenerator : public RefCounted {
	GDCLASS(RandomNumberGenerator, RefCounted);

protected:
	RandomPCG randbase;

	static void _bind_methods();

public:
	_FORCE_INLINE_ void set_seed(uint64_t p_seed) { randbase.seed(p_seed); }
	_FORCE_INLINE_ uint64_t get_seed() { return randbase.get_seed(); }

	_FORCE_INLINE_ void set_state(uint64_t p_state) { randbase.set_state(p_state); }
```

Source: `core/math/random_number_generator.h` — RandomNumberGenerator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/math/random_pcg.h:54` — RandomPCG
- Code: `core/math/random_number_generator.h:36` — RandomNumberGenerator

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [Device runtime](#topic-device-runtime)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Native support algorithms](#topic-native-support-algorithms)

<a id="topic-resource-loading"></a>
## 16. Resource loading and caching

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Packed resource files](#topic-packed-resource-files), [Resource formats and serialization](#topic-resource-formats).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Loading is a registry-based service rather than a single file-format implementation.

ResourceLoader selects registered ResourceFormatLoader instances, applies cache modes, reports dependencies, and tracks threaded load tasks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
private:
	friend class ResourceCache;

	String name;
	String path_cache;
	String scene_unique_id;

#ifdef TOOLS_ENABLED
	uint64_t last_modified_time = 0;
	uint64_t import_last_modified_time = 0;
	String import_path;
#endif

	enum EmitChangedState {
```

Source: `core/io/resource.h` — ResourceCache (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource, ResourceLoader::ThreadLoadTask

Evidence:
- Code: `core/io/resource_loader.h:40` — ResourceLoader
- Code: `core/io/resource_loader.h` — ResourceLoader::ThreadLoadTask
- Code: `core/io/resource.h:70` — ResourceCache

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [HDR, JPEG, and KTX loading](#topic-image-format-loading)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-resources"></a>
## 17. Resources and asset lifecycle

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Native extension loading](#topic-native-extensions), [Packed scene serialization](#topic-packed-scenes), [Physics shapes, objects, and collision reports](#topic-physics-collision), [Meshes, materials, textures, and instancing](#topic-rendering-assets), [Script resources and instances](#topic-scripting), [Skeletal modifiers and inverse kinematics](#topic-skeletal-ik).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Resources are shared data objects used by scenes, scripts, meshes, shapes, materials, and extensions.

A Resource is a RefCounted asset value with a path, name, scene-local option, and loader or saver lifecycle.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Resource : public RefCounted {
	GDCLASS(Resource, RefCounted);

public:
	static constexpr AncestralClass static_ancestral_class = AncestralClass::RESOURCE;

	static void register_custom_data_to_otdb();
	virtual String get_base_extension() const { return "res"; }

protected:
	static void _add_resource_base_extension_to_classdb(const String &p_extension, const String &p_class);
	struct DuplicateParams {
		bool deep = false;
```

Source: `core/io/resource.h` — Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource, Mesh, Material, GDExtension

Evidence:
- Code: `core/io/resource.h:44` — Resource
- Code: `doc/classes/Resource.xml` — Resource.resource_path

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Engine object model](#topic-object-model)
- [Allocation and reference ownership](#topic-allocation)
- [Audio buses, streams, and effects](#topic-audio-bus-processing)
- [Cryptographic resources and contexts](#topic-crypto-resources)
- [Editing selection history](#topic-editing-selection-history)
- [Editor plugins and customization hooks](#topic-editor-extensibility)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Project filesystem index](#topic-filesystem-project-index)
- [GDExtension libraries](#topic-gdextension-libraries)
- [Glyph and face caching](#topic-glyph-caching)
- [GPU command encoding](#topic-gpu-command-encoding)
- [D3D12 memory allocation](#topic-gpu-memory-suballocation)
- [Image codec integration](#topic-image-codec-registration)
- [Input events and actions](#topic-input-events-actions)
- [Ogg Vorbis audio streams](#topic-ogg-vorbis-audio-streams)
- [Packed resource files](#topic-packed-resource-files)
- [PNG image codec](#topic-png-image-codec)
- [Project settings and feature overrides](#topic-project-settings)
- [Textures, meshes, and materials](#topic-rendering-resources)
- [Resource-backed assets](#topic-resource-assets)
- [Resource-bundle data](#topic-resource-bundle-data)
- [Resource dependency resolution](#topic-resource-dependency-resolution)
- [Resource formats and serialization](#topic-resource-formats)
- [Resource and server identifiers](#topic-resource-identifiers)
- [Standalone resource importing](#topic-resource-importing)
- [Resource loading and caching](#topic-resource-loading)
- [Asynchronous resource previews](#topic-resource-previews)
- [Resource-specific editing](#topic-resource-specific-editors)
- [Shader programs and material parameters](#topic-shader-programs)
- [Skeleton modification and physical-bone simulation](#topic-skeleton-modifiers)
- [Themes and style boxes](#topic-themes-and-style-boxes)
- [Theora video streams](#topic-theora-video-streams)
- [3D shapes, worlds, and skins](#topic-three-dimensional-content)
- [Tile libraries, cells, and patterns](#topic-tile-libraries)
- [2D shapes, tiles, and paths](#topic-two-dimensional-content)
- [Variant text parsing and writing](#topic-variant-text-serialization)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-runtime-metadata"></a>
## 18. Runtime class metadata

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Virtual implementation extensions](#topic-extensions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This metadata is the bridge between native engine classes and the exposed API.

The runtime records classes, inheritance, methods, properties, signals, and instantiation capability for engine objects.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

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

RID, Resource

Evidence:
- Code: `core/object/class_db.h:62` — ClassDB

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Native extension loading](#topic-native-extensions)
- [Class metadata and reflection](#topic-reflection)
- [Reflection metadata](#topic-reflection-metadata)
- [Shader interface mapping and reflection](#topic-shader-interface-analysis)
- [Shader interface metadata](#topic-shader-interface-metadata)
- [Shader reflection](#topic-shader-reflection)
- [SPIR-V shader reflection](#topic-spirv-shader-reflection)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-string-names-paths"></a>
## 19. Strings, interned names, and node paths

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Translation and locale selection](#topic-localization), [Profiling name interning](#topic-profiling-interning), [Variant text parsing and writing](#topic-variant-text-serialization).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The string layer implements Unicode strings, hash-backed interned StringName values, path and subpath storage in NodePath, and fuzzy-search helpers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct FuzzySearchToken {
	int idx = -1;
	String string;

	bool try_exact_match(FuzzyTokenMatch &p_match, const String &p_target, int p_offset) const;
	bool try_fuzzy_match(FuzzyTokenMatch &p_match, const String &p_target, int p_offset, int p_miss_budget) const;
};

class FuzzyTokenMatch {
	friend struct FuzzySearchToken;
	friend class FuzzySearchMatch;
	friend class FuzzySearch;
```

Source: `core/string/fuzzy_search.h` — FuzzySearch (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

StringName, NodePath

Evidence:
- Code: `core/string/ustring.h:39` — String
- Code: `core/string/string_name.h:40` — StringName
- Code: `core/string/node_path.h:39` — NodePath
- Code: `core/string/fuzzy_search.h:49` — FuzzySearch

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene replication configuration](#topic-scene-replication-configuration)

<a id="topic-synchronization-primitives"></a>
## 20. Threads and synchronization

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The OS layer defines Thread, mutexes, condition variables, read-write locks, semaphores, spin locks, and safe binary mutexes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ConditionVariable {
	mutable THREADING_NAMESPACE::condition_variable condition;

public:
	template <typename BinaryMutexT>
	_ALWAYS_INLINE_ void wait(const MutexLock<BinaryMutexT> &p_lock) const {
		condition.wait(p_lock._get_lock());
	}

	template <int Tag>
	_ALWAYS_INLINE_ void wait(const MutexLock<SafeBinaryMutex<Tag>> &p_lock) const {
		condition.wait(p_lock.mutex._get_lock());
	}
```

Source: `core/os/condition_variable.h` — ConditionVariable (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `core/os/thread.h:36` — Thread
- Code: `core/os/mutex.h:47` — MutexLock
- Code: `core/os/condition_variable.h:53` — ConditionVariable
- Code: `core/os/rw_lock.h:44` — RWLock

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Event queue and watches](#topic-sdl-event-queue)
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)
- [Task scheduling and synchronization](#topic-task-scheduling)

<a id="topic-physics-space-queries"></a>
## 21. 3D physics query contracts

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The public query objects retain typed parameters while server-side direct-space operations consume them.

Physics-server types define ray, point, shape, and motion parameter and result records; RefCounted query objects expose those contracts to callers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class PhysicsRayQueryParameters3D : public RefCounted {
	GDCLASS(PhysicsRayQueryParameters3D, RefCounted);

	PS3DT::RayParameters parameters;

protected:
	static void _bind_methods();

public:
	static Ref<PhysicsRayQueryParameters3D> create(Vector3 p_from, Vector3 p_to, uint32_t p_mask, const TypedArray<RID> &p_exclude);
	const PS3DT::RayParameters &get_parameters() const { return parameters; }

	void set_from(const Vector3 &p_from) { parameters.from = p_from; }
```

Source: `servers/physics_3d/queries/physics_ray_query_parameters_3d.h` — PhysicsRayQueryParameters3D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PhysicsRayQueryParameters3D

Evidence:
- Code: `servers/physics_3d/physics_server_3d_types.h` — PS3DT::RayParameters, RayResult, PointParameters, ShapeParameters, MotionParameters, MotionResult
- Code: `servers/physics_3d/queries/physics_ray_query_parameters_3d.h:36` — PhysicsRayQueryParameters3D
- Code: `servers/physics_3d/queries/physics_testmotion_query_result_3d.h:36` — PhysicsTestMotionResult3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision shapes](#topic-collision-shapes)
- [Motion blur](#topic-motion-blur)
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Engine object model](#topic-object-model)
- [Physics queries and motion tests](#topic-physics-queries)
- [Ray query state](#topic-ray-query)

<a id="topic-animation-blending"></a>
## 22. Animation blending and playback

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Animation graph resources retain nodes, blend points, transitions, and node connections.

AnimationMixer, AnimationPlayer, AnimationTree, blend spaces, blend trees, and state machines compose animation playback.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class AnimationNodeBlendTree : public AnimationRootNode {
	GDCLASS(AnimationNodeBlendTree, AnimationRootNode);

	struct Node {
		Ref<AnimationNode> node;
		Vector2 position;
		LocalVector<StringName> connections;
	};

	AHashMap<StringName, Node> nodes;

	Vector2 graph_offset;
```

Source: `scene/animation/animation_blend_tree.h` — AnimationNodeBlendTree (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

AnimationMixer, AnimationNodeBlendTree, AnimationNodeStateMachine

Evidence:
- Code: `scene/animation/animation_mixer.h:42` — AnimationMixer
- Code: `scene/animation/animation_blend_tree.h:407` — AnimationNodeBlendTree
- Code: `scene/animation/animation_node_state_machine.h:109` — AnimationNodeStateMachine

<a id="topic-canvas-and-viewports"></a>
## 23. Canvas and viewport presentation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [GUI controls and layout](#topic-gui-controls), [Display, camera, video, and movie services](#topic-platform-presentation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

CanvasItem is the 2D visual branch; Window is a Viewport subclass.

CanvasItem, CanvasLayer, Viewport, and Window extend the Node tree with visual presentation and window-facing state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CanvasItem : public Node {
	GDCLASS(CanvasItem, Node);

	friend class CanvasLayer;

public:
	static constexpr AncestralClass static_ancestral_class = AncestralClass::CANVAS_ITEM;

	enum TextureFilter {
		TEXTURE_FILTER_PARENT_NODE,
		TEXTURE_FILTER_NEAREST,
		TEXTURE_FILTER_LINEAR,
		TEXTURE_FILTER_NEAREST_WITH_MIPMAPS,
```

Source: `scene/main/canvas_item.h` — class CanvasItem : public Node (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node

Evidence:
- Code: `scene/main/canvas_item.h:47` — class CanvasItem : public Node
- Code: `scene/main/canvas_layer.h:37` — class CanvasLayer : public Node
- Code: `scene/main/viewport.h:96` — class Viewport : public Node
- Code: `scene/main/window.h:42` — class Window : public Viewport

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Scene graph and lifecycle](#topic-scene-tree)
- [Swapchain presentation](#topic-vulkan-swapchain-presentation)

<a id="topic-navigation"></a>
## 24. Navigation maps and path queries

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Both 2D and 3D APIs expose equivalent navigation query objects and server services.

Navigation agents use server maps, regions, path-query parameters, and path-query results to follow a target position.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="NavigationPathQueryParameters2D" inherits="RefCounted" api_type="core" experimental="" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Provides parameters for 2D navigation path queries.
	</brief_description>
	<description>
		By changing various properties of this object, such as the start and target position, you can configure path queries to the [NavigationServer2D].
	</description>
	<tutorials>
		<link title="Using NavigationPathQueryObjects">$DOCS_URL/tutorials/navigation/navigation_using_navigationpathqueryobjects.html</link>
	</tutorials>
	<members>
		<member name="excluded_regions" type="RID[]" setter="set_excluded_regions" getter="get_excluded_regions" default="[]">
			The list of region [RID]s that will be excluded from the path query. Use [method NavigationRegion2D.get_rid] to get the [RID] associated with a [NavigationRegion2D] node.
```

Source: `doc/classes/NavigationPathQueryParameters2D.xml` — NavigationPathQueryParameters2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

NavigationPathQueryParameters2D, NavigationPathQueryResult2D

Evidence:
- Code: `servers/navigation_server_2d.h` — NavigationServer2D
- Code: `doc/classes/NavigationPathQueryParameters2D.xml:2` — NavigationPathQueryParameters2D
- Code: `doc/classes/NavigationPathQueryResult2D.xml:2` — NavigationPathQueryResult2D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Navigation agents and regions](#topic-navigation-agents)
- [Navigation region construction](#topic-navigation-regions)

<a id="topic-property-introspection"></a>
## 25. Property-list and scene-property helpers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These helpers bridge reflective property operations with scene state.

PropertyListHelper resolves slash-delimited property names to property records, while PropertyUtils compares properties and walks scene-state pack data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const PropertyListHelper::Property *PropertyListHelper::_get_property(const String &p_property, int *r_index, bool p_allow_oob) const {
	const Vector<String> components = p_property.rsplit("/", true, 1);
	if (components.size() < 2 || !components[0].begins_with(prefix)) {
		return nullptr;
	}

	const String index_string = components[0].trim_prefix(prefix);
	if (!index_string.is_valid_int()) {
		return nullptr;
	}

	int index = index_string.to_int();
	if (index < 0 || (!p_allow_oob && index >= _call_array_length_getter())) {
```

Source: `scene/property_list_helper.cpp` — PropertyListHelper::_get_property (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/property_list_helper.h` — class PropertyListHelper and PropertyListHelper::Property
- Code: `scene/property_list_helper.cpp:56` — PropertyListHelper::_get_property
- Code: `scene/property_utils.cpp` — PropertyUtils scene-state PackState accesses

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)

<a id="topic-runtime-configuration"></a>
## 26. Runtime configuration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Engine bootstrap](#topic-engine-bootstrap).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The main executable and a Godot project file both provide concrete configuration inputs.

Runtime configuration reads project settings such as the main scene and boot-splash options; the application project file also declares its main scene.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
// an early exit with that error code.
int Main::start() {
	GodotProfileZone("start");
	OS::get_singleton()->benchmark_begin_measure("Startup", "Main::Start");

	ERR_FAIL_COND_V(!_start_success, EXIT_FAILURE);

	bool has_icon = false;
	String positional_arg;
	String game_path;
	String script;
	String main_loop_type;
	bool check_only = false;
```

Source: `main/main.cpp` — Main::start (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `main/main.cpp:4016` — Main::start
- Code: `platform/android/java/app/src/instrumented/assets/project.godot:15` — [application]

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Input action and shortcut configuration](#topic-input-action-configuration)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-scene-graph"></a>
## 27. Scene graph nodes

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [GUI controls and layout](#topic-gui-control-layout), [Navigation agents and regions](#topic-navigation-agents), [CPU and GPU particle systems](#topic-particle-systems), [Skeletal modifiers and inverse kinematics](#topic-skeletal-modifiers-and-ik), [3D physics nodes](#topic-three-dimensional-physics), [Tile map layer data](#topic-tilemap-layer-data), [2D physics nodes](#topic-two-dimensional-physics).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

2D nodes inherit from CanvasItem through Node2D, while 3D nodes inherit from Node through Node3D.

Node2D and Node3D are scene graph node bases for specialized runtime behavior.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Node2D : public CanvasItem {
	GDCLASS(Node2D, CanvasItem);

	mutable MTFlag xform_dirty;
	mutable Point2 position;
	mutable real_t rotation = 0.0;
	mutable Size2 scale = Vector2(1, 1);
	mutable real_t skew = 0.0;

	Transform2D transform;

	_FORCE_INLINE_ bool _is_xform_dirty() const { return is_group_processing() ? xform_dirty.mt.is_set() : xform_dirty.st; }
	void _set_xform_dirty(bool p_dirty) const;
```

Source: `scene/2d/node_2d.h` — Node2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node2D, Node3D

Evidence:
- Code: `scene/2d/node_2d.h:35` — Node2D
- Code: `scene/3d/node_3d.h:50` — Node3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [Device runtime](#topic-device-runtime)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene graph and lifecycle](#topic-scene-tree)

<a id="topic-scene-tree-execution"></a>
## 28. SceneTree execution

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the runtime coordinator above individual nodes.

A SceneTree schedules the Node hierarchy and maintains scene-level timer and tween processing collections.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class SceneTreeTimer : public RefCounted {
	GDCLASS(SceneTreeTimer, RefCounted);

	double time_left = 0.0;
	bool process_always = true;
	bool process_in_physics = false;
	bool ignore_time_scale = false;

protected:
	static void _bind_methods();

public:
	void set_time_left(double p_time);
```

Source: `scene/main/scene_tree.h` — class SceneTree (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node

Evidence:
- Code: `scene/main/scene_tree.h:89` — class SceneTree
- Code: `scene/main/scene_tree.cpp` — SceneTree timer and tween processing blocks

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene graph and lifecycle](#topic-scene-tree)

<a id="topic-xr-tracking"></a>
## 29. XR tracking and poses

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Trackers classify sources and poses carry the spatial data exposed for those sources.

XRServer coordinates interfaces and tracker types, while positional, body, controller, face, and hand trackers publish XR pose state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class XRBodyTracker : public XRPositionalTracker {
	GDCLASS(XRBodyTracker, XRPositionalTracker);
	_THREAD_SAFE_CLASS_

public:
	enum BodyFlags {
		BODY_FLAG_UPPER_BODY_SUPPORTED = 1,
		BODY_FLAG_LOWER_BODY_SUPPORTED = 2,
		BODY_FLAG_HANDS_SUPPORTED = 4,
	};

	enum Joint {
		JOINT_ROOT,
```

Source: `servers/xr/xr_body_tracker.h` — XRBodyTracker (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

XRTracker, XRPose

Evidence:
- Code: `servers/xr/xr_server.h` — XRServer::RenderState
- Code: `servers/xr/xr_tracker.h:40` — XRTracker
- Code: `servers/xr/xr_pose.h:36` — XRPose
- Code: `servers/xr/xr_positional_tracker.h:44` — XRPositionalTracker
- Code: `servers/xr/xr_body_tracker.h:35` — XRBodyTracker
- Code: `servers/xr/xr_hand_tracker.h:35` — XRHandTracker

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Pluggable and wrapped server backends](#topic-pluggable-server-backends)
- [Skeletal modifiers and inverse kinematics](#topic-skeletal-modifiers-and-ik)

<a id="topic-openxr-binding-modifiers"></a>
## 30. Binding modifiers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation has a base modifier plus profile-level and action-binding subclasses, including D-pad and analog-threshold modifiers.

Binding modifiers alter interaction-profile or per-action binding data before it is submitted to OpenXR.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class OpenXRDpadBindingModifier : public OpenXRIPBindingModifier {
	GDCLASS(OpenXRDpadBindingModifier, OpenXRIPBindingModifier);

private:
	PackedByteArray dpad_bindings_data;
	XrInteractionProfileDpadBindingEXT *dpad_bindings = nullptr;
	String input_path;
	Ref<OpenXRActionSet> action_set;
	Ref<OpenXRHapticBase> on_haptic;
	Ref<OpenXRHapticBase> off_haptic;

protected:
	static void _bind_methods();
```

Source: `modules/openxr/extensions/openxr_dpad_binding_extension.h` — class OpenXRDpadBindingModifier : public OpenXRIPBindingModifier (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OpenXRIPBinding

Evidence:
- Code: `modules/openxr/extensions/openxr_dpad_binding_extension.h:61` — class OpenXRDpadBindingModifier : public OpenXRIPBindingModifier
- Code: `modules/openxr/extensions/openxr_valve_analog_threshold_extension.h:60` — class OpenXRAnalogThresholdModifier : public OpenXRActionBindingModifier
- Code: `modules/openxr/doc_classes/OpenXRBindingModifier.xml` — OpenXRBindingModifier::_get_ip_modification

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [OpenXR runtime integration](#topic-openxr-runtime-integration)

<a id="topic-csharp-translation-extraction"></a>
## 31. C# translation extraction

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The parser builds parse options from project define constants and resolves invocation symbols through a semantic model.

The editor parses C# source into a Roslyn compilation and extracts constant translation-call arguments and nearby comments.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
public partial class CsTranslationParserPlugin : EditorTranslationParserPlugin
{

    private class CommentData
    {
        public string Comment = "";
        public int StartLine;
        public int EndLine;
        public bool Newline = true;
    }

    private List<MetadataReference>? _projectReferences;
    private Array<string[]> _ret = new Array<string[]>();
```

Source: `modules/mono/editor/GodotTools/GodotTools/CsTranslationParserPlugin.cs` — CsTranslationParserPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools/CsTranslationParserPlugin.cs:16` — CsTranslationParserPlugin
- Code: `modules/mono/editor/GodotTools/GodotTools/CsTranslationParserPlugin.cs:89` — GetProjectDefineConstants

<a id="topic-dotnet-editor-build-integration"></a>
## 32. Editor build integration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation includes project-file editing, a build logger that writes an issues CSV, build management, and problem views.

Editor tools generate .NET projects, invoke builds, and surface CSV diagnostics.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
{
    public static class BuildManager
    {
        private static BuildInfo? _buildInProgress;

        public const string MsBuildIssuesFileName = "msbuild_issues.csv";
        private const string MsBuildLogFileName = "msbuild_log.txt";

        public delegate void BuildLaunchFailedEventHandler(BuildInfo buildInfo, string reason);

        public static event BuildLaunchFailedEventHandler? BuildLaunchFailed;
        public static event Action<BuildInfo>? BuildStarted;
        public static event Action<BuildResult>? BuildFinished;
        public static event Action<string?>? StdOutputReceived;
```

Source: `modules/mono/editor/GodotTools/GodotTools/Build/BuildManager.cs` — BuildManager (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BuildDiagnostic

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools.ProjectEditor/ProjectGenerator.cs` — ProjectGenerator.GenGameProject
- Code: `modules/mono/editor/GodotTools/GodotTools.BuildLogger/GodotBuildLogger.cs` — GodotBuildLogger.Initialize
- Code: `modules/mono/editor/GodotTools/GodotTools/Build/BuildManager.cs:12` — BuildManager
- Code: `modules/mono/editor/GodotTools/GodotTools/Build/BuildProblemsView.cs` — BuildProblemsView.ReadDiagnosticsFromFile

<a id="topic-extension-api-compatibility"></a>
## 33. Extension API compatibility baselines

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The baseline files record expected API changes such as added arguments, changed types, and removed APIs.

Version-pair directories retain expected extension-API validation differences, and the validator aggregates their relevant diagnostic lines.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
get_expected_output()
{
  local parts=()
  IFS='_' read -ra parts <<< "$(basename "$1")"

  if [[ "${#parts[@]}" == "2" ]]; then
    while read -r file; do
      cat "$file" >> "$expected_errors"
    done <<< "$(find "$1" -type f -name "*.txt")"

    next="$(find "$api_validation_dir" -type d -name "${parts[1]}*")"
    if [[ "$next" != "" ]]; then
      get_expected_output "$next"
```

Source: `misc/scripts/validate_extension_api.sh` — get_expected_output (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `misc/scripts/validate_extension_api.sh:37` — get_expected_output
- Code: `misc/extension_api_validation/README.md` — expected output format

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDExtension libraries](#topic-gdextension-libraries)
- [Native extension loading](#topic-native-extensions)
- [Virtual implementation extensions](#topic-extensions)

<a id="topic-fbx-gltf-scene-import"></a>
## 34. FBX scene import through GLTF structures

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation maps FBX scene, node, mesh, material, texture, animation, skin, and light data into GLTF-oriented state.

FBXDocument and FBXState specialize GLTF document and state structures while importers expose UFBX and FBX2GLTF editor entry points.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorSceneFormatImporterUFBX : public EditorSceneFormatImporter {
	GDCLASS(EditorSceneFormatImporterUFBX, EditorSceneFormatImporter);

public:
	enum FBX_IMPORTER_TYPE {
		FBX_IMPORTER_UFBX,
		FBX_IMPORTER_FBX2GLTF,
	};
	virtual void get_extensions(List<String> *r_extensions) const override;
	virtual Node *import_scene(const String &p_path, uint32_t p_flags,
			const HashMap<StringName, Variant> &p_options,
			List<String> *r_missing_deps, Error *r_err = nullptr) override;
	virtual void get_import_options(const String &p_path,
```

Source: `modules/fbx/editor/editor_scene_importer_ufbx.h` — EditorSceneFormatImporterUFBX (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FBXDocument, FBXState

Evidence:
- Code: `modules/fbx/fbx_document.cpp` — FBXDocument FBX-to-GLTF conversion paths
- Code: `modules/fbx/editor/editor_scene_importer_ufbx.h:38` — EditorSceneFormatImporterUFBX

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)

<a id="topic-fixture-contracts"></a>
## 35. Fixture contracts

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Contextual completion contracts](#topic-completion-contexts), [Diagnostic expectations](#topic-diagnostic-expectations), [Language-server semantic fixtures](#topic-lsp-semantic-fixtures).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

A source fixture provides the case; its paired stored artifact provides the expected observable result.

The suite stores GDScript source alongside expected outcome files and test assets, so behavior is expressed as a checked fixture contract.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
func test():
	var lambda_0 := func() -> void:
		print(0)
	lambda_0.call()

	var lambda_1 := func(printed: int) -> void:
		print(printed)
	lambda_1.call(1)

	var lambda_2 := func(identity: int) -> int:
		return identity
	print(lambda_2.call(2))
```

Source: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd` — test() (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd:1` — test()
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.out` — expected output

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Script resources and instances](#topic-scripting)

<a id="topic-gdscript-editor-services"></a>
## 36. GDScript editor services

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The services are built only for editor builds.

Editor tooling consumes GDScript parser output to generate documentation, color syntax, extract translations, and provide completion-related behavior.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GDScriptSyntaxHighlighter : public EditorSyntaxHighlighter {
	GDCLASS(GDScriptSyntaxHighlighter, EditorSyntaxHighlighter)

private:
	struct ColorRegion {
		enum Type {
			TYPE_NONE,
			TYPE_STRING, // `"` and `'`, optional prefix `&`, `^`, or `r`.
			TYPE_MULTILINE_STRING, // `"""` and `'''`, optional prefix `r`.
			TYPE_COMMENT, // `#` and `##`.
			TYPE_CODE_REGION, // `#region` and `#endregion`.
		};
```

Source: `modules/gdscript/editor/gdscript_highlighter.h` — GDScriptSyntaxHighlighter (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDScriptParser::Node, GDScriptTokenizer::Token

Evidence:
- Code: `modules/gdscript/editor/gdscript_docgen.cpp` — GDScriptDocGen AST traversal
- Code: `modules/gdscript/editor/gdscript_highlighter.h:35` — GDScriptSyntaxHighlighter

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Contextual completion contracts](#topic-completion-contexts)
- [Shader editing and language plugins](#topic-shader-editing-and-language-plugins)
- [Script resources and instances](#topic-scripting)

<a id="topic-gdscript-language-server"></a>
## 37. GDScript language-server support

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

LSP-shaped request and response structures are defined in godot_lsp.h and served by JSONRPC-based protocol classes.

The language server reuses parser-derived symbol data to manage documents, resolve symbols, publish diagnostics, provide links, and build workspace edits.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const LSP::DocumentSymbol *GDScriptWorkspace::resolve_symbol(const LSP::TextDocumentPositionParams &p_doc_pos, const String &p_symbol_name, bool p_func_required) {
	const LSP::DocumentSymbol *symbol = nullptr;

	String path = get_file_path(p_doc_pos.textDocument.uri);

	const ExtendGDScriptParser *parser = GDScriptLanguageProtocol::get_singleton()->get_parse_result(path);
	if (parser) {
		String symbol_name = p_symbol_name;
		if (symbol_name.get_slice_count("(") > 0) {
			symbol_name = symbol_name.get_slicec('(', 0);
		}

		LSP::Position pos = p_doc_pos.position;
```

Source: `modules/gdscript/language_server/gdscript_workspace.cpp` — GDScriptWorkspace::resolve_symbol (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

LSP::TextDocumentItem, LSP::DocumentSymbol, GDScriptWorkspace

Evidence:
- Code: `modules/gdscript/language_server/godot_lsp.h` — TextDocumentItem, DocumentSymbol, Diagnostic, and WorkspaceEdit
- Code: `modules/gdscript/language_server/gdscript_workspace.cpp:646` — GDScriptWorkspace::resolve_symbol

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Language-server semantic fixtures](#topic-lsp-semantic-fixtures)

<a id="topic-gdscript-static-analysis"></a>
## 38. GDScript static analysis

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [GDScript bytecode compilation and execution](#topic-gdscript-bytecode-runtime).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied analyzer tests exercise typed arrays and dictionaries, enum types, override signatures, casts, return types, and invalid assignments.

The analyzer consumes parsed script trees to resolve inheritance, infer and check data types, validate annotations, and diagnose invalid expressions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GDScriptAnalyzer {
	GDScriptParser *parser = nullptr;

	template <typename Fn>
	class Finally {
		Fn fn;

	public:
		Finally(Fn p_fn) :
				fn(p_fn) {}
		~Finally() {
			fn();
		}
```

Source: `modules/gdscript/gdscript_analyzer.h` — GDScriptAnalyzer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDScriptParser::Node, GDScript

Evidence:
- Code: `modules/gdscript/gdscript_analyzer.h:38` — GDScriptAnalyzer
- Code: `modules/gdscript/tests/scripts/analyzer/errors/typed_dictionary.out` — typed dictionary analyzer expectations

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Script resources and instances](#topic-scripting)

<a id="topic-image-format-loading"></a>
## 39. HDR, JPEG, and KTX loading

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Separate module adapters implement an HDR ImageFormatLoader, a libjpeg-turbo ImageFormatLoader, and a KTX ResourceFormatLoader.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ImageLoaderHDR : public ImageFormatLoader {
public:
	virtual Error load_image(Ref<Image> p_image, Ref<FileAccess> f, BitField<ImageFormatLoader::LoaderFlags> p_flags, float p_scale);
	virtual void get_recognized_extensions(List<String> *p_extensions) const;

	ImageLoaderHDR();
};
```

Source: `modules/hdr/image_loader_hdr.h` — ImageLoaderHDR : public ImageFormatLoader (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/hdr/image_loader_hdr.h:35` — ImageLoaderHDR : public ImageFormatLoader
- Code: `modules/jpg/image_loader_libjpeg_turbo.h:35` — ImageLoaderLibJPEGTurbo : public ImageFormatLoader
- Code: `modules/ktx/texture_loader_ktx.h:35` — ResourceFormatKTX : public ResourceFormatLoader

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Image decode pipelines](#topic-image-decode-pipeline)
- [JPEG decompression and coefficient transcoding](#topic-jpeg-decompression-transcoding)
- [KTX texture containers](#topic-ktx-texture-container)
- [Raster image input](#topic-raster-image-input)
- [Resource loading and caching](#topic-resource-loading)

<a id="topic-high-level-rpc"></a>
## 40. High-level RPC routing

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

RPC configuration is cached per node and packet processing distinguishes remote-call, spawn, despawn, sync, and system commands.

SceneMultiplayer routes RPC calls, peer state, connection commands, and RPC configuration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class SceneCacheInterface : public RefCounted {
	GDCLASS(SceneCacheInterface, RefCounted);

private:
	SceneMultiplayer *multiplayer = nullptr;

	//path sent caches
	struct NodeCache {
		int cache_id = 0;
		HashMap<int, int> recv_ids; // peer id, remote cache id
		HashMap<int, bool> confirmed_peers; // peer id, confirmed
	};
```

Source: `modules/multiplayer/scene_cache_interface.h` — SceneCacheInterface (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/multiplayer/scene_multiplayer.h:64` — SceneMultiplayer
- Code: `modules/multiplayer/scene_rpc_interface.h:41` — SceneRPCInterface
- Code: `modules/multiplayer/scene_cache_interface.h:38` — SceneCacheInterface
- External (official, verified): [High-level multiplayer — Godot Engine](https://docs.godotengine.org/en/latest/tutorials/networking/high_level_multiplayer.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [Godot member exposure](#topic-godot-member-exposure)
- [Packets, HTTP, JSON, and JSON-RPC](#topic-network-data-exchange)

<a id="topic-ide-messaging-protocol"></a>
## 41. IDE messaging protocol

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Peer, decoder, client, server, and CLI forwarding code share Message and MessageContent as the protocol envelope.

IDE clients and the editor exchange framed requests and responses after a handshake, with JSON bodies for typed requests.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
{
    public class Message
    {
        public MessageKind Kind { get; }
        public string Id { get; }
        public MessageContent Content { get; }

        public Message(MessageKind kind, string id, MessageContent content)
        {
            Kind = kind;
            Id = id;
            Content = content;
        }
```

Source: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Message.cs` — Message (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Message, GodotIdeMetadata

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Message.cs:3` — Message
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/MessageDecoder.cs:6` — MessageDecoder
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Peer.cs:15` — Peer
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Requests/Requests.cs:10` — Request

<a id="topic-image-codec-registration"></a>
## 42. Image codec integration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied modules cover BMP loading, DDS resource loading and saving, and several block-compression implementations.

Image codec modules provide loader, saver, compressor, and decompressor implementations against engine image and resource interfaces.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ImageLoaderBMP : public ImageFormatLoader {
protected:
	static const unsigned BITMAP_SIGNATURE = 0x4d42;

	static const unsigned BITMAP_FILE_HEADER_SIZE = 14; // bmp_file_header_s
	static const unsigned BITMAP_INFO_HEADER_MIN_SIZE = 40; // bmp_info_header_s

	enum bmp_compression_s {
		BI_RGB = 0x00,
		BI_RLE8 = 0x01, // compressed
		BI_RLE4 = 0x02, // compressed
		BI_BITFIELDS = 0x03,
		BI_JPEG = 0x04,
```

Source: `modules/bmp/image_loader_bmp.h` — ImageLoaderBMP (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Image, DDSFormatInfo

Evidence:
- Code: `modules/bmp/image_loader_bmp.h:35` — ImageLoaderBMP
- Code: `modules/dds/image_saver_dds.h` — ImageSaverDDS
- Code: `modules/astcenc/image_compress_astcenc.h` — register_image_compress_astcenc

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [ASTC block coding](#topic-astc-block-coding)
- [GPU block texture conversion](#topic-block-texture-transcoding)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-mp3-audio-resources"></a>
## 43. MP3 audio resources

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

AudioStreamMP3 and AudioStreamPlaybackMP3 form the runtime pair, while ResourceImporterMP3 handles editor import.

The MP3 module loads MP3 stream data, supplies resampled playback, and imports MP3 assets.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class AudioStreamMP3;

class AudioStreamPlaybackMP3 : public AudioStreamPlaybackResampled {
	GDCLASS(AudioStreamPlaybackMP3, AudioStreamPlaybackResampled);

	enum {
		FADE_SIZE = 256
	};
	AudioFrame loop_fade[FADE_SIZE];
	int loop_fade_remaining = FADE_SIZE;

	bool looping_override = false;
	bool looping = false;
```

Source: `modules/mp3/audio_stream_mp3.h` — AudioStreamMP3 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/mp3/audio_stream_mp3.h:37` — AudioStreamMP3
- Code: `modules/mp3/audio_stream_mp3.h:39` — AudioStreamPlaybackMP3
- Code: `modules/mp3/resource_importer_mp3.h:35` — ResourceImporterMP3
- External (official, unverified (source anchor not found)): [AudioStreamMP3 — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_audiostreammp3.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-script-registration-metadata"></a>
## 44. Managed script registration metadata

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The generator emits ScriptPathAttribute declarations and an AssemblyHasScriptsAttribute containing discovered script types.

C# script classes receive generated script-path and assembly script-type metadata for engine discovery.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
[AttributeUsage(AttributeTargets.Assembly)]
    public sealed class AssemblyHasScriptsAttribute : Attribute
    {
        /// <summary>
        /// If the Godot scripts contained in the assembly require lookup
        /// and can't rely on <see cref="ScriptTypes"/>.
        /// </summary>
        [MemberNotNullWhen(false, nameof(ScriptTypes))]
        public bool RequiresLookup { get; }

        /// <summary>
        /// The collection of types that implement a Godot script.
        /// </summary>
        public Type[]? ScriptTypes { get; }
```

Source: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/AssemblyHasScriptsAttribute.cs` — AssemblyHasScriptsAttribute (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ScriptPathAttribute, AssemblyHasScriptsAttribute

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPathAttributeGenerator.cs` — ScriptPathAttributeGenerator.Execute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/ScriptPathAttribute.cs:9` — ScriptPathAttribute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/AssemblyHasScriptsAttribute.cs:13` — AssemblyHasScriptsAttribute

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Script resources and instances](#topic-scripting)

<a id="topic-tls-crypto-backend"></a>
## 45. Mbed TLS-backed crypto and transport

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Mbed TLS module supplies Godot Crypto, certificate, TLS context, DTLS server, and TLS peer implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class DTLSServerMbedTLS : public DTLSServer {
private:
	static DTLSServer *_create_func(bool p_notify_postinitialize);
	Ref<TLSOptions> tls_options;
	Ref<CookieContextMbedTLS> cookies;

public:
	static void initialize();
	static void finalize();

	Error setup(Ref<TLSOptions> p_options) override;
	void stop() override;
	Ref<PacketPeerDTLS> take_connection(Ref<PacketPeerUDP> p_peer) override;
```

Source: `modules/mbedtls/dtls_server_mbedtls.h` — DTLSServerMbedTLS (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/mbedtls/crypto_mbedtls.h` — CryptoMbedTLS, CryptoKeyMbedTLS, X509CertificateMbedTLS, and HMACContextMbedTLS
- Code: `modules/mbedtls/dtls_server_mbedtls.h:37` — DTLSServerMbedTLS
- Code: `modules/mbedtls/stream_peer_mbedtls.h:37` — StreamPeerMbedTLS

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Cryptographic keys, hashing, and TLS support](#topic-cryptography)
- [Godot ENet socket adaptation](#topic-godot-enet-socket-adaptation)
- [TLS connection and session state](#topic-tls-connection-state)
- [Cryptographic resources and contexts](#topic-crypto-resources)
- [Networking and transport I/O](#topic-network-io)
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)
- [Byte streams and socket servers](#topic-stream-networking)
- [X.509 certificate processing](#topic-x509-certificate-processing)

<a id="topic-native-platform-adapters"></a>
## 46. Native platform adapters

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Apple embedded hosting](#topic-apple-embedded-hosting), [Audio backend adapters](#topic-audio-backend-adapters), [MIDI input adapters](#topic-midi-input-adapters).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Unix and Windows have parallel implementations for several services, while Apple has additional embedded implementations.

Platform adapters implement filesystem, sockets, IP resolution, operating-system services, and threads behind engine interfaces selected by the driver build.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class FileAccessUnix : public FileAccess {
	GDSOFTCLASS(FileAccessUnix, FileAccess);
	FILE *f = nullptr;
	int flags = 0;
	void check_errors(bool p_write = false) const;
	mutable Error last_error = OK;
	String save_path;
	String path;
	String path_src;

	void _close();

#if defined(TOOLS_ENABLED)
```

Source: `drivers/unix/file_access_unix.h` — class FileAccessUnix : public FileAccess (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `drivers/unix/file_access_unix.h:39` — class FileAccessUnix : public FileAccess
- Code: `drivers/unix/net_socket_unix.h:40` — class NetSocketUnix : public NetSocket
- Code: `drivers/windows/net_socket_winsock.h:40` — class NetSocketWinSock : public NetSocket

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)

<a id="topic-navigation-map-composition"></a>
## 47. Navigation map composition

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Navigation avoidance](#topic-navigation-avoidance), [Navigation path queries](#topic-navigation-path-queries).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Parallel 2D and 3D implementations expose map-owned region, link, agent, and obstacle collections and iteration snapshots.

Navigation maps collect regions, links, agents, and obstacles and build per-frame read iterations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NavRegionIteration2D;
struct NavMapIteration2D;

struct NavMapIterationBuild2D {
	Vector2 merge_rasterizer_cell_size;
	bool use_edge_connections = true;
	real_t edge_connection_margin;
	real_t link_connection_radius;
	Nav2D::PerformanceData performance_data;
	int polygon_count = 0;
	int free_edge_count = 0;

	HashMap<Nav2D::EdgeKey, Nav2D::EdgeConnectionPair, Nav2D::EdgeKey> iter_connection_pairs_map;
	LocalVector<Nav2D::Connection> iter_free_edges;
```

Source: `modules/navigation_2d/2d/nav_map_iteration_2d.h` — NavMapIteration2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/navigation_2d/nav_map_2d.h:51` — NavMap2D
- Code: `modules/navigation_2d/2d/nav_map_iteration_2d.h:43` — NavMapIteration2D
- Code: `modules/navigation_3d/nav_map_3d.h:53` — NavMap3D
- Code: `modules/navigation_3d/3d/nav_map_iteration_3d.h:43` — NavMapIteration3D
- External (official, unverified (source anchor not found)): [2D navigation overview — Godot Engine](https://docs.godotengine.org/en/latest/tutorials/navigation/navigation_introduction_2d.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Navigation agents and regions](#topic-navigation-agents)
- [Navigation region construction](#topic-navigation-regions)

<a id="topic-navigation-mesh-construction"></a>
## 48. Navigation mesh construction

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

2D construction uses polygon geometry and 3D construction uses Recast-backed mesh generation when configured.

The 2D and 3D generators turn source geometry into navigation mesh data and connect region or link geometry into map iterations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NavMeshGenerator2D : public Object {
	GDSOFTCLASS(NavMeshGenerator2D, Object);

	static NavMeshGenerator2D *singleton;

	static Mutex baking_navmesh_mutex;
	static Mutex generator_task_mutex;

	static RWLock generator_parsers_rwlock;
	static LocalVector<NavMeshGeometryParser2D *> generator_parsers;

	static bool use_threads;
	static bool baking_use_multiple_threads;
```

Source: `modules/navigation_2d/2d/nav_mesh_generator_2d.h` — NavMeshGenerator2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/navigation_2d/2d/nav_mesh_generator_2d.h:43` — NavMeshGenerator2D
- Code: `modules/navigation_2d/2d/nav_region_builder_2d.h:37` — NavRegionBuilder2D
- Code: `modules/navigation_3d/3d/nav_mesh_generator_3d.h:41` — NavMeshGenerator3D
- Code: `modules/navigation_3d/3d/nav_region_builder_3d.h:37` — NavRegionBuilder3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry resources](#topic-geometry-resources)
- [Navigation agents and regions](#topic-navigation-agents)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)

<a id="topic-objectdb-snapshots"></a>
## 49. ObjectDB snapshots

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The runtime collector transports snapshot data while editor classes turn it into inspectable data and controls.

The debug-only object database profiler collects serialized object snapshots and renders summary, class, node, object, and ref-counted views.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class SnapshotView : public Control {
	GDCLASS(SnapshotView, Control);

protected:
	GameStateSnapshot *snapshot_data = nullptr;
	GameStateSnapshot *diff_data = nullptr;

	Vector<TreeItem *> _get_children_recursive(Tree *p_tree);

public:
	String view_name;

	virtual void show_snapshot(GameStateSnapshot *p_data, GameStateSnapshot *p_diff_data = nullptr);
```

Source: `modules/objectdb_profiler/editor/data_viewers/snapshot_view.h` — SnapshotView (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/objectdb_profiler/snapshot_collector.h:44` — SnapshotCollector
- Code: `modules/objectdb_profiler/editor/snapshot_data.h:35` — GameStateSnapshot
- Code: `modules/objectdb_profiler/editor/data_viewers/snapshot_view.h:40` — SnapshotView
- Code: `modules/objectdb_profiler/editor/objectdb_profiler_panel.h:43` — ObjectDBProfilerPanel

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Engine object model](#topic-object-model)
- [Scene graph and lifecycle](#topic-scene-tree)

<a id="topic-ogg-vorbis-audio-streams"></a>
## 50. Ogg Vorbis audio streams

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module also provides an editor resource importer.

AudioStreamOggVorbis and its playback class provide an audio-stream resource and resampled playback implementation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class AudioStreamOggVorbis;

class AudioStreamPlaybackOggVorbis : public AudioStreamPlaybackResampled {
	GDCLASS(AudioStreamPlaybackOggVorbis, AudioStreamPlaybackResampled);

	uint32_t frames_mixed = 0;
	bool active = false;
	bool looping_override = false;
	bool looping = false;
	int loops = 0;

	enum {
		FADE_SIZE = 256
```

Source: `modules/vorbis/audio_stream_ogg_vorbis.h` — AudioStreamOggVorbis (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

AudioStreamOggVorbis, AudioStreamPlaybackOggVorbis

Evidence:
- Code: `modules/vorbis/audio_stream_ogg_vorbis.h:40` — AudioStreamOggVorbis
- Code: `modules/vorbis/audio_stream_ogg_vorbis.h:42` — AudioStreamPlaybackOggVorbis
- Code: `modules/vorbis/resource_importer_ogg_vorbis.h:37` — ResourceImporterOggVorbis

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-ogg-packet-sequences"></a>
## 51. Ogg packet sequences

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module defines a Resource plus a ref-counted playback companion.

OggPacketSequence persists packet data, page granule positions, and sample-rate information for sequence playback.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class OggPacketSequencePlayback;

class OggPacketSequence : public Resource {
	GDCLASS(OggPacketSequence, Resource);

	friend class OggPacketSequencePlayback;

	// List of pages, each of which is a list of packets on that page. The innermost PackedByteArrays contain complete ogg packets.
	Vector<Vector<PackedByteArray>> page_data;

	// List of the granule position for each page.
	Vector<uint64_t> page_granule_positions;
```

Source: `modules/ogg/ogg_packet_sequence.h` — OggPacketSequence (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OggPacketSequence

Evidence:
- Code: `modules/ogg/ogg_packet_sequence.h:41` — OggPacketSequence
- Code: `modules/ogg/ogg_packet_sequence.h:39` — OggPacketSequencePlayback
- External (official, unverified (source anchor not found)): [OggPacketSequence — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_oggpacketsequence.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Theora block video coding](#topic-theora-block-video-codec)
- [Vorbis block analysis and synthesis](#topic-vorbis-block-synthesis)
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-openxr-action-configuration"></a>
## 52. OpenXR action configuration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The action-map subtree uses Resource-derived configuration objects for the OpenXR input graph.

OpenXRActionMap persists action sets, actions, interaction profiles, bindings, modifiers, and haptic feedback configuration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class OpenXRActionSet;

class OpenXRAction : public Resource {
	GDCLASS(OpenXRAction, Resource);

public:
	enum ActionType {
		OPENXR_ACTION_BOOL,
		OPENXR_ACTION_FLOAT,
		OPENXR_ACTION_VECTOR2,
		OPENXR_ACTION_POSE,
		OPENXR_ACTION_HAPTIC,
		OPENXR_ACTION_MAX
```

Source: `modules/openxr/action_map/openxr_action.h` — OpenXRAction (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OpenXRActionMap

Evidence:
- Code: `modules/openxr/action_map/openxr_action_map.h:39` — OpenXRActionMap
- Code: `modules/openxr/action_map/openxr_action.h:37` — OpenXRAction
- Code: `modules/openxr/action_map/openxr_binding_modifier.h:41` — OpenXRBindingModifier
- Code: `modules/openxr/action_map/openxr_haptic_feedback.h:37` — OpenXRHapticBase
- External (official, unverified (source anchor not found)): [OpenXRActionMap — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_openxractionmap.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)

<a id="topic-openxr-runtime-integration"></a>
## 53. OpenXR runtime integration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [OpenXR composition layers](#topic-openxr-composition-layers), [OpenXR extension wrappers](#topic-openxr-extension-wrappers), [OpenXR render models](#topic-openxr-render-models).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the integration boundary; later OpenXR topics specialize its action, extension, spatial, and rendering behavior.

The module places OpenXRInterface over OpenXRAPI so the engine can interact with an OpenXR runtime.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class OpenXRAPI {
public:
	class OpenXRSwapChainInfo {
	private:
		XrSwapchain swapchain = XR_NULL_HANDLE;
		void *swapchain_graphics_data = nullptr;
		uint32_t image_index = 0;
		bool image_acquired = false;
		bool skip_acquire_swapchain = false;

		static Vector<OpenXRSwapChainInfo> free_queue;

	public:
```

Source: `modules/openxr/openxr_api.h` — class OpenXRAPI (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OpenXRActionMap, OpenXRFutureResult

Evidence:
- Code: `modules/openxr/openxr_interface.h:66` — class OpenXRInterface : public XRInterface
- Code: `modules/openxr/openxr_api.h:56` — class OpenXRAPI
- External (official, unverified (source anchor not found)): [The OpenXR 1.1 Specification](https://registry.khronos.org/OpenXR/specs/1.1/html/xrspec.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Binding modifiers](#topic-openxr-binding-modifiers)
- [OpenXR structure chains](#topic-openxr-structure-chains)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)

<a id="topic-png-image-codec"></a>
## 54. PNG image codec

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

PNG support is isolated as a driver with loader and saver classes.

The PNG driver supplies image loading and resource saving implementations and can build bundled libpng sources.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ImageLoaderPNG : public ImageFormatLoader {
private:
	static Vector<uint8_t> lossless_pack_png(const Ref<Image> &p_image);
	static Ref<Image> lossless_unpack_png(const Vector<uint8_t> &p_data);
	static Ref<Image> unpack_mem_png(const uint8_t *p_png, int p_size);
	static Ref<Image> load_mem_png(const uint8_t *p_png, int p_size);

public:
	virtual Error load_image(Ref<Image> p_image, Ref<FileAccess> f, BitField<ImageFormatLoader::LoaderFlags> p_flags, float p_scale);
	virtual void get_recognized_extensions(List<String> *p_extensions) const;
	ImageLoaderPNG();
};
```

Source: `drivers/png/image_loader_png.h` — class ImageLoaderPNG : public ImageFormatLoader (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `drivers/png/image_loader_png.h:35` — class ImageLoaderPNG : public ImageFormatLoader
- Code: `drivers/png/resource_saver_png.h:36` — class ResourceSaverPNG : public ResourceFormatSaver
- Code: `drivers/png/SCsub` — builtin_libpng source selection

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [PNG streaming I/O and row transforms](#topic-png-stream-transforms)
- [Raster image input](#topic-raster-image-input)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-procedural-noise-textures"></a>
## 55. Procedural noise textures

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module supplies an abstract Noise resource, FastNoiseLite implementation, texture resources, editor preview support, and tests.

Noise generators produce values and image data that NoiseTexture2D and NoiseTexture3D turn into textures.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class FastNoiseLite : public Noise {
	GDCLASS(FastNoiseLite, Noise);
	OBJ_SAVE_TYPE(FastNoiseLite);

public:
	enum NoiseType {
		TYPE_SIMPLEX = _FastNoiseLite::NoiseType_OpenSimplex2,
		TYPE_SIMPLEX_SMOOTH = _FastNoiseLite::NoiseType_OpenSimplex2S,
		TYPE_CELLULAR = _FastNoiseLite::NoiseType_Cellular,
		TYPE_PERLIN = _FastNoiseLite::NoiseType_Perlin,
		TYPE_VALUE_CUBIC = _FastNoiseLite::NoiseType_ValueCubic,
		TYPE_VALUE = _FastNoiseLite::NoiseType_Value,
	};
```

Source: `modules/noise/fastnoise_lite.h` — FastNoiseLite (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/noise/noise.h:36` — Noise
- Code: `modules/noise/fastnoise_lite.h:35` — FastNoiseLite
- Code: `modules/noise/noise_texture_2d.h:40` — NoiseTexture2D
- Code: `modules/noise/noise_texture_3d.h:40` — NoiseTexture3D
- External (official, unverified (source anchor not found)): [Noise — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_noise.html), accessed 2026-07-15

<a id="topic-raycast-occlusion-culling"></a>
## 56. Raycast occlusion culling

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Embree-backed lightmap and static raycaster classes supplement the central occlusion-cull implementation.

The raycast module represents occluders, instances, scenarios, and raycast thread data to provide renderer-scene occlusion culling.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class LightmapRaycasterEmbree : public LightmapRaycaster {
	GDCLASS(LightmapRaycasterEmbree, LightmapRaycaster);

private:
	struct AlphaTextureData {
		Vector<uint8_t> data;
		Vector2i size;

		uint8_t sample(float u, float v) const;
	};

	RTCDevice embree_device;
	RTCScene embree_scene;
```

Source: `modules/raycast/lightmap_raycaster_embree.h` — class LightmapRaycasterEmbree : public LightmapRaycaster (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RaycastOccluder

Evidence:
- Code: `modules/raycast/raycast_occlusion_cull.h` — RaycastOcclusionCull; Occluder, OccluderInstance, Scenario, RaycastThreadData
- Code: `modules/raycast/static_raycaster_embree.h:39` — class StaticRaycasterEmbree : public StaticRaycaster
- Code: `modules/raycast/lightmap_raycaster_embree.h:41` — class LightmapRaycasterEmbree : public LightmapRaycaster

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Primitive intersection](#topic-primitive-intersection)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Task scheduling and synchronization](#topic-task-scheduling)

<a id="topic-regular-expression-results"></a>
## 57. Regular-expression matching

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

RegExMatch is the cross-call result object for search operations rather than a local temporary.

The regex module exposes compiled regular expressions and RefCounted match results with ranges.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
int RegEx::_sub(const String &p_subject, const String &p_replacement, int p_offset, int p_end, uint32_t p_flags, String &r_output) const {
	// `safety_zone` is the number of chars we allocate in addition to the number of chars expected in order to
	// guard against the PCRE API writing one additional `\0` at the end. PCRE's API docs are unclear on whether
	// PCRE understands outlength in `pcre2_substitute(`) as counting an implicit additional terminating char or
	// not. Always allocating one char more than telling PCRE has us on the safe side.
	const int safety_zone = 1;

	PCRE2_SIZE olength = p_subject.length() + 1; // Space for output string and one terminating `\0` character.
	Vector<char32_t> output;
	output.resize(olength + safety_zone);

	PCRE2_SIZE length = p_subject.length();
	if (p_end >= 0 && (uint32_t)p_end < length) {
```

Source: `modules/regex/regex.cpp` — RegEx::_sub (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RegExMatch

Evidence:
- Code: `modules/regex/regex.h` — class RegExMatch; struct Range; class RegEx
- Code: `modules/regex/regex.cpp:293` — RegEx::_sub

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Engine object model](#topic-object-model)

<a id="topic-rendering-backends"></a>
## 58. Rendering backend abstraction

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The source partition contains separate backend implementations rather than one graphics API implementation.

Rendering backends specialize common context and device-driver abstractions for Vulkan, GLES3, Direct3D 12, and Metal.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// - D3D12 structs are zero-initialized and fields not requiring a non-zero value are omitted (except in cases where expresivity reasons apply).
class RenderingDeviceDriverD3D12 : public RenderingDeviceDriver {
	/*****************/
	/**** GENERIC ****/
	/*****************/

	struct D3D12Format {
		DXGI_FORMAT family = DXGI_FORMAT_UNKNOWN;
		DXGI_FORMAT general_format = DXGI_FORMAT_UNKNOWN;
		UINT swizzle = D3D12_DEFAULT_SHADER_4_COMPONENT_MAPPING;
		DXGI_FORMAT dsv_format = DXGI_FORMAT_UNKNOWN;
	};

	static const D3D12Format RD_TO_D3D12_FORMAT[RDD::DATA_FORMAT_MAX];
```

Source: `drivers/d3d12/rendering_device_driver_d3d12.h` — class RenderingDeviceDriverD3D12 : public RenderingDeviceDriver (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RenderingDeviceDriverVulkan::BufferInfo, RenderingDeviceDriverVulkan::CommandBufferInfo

Evidence:
- Code: `drivers/vulkan/rendering_device_driver_vulkan.h:54` — class RenderingDeviceDriverVulkan : public RenderingDeviceDriver
- Code: `drivers/d3d12/rendering_device_driver_d3d12.h:61` — class RenderingDeviceDriverD3D12 : public RenderingDeviceDriver
- Code: `drivers/metal/rendering_device_driver_metal.h` — class RenderingDeviceDriverMetal : public RenderingDeviceDriver
- Code: `drivers/gles3/rasterizer_gles3.h:53` — class RasterizerGLES3 : public RendererCompositor

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Metal-cpp object bridge](#topic-metal-cpp-object-bridge)
- [Android rendering and input](#topic-android-render-input)
- [Compile-time platform backends](#topic-sdl-platform-backends)

<a id="topic-scene-replication-configuration"></a>
## 59. Scene replication configuration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Replicated property synchronization](#topic-replicated-property-synchronization), [Replicated spawning](#topic-replicated-spawning).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Each configured property stores a spawn flag and a replication mode, and the configuration maintains derived lists for replication phases.

A saved SceneReplicationConfig separates NodePath properties into spawn, sync, and watch lists.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const List<NodePath> &SceneReplicationConfig::get_spawn_properties() {
	if (dirty) {
		_update();
	}
	return spawn_props;
}

const List<NodePath> &SceneReplicationConfig::get_sync_properties() {
	if (dirty) {
		_update();
	}
	return sync_props;
}
```

Source: `modules/multiplayer/scene_replication_config.cpp` — SceneReplicationConfig::get_spawn_properties (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SceneReplicationConfig, ReplicationProperty

Evidence:
- Code: `modules/multiplayer/scene_replication_config.h:36` — SceneReplicationConfig
- Code: `modules/multiplayer/scene_replication_config.cpp:252` — SceneReplicationConfig::get_spawn_properties
- External (official, unverified (source anchor not found)): [SceneReplicationConfig — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_scenereplicationconfig.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [Strings, interned names, and node paths](#topic-string-names-paths)

<a id="topic-signal-awaitability"></a>
## 60. Signal awaitability

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SignalAwaiter bridges signal completion back to managed await continuations.

A Signal combines an owner and a signal name and exposes a custom awaiter for asynchronous notification.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
{
    public class SignalAwaiter : IAwaiter<Variant[]>, IAwaitable<Variant[]>
    {
        private bool _completed;
        private Variant[] _result;
        private Action _continuation;

        public SignalAwaiter(GodotObject source, StringName signal, GodotObject target)
        {
            var awaiterGcHandle = CustomGCHandle.AllocStrong(this);
            using godot_string_name signalSrc = NativeFuncs.godotsharp_string_name_new_copy(
                (godot_string_name)(signal?.NativeValue ?? default));
            NativeFuncs.godotsharp_internal_signal_awaiter_connect(GodotObject.GetPtr(source), in signalSrc,
                GodotObject.GetPtr(target), GCHandle.ToIntPtr(awaiterGcHandle));
```

Source: `modules/mono/glue/GodotSharp/GodotSharp/Core/SignalAwaiter.cs` — SignalAwaiter (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant, ManagedCallbacks

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Signal.cs` — Signal.GetAwaiter
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/SignalAwaiter.cs:7` — SignalAwaiter
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs` — ManagedCallbacks.SignalAwaiter_SignalCallback

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic invocation and signals](#topic-dynamic-invocation-signals)
- [Android instrumented integration tests](#topic-android-instrumented-tests)
- [Android plugins and signals](#topic-android-plugin-signals)
- [Godot member exposure](#topic-godot-member-exposure)
- [managed C# script bridge and source generation](#topic-managed-csharp-scripting)
- [Class metadata and reflection](#topic-reflection)
- [Reflection metadata](#topic-reflection-metadata)
- [Runtime class metadata](#topic-runtime-metadata)
- [Scene tree and signal connections](#topic-scene-tree-and-signal-connections)

<a id="topic-theora-video-streams"></a>
## 61. Theora video streams

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The stream, loader, playback, and writer types separate runtime playback from editor export behavior.

The Theora module defines a video-stream resource, playback implementation, resource loader, and OGV movie writer.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class MovieWriterOGV : public MovieWriter {
	GDCLASS(MovieWriterOGV, MovieWriter)

	uint32_t mix_rate = 48000;
	AudioServer::SpeakerMode speaker_mode = AudioServer::SPEAKER_MODE_STEREO;
	String base_path;
	uint32_t frame_count = 0;
	uint32_t fps = 0;
	uint32_t audio_ch = 0;
	uint32_t audio_frames = 0;

	Ref<FileAccess> f;
```

Source: `modules/theora/editor/movie_writer_ogv.h` — class MovieWriterOGV : public MovieWriter (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VideoStreamTheora

Evidence:
- Code: `modules/theora/video_stream_theora.h` — VideoStreamPlaybackTheora, VideoStreamTheora, ResourceFormatLoaderTheora
- Code: `modules/theora/editor/movie_writer_ogv.h:40` — class MovieWriterOGV : public MovieWriter

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)
- [Resources and asset lifecycle](#topic-resources)
- [Theora block video coding](#topic-theora-block-video-codec)

<a id="topic-upnp-device-control"></a>
## 62. UPnP device control

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

UPNPDevice is the public device object; UPNPMiniUPNP and UPNPDeviceMiniUPNP provide the third-party-backed forms.

The UPnP module represents generic devices and MiniUPnP-backed specializations behind RefCounted APIs.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class UPNP : public RefCounted {
	GDCLASS(UPNP, RefCounted);

protected:
	static void _bind_methods();

	static UPNP *(*_create)(bool p_notify_postinitialize);

public:
	enum UPNPResult {
		UPNP_RESULT_SUCCESS,
		UPNP_RESULT_NOT_AUTHORIZED,
		UPNP_RESULT_PORT_MAPPING_NOT_FOUND,
```

Source: `modules/upnp/upnp.h` — class UPNP : public RefCounted (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UPNPDevice

Evidence:
- Code: `modules/upnp/upnp.h:37` — class UPNP : public RefCounted
- Code: `modules/upnp/upnp_device.h:36` — class UPNPDevice : public RefCounted
- Code: `modules/upnp/upnp_miniupnp.h:39` — class UPNPMiniUPNP : public UPNP

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Engine object model](#topic-object-model)
- [UPnP device discovery](#topic-upnp-device-discovery)

<a id="topic-visual-shader-module-build"></a>
## 63. Visual Shader source partition

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This inventory establishes the module's source partition without relying on unlisted implementation bodies.

The Visual Shader build script compiles the module's C++ sources, its node sources, and editor sources when building the editor.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
env_visual_shader.add_source_files(env.modules_sources, "*.cpp")
env_visual_shader.add_source_files(env.modules_sources, "vs_nodes/*.cpp")

if env.editor_build:
    env_visual_shader.add_source_files(env.modules_sources, "editor/*.cpp")
```

Source: `modules/visual_shader/SCsub` — env_visual_shader.add_source_files (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/visual_shader/SCsub:11` — env_visual_shader.add_source_files

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Scene graph and lifecycle](#topic-scene-tree)
- [Script resources and instances](#topic-scripting)
- [Visual shader vector operations](#topic-visual-shader-vector-operations)

<a id="topic-visual-shader-nodes"></a>
## 64. Visual shader nodes

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Visual shader vector operations](#topic-visual-shader-vector-operations).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module defines input, output, parameter, expression, varying, math, texture, and particle node families.

Node resources provide reusable shader operations within typed graphs of node records and connection records.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class VisualShaderNodeParameter;
class VisualShaderNode;

class VisualShader : public Shader {
	GDCLASS(VisualShader, Shader);

public:
	enum Type {
		TYPE_VERTEX,
		TYPE_FRAGMENT,
		TYPE_LIGHT,
		TYPE_START,
		TYPE_PROCESS,
```

Source: `modules/visual_shader/visual_shader.h` — VisualShaderNode (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VisualShaderNode

Evidence:
- Code: `modules/visual_shader/visual_shader.h:40` — VisualShaderNode
- Code: `modules/visual_shader/vs_nodes/visual_shader_nodes.h:911` — VisualShaderNodeVectorOp
- Code: `modules/visual_shader/vs_nodes/visual_shader_particle_nodes.h:39` — VisualShaderNodeParticleEmitter

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Resource-backed assets](#topic-resource-assets)
- [Scene graph and lifecycle](#topic-scene-tree)

<a id="topic-webp-image-io"></a>
## 65. WebP image I/O

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

It conditionally compiles bundled libwebp sources through its module build script.

The WebP module implements image loading, saving, and shared buffer helpers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ImageLoaderWebP : public ImageFormatLoader {
public:
	virtual Error load_image(Ref<Image> p_image, Ref<FileAccess> f, BitField<ImageFormatLoader::LoaderFlags> p_flags, float p_scale);
	virtual void get_recognized_extensions(List<String> *p_extensions) const;
	ImageLoaderWebP();
};
```

Source: `modules/webp/image_loader_webp.h` — ImageLoaderWebP (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ImageLoaderWebP, ResourceSaverWebP

Evidence:
- Code: `modules/webp/image_loader_webp.h:35` — ImageLoaderWebP
- Code: `modules/webp/resource_saver_webp.h:36` — ResourceSaverWebP
- Code: `modules/webp/webp_common.h:2` — webp_common

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry buffer storage](#topic-buffer-storage)
- [Image decode pipelines](#topic-image-decode-pipeline)
- [WebP chunk parsing, incremental decode, and animation](#topic-webp-riff-decoding)

<a id="topic-webrtc-peer-connections"></a>
## 66. WebRTC peer connections

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [WebRTC data channels](#topic-webrtc-data-channels), [WebRTC multiplayer mesh](#topic-webrtc-multiplayer-mesh).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The core class has extension and JavaScript-backed implementations.

A WebRTC peer connection exposes connection state, negotiation, ICE candidates, and data-channel creation to Godot.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class WebRTCPeerConnection : public RefCounted {
	GDCLASS(WebRTCPeerConnection, RefCounted);

public:
	enum ConnectionState {
		STATE_NEW,
		STATE_CONNECTING,
		STATE_CONNECTED,
		STATE_DISCONNECTED,
		STATE_FAILED,
		STATE_CLOSED
	};
```

Source: `modules/webrtc/webrtc_peer_connection.h` — WebRTCPeerConnection (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebRTCPeerConnection

Evidence:
- Code: `modules/webrtc/webrtc_peer_connection.h:35` — WebRTCPeerConnection
- Code: `modules/webrtc/webrtc_peer_connection_extension.h:37` — WebRTCPeerConnectionExtension
- External (primary, verified): [WebRTC: Real-Time Communication in Browsers](https://www.w3.org/TR/webrtc/), accessed 2026-07-16

<a id="topic-websocket-peers"></a>
## 67. WebSocket peers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [WebSocket multiplayer](#topic-websocket-multiplayer).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The browser build includes a JavaScript socket bridge, while native builds include a WSLay-based peer.

WebSocketPeer defines packet-oriented socket behavior and has native and browser-backed implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EMWSPeer : public WebSocketPeer {
private:
	int peer_sock = -1;

	State ready_state = STATE_CLOSED;
	Vector<uint8_t> packet_buffer;
	PacketBuffer<uint8_t> in_buffer;
	uint8_t was_string = 0;
	int close_code = -1;
	String close_reason;
	String selected_protocol;
	String requested_url;
```

Source: `modules/websocket/emws_peer.h` — EMWSPeer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebSocketPeer

Evidence:
- Code: `modules/websocket/websocket_peer.h:36` — WebSocketPeer
- Code: `modules/websocket/wsl_peer.h:45` — WSLPeer
- Code: `modules/websocket/emws_peer.h:56` — EMWSPeer
- External (primary, verified): [RFC 6455: The WebSocket Protocol](https://www.rfc-editor.org/rfc/rfc6455), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-webxr-session-bridge"></a>
## 68. WebXR session bridge

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The JavaScript bridge manages sessions, reference spaces, frame data, views, and input-source data for the C++ implementation.

WebXR interface state receives browser session callbacks and input sources through web-platform bindings.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```javascript
const GodotWebXR = {
	$GodotWebXR__deps: ['$MainLoop', '$GL', '$GodotRuntime', '$runtimeKeepalivePush', '$runtimeKeepalivePop'],
	$GodotWebXR: {
		gl: null,

		session: null,
		gl_binding: null,
		layer: null,
		space: null,
		frame: null,
		pose: null,
		view_count: 1,
		input_sources: new Array(16),
```

Source: `modules/webxr/native/library_godot_webxr.js` — GodotWebXR (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebXRInterfaceJS, WebXRInputSource

Evidence:
- Code: `modules/webxr/webxr_interface.h:40` — WebXRInterface
- Code: `modules/webxr/webxr_interface_js.h` — WebXRInterfaceJS::InputSource
- Code: `modules/webxr/native/library_godot_webxr.js:31` — GodotWebXR
- External (primary, verified): [WebXR Device API](https://www.w3.org/TR/webxr/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Shader interface metadata](#topic-shader-interface-metadata)

<a id="topic-zip-archive-io"></a>
## 69. ZIP archive I/O

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module includes tests for compressed and uncompressed archives.

ZIPReader reads archives and ZIPPacker creates archives through reference-counted API objects.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ZIPPacker : public RefCounted {
	GDCLASS(ZIPPacker, RefCounted);

	Ref<FileAccess> fa;
	zipFile zf = nullptr;
	int compression_level = Z_DEFAULT_COMPRESSION;
	HashSet<String> directories;

protected:
	static void _bind_methods();

#ifndef DISABLE_DEPRECATED
	Error _start_file_bind_compat_115946(const String &p_path);
```

Source: `modules/zip/zip_packer.h` — ZIPPacker (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ZIPReader, ZIPPacker

Evidence:
- Code: `modules/zip/zip_reader.h:38` — ZIPReader
- Code: `modules/zip/zip_packer.h:38` — ZIPPacker
- Code: `modules/zip/tests/test_zip.h` — ZipTest

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [Target-platform export](#topic-target-platform-export)
- [ZIP64 archive I/O](#topic-zip64-archive-io)

<a id="topic-gltf-materials-textures"></a>
## 70. glTF material and texture conversion

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

GLTFDocument maps glTF material dictionaries, PBR metallic-roughness values, textures, samplers, UV coordinates, color transforms, alpha mode, and texture-transform extensions to Godot material resources.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GLTFTexture : public Resource {
	GDCLASS(GLTFTexture, Resource);

private:
	GLTFImageIndex src_image = -1;
	GLTFTextureSamplerIndex sampler = -1;

protected:
	static void _bind_methods();

public:
	GLTFImageIndex get_src_image() const;
	void set_src_image(GLTFImageIndex val);
```

Source: `modules/gltf/structures/gltf_texture.h` — GLTFTexture (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GLTFState

Evidence:
- Code: `modules/gltf/gltf_document.cpp` — material parsing and KHR_texture_transform handling
- Code: `modules/gltf/structures/gltf_texture.h:37` — GLTFTexture

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Alpha-plane coding](#topic-alpha-plane-coding)
- [FBX scene import through GLTF structures](#topic-fbx-gltf-scene-import)
- [Basis file layout](#topic-basis-container-layout)
- [Basis texture transcoding](#topic-basis-transcoding)
- [GPU block texture conversion](#topic-block-texture-transcoding)
- [KTX texture containers](#topic-ktx-texture-container)
- [KTX2 container transcoding](#topic-ktx2-container-transcoding)
- [Raster image input](#topic-raster-image-input)
- [Textures, meshes, and materials](#topic-rendering-resources)
- [Standalone resource importing](#topic-resource-importing)
- [Texture resources and fallback placeholders](#topic-textures-and-placeholders)
- [Instancing](#topic-instancing)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Resource-backed assets](#topic-resource-assets)

<a id="topic-gltf-node-hierarchy"></a>
## 71. glTF scene-node hierarchy

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

GLTFDocument reconstructs parent and child relationships among indexed GLTF node records and attaches meshes, cameras, lights, skins, and skeleton membership through node indexes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GLTFNode : public Resource {
	GDCLASS(GLTFNode, Resource);
	friend class GLTFDocument;
	friend class SkinTool;
	friend class FBXDocument;

private:
	String original_name;
	GLTFNodeIndex parent = -1;
	int height = -1;
	Transform3D transform;
	GLTFMeshIndex mesh = -1;
	GLTFCameraIndex camera = -1;
```

Source: `modules/gltf/structures/gltf_node.h` — GLTFNode (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GLTFNode, GLTFState

Evidence:
- Code: `modules/gltf/structures/gltf_node.h:37` — GLTFNode
- Code: `modules/gltf/gltf_document.cpp` — scene and node parsing paths

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [FBX scene import through GLTF structures](#topic-fbx-gltf-scene-import)
- [Canvas and viewport presentation](#topic-canvas-and-viewports)
- [COLLADA scene interchange](#topic-collada-import)
- [Expression parsing and evaluation](#topic-expression-evaluation)
- [HTTP and multiplayer](#topic-networking)
- [ObjectDB snapshots](#topic-objectdb-snapshots)
- [Packed scene serialization](#topic-packed-scenes)
- [Post-import processing](#topic-post-import-processing)
- [Scene graph nodes](#topic-scene-graph)
- [Scene importing](#topic-scene-importing)
- [Scene graph and lifecycle](#topic-scene-tree)
- [SceneTree execution](#topic-scene-tree-execution)
- [3D physics nodes](#topic-three-dimensional-physics)
- [2D physics nodes](#topic-two-dimensional-physics)
- [Visual Shader source partition](#topic-visual-shader-module-build)
- [Visual shader nodes](#topic-visual-shader-nodes)
- [Skeletons, animation, and ragdolls](#topic-skeletal-ragdoll)

<a id="topic-interactive-music-transitions"></a>
## 72. interactive music transition tables

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

AudioStreamInteractive combines clips with a TransitionKey-indexed transition map so playback can select a rule for a clip-to-clip change.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
for (int i = 0; i < keys.size(); i++) {
		const Transition &tr = transition_map[TransitionKey(keys[i].x, keys[i].y)];
		Dictionary data;
		data["from_time"] = tr.from_time;
		data["to_time"] = tr.to_time;
		data["fade_mode"] = tr.fade_mode;
		data["fade_beats"] = tr.fade_beats;
		if (tr.use_filler_clip) {
			data["use_filler_clip"] = true;
			data["filler_clip"] = tr.filler_clip;
		}
		if (tr.hold_previous) {
			data["hold_previous"] = true;
		}
```

Source: `modules/interactive_music/audio_stream_interactive.cpp` — transition_map[TransitionKey(keys[i].x, keys[i].y)] (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

AudioStreamInteractive, AudioStreamInteractive Transition

Evidence:
- Code: `modules/interactive_music/audio_stream_interactive.h` — AudioStreamInteractive::Clip, Transition, TransitionKey, and TransitionKeyHasher
- Code: `modules/interactive_music/audio_stream_interactive.cpp:202` — transition_map[TransitionKey(keys[i].x, keys[i].y)]

<a id="topic-mobile-vr-interface"></a>
## 73. mobile XR interface

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

MobileVRInterface is the module's concrete XRInterface implementation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class MobileVRInterface : public XRInterface {
	GDCLASS(MobileVRInterface, XRInterface);

private:
	bool initialized = false;
	XRInterface::TrackingStatus tracking_state;
	XRPose::TrackingConfidence tracking_confidence = XRPose::XR_TRACKING_CONFIDENCE_NONE;

	// Just set some defaults for these. At some point we need to look at adding a lookup table for common device + headset combos and/or support reading cardboard QR codes
	double eye_height = 1.85;
	uint64_t last_ticks = 0;

	double intraocular_dist = 6.0;
```

Source: `modules/mobile_vr/mobile_vr_interface.h` — MobileVRInterface : public XRInterface (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/mobile_vr/mobile_vr_interface.h:48` — MobileVRInterface : public XRInterface
- Code: `modules/mobile_vr/register_types.cpp` — mobile_vr module registration

<a id="topic-physics-2d-collision-pipeline"></a>
## 74. native 2D collision pipeline

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Physics spaces hold collision objects; the 2D BVH broad phase finds candidate collision objects and the solver dispatches shape-pair routines, including separating-axis tests.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
template <typename ShapeA, typename ShapeB, bool castA = false, bool castB = false, bool withMargin = false>
class SeparatorAxisTest2D {
	const ShapeA *shape_A = nullptr;
	const ShapeB *shape_B = nullptr;
	const Transform2D *transform_A = nullptr;
	const Transform2D *transform_B = nullptr;
	real_t best_depth = 1e15;
	Vector2 best_axis;
#ifdef DEBUG_ENABLED
	int best_axis_count = 0;
	int best_axis_index = -1;
#endif
	Vector2 motion_A;
	Vector2 motion_B;
```

Source: `modules/godot_physics_2d/godot_collision_solver_2d_sat.cpp` — SeparatorAxisTest2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GodotSpace2D, GodotCollisionObject2D

Evidence:
- Code: `modules/godot_physics_2d/godot_broad_phase_2d_bvh.h:39` — GodotBroadPhase2DBVH
- Code: `modules/godot_physics_2d/godot_collision_solver_2d_sat.cpp:175` — SeparatorAxisTest2D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH construction](#topic-bvh-construction)
- [Collision shapes](#topic-collision-shapes)
- [native 3D collision pipeline](#topic-physics-3d-collision-pipeline)
- [2D physics nodes](#topic-two-dimensional-physics)
- [Spatial indexing and ray queries](#topic-spatial-indexing)

<a id="topic-physics-3d-collision-pipeline"></a>
## 75. native 3D collision pipeline

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The native 3D server manages spaces, bodies, shapes, and joints; its collision code includes GJK/EPA support and SAT shape-pair routines.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
template <typename ShapeA, typename ShapeB, bool withMargin = false>
class SeparatorAxisTest {
	const ShapeA *shape_A = nullptr;
	const ShapeB *shape_B = nullptr;
	const Transform3D *transform_A = nullptr;
	const Transform3D *transform_B = nullptr;
	real_t best_depth = 1e15;
	_CollectorCallback *callback = nullptr;
	real_t margin_A = 0.0;
	real_t margin_B = 0.0;
	Vector3 separator_axis;

public:
	Vector3 best_axis;
```

Source: `modules/godot_physics_3d/godot_collision_solver_3d_sat.cpp` — SeparatorAxisTest (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GodotSpace3D, GodotCollisionObject3D

Evidence:
- Code: `modules/godot_physics_3d/godot_physics_server_3d.h:41` — GodotPhysicsServer3D
- Code: `modules/godot_physics_3d/gjk_epa.cpp` — GJK and EPA implementation
- Code: `modules/godot_physics_3d/godot_collision_solver_3d_sat.cpp:618` — SeparatorAxisTest

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision shapes](#topic-collision-shapes)
- [Convex support and penetration](#topic-convex-collision)
- [native 2D collision pipeline](#topic-physics-2d-collision-pipeline)
- [3D physics nodes](#topic-three-dimensional-physics)

<a id="topic-scene-authoring"></a>
## 76. 2D and 3D scene authoring

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Scene tree and signal connections](#topic-scene-tree-and-signal-connections).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The partition implements distinct editing surfaces rather than a single dimension-agnostic viewport.

Scene authoring is split between CanvasItemEditor for 2D work and Node3DEditor with Node3DEditorViewport for 3D work.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Node3DEditor : public VBoxContainer {
	GDCLASS(Node3DEditor, VBoxContainer);

public:
	static const unsigned int VIEWPORTS_COUNT = 4;

	enum ToolMode {
		TOOL_MODE_TRANSFORM,
		TOOL_MODE_MOVE,
		TOOL_MODE_ROTATE,
		TOOL_MODE_SCALE,
		TOOL_MODE_SELECT,
		TOOL_MODE_LIST_SELECT,
```

Source: `editor/scene/3d/node_3d_editor_plugin.h` — Node3DEditor : public VBoxContainer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/canvas_item_editor_plugin.h:75` — CanvasItemEditor : public VBoxContainer
- Code: `editor/scene/3d/node_3d_editor_plugin.h:60` — Node3DEditor : public VBoxContainer
- Code: `editor/scene/3d/node_3d_editor_viewport.h:146` — Node3DEditorViewport : public Control

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-3d-gizmo-authoring"></a>
## 77. 3D gizmo authoring

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The editor distributes type-specific 3D visualization behavior across dedicated gizmo plug-ins.

Node3D editor gizmo plug-ins visualize and manipulate cameras, lights, meshes, occluders, particle emitters, physics objects, skeleton tools, and visibility volumes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Camera3DGizmoPlugin : public EditorNode3DGizmoPlugin {
	GDCLASS(Camera3DGizmoPlugin, EditorNode3DGizmoPlugin);

private:
	static float _find_closest_angle_to_half_pi_arc(const Vector3 &p_from, const Vector3 &p_to, float p_arc_radius, const Transform3D &p_arc_xform);

public:
	bool has_gizmo(Node3D *p_spatial) override;
	String get_gizmo_name() const override;
	int get_priority() const override;

	String get_handle_name(const EditorNode3DGizmo *p_gizmo, int p_id, bool p_secondary) const override;
	Variant get_handle_value(const EditorNode3DGizmo *p_gizmo, int p_id, bool p_secondary) const override;
```

Source: `editor/scene/3d/gizmos/camera_3d_gizmo_plugin.h` — Camera3DGizmoPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/3d/gizmos/camera_3d_gizmo_plugin.h:35` — Camera3DGizmoPlugin
- Code: `editor/scene/3d/gizmos/light_3d_gizmo_plugin.h:35` — Light3DGizmoPlugin
- Code: `editor/scene/3d/gizmos/mesh_instance_3d_gizmo_plugin.h:36` — MeshInstance3DGizmoPlugin
- Code: `editor/scene/3d/gizmos/physics/collision_shape_3d_gizmo_plugin.h:37` — CollisionShape3DGizmoPlugin

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Skeletons, animation, and ragdolls](#topic-skeletal-ragdoll)

<a id="topic-android-export-pipeline"></a>
## 78. Android export pipeline

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Android Gradle assembly](#topic-android-gradle-assembly), [APK signing and verification](#topic-apk-signing).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

`EditorExportPlatformAndroid` is the editor-facing exporter and its helpers hold export-specific data structures.

Android export derives Gradle project files, manifest content, localized project names, icons, ABIs, and plugin configuration from project and export inputs.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorExportPlatformAndroid : public EditorExportPlatform {
	GDCLASS(EditorExportPlatformAndroid, EditorExportPlatform);

	Ref<ImageTexture> logo;
	Ref<ImageTexture> run_icon;

	struct Device {
		String id;
		String name;
		String description;
		int api_level = 0;
		String architecture;
	};
```

Source: `platform/android/export/export_plugin.h` — EditorExportPlatformAndroid (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

APKExportData, PluginConfigAndroid, LauncherIcon, CustomExportData

Evidence:
- Code: `platform/android/export/export_plugin.h:64` — EditorExportPlatformAndroid
- Code: `platform/android/export/export_plugin.h:78` — APKExportData
- Code: `platform/android/export/godot_plugin_config.h:53` — PluginConfigAndroid
- Code: `platform/android/export/gradle_export_util.h:62` — CustomExportData

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Encoder configuration](#topic-encoder-configuration)

<a id="topic-resource-previews"></a>
## 79. Asynchronous resource previews

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Preview generation is a shared inspector-side service rather than a feature embedded in each resource editor.

Resource previews use generators selected by handles, queue source paths or in-memory resources, and cache generated full and small textures with metadata.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void EditorResourcePreview::_thread_func(void *ud) {
	EditorResourcePreview *erp = (EditorResourcePreview *)ud;
	erp->_thread();
}

void EditorResourcePreview::_preview_ready(const String &p_path, int p_hash, const Ref<Texture2D> &p_texture, const Ref<Texture2D> &p_small_texture, const Callable &p_callback, const Dictionary &p_metadata) {
	{
		MutexLock lock(preview_mutex);

		uint64_t modified_time = 0;

		if (!p_path.begins_with("ID:")) {
			modified_time = FileAccess::get_modified_time(p_path);
```

Source: `editor/inspector/editor_resource_preview.cpp` — EditorResourcePreview::_thread (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PreviewRequest, ResourcePreviewCacheEntry

Evidence:
- Code: `editor/inspector/editor_resource_preview.h:41` — EditorResourcePreviewGenerator
- Code: `editor/inspector/editor_resource_preview.h` — EditorResourcePreview::QueueItem
- Code: `editor/inspector/editor_resource_preview.cpp:425` — EditorResourcePreview::_thread

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Resource-backed assets](#topic-resource-assets)
- [Editor plugins and customization hooks](#topic-editor-extensibility)
- [Localization and translation-template generation](#topic-localization-and-template-generation)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-build-time-source-generation"></a>
## 80. Build-time source generation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These helpers convert build inputs into generated C++ declarations and byte data.

The Python build helpers generate C++ tables for documentation data, exporter registration, compressed documentation, and compressed translations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```python
def doc_data_class_path_builder(target, source, env):
    paths = dict(sorted(source[0].read().items()))
    data = "\n".join([f'\t{{"{key}", "{value}"}},' for key, value in paths.items()])
    with methods.generated_wrapper(str(target[0])) as file:
        file.write(
            f"""\
struct _DocDataClassPath {{
	const char *name;
	const char *path;
}};

inline constexpr int _doc_data_class_path_count = {len(paths)};
inline constexpr _DocDataClassPath _doc_data_class_paths[{len(paths) + 1}] = {{
```

Source: `editor/editor_builders.py` — doc_data_class_path_builder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorTranslationList

Evidence:
- Code: `editor/editor_builders.py:12` — doc_data_class_path_builder
- Code: `editor/editor_builders.py:32` — register_exporters_builder
- Code: `editor/editor_builders.py:55` — make_doc_header
- Code: `editor/editor_builders.py:71` — make_translations

<a id="topic-debug-adapter-protocol"></a>
## 81. Debug Adapter Protocol integration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The protocol layer translates debugger and editor state into a language-agnostic wire format.

The editor's debug adapter serializes protocol content as JSON, manages protocol data types, and starts a local protocol server.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
Vector<uint8_t> DAPeer::format_output(const Dictionary &p_params) const {
	const Vector<uint8_t> &content = Variant(p_params).to_json_string().to_utf8_buffer();
	Vector<uint8_t> response = vformat("Content-Length: %d\r\n\r\n", content.size()).to_utf8_buffer();

	response.append_array(content);
	return response;
}

Error DebugAdapterProtocol::on_client_connected() {
	ERR_FAIL_COND_V_MSG(clients.size() >= DAP_MAX_CLIENTS, FAILED, "Max client limits reached");

	Ref<StreamPeerTCP> tcp_peer = server->take_connection();
	ERR_FAIL_COND_V_MSG(tcp_peer.is_null(), FAILED, "Failed to take incoming DAP connection.");
	tcp_peer->set_no_delay(true);
```

Source: `editor/debugger/debug_adapter/debug_adapter_protocol.cpp` — Variant(p_params).to_json_string().to_utf8_buffer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

DAP::Source, DAP::Breakpoint, DAP::Capabilities

Evidence:
- Code: `editor/debugger/debug_adapter/debug_adapter_protocol.cpp:139` — Variant(p_params).to_json_string().to_utf8_buffer
- Code: `editor/debugger/debug_adapter/debug_adapter_types.h` — Source, Breakpoint, Capabilities, Message, Scope, StackFrame, Thread, and Variable
- External (primary, unverified (source anchor not found)): [Debug Adapter Protocol Overview](https://microsoft.github.io/debug-adapter-protocol/overview.html), accessed 2026-07-15

<a id="topic-editor-and-project-settings"></a>
## 82. Editor and project settings

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Input action and shortcut configuration](#topic-input-action-configuration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Editor preferences and project configuration are separate concerns in this partition.

EditorSettings persists editor-side setting containers and project metadata, while dedicated dialogs expose editor, project, autoload, action-map, and input-event configuration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ProjectSettingsEditor : public AcceptDialog {
	GDCLASS(ProjectSettingsEditor, AcceptDialog);

	inline static ProjectSettingsEditor *singleton = nullptr;

	enum {
		FEATURE_ALL,
		FEATURE_CUSTOM,
		FEATURE_FIRST,
	};

	ProjectSettings *ps = nullptr;
	Timer *timer = nullptr;
```

Source: `editor/settings/project_settings_editor.h` — ProjectSettingsEditor : public AcceptDialog (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorSettings

Evidence:
- Code: `editor/settings/editor_settings.h` — EditorSettings, Plugin, and VariantContainer
- Code: `editor/settings/editor_settings.cpp` — EditorSettings::_get_project_metadata_path and metadata load/save code
- Code: `editor/settings/project_settings_editor.h:51` — ProjectSettingsEditor : public AcceptDialog
- Code: `editor/settings/editor_autoload_settings.h` — EditorAutoloadSettings and AutoloadInfo

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Android export pipeline](#topic-android-export-pipeline)
- [Apple embedded packaging and signing](#topic-apple-embedded-packaging)
- [Audio buses, streams, and effects](#topic-audio-bus-processing)
- [Contextual completion contracts](#topic-completion-contexts)
- [Engine bootstrap](#topic-engine-bootstrap)
- [Frame timing and physics stepping](#topic-frame-timing)
- [High-level RPC routing](#topic-high-level-rpc)
- [ISA and platform portability](#topic-isa-portability)
- [OpenXR action configuration](#topic-openxr-action-configuration)
- [PSA key lifecycle and storage](#topic-psa-key-lifecycle)
- [Textures, meshes, and materials](#topic-rendering-resources)
- [Runtime configuration](#topic-runtime-configuration)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [TLS connection and session state](#topic-tls-connection-state)
- [Render-pass configuration](#topic-vulkan-render-pass-configuration)
- [Swapchain presentation](#topic-vulkan-swapchain-presentation)
- [Worker-based parallelism](#topic-worker-parallelism)
- [Encoder configuration](#topic-encoder-configuration)
- [Project settings and feature overrides](#topic-project-settings)
- [Class metadata and reflection](#topic-reflection)

<a id="topic-editor-authoring-workspaces"></a>
## 83. Editor authoring workspaces

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These are editor-facing authoring surfaces rather than runtime rendering backends.

The editor implements dock management, filesystem browsing, scene-tree editing, animation editing, asset-library browsing, audio editing, and debugger views as specialized controls and plugins.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class AnimationTrackEditor : public VBoxContainer {
	GDCLASS(AnimationTrackEditor, VBoxContainer);
	friend class AnimationTimelineEdit;
	friend class AnimationBezierTrackEdit;
	friend class AnimationMarkerKeyEditEditor;

	Ref<Animation> animation;
	bool read_only = false;
	Node *root = nullptr;

	AcceptDialog *read_only_dialog = nullptr;

	MenuButton *edit = nullptr;
```

Source: `editor/animation/animation_track_editor.h` — class AnimationTrackEditor : public VBoxContainer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/docks/editor_dock_manager.h:81` — class EditorDockManager : public Object
- Code: `editor/animation/animation_track_editor.h:606` — class AnimationTrackEditor : public VBoxContainer
- Code: `editor/docks/scene_tree_dock.h:51` — class SceneTreeDock : public EditorDock
- Code: `editor/debugger/script_editor_debugger.h:59` — class ScriptEditorDebugger : public MarginContainer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-editor-build-composition"></a>
## 84. Editor build composition

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Editor source aggregation is expressed as Python-based SCons configuration alongside the C++ implementation.

SCsub build scripts add C++ source files to the editor source set and invoke child SConscript files for nested editor subsystems.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```python
# See also `scene/theme/icons/default_theme_icons_builders.py`.
def make_editor_icons_action(target, source, env):
    icons_names = []
    icons_raw = []
    icons_med = []
    icons_big = []

    for idx, svg in enumerate(source):
        path = str(svg)
        with open(path, encoding="utf-8", newline="\n") as file:
            icons_raw.append(methods.to_raw_cstring(file.read()))

        name = os.path.splitext(os.path.basename(path))[0]
        icons_names.append(f'"{name}"')
```

Source: `editor/icons/editor_icons_builders.py` — make_editor_icons_action (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/import/SCsub:6` — env.add_source_files
- Code: `editor/scene/2d/SCsub:8` — SConscript
- Code: `editor/icons/editor_icons_builders.py:9` — make_editor_icons_action

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [SCons module build configuration](#topic-module-build-configuration)

<a id="topic-editor-scene-sessions"></a>
## 85. Editor scene sessions

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Editing selection history](#topic-editing-selection-history), [Editor plugin state and custom types](#topic-editor-plugin-state), [Undo and redo history](#topic-undo-redo-history).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

EditedScene is the durable editor-side record for an open scene rather than the scene tree itself.

EditorData represents each open editor scene with its root, project path, plugin state, selection, undo history ID, and timing/version metadata.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void EditorData::load_editor_plugin_states_from_config(const Ref<ConfigFile> &p_config_file, int p_idx) {
	ERR_FAIL_INDEX(p_idx, edited_scene.size());
	EditedScene &es = edited_scene.write[p_idx];

	const Vector<String> esl = p_config_file->get_section_keys("editor_states");

	Dictionary states;
	for (const String &E : esl) {
		const Variant state = p_config_file->get_value("editor_states", E);
		if (state.get_type() != Variant::NIL) {
			states[E] = state;
		}
	}
```

Source: `editor/editor_data.cpp` — EditorData::load_editor_plugin_states_from_config (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditedScene

Evidence:
- Code: `editor/editor_data.h` — EditorData::EditedScene
- Code: `editor/editor_data.cpp:385` — EditorData::load_editor_plugin_states_from_config

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Class metadata and reflection](#topic-reflection)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Undo and redo history](#topic-undo-redo-history)

<a id="topic-editor-theming"></a>
## 86. Editor theming

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Theme construction is centralized in EditorThemeManager while visual variants reside in separate builders.

Editor theming builds a Theme from ThemeConfiguration and settings, with distinct classic and modern builders plus font, icon, color-map, and scale support.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ThemeClassic {
public:
	static void populate_shared_styles(const Ref<EditorTheme> &p_theme, EditorThemeManager::ThemeConfiguration &p_config);
	static void populate_standard_styles(const Ref<EditorTheme> &p_theme, EditorThemeManager::ThemeConfiguration &p_config);
	static void populate_editor_styles(const Ref<EditorTheme> &p_theme, EditorThemeManager::ThemeConfiguration &p_config);
};
```

Source: `editor/themes/theme_classic.h` — ThemeClassic (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/themes/editor_theme_manager.h` — EditorThemeManager::ThemeConfiguration
- Code: `editor/themes/editor_theme_manager.cpp` — text editor color-theme settings and custom theme path handling
- Code: `editor/themes/theme_classic.h:35` — ThemeClassic
- Code: `editor/themes/theme_modern.h:35` — ThemeModern
- External (official, unverified (source anchor not found)): [Theme — Godot Engine documentation](https://docs.godotengine.org/en/stable/classes/class_theme.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Themes and style boxes](#topic-themes-and-style-boxes)
- [UI themes and style boxes](#topic-ui-theming)

<a id="topic-linuxbsd-desktop-integration"></a>
## 87. Linux/BSD desktop integration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Wayland window backend](#topic-wayland-window-backend).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The platform layer contains both generic Linux/BSD OS code and desktop-protocol integrations.

Linux/BSD platform code integrates OS services, portals, screensaver handling, TTS, X11, and Wayland display servers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class FreeDesktopPortalDesktop : public Object {
	GDSOFTCLASS(FreeDesktopPortalDesktop, Object);

private:
	bool unsupported = false;

	enum ReadVariantType {
		VAR_TYPE_UINT32, // u
		VAR_TYPE_BOOL, // b
		VAR_TYPE_COLOR, // (ddd)
	};

	static bool try_parse_variant(DBusMessage *p_reply_message, ReadVariantType p_type, void *r_value);
```

Source: `platform/linuxbsd/freedesktop_portal_desktop.h` — FreeDesktopPortalDesktop (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/linuxbsd/os_linuxbsd.h:48` — OS_LinuxBSD
- Code: `platform/linuxbsd/freedesktop_portal_desktop.h:45` — FreeDesktopPortalDesktop
- Code: `platform/linuxbsd/freedesktop_screensaver.h:37` — FreeDesktopScreenSaver
- Code: `platform/linuxbsd/tts_linux.h:49` — TTS_Linux

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Wayland window backend](#topic-wayland-window-backend)

<a id="topic-localization-and-template-generation"></a>
## 88. Localization and translation-template generation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Message context, singular text, plural text, comments, and source locations are explicitly handled during template generation.

Localization tooling parses translation inputs, exposes locale selection and preview, edits localization settings, and generates message maps and POT-style template output.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorTranslationPreviewMenu : public PopupMenu {
	GDCLASS(EditorTranslationPreviewMenu, PopupMenu);

	static constexpr int MENU_ID_PSEUDOLOCALIZATION = 1;

	void _prepare();
	void _pressed(int p_index);

protected:
	void _notification(int p_what);

public:
	EditorTranslationPreviewMenu();
```

Source: `editor/translations/editor_translation_preview_menu.h` — EditorTranslationPreviewMenu : public PopupMenu (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Translation Message

Evidence:
- Code: `editor/translations/editor_translation_parser.h` — EditorTranslationParserPlugin and EditorTranslationParser
- Code: `editor/translations/template_generator.cpp` — TranslationTemplateGenerator message context, plural, comment, location, MessageKey, and MessageMap handling
- Code: `editor/translations/editor_translation_preview_menu.h:35` — EditorTranslationPreviewMenu : public PopupMenu
- External (official, unverified (source anchor not found)): [Translation — Godot Engine documentation](https://docs.godotengine.org/en/stable/classes/class_translation.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Asynchronous resource previews](#topic-resource-previews)

<a id="topic-native-desktop-services"></a>
## 89. Native desktop services

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These services are separate native integrations rather than methods embedded solely in window management.

macOS and Windows platform code includes native-menu and text-to-speech adapters alongside their display servers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NativeMenuMacOS : public NativeMenu {
	GDCLASS(NativeMenuMacOS, NativeMenu)

	struct MenuData {
		NSMenu *menu = nullptr;

		Callable open_cb;
		Callable close_cb;
		bool is_open = false;
		bool is_system = false;
	};

	mutable RID_PtrOwner<MenuData> menus;
```

Source: `platform/macos/native_menu_macos.h` — NativeMenuMacOS (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/macos/native_menu_macos.h:40` — NativeMenuMacOS
- Code: `platform/macos/tts_macos.h:46` — TTSUtterance
- Code: `platform/windows/native_menu_windows.h:40` — NativeMenuWindows
- Code: `platform/windows/tts_windows.h:40` — TTS_Windows

<a id="topic-platform-exports"></a>
## 90. Platform export packaging

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Export code is platform-specific but follows EditorExportPlatform-derived services.

Platform export plug-ins implement macOS, Web, Windows, and visionOS export behavior; the Windows exporter also includes PE template modification structures.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorExportPlatformMacOS : public EditorExportPlatform {
	GDCLASS(EditorExportPlatformMacOS, EditorExportPlatform);

	int version_code = 0;

	Ref<ImageTexture> logo;

	struct SSHCleanupCommand {
		String host;
		String port;
		Vector<String> ssh_args;
		String cmd_args;
		bool wait = false;
```

Source: `platform/macos/export/export_plugin.h` — EditorExportPlatformMacOS (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/macos/export/export_plugin.h:52` — EditorExportPlatformMacOS
- Code: `platform/web/export/export_plugin.h:42` — EditorExportPlatformWeb
- Code: `platform/windows/export/template_modifier.h:37` — TemplateModifier
- Code: `platform/visionos/export/export_plugin.h:35` — EditorExportPlatformVisionOS

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Target-platform export](#topic-target-platform-export)

<a id="topic-project-catalog"></a>
## 91. Project catalog and selection

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The project manager operates on structured records rather than treating projects as paths alone.

ProjectList models known projects with name, path, tags, main scene, version, unsupported features, last-edit time, and favorite state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
struct ProjectListComparator {
	ProjectList::FilterOption order_option = ProjectList::FilterOption::EDIT_DATE;

	// operator<
	_FORCE_INLINE_ bool operator()(const ProjectList::Item &a, const ProjectList::Item &b) const {
		if (a.favorite && !b.favorite) {
			return true;
		}
		if (b.favorite && !a.favorite) {
			return false;
		}
		switch (order_option) {
			case ProjectList::PATH:
```

Source: `editor/project_manager/project_list.cpp` — ProjectList (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ProjectCatalog, ProjectCatalogItem

Evidence:
- Code: `editor/project_manager/project_list.h` — ProjectList::Item
- Code: `editor/project_manager/project_list.cpp:71` — ProjectList

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Motion blur](#topic-motion-blur)
- [Main loop, OS, and time services](#topic-runtime-loop-time)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-filesystem-project-index"></a>
## 92. Project filesystem index

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Export presets](#topic-export-presets), [Resource dependency resolution](#topic-resource-dependency-resolution).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The index separates directory topology from per-file metadata and supports rescanning.

The filesystem service scans project directories into a tree of file records with resource type, UID, import state, dependency list, and script-class metadata.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void EditorFileSystem::_scan_filesystem() {
	// On the first scan, the first_scan_root_dir is created in _first_scan_filesystem.
	ERR_FAIL_COND(!scanning || new_filesystem || (first_scan && !first_scan_root_dir));

	//read .fscache
	String cpath;

	sources_changed.clear();
	file_cache.clear();

	String fscache = EditorPaths::get_singleton()->get_project_settings_dir().path_join(CACHE_FILE_NAME);
	{
		Ref<FileAccess> f = FileAccess::open(fscache, FileAccess::READ);
```

Source: `editor/file_system/editor_file_system.cpp` — EditorFileSystem::_scan_filesystem (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorFileSystemDirectory, EditorFileSystemDirectory::FileInfo

Evidence:
- Code: `editor/file_system/editor_file_system.h:46` — EditorFileSystemDirectory
- Code: `editor/file_system/editor_file_system.h:165` — EditorFileSystemDirectory::FileInfo
- Code: `editor/file_system/editor_file_system.cpp:409` — EditorFileSystem::_scan_filesystem

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Resources and asset lifecycle](#topic-resources)
- [Script resources and instances](#topic-scripting)

<a id="topic-project-upgrade"></a>
## 93. Project source migration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Migration is source-aware instead of a blind string replacement.

ProjectConverter3To4 retains whether each source line is a comment while applying named rename, conversion, and check-only passes to project source.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct SourceLine {
	String line;
	bool is_comment;
};

class RegEx;
template <typename T>
class Ref;

class ProjectConverter3To4 {
	class RegExContainer;

	uint64_t maximum_file_size;
```

Source: `editor/project_upgrade/project_converter_3_to_4.h` — SourceLine (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SourceLine

Evidence:
- Code: `editor/project_upgrade/project_converter_3_to_4.h:39` — SourceLine
- Code: `editor/project_upgrade/project_converter_3_to_4.h:48` — ProjectConverter3To4
- Code: `editor/project_upgrade/renames_map_3_to_4.h:35` — RenamesMap3To4

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Character-encoding conversion](#topic-character-encoding-conversion)

<a id="topic-property-inspection"></a>
## 94. Reflective property inspection

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Specialized controls are selected by property shape while inspector plug-ins can supply custom editors.

The inspector reads Object properties into EditorProperty controls and supports custom parsing, revert, copy, paste, keying, pinning, and favorites.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorPropertyRevert {
public:
	static Variant get_property_revert_value(Object *p_object, const StringName &p_property, bool *r_is_valid);
	static bool can_property_revert(Object *p_object, const StringName &p_property, const Variant *p_custom_current_value = nullptr);
};

class EditorInspectorActionButton : public Button {
	GDCLASS(EditorInspectorActionButton, Button);

	StringName icon_name;

protected:
	void _notification(int p_what);
```

Source: `editor/inspector/editor_inspector.h` — EditorProperty (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/inspector/editor_inspector.h:70` — EditorProperty
- Code: `editor/inspector/editor_inspector.h:341` — EditorInspectorPlugin
- Code: `editor/inspector/editor_properties.h:795` — EditorInspectorDefaultPlugin

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)
- [Editor plugins and customization hooks](#topic-editor-extensibility)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)

<a id="topic-android-remote-engine-service"></a>
## 95. Remote Android engine service

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The service defines numeric message codes and bundle keys, while a fragment binds to and receives remote engine state.

Remote engine service messages convey engine status, errors, surface packages, dimensions, host tokens, and command-line parameters between Android UI and service code.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```kotlin
*/
open class GodotService : Service() {

	companion object {
		private val TAG = GodotService::class.java.simpleName

		const val EXTRA_MSG_PAYLOAD = "extraMsgPayload"

		// Keys to store / retrieve msg payloads
		const val KEY_COMMAND_LINE_PARAMETERS = "commandLineParameters"
		const val KEY_HOST_TOKEN = "hostToken"
		const val KEY_DISPLAY_ID = "displayId"
		const val KEY_WIDTH = "width"
		const val KEY_HEIGHT = "height"
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt` — GodotService (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GodotService.EngineStatus

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt:2` — GodotService
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt:79` — KEY_ENGINE_STATUS
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/RemoteGodotFragment.kt:2` — RemoteGodotFragment

<a id="topic-resource-specific-editors"></a>
## 96. Resource-specific editing

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The implementation chooses type-specific editing experiences for these resource families.

Dedicated editor controls and docks edit gradients, curves, materials, sprite frames, mesh libraries, textures, packed scenes, and resource preloaders.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ResourcePreloaderEditor : public EditorDock {
	GDCLASS(ResourcePreloaderEditor, EditorDock);

	enum {
		BUTTON_OPEN_SCENE,
		BUTTON_EDIT_RESOURCE,
		BUTTON_REMOVE
	};

	Button *load = nullptr;
	Button *paste = nullptr;
	FilterLineEdit *search = nullptr;
	MarginContainer *mc = nullptr;
```

Source: `editor/scene/resource_preloader_editor_plugin.h` — ResourcePreloaderEditor : public EditorDock (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/curve_editor_plugin.h` — CurveEdit and CurveEditor
- Code: `editor/scene/sprite_frames_editor_plugin.h:72` — SpriteFramesEditor : public EditorDock
- Code: `editor/scene/texture/texture_editor_plugin.h` — TexturePreview and TextureEditorPlugin
- Code: `editor/scene/resource_preloader_editor_plugin.h:43` — ResourcePreloaderEditor : public EditorDock

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-run-management"></a>
## 97. Running projects and instances

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Run UI, process embedding, device selection, and multiple-instance arguments are implemented as related editor services.

The run subsystem starts editor launches, exposes native-device runs, supports configurable multiple instances, and hosts embedded game-view processes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorRun {
public:
	enum Status {
		STATUS_PLAY,
		STATUS_PAUSED,
		STATUS_STOP
	};

	List<ProcessID> pids;

	struct WindowPlacement {
		int screen = 0;
		Point2i position = Point2i(INT_MAX, INT_MAX);
```

Source: `editor/run/editor_run.h` — EditorRun (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/run/editor_run.h:44` — EditorRun
- Code: `editor/run/editor_run_native.h:38` — EditorRunNative
- Code: `editor/run/run_instances_dialog.h` — RunInstancesDialog::InstanceData
- Code: `editor/run/embedded_process.h:85` — EmbeddedProcess

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Manifold mesh interchange](#topic-manifold-mesh-data)
- [Skeleton modification and physical-bone simulation](#topic-skeleton-modifiers)
- [Target-platform export](#topic-target-platform-export)

<a id="topic-scene-importing"></a>
## 98. Scene importing

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [COLLADA scene interchange](#topic-collada-import), [Post-import processing](#topic-post-import-processing).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the central pipeline for bringing 3D scene formats into the editor.

Scene importers expose extensions, options, flags, and an import operation that produces a Node-based scene.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorSceneFormatImporter : public RefCounted {
	GDCLASS(EditorSceneFormatImporter, RefCounted);

	List<ResourceImporter::ImportOption> *current_option_list = nullptr;

protected:
	static void _bind_methods();

	Node *import_scene_wrapper(const String &p_path, uint32_t p_flags, const Dictionary &p_options);
	Ref<Animation> import_animation_wrapper(const String &p_path, uint32_t p_flags, const Dictionary &p_options);

	GDVIRTUAL0RC_REQUIRED(Vector<String>, _get_extensions)
	GDVIRTUAL3R_REQUIRED(Object *, _import_scene, String, uint32_t, Dictionary)
```

Source: `editor/import/3d/resource_importer_scene.h` — EditorSceneFormatImporter (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ColladaParseState

Evidence:
- Code: `editor/import/3d/resource_importer_scene.h:49` — EditorSceneFormatImporter
- Code: `editor/import/3d/resource_importer_scene.h:155` — ResourceImporterScene

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene graph and lifecycle](#topic-scene-tree)

<a id="topic-shader-editing-and-language-plugins"></a>
## 99. Shader editing and language plugins

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Platform-selective shader baking](#topic-platform-selective-shader-baking).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Language selection is represented by EditorShaderLanguagePlugin instances rather than hard-coded solely in one text editor.

Shader editing uses a ShaderEditorPlugin, a text shader editor, syntax highlighting, shader-language plugins, creation dialogs, and shader-file editing.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorShaderLanguagePlugin : public RefCounted {
	GDCLASS(EditorShaderLanguagePlugin, RefCounted);

	static Vector<Ref<EditorShaderLanguagePlugin>> shader_languages;
	static Vector<Vector2i> language_variation_map;

public:
	static void register_shader_language(const Ref<EditorShaderLanguagePlugin> &p_shader_language);
	static void clear_registered_shader_languages();
	static const Vector<Ref<EditorShaderLanguagePlugin>> get_shader_languages_read_only();
	static int get_shader_language_variation_count();
	static Ref<EditorShaderLanguagePlugin> get_shader_language_for_index(int p_index);
	static String get_file_extension_for_index(int p_index);
```

Source: `editor/shader/editor_shader_language_plugin.h` — EditorShaderLanguagePlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/shader/shader_editor_plugin.h` — ShaderEditorPlugin and EditedShader
- Code: `editor/shader/text_shader_editor.h` — GDShaderSyntaxHighlighter, ShaderTextEditor, and TextShaderEditor
- Code: `editor/shader/editor_shader_language_plugin.h:35` — EditorShaderLanguagePlugin
- Code: `editor/shader/shader_create_dialog.h` — ShaderCreateDialog and ShaderTypeData
- External (official, unverified (source anchor not found)): [Shading language — Godot Engine documentation](https://docs.godotengine.org/en/stable/tutorials/shaders/shader_reference/shading_language.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript editor services](#topic-gdscript-editor-services)

<a id="topic-resource-importing"></a>
## 100. Standalone resource importing

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

This is separate from scene import and produces editor resource assets from individual files.

Resource-importer classes handle images, textures, texture atlases, SVG, fonts, WAV audio, CSV translations, shaders, and bitmaps.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ResourceImporterBitMap : public ResourceImporter {
	GDCLASS(ResourceImporterBitMap, ResourceImporter);

public:
	virtual String get_importer_name() const override;
	virtual String get_visible_name() const override;
	virtual void get_recognized_extensions(List<String> *p_extensions) const override;
	virtual String get_save_extension() const override;
	virtual String get_resource_type() const override;

	virtual int get_preset_count() const override;
	virtual String get_preset_name(int p_idx) const override;
```

Source: `editor/import/resource_importer_bitmask.h` — ResourceImporterBitMap (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/import/resource_importer_image.h:35` — ResourceImporterImage
- Code: `editor/import/resource_importer_texture.h:40` — ResourceImporterTexture
- Code: `editor/import/resource_importer_texture_atlas.h:35` — ResourceImporterTextureAtlas
- Code: `editor/import/resource_importer_svg.h:35` — ResourceImporterSVG
- Code: `editor/import/resource_importer_dynamic_font.h:35` — ResourceImporterDynamicFont
- Code: `editor/import/resource_importer_wav.h:35` — ResourceImporterWAV
- Code: `editor/import/resource_importer_csv_translation.h:35` — ResourceImporterCSVTranslation
- Code: `editor/import/resource_importer_shader_file.h:35` — ResourceImporterShaderFile
- Code: `editor/import/resource_importer_bitmask.h:35` — ResourceImporterBitMap

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-tile-authoring"></a>
## 101. Tile authoring

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Tile editing is implemented as several coordinated editor tools rather than a single monolithic panel.

Tile tools edit atlas sources, per-tile data, terrain data, proxies, map layers, and scene-collection sources.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class TileDataEditor : public VBoxContainer {
	GDCLASS(TileDataEditor, VBoxContainer);

private:
	bool _tile_set_changed_update_needed = false;
	void _tile_set_changed_plan_update();
	void _tile_set_changed_deferred_update();

protected:
	Ref<TileSet> tile_set;
	TileData *_get_tile_data(TileMapCell p_cell);
	virtual void _tile_set_changed() {}
```

Source: `editor/scene/2d/tiles/tile_data_editors.h` — TileDataEditor (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/2d/tiles/tile_set_atlas_source_editor.h:43` — TileSetAtlasSourceEditor
- Code: `editor/scene/2d/tiles/tile_data_editors.h:42` — TileDataEditor
- Code: `editor/scene/2d/tiles/tile_map_layer_editor.h:51` — TileMapLayerEditor
- Code: `editor/scene/2d/tiles/tile_proxies_manager_dialog.h:41` — TileProxiesManagerDialog

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Tile libraries, cells, and patterns](#topic-tile-libraries)

<a id="topic-version-control-integration"></a>
## 102. Version-control integration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The interface separates VCS-provided data structures from the editor-facing plugin.

Version-control integration exchanges structured status, commit, file-diff, hunk, and line data through EditorVCSInterface and presents it through VersionControlEditorPlugin.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class VersionControlEditorPlugin : public EditorPlugin {
	GDCLASS(VersionControlEditorPlugin, EditorPlugin)

public:
	enum ButtonType {
		BUTTON_TYPE_OPEN = 0,
		BUTTON_TYPE_DISCARD = 1,
	};

	enum DiffViewType {
		DIFF_VIEW_TYPE_SPLIT = 0,
		DIFF_VIEW_TYPE_UNIFIED = 1,
	};
```

Source: `editor/version_control/version_control_editor_plugin.h` — VersionControlEditorPlugin : public EditorPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VCS Diff File, VCS Diff Hunk, VCS Diff Line

Evidence:
- Code: `editor/version_control/editor_vcs_interface.h` — EditorVCSInterface::DiffLine, DiffHunk, DiffFile, Commit, and StatusFile
- Code: `editor/version_control/version_control_editor_plugin.h:44` — VersionControlEditorPlugin : public EditorPlugin
- External (official, unverified (source anchor not found)): [Version control systems — Godot Engine documentation](https://docs.godotengine.org/en/4.4/tutorials/best_practices/version_control_systems.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GUI controls and layout](#topic-gui-control-layout)
- [Scene construction and commit](#topic-scene-commit)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-ios-embedded-adapter"></a>
## 103. iOS embedded adapter

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The iOS build lists device metrics, display, view, main, and OS Objective-C++ sources in one library.

iOS platform classes adapt the engine’s OS, display, and view responsibilities for an embedded Apple target.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class DisplayServerIOS final : public DisplayServerAppleEmbedded {
	GDSOFTCLASS(DisplayServerIOS, DisplayServerAppleEmbedded);

	_THREAD_SAFE_CLASS_

	DisplayServerIOS(const String &p_rendering_driver, DisplayServerEnums::WindowMode p_mode, DisplayServerEnums::VSyncMode p_vsync_mode, uint32_t p_flags, const Vector2i *p_position, const Vector2i &p_resolution, int p_screen, DisplayServerEnums::Context p_context, int64_t p_parent_window, Error &r_error);
	~DisplayServerIOS();

public:
	static DisplayServerIOS *get_singleton();

	static void register_ios_driver();
	static DisplayServer *create_func(const String &p_rendering_driver, DisplayServerEnums::WindowMode p_mode, DisplayServerEnums::VSyncMode p_vsync_mode, uint32_t p_flags, const Vector2i *p_position, const Vector2i &p_resolution, int p_screen, DisplayServerEnums::Context p_context, int64_t p_parent_window, Error &r_error);
```

Source: `platform/ios/display_server_ios.h` — DisplayServerIOS (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/ios/SCsub:10` — ios_lib
- Code: `platform/ios/ios.h:35` — iOS
- Code: `platform/ios/display_server_ios.h:35` — DisplayServerIOS
- Code: `platform/ios/os_ios.h:37` — OS_IOS

<a id="topic-aabb-tree-construction"></a>
## 104. AABB tree construction

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Imagine checking a moving object against a large mesh. Testing every triangle is wasteful when most triangles are far away. A bounding-volume tree first rejects whole regions of space, then examines only the triangles in a region that might matter.

This is a geometry-acceleration structure within the supplied Jolt partition.

Jolt's AABBTreeBuilder stores nodes and indexed triangles and accepts a maximum-triangles-per-leaf setting while building its tree.

### Vocabulary used in this section

- **axis-aligned bounding box (AABB)** — A box whose faces stay parallel to the coordinate axes. Here, each tree node uses one to mark the region covered by its contents.
- **indexed triangle** — A triangle stored in a shared collection and referred to by position in that collection rather than copied into every tree node.
- **internal node and leaf** — An internal node points to child nodes; a leaf stops splitting and holds a bounded batch of triangles.
- **Jolt** — The physics library vendored under Godot's third-party source tree for this implementation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
uint AABBTreeBuilder::BuildInternal(const TriangleSplitter::Range &inTriangles)
{
	// Check if there are too many triangles left
	if (inTriangles.Count() > mMaxTrianglesPerLeaf)
	{
		// Split triangles in two batches
		TriangleSplitter::Range left, right;
		if (!mTriangleSplitter.Split(inTriangles, left, right))
		{
			// When the trace below triggers:
			//
			// This code builds a tree structure to accelerate collision detection.
			// At top level it will start with all triangles in a mesh and then divides the triangles into two batches.
```

Source: `thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.cpp` — AABBTreeBuilder (excerpt selected locally from grounded evidence)

**What to notice:** `inTriangles` is the current batch. The size check compares it with `mMaxTrianglesPerLeaf`; only an oversized batch is split into left and right work. Those recursive calls create child nodes, so a later collision query can first reject a whole bounding box before considering the triangles held by a leaf.

### Concrete structures to recognize

AABBTreeBuilder

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.h:35` — AABBTreeBuilder
- Code: `thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.cpp:7` — AABBTreeBuilder

<a id="topic-allocation"></a>
## 105. Allocation and reference ownership

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Builders receive allocator callbacks, while shared runtime objects inherit or use RefCount-based ownership.

Reference-counted objects, FastAllocator blocks, cached allocation, aligned allocation, and monitored allocation support resource and BVH memory management.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class RefCount
  {
  public:
    RefCount(int val = 0) : refCounter(val) {}
    virtual ~RefCount() {};
  
    virtual RefCount* refInc() { refCounter.fetch_add(1); return this; }
    virtual void refDec() { if (refCounter.fetch_add(-1) == 1) delete this; }
  private:
    std::atomic<size_t> refCounter;
  };
  
  ////////////////////////////////////////////////////////////////////////////////
```

Source: `thirdparty/embree/common/sys/ref.h` — RefCount (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RefCount, Ref, FastAllocator, CachedAllocator

Evidence:
- Code: `thirdparty/embree/common/sys/ref.h:15` — RefCount
- Code: `thirdparty/embree/common/sys/ref.h:32` — Ref
- Code: `thirdparty/embree/kernels/common/alloc.h:15` — FastAllocator
- Code: `thirdparty/embree/kernels/common/alloc.h:256` — CachedAllocator

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry buffer storage](#topic-buffer-storage)
- [Class-reference generation](#topic-class-reference-generation)
- [Godot-specific third-party adaptation](#topic-godot-thirdparty-adaptation)
- [D3D12 memory allocation](#topic-gpu-memory-suballocation)
- [Histograms and Huffman codes](#topic-histograms-and-huffman-codes)
- [HTTP and multiplayer](#topic-networking)
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Engine object model](#topic-object-model)
- [Pseudo-random generation](#topic-random-generation)
- [ZIP archive I/O](#topic-zip-archive-io)
- [BVH construction](#topic-bvh-construction)
- [Resources and asset lifecycle](#topic-resources)
- [Spatial indexing and ray queries](#topic-spatial-indexing)

<a id="topic-basis-transcoding"></a>
## 106. Basis texture transcoding

**Scope:** Vendored: basis_universal

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Basis file layout](#topic-basis-container-layout), [GPU block texture conversion](#topic-block-texture-transcoding), [KTX2 container transcoding](#topic-ktx2-container-transcoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation exposes a container-level transcoder plus specialized ETC1S, UASTC, and ASTC paths.

The Basis Universal subsystem decodes ETC1S and UASTC texture slices into selected texture formats through high-level and low-level transcoders.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class basisu_transcoder;

	// This struct holds all state used during transcoding. For video, it needs to persist between image transcodes (it holds the previous frame).
	// For threading you can use one state per thread.
	struct basisu_transcoder_state
	{
		struct block_preds
		{
			uint16_t m_endpoint_index;
			uint8_t m_pred_bits;
		};

		basisu::vector<block_preds> m_block_endpoint_preds[2];
```

Source: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — basisu_transcoder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BasisFileHeader, BasisSliceDescriptor

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:1` — basisu_transcoder
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:219` — basisu_lowlevel_etc1s_transcoder

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Texture compression pipeline](#topic-texture-compression-pipeline)

<a id="topic-font-table-access"></a>
## 107. Binary font-table access

**Scope:** Vendored: graphite

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Graphite SILF rule execution](#topic-graphite-rule-execution), [Graphite segment shaping](#topic-graphite-shaping).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code directly models table records, offsets, and character-to-glyph lookup structures.

Graphite and HarfBuzz read binary font tables, including cmap mappings, to obtain glyph data and layout behavior.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Face::Table
{
    const Face *            _f;
    mutable const byte *    _p;
    size_t                  _sz;
    bool                    _compressed;

    Error decompress();

    void release();

public:
    Table() throw();
```

Source: `thirdparty/graphite/src/inc/Face.h` — Face::Table (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Face, hb_blob_t, hb_face_t

Evidence:
- Code: `thirdparty/graphite/src/inc/Face.h:146` — Face::Table
- Code: `thirdparty/graphite/src/CmapCache.cpp` — bmp_subtable, smp_subtable, and Cmap
- Code: `thirdparty/harfbuzz/src/hb-open-file.hh` — OpenTypeFontFile and OpenTypeFontFace
- Code: `thirdparty/harfbuzz/src/hb-ot-cmap-table.hh` — OT::cmap and cmap::accelerator_t
- External (official, unverified (source anchor not found)): [cmap — Character To Glyph Index Mapping Table](https://learn.microsoft.com/en-us/typography/opentype/spec/cmap), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [SFNT table loading](#topic-sfnt-tables)

<a id="topic-block-texture-encoding"></a>
## 108. Block texture encoding

**Scope:** Vendored: cvtt

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is a separate encoder subsystem from Basis transcoding.

CVTT compression input uses pixel blocks, options, encoding plans, and format-specific BC6H, BC7, ETC, and S3TC computation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const BC6HData::ModeDescriptor *desc = BC6HData::g_modeDescriptors[mode];

        const size_t headerBits = modeInfo.m_partitioned ? 82 : 65;

        for (int subset = 0; subset < 2; subset++)
        {
            for (int epi = 0; epi < 2; epi++)
            {
                for (int ch = 0; ch < 3; ch++)
                    eps[subset][epi][ch] = ParallelMath::Extract(bestEndPoints[subset][epi][ch], block);
            }
        }
```

Source: `thirdparty/cvtt/ConvectionKernels_BC67.cpp` — BC6HData::ModeDescriptor (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/cvtt/ConvectionKernels.h` — cvtt::PixelBlockU8
- Code: `thirdparty/cvtt/ConvectionKernels.h:254` — cvtt::Options
- Code: `thirdparty/cvtt/ConvectionKernels_BC67.cpp:3200` — BC6HData::ModeDescriptor
- Code: `thirdparty/cvtt/ConvectionKernels_ETC.cpp:47` — cvtt::Internal::ETCComputer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU block texture conversion](#topic-block-texture-transcoding)
- [Texture block codecs](#topic-texture-block-codecs)

<a id="topic-stream-networking"></a>
## 109. Byte streams and socket servers

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

StreamPeerBuffer provides the same stream cursor model over an in-memory byte array.

StreamPeer exposes typed and raw byte-stream I/O, specialized peers implement TCP, TLS, compression, and Unix-domain sockets, and socket servers accept peer connections.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<description>
		StreamPeer is an abstract base class mostly used for stream-based protocols (such as TCP). It provides an API for sending and receiving data through streams as raw data or strings.
		[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="get_8">
			<return type="int" />
			<description>
				Gets a signed byte from the stream.
			</description>
		</method>
		<method name="get_16">
```

Source: `doc/classes/StreamPeer.xml` — StreamPeer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

StreamPeerBuffer, RID

Evidence:
- Code: `doc/classes/StreamPeer.xml:2` — StreamPeer
- Code: `doc/classes/StreamPeerBuffer.xml:2` — StreamPeerBuffer
- Code: `doc/classes/TCPServer.xml:2` — TCPServer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)
- [TLS connection and session state](#topic-tls-connection-state)
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-class-reference-generation"></a>
## 110. Class-reference generation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Python tool owns the intermediate documentation model and output formatting.

Documentation generation parses class-reference XML into class, type, property, parameter, signal, method, and constant definitions before producing reStructuredText.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```python
try:
            state.parse_class(data[0], data[1])
        except Exception as e:
            print_error(f"{name}.xml: Exception while parsing class: {e}", state)

    state.sort_classes()

    pattern = re.compile(args.filter)

    # Create the output folder recursively if it doesn't already exist.
    os.makedirs(args.output, exist_ok=True)

    print("Generating the RST class reference...")
```

Source: `doc/tools/make_rst.py` — State.parse_class (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ClassDef, TypeName, ClassStatusProgress

Evidence:
- Code: `doc/tools/make_rst.py` — State.parse_class
- Code: `doc/tools/make_rst.py:210` — TypeName.from_element
- Code: `doc/tools/make_rst.py:91` — make_rst_class

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [Dynamic invocation and signals](#topic-dynamic-invocation-signals)

<a id="topic-collision-shapes"></a>
## 111. Collision shapes

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Collision filtering](#topic-collision-filtering), [Convex support and penetration](#topic-convex-collision), [Geometry preprocessing](#topic-geometry-preprocessing), [Narrow-phase collision queries](#topic-narrow-phase).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Shape hierarchy is the common geometry contract for bodies and collision queries.

Collision shapes define the geometric representation used for ray casts, shape casts, point tests, overlap tests, triangle extraction, and collision dispatch.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// Dispatch function, main function to handle collisions between shapes
class JPH_EXPORT CollisionDispatch
{
public:
	/// Collide 2 shapes and pass any collision on to ioCollector
	/// @param inShape1 The first shape
	/// @param inShape2 The second shape
	/// @param inScale1 Local space scale of shape 1 (scales relative to its center of mass)
	/// @param inScale2 Local space scale of shape 2 (scales relative to its center of mass)
	/// @param inCenterOfMassTransform1 Transform to transform center of mass of shape 1 into world space
	/// @param inCenterOfMassTransform2 Transform to transform center of mass of shape 2 into world space
	/// @param inSubShapeIDCreator1 Class that tracks the current sub shape ID for shape 1
	/// @param inSubShapeIDCreator2 Class that tracks the current sub shape ID for shape 2
	/// @param inCollideShapeSettings Options for the CollideShape test
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionDispatch.h` — CollisionDispatch (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shape, ShapeSettings

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h` — ShapeSettings, Shape, ShapeFunctions
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionDispatch.h:18` — CollisionDispatch

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [native 2D collision pipeline](#topic-physics-2d-collision-pipeline)
- [native 3D collision pipeline](#topic-physics-3d-collision-pipeline)
- [Physics shapes, objects, and collision reports](#topic-physics-collision)
- [Physics queries and motion tests](#topic-physics-queries)
- [2D and 3D physics queries](#topic-physics-server-queries)
- [3D physics query contracts](#topic-physics-space-queries)
- [Texture resources and fallback placeholders](#topic-textures-and-placeholders)
- [Geometry families](#topic-geometry-families)
- [OpenXR dispatch forwarding](#topic-openxr-dispatch)
- [Ray query state](#topic-ray-query)

<a id="topic-color-font-paint"></a>
## 112. Color-font paint processing

**Scope:** Vendored: harfbuzz

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Raster font rendering](#topic-raster-font-rendering), [Vector font export](#topic-vector-font-export).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The paint layer is independent of a particular raster or vector output format.

HarfBuzz supplies paint functions plus bounded and extents contexts to evaluate color-font paint operations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
typedef struct  hb_paint_bounded_context_t hb_paint_bounded_context_t;

struct hb_paint_bounded_context_t
{
  void clear ()
  {
    clips = 0;
    bounded = true;
    groups.clear ();
  }

  hb_paint_bounded_context_t ()
  {
```

Source: `thirdparty/harfbuzz/src/hb-paint-bounded.hh` — hb_paint_bounded_context_t (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_paint_funcs_t, hb_paint_extents_context_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-paint.hh:53` — hb_paint_funcs_t
- Code: `thirdparty/harfbuzz/src/hb-paint-bounded.hh:34` — hb_paint_bounded_context_t

<a id="topic-compact-heightfield"></a>
## 113. Compact heightfield representation

**Scope:** Vendored: recastnavigation

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Navigation region construction](#topic-navigation-regions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Region algorithms index cells and spans through the heightfield's width and height.

A compact heightfield stores grid dimensions, compact cells, spans, and per-span area data for navigation processing.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/// @see rcCompactHeightfield, rcCompactSpan, rcBuildDistanceField, rcBuildRegionsMonotone, rcConfig
bool rcBuildRegionsMonotone(rcContext* ctx, rcCompactHeightfield& chf,
							const int borderSize, const int minRegionArea, const int mergeRegionArea)
{
	rcAssert(ctx);
	
	rcScopedTimer timer(ctx, RC_TIMER_BUILD_REGIONS);
	
	const int w = chf.width;
	const int h = chf.height;
	unsigned short id = 1;
	
	rcScopedDelete<unsigned short> srcReg((unsigned short*)rcAlloc(sizeof(unsigned short)*chf.spanCount, RC_ALLOC_TEMP));
	if (!srcReg)
```

Source: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp` — rcBuildRegions (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

rcCompactHeightfield

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp:1247` — rcBuildRegions

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Navigation contours and polygons](#topic-navmesh-contours-polygons)
- [Navigation heightfields](#topic-navmesh-heightfields)
- [Navigation agents and regions](#topic-navigation-agents)

<a id="topic-cryptography"></a>
## 114. Cryptographic keys, hashing, and TLS support

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

CryptoKey is a Resource, so key material participates in the engine resource lifecycle.

Crypto, CryptoKey, HashingContext, and HMACContext expose key generation, signatures, encryption, hashing, and message authentication APIs.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<description>
		The Crypto class provides access to advanced cryptographic functionalities.
		Currently, this includes asymmetric key encryption/decryption, signing/verification, and generating cryptographically secure random bytes, RSA keys, HMAC digests, and self-signed [X509Certificate]s.
		[codeblocks]
		[gdscript]
		var crypto = Crypto.new()

		# Generate new RSA key.
		var key = crypto.generate_rsa(4096)

		# Generate new self-signed certificate with the given key.
		var cert = crypto.generate_self_signed_certificate(key, "CN=mydomain.com,O=My Game Company,C=IT")

		# Save key and certificate in the user folder.
```

Source: `doc/classes/Crypto.xml` — Crypto (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `doc/classes/Crypto.xml:2` — Crypto
- Code: `doc/classes/CryptoKey.xml:2` — CryptoKey
- Code: `doc/classes/HMACContext.xml:2` — HMACContext

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Cryptographic resources and contexts](#topic-crypto-resources)
- [Resource formats and serialization](#topic-resource-formats)
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-device-runtime"></a>
## 115. Device runtime

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Scene construction and commit](#topic-scene-commit).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

It is the API-level runtime owner behind scene and buffer creation.

A Device configures Embree execution, exposes device and thread error messages, and limits tasking resources.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const char* Device::getThreadLastErrorMessage()
  {
    RTCErrorMessage* stored_error = g_errorHandler.error();
    return stored_error->msg.c_str();
  }

  void Device::process_error(Device* device, RTCError error, const char* str)
  {
    /* store global error code when device construction failed */
    if (!device)
      return setThreadErrorCode(error, str ? std::string(str) : std::string());

    /* print error when in verbose mode */
```

Source: `thirdparty/embree/kernels/common/device.cpp` — Device::getThreadLastErrorMessage (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCDevice, Device

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h:11` — RTCDevice
- Code: `thirdparty/embree/kernels/common/device.h:16` — Device
- Code: `thirdparty/embree/kernels/common/device.cpp:302` — Device::getThreadLastErrorMessage

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Diagnostic expectations](#topic-diagnostic-expectations)
- [Dynamic runtime values](#topic-dynamic-variant)
- [Engine bootstrap](#topic-engine-bootstrap)
- [Gamepad mapping and classification](#topic-gamepad-mapping)
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [HID device enumeration and backends](#topic-hid-device-enumeration)
- [MetalFX scaling and interpolation](#topic-metalfx-scaling)
- [Engine object model](#topic-object-model)
- [OpenXR composition layers](#topic-openxr-composition-layers)
- [OpenXR dispatch forwarding](#topic-openxr-dispatch)
- [OpenXR extension wrappers](#topic-openxr-extension-wrappers)
- [OpenXR render models](#topic-openxr-render-models)
- [OpenXR runtime integration](#topic-openxr-runtime-integration)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Pseudo-random generation](#topic-random-generation)
- [Class metadata and reflection](#topic-reflection)
- [Rendering backend abstraction](#topic-rendering-backends)
- [Running projects and instances](#topic-run-management)
- [Runtime configuration](#topic-runtime-configuration)
- [Main loop, OS, and time services](#topic-runtime-loop-time)
- [Runtime class metadata](#topic-runtime-metadata)
- [Scene graph nodes](#topic-scene-graph)
- [Event queue and watches](#topic-sdl-event-queue)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [UPnP device discovery](#topic-upnp-device-discovery)
- [UPnP port-mapping control](#topic-upnp-port-mapping)
- [Resource-backed assets](#topic-resource-assets)
- [Task scheduling and synchronization](#topic-task-scheduling)

<a id="topic-enet-host-peer-transport"></a>
## 116. ENet host and peer transport

**Scope:** Vendored: enet

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [ENet wire commands](#topic-enet-wire-commands).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ENet creates a host with a bounded peer allocation and manages peer state changes when connecting or disconnecting.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
ENetHost *
enet_host_create (const ENetAddress * address, size_t peerCount, size_t channelLimit, enet_uint32 incomingBandwidth, enet_uint32 outgoingBandwidth)
{
    ENetHost * host;
    ENetPeer * currentPeer;

    if (peerCount > ENET_PROTOCOL_MAXIMUM_PEER_ID)
      return NULL;

    host = (ENetHost *) enet_malloc (sizeof (ENetHost));
    if (host == NULL)
      return NULL;
    memset (host, 0, sizeof (ENetHost));
```

Source: `thirdparty/enet/host.c` — enet_host_create (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ENetHost, ENetPeer

Evidence:
- Code: `thirdparty/enet/host.c:29` — enet_host_create
- Code: `thirdparty/enet/protocol.c:36` — enet_protocol_change_state
- Code: `thirdparty/enet/enet/enet.h:41` — _ENetHost
- Code: `thirdparty/enet/enet/enet.h:271` — _ENetPeer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Godot ENet socket adaptation](#topic-godot-enet-socket-adaptation)
- [Wraparound network time](#topic-wraparound-network-time)
- [ENet transport and multiplayer adaptation](#topic-enet-transport-and-multiplayer)

<a id="topic-editor-extensibility"></a>
## 117. Editor plugins and customization hooks

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Editor-facing APIs are distinct from exported-project APIs and specialize controls, property editing, importing, and export behavior.

EditorPlugin subclasses register inspector, import, export, debugger, and resource-preview hooks in the editor.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
public partial class MySpecialPlugin : EditorImportPlugin
		{
			public override string _GetImporterName()
			{
				return "my.special.plugin";
			}

			public override string _GetVisibleName()
			{
				return "Special Mesh";
			}

			public override string[] _GetRecognizedExtensions()
```

Source: `doc/classes/EditorImportPlugin.xml` — EditorImportPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource

Evidence:
- Code: `doc/classes/EditorPlugin.xml:2` — EditorPlugin
- Code: `doc/classes/EditorInspectorPlugin.xml:2` — EditorInspectorPlugin
- Code: `doc/classes/EditorImportPlugin.xml:2` — EditorImportPlugin

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Custom inspector property editors](#topic-inspector-property-editors)
- [Reflective property inspection](#topic-property-inspection)
- [Asynchronous resource previews](#topic-resource-previews)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-entropy-bitstreams"></a>
## 118. Entropy bitstreams

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The same utility area contains both writing and parsing primitives.

VP8 and VP8L writers serialize codec decisions into range-coded or raw bit-level byte streams while readers maintain cursor, range, and value state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#if (BITTRACE > 0)
struct VP8BitReader;
extern void BitTrace(const struct VP8BitReader* const br, const char label[]);
#define BT_TRACK(br) BitTrace(br, label)
#define VP8Get(BR, L) VP8GetValue(BR, 1, L)
#else
#define BT_TRACK(br)
// We'll REMOVE the 'const char label[]' from all signatures and calls (!!):
#define VP8GetValue(BR, N, L) VP8GetValue(BR, N)
#define VP8Get(BR, L) VP8GetValue(BR, 1, L)
#define VP8GetSignedValue(BR, N, L) VP8GetSignedValue(BR, N)
#define VP8GetBit(BR, P, L) VP8GetBit(BR, P)
#define VP8GetBitAlt(BR, P, L) VP8GetBitAlt(BR, P)
#define VP8GetSigned(BR, V, L) VP8GetSigned(BR, V)
```

Source: `thirdparty/libwebp/src/utils/bit_reader_utils.h` — struct VP8BitReader (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/libwebp/src/utils/bit_writer_utils.h` — struct VP8BitWriter, VP8LBitWriter
- Code: `thirdparty/libwebp/src/utils/bit_writer_utils.c` — VP8PutBit, VP8LBitWriterAppend
- Code: `thirdparty/libwebp/src/utils/bit_reader_utils.h:33` — struct VP8BitReader
- Code: `thirdparty/libwebp/src/utils/bit_reader_inl_utils.h` — VP8GetBit, VP8GetSignedValue

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Endian-safe binary I/O](#topic-endian-safe-binary-io)
- [Lossless pixel transform pipeline](#topic-lossless-transform-pipeline)

<a id="topic-sdl-event-queue"></a>
## 119. Event queue and watches

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Queued entries can carry temporary memory associated with an event payload.

The event runtime stores SDL_Event values in a mutex-protected linked queue with an atomic count and event-watch support.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
typedef struct SDL_EventEntry
{
    SDL_Event event;
    SDL_TemporaryMemory *memory;
    struct SDL_EventEntry *prev;
    struct SDL_EventEntry *next;
} SDL_EventEntry;

static struct
{
    SDL_Mutex *lock;
    bool active;
    SDL_AtomicInt count;
```

Source: `thirdparty/sdl/events/SDL_events.c` — SDL_EventEntry (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_Event, SDL_EventEntry

Evidence:
- Code: `thirdparty/sdl/events/SDL_events.c:127` — SDL_EventEntry
- Code: `thirdparty/sdl/events/SDL_events.c:144` — SDL_EventQ
- Code: `thirdparty/sdl/events/SDL_eventwatch_c.h:30` — SDL_EventWatchList
- External (official, verified): [SDL3 CategoryEvents](https://wiki.libsdl.org/SDL3/CategoryEvents), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Device runtime](#topic-device-runtime)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Platform display and window adaptation](#topic-platform-display-adaptation)
- [Threads and synchronization](#topic-synchronization-primitives)

<a id="topic-explicit-drm-syncobj"></a>
## 120. Explicit DRM synchronization objects

**Scope:** Vendored: wayland-protocols

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The schema models synchronization as externally exchanged timeline objects and 64-bit point values.

Wayland protocol objects import DRM synchronization timelines and attach acquire and release timeline points to a surface commit.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
affected by this request, in particular timeline points set by
        set_acquire_point and set_release_point are not unset.
      </description>
    </request>
  </interface>

  <interface name="wp_linux_drm_syncobj_surface_v1" version="1">
    <description summary="per-surface explicit synchronization">
      This object is an add-on interface for wl_surface to enable explicit
      synchronization.

      Each surface can be associated with only one object of this interface at
      any time.
```

Source: `thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml` — set_acquire_point and set_release_point (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

wp_linux_drm_syncobj_timeline_v1

Evidence:
- Code: `thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml` — interface wp_linux_drm_syncobj_manager_v1
- Code: `thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml` — interface wp_linux_drm_syncobj_timeline_v1
- Code: `thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml:121` — set_acquire_point and set_release_point
- External (primary, unverified (source anchor not found)): [linux-drm-syncobj-v1 Wayland protocol schema](https://gitlab.freedesktop.org/wayland/wayland-protocols/-/blob/main/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Wayland window backend](#topic-wayland-window-backend)

<a id="topic-file-and-config-persistence"></a>
## 121. Files, directories, and configuration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is direct application data persistence rather than the Resource import and loading path.

FileAccess and DirAccess read or write project and user paths, while ConfigFile stores section and key values.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="ConfigFile" inherits="RefCounted" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Helper class to handle INI-style files.
	</brief_description>
	<description>
		This helper class can be used to store [Variant] values on the filesystem using INI-style formatting. The stored values are identified by a section and a key:
		[codeblock lang=text]
		[section]
		some_key=42
		string_example="Hello World3D!"
		a_vector=Vector3(1, 0, 2)
		[/codeblock]
		The stored data can be saved to or parsed from a file, though ConfigFile objects can also be used directly without accessing the filesystem.
```

Source: `doc/classes/ConfigFile.xml` — ConfigFile (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ConfigFile

Evidence:
- Code: `doc/classes/FileAccess.xml:2` — FileAccess
- Code: `doc/classes/DirAccess.xml:2` — DirAccess
- Code: `doc/classes/ConfigFile.xml:2` — ConfigFile

<a id="topic-font-driver-modules"></a>
## 122. Font driver modules

**Scope:** Vendored: freetype

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Glyph and face caching](#topic-glyph-caching), [Glyph rasterization](#topic-glyph-rasterization), [PostScript font services](#topic-postscript-font-services).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Each format family has a driver class and generally separates driver, object, loader, and glyph-loading concerns.

Format-specific driver classes expose CFF, CID, PCF, PFR, Type 1, Type 42, Windows FNT, and TrueType implementations to the FreeType module layer.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
FT_CALLBACK_TABLE_DEF
  const FT_Driver_ClassRec  winfnt_driver_class =
  {
    {
      FT_MODULE_FONT_DRIVER        |
      FT_MODULE_DRIVER_NO_OUTLINES,
      sizeof ( FT_DriverRec ),

      "winfonts",
      0x10000L,
      0x20000L,

      NULL, /* module-specific interface */
```

Source: `thirdparty/freetype/src/winfonts/winfnt.c` — winfnt_driver_class (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PFR_FaceRec

Evidence:
- Code: `thirdparty/freetype/src/cff/cffdrivr.h:28` — cff_driver_class
- Code: `thirdparty/freetype/src/truetype/ttdriver.h:28` — tt_driver_class
- Code: `thirdparty/freetype/src/winfonts/winfnt.c:1179` — winfnt_driver_class
- External (official, verified): [Module Management - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-module_management.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [CFF font subsetting](#topic-cff-font-subsetting)
- [FreeType module composition](#topic-freetype-module-composition)

<a id="topic-font-streams"></a>
## 123. Font streams

**Scope:** Vendored: freetype

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Compressed font stream adapters](#topic-compressed-font-stream-adapters), [SFNT table loading](#topic-sfnt-tables).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the byte-access boundary used by the supplied compression and font-reading code.

Font input is represented by an FT_Stream that can deliver bytes through an in-memory base or a read callback.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static FT_Error
  ft_bzip2_file_fill_input( FT_BZip2File  zip )
  {
    bz_stream*  bzstream = &zip->bzstream;
    FT_Stream   stream    = zip->source;
    FT_ULong    size;


    if ( stream->read )
    {
      size = stream->read( stream, stream->pos, zip->input,
                           FT_BZIP2_BUFFER_SIZE );
      if ( size == 0 )
      {
```

Source: `thirdparty/freetype/src/bzip2/ftbzip2.c` — ft_bzip2_file_fill_input (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:249` — ft_bzip2_file_fill_input
- External (official, verified): [BZIP2 Streams - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-bzip2.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Intersection and occlusion filters](#topic-filter-callbacks)
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-font-subsetting"></a>
## 124. Font subsetting

**Scope:** Vendored: harfbuzz

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [CFF font subsetting](#topic-cff-font-subsetting), [OpenType table routing](#topic-opentype-table-routing), [Variable-font table subsetting](#topic-variable-font-subsetting).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The partition contains input, planning, table-specific processing, serialization, and CFF-specific subset code.

HarfBuzz represents caller selections as a subset input and derives an in-memory plan for rewriting a font.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
struct hb_subset_input_t
{
  HB_INTERNAL hb_subset_input_t ();

  ~hb_subset_input_t ()
  {
    sets.~sets_t ();

#ifdef HB_EXPERIMENTAL_API
    for (auto _ : name_table_overrides.values ())
      _.fini ();
#endif
  }
```

Source: `thirdparty/harfbuzz/src/hb-subset-input.hh` — hb_subset_input_t (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_subset_input_t, hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-input.hh:84` — hb_subset_input_t
- Code: `thirdparty/harfbuzz/src/hb-subset-plan.hh:128` — hb_subset_plan_t

<a id="topic-gamepad-motion-fusion"></a>
## 125. Gamepad motion fusion and calibration

**Scope:** Vendored: gamepadmotionhelpers

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation uses vector and quaternion operations, smoothing, stillness checks, and calibration confidence settings.

GamepadMotion maintains orientation, gravity, processed acceleration, and automatic gyro-calibration state from gyro and accelerometer samples.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
// You can ignore anything in namespace GamepadMotionHelpers.
class GamepadMotionSettings;
class GamepadMotion;

namespace GamepadMotionHelpers
{
	struct GyroCalibration
	{
		float X;
		float Y;
		float Z;
		float AccelMagnitude;
		int NumSamples;
	};
```

Source: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — GamepadMotion (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp:11` — GamepadMotion
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp:97` — AutoCalibration
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp:130` — Motion

<a id="topic-buffer-storage"></a>
## 126. Geometry buffer storage

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The public API also defines opaque RTCBuffer handles and buffer-type slots for indices, vertices, attributes, normals, transforms, and related data.

A Buffer is a reference-counted storage abstraction, while BufferView and RawBufferView expose views used to bind geometry data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* Types of buffers */
enum RTCBufferType
{
  RTC_BUFFER_TYPE_INDEX            = 0,
  RTC_BUFFER_TYPE_VERTEX           = 1,
  RTC_BUFFER_TYPE_VERTEX_ATTRIBUTE = 2,
  RTC_BUFFER_TYPE_NORMAL           = 3,
  RTC_BUFFER_TYPE_TANGENT          = 4,
  RTC_BUFFER_TYPE_NORMAL_DERIVATIVE = 5,

  RTC_BUFFER_TYPE_GRID                 = 8,

  RTC_BUFFER_TYPE_FACE                 = 16,
  RTC_BUFFER_TYPE_LEVEL                = 17,
```

Source: `thirdparty/embree/include/embree4/rtcore_buffer.h` — RTCBufferType (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCBuffer, Buffer

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h:11` — RTCBufferType
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h:36` — RTCBuffer
- Code: `thirdparty/embree/kernels/common/buffer.h:366` — BufferView

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [glTF binary accessor decoding and encoding](#topic-gltf-accessors)
- [Shader interface mapping and reflection](#topic-shader-interface-analysis)
- [Web JavaScript bridge](#topic-web-javascript-bridge)
- [WebP image I/O](#topic-webp-image-io)
- [Geometry resources](#topic-geometry-resources)

<a id="topic-geometry-resources"></a>
## 127. Geometry resources

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Geometry families](#topic-geometry-families), [Instancing](#topic-instancing), [Motion blur](#topic-motion-blur), [Primitive intersection](#topic-primitive-intersection), [Primitive references](#topic-primitive-references), [Scene construction and commit](#topic-scene-commit).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Geometry is the scene content abstraction that builders and intersectors address by geometry ID.

A Geometry represents typed primitive data, supplies bounds, and provides the common base used by mesh, curve, point, line, grid, user, and instance implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Scene;
  class Geometry;

  struct GeometryCounts 
  {
    __forceinline GeometryCounts()
      : numFilterFunctions(0),
        numTriangles(0), numMBTriangles(0), 
        numQuads(0), numMBQuads(0), 
        numBezierCurves(0), numMBBezierCurves(0), 
        numLineSegments(0), numMBLineSegments(0), 
        numSubdivPatches(0), numMBSubdivPatches(0), 
        numUserGeometries(0), numMBUserGeometries(0), 
        numInstancesCheap(0), numMBInstancesCheap(0),
```

Source: `thirdparty/embree/kernels/common/geometry.h` — Geometry (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCGeometry, Geometry

Evidence:
- Code: `thirdparty/embree/kernels/common/geometry.h:16` — Geometry
- Code: `thirdparty/embree/kernels/common/scene_triangle_mesh.h:12` — TriangleMesh
- Code: `thirdparty/embree/kernels/common/scene_instance.h:14` — Instance

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Geometry buffer storage](#topic-buffer-storage)
- [BVH traversal](#topic-bvh-traversal)
- [Catmull–Clark patch construction](#topic-catmull-clark-patch-construction)
- [Intersection and occlusion filters](#topic-filter-callbacks)
- [Geometry preprocessing](#topic-geometry-preprocessing)
- [Intersection results](#topic-hit-results)
- [Mesh geometry algorithms](#topic-mesh-geometry-algorithms)
- [Navigation mesh construction](#topic-navigation-mesh-construction)
- [Ray–primitive intersection](#topic-ray-primitive-intersection)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Spatial indexing and ray queries](#topic-spatial-indexing)
- [Transform interpolation](#topic-transform-interpolation)
- [Triangle intersection algorithms](#topic-triangle-intersection-algorithms)
- [2D shapes, tiles, and paths](#topic-two-dimensional-content)
- [Instancing](#topic-instancing)
- [Primitive references](#topic-primitive-references)

<a id="topic-godot-thirdparty-adaptation"></a>
## 128. Godot-specific third-party adaptation

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These files document local deltas from upstream vendored code.

The repository applies VMA allocator integration patches, redirects Vulkan includes to Godot's Vulkan header, adds lazy-allocation statistics, and configures portability macros for bundled dependencies.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#ifdef BIG_ENDIAN_ENABLED
#define WORDS_BIGENDIAN
#endif

#endif /* CONFIG_H */
```

Source: `thirdparty/wslay/config.h` — WORDS_BIGENDIAN (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/vulkan/patches/0001-VKEnumStringHelper-godot-vulkan.patch` — drivers/vulkan/godot_vulkan.h include
- Code: `thirdparty/vulkan/patches/0002-VMA-godot-vulkan.patch` — drivers/vulkan/godot_vulkan.h include
- Code: `thirdparty/vulkan/patches/0003-VMA-add-vmaCalculateLazilyAllocatedBytes.patch:19` — vmaCalculateLazilyAllocatedBytes
- Code: `thirdparty/wslay/config.h:15` — WORDS_BIGENDIAN
- Code: `thirdparty/wslay/patches/0001-msvc-build-fix.patch:11` — SSIZE_T ssize_t

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)

<a id="topic-vulkan-pipeline-configuration"></a>
## 129. Graphics pipeline configuration

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The header represents pipeline setup as a composed configuration record rather than a local rendering algorithm.

`GraphicsPipelineCreateInfo` groups shader stages with separate vertex-input, assembly, viewport, rasterization, multisample, depth-stencil, color-blend, and dynamic-state records.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Concrete structures to recognize

GraphicsPipelineCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:56895` — struct GraphicsPipelineCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:21783` — struct PipelineShaderStageCreateInfo

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Glyph rasterization](#topic-glyph-rasterization)

<a id="topic-graphics-procedure-loading"></a>
## 130. Graphics procedure loading

**Scope:** Vendored: glad

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The source is a vendored generated loader rather than a rendering implementation.

Generated glad loaders resolve OpenGL, EGL, and GLX procedure entry points and track available versions and extensions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct _glad_egl_userptr {
    void *handle;
    PFNEGLGETPROCADDRESSPROC get_proc_address_ptr;
};

static GLADapiproc glad_egl_get_proc(void *vuserptr, const char* name) {
    struct _glad_egl_userptr userptr = *(struct _glad_egl_userptr*) vuserptr;
    GLADapiproc result = NULL;

    result = glad_dlsym_handle(userptr.handle, name);
    if (result == NULL) {
        result = GLAD_GNUC_EXTENSION (GLADapiproc) userptr.get_proc_address_ptr(name);
    }
```

Source: `thirdparty/glad/egl.c` — _glad_egl_userptr (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/glad/gl.c:2446` — _glad_gl_userptr
- Code: `thirdparty/glad/egl.c:343` — _glad_egl_userptr
- External (official, verified): [Khronos Combined OpenGL Registry](https://registry.khronos.org/OpenGL/), accessed 2026-07-16

<a id="topic-hid-device-enumeration"></a>
## 131. HID device enumeration and backends

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Gamepad mapping and classification](#topic-gamepad-mapping).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Linux, macOS, Windows, libusb, and SDL joystick HIDAPI code are present in this partition.

SDL enumerates HID devices into linked device-information records and routes opened devices through platform and driver backends.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct SDL_hid_device_info *SDL_hid_enumerate(unsigned short vendor_id, unsigned short product_id)
{
    struct hid_device_info *driver_devs = NULL;
    struct hid_device_info *usb_devs = NULL;
    struct hid_device_info *raw_devs = NULL;
    struct hid_device_info *dev;
    struct SDL_hid_device_info *devs = NULL, *last = NULL;

    if (SDL_hidapi_refcount == 0 && SDL_hid_init() < 0) {
        return NULL;
    }

    // Collect the available devices
```

Source: `thirdparty/sdl/hidapi/SDL_hidapi.c` — SDL_hid_enumerate (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/hidapi/SDL_hidapi.c:1410` — SDL_hid_enumerate
- Code: `thirdparty/sdl/hidapi/hidapi/hidapi.h:152` — hid_device_info
- Code: `thirdparty/sdl/joystick/hidapi/SDL_hidapijoystick_c.h:68` — SDL_HIDAPI_Device

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Compile-time platform backends](#topic-sdl-platform-backends)

<a id="topic-half-edge-topology"></a>
## 132. Half-edge topology

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Catmull–Clark patch construction](#topic-catmull-clark-patch-construction).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Subdivision connectivity is represented by HalfEdge navigation and by Catmull–Clark rings that inspect neighboring edges, faces, and crease information.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<typename Vertex, typename Vertex_t = Vertex>
    struct __aligned(64) CatmullClark1RingT
  {
    ALIGNED_STRUCT_(64);
    
    int border_index;                                   //!< edge index where border starts
    unsigned int face_valence;                          //!< number of adjacent quad faces
    unsigned int edge_valence;                          //!< number of adjacent edges (2*face_valence)
    float vertex_crease_weight;                         //!< weight of vertex crease (0 if no vertex crease)
    DynamicStackArray<float,16,MAX_RING_FACE_VALENCE> crease_weight; //!< edge crease weights for each adjacent edge
    float vertex_level;                                 //!< maximum level of all adjacent edges
    float edge_level;                                   //!< level of first edge
    unsigned int eval_start_index;                      //!< topology dependent index to start evaluation
    unsigned int eval_unique_identifier;                //!< topology dependent unique identifier for this ring
```

Source: `thirdparty/embree/kernels/subdiv/catmullclark_ring.h` — CatmullClark1RingT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

HalfEdge

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/half_edge.h:10` — HalfEdge
- Code: `thirdparty/embree/kernels/subdiv/catmullclark_ring.h:18` — CatmullClark1RingT

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Catmull–Clark patch construction](#topic-catmull-clark-patch-construction)
- [Halfedge topology](#topic-halfedge-topology)
- [Navigation agents and regions](#topic-navigation-agents)
- [Subdivision surface evaluation](#topic-subdivision-surface-evaluation)

<a id="topic-histograms-and-huffman-codes"></a>
## 133. Histograms and Huffman codes

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The lossless encoder can use multiple histogram groups over image tiles.

Histograms built from backward-reference symbols are clustered and converted into Huffman code lengths, codes, and coded tree headers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static void StoreHuffmanTreeOfHuffmanTreeToBitMask(
    VP8LBitWriter* const bw, const uint8_t* code_length_bitdepth) {
  // RFC 1951 will calm you down if you are worried about this funny sequence.
  // This sequence is tuned from that, but more weighted for lower symbol count,
  // and more spiking histograms.
  static const uint8_t kStorageOrder[CODE_LENGTH_CODES] = {
    17, 18, 0, 1, 2, 3, 4, 5, 16, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  };
  int i;
  // Throw away trailing zeros:
  int codes_to_store = CODE_LENGTH_CODES;
  for (; codes_to_store > 4; --codes_to_store) {
    if (code_length_bitdepth[kStorageOrder[codes_to_store - 1]] != 0) {
```

Source: `thirdparty/libwebp/src/enc/vp8l_enc.c` — StoreHuffmanTree (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VP8LBackwardRefs

Evidence:
- Code: `thirdparty/libwebp/src/enc/histogram_enc.c` — VP8LHistogram, VP8LGetHistoImageSymbols
- Code: `thirdparty/libwebp/src/utils/huffman_encode_utils.h` — HuffmanTreeToken, HuffmanTreeCode
- Code: `thirdparty/libwebp/src/enc/vp8l_enc.c` — StoreHuffmanTree

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [Brotli stream decompression](#topic-brotli-stream-decompression)
- [Bitstream and Huffman decoding](#topic-prefix-code-decoding)

<a id="topic-image-decode-pipeline"></a>
## 134. Image decode pipelines

**Scope:** Vendored: libjpeg-turbo

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [JPEG decompression and coefficient transcoding](#topic-jpeg-decompression-transcoding), [PNG streaming I/O and row transforms](#topic-png-stream-transforms), [WebP chunk parsing, incremental decode, and animation](#topic-webp-riff-decoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the common lifecycle behind the JPEG, PNG, and WebP implementations in this partition.

The vendored image libraries parse encoded input, keep decoder state, and emit raster rows or planes; JPEG has scanline and upsampling modules, PNG applies row transforms, and WebP routes decoded output through VP8 I/O.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
GLOBAL(void)
_jinit_upsampler(j_decompress_ptr cinfo)
{
  my_upsample_ptr upsample;
  int ci;
  jpeg_component_info *compptr;
  boolean need_buffer, do_fancy;
  int h_in_group, v_in_group, h_out_group, v_out_group;

#ifdef D_LOSSLESS_SUPPORTED
  if (cinfo->master->lossless) {
#if BITS_IN_JSAMPLE == 8
    if (cinfo->data_precision > BITS_IN_JSAMPLE || cinfo->data_precision < 2)
#else
```

Source: `thirdparty/libjpeg-turbo/src/jdsample.c` — jinit_upsampler (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JPEG Decompression Context, PNG Information State, WebP Decoder State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jdsample.c` — jinit_upsampler
- Code: `thirdparty/libpng/pngtrans.c:20` — png_set_bgr
- Code: `thirdparty/libwebp/src/dec/io_dec.c` — VP8InitIoInternal
- External (official, unverified (source anchor not found)): [Portable Network Graphics (PNG) Specification (Third Edition)](https://www.w3.org/TR/png-3/), accessed 2026-07-15
- External (official, unverified (source anchor not found)): [WebP Container Specification](https://developers.google.com/speed/webp/docs/riff_container), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Brotli stream decompression](#topic-brotli-stream-decompression)
- [HDR, JPEG, and KTX loading](#topic-image-format-loading)
- [Picture planes and pixel ownership](#topic-input-picture-planes)
- [JPEG decompression and coefficient transcoding](#topic-jpeg-decompression-transcoding)
- [PNG streaming I/O and row transforms](#topic-png-stream-transforms)
- [Raster font rendering](#topic-raster-font-rendering)
- [Raster image input](#topic-raster-image-input)
- [WebP image I/O](#topic-webp-image-io)
- [WebP chunk parsing, incremental decode, and animation](#topic-webp-riff-decoding)

<a id="topic-locale-resolution"></a>
## 135. Locale resolution

**Scope:** Vendored: icu4c

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation exposes both object-oriented LocaleMatcher APIs and lower-level locale parsing code.

ICU parses locale identifiers and matches desired locales against supported locales using LSR values, likely-subtag data, and locale-distance data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct LSR final : public UMemory {
    static constexpr int32_t REGION_INDEX_LIMIT = 1001 + 26 * 26;

    static constexpr int32_t EXPLICIT_LSR = 7;
    static constexpr int32_t EXPLICIT_LANGUAGE = 4;
    static constexpr int32_t EXPLICIT_SCRIPT = 2;
    static constexpr int32_t EXPLICIT_REGION = 1;
    static constexpr int32_t IMPLICIT_LSR = 0;
    static constexpr int32_t DONT_CARE_FLAGS = 0;

    const char *language;
    const char *script;
    const char *region;
```

Source: `thirdparty/icu4c/common/lsr.h` — LSR (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Locale, LSR, LocaleMatcher

Evidence:
- Code: `thirdparty/icu4c/common/unicode/localematcher.h:28` — LocaleMatcher
- Code: `thirdparty/icu4c/common/lsr.h:17` — LSR

<a id="topic-manifold-mesh-data"></a>
## 136. Manifold mesh interchange

**Scope:** Vendored: manifold

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Halfedge topology](#topic-halfedge-topology).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the external mesh-shaped data carrier at the geometry boundary.

MeshGLP carries indexed triangles, per-vertex properties, merge information, and run metadata into Manifold construction.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template <typename Precision, typename I = uint32_t>
struct MeshGLP {
  /// Number of property vertices
  I NumVert() const { return vertProperties.size() / numProp; };
  /// Number of triangles
  I NumTri() const { return triVerts.size() / 3; };
  /// Number of properties per vertex, always >= 3.
  I numProp = 3;
  /// Flat, GL-style interleaved list of all vertex properties: propVal =
  /// vertProperties[vert * numProp + propIdx]. The first three properties are
  /// always the position x, y, z. The stride of the array is numProp.
  std::vector<Precision> vertProperties;
  /// The vertex indices of the three triangle corners in CCW (from the outside)
  /// order, for each triangle.
```

Source: `thirdparty/manifold/include/manifold/manifold.h` — struct MeshGLP (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MeshGLP

Evidence:
- Code: `thirdparty/manifold/include/manifold/manifold.h:113` — struct MeshGLP
- Code: `thirdparty/manifold/src/constructors.cpp` — Manifold constructors accepting MeshGLP

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Running projects and instances](#topic-run-management)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)

<a id="topic-multichannel-distance-fields"></a>
## 137. Multi-channel distance fields

**Scope:** Vendored: msdfgen

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

MSDFgen represents vector shapes, colors their edges, and generates distance-field pixels for one, three, or four channels.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
template <>
class DistancePixelConversion<MultiDistance> {
    DistanceMapping mapping;
public:
    typedef BitmapSection<float, 3> BitmapSectionType;
    inline explicit DistancePixelConversion(DistanceMapping mapping) : mapping(mapping) { }
    inline void operator()(float *pixels, const MultiDistance &distance) const {
        pixels[0] = float(mapping(distance.r));
        pixels[1] = float(mapping(distance.g));
        pixels[2] = float(mapping(distance.b));
    }
};

template <>
```

Source: `thirdparty/msdfgen/core/msdfgen.cpp` — DistancePixelConversion<MultiDistance> (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

msdfgen::Shape, msdfgen::BitmapSection<float, 3>

Evidence:
- Code: `thirdparty/msdfgen/msdfgen.h` — msdfgen public interface
- Code: `thirdparty/msdfgen/core/msdfgen.cpp:26` — DistancePixelConversion<MultiDistance>
- External (primary, verified): [Chlumsky/msdfgen](https://github.com/Chlumsky/msdfgen), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Vector font export](#topic-vector-font-export)

<a id="topic-native-support-algorithms"></a>
## 138. Native support algorithms

**Scope:** Vendored: misc

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The support subtree implements deterministic PCG random state evolution alongside compression, noise, color, audio, and packing algorithms.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
// Source from http://www.pcg-random.org/downloads/pcg-c-basic-0.9.zip
void pcg32_srandom_r(pcg32_random_t* rng, uint64_t initstate, uint64_t initseq)
{
    rng->state = 0U;
    rng->inc = (initseq << 1u) | 1u;
    pcg32_random_r(rng);
    rng->state += initstate;
    pcg32_random_r(rng);
}

// Source from https://github.com/imneme/pcg-c-basic/blob/master/pcg_basic.c
// pcg32_boundedrand_r(rng, bound):
//     Generate a uniformly distributed number, r, where 0 <= r < bound
uint32_t pcg32_boundedrand_r(pcg32_random_t *rng, uint32_t bound) {
```

Source: `thirdparty/misc/pcg.cpp` — pcg32_srandom_r (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

pcg32_random_t

Evidence:
- Code: `thirdparty/misc/pcg.cpp:18` — pcg32_srandom_r
- Code: `thirdparty/misc/FastNoiseLite.h:48` — FastNoiseLite

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Pseudo-random generation](#topic-random-generation)

<a id="topic-navmesh-heightfields"></a>
## 139. Navigation heightfields

**Scope:** Vendored: recastnavigation

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Navigation contours and polygons](#topic-navmesh-contours-polygons).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Recast rasterizes input triangles into a heightfield and converts its spans into a compact heightfield with adjacency.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
bool rcBuildCompactHeightfield(rcContext* context, const int walkableHeight, const int walkableClimb,
                               const rcHeightfield& heightfield, rcCompactHeightfield& compactHeightfield)
{
	rcAssert(context);

	rcScopedTimer timer(context, RC_TIMER_BUILD_COMPACTHEIGHTFIELD);

	const int xSize = heightfield.width;
	const int zSize = heightfield.height;
	const int spanCount = rcGetHeightFieldSpanCount(context, heightfield);

	// Fill in header.
	compactHeightfield.width = xSize;
```

Source: `thirdparty/recastnavigation/Recast/Source/Recast.cpp` — rcBuildCompactHeightfield (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

rcHeightfield, rcCompactHeightfield, rcCompactSpan

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Include/Recast.h:293` — rcHeightfield
- Code: `thirdparty/recastnavigation/Recast/Source/Recast.cpp:403` — rcBuildCompactHeightfield
- External (primary, verified): [recastnavigation/recastnavigation](https://github.com/recastnavigation/recastnavigation), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Compact heightfield representation](#topic-compact-heightfield)

<a id="topic-ogg-pages-and-packets"></a>
## 140. Ogg pages, packets, and bit packing

**Scope:** Vendored: libogg

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Theora block video coding](#topic-theora-block-video-codec), [Vorbis block analysis and synthesis](#topic-vorbis-block-synthesis).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the shared container layer used by the supplied Vorbis and Theora implementations.

The Ogg implementation packs variable-sized words into octet streams and maintains stream state that frames packets into page headers and bodies.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
void oggpack_writeinit(oggpack_buffer *b){
  memset(b,0,sizeof(*b));
  b->ptr=b->buffer=_ogg_malloc(BUFFER_INCREMENT);
  if (!b->buffer)
    return;
  b->buffer[0]='\0';
  b->storage=BUFFER_INCREMENT;
}

void oggpackB_writeinit(oggpack_buffer *b){
  oggpack_writeinit(b);
}
```

Source: `thirdparty/libogg/bitwise.c` — oggpack_buffer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Ogg Stream State, Ogg Packet

Evidence:
- Code: `thirdparty/libogg/bitwise.c:39` — oggpack_buffer
- Code: `thirdparty/libogg/ogg/ogg.h:49` — ogg_stream_state
- External (primary, unverified (source anchor not found)): [Page Multiplexing and Ordering in a Physical Ogg Stream](https://www.xiph.org/ogg/doc/ogg-multiplex.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-openxr-structure-chains"></a>
## 141. OpenXR structure chains

**Scope:** Vendored: openxr

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [OpenXR runtime loading](#topic-openxr-runtime-loading).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

OpenXR input and output structures carry a typed structure identifier and a `next` pointer for extensible structure chains.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
// Create a debug utils messenger if the create structure is in the "next" chain
        const auto *next_header = reinterpret_cast<const XrBaseInStructure *>(info->next);
        const XrDebugUtilsMessengerCreateInfoEXT *dbg_utils_create_info = nullptr;
        while (next_header != nullptr) {
            if (next_header->type == XR_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT) {
                LoaderLogger::LogInfoMessage("xrCreateInstance", "Found XrDebugUtilsMessengerCreateInfoEXT in \'next\' chain.");
                dbg_utils_create_info = reinterpret_cast<const XrDebugUtilsMessengerCreateInfoEXT *>(next_header);
                XrDebugUtilsMessengerEXT messenger;
                result = LoaderTrampolineCreateDebugUtilsMessengerEXT(loader_instance->GetInstanceHandle(), dbg_utils_create_info,
                                                                      &messenger);
                if (XR_FAILED(result)) {
                    return XR_ERROR_VALIDATION_FAILURE;
                }
                loader_instance->SetDefaultDebugUtilsMessenger(messenger);
```

Source: `thirdparty/openxr/src/loader/loader_core.cpp` — reinterpret_cast<const XrBaseInStructure *>(info->next) (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

XrInstanceCreateInfo, XrBaseInStructure

Evidence:
- Code: `thirdparty/openxr/include/openxr/openxr.h:1202` — XrInstanceCreateInfo
- Code: `thirdparty/openxr/src/loader/loader_core.cpp:283` — reinterpret_cast<const XrBaseInStructure *>(info->next)
- External (official, unverified (source anchor not found)): [The OpenXR™ 1.1.60 Specification (with all registered extensions)](https://registry.khronos.org/OpenXR/specs/1.1/html/xrspec.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [OpenXR runtime integration](#topic-openxr-runtime-integration)

<a id="topic-psa-key-lifecycle"></a>
## 142. PSA key lifecycle and storage

**Scope:** Vendored: mbedtls

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation distinguishes volatile identifier ranges and contains dedicated storage and slot-management code.

PSA key records retain configuration-selected algorithms, type, lifetime, identifier, and policy attributes from initialization or import through slot management and optional storage.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
psa_status_t psa_import_key_into_slot(
    const psa_key_attributes_t *attributes,
    const uint8_t *data, size_t data_length,
    uint8_t *key_buffer, size_t key_buffer_size,
    size_t *key_buffer_length, size_t *bits)
{
    psa_status_t status = PSA_ERROR_CORRUPTION_DETECTED;
    psa_key_type_t type = attributes->type;

    /* zero-length keys are never supported. */
    if (data_length == 0) {
        return PSA_ERROR_NOT_SUPPORTED;
    }
```

Source: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto.c` — psa_import_key (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

psa_key_attributes_t

Evidence:
- Code: `thirdparty/mbedtls/tf-psa-crypto/include/psa/crypto_struct.h` — struct psa_key_attributes_s; PSA_KEY_ATTRIBUTES_INIT
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto.c:1762` — psa_import_key
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_slot_management.h` — PSA_KEY_ID_VOLATILE_MIN; PSA_KEY_ID_VOLATILE_MAX
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_storage.c` — persistent key storage implementation
- External (official, verified): [PSA Certified Crypto API 1.1](https://arm-software.github.io/psa-api/crypto/1.1/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)

<a id="topic-curve-and-patch-bases"></a>
## 143. Parametric curve bases

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Subdivision surface evaluation](#topic-subdivision-surface-evaluation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These basis implementations are reused by the subdivision-side curve and patch evaluators.

Bezier, B-spline, and Catmull-Rom basis evaluators generate positions and derivatives from four control points.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
{
  class BezierBasis
  {
  public:

    template<typename T>
      static __forceinline Vec4<T> eval(const T& u) 
    {
      const T t1 = u;
      const T t0 = 1.0f-t1;
      const T B0 = t0 * t0 * t0;
      const T B1 = 3.0f * t1 * (t0 * t0);
      const T B2 = 3.0f * (t1 * t1) * t0;
      const T B3 = t1 * t1 * t1;
```

Source: `thirdparty/embree/kernels/subdiv/bezier_curve.h` — BezierBasis (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/bezier_curve.h:12` — BezierBasis
- Code: `thirdparty/embree/kernels/subdiv/bspline_curve.h:11` — BSplineBasis
- Code: `thirdparty/embree/kernels/subdiv/catmullrom_curve.h:20` — CatmullRomBasis

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GUI controls and layout](#topic-gui-control-layout)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-physics-spaces"></a>
## 144. Physics spaces, bodies, shapes, and joints

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Physics queries and motion tests](#topic-physics-queries).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation provides parallel 2D and 3D API families.

Each physics server models a self-contained space containing bodies, areas, joints, and shapes, and exposes APIs to create and configure them.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="PhysicsServer2D" inherits="Object" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		A server interface for low-level 2D physics access.
	</brief_description>
	<description>
		PhysicsServer2D is the server responsible for all 2D physics. It can directly create and manipulate all physics objects:
		- A [i]space[/i] is a self-contained world for a physics simulation. It contains bodies, areas, and joints. Its state can be queried for collision and intersection information, and several parameters of the simulation can be modified.
		- A [i]shape[/i] is a geometric shape such as a circle, a rectangle, a capsule, or a polygon. It can be used for collision detection by adding it to a body/area, possibly with an extra transformation relative to the body/area's origin. Bodies/areas can have multiple (transformed) shapes added to them, and a single shape can be added to bodies/areas multiple times with different local transformations.
		- A [i]body[/i] is a physical object which can be in static, kinematic, or rigid mode. Its state (such as position and velocity) can be queried and updated. A force integration callback can be set to customize the body's physics.
		- An [i]area[/i] is a region in space which can be used to detect bodies and areas e
…[truncated; see coverage report]
```

Source: `doc/classes/PhysicsServer2D.xml` — PhysicsServer2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PhysicsMaterial, RID

Evidence:
- Code: `doc/classes/PhysicsServer2D.xml:2` — PhysicsServer2D
- Code: `doc/classes/PhysicsServer3D.xml:2` — PhysicsServer3D

<a id="topic-input-picture-planes"></a>
## 145. Picture planes and pixel ownership

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Alpha-plane coding](#topic-alpha-plane-coding), [Encoder configuration](#topic-encoder-configuration), [Image quality metrics](#topic-image-quality-metrics), [Lossless pixel transform pipeline](#topic-lossless-transform-pipeline), [Lossy macroblock encoding](#topic-lossy-macroblock-encoding), [Spatial predictive filters](#topic-spatial-predictive-filters), [YUV/RGB conversion and chroma processing](#topic-yuv-rgb-conversion).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The selected representation determines which pixel fields the encoder treats as native input.

WebPPicture represents one image as either ARGB pixels or YUV(A) planes with dimensions, strides, optional alpha, and private allocation pointers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

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

WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h:34` — struct WebPPicture
- Code: `thirdparty/libwebp/src/enc/iterator_enc.c:136` — VP8IteratorImport
- External (official, verified): [WebP API Documentation](https://developers.google.com/speed/webp/docs/api), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Alpha-plane coding](#topic-alpha-plane-coding)
- [Geometry and transform values](#topic-geometry-transforms)
- [Image decode pipelines](#topic-image-decode-pipeline)
- [YUV/RGB conversion and chroma processing](#topic-yuv-rgb-conversion)

<a id="topic-polygon-clipping"></a>
## 146. Polygon clipping and result trees

**Scope:** Vendored: clipper2

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Integer and scaled floating-point front ends share the ClipperBase execution machinery.

Clipper accepts subject and clip paths, runs a specified clip and fill rule, and builds closed or open paths or a hierarchy.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ClipperD : public ClipperBase {
	private:
		double scale_ = 1.0, invScale_ = 1.0;
#ifdef USINGZ
		ZCallbackD zCallbackD_ = nullptr;
#endif
		void BuildPathsD(PathsD& solutionClosed, PathsD* solutionOpen);
		void BuildTreeD(PolyPathD& polytree, PathsD& open_paths);
	public:
		explicit ClipperD(int precision = 2) : ClipperBase()
		{
			CheckPrecisionRange(precision, error_code_);
			// to optimize scaling / descaling precision
```

Source: `thirdparty/clipper2/include/clipper2/clipper.engine.h` — ClipperD (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PolyPathD

Evidence:
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h` — Clipper64::Execute
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h:520` — ClipperD
- Code: `thirdparty/clipper2/include/clipper2/clipper.offset.h:32` — ClipperOffset

<a id="topic-portable-math-fallbacks"></a>
## 147. Portable mathematical fallbacks

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The bundled code includes trigonometric, logarithmic, exponential, power, and floating-point classification operations.

SDL dispatches mathematical functions to platform library functions when present or to bundled fdlibm-derived implementations otherwise.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* __ieee754_sqrt(x)
 * Return correctly rounded sqrt.
 *           ------------------------------------------
 *	     |  Use the hardware sqrt if you have one |
 *           ------------------------------------------
 * Method:
 *   Bit by bit method using integer arithmetic. (Slow, but portable)
 *   1. Normalization
 *	Scale x to y in [1,4) with even powers of 2:
 *	find an integer k such that  1 <= (y=x*2^(2k)) < 4, then
 *		sqrt(x) = 2^k * sqrt(y)
 *   2. Bit by bit computation
 *	Let q  = sqrt(y) truncated to i bit after binary point (q = 1),
```

Source: `thirdparty/sdl/libm/e_sqrt.c` — __ieee754_sqrt (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/sdl/stdlib/SDL_stdlib.c:27` — SDL_atan
- Code: `thirdparty/sdl/libm/math_libm.h:29` — SDL_uclibc_atan
- Code: `thirdparty/sdl/libm/e_sqrt.c:13` — __ieee754_sqrt

<a id="topic-ray-query"></a>
## 148. Ray query state

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [BVH traversal](#topic-bvh-traversal), [Intersection results](#topic-hit-results), [Motion blur](#topic-motion-blur), [Ray packets and streams](#topic-packet-queries), [Primitive intersection](#topic-primitive-intersection).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The API provides scalar rays plus 4-, 8-, 16-, and compile-time-N packet forms.

A ray query carries origin, direction, near and far distances, time, mask, ID, and flags for intersection or occlusion traversal.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* Ray structure for a single ray */
struct RTC_ALIGN(16) RTCRay
{
  float org_x;        // x coordinate of ray origin
  float org_y;        // y coordinate of ray origin
  float org_z;        // z coordinate of ray origin
  float tnear;        // start of ray segment

  float dir_x;        // x coordinate of ray direction
  float dir_y;        // y coordinate of ray direction
  float dir_z;        // z coordinate of ray direction
  float time;         // time of this ray for motion blur

  float tfar;         // end of ray segment (set to hit distance)
```

Source: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCRay (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCRay, RTCRayHit

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:11` — RTCRay
- Code: `thirdparty/embree/kernels/common/ray.h:15` — RayK
- Code: `thirdparty/embree/kernels/common/ray.h:1203` — RayStreamAOS

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [Motion blur](#topic-motion-blur)
- [Primitive intersection](#topic-primitive-intersection)
- [Collision shapes](#topic-collision-shapes)
- [Motion-blur geometry](#topic-motion-blur-geometry)
- [Physics queries and motion tests](#topic-physics-queries)
- [2D and 3D physics queries](#topic-physics-server-queries)
- [3D physics query contracts](#topic-physics-space-queries)
- [Ray–primitive intersection](#topic-ray-primitive-intersection)
- [SIMD ray packets](#topic-simd-ray-packets)
- [Spatial indexing and ray queries](#topic-spatial-indexing)
- [Triangle intersection algorithms](#topic-triangle-intersection-algorithms)
- [Main loop, OS, and time services](#topic-runtime-loop-time)

<a id="topic-ray-primitive-intersection"></a>
## 149. Ray–primitive intersection

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Triangle intersection algorithms](#topic-triangle-intersection-algorithms).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Embree provides separate kernels for several primitive families while retaining the same hit-candidate flow.

Ray-primitive intersection evaluates rays against triangle, sphere, and rounded-line geometry and forwards valid hit candidates to an epilog.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<int M>
      struct RoundLinearCurveIntersector1
      {
        typedef CurvePrecalculations1 Precalculations;

        template<typename Ray>
        struct ray_tfar {
          Ray& ray;
          __forceinline ray_tfar(Ray& ray) : ray(ray) {}
          __forceinline vfloat<M> operator() () const { return ray.tfar; };
        };
	
        template<typename Ray, typename Epilog>
        static __forceinline bool intersect(const vbool<M>& valid_i,
```

Source: `thirdparty/embree/kernels/geometry/roundline_intersector.h` — RoundLinearCurveIntersector1 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SubGrid, TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h:62` — MoellerTrumboreIntersector1
- Code: `thirdparty/embree/kernels/geometry/sphere_intersector.h:76` — SphereIntersector1
- Code: `thirdparty/embree/kernels/geometry/roundline_intersector.h:650` — RoundLinearCurveIntersector1

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Geometry resources](#topic-geometry-resources)
- [Intersection results](#topic-hit-results)
- [Primitive intersection](#topic-primitive-intersection)
- [Ray query state](#topic-ray-query)

<a id="topic-regex-compile-match"></a>
## 150. Regular-expression compilation and matching

**Scope:** Vendored: pcre2

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Regular-expression JIT code generation](#topic-regex-jit-codegen).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

PCRE2 compiles patterns into code objects and matches subjects with traditional and DFA matching engines.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
The indices for several classes are stored in pcre2_compile.h - these must
be kept in sync with posix_names, posix_name_lengths, posix_class_maps,
and posix_substitutes. */

static const char posix_names[] =
  STRING_alpha0 STRING_lower0 STRING_upper0 STRING_alnum0
  STRING_ascii0 STRING_blank0 STRING_cntrl0 STRING_digit0
  STRING_graph0 STRING_print0 STRING_punct0 STRING_space0
  STRING_word0  STRING_xdigit;

static const uint8_t posix_name_lengths[] = {
  5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 6, 0 };
```

Source: `thirdparty/pcre2/src/pcre2_compile.c` — pcre2_compile (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

pcre2_code, pcre2_match_data

Evidence:
- Code: `thirdparty/pcre2/src/pcre2_compile.c:42` — pcre2_compile
- Code: `thirdparty/pcre2/src/pcre2_dfa_match.c:42` — pcre2_dfa_match
- External (official, unverified (source anchor not found)): [pcre2api man page](https://www.pcre.org/current/doc/html/pcre2api.html), accessed 2026-07-15

<a id="topic-vulkan-render-pass-configuration"></a>
## 151. Render-pass configuration

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied C++ binding also declares the newer `RenderPassCreateInfo2` form.

`RenderPassCreateInfo` combines attachment descriptions, subpass descriptions, and subpass dependencies into one externally exchanged rendering configuration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Concrete structures to recognize

RenderPassCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:141275` — struct RenderPassCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:141984` — struct RenderPassCreateInfo2

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)

<a id="topic-resource-bundle-data"></a>
## 152. Resource-bundle data

**Scope:** Vendored: icu4c

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The code has both C-facing UResourceBundle state and C++ ResourceBundle wrappers.

ICU resource bundles expose strings, binary values, integer vectors, keys, names, and locales over loaded resource data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const uint8_t *ResourceBundle::getBinary(int32_t& len, UErrorCode& status) const {
    return ures_getBinary(fResource, &len, &status);
}

const int32_t *ResourceBundle::getIntVector(int32_t& len, UErrorCode& status) const {
    return ures_getIntVector(fResource, &len, &status);
}

uint32_t ResourceBundle::getUInt(UErrorCode& status) const {
    return ures_getUInt(fResource, &status);
}

int32_t ResourceBundle::getInt(UErrorCode& status) const {
```

Source: `thirdparty/icu4c/common/resbund.cpp` — ResourceBundle::getBinary (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UResourceBundle, ResourceDataValue

Evidence:
- Code: `thirdparty/icu4c/common/uresimp.h:63` — UResourceBundle
- Code: `thirdparty/icu4c/common/resbund.cpp:258` — ResourceBundle::getBinary

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-rigid-body-motion"></a>
## 153. Rigid-body motion

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Constraint solving](#topic-constraints), [Vehicle dynamics](#topic-vehicle-dynamics).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

MotionProperties is the mutable dynamic state associated with movable rigid bodies.

Rigid-body motion uses Body transform state plus MotionProperties to model static, kinematic, and dynamic behavior, permitted degrees of freedom, mass, inertia, and velocities.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void MotionProperties::SetMassProperties(EAllowedDOFs inAllowedDOFs, const MassProperties &inMassProperties)
{
	// Store allowed DOFs
	mAllowedDOFs = inAllowedDOFs;

	// Decompose DOFs
	uint allowed_translation_axis = uint(inAllowedDOFs) & 0b111;
	uint allowed_rotation_axis = (uint(inAllowedDOFs) >> 3) & 0b111;

	// Set inverse mass
	if (allowed_translation_axis == 0)
	{
		// No translation possible
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionProperties.cpp` — MotionProperties::SetMassProperties (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionType.h:10` — EMotionType
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionProperties.cpp:12` — MotionProperties::SetMassProperties
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/AllowedDOFs.h:10` — EAllowedDOFs

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Instancing](#topic-instancing)
- [Motion blur](#topic-motion-blur)

<a id="topic-sdl-runtime-properties"></a>
## 154. Runtime property groups and hints

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Abstract I/O streams](#topic-sdl-iostreams).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation uses properties in core services and stream metadata.

SDL runtime property groups associate named values with an ID, while hints retrieve configuration from property and environment sources.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
const char *SDL_GetHint(const char *name)
{
    if (!name) {
        return NULL;
    }

    const char *result = GetHintEnvironmentVariable(name);

    const SDL_PropertiesID hints = GetHintProperties(false);
    if (hints) {
        SDL_LockProperties(hints);

        SDL_Hint *hint = (SDL_Hint *)SDL_GetPointerProperty(hints, name, NULL);
```

Source: `thirdparty/sdl/SDL_hints.c` — SDL_GetHint (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_PropertiesID

Evidence:
- Code: `thirdparty/sdl/SDL_properties.c:542` — SDL_GetStringProperty
- Code: `thirdparty/sdl/SDL_hints.c:258` — SDL_GetHint
- Code: `thirdparty/sdl/include/SDL3/SDL_properties.h:28` — SDL_PropertiesID

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Editing selection history](#topic-editing-selection-history)
- [managed C# script bridge and source generation](#topic-managed-csharp-scripting)
- [Manifold mesh interchange](#topic-manifold-mesh-data)
- [Reflective property inspection](#topic-property-inspection)
- [Property-list and scene-property helpers](#topic-property-introspection)
- [Class metadata and reflection](#topic-reflection)
- [Reflection metadata](#topic-reflection-metadata)
- [Replicated property synchronization](#topic-replicated-property-synchronization)
- [Runtime class metadata](#topic-runtime-metadata)
- [Scene replication configuration](#topic-scene-replication-configuration)
- [Scene state inspection](#topic-scene-state)
- [Text shaping plans](#topic-text-shaping-plans)

<a id="topic-spirv-intermediate-representation"></a>
## 155. SPIR-V intermediate representation

**Scope:** Background concept

> **Background concept — cited from primary sources.** This foundational idea is exercised by the codebase but not defined in one place, so it is explained from external references rather than a source excerpt.

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Shader reflection](#topic-shader-reflection).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SPIRV-Cross stores these concepts in its common IR structures and compiler accessors.

The shader tooling parses SPIR-V into typed values, variables, blocks, functions, decorations, and entry-point metadata.

### Concrete structures to recognize

SPIRType, SPIRVariable

Evidence:
- Code: `thirdparty/spirv-cross/spirv_common.hpp` — spirv_cross::SPIRType
- Code: `thirdparty/spirv-cross/spirv_common.hpp` — spirv_cross::SPIRVariable
- Code: `thirdparty/spirv-cross/spirv_parser.hpp` — spirv_cross::Parser
- External (official, verified): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Shader programs and material parameters](#topic-shader-programs)
- [SPIR-V generation](#topic-spirv-generation)
- [Shader interface metadata](#topic-shader-interface-metadata)
- [SPIR-V shader reflection](#topic-spirv-shader-reflection)
- [SPIR-V rewriting and optimization](#topic-spirv-rewriting)

<a id="topic-spirv-rewriting"></a>
## 156. SPIR-V rewriting and optimization

**Scope:** Vendored: re-spirv

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

re-spirv parses a word-addressed shader into instructions, functions, blocks, results, decorations, and specializations before optimization.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
struct OptimizerContext {
        const Shader &shader;
        std::vector<uint32_t> &instructionAdjacentListIndices;
        std::vector<uint32_t> &instructionInDegrees;
        std::vector<uint32_t> &instructionOutDegrees;
        std::vector<ListNode> &listNodes;
        std::vector<Resolution> &resolutions;
        std::vector<uint8_t> &optimizedData;
        Options options;
    };

    static void optimizerEliminateInstruction(uint32_t pInstructionIndex, OptimizerContext &rContext) {
        uint32_t *optimizedWords = reinterpret_cast<uint32_t *>(rContext.optimizedData.data());
```

Source: `thirdparty/re-spirv/re-spirv.cpp` — OptimizerContext (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shader, Instruction, Optimizer

Evidence:
- Code: `thirdparty/re-spirv/re-spirv.h:169` — struct Shader
- Code: `thirdparty/re-spirv/re-spirv.cpp:2307` — OptimizerContext

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Shader interface metadata](#topic-shader-interface-metadata)
- [Shader reflection](#topic-shader-reflection)
- [SPIR-V generation](#topic-spirv-generation)
- [SPIR-V intermediate representation](#topic-spirv-intermediate-representation)
- [SPIR-V shader reflection](#topic-spirv-shader-reflection)

<a id="topic-shader-programs"></a>
## 157. Shader programs and material parameters

**Scope:** First-party

**Builds on:** [RenderingDevice descriptors and handles](#topic-rendering-device-resources).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The high-level Shader API is distinct from RDShaderFile and RDShaderSPIRV, which are RenderingDevice-facing representations.

A Shader resource supplies custom shader source, ShaderInclude supplies reusable source fragments, and ShaderMaterial binds a Shader with parameter values.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="RDShaderSPIRV" inherits="Resource" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		SPIR-V intermediate representation as part of an [RDShaderFile] (used by [RenderingDevice]).
	</brief_description>
	<description>
		[RDShaderSPIRV] represents an [RDShaderFile]'s [url=https://www.khronos.org/spir/]SPIR-V[/url] code for various shader stages, as well as possible compilation error messages. SPIR-V is a low-level intermediate shader representation. This intermediate representation is not used directly by GPUs for rendering, but it can be compiled into binary shaders that GPUs can understand. Unlike compiled shaders, SPIR-V is portable across GPU models and driver versions.
		This object is used by [RenderingDevice].
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="get_stage_bytecode" qualifiers="const">
			<return type="PackedByteArray" />
```

Source: `doc/classes/RDShaderSPIRV.xml` — RDShaderSPIRV (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ShaderMaterial, RDPipelineSpecializationConstant

Evidence:
- Code: `doc/classes/Shader.xml:2` — Shader
- Code: `doc/classes/ShaderInclude.xml:2` — ShaderInclude
- Code: `doc/classes/ShaderMaterial.xml:2` — ShaderMaterial
- Code: `doc/classes/RDShaderSPIRV.xml:2` — RDShaderSPIRV

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resources and asset lifecycle](#topic-resources)
- [Shader interface metadata](#topic-shader-interface-metadata)
- [Shader reflection](#topic-shader-reflection)
- [SPIR-V generation](#topic-spirv-generation)
- [SPIR-V intermediate representation](#topic-spirv-intermediate-representation)
- [SPIR-V shader reflection](#topic-spirv-shader-reflection)

<a id="topic-shader-source-compilation"></a>
## 158. Shader source compilation

**Scope:** Vendored: glslang

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Shader interface mapping and reflection](#topic-shader-interface-analysis).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The public API separates an individual shader object from a program that links shader objects.

TShader and TProgram compile shader source after preprocessed tokens into stage intermediates and program-level diagnostics.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
//
class TIntermediate {
public:
    explicit TIntermediate(EShLanguage l, int v = 0, EProfile p = ENoProfile) :
        language(l),
        profile(p), version(v),
        treeRoot(nullptr),
        resources(TBuiltInResource{}),
        numEntryPoints(0), numErrors(0), numPushConstants(0), recursive(false),
        invertY(false),
        dxPositionW(false),
        enhancedMsgs(false),
        debugInfo(false),
        useStorageBuffer(false),
```

Source: `thirdparty/glslang/glslang/MachineIndependent/localintermediate.h` — TIntermediate (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TShader, TProgram

Evidence:
- Code: `thirdparty/glslang/glslang/Public/ShaderLang.h` — glslang::TShader
- Code: `thirdparty/glslang/glslang/Public/ShaderLang.h` — glslang::TProgram
- Code: `thirdparty/glslang/glslang/MachineIndependent/localintermediate.h:95` — TIntermediate
- External (primary, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.html), accessed 2026-07-15
- External (primary, verified): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-16

<a id="topic-shader-language-front-end"></a>
## 159. Shader-language front end

**Scope:** Vendored: glslang

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [SPIR-V generation](#topic-spirv-generation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This implementation provides the internal representation consumed by the SPIR-V lowering code.

The vendored glslang front end models shader types, symbols, source locations, parse context, and a typed intermediate tree.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
//
class TIntermNode {
public:
    POOL_ALLOCATOR_NEW_DELETE(glslang::GetThreadPoolAllocator())

    TIntermNode() { loc.init(); }
    virtual const glslang::TSourceLoc& getLoc() const { return loc; }
    virtual void setLoc(const glslang::TSourceLoc& l) { loc = l; }
    virtual void traverse(glslang::TIntermTraverser*) = 0;
    virtual       glslang::TIntermVariableDecl*  getAsVariableDecl()        { return nullptr; }
    virtual       glslang::TIntermTyped*         getAsTyped()               { return nullptr; }
    virtual       glslang::TIntermOperator*      getAsOperator()            { return nullptr; }
    virtual       glslang::TIntermConstantUnion* getAsConstantUnion()       { return nullptr; }
    virtual       glslang::TIntermAggregate*     getAsAggregate()           { return nullptr; }
```

Source: `thirdparty/glslang/glslang/Include/intermediate.h` — TIntermNode (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

spv::Module, spv::Function, spv::Block

Evidence:
- Code: `thirdparty/glslang/glslang/Include/intermediate.h:44` — TIntermNode
- Code: `thirdparty/glslang/glslang/MachineIndependent/ParseHelper.h:42` — TParseContext
- External (official, verified): [OpenGL / OpenGL ES Reference Compiler](https://www.khronos.org/opengles/sdk/Reference-Compiler/), accessed 2026-07-16

<a id="topic-skeleton-modifiers"></a>
## 160. Skeleton modification and physical-bone simulation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Physical-bone adapters transfer simulated transforms back to visual bones.

2D modifier stacks order Resource-based bone modifications, while SkeletonModifier3D nodes run after AnimationMixer playback and can implement IK, constraints, and skeleton physics.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="SkeletonModification2DPhysicalBones" inherits="SkeletonModification2D" api_type="core" experimental="Physical bones may be changed in the future to perform the position update of [Bone2D] on their own, without needing this resource." xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		A modification that applies the transforms of [PhysicalBone2D] nodes to [Bone2D] nodes.
	</brief_description>
	<description>
		This modification takes the transforms of [PhysicalBone2D] nodes and applies them to [Bone2D] nodes. This allows the [Bone2D] nodes to react to physics thanks to the linked [PhysicalBone2D] nodes.
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="fetch_physical_bones">
			<return type="void" />
			<description>
```

Source: `doc/classes/SkeletonModification2DPhysicalBones.xml` — SkeletonModification2DPhysicalBones (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SkeletonModificationStack2D, SkeletonProfile

Evidence:
- Code: `doc/classes/SkeletonModificationStack2D.xml:2` — SkeletonModificationStack2D
- Code: `doc/classes/SkeletonModifier3D.xml:2` — SkeletonModifier3D
- Code: `doc/classes/SkeletonModification2DPhysicalBones.xml:2` — SkeletonModification2DPhysicalBones

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resources and asset lifecycle](#topic-resources)
- [Running projects and instances](#topic-run-management)
- [Skeletons, animation, and ragdolls](#topic-skeletal-ragdoll)
- [Skeletal modifiers and inverse kinematics](#topic-skeletal-modifiers-and-ik)
- [2D skeleton modification stacks](#topic-skeleton-modification)

<a id="topic-vulkan-swapchain-presentation"></a>
## 161. Swapchain presentation

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Presentation is represented by configuration and submission records in the generated API layer.

The binding pairs swapchain configuration with a presentation request whose wait semaphores, swapchains, and image indices are array inputs.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Concrete structures to recognize

SwapchainCreateInfoKHR, PresentInfoKHR, SwapchainKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:154367` — struct SwapchainCreateInfoKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:135085` — struct PresentInfoKHR

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [Dynamic arrays and dictionaries](#topic-variant-containers)
- [Canvas and viewport presentation](#topic-canvas-and-viewports)
- [Metal drawable presentation](#topic-metal-presentation)

<a id="topic-task-scheduling"></a>
## 162. Task scheduling and synchronization

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The scheduler can be selected between internal, TBB, and PPL implementations through configuration macros.

The internal scheduler represents work as tasks, queues, threads, and thread pools, and provides barriers and atomic helpers for synchronization.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<typename T>
    __forceinline void _atomic_min(std::atomic<T>& aref, const T& bref)
  {
    const T b = bref.load();
    while (true) {
      T a = aref.load();
      if (a <= b) break;
      if (aref.compare_exchange_strong(a,b)) break;
    }
  }

  template<typename T>
    __forceinline void _atomic_max(std::atomic<T>& aref, const T& bref)
  {
```

Source: `thirdparty/embree/common/sys/atomic.h` — _atomic_min (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TaskScheduler, Task, TaskQueue, ThreadPool, BarrierActive

Evidence:
- Code: `thirdparty/embree/common/tasking/taskscheduler.h:6` — TASKING_INTERNAL
- Code: `thirdparty/embree/common/tasking/taskschedulerinternal.h:126` — TaskQueue
- Code: `thirdparty/embree/common/sys/barrier.h:38` — BarrierActive
- Code: `thirdparty/embree/common/sys/atomic.h:39` — _atomic_min

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)
- [Deferred calls and worker tasks](#topic-deferred-execution)
- [Device runtime](#topic-device-runtime)
- [Encoder configuration](#topic-encoder-configuration)
- [Raycast occlusion culling](#topic-raycast-occlusion-culling)
- [Threads and synchronization](#topic-synchronization-primitives)
- [Wayland window backend](#topic-wayland-window-backend)
- [Worker-based parallelism](#topic-worker-parallelism)

<a id="topic-temporal-upscaling-dispatch"></a>
## 163. Temporal upscaling dispatch

**Scope:** Vendored: amd-fsr2

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation contains temporal resource selection and named shader-pass sources.

Vendored FSR2 dispatch code selects alternating frame resources, handles accumulation reset, and issues compute work based on render and display dimensions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
static void fsr2DebugCheckDispatch(FfxFsr2Context_Private* context, const FfxFsr2DispatchDescription* params)
{
    if (params->commandList == nullptr)
    {
        context->contextDescription.fpMessage(FFX_FSR2_MESSAGE_TYPE_ERROR, L"commandList is null");
    }

    if (params->color.resource == nullptr)
    {
        context->contextDescription.fpMessage(FFX_FSR2_MESSAGE_TYPE_ERROR, L"color resource is null");
    }

    if (params->depth.resource == nullptr)
```

Source: `thirdparty/amd-fsr2/ffx_fsr2.cpp` — fsr2Dispatch (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/amd-fsr2/ffx_fsr2.cpp:771` — fsr2Dispatch
- Code: `thirdparty/amd-fsr2/shaders/ffx_fsr2_accumulate_pass.glsl` — accumulation compute pass

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [OpenXR dispatch forwarding](#topic-openxr-dispatch)
- [Resource-backed assets](#topic-resource-assets)

<a id="topic-text-and-localization"></a>
## 164. Text backends and translation domains

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

TextLine and TextParagraph are TextServer abstractions for a single line or paragraph.

TextServerManager registers and selects TextServer implementations, while TranslationServer stores language data in named TranslationDomain collections.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="TextServerManager" inherits="Object" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		A singleton for managing [TextServer] implementations.
	</brief_description>
	<description>
		[TextServerManager] is the API backend for loading, enumerating, and switching [TextServer]s.
		[b]Note:[/b] Switching text server at runtime is possible, but will invalidate all fonts and text buffers. Make sure to unload all controls, fonts, and themes before doing so.
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="add_interface">
			<return type="void" />
```

Source: `doc/classes/TextServerManager.xml` — TextServerManager (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TranslationDomain, TextLine

Evidence:
- Code: `doc/classes/TextServerManager.xml:2` — TextServerManager
- Code: `doc/classes/TranslationServer.xml:2` — TranslationServer
- Code: `doc/classes/TranslationDomain.xml:2` — TranslationDomain

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Translation and locale selection](#topic-localization)

<a id="topic-text-shaping-plans"></a>
## 165. Text shaping plans

**Scope:** Vendored: harfbuzz

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The public shaping entry points accept features and an optional shaper list, while plan creation accepts matching configuration.

A shaping-plan key records segment properties, user features, coordinates, and shaper selection so shaping can use a plan object.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
struct hb_shape_plan_key_t
{
  hb_segment_properties_t  props;

  const hb_feature_t      *user_features;
  unsigned int             num_user_features;

#ifndef HB_NO_OT_SHAPE
  hb_ot_shape_plan_key_t   ot;
#endif

  hb_shape_func_t         *shaper_func;
  const char              *shaper_name;
```

Source: `thirdparty/harfbuzz/src/hb-shape-plan.hh` — hb_shape_plan_key_t (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_shape_plan_t, hb_shape_plan_key_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-shape-plan.hh:35` — hb_shape_plan_key_t
- Code: `thirdparty/harfbuzz/src/hb-shape.h:44` — hb_shape

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)

<a id="topic-texture-block-codecs"></a>
## 166. Texture block codecs

**Scope:** Vendored: etcpak

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

etcpak exposes ETC/EAC and BC block compression plus block-decoding entry points that operate on typed source and destination buffers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void DecodeRGBlock( const void* src, void* dst, size_t width );
void DecodeRGBBlock( const void* src, void* dst, size_t width );
void DecodeRGBABlock( const void* src, void* dst, size_t width );

#endif
```

Source: `thirdparty/etcpak/DecodeRGB.hpp` — DecodeRGBBlock (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/etcpak/ProcessRGB.hpp:6` — CompressEtc1Rgb
- Code: `thirdparty/etcpak/ProcessDxtc.hpp:12` — CompressBc5
- Code: `thirdparty/etcpak/DecodeRGB.hpp:9` — DecodeRGBBlock

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Block texture encoding](#topic-block-texture-encoding)

<a id="topic-texture-compression-pipeline"></a>
## 167. Texture compression pipeline

**Scope:** Vendored: basis_universal

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied implementation includes ETC-family, UASTC, ASTC, GPU texture, and container-output code.

Vendored Basis Universal code separates frontend block/codebook work, backend encoding, and Basis or KTX2 output creation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct basis_compressor_params
	{
		basis_compressor_params() :
			m_compression_level((int)BASISU_DEFAULT_COMPRESSION_LEVEL, 0, (int)BASISU_MAX_COMPRESSION_LEVEL),
			m_selector_rdo_thresh(BASISU_DEFAULT_SELECTOR_RDO_THRESH, 0.0f, 1e+10f),
			m_endpoint_rdo_thresh(BASISU_DEFAULT_ENDPOINT_RDO_THRESH, 0.0f, 1e+10f),
			m_mip_scale(1.0f, .000125f, 4.0f),
			m_mip_smallest_dimension(1, 1, 16384),
			m_etc1s_max_endpoint_clusters(512),
			m_etc1s_max_selector_clusters(512),
			m_etc1s_quality_level(-1),
			m_pack_uastc_ldr_4x4_flags(cPackUASTCLevelDefault),
			m_rdo_uastc_ldr_4x4_quality_scalar(1.0f, 0.001f, 50.0f),
```

Source: `thirdparty/basis_universal/encoder/basisu_comp.h` — basis_compressor (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/basis_universal/encoder/basisu_frontend.h:1` — basisu_frontend
- Code: `thirdparty/basis_universal/encoder/basisu_backend.h` — basisu_backend and basisu_backend_output
- Code: `thirdparty/basis_universal/encoder/basisu_comp.h:581` — basis_compressor

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [KTX2 container transcoding](#topic-ktx2-container-transcoding)
- [Basis file layout](#topic-basis-container-layout)
- [Basis texture transcoding](#topic-basis-transcoding)

<a id="topic-texture-format-description"></a>
## 168. Texture format descriptions

**Scope:** Vendored: libktx

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [KTX texture containers](#topic-ktx-texture-container).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The descriptor is an exchanged binary description of texel layout rather than an opaque GPU-native layout.

Khronos Data Format descriptors represent image-format layout using header-word and sample-word accessors, while helper code creates, interprets, queries, and prints those descriptors.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static uint32_t *writeHeader(int numSamples, int bytes, int suffix,
                             channels_infotype infotype)
{
    uint32_t *DFD = (uint32_t *) malloc(sizeof(uint32_t) *
                                        (1 + KHR_DF_WORD_SAMPLESTART +
                                         numSamples * KHR_DF_WORD_SAMPLEWORDS));
    uint32_t* BDFD = DFD+1;
    DFD[0] = sizeof(uint32_t) *
        (1 + KHR_DF_WORD_SAMPLESTART +
         numSamples * KHR_DF_WORD_SAMPLEWORDS);
    BDFD[KHR_DF_WORD_VENDORID] =
        (KHR_DF_VENDORID_KHRONOS << KHR_DF_SHIFT_VENDORID) |
        (KHR_DF_KHR_DESCRIPTORTYPE_BASICFORMAT << KHR_DF_SHIFT_DESCRIPTORTYPE);
```

Source: `thirdparty/libktx/external/dfdutils/createdfd.c` — writeHeader (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

KTX Level Index Entry

Evidence:
- Code: `thirdparty/libktx/include/KHR/khr_df.h` — khr_word_e
- Code: `thirdparty/libktx/external/dfdutils/createdfd.c:27` — writeHeader
- Code: `thirdparty/libktx/external/dfdutils/interpretdfd.c:64` — interpretDFD
- External (official, verified): [Khronos Data Format Specification](https://registry.khronos.org/DataFormat/specs/1.4/dataformat.1.4.html), accessed 2026-07-16

<a id="topic-sdl-thread-abstractions"></a>
## 169. Thread and synchronization abstractions

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The event queue and device code can use these facilities rather than directly naming every platform primitive.

SDL implements mutex, semaphore, condition, read/write lock, thread, and TLS abstractions with generic, pthread, and Windows backends.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
void SDL_LockMutex(SDL_Mutex *mutex) SDL_NO_THREAD_SAFETY_ANALYSIS // clang doesn't know about NULL mutexes
{
    if (mutex) {
#ifdef FAKE_RECURSIVE_MUTEX
        pthread_t this_thread = pthread_self();
        if (mutex->owner == this_thread) {
            ++mutex->recursive;
        } else {
            /* The order of operations is important.
               We set the locking thread id after we obtain the lock
               so unlocks from other threads will fail.
             */
            const int rc = pthread_mutex_lock(&mutex->id);
```

Source: `thirdparty/sdl/thread/pthread/SDL_sysmutex.c` — SDL_LockMutex (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/sdl/include/SDL3/SDL_mutex.h:296` — SDL_Mutex
- Code: `thirdparty/sdl/thread/pthread/SDL_sysmutex.c:61` — SDL_LockMutex
- Code: `thirdparty/sdl/thread/windows/SDL_sysmutex.c:182` — SDL_CreateMutex

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Compile-time platform backends](#topic-sdl-platform-backends)
- [Deferred calls and worker tasks](#topic-deferred-execution)
- [Native platform adapters](#topic-native-platform-adapters)
- [Task scheduling and synchronization](#topic-task-scheduling)
- [Threads and synchronization](#topic-synchronization-primitives)
- [TLS connection and session state](#topic-tls-connection-state)
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-concurrency"></a>
## 170. Thread synchronization

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The class documentation specifies lifetime constraints to avoid cleanup races and deadlocks.

Mutex synchronizes access to a critical section between Thread instances and is documented as a reentrant binary semaphore.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
Locks this [Mutex], blocks until it is unlocked by the current owner.
				[b]Note:[/b] This function returns without blocking if the thread already has ownership of the mutex.
			</description>
		</method>
		<method name="try_lock">
			<return type="bool" />
			<description>
				Tries locking this [Mutex], but does not block. Returns [code]true[/code] on success, [code]false[/code] otherwise.
				[b]Note:[/b] This function returns [code]true[/code] if the thread already has ownership of the mutex.
			</description>
		</method>
		<method name="unlock">
			<return type="void" />
			<description>
```

Source: `doc/classes/Mutex.xml` — Mutex (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `doc/classes/Mutex.xml:2` — Mutex

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Encoder configuration](#topic-encoder-configuration)
- [Raycast occlusion culling](#topic-raycast-occlusion-culling)
- [Event queue and watches](#topic-sdl-event-queue)
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)
- [Threads and synchronization](#topic-synchronization-primitives)
- [Task scheduling and synchronization](#topic-task-scheduling)
- [Wayland window backend](#topic-wayland-window-backend)
- [Worker-based parallelism](#topic-worker-parallelism)

<a id="topic-mesh-provenance"></a>
## 171. Triangle provenance

**Scope:** Vendored: manifold

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The relation data is used during boolean-result construction and property transfer.

CSG output keeps triangle provenance through TriRef records and MeshRelationD triRef collections so result triangles can be related to source meshes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct TriRef {
  /// The unique ID of the mesh instance of this triangle. If .meshID and .tri
  /// match for two triangles, then they are coplanar and came from the same
  /// face.
  int meshID;
  /// The OriginalID of the mesh this triangle came from. This ID is ideal for
  /// reapplying properties like UV coordinates to the output mesh.
  int originalID;
  /// Probably the triangle index of the original triangle this was part of:
  /// Mesh.triVerts[tri], but it's an input, so just pass it along unchanged.
  int faceID;
  /// Triangles with the same coplanar ID are coplanar.
  int coplanarID;
```

Source: `thirdparty/manifold/src/shared.h` — struct TriRef (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriRef, MeshRelationD

Evidence:
- Code: `thirdparty/manifold/src/shared.h:135` — struct TriRef
- Code: `thirdparty/manifold/src/impl.h` — struct Manifold::Impl::MeshRelationD
- Code: `thirdparty/manifold/src/boolean_result.cpp` — outR.meshRelation_.triRef; struct MapTriRef

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)

<a id="topic-ui-theming"></a>
## 172. UI themes and style boxes

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Panel, PanelContainer, PopupPanel, separators, and controls consume these style definitions.

Theme resources apply reusable colors, constants, fonts, icons, and StyleBoxes across Control and Window branches, while StyleBox defines a visual box treatment.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="PanelContainer" inherits="Container" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		A container that keeps its child controls within the area of a [StyleBox].
	</brief_description>
	<description>
		A container that keeps its child controls within the area of a [StyleBox]. Useful for giving controls an outline.
	</description>
	<tutorials>
		<link title="Using Containers">$DOCS_URL/tutorials/ui/gui_containers.html</link>
		<link title="2D Role Playing Game (RPG) Demo">https://godotengine.org/asset-library/asset/2729</link>
	</tutorials>
	<members>
		<member name="mouse_filter" type="int" setter="set_mouse_filter" getter="get_mouse_filter" overrides="Control" enum="Control.MouseFilter" default="0" />
```

Source: `doc/classes/PanelContainer.xml` — PanelContainer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Theme, StyleBox

Evidence:
- Code: `doc/classes/Theme.xml:2` — Theme
- Code: `doc/classes/StyleBox.xml:2` — StyleBox
- Code: `doc/classes/PanelContainer.xml:2` — PanelContainer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor theming](#topic-editor-theming)
- [GUI controls and layout](#topic-gui-control-layout)
- [Resource-backed assets](#topic-resource-assets)
- [Themes and style boxes](#topic-themes-and-style-boxes)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-upnp-device-discovery"></a>
## 173. UPnP device discovery

**Scope:** Vendored: miniupnpc

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [UPnP port-mapping control](#topic-upnp-port-mapping).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

MiniUPnPc discovers UPnP devices, represents them as a linked device list, and parses an IGD description into URLs and service data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
MINIUPNP_LIBSPEC int
UPNP_GetValidIGD(struct UPNPDev * devlist,
                 struct UPNPUrls * urls,
                 struct IGDdatas * data,
                 char * lanaddr, int lanaddrlen,
                 char * wanaddr, int wanaddrlen)
{
	struct xml_desc {
		char lanaddr[40];
		char wanaddr[40];
		char * xml;
		int size;
		int is_igd;
	} * desc = NULL;
```

Source: `thirdparty/miniupnpc/src/miniupnpc.c` — UPNP_GetValidIGD (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UPNPDev, UPNPUrls, IGDdatas

Evidence:
- Code: `thirdparty/miniupnpc/src/minissdpc.h:29` — getDevicesFromMiniSSDPD
- Code: `thirdparty/miniupnpc/src/miniupnpc.c:505` — UPNP_GetValidIGD

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [UPnP device control](#topic-upnp-device-control)
- [UPnP port-mapping control](#topic-upnp-port-mapping)

<a id="topic-unicode-data"></a>
## 174. Unicode data and properties

**Scope:** Vendored: harfbuzz

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Text break iteration](#topic-break-iteration), [Character-encoding conversion](#topic-character-encoding-conversion), [Identifier spoof checking](#topic-identifier-spoof-checking), [Unicode normalization](#topic-unicode-normalization).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Generated tables are paired with implementations that expose properties and Unicode processing APIs.

The included code stores generated Unicode script, emoji, normalization, bidi, case, and general-character-property data for lookup-oriented services.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
static const hb_script_t _hb_ucd_sc_map[176]=
{
                   HB_SCRIPT_COMMON,              HB_SCRIPT_INHERITED,
                  HB_SCRIPT_UNKNOWN,                 HB_SCRIPT_ARABIC,
                 HB_SCRIPT_ARMENIAN,                HB_SCRIPT_BENGALI,
                 HB_SCRIPT_CYRILLIC,             HB_SCRIPT_DEVANAGARI,
                 HB_SCRIPT_GEORGIAN,                  HB_SCRIPT_GREEK,
                 HB_SCRIPT_GUJARATI,               HB_SCRIPT_GURMUKHI,
                   HB_SCRIPT_HANGUL,                    HB_SCRIPT_HAN,
                   HB_SCRIPT_HEBREW,               HB_SCRIPT_HIRAGANA,
                  HB_SCRIPT_KANNADA,               HB_SCRIPT_KATAKANA,
                      HB_SCRIPT_LAO,                  HB_SCRIPT_LATIN,
                HB_SCRIPT_MALAYALAM,                  HB_SCRIPT_ORIYA,
```

Source: `thirdparty/harfbuzz/src/hb-ucd-table.hh` — _hb_ucd_sc_map (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UCPTrie, hb_unicode_funcs_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-ucd-table.hh:17` — _hb_ucd_sc_map
- Code: `thirdparty/icu4c/common/uchar_props_data.h:14` — propsTrie_index

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Script resources and instances](#topic-scripting)
- [Unicode normalization](#topic-unicode-normalization)

<a id="topic-vi-native-surface-creation"></a>
## 175. VI native surface creation

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is a platform-specific Vulkan surface extension declaration.

The `VK_NN_vi_surface` header declares input for creating a `VkSurfaceKHR` from a Nintendo VI window handle.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
typedef VkFlags VkViSurfaceCreateFlagsNN;
typedef struct VkViSurfaceCreateInfoNN {
    VkStructureType             sType;
    const void*                 pNext;
    VkViSurfaceCreateFlagsNN    flags;
    void*                       window;
} VkViSurfaceCreateInfoNN;

typedef VkResult (VKAPI_PTR *PFN_vkCreateViSurfaceNN)(VkInstance instance, const VkViSurfaceCreateInfoNN* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);

#ifndef VK_NO_PROTOTYPES
#ifndef VK_ONLY_EXPORTED_PROTOTYPES
VKAPI_ATTR VkResult VKAPI_CALL vkCreateViSurfaceNN(
    VkInstance                                  instance,
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VkViSurfaceCreateInfoNN (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VkViSurfaceCreateInfoNN, VkSurfaceKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:27` — VkViSurfaceCreateInfoNN
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:38` — vkCreateViSurfaceNN
- External (official, unverified (source anchor not found)): [vkCreateViSurfaceNN(3) :: Vulkan Documentation Project](https://docs.vulkan.org/refpages/latest/refpages/source/vkCreateViSurfaceNN.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Vulkan extensible records](#topic-vulkan-extensible-records)

<a id="topic-viewport-render-data"></a>
## 176. Viewport render-frame data

**Scope:** First-party

**Builds on:** [RenderingDevice descriptors and handles](#topic-rendering-device-resources).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These are internal rendering-server objects rather than general script-created scene objects.

RenderData and RenderSceneData exist for a viewport frame, while RenderSceneBuffersConfiguration configures buffers recreated when a viewport changes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="RenderData" inherits="Object" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Abstract render data object, holds frame data related to rendering a single frame of a viewport.
	</brief_description>
	<description>
		Abstract render data object, exists for the duration of rendering a single viewport. See also [RenderDataRD], [RenderSceneData], and [RenderSceneDataRD].
		[b]Note:[/b] This is an internal rendering server object. Do not instantiate this class from a script.
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="get_camera_attributes" qualifiers="const">
			<return type="RID" />
```

Source: `doc/classes/RenderData.xml` — RenderData (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RenderSceneBuffersConfiguration, RenderData

Evidence:
- Code: `doc/classes/RenderData.xml:2` — RenderData
- Code: `doc/classes/RenderSceneBuffers.xml:2` — RenderSceneBuffers
- Code: `doc/classes/RenderSceneBuffersConfiguration.xml:2` — RenderSceneBuffersConfiguration

<a id="topic-vulkan-extensible-records"></a>
## 177. Vulkan extensible records

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The convention appears across the generated binding's externally exchanged configuration records.

Many Vulkan input records carry a `pNext` pointer, while the C VI surface record also carries an `sType` discriminator, forming an extensible record convention.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
typedef VkFlags VkViSurfaceCreateFlagsNN;
typedef struct VkViSurfaceCreateInfoNN {
    VkStructureType             sType;
    const void*                 pNext;
    VkViSurfaceCreateFlagsNN    flags;
    void*                       window;
} VkViSurfaceCreateInfoNN;

typedef VkResult (VKAPI_PTR *PFN_vkCreateViSurfaceNN)(VkInstance instance, const VkViSurfaceCreateInfoNN* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);

#ifndef VK_NO_PROTOTYPES
#ifndef VK_ONLY_EXPORTED_PROTOTYPES
VKAPI_ATTR VkResult VKAPI_CALL vkCreateViSurfaceNN(
    VkInstance                                  instance,
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VkViSurfaceCreateInfoNN (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VkViSurfaceCreateInfoNN, GraphicsPipelineCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:27` — VkViSurfaceCreateInfoNN
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:56895` — struct GraphicsPipelineCreateInfo

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [VI native surface creation](#topic-vi-native-surface-creation)

<a id="topic-extension-structure-chains"></a>
## 178. Vulkan extension structure chains

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied platform records use the common Vulkan extensible-structure shape.

The structs expose `sType` and `pNext` fields so extension records can be carried through Vulkan creation and query calls.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
typedef VkFlags VkMetalSurfaceCreateFlagsEXT;
typedef struct VkMetalSurfaceCreateInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkMetalSurfaceCreateFlagsEXT    flags;
    const CAMetalLayer*             pLayer;
} VkMetalSurfaceCreateInfoEXT;

typedef VkResult (VKAPI_PTR *PFN_vkCreateMetalSurfaceEXT)(VkInstance instance, const VkMetalSurfaceCreateInfoEXT* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);

#ifndef VK_NO_PROTOTYPES
#ifndef VK_ONLY_EXPORTED_PROTOTYPES
VKAPI_ATTR VkResult VKAPI_CALL vkCreateMetalSurfaceEXT(
    VkInstance                                  instance,
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_metal.h` — VkMetalSurfaceCreateInfoEXT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VkMetalSurfaceCreateInfoEXT, VkExternalFormatQNX

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_metal.h:33` — VkMetalSurfaceCreateInfoEXT
- Code: `thirdparty/vulkan/include/vulkan/vulkan_screen.h:87` — VkExternalFormatQNX
- External (official, unverified (source anchor not found)): [VkExternalFormatQNX Manual Page](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatQNX.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Virtual implementation extensions](#topic-extensions)

<a id="topic-vulkan-instance-setup"></a>
## 179. Vulkan instance setup

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the outer configuration record for the Vulkan API surface represented by the supplied headers.

The binding models instance setup as an `InstanceCreateInfo` record containing optional application metadata plus enabled layer and extension name arrays.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Concrete structures to recognize

InstanceCreateInfo, ApplicationInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:67360` — struct InstanceCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:6883` — struct ApplicationInfo

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Virtual implementation extensions](#topic-extensions)
- [Instancing](#topic-instancing)
- [Class metadata and reflection](#topic-reflection)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)

<a id="topic-vulkan-video-session-configuration"></a>
## 180. Vulkan video session configuration

**Scope:** Vendored: vulkan

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This partition declares video API data contracts; it does not show a video-processing implementation.

The generated records model a video session through a video profile and standard-header-version information, with separate records for decode and encode operations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Concrete structures to recognize

VideoSessionCreateInfoKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:172949` — struct VideoSessionCreateInfoKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:162293` — struct VideoDecodeInfoKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:169844` — struct VideoEncodeInfoKHR

<a id="topic-websocket-frame-events"></a>
## 181. WebSocket framing and event contexts

**Scope:** Vendored: wslay

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation is organized around incremental frame reads and writes plus event callbacks.

Wslay separates frame-level I/O from event-level message handling through frame contexts, event contexts, callback tables, message sources, and send/control queues.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct wslay_event_context {
  /* config status, bitwise OR of enum wslay_event_config values*/
  uint32_t config;
  /* maximum message length that can be received */
  uint64_t max_recv_msg_length;
  /* 1 if initialized for server, otherwise 0 */
  uint8_t server;
  /* bitwise OR of enum wslay_event_close_status values */
  uint8_t close_status;
  /* status code in received close control frame */
  uint16_t status_code_recv;
  /* status code in sent close control frame */
  uint16_t status_code_sent;
```

Source: `thirdparty/wslay/wslay_event.h` — struct wslay_event_context (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

wslay_event_context

Evidence:
- Code: `thirdparty/wslay/wslay/wslay.h` — wslay_frame_context, wslay_event_context, wslay_event_callbacks
- Code: `thirdparty/wslay/wslay_event.h:80` — struct wslay_event_context
- Code: `thirdparty/wslay/wslay_frame.c` — wslay_frame_recv and wslay_frame_send
- External (primary, unverified (source anchor not found)): [RFC 6455: The WebSocket Protocol](https://www.rfc-editor.org/rfc/rfc6455.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Intersection and occlusion filters](#topic-filter-callbacks)
- [GUI controls and layout](#topic-gui-control-layout)
- [Control-tree user interfaces](#topic-ui-composition)
- [WebSocket multiplayer](#topic-websocket-multiplayer)

<a id="topic-wraparound-network-time"></a>
## 182. Wraparound network time

**Scope:** Vendored: enet

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ENet compares and subtracts time values with an overflow threshold so ordering and differences remain defined across its configured time wraparound.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#define ENET_TIME_OVERFLOW 86400000

#define ENET_TIME_LESS(a, b) ((a) - (b) >= ENET_TIME_OVERFLOW)
#define ENET_TIME_GREATER(a, b) ((b) - (a) >= ENET_TIME_OVERFLOW)
#define ENET_TIME_LESS_EQUAL(a, b) (! ENET_TIME_GREATER (a, b))
#define ENET_TIME_GREATER_EQUAL(a, b) (! ENET_TIME_LESS (a, b))

#define ENET_TIME_DIFFERENCE(a, b) ((a) - (b) >= ENET_TIME_OVERFLOW ? (b) - (a) : (a) - (b))

#endif /* __ENET_TIME_H__ */
```

Source: `thirdparty/enet/enet/time.h` — ENET_TIME_OVERFLOW (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/enet/enet/time.h:8` — ENET_TIME_OVERFLOW
- Code: `thirdparty/enet/enet/time.h:10` — ENET_TIME_LESS
- Code: `thirdparty/enet/enet/time.h:15` — ENET_TIME_DIFFERENCE

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [ENet host and peer transport](#topic-enet-host-peer-transport)
- [ENet transport and multiplayer adaptation](#topic-enet-transport-and-multiplayer)
- [Motion blur](#topic-motion-blur)
- [Main loop, OS, and time services](#topic-runtime-loop-time)

<a id="topic-zip64-archive-io"></a>
## 183. ZIP64 archive I/O

**Scope:** Vendored: minizip

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

MiniZip reads and writes ZIP archives while its I/O API abstracts file opening, seeking, telling, and allocation callbacks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
*/
extern int ZEXPORT zipOpenNewFileInZip4_64(zipFile file, const char* filename, const zip_fileinfo* zipfi,
                                           const void* extrafield_local, uInt size_extrafield_local,
                                           const void* extrafield_global, uInt size_extrafield_global,
                                           const char* comment, int method, int level, int raw,
                                           int windowBits,int memLevel, int strategy,
                                           const char* password, uLong crcForCrypting,
                                           uLong versionMadeBy, uLong flagBase, int zip64) {
    zip64_internal* zi;
    uInt size_filename;
    uInt size_comment;
    uInt i;
    int err = ZIP_OK;
```

Source: `thirdparty/minizip/zip.c` — zipOpenNewFileInZip4_64 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

zlib_filefunc64_32_def

Evidence:
- Code: `thirdparty/minizip/ioapi.h:163` — zlib_filefunc64_32_def
- Code: `thirdparty/minizip/zip.c:1302` — zipOpenNewFileInZip4_64

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [ZIP archive I/O](#topic-zip-archive-io)

<a id="topic-zstandard-stream-codec"></a>
## 184. Zstandard stream compression

**Scope:** Vendored: zstd

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation is divided into common, compress, and decompress subtrees.

The bundled code provides Zstandard stream compression contexts, decompression contexts, dictionaries, entropy tables, hash-based match finders, and optional worker pools.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct ZSTD_DCtx_s
{
    const ZSTD_seqSymbol* LLTptr;
    const ZSTD_seqSymbol* MLTptr;
    const ZSTD_seqSymbol* OFTptr;
    const HUF_DTable* HUFptr;
    ZSTD_entropyDTables_t entropy;
    U32 workspace[HUF_DECOMPRESS_WORKSPACE_SIZE_U32];   /* space needed when building huffman tables */
    const void* previousDstEnd;   /* detect continuity */
    const void* prefixStart;      /* start of current segment */
    const void* virtualStart;     /* virtual start of previous segment if it was just before current one */
    const void* dictEnd;          /* end of previous segment */
    size_t expected;
```

Source: `thirdparty/zstd/decompress/zstd_decompress_internal.h` — ZSTD_DCtx_s (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/zstd/zstd.h` — ZSTD_CCtx, ZSTD_DCtx, ZSTD_CDict, ZSTD_DDict
- Code: `thirdparty/zstd/compress/zstd_compress_internal.h` — ZSTD_CCtx_s and ZSTD_MatchState_t
- Code: `thirdparty/zstd/decompress/zstd_decompress_internal.h:126` — ZSTD_DCtx_s
- Code: `thirdparty/zstd/common/pool.h` — POOL_create and POOL_joinJobs
- External (primary, unverified (source anchor not found)): [RFC 8878: Zstandard Compression and the application/zstd Media Type](https://www.rfc-editor.org/rfc/rfc8878.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-zlib-stream-codec"></a>
## 185. zlib stream compression

**Scope:** Vendored: zlib

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Checksum, compression, and decompression code are separate source modules behind zlib.h.

The bundled zlib sources implement Adler-32 and CRC-32 checksums, DEFLATE compression trees, and inflate state for a public streaming compression interface.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* adler32.c -- compute the Adler-32 checksum of a data stream
 * Copyright (C) 1995-2011, 2016 Mark Adler
 * For conditions of distribution and use, see copyright notice in zlib.h
 */

/* @(#) $Id$ */

#include "zutil.h"

#define BASE 65521U     /* largest prime smaller than 65536 */
#define NMAX 5552
/* NMAX is the largest n such that 255n(n+1)/2 + (n+1)(BASE-1) <= 2^32-1 */
```

Source: `thirdparty/zlib/adler32.c` — adler32 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

z_stream

Evidence:
- Code: `thirdparty/zlib/zlib.h` — z_stream and struct internal_state
- Code: `thirdparty/zlib/adler32.c:1` — adler32
- Code: `thirdparty/zlib/deflate.h` — dyn_ltree, dyn_dtree, bl_tree
- Code: `thirdparty/zlib/inflate.h:82` — struct inflate_state
- External (primary, unverified (source anchor not found)): [RFC 1950: ZLIB Compressed Data Format Specification](https://www.rfc-editor.org/rfc/rfc1950), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Shader interface metadata](#topic-shader-interface-metadata)
- [OpenEXR image decoding](#topic-openexr-image-decoding)

<a id="topic-reflection"></a>
## 186. Class metadata and reflection

**Scope:** First-party

**Builds on:** [Engine object model](#topic-object-model).

**This prepares you for:** [Native extension loading](#topic-native-extensions), [Script resources and instances](#topic-scripting).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Registration also records construction behavior and API exposure for native and extension classes.

ClassDB records Object-derived classes, properties, methods, signals, and constants for runtime lookup.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="ClassDB" inherits="Object" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		A class information repository.
	</brief_description>
	<description>
		Provides access to metadata stored for every available engine class.
		[b]Note:[/b] Script-defined classes with [code]class_name[/code] are not part of [ClassDB], so they will not return reflection data such as a method or property list. However, [GDExtension]-defined classes [i]are[/i] part of [ClassDB], so they will return reflection data.
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="can_instantiate" qualifiers="const">
			<return type="bool" />
```

Source: `doc/classes/ClassDB.xml` — ClassDB (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ClassInfo, Variant

Evidence:
- Code: `core/object/class_db.h` — ClassDB::ClassInfo
- Code: `doc/classes/ClassDB.xml:2` — ClassDB

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Android plugins and signals](#topic-android-plugin-signals)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Editor scene sessions](#topic-editor-scene-sessions)
- [Project filesystem index](#topic-filesystem-project-index)
- [Godot member exposure](#topic-godot-member-exposure)
- [Manifold mesh interchange](#topic-manifold-mesh-data)
- [PNG streaming I/O and row transforms](#topic-png-stream-transforms)
- [Profiling name interning](#topic-profiling-interning)
- [Project settings and feature overrides](#topic-project-settings)
- [Reflection metadata](#topic-reflection-metadata)
- [Asynchronous resource previews](#topic-resource-previews)
- [Managed script registration metadata](#topic-script-registration-metadata)
- [Shader reflection](#topic-shader-reflection)
- [SPIR-V intermediate representation](#topic-spirv-intermediate-representation)
- [SPIR-V shader reflection](#topic-spirv-shader-reflection)
- [Vulkan instance setup](#topic-vulkan-instance-setup)
- [Runtime class metadata](#topic-runtime-metadata)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-variant-containers"></a>
## 187. Dynamic arrays and dictionaries

**Scope:** First-party

**Builds on:** [Dynamic runtime values](#topic-dynamic-variant).

**This prepares you for:** [Generic container infrastructure](#topic-generic-containers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Array and Dictionary store Variant values, while typed validators and typed wrappers constrain their element, key, or value types.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct ArrayPrivate;
struct ContainerType;

class Array {
	mutable ArrayPrivate *_p;
	void _ref(const Array &p_from) const;
	void _unref() const;

public:
	struct ConstIterator {
		_FORCE_INLINE_ const Variant &operator*() const { return *element_ptr; }
		_FORCE_INLINE_ const Variant *operator->() const { return element_ptr; }
```

Source: `core/variant/array.h` — Array (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Array, Dictionary, Variant

Evidence:
- Code: `core/variant/array.h:47` — Array
- Code: `core/variant/dictionary.h:46` — Dictionary
- Code: `core/variant/container_type_validate.h:43` — ContainerTypeValidate
- Code: `core/variant/typed_dictionary.h:36` — TypedDictionary

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Text break iteration](#topic-break-iteration)
- [Ray packets and streams](#topic-packet-queries)
- [Swapchain presentation](#topic-vulkan-swapchain-presentation)

<a id="topic-dynamic-invocation-signals"></a>
## 188. Dynamic invocation and signals

**Scope:** First-party

**Builds on:** [Dynamic runtime values](#topic-dynamic-variant), [Object identity and lifetime](#topic-object-identity-lifetime).

**This prepares you for:** [Deferred calls and worker tasks](#topic-deferred-execution), [Undo and redo actions](#topic-undo-redo).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The object and callable layers call Object methods and signal handlers with Variant argument arrays, including bound and unbound callable forms.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Variant;
class CallableCustom;

// This is an abstraction of things that can be called.
// It is used for signals and other cases where efficient calling of functions
// is required. It is designed for the standard case (object and method)
// but can be optimized or customized.

// Enforce 16 bytes with `alignas` to avoid arch-specific alignment issues on x86 vs armv7.

class Callable {
	alignas(8) StringName method;
	union {
		uint64_t object = 0;
```

Source: `core/variant/callable.h` — Callable (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Object, Variant

Evidence:
- Code: `core/object/object.h` — Object::callp
- Code: `core/variant/callable.h:48` — Callable
- Code: `core/variant/callable_bind.h:36` — CallableCustomBind

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class-reference generation](#topic-class-reference-generation)
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Signal awaitability](#topic-signal-awaitability)
- [Engine object model](#topic-object-model)

<a id="topic-expression-evaluation"></a>
## 189. Expression parsing and evaluation

**Scope:** First-party

**Builds on:** [Dynamic runtime values](#topic-dynamic-variant).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Expression defines token and expression-node types and evaluates expression nodes against Variant inputs and built-in functions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
Variant Expression::execute(const Array &p_inputs, Object *p_base, bool p_show_error, bool p_const_calls_only) {
	ERR_FAIL_COND_V_MSG(error_set, Variant(), vformat("There was previously a parse error: %s.", error_str));

	execution_error = false;
	Variant output;
	String error_txt;
	bool err = _execute(p_inputs, p_base, root, output, p_const_calls_only, error_txt);
	if (err) {
		execution_error = true;
		error_str = error_txt;
		ERR_FAIL_COND_V_MSG(p_show_error, Variant(), error_str);
	}
```

Source: `core/math/expression.cpp` — Expression::execute (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant

Evidence:
- Code: `core/math/expression.h:35` — Expression
- Code: `core/math/expression.h:247` — Expression::ENode
- Code: `core/math/expression.cpp:1494` — Expression::execute

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Scene graph and lifecycle](#topic-scene-tree)

<a id="topic-mesh-geometry-algorithms"></a>
## 190. Mesh geometry algorithms

**Scope:** First-party

**Builds on:** [Geometry and transform values](#topic-geometry-transforms).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation uses geometry primitives to build Delaunay triangulations, convex hulls, quick hulls, and polygon triangulations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// Ole Kniemeyer, MAXON Computer GmbH
class ConvexHullComputer {
public:
	class Edge {
	private:
		int32_t next = 0;
		int32_t reverse = 0;
		int32_t target_vertex = 0;

		friend class ConvexHullComputer;

	public:
		int32_t get_next_relative() const {
			return next;
```

Source: `core/math/convex_hull.h` — ConvexHullComputer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMesh, AABB

Evidence:
- Code: `core/math/delaunay_2d.h:36` — Delaunay2D
- Code: `core/math/delaunay_3d.h:44` — Delaunay3D
- Code: `core/math/convex_hull.h:53` — ConvexHullComputer
- Code: `core/math/triangulate.h:40` — Triangulate

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry resources](#topic-geometry-resources)

<a id="topic-packed-resource-files"></a>
## 191. Packed resource files

**Scope:** First-party

**Builds on:** [Resource loading and caching](#topic-resource-loading).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

PackedData is the virtual-file index behind pack-backed file access.

Packed data records pack paths, offsets, sizes, hashes, encryption flags, bundle flags, and delta flags, then supplies files to the resource-loading service through PackSource.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class PackedData {
	friend class FileAccessPack;
	friend class DirAccessPack;
	friend class PackSource;

public:
	struct PackedFile {
		String pack;
		uint64_t offset; //if offset is ZERO, the file was ERASED
		uint64_t size;
		uint8_t md5[16];
		PackSource *src = nullptr;
		bool encrypted;
```

Source: `core/io/file_access_pack.h` — PackedData (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PackedData::PackedFile

Evidence:
- Code: `core/io/file_access_pack.h:64` — PackedData
- Code: `core/io/file_access_pack.h:164` — PackedData::PackedFile
- Code: `core/io/file_access_pack.h:62` — PackSource

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-profiling-interning"></a>
## 192. Profiling name interning

**Scope:** First-party

**Builds on:** [Strings, interned names, and node paths](#topic-string-names-paths).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The profiling implementation interns StringName metadata and source-location data for tracing backends.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
struct TracyInternTable {
	constexpr static uint32_t TABLE_BITS = 16;
	constexpr static uint32_t TABLE_LEN = 1 << TABLE_BITS;
	constexpr static uint32_t TABLE_MASK = TABLE_LEN - 1;

	static inline BinaryMutex mutex;

	static inline SourceLocationInternData *source_location_table[TABLE_LEN];
	static inline PagedAllocator<SourceLocationInternData> source_location_allocator;

	static inline StringInternData *string_table[TABLE_LEN];
	static inline PagedAllocator<StringInternData> string_allocator;
};
```

Source: `core/profiling/profiling.cpp` — TracyInternTable (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

StringName

Evidence:
- Code: `core/profiling/profiling.cpp:71` — TracyInternTable
- Code: `core/profiling/profiling.cpp:115` — intern_source_location

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Compile-time platform backends](#topic-sdl-platform-backends)

<a id="topic-reflection-metadata"></a>
## 193. Reflection metadata

**Scope:** First-party

**Builds on:** [Object identity and lifetime](#topic-object-identity-lifetime).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ClassDB, GDType, MethodInfo, and PropertyInfo describe Object classes, inheritance, methods, properties, constants, enums, and signals.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

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

Object, GDType, MethodInfo, PropertyInfo

Evidence:
- Code: `core/object/class_db.h:62` — ClassDB
- Code: `core/object/gdtype.h:38` — GDType
- Code: `core/object/method_info.h:48` — MethodInfo
- Code: `core/object/property_info.h:124` — PropertyInfo

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)
- [Class metadata and reflection](#topic-reflection)
- [Native extension loading](#topic-native-extensions)
- [Runtime class metadata](#topic-runtime-metadata)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-resource-formats"></a>
## 194. Resource formats and serialization

**Scope:** First-party

**Builds on:** [Resource loading and caching](#topic-resource-loading).

**This prepares you for:** [Resource and server identifiers](#topic-resource-identifiers), [Scene state inspection](#topic-scene-state), [Texture resources and fallback placeholders](#topic-textures-and-placeholders), [Tile libraries, cells, and patterns](#topic-tile-libraries).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Format-specific code is separated from the common resource-loading service.

The resource-loading service uses binary, JSON, image, crypto, and extension resource-format loaders and savers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GDExtensionResourceLoader : public ResourceFormatLoader {
	GDSOFTCLASS(GDExtensionResourceLoader, ResourceFormatLoader);

public:
	static Error load_gdextension_resource(const String &p_path, Ref<GDExtension> &p_extension);

	virtual Ref<Resource> load(const String &p_path, const String &p_original_path, Error *r_error, bool p_use_sub_threads = false, float *r_progress = nullptr, CacheMode p_cache_mode = CACHE_MODE_REUSE) override;
	virtual void get_recognized_extensions(List<String> *p_extensions) const override;
	virtual bool handles_type(const String &p_type) const override;
	virtual String get_resource_type(const String &p_path) const override;
#ifdef TOOLS_ENABLED
	virtual void get_classes_used(const String &p_path, HashSet<StringName> *r_classes) override;
#endif // TOOLS_ENABLED
```

Source: `core/extension/gdextension_resource_format.h` — GDExtensionResourceLoader (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource

Evidence:
- Code: `core/io/resource_format_binary.h:90` — ResourceFormatLoaderBinary
- Code: `core/io/json_resource_format.h:37` — ResourceFormatLoaderJSON
- Code: `core/io/image_resource_format.h:36` — ResourceFormatLoaderImage
- Code: `core/extension/gdextension_resource_format.h:36` — GDExtensionResourceLoader

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Cryptographic keys, hashing, and TLS support](#topic-cryptography)
- [Virtual implementation extensions](#topic-extensions)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-script-hosting"></a>
## 195. Script languages and instances

**Scope:** First-party

**Builds on:** [Object identity and lifetime](#topic-object-identity-lifetime).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The scripting layer attaches Object-facing script instances, manages registered script languages, and provides extension-backed script implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ScriptInstance {
public:
	virtual bool set(const StringName &p_name, const Variant &p_value) = 0;
	virtual bool get(const StringName &p_name, Variant &r_ret) const = 0;
	virtual void get_property_list(List<PropertyInfo> *p_properties) const = 0;
	virtual Variant::Type get_property_type(const StringName &p_name, bool *r_is_valid = nullptr) const = 0;
	virtual void validate_property(PropertyInfo &p_property) const = 0;

	virtual bool property_can_revert(const StringName &p_name) const = 0;
	virtual bool property_get_revert(const StringName &p_name, Variant &r_ret) const = 0;

	virtual Object *get_owner() { return nullptr; }
	virtual void get_property_state(List<Pair<StringName, Variant>> &r_state);
```

Source: `core/object/script_instance.h` — ScriptInstance (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Object, Variant

Evidence:
- Code: `core/object/script_language.h:50` — ScriptServer
- Code: `core/object/script_language.h:40` — ScriptLanguage
- Code: `core/object/script_instance.h:38` — ScriptInstance
- Code: `core/object/script_language_extension.h:758` — ScriptInstanceExtension

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Virtual implementation extensions](#topic-extensions)
- [Engine object model](#topic-object-model)
- [Script resources and instances](#topic-scripting)

<a id="topic-spatial-indexing"></a>
## 196. Spatial indexing and ray queries

**Scope:** First-party

**Builds on:** [Geometry and transform values](#topic-geometry-transforms).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Spatial trees use geometry bounds to maintain items, cull them, refit nodes, and accelerate triangle-mesh and static-ray queries.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template <typename T, int NUM_TREES, int MAX_CHILDREN, int MAX_ITEMS, typename USER_PAIR_TEST_FUNCTION = BVH_DummyPairTestFunction<T>, typename USER_CULL_TEST_FUNCTION = BVH_DummyCullTestFunction<T>, bool USE_PAIRS = false, typename BOUNDS = AABB, typename POINT = Vector3>
class BVH_Tree {
	friend class BVH;

#include "core/math/bvh_pair.inc"
#include "core/math/bvh_structs.inc"

public:
	BVH_Tree() {
		for (int n = 0; n < NUM_TREES; n++) {
			_root_node_id[n] = BVHCommon::INVALID;
		}

		// disallow zero leaf ids
```

Source: `core/math/bvh_tree.h` — BVH_Tree (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMesh, AABB

Evidence:
- Code: `core/math/bvh_tree.h:169` — BVH_Tree
- Code: `core/math/dynamic_bvh.h:56` — DynamicBVH
- Code: `core/math/triangle_mesh.h:36` — TriangleMesh
- Code: `core/math/static_raycaster.h:35` — StaticRaycaster

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Geometry resources](#topic-geometry-resources)
- [Primitive references](#topic-primitive-references)
- [Ray query state](#topic-ray-query)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Allocation and reference ownership](#topic-allocation)
- [BVH split heuristics](#topic-builder-heuristics)
- [BVH construction](#topic-bvh-construction)
- [BVH traversal](#topic-bvh-traversal)
- [Motion blur](#topic-motion-blur)
- [native 2D collision pipeline](#topic-physics-2d-collision-pipeline)
- [Scene construction and commit](#topic-scene-commit)

<a id="topic-transform-interpolation"></a>
## 197. Transform interpolation

**Scope:** First-party

**Builds on:** [Geometry and transform values](#topic-geometry-transforms).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

TransformInterpolator and InterpolatedProperty calculate intermediate geometry transforms and values.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class TransformInterpolator {
public:
	enum Method {
		INTERP_LERP,
		INTERP_SLERP,
		INTERP_SCALED_SLERP,
	};

private:
	_FORCE_INLINE_ static bool _sign(real_t p_val) { return p_val >= 0; }
	static real_t _vec3_sum(const Vector3 &p_pt) { return p_pt.x + p_pt.y + p_pt.z; }
	static real_t _vec3_normalize(Vector3 &p_vec);
	_FORCE_INLINE_ static bool _vec3_is_equal_approx(const Vector3 &p_a, const Vector3 &p_b, real_t p_tolerance) {
```

Source: `core/math/transform_interpolator.h` — TransformInterpolator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Transform3D

Evidence:
- Code: `core/math/transform_interpolator.h:51` — TransformInterpolator
- Code: `core/templates/interpolated_property.h:45` — InterpolatedProperty

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry resources](#topic-geometry-resources)
- [MetalFX scaling and interpolation](#topic-metalfx-scaling)

<a id="topic-localization"></a>
## 198. Translation and locale selection

**Scope:** First-party

**Builds on:** [Strings, interned names, and node paths](#topic-string-names-paths).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Translation catalogs map StringName keys to messages, TranslationDomain selects applicable catalogs, and TranslationServer manages locale, fallback, and plural-rule services.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class PluralRules : public Object {
	GDSOFTCLASS(PluralRules, Object);

	mutable LRUCache<int, int> cache;

	// These two fields are initialized in the constructor.
	const int nplurals;
	const String plural;

	// Cache temporary variables related to `evaluate()` to make it faster.
	class EQNode : public RefCounted {
		GDSOFTCLASS(EQNode, RefCounted);
```

Source: `core/string/plural_rules.h` — PluralRules (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Translation, TranslationDomain, StringName

Evidence:
- Code: `core/string/translation.h` — Translation::MessageKey
- Code: `core/string/translation_domain.h:37` — TranslationDomain
- Code: `core/string/translation_server.h:36` — TranslationServer
- Code: `core/string/plural_rules.h:38` — PluralRules

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Text backends and translation domains](#topic-text-and-localization)

<a id="topic-variant-text-serialization"></a>
## 199. Variant text parsing and writing

**Scope:** First-party

**Builds on:** [Dynamic runtime values](#topic-dynamic-variant), [Strings, interned names, and node paths](#topic-string-names-paths).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

VariantParser and VariantWriter serialize Variant values as String text through stream, token, tag, and resource-parser abstractions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
Error VariantParser::parse_value(Token &r_token, Variant &r_value, Stream *p_stream, int &r_line, String &r_err_str, ResourceParser *p_res_parser) {
	if (r_token.type == TK_CURLY_BRACKET_OPEN) {
		Dictionary d;
		Error err = _parse_dictionary(d, p_stream, r_line, r_err_str, p_res_parser);
		if (err) {
			return err;
		}
		r_value = d;
		return OK;
	} else if (r_token.type == TK_BRACKET_OPEN) {
		Array a;
		Error err = _parse_array(a, p_stream, r_line, r_err_str, p_res_parser);
		if (err) {
```

Source: `core/variant/variant_parser.cpp` — VariantParser::parse (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant, Array, Dictionary

Evidence:
- Code: `core/variant/variant_parser.h:37` — VariantParser
- Code: `core/variant/variant_parser.h:157` — VariantWriter
- Code: `core/variant/variant_parser.cpp:1963` — VariantParser::parse

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Ray packets and streams](#topic-packet-queries)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-navigation-queries"></a>
## 200. 2D and 3D navigation queries

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

2D uses Vector2 paths and 3D uses Vector3 paths.

RefCounted navigation query parameter and result objects exchange path points, path types, and path-owner identifiers with the 2D and 3D navigation server APIs.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NavigationPathQueryParameters2D : public RefCounted {
	GDCLASS(NavigationPathQueryParameters2D, RefCounted);

protected:
	static void _bind_methods();

public:
	enum PathfindingAlgorithm {
		PATHFINDING_ALGORITHM_ASTAR = NavigationEnums2D::PATHFINDING_ALGORITHM_ASTAR,
	};

	enum PathPostProcessing {
		PATH_POSTPROCESSING_CORRIDORFUNNEL = NavigationEnums2D::PATH_POSTPROCESSING_CORRIDORFUNNEL,
```

Source: `servers/navigation_2d/navigation_path_query_parameters_2d.h` — class NavigationPathQueryParameters2D : public RefCounted (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

NavigationPathQueryResult2D

Evidence:
- Code: `servers/navigation_2d/navigation_path_query_parameters_2d.h:37` — class NavigationPathQueryParameters2D : public RefCounted
- Code: `servers/navigation_2d/navigation_path_query_result_2d.h` — NavigationPathQueryResult2D path getters
- Code: `servers/navigation_3d/navigation_path_query_parameters_3d.h:37` — class NavigationPathQueryParameters3D : public RefCounted
- Code: `servers/navigation_3d/navigation_path_query_result_3d.h` — NavigationPathQueryResult3D path getters

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Navigation agents and regions](#topic-navigation-agents)
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Engine object model](#topic-object-model)

<a id="topic-physics-server-queries"></a>
## 201. 2D and 3D physics queries

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The 2D and 3D partitions mirror this server-plus-query arrangement.

Physics servers expose direct body and space state APIs, while RefCounted query parameter and result objects package ray, point, shape, and motion requests.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class PhysicsDirectSpaceState2D : public Object {
	GDCLASS(PhysicsDirectSpaceState2D, Object);

	Dictionary _intersect_ray(RequiredParam<PhysicsRayQueryParameters2D> rp_ray_query);
	TypedArray<Dictionary> _intersect_point(RequiredParam<PhysicsPointQueryParameters2D> rp_point_query, int p_max_results = 32);
	TypedArray<Dictionary> _intersect_shape(RequiredParam<PhysicsShapeQueryParameters2D> rp_shape_query, int p_max_results = 32);
	Vector<real_t> _cast_motion(RequiredParam<PhysicsShapeQueryParameters2D> rp_shape_query);
	TypedArray<Vector2> _collide_shape(RequiredParam<PhysicsShapeQueryParameters2D> rp_shape_query, int p_max_results = 32);
	Dictionary _get_rest_info(RequiredParam<PhysicsShapeQueryParameters2D> rp_shape_query);

protected:
	static void _bind_methods();
```

Source: `servers/physics_2d/direct_states/physics_direct_space_state_2d.h` — class PhysicsDirectSpaceState2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `servers/physics_2d/physics_server_2d.h:42` — class PhysicsServer2D
- Code: `servers/physics_2d/direct_states/physics_direct_space_state_2d.h:39` — class PhysicsDirectSpaceState2D
- Code: `servers/physics_2d/queries/physics_ray_query_parameters_2d.h:39` — class PhysicsRayQueryParameters2D : public RefCounted
- Code: `servers/physics_3d/physics_server_3d.h:51` — class PhysicsServer3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision shapes](#topic-collision-shapes)
- [Motion blur](#topic-motion-blur)
- [Object identity and lifetime](#topic-object-identity-lifetime)
- [Engine object model](#topic-object-model)
- [Ray query state](#topic-ray-query)

<a id="topic-two-dimensional-physics"></a>
## 202. 2D physics nodes

**Scope:** First-party

**Builds on:** [Scene graph nodes](#topic-scene-graph).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The hierarchy includes static, rigid, character, animatable, and joint node types.

2D physics scene graph nodes provide collision objects, bodies, joints, casts, and collision results.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CollisionObject2D : public Node2D {
	GDCLASS(CollisionObject2D, Node2D);

public:
	static constexpr AncestralClass static_ancestral_class = AncestralClass::COLLISION_OBJECT_2D;

	enum DisableMode {
		DISABLE_MODE_REMOVE,
		DISABLE_MODE_MAKE_STATIC,
		DISABLE_MODE_KEEP_ACTIVE,
	};

private:
```

Source: `scene/2d/physics/collision_object_2d.h` — CollisionObject2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CollisionObject2D, RigidBody2D, KinematicCollision2D

Evidence:
- Code: `scene/2d/physics/collision_object_2d.h:39` — CollisionObject2D
- Code: `scene/2d/physics/rigid_body_2d.h:39` — RigidBody2D
- Code: `scene/2d/physics/kinematic_collision_2d.h:39` — KinematicCollision2D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [native 2D collision pipeline](#topic-physics-2d-collision-pipeline)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-three-dimensional-physics"></a>
## 203. 3D physics nodes

**Scope:** First-party

**Builds on:** [Scene graph nodes](#topic-scene-graph).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The hierarchy includes static, rigid, character, physical-bone, soft-body, and vehicle implementations.

3D physics scene graph nodes provide collision objects, bodies, joints, casts, soft bodies, and vehicles.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CollisionObject3D : public Node3D {
	GDCLASS(CollisionObject3D, Node3D);

public:
	static constexpr AncestralClass static_ancestral_class = AncestralClass::COLLISION_OBJECT_3D;

	enum DisableMode {
		DISABLE_MODE_REMOVE,
		DISABLE_MODE_MAKE_STATIC,
		DISABLE_MODE_KEEP_ACTIVE,
	};

private:
```

Source: `scene/3d/physics/collision_object_3d.h` — CollisionObject3D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CollisionObject3D, RigidBody3D, SoftBody3D

Evidence:
- Code: `scene/3d/physics/collision_object_3d.h:38` — CollisionObject3D
- Code: `scene/3d/physics/rigid_body_3d.h:38` — RigidBody3D
- Code: `scene/3d/physics/soft_body_3d.h:39` — SoftBody3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [native 3D collision pipeline](#topic-physics-3d-collision-pipeline)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-particle-systems"></a>
## 204. CPU and GPU particle systems

**Scope:** First-party

**Builds on:** [Scene graph nodes](#topic-scene-graph).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module also includes 3D GPU collision and attractor node families.

2D and 3D particle nodes provide separate CPU and GPU particle-system implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CPUParticles2D : public Node2D {
private:
	GDCLASS(CPUParticles2D, Node2D);

public:
	enum DrawOrder {
		DRAW_ORDER_INDEX,
		DRAW_ORDER_LIFETIME,
	};

	enum Parameter {
		PARAM_INITIAL_LINEAR_VELOCITY,
		PARAM_ANGULAR_VELOCITY,
```

Source: `scene/2d/cpu_particles_2d.h` — CPUParticles2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CPUParticles2D, GPUParticles3D

Evidence:
- Code: `scene/2d/cpu_particles_2d.h:39` — CPUParticles2D
- Code: `scene/3d/gpu_particles_3d.h:38` — GPUParticles3D
- Code: `scene/3d/gpu_particles_collision_3d.h:38` — GPUParticlesCollision3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Renderer-owned resource storage](#topic-renderer-storage)

<a id="topic-platform-presentation"></a>
## 205. Display, camera, video, and movie services

**Scope:** First-party

**Builds on:** [Canvas and viewport presentation](#topic-canvas-and-viewports).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

A headless DisplayServer implementation is also present.

DisplayServer abstracts display functionality, CameraServer manages CameraFeed objects, VideoStreamPlayer presents video in a Control, and MovieWriter emits movie output.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class VideoStreamPlayer : public Control {
	GDCLASS(VideoStreamPlayer, Control);

	Ref<VideoStreamPlayback> playback;
	Ref<VideoStream> stream;

	int sp_get_channel_count() const;
	bool mix(AudioFrame *p_buffer, int p_frames);

	Ref<Texture2D> texture;
	Size2 texture_size;
	void texture_changed(const Ref<Texture2D> &p_texture);
```

Source: `scene/gui/video_stream_player.h` — class VideoStreamPlayer : public Control (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `servers/display/display_server.h:62` — class DisplayServer
- Code: `servers/display/display_server_headless.h:38` — class DisplayServerHeadless : public DisplayServer
- Code: `servers/camera/camera_server.h:48` — class CameraServer
- Code: `scene/gui/video_stream_player.h:37` — class VideoStreamPlayer : public Control
- Code: `servers/movie_writer/movie_writer.h:37` — class MovieWriter

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GUI controls and layout](#topic-gui-control-layout)
- [Platform display and window adaptation](#topic-platform-display-adaptation)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-engine-bootstrap"></a>
## 206. Engine bootstrap

**Scope:** First-party

**Builds on:** [Runtime configuration](#topic-runtime-configuration).

**This prepares you for:** [Android platform host](#topic-android-platform-host), [Frame timing and physics stepping](#topic-frame-timing), [Performance monitors](#topic-performance-monitors).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

`Main` is the common startup layer before a selected platform host runs the engine.

Engine bootstrap consumes runtime configuration to initialize execution and select the configured main scene.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
Error Main::setup(const char *execpath, int argc, char *argv[], bool p_second_phase) {
	GodotProfileZone("setup");
	Thread::make_main_thread();
	set_current_thread_safe_for_nodes(true);

	OS::get_singleton()->initialize();

	CoreGlobals::print_ready = true;

#if !defined(OVERRIDE_PATH_ENABLED) && !defined(TOOLS_ENABLED)
	String old_cwd = OS::get_singleton()->get_cwd();
#if defined(MACOS_ENABLED) || defined(APPLE_EMBEDDED_ENABLED)
	String new_cwd = OS::get_singleton()->get_bundle_resource_dir();
```

Source: `main/main.cpp` — Main::setup (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `main/main.cpp:744` — Main::setup
- Code: `main/main.cpp:4016` — Main::start
- Code: `main/main.h:39` — Main

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-gui-control-layout"></a>
## 207. GUI controls and layout

**Scope:** First-party

**Builds on:** [Scene graph nodes](#topic-scene-graph).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation includes buttons, text and code editing, dialogs, file selection, color picking, graph editing, and flow layout.

Control-derived widgets and Container-derived layouts implement the scene GUI layer.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CodeEdit : public TextEdit {
	GDCLASS(CodeEdit, TextEdit)

public:
	// Keep enums in sync with:
	// core/object/script_language.h - ScriptLanguage::CodeCompletionKind
	enum CodeCompletionKind {
		KIND_CLASS,
		KIND_FUNCTION,
		KIND_SIGNAL,
		KIND_VARIABLE,
		KIND_MEMBER,
		KIND_ENUM,
```

Source: `scene/gui/code_edit.h` — CodeEdit (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Control, Container, CodeEdit

Evidence:
- Code: `scene/gui/control.h:45` — Control
- Code: `scene/gui/container.h:35` — Container
- Code: `scene/gui/code_edit.h:36` — CodeEdit

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Parametric curve bases](#topic-curve-and-patch-bases)
- [GUI controls and layout](#topic-gui-controls)
- [ISA and platform portability](#topic-isa-portability)
- [Display, camera, video, and movie services](#topic-platform-presentation)
- [Regular-expression JIT code generation](#topic-regex-jit-codegen)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Themes and style boxes](#topic-themes-and-style-boxes)
- [Control-tree user interfaces](#topic-ui-composition)
- [UI themes and style boxes](#topic-ui-theming)
- [Version-control integration](#topic-version-control-integration)
- [WebSocket framing and event contexts](#topic-websocket-frame-events)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-gui-controls"></a>
## 208. GUI controls and layout

**Scope:** First-party

**Builds on:** [Canvas and viewport presentation](#topic-canvas-and-viewports).

**This prepares you for:** [Text layout and editing](#topic-text-interface).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The partition includes both leaf controls and container controls.

Control-derived GUI classes implement selection, input, scrolling, split layout, menus, and graph editing within the CanvasItem tree.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GraphEdit : public Control {
	GDCLASS(GraphEdit, Control);

public:
	struct Connection : RefCounted {
		StringName from_node;
		StringName to_node;
		int from_port = 0;
		int to_port = 0;
		float activity = 0.0;
		bool keep_alive = true;

	private:
```

Source: `scene/gui/graph_edit.h` — class GraphEdit : public Control (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GraphEditConnection

Evidence:
- Code: `scene/gui/scroll_bar.h:35` — class ScrollBar : public Range
- Code: `scene/gui/split_container.h:98` — class SplitContainer : public Container
- Code: `scene/gui/tree.h:469` — class Tree : public Control
- Code: `scene/gui/graph_edit.h:113` — class GraphEdit : public Control

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GUI controls and layout](#topic-gui-control-layout)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-networking"></a>
## 209. HTTP and multiplayer

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The two APIs appear as separate scene-main facilities.

HTTPRequest is a Node API for HTTP work, while MultiplayerAPI and MultiplayerPeer define reference-counted multiplayer and packet-peer interfaces.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class HTTPRequest : public Node {
	GDCLASS(HTTPRequest, Node);

public:
	enum Result {
		RESULT_SUCCESS,
		RESULT_CHUNKED_BODY_SIZE_MISMATCH,
		RESULT_CANT_CONNECT,
		RESULT_CANT_RESOLVE,
		RESULT_CONNECTION_ERROR,
		RESULT_TLS_HANDSHAKE_ERROR,
		RESULT_NO_RESPONSE,
		RESULT_BODY_SIZE_LIMIT_EXCEEDED,
```

Source: `scene/main/http_request.h` — class HTTPRequest : public Node (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/main/http_request.h:41` — class HTTPRequest : public Node
- Code: `scene/main/multiplayer_api.h:36` — class MultiplayerAPI : public RefCounted
- Code: `scene/main/multiplayer_peer.h:38` — class MultiplayerPeer : public PacketPeer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [BVH traversal](#topic-bvh-traversal)
- [ENet transport and multiplayer adaptation](#topic-enet-transport-and-multiplayer)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Ray packets and streams](#topic-packet-queries)
- [Scene graph and lifecycle](#topic-scene-tree)
- [WebSocket multiplayer](#topic-websocket-multiplayer)

<a id="topic-navigation-agents"></a>
## 210. Navigation agents and regions

**Scope:** First-party

**Builds on:** [Scene graph nodes](#topic-scene-graph).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Separate 2D and 3D implementations retain paths, path metadata, links, regions, and obstacles.

Navigation agents consume navigation-result paths, while regions, links, and obstacles describe navigation-world participation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NavigationAgent2D : public Node {
	GDCLASS(NavigationAgent2D, Node);

	Node2D *agent_parent = nullptr;

	RID agent;
	RID map_override;

	bool avoidance_enabled = false;
	uint32_t avoidance_layers = 1;
	uint32_t avoidance_mask = 1;
	real_t avoidance_priority = 1.0;
	uint32_t navigation_layers = 1;
```

Source: `scene/2d/navigation/navigation_agent_2d.h` — NavigationAgent2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

NavigationAgent2D, NavigationAgent3D, NavigationRegion3D

Evidence:
- Code: `scene/2d/navigation/navigation_agent_2d.h:40` — NavigationAgent2D
- Code: `scene/3d/navigation/navigation_agent_3d.h:42` — NavigationAgent3D
- Code: `scene/3d/navigation/navigation_region_3d.h:36` — NavigationRegion3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Compact heightfield representation](#topic-compact-heightfield)
- [Editing selection history](#topic-editing-selection-history)
- [Half-edge topology](#topic-half-edge-topology)
- [Navigation maps and path queries](#topic-navigation)
- [Navigation map composition](#topic-navigation-map-composition)
- [Navigation mesh construction](#topic-navigation-mesh-construction)
- [2D and 3D navigation queries](#topic-navigation-queries)
- [Navigation region construction](#topic-navigation-regions)
- [Tile map layer data](#topic-tilemap-layer-data)

<a id="topic-pluggable-server-backends"></a>
## 211. Pluggable and wrapped server backends

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The same server-facing abstractions support fallback, extension-provided, and wrapped implementations.

C++ inheritance supplies extension, dummy, and multithread-wrapped implementations behind the physics, text, XR, and rendering server interfaces.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class PhysicsServer3DDummy : public PhysicsServer3D {
	GDCLASS(PhysicsServer3DDummy, PhysicsServer3D);

	PhysicsDirectBodyState3DDummy *body_state_dummy = nullptr;
	PhysicsDirectSpaceState3DDummy *space_state_dummy = nullptr;

public:
	virtual RID world_boundary_shape_create() override { return RID(); }
	virtual RID separation_ray_shape_create() override { return RID(); }
	virtual RID sphere_shape_create() override { return RID(); }
	virtual RID box_shape_create() override { return RID(); }
	virtual RID capsule_shape_create() override { return RID(); }
	virtual RID cylinder_shape_create() override { return RID(); }
```

Source: `servers/physics_3d/physics_server_3d_dummy.h` — PhysicsServer3DDummy (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `servers/physics_3d/physics_server_3d_extension.h:203` — PhysicsServer3DExtension
- Code: `servers/physics_3d/physics_server_3d_dummy.h:131` — PhysicsServer3DDummy
- Code: `servers/text/text_server_extension.h:39` — TextServerExtension
- Code: `servers/xr/xr_interface_extension.h:36` — XRInterfaceExtension

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Virtual implementation extensions](#topic-extensions)
- [XR tracking and poses](#topic-xr-tracking)

<a id="topic-resource-assets"></a>
## 212. Resource-backed assets

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Audio buses, streams, and effects](#topic-audio-bus-processing), [Textures, meshes, and materials](#topic-rendering-resources), [2D skeleton modification stacks](#topic-skeleton-modification), [Themes and style boxes](#topic-themes-and-style-boxes), [3D shapes, worlds, and skins](#topic-three-dimensional-content), [2D shapes, tiles, and paths](#topic-two-dimensional-content).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The resource family is the common content representation used throughout the inspected scene partition.

Resource-derived classes represent reusable scene content such as textures, meshes, materials, themes, animations, shapes, and packed scenes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Animation : public Resource {
	GDCLASS(Animation, Resource);
	RES_BASE_EXTENSION("anim");

public:
	typedef uint64_t TrackCacheID;

	static inline String PARAMETERS_BASE_PATH = "parameters/";
	static constexpr real_t DEFAULT_STEP = 1.0 / 30;

	enum TrackType : uint8_t {
		TYPE_VALUE, // Set a value in a property, can be interpolated.
		TYPE_POSITION_3D, // Position 3D track, can be compressed.
```

Source: `scene/resources/animation.h` — class Animation : public Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PackedScene, Animation, TileMapPattern

Evidence:
- Code: `scene/resources/packed_scene.h:248` — class PackedScene : public Resource
- Code: `scene/resources/animation.h:38` — class Animation : public Resource
- Code: `scene/resources/texture.h:37` — class Texture : public Resource
- Code: `scene/resources/theme.h:38` — class Theme : public Resource

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [GPU command encoding](#topic-gpu-command-encoding)
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [Physics shapes, objects, and collision reports](#topic-physics-collision)
- [Renderer-owned resource storage](#topic-renderer-storage)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Asynchronous resource previews](#topic-resource-previews)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene state inspection](#topic-scene-state)
- [Skeletal modifiers and inverse kinematics](#topic-skeletal-ik)
- [Temporal upscaling dispatch](#topic-temporal-upscaling-dispatch)
- [Texture resources and fallback placeholders](#topic-textures-and-placeholders)
- [UI themes and style boxes](#topic-ui-theming)
- [Visual shader nodes](#topic-visual-shader-nodes)
- [Resources and asset lifecycle](#topic-resources)
- [Scene construction and commit](#topic-scene-commit)

<a id="topic-scene-tree"></a>
## 213. Scene graph and lifecycle

**Scope:** First-party

**Builds on:** [Engine object model](#topic-object-model).

**This prepares you for:** [Input events and actions](#topic-input-actions), [Packed scene serialization](#topic-packed-scenes), [Physics shapes, objects, and collision reports](#topic-physics-collision), [Meshes, materials, textures, and instancing](#topic-rendering-assets), [Control-tree user interfaces](#topic-ui-composition).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Nodes carry parent, owner, children, name, tree, processing, and RPC configuration state.

A scene tree arranges Object-derived Node instances into a parent-child hierarchy that SceneTree owns and updates.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class SceneTreeTimer : public RefCounted {
	GDCLASS(SceneTreeTimer, RefCounted);

	double time_left = 0.0;
	bool process_always = true;
	bool process_in_physics = false;
	bool ignore_time_scale = false;

protected:
	static void _bind_methods();

public:
	void set_time_left(double p_time);
```

Source: `scene/main/scene_tree.h` — SceneTree (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node, SceneState

Evidence:
- Code: `scene/main/node.h` — Node::Data
- Code: `scene/main/node.h` — Node::add_child
- Code: `scene/main/scene_tree.h:89` — SceneTree

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Canvas and viewport presentation](#topic-canvas-and-viewports)
- [COLLADA scene interchange](#topic-collada-import)
- [Expression parsing and evaluation](#topic-expression-evaluation)
- [HTTP and multiplayer](#topic-networking)
- [ObjectDB snapshots](#topic-objectdb-snapshots)
- [Post-import processing](#topic-post-import-processing)
- [Scene graph nodes](#topic-scene-graph)
- [Scene importing](#topic-scene-importing)
- [SceneTree execution](#topic-scene-tree-execution)
- [Visual Shader source partition](#topic-visual-shader-module-build)
- [Visual shader nodes](#topic-visual-shader-nodes)
- [Scene tree and signal connections](#topic-scene-tree-and-signal-connections)

<a id="topic-skeletal-modifiers-and-ik"></a>
## 214. Skeletal modifiers and inverse kinematics

**Scope:** First-party

**Builds on:** [Scene graph nodes](#topic-scene-graph).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module contains multiple IK algorithms and skeletal runtime modifiers.

SkeletonModifier3D subclasses apply bone constraints, IK solvers, retargeting, spring-bone simulation, and XR tracker data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ChainIK3D : public IKModifier3D {
	GDCLASS(ChainIK3D, IKModifier3D);

public:
	struct ChainIK3DSetting : public IKModifier3DSetting {
#ifdef TOOLS_ENABLED
		// Note:
		// To cache global rest on global pose in SkeletonModifier process.
		// Since gizmo drawing might be processed after SkeletonModifier process,
		// so the gizmo which depend on modified pose is not drawn correctly.
		// Especially, limitation sphere is needed this since it bound mutable bone axis which retrieve by bone pose to the parent bone rest.
		Transform3D root_global_rest;
#endif // TOOLS_ENABLED
```

Source: `scene/3d/chain_ik_3d.h` — ChainIK3D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SkeletonModifier3D, ChainIK3D, SpringBoneSimulator3D

Evidence:
- Code: `scene/3d/skeleton_modifier_3d.h:36` — SkeletonModifier3D
- Code: `scene/3d/chain_ik_3d.h:35` — ChainIK3D
- Code: `scene/3d/spring_bone_simulator_3d.h:37` — SpringBoneSimulator3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Skeleton modification and physical-bone simulation](#topic-skeleton-modifiers)
- [XR tracking and poses](#topic-xr-tracking)

<a id="topic-tilemap-layer-data"></a>
## 215. Tile map layer data

**Scope:** First-party

**Builds on:** [Scene graph nodes](#topic-scene-graph).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

A TileMap node coordinates TileMapLayer instances.

TileMapLayer stores cell data keyed by grid coordinates and derives rendering, physics, navigation, and debug quadrants from that data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#endif // NAVIGATION_2D_DISABLED
class TileMapLayer;
class TerrainConstraint;

enum TileMapDataFormat {
	TILE_MAP_DATA_FORMAT_1 = 0,
	TILE_MAP_DATA_FORMAT_2,
	TILE_MAP_DATA_FORMAT_3,
	TILE_MAP_DATA_FORMAT_MAX,
};

class TileMap : public Node2D {
	GDCLASS(TileMap, Node2D)
```

Source: `scene/2d/tile_map.h` — TileMap (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TileMapLayer, TileMapLayerCellData

Evidence:
- Code: `scene/2d/tile_map_layer.h:332` — TileMapLayer
- Code: `scene/2d/tile_map_layer.h:105` — CellData
- Code: `scene/2d/tile_map.h:51` — TileMap

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Navigation agents and regions](#topic-navigation-agents)
- [Tile libraries, cells, and patterns](#topic-tile-libraries)

<a id="topic-apple-embedded-hosting"></a>
## 216. Apple embedded hosting

**Scope:** First-party

**Builds on:** [Native platform adapters](#topic-native-platform-adapters).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This layer packages the engine for Apple embedded targets.

Apple embedded hosting uses SwiftUI and Objective-C++ alongside display, keyboard, view-controller, text-to-speech, and Vulkan-context platform adapters.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// "Embedded" as in "Embedded Device".
class DisplayServerAppleEmbedded : public DisplayServer {
	GDSOFTCLASS(DisplayServerAppleEmbedded, DisplayServer);

	_THREAD_SAFE_CLASS_

#if defined(RD_ENABLED)
	RenderingContextDriver *rendering_context = nullptr;
	RenderingDevice *rendering_device = nullptr;
#endif
	NativeMenu *native_menu = nullptr;

	id tts = nullptr;
```

Source: `drivers/apple_embedded/display_server_apple_embedded.h` — class DisplayServerAppleEmbedded : public DisplayServer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `drivers/apple_embedded/app.swift` — GodotSwiftUIViewController and SwiftUIApp
- Code: `drivers/apple_embedded/display_server_apple_embedded.h:62` — class DisplayServerAppleEmbedded : public DisplayServer
- Code: `drivers/apple_embedded/rendering_context_driver_vulkan_apple_embedded.h:39` — class RenderingContextDriverVulkanAppleEmbedded : public RenderingContextDriverVulkan

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Apple embedded packaging and signing](#topic-apple-embedded-packaging)
- [Keyboard and MIDI parsing](#topic-input-midi)

<a id="topic-audio-backend-adapters"></a>
## 217. Audio backend adapters

**Scope:** First-party

**Builds on:** [Native platform adapters](#topic-native-platform-adapters).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementations derive from the common AudioDriver interface.

Audio output is provided by platform adapters selected with the driver build for ALSA, PulseAudio, CoreAudio, WASAPI, and XAudio2.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

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

Evidence:
- Code: `drivers/alsa/audio_driver_alsa.h:43` — class AudioDriverALSA : public AudioDriver
- Code: `drivers/pulseaudio/audio_driver_pulseaudio.h:46` — class AudioDriverPulseAudio : public AudioDriver
- Code: `drivers/coreaudio/audio_driver_coreaudio.h:44` — class AudioDriverCoreAudio : public AudioDriver
- Code: `drivers/wasapi/audio_driver_wasapi.h:45` — class AudioDriverWASAPI : public AudioDriver

<a id="topic-completion-contexts"></a>
## 218. Contextual completion contracts

**Scope:** First-party

**Builds on:** [Fixture contracts](#topic-fixture-contracts).

**This prepares you for:** [Scene-aware test context](#topic-scene-contexts).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The configuration records assertions over displayed or inserted completion text.

Completion fixture contracts pair a cursor-bearing script with configuration rules that include or exclude candidates.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
[output]
include=[
    {"display": "DrawMode",
     "location": 256},
    {"display": "Anchor",
     "location": 257},
    {"display": "TextureRepeat",
     "location": 258},
]
```

Source: `modules/gdscript/tests/scripts/completion/builtin_enum/builtin_enum_autocomplete.cfg` — [output] (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Completion Test Configuration, Test Script Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/completion/builtin_enum/builtin_enum_autocomplete.gd:3` — _ready()
- Code: `modules/gdscript/tests/scripts/completion/builtin_enum/builtin_enum_autocomplete.cfg:1` — [output]

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript editor services](#topic-gdscript-editor-services)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [Script resources and instances](#topic-scripting)

<a id="topic-diagnostic-expectations"></a>
## 219. Diagnostic expectations

**Scope:** First-party

**Builds on:** [Fixture contracts](#topic-fixture-contracts).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This makes diagnostic behavior part of the tested interface.

Diagnostic fixture contracts preserve parser errors, analyzer warnings, and runtime errors as expected textual results.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
GDTEST_OK
~~ WARNING at line 2: (UNUSED_VARIABLE) The local variable "unused" is declared but never used in the block. If this is intended, prefix it with an underscore: "_unused".
```

Source: `modules/gdscript/tests/scripts/analyzer/warnings/unused_variable.out` — UNUSED_VARIABLE (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/parser/warnings/integer_division.out:2` — INTEGER_DIVISION
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/unused_variable.out:2` — UNUSED_VARIABLE
- Code: `modules/gdscript/tests/scripts/runtime/errors/invalid_cast.out` — expected runtime errors

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Image quality metrics](#topic-image-quality-metrics)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)

<a id="topic-enet-transport-and-multiplayer"></a>
## 220. ENet transport and multiplayer adaptation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The connection implementation also invokes the engine Compression API for payload data.

ENetConnection and ENetPacketPeer wrap connection and peer behavior, while ENetMultiplayerPeer adapts the transport to MultiplayerPeer.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ENetConnection : public RefCounted {
	GDCLASS(ENetConnection, RefCounted);

public:
	enum CompressionMode {
		COMPRESS_NONE = 0,
		COMPRESS_RANGE_CODER,
		COMPRESS_FASTLZ,
		COMPRESS_ZLIB,
		COMPRESS_ZSTD,
	};

	enum HostStatistic {
```

Source: `modules/enet/enet_connection.h` — ENetConnection (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ENetConnection, ENetPacketPeer, ENetMultiplayerPeer

Evidence:
- Code: `modules/enet/enet_connection.h:43` — ENetConnection
- Code: `modules/enet/enet_multiplayer_peer.h:39` — ENetMultiplayerPeer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [ENet host and peer transport](#topic-enet-host-peer-transport)
- [ENet wire commands](#topic-enet-wire-commands)
- [Godot ENet socket adaptation](#topic-godot-enet-socket-adaptation)
- [HTTP and multiplayer](#topic-networking)
- [Wraparound network time](#topic-wraparound-network-time)

<a id="topic-gdscript-bytecode-runtime"></a>
## 221. GDScript bytecode compilation and execution

**Scope:** First-party

**Builds on:** [GDScript static analysis](#topic-gdscript-static-analysis).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

GDScriptFunction stores execution metadata, while GDScriptInstance binds a script to an engine object.

The compiler lowers analyzed script trees through code-generation classes, and the VM executes functions using validated calls, getters, setters, and container operations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GDScriptByteCodeGenerator : public GDScriptCodeGenerator {
	struct StackSlot {
		Variant::Type type = Variant::NIL;
		bool can_contain_object = true;
		Vector<int> bytecode_indices;

		StackSlot() = default;
		StackSlot(Variant::Type p_type, bool p_can_contain_object) :
				type(p_type), can_contain_object(p_can_contain_object) {}
	};

	struct CallTarget {
		Address target;
```

Source: `modules/gdscript/gdscript_byte_codegen.h` — GDScriptByteCodeGenerator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDScript, GDScriptInstance, GDScriptFunction

Evidence:
- Code: `modules/gdscript/gdscript_byte_codegen.h:40` — GDScriptByteCodeGenerator
- Code: `modules/gdscript/gdscript_vm.cpp` — GDScriptFunction VM execution

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Graphite SILF rule execution](#topic-graphite-rule-execution)
- [Script resources and instances](#topic-scripting)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-gpu-texture-compression"></a>
## 222. GPU block-compression dispatch

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module packages BC1, BC4, BC6H, alpha-stitching, and RGB-to-RGBA shader sources.

The Betsy compressor relies on GLSL compute shaders that fetch source texels or blocks and write compressed results to storage images.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class BetsyCompressor : public Object {
	GDSOFTCLASS(BetsyCompressor, Object);

	mutable CommandQueueMT command_queue;
	bool exit = false;
	WorkerThreadPool::TaskID task_id = WorkerThreadPool::INVALID_TASK_ID;

	struct BetsyShader {
		RID compiled;
		RID pipeline;
	};

	// Resources shared by all compression formats.
```

Source: `modules/betsy/image_compress_betsy.h` — BetsyCompressor (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Image

Evidence:
- Code: `modules/betsy/image_compress_betsy.h:96` — BetsyCompressor
- Code: `modules/betsy/alpha_stitch.glsl:16` — main

<a id="topic-godot-member-exposure"></a>
## 223. Godot member exposure

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The source generators inspect compatible C# members and create the metadata queried by ScriptManagerBridge.

Generator output registers exported members, signals, and RPC metadata as engine-visible method and property descriptions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
[Generator]
    public class ScriptMethodsGenerator : ISourceGenerator
    {
        public void Initialize(GeneratorInitializationContext context)
        {
        }

        public void Execute(GeneratorExecutionContext context)
        {
            if (context.IsGodotSourceGeneratorDisabled("ScriptMethods"))
                return;

            INamedTypeSymbol[] godotClasses = context
                .Compilation.SyntaxTrees
```

Source: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptMethodsGenerator.cs` — ScriptMethodsGenerator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Variant

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPropertiesGenerator.cs:13` — ScriptPropertiesGenerator
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptMethodsGenerator.cs:11` — ScriptMethodsGenerator
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptSignalsGenerator.cs:11` — ScriptSignalsGenerator
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ScriptManagerBridge.cs:20` — ScriptManagerBridge

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [High-level RPC routing](#topic-high-level-rpc)
- [Class metadata and reflection](#topic-reflection)
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-lsp-semantic-fixtures"></a>
## 224. Language-server semantic fixtures

**Scope:** First-party

**Builds on:** [Fixture contracts](#topic-fixture-contracts).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These inputs are distinct from completion assertions because they model symbols and source ranges.

Language-server fixture contracts provide nested declarations, local scopes, documentation comments, and lambdas for semantic-editor tests.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
const LSP::DocumentSymbol *test_resolve_symbol_at(const String &p_uri, const LSP::Position p_pos, const String &p_expected_uri, const String &p_expected_name, const LSP::Range &p_expected_range) {
	Ref<GDScriptWorkspace> workspace = GDScriptLanguageProtocol::get_singleton()->get_workspace();

	LSP::TextDocumentPositionParams params = pos_in(p_uri, p_pos);
	const LSP::DocumentSymbol *symbol = workspace->resolve_symbol(params);
	CHECK(symbol);

	if (symbol) {
		CHECK_EQ(symbol->uri, p_expected_uri);
		CHECK_EQ(symbol->name, p_expected_name);
		CHECK_EQ(symbol->selectionRange, p_expected_range);
	}
```

Source: `modules/gdscript/tests/test_lsp.h` — test_resolve_symbol_at (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Test Script Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/lsp/class.gd:114` — Inner三.NestedInInner3
- Code: `modules/gdscript/tests/test_lsp.h:133` — test_resolve_symbol_at

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript language-server support](#topic-gdscript-language-server)

<a id="topic-midi-input-adapters"></a>
## 225. MIDI input adapters

**Scope:** First-party

**Builds on:** [Native platform adapters](#topic-native-platform-adapters).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Each named implementation derives from the common MIDIDriver interface.

MIDI input is supplied by platform adapters for ALSA MIDI, CoreMIDI, and WinMM.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class MIDIDriverALSAMidi : public MIDIDriver {
	Thread thread;
	Mutex mutex;

	struct InputConnection {
		InputConnection() = default;
		InputConnection(int p_device_index, snd_rawmidi_t *p_rawmidi);

		Parser parser;
		snd_rawmidi_t *rawmidi_ptr = nullptr;

		// Read in and parse available data, forwarding complete events to Input.
		void read();
```

Source: `drivers/alsamidi/midi_driver_alsamidi.h` — class MIDIDriverALSAMidi : public MIDIDriver (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `drivers/alsamidi/midi_driver_alsamidi.h:43` — class MIDIDriverALSAMidi : public MIDIDriver
- Code: `drivers/coremidi/midi_driver_coremidi.h:43` — class MIDIDriverCoreMidi : public MIDIDriver
- Code: `drivers/winmidi/midi_driver_winmidi.h:42` — class MIDIDriverWinMidi : public MIDIDriver

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Keyboard and MIDI parsing](#topic-input-midi)

<a id="topic-navigation-avoidance"></a>
## 226. Navigation avoidance

**Scope:** First-party

**Builds on:** [Navigation map composition](#topic-navigation-map-composition).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Both dimensions model agent positions and velocities plus obstacle geometry, and their maps expose agent and obstacle collections.

Agents and obstacles are map members that provide avoidance data alongside pathfinding data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NavAgent2D : public NavRid2D {
	Vector2 position;
	Vector2 target_position;
	Vector2 velocity;
	Vector2 velocity_forced;
	real_t radius = NavigationDefaults2D::AVOIDANCE_AGENT_RADIUS;
	real_t max_speed = NavigationDefaults2D::AVOIDANCE_AGENT_MAX_SPEED;
	real_t time_horizon_agents = NavigationDefaults2D::AVOIDANCE_AGENT_TIME_HORIZON_AGENTS;
	real_t time_horizon_obstacles = NavigationDefaults2D::AVOIDANCE_AGENT_TIME_HORIZON_OBSTACLES;
	int max_neighbors = NavigationDefaults2D::AVOIDANCE_AGENT_MAX_NEIGHBORS;
	real_t neighbor_distance = NavigationDefaults2D::AVOIDANCE_AGENT_NEIGHBOR_DISTANCE;
	Vector2 safe_velocity;
	bool clamp_speed = true; // Experimental, clamps velocity to max_speed.
```

Source: `modules/navigation_2d/nav_agent_2d.h` — NavAgent2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/navigation_2d/nav_agent_2d.h:45` — NavAgent2D
- Code: `modules/navigation_2d/nav_obstacle_2d.h:42` — NavObstacle2D
- Code: `modules/navigation_3d/nav_agent_3d.h:46` — NavAgent3D
- Code: `modules/navigation_3d/nav_obstacle_3d.h:41` — NavObstacle3D

<a id="topic-navigation-path-queries"></a>
## 227. Navigation path queries

**Scope:** First-party

**Builds on:** [Navigation map composition](#topic-navigation-map-composition).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The query implementations maintain begin and end polygons, search data, path points, and path-return constraints.

Path queries consume a map read iteration, select polygons, and construct a path corridor in 2D or 3D.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NavMeshQueries2D {
public:
	struct PathQuerySlot {
		LocalVector<Nav2D::NavigationPoly> path_corridor;
		Heap<Nav2D::NavigationPoly *, Nav2D::NavPolyTravelCostGreaterThan, Nav2D::NavPolyHeapIndexer> traversable_polys;
		bool in_use = false;
		uint32_t slot_index = 0;
		AHashMap<const Nav2D::Polygon *, uint32_t> poly_to_id;
	};

	struct NavMeshPathQueryTask2D {
		enum TaskStatus {
			QUERY_STARTED,
```

Source: `modules/navigation_2d/2d/nav_mesh_queries_2d.h` — NavMeshQueries2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/navigation_2d/2d/nav_mesh_queries_2d.h:47` — NavMeshQueries2D
- Code: `modules/navigation_3d/3d/nav_mesh_queries_3d.h:47` — NavMeshQueries3D
- Code: `modules/navigation_2d/nav_utils_2d.h` — Nav2D::NavigationPoly
- Code: `modules/navigation_3d/nav_utils_3d.h` — Nav3D::NavigationPoly

<a id="topic-openxr-composition-layers"></a>
## 228. OpenXR composition layers

**Scope:** First-party

**Builds on:** [OpenXR runtime integration](#topic-openxr-runtime-integration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The scene layer has quad, cylinder, and equirect subclasses, while the extension owns swapchain and composition-layer state.

Composition-layer scene nodes and an extension submit specialized layer descriptions to the OpenXR runtime.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class OpenXRCompositionLayer : public Node3D {
	GDCLASS(OpenXRCompositionLayer, Node3D);

public:
	// Must be identical to Filter enum definition in OpenXRCompositionLayerExtension.
	enum Filter {
		FILTER_NEAREST,
		FILTER_LINEAR,
		FILTER_CUBIC,
	};

	// Must be identical to MipmapMode enum definition in OpenXRCompositionLayerExtension.
	enum MipmapMode {
```

Source: `modules/openxr/scene/openxr_composition_layer.h` — class OpenXRCompositionLayer : public Node3D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/openxr/scene/openxr_composition_layer.h:45` — class OpenXRCompositionLayer : public Node3D
- Code: `modules/openxr/scene/openxr_composition_layer_quad.h:37` — class OpenXRCompositionLayerQuad : public OpenXRCompositionLayer
- Code: `modules/openxr/extensions/openxr_composition_layer_extension.h` — OpenXRCompositionLayerExtension::SwapchainState and CompositionLayer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Virtual implementation extensions](#topic-extensions)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-openxr-extension-wrappers"></a>
## 229. OpenXR extension wrappers

**Scope:** First-party

**Builds on:** [OpenXR runtime integration](#topic-openxr-runtime-integration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied extensions cover platform graphics bindings, controller profiles, composition layers, hand tracking, render models, futures, and spatial capabilities.

Extension wrappers isolate optional OpenXR runtime features behind a common base interface.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
if (p_use_next_extensions) {
		for (OpenXRExtensionWrapper *wrapper : OpenXRAPI::get_registered_extension_wrappers()) {
			void *np = wrapper->set_swapchain_create_info_and_get_next_pointer(next_pointer);
			if (np != nullptr) {
				next_pointer = np;
			}
		}
	}

	XrSwapchainCreateInfo swapchain_create_info = {
		XR_TYPE_SWAPCHAIN_CREATE_INFO, // type
		next_pointer, // next
		p_create_flags, // createFlags
		p_usage_flags, // usageFlags
```

Source: `modules/openxr/openxr_api.cpp` — OpenXRAPI::get_registered_extension_wrappers (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OpenXRFutureResult, OpenXRRenderModelData

Evidence:
- Code: `modules/openxr/extensions/openxr_extension_wrapper.h:52` — class OpenXRExtensionWrapper : public Object
- Code: `modules/openxr/openxr_api.cpp:106` — OpenXRAPI::get_registered_extension_wrappers
- Code: `modules/openxr/extensions/SCsub` — extension and platform source selection

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Virtual implementation extensions](#topic-extensions)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Shader interface metadata](#topic-shader-interface-metadata)

<a id="topic-openxr-render-models"></a>
## 230. OpenXR render models

**Scope:** First-party

**Builds on:** [OpenXR runtime integration](#topic-openxr-runtime-integration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Render-model data is separate from interaction-profile suggestions, allowing runtime-provided device models to drive presentation.

The render-model extension tracks runtime render models, and scene nodes display models or manage their activation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class OpenXRRenderModel : public Node3D {
	GDCLASS(OpenXRRenderModel, Node3D);

private:
	RID render_model;
	Node3D *scene = nullptr;
	HashMap<String, Node3D *> animatable_nodes;

	void _load_render_model_scene();
	void _on_render_model_top_level_path_changed(RID p_render_model);

protected:
	static void _bind_methods();
```

Source: `modules/openxr/scene/openxr_render_model.h` — class OpenXRRenderModel : public Node3D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

OpenXRRenderModelData

Evidence:
- Code: `modules/openxr/extensions/openxr_render_model_extension.h` — class OpenXRRenderModelData; class OpenXRRenderModelExtension
- Code: `modules/openxr/scene/openxr_render_model.h:40` — class OpenXRRenderModel : public Node3D
- Code: `modules/openxr/scene/openxr_render_model_manager.h:42` — class OpenXRRenderModelManager : public Node3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Virtual implementation extensions](#topic-extensions)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-plugin-and-assembly-reload"></a>
## 231. Plugin and assembly reload

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The loader separates shared assemblies from plugin-specific dependency resolution and records a weak reference for load-context tracking.

Editor plug-ins load into AssemblyLoadContext instances that resolve dependencies and may be collectible.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
{
    public partial class HotReloadAssemblyWatcher : Node
    {
#nullable disable
        private Timer _watchTimer;
#nullable enable

        public override void _Notification(int what)
        {
            if (what == Node.NotificationWMWindowFocusIn)
            {
                RestartTimer();

                if (Internal.IsAssembliesReloadingNeeded())
```

Source: `modules/mono/editor/GodotTools/GodotTools/HotReloadAssemblyWatcher.cs` — HotReloadAssemblyWatcher (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/PluginLoadContext.cs:9` — PluginLoadContext
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/Main.cs:149` — LoadPlugin
- Code: `modules/mono/editor/GodotTools/GodotTools/HotReloadAssemblyWatcher.cs:8` — HotReloadAssemblyWatcher

<a id="topic-replicated-property-synchronization"></a>
## 232. Replicated property synchronization

**Scope:** First-party

**Builds on:** [Scene replication configuration](#topic-scene-replication-configuration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The synchronizer builds watchers from configured paths and participates in replication start and stop through object configuration calls.

Synchronizers observe configured properties and forward sync property lists to scene replication.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
Object *MultiplayerSynchronizer::_get_prop_target(Object *p_obj, const NodePath &p_path) {
	if (p_path.get_name_count() == 0) {
		return p_obj;
	}
	Node *node = Object::cast_to<Node>(p_obj);
	ERR_FAIL_COND_V_MSG(!node || !node->has_node(p_path), nullptr, vformat("Node '%s' not found.", p_path));
	return node->get_node(p_path);
}

void MultiplayerSynchronizer::_stop() {
#ifdef TOOLS_ENABLED
	if (Engine::get_singleton()->is_editor_hint()) {
		return;
```

Source: `modules/multiplayer/multiplayer_synchronizer.cpp` — MultiplayerSynchronizer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SceneReplicationConfig

Evidence:
- Code: `modules/multiplayer/multiplayer_synchronizer.h:37` — MultiplayerSynchronizer
- Code: `modules/multiplayer/multiplayer_synchronizer.cpp:37` — MultiplayerSynchronizer
- Code: `modules/multiplayer/scene_replication_interface.cpp:46` — SceneReplicationInterface
- External (official, unverified (source anchor not found)): [MultiplayerSynchronizer — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_multiplayersynchronizer.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)

<a id="topic-replicated-spawning"></a>
## 233. Replicated spawning

**Scope:** First-party

**Builds on:** [Scene replication configuration](#topic-scene-replication-configuration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Spawner state records tracked nodes and their spawn data, while the replication interface sends spawn and despawn commands.

The spawner tracks spawnable scenes and sends configured spawn property lists while creating or removing nodes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const Variant MultiplayerSpawner::get_spawn_argument(const ObjectID &p_id) const {
	const SpawnInfo *info = tracked_nodes.getptr(p_id);
	return info ? info->args : Variant();
}

Node *MultiplayerSpawner::instantiate_scene(int p_id) {
	ERR_FAIL_COND_V_MSG(spawn_limit && spawn_limit <= tracked_nodes.size(), nullptr, "Spawn limit reached!");
	ERR_FAIL_UNSIGNED_INDEX_V((uint32_t)p_id, spawnable_scenes.size(), nullptr);
	SpawnableScene &sc = spawnable_scenes[p_id];
	if (sc.cache.is_null()) {
		sc.cache = ResourceLoader::load(sc.path);
	}
	ERR_FAIL_COND_V_MSG(sc.cache.is_null(), nullptr, "Invalid spawnable scene: " + sc.path);
```

Source: `modules/multiplayer/multiplayer_spawner.cpp` — MultiplayerSpawner::get_spawn_argument (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SceneReplicationConfig

Evidence:
- Code: `modules/multiplayer/multiplayer_spawner.h:37` — MultiplayerSpawner
- Code: `modules/multiplayer/multiplayer_spawner.cpp:294` — MultiplayerSpawner::get_spawn_argument
- Code: `modules/multiplayer/scene_replication_interface.cpp:46` — SceneReplicationInterface
- External (official, unverified (source anchor not found)): [MultiplayerSpawner — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_multiplayerspawner.html), accessed 2026-07-15

<a id="topic-module-build-configuration"></a>
## 234. SCons module build configuration

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Python SCons scripts determine whether modules can build and add C++ source files, generated shader headers, or third-party sources to the build environment.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```python
def can_build(env, platform):
    return not env["disable_physics_2d"]


def configure(env):
    pass
```

Source: `modules/godot_physics_2d/config.py` — can_build (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `modules/godot_physics_2d/config.py:1` — can_build
- Code: `modules/lightmapper_rd/SCsub` — GLSL_HEADER and add_source_files
- Code: `modules/jolt_physics/SCsub` — thirdparty_sources and add_source_files

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor build composition](#topic-editor-build-composition)
- [Platform-selective shader baking](#topic-platform-selective-shader-baking)

<a id="topic-visual-shader-vector-operations"></a>
## 235. Visual shader vector operations

**Scope:** First-party

**Builds on:** [Visual shader nodes](#topic-visual-shader-nodes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The node set includes operator, function, distance, refraction, compose, and decompose variants.

Vector-operation nodes apply arithmetic and vector functions as operations within a visual shader graph.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class VisualShaderNodeVectorOp : public VisualShaderNodeVectorBase {
	GDCLASS(VisualShaderNodeVectorOp, VisualShaderNodeVectorBase);

public:
	enum Operator {
		OP_ADD,
		OP_SUB,
		OP_MUL,
		OP_DIV,
		OP_MOD,
		OP_POW,
		OP_MAX,
		OP_MIN,
```

Source: `modules/visual_shader/vs_nodes/visual_shader_nodes.h` — VisualShaderNodeVectorOp (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VisualShaderNodeVectorOp

Evidence:
- Code: `modules/visual_shader/vs_nodes/visual_shader_nodes.h:911` — VisualShaderNodeVectorOp
- Code: `modules/visual_shader/vs_nodes/visual_shader_nodes.h:1892` — VisualShaderNodeVectorRefract

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Vector font export](#topic-vector-font-export)
- [Visual Shader source partition](#topic-visual-shader-module-build)

<a id="topic-webrtc-data-channels"></a>
## 236. WebRTC data channels

**Scope:** First-party

**Builds on:** [WebRTC peer connections](#topic-webrtc-peer-connections).

**This prepares you for:** [WebRTC multiplayer mesh](#topic-webrtc-multiplayer-mesh).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module supplies a base channel plus extension and JavaScript implementations.

A WebRTC data channel carries PacketPeer data through a WebRTC peer connection.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class WebRTCDataChannel : public PacketPeer {
	GDCLASS(WebRTCDataChannel, PacketPeer);

public:
	enum WriteMode {
		WRITE_MODE_TEXT,
		WRITE_MODE_BINARY,
	};

	enum ChannelState {
		STATE_CONNECTING,
		STATE_OPEN,
		STATE_CLOSING,
```

Source: `modules/webrtc/webrtc_data_channel.h` — WebRTCDataChannel (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebRTCDataChannel

Evidence:
- Code: `modules/webrtc/webrtc_data_channel.h:35` — WebRTCDataChannel
- Code: `modules/webrtc/webrtc_data_channel_js.h:37` — WebRTCDataChannelJS
- External (primary, verified): [WebRTC: Real-Time Communication in Browsers](https://www.w3.org/TR/webrtc/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Packets, HTTP, JSON, and JSON-RPC](#topic-network-data-exchange)
- [Networking and transport I/O](#topic-network-io)

<a id="topic-websocket-multiplayer"></a>
## 237. WebSocket multiplayer

**Scope:** First-party

**Builds on:** [WebSocket peers](#topic-websocket-peers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Its packet and pending-peer records are the cross-cutting state used by the multiplayer transport.

WebSocketMultiplayerPeer tracks WebSocket peers plus pending peers and packets.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class WebSocketMultiplayerPeer : public MultiplayerPeer {
	GDCLASS(WebSocketMultiplayerPeer, MultiplayerPeer);

private:
	Ref<WebSocketPeer> _create_peer();

protected:
	enum {
		SYS_NONE = 0,
		SYS_ADD = 1,
		SYS_DEL = 2,
		SYS_ID = 3,
```

Source: `modules/websocket/websocket_multiplayer_peer.h` — WebSocketMultiplayerPeer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebSocketMultiplayerPeer, WebSocketMultiplayerPacket

Evidence:
- Code: `modules/websocket/websocket_multiplayer_peer.h:39` — WebSocketMultiplayerPeer
- Code: `modules/websocket/websocket_multiplayer_peer.h` — WebSocketMultiplayerPeer::Packet
- Code: `modules/websocket/websocket_multiplayer_peer.h` — WebSocketMultiplayerPeer::PendingPeer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [WebSocket framing and event contexts](#topic-websocket-frame-events)
- [HTTP and multiplayer](#topic-networking)

<a id="topic-gltf-accessors"></a>
## 238. glTF binary accessor decoding and encoding

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Mesh attributes, indices, animation samples, and other numeric glTF values pass through this layer.

GLTFAccessor converts typed values to and from GLTFState’s indexed buffer views, including sparse data when the sparse representation is smaller than dense data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void GLTFBufferView::_bind_methods() {
	ClassDB::bind_method(D_METHOD("load_buffer_view_data", "state"), &GLTFBufferView::load_buffer_view_data);

	ClassDB::bind_static_method("GLTFBufferView", D_METHOD("from_dictionary", "dictionary"), &GLTFBufferView::from_dictionary);
	ClassDB::bind_method(D_METHOD("to_dictionary"), &GLTFBufferView::to_dictionary);

	ClassDB::bind_method(D_METHOD("get_buffer"), &GLTFBufferView::get_buffer);
	ClassDB::bind_method(D_METHOD("set_buffer", "buffer"), &GLTFBufferView::set_buffer);
	ClassDB::bind_method(D_METHOD("get_byte_offset"), &GLTFBufferView::get_byte_offset);
	ClassDB::bind_method(D_METHOD("set_byte_offset", "byte_offset"), &GLTFBufferView::set_byte_offset);
	ClassDB::bind_method(D_METHOD("get_byte_length"), &GLTFBufferView::get_byte_length);
	ClassDB::bind_method(D_METHOD("set_byte_length", "byte_length"), &GLTFBufferView::set_byte_length);
	ClassDB::bind_method(D_METHOD("get_byte_stride"), &GLTFBufferView::get_byte_stride);
	ClassDB::bind_method(D_METHOD("set_byte_stride", "byte_stride"), &GLTFBufferView::set_byte_stride);
```

Source: `modules/gltf/structures/gltf_buffer_view.cpp` — GLTFBufferView::load_buffer_view_data (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GLTFAccessor, GLTFBufferView, GLTFState

Evidence:
- Code: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor::_decode_as_numbers and encode_new_sparse_accessor_from_vec3s
- Code: `modules/gltf/structures/gltf_buffer_view.cpp:39` — GLTFBufferView::load_buffer_view_data
- External (primary, unverified (source anchor not found)): [glTF 2.0 Specification](https://registry.khronos.org/glTF/specs/2.0/glTF-2.0.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry buffer storage](#topic-buffer-storage)

<a id="topic-managed-csharp-scripting"></a>
## 239. managed C# script bridge and source generation

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Mono module represents managed scripts through CSharpScript and CSharpInstance, while its SDK tests use partial C# types and source generators for properties, methods, signals, serialization, and script paths.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```csharp
public class ScriptPropertiesGeneratorTests
{
    [Fact]
    public async Task ExportedFields()
    {
        await CSharpSourceGeneratorVerifier<ScriptPropertiesGenerator>.Verify(
            new string[] { "ExportedFields.cs", "MoreExportedFields.cs" },
            new string[] { "ExportedFields_ScriptProperties.generated.cs" }
        );
    }

    [Fact]
    public async Task ExportedProperties()
```

Source: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/ScriptPropertiesGeneratorTests.cs` — ScriptPropertiesGeneratorTests (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CSharpScript

Evidence:
- Code: `modules/mono/csharp_script.h` — CSharpScript and CSharpInstance
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/ScriptPropertiesGeneratorTests.cs:6` — ScriptPropertiesGeneratorTests

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Script resources and instances](#topic-scripting)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-apk-signing"></a>
## 240. APK signing and verification

**Scope:** First-party

**Builds on:** [Android export pipeline](#topic-android-export-pipeline).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Android editor has a Kotlin utility that imports `ApkSigner` and `ApkVerifier`.

APK signing and verification operate on Android export pipeline artifacts through the bundled apksig implementation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```java
*/
        public ApkSigner build() {
            if (mV3SigningExplicitlyDisabled && mV3SigningExplicitlyEnabled) {
                throw new IllegalStateException(
                        "Builder configured to both enable and disable APK "
                                + "Signature Scheme v3 signing");
            }

            if (mV3SigningExplicitlyDisabled) {
                mV3SigningEnabled = false;
            }

            if (mV3SigningExplicitlyEnabled) {
                mV3SigningEnabled = true;
```

Source: `platform/android/java/editor/src/main/java/com/android/apksig/ApkSigner.java` — ApkSigner (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/editor/src/main/java/org/godotengine/editor/utils/ApkSignerUtil.kt:109` — signApk
- Code: `platform/android/java/editor/src/main/java/com/android/apksig/ApkSigner.java:72` — ApkSigner
- Code: `platform/android/java/editor/src/main/java/com/android/apksig/ApkVerifier.java:34` — ApkVerifier

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)

<a id="topic-android-gradle-assembly"></a>
## 241. Android Gradle assembly

**Scope:** First-party

**Builds on:** [Android export pipeline](#topic-android-export-pipeline).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Android build root explicitly includes all five module names.

Gradle settings and build logic assemble app, library, editor, native-source-index, and install-time asset-pack modules for the Android export pipeline.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
*/
def generateBuildTasks(String flavor = "template", String edition = "standard", String androidDistro = "android") {
    if (!supportedFlavors.contains(flavor)) {
        throw new GradleException("Invalid build flavor: $flavor")
    }
    if (!supportedAndroidDistributions.contains(androidDistro)) {
        throw new GradleException("Invalid Android distribution: $androidDistro")
    }
    if (!supportedEditions.contains(edition)) {
        throw new GradleException("Invalid build edition: $edition")
    }
    if (edition == "mono" && flavor != "template") {
        throw new GradleException("'mono' edition only supports the 'template' flavor.")
    }
```

Source: `platform/android/java/build.gradle` — generateBuildTasks (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/settings.gradle:23` — include ':app'
- Code: `platform/android/java/settings.gradle:24` — include ':lib'
- Code: `platform/android/java/build.gradle:101` — generateBuildTasks
- External (official, verified): [Gradle DSL Reference](https://docs.gradle.org/current/dsl/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [Motion blur](#topic-motion-blur)
- [Main loop, OS, and time services](#topic-runtime-loop-time)

<a id="topic-browser-runtime-adapters"></a>
## 242. Browser runtime adapters

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The browser implementation is split between C++ platform objects and JavaScript libraries compiled into the Web build.

Web audio, display, input, fetch, MIDI, filesystem, and WebGL adapters call JavaScript libraries and exchange data with C++ Web platform classes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```javascript
// Update computed latency
			GodotAudio.interval = setInterval(function () {
				let computed_latency = 0;
				if (ctx.baseLatency) {
					computed_latency += GodotAudio.ctx.baseLatency;
				}
				if (ctx.outputLatency) {
					computed_latency += GodotAudio.ctx.outputLatency;
				}
				onlatencyupdate(computed_latency);
			}, 1000);
			GodotOS.atexit(GodotAudio.close_async);

			const path = GodotConfig.locate_file('godot.audio.position.worklet.js');
```

Source: `platform/web/js/libs/library_godot_audio.js` — GodotAudio (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JavaScriptObjectImpl

Evidence:
- Code: `platform/web/audio_driver_web.h:114` — AudioDriverWorklet
- Code: `platform/web/http_client_web.h:58` — HTTPClientWeb
- Code: `platform/web/webmidi_driver.h:37` — MIDIDriverWebMidi
- Code: `platform/web/js/libs/library_godot_audio.js:61` — GodotAudio

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Keyboard and MIDI parsing](#topic-input-midi)

<a id="topic-collada-import"></a>
## 243. COLLADA scene interchange

**Scope:** First-party

**Builds on:** [Scene importing](#topic-scene-importing).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

COLLADA is represented as an intermediate parser state before editor scene construction.

The COLLADA parser supplies a scene importer with image, material, effect, mesh, skin, morph, node, visual-scene, and animation data collected in Collada::State.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
struct ColladaImport {
	Collada collada;
	Node3D *scene = nullptr;

	Vector<Ref<Animation>> animations;

	struct NodeMap {
		//String path;
		Node3D *node = nullptr;
		int bone = -1;
		List<int> anim_tracks;
	};
```

Source: `editor/import/3d/editor_import_collada.cpp` — ColladaImport (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ColladaParseState

Evidence:
- Code: `editor/import/3d/collada.h` — Collada::State
- Code: `editor/import/3d/editor_import_collada.cpp:48` — ColladaImport

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene graph and lifecycle](#topic-scene-tree)

<a id="topic-editing-selection-history"></a>
## 244. Editing selection history

**Scope:** First-party

**Builds on:** [Editor scene sessions](#topic-editor-scene-sessions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The history stores object paths and the currently selected point in that path history.

Selection history gives an editor scene session back-and-forward navigation across selected objects and nested-resource properties.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void EditorSelectionHistory::add_object(ObjectID p_object, const String &p_property, bool p_inspector_only) {
	Object *obj = ObjectDB::get_instance(p_object);
	ERR_FAIL_NULL(obj);
	RefCounted *r = Object::cast_to<RefCounted>(obj);
	_Object o;
	if (r) {
		o.ref = Ref<RefCounted>(r);
	}
	o.object = p_object;
	o.property = p_property;
	o.inspector_only = p_inspector_only;

	bool has_prev = current_elem_idx >= 0 && current_elem_idx < history.size();
```

Source: `editor/editor_data.cpp` — EditorSelectionHistory::add_object (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditedScene

Evidence:
- Code: `editor/editor_data.h:47` — EditorSelectionHistory
- Code: `editor/editor_data.cpp:110` — EditorSelectionHistory::add_object

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Navigation agents and regions](#topic-navigation-agents)
- [Resources and asset lifecycle](#topic-resources)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)

<a id="topic-editor-plugin-state"></a>
## 245. Editor plugin state and custom types

**Scope:** First-party

**Builds on:** [Editor scene sessions](#topic-editor-scene-sessions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

EditorData selects a main editor plugin for an object and retains plugin-provided type metadata.

Editor plugins can handle objects, register custom types, and serialize plugin state into the active scene's editor state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
EditorPlugin *EditorData::get_handling_main_editor(Object *p_object) {
	// We need to iterate backwards so that we can check user-created plugins first.
	// Otherwise, it would not be possible for plugins to handle CanvasItem and Spatial nodes.
	for (int i = editor_plugins.size() - 1; i > -1; i--) {
		if (editor_plugins[i]->has_main_screen() && editor_plugins[i]->handles(p_object)) {
			return editor_plugins[i];
		}
	}

	return nullptr;
}

Vector<EditorPlugin *> EditorData::get_handling_sub_editors(Object *p_object) {
```

Source: `editor/editor_data.cpp` — EditorData::get_handling_main_editor (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditedScene

Evidence:
- Code: `editor/editor_data.h` — EditorData::CustomType
- Code: `editor/editor_data.cpp:260` — EditorData::get_handling_main_editor
- Code: `editor/editor_data.cpp:334` — EditorData::get_scene_editor_states_with_selection

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-export-presets"></a>
## 246. Export presets

**Scope:** First-party

**Builds on:** [Project filesystem index](#topic-filesystem-project-index).

**This prepares you for:** [Export orchestration](#topic-export-orchestration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Preset configuration is target-platform-specific while retaining resource-selection policy.

An export preset selects indexed project files, applies include, exclude, and customized-file rules, and records output path, features, patches, and encryption options.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
bool EditorExportPreset::_set(const StringName &p_name, const Variant &p_value) {
	values[p_name] = p_value;
	EditorExport::singleton->save_presets();
	if (update_visibility.has(p_name)) {
		if (update_visibility[p_name]) {
			update_value_overrides();
			notify_property_list_changed();
		}
		return true;
	}

	return false;
}
```

Source: `editor/export/editor_export_preset.cpp` — EditorExportPreset (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorExportPreset

Evidence:
- Code: `editor/export/editor_export_preset.h:38` — EditorExportPreset
- Code: `editor/export/editor_export_preset.cpp:41` — EditorExportPreset

<a id="topic-input-action-configuration"></a>
## 247. Input action and shortcut configuration

**Scope:** First-party

**Builds on:** [Editor and project settings](#topic-editor-and-project-settings).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation supports keyboard, mouse, and joypad event capture in a shared line-edit control.

Input action configuration depends on editor and project settings because ActionMapEditor, event search, event capture, and input-event configuration edit action-event mappings.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// Separate from ActionMapEditor for code cleanliness and separation of responsibilities.
class InputEventConfigurationDialog : public ConfirmationDialog {
	GDCLASS(InputEventConfigurationDialog, ConfirmationDialog)
private:
	struct IconCache {
		Ref<Texture2D> keyboard;
		Ref<Texture2D> mouse;
		Ref<Texture2D> mouse_left_button;
		Ref<Texture2D> mouse_right_button;
		Ref<Texture2D> mouse_middle_button;
		Ref<Texture2D> mouse_wheel_up;
		Ref<Texture2D> mouse_wheel_down;
		Ref<Texture2D> mouse_wheel_left;
		Ref<Texture2D> mouse_wheel_right;
```

Source: `editor/settings/input_event_configuration_dialog.h` — InputEventConfigurationDialog (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/settings/action_map_editor.h` — ActionMapEditor and ActionInfo
- Code: `editor/settings/input_event_configuration_dialog.h:42` — InputEventConfigurationDialog
- Code: `editor/settings/event_listener_line_edit.cpp` — EventListenerLineEdit input event handling
- External (official, unverified (source anchor not found)): [Using InputEvent — Godot Engine documentation](https://docs.godotengine.org/en/stable/tutorials/inputs/inputevent.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Encoder configuration](#topic-encoder-configuration)
- [Runtime configuration](#topic-runtime-configuration)

<a id="topic-platform-display-adaptation"></a>
## 248. Platform display and window adaptation

**Scope:** First-party

**Builds on:** [Input events and actions](#topic-input-events-actions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The same engine-facing display role is implemented separately for X11, macOS, Windows, and Web.

DisplayServer implementations adapt platform windows and events; the Web adapter constructs InputEvent instances and forwards them to Input.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
if (!Thread::is_main_thread()) {
		callable_mp_static(DisplayServerWeb::_mouse_button_callback).call_deferred(p_pressed, p_button, p_x, p_y, p_modifiers);
		return true;
	}
#endif

	return _mouse_button_callback(p_pressed, p_button, p_x, p_y, p_modifiers);
}

int DisplayServerWeb::_mouse_button_callback(int p_pressed, int p_button, double p_x, double p_y, int p_modifiers) {
	DisplayServerWeb *ds = get_singleton();

	Point2 pos(p_x, p_y);
	Ref<InputEventMouseButton> ev;
```

Source: `platform/web/display_server_web.cpp` — DisplayServerWeb::_mouse_button_callback (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

InputEvent

Evidence:
- Code: `platform/web/display_server_web.cpp:232` — DisplayServerWeb::_mouse_button_callback
- Code: `platform/linuxbsd/x11/display_server_x11.h:104` — DisplayServerX11
- Code: `platform/macos/display_server_macos.h:71` — DisplayServerMacOS
- Code: `platform/windows/display_server_windows.h:195` — DisplayServerWindows

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Input events and actions](#topic-input-actions)
- [Display, camera, video, and movie services](#topic-platform-presentation)
- [Event queue and watches](#topic-sdl-event-queue)

<a id="topic-platform-selective-shader-baking"></a>
## 249. Platform-selective shader baking

**Scope:** First-party

**Builds on:** [Shader editing and language plugins](#topic-shader-editing-and-language-plugins).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The platform split occurs in build manifests and is reflected in three platform adapter classes.

Platform-selective shader baking depends on shader editing and compiles distinct Vulkan, D3D12, or Metal export-plugin sources when matching SCons environment flags are enabled.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ShaderBakerExportPluginPlatformD3D12 : public ShaderBakerExportPluginPlatform {
	GDCLASS(ShaderBakerExportPluginPlatformD3D12, ShaderBakerExportPluginPlatform);

private:
	void *lib_d3d12 = nullptr;

public:
	virtual RenderingShaderContainerFormat *create_shader_container_format(const Ref<EditorExportPlatform> &p_platform, const Ref<EditorExportPreset> &p_preset) override;
	virtual bool matches_driver(const String &p_driver) override;
	virtual ~ShaderBakerExportPluginPlatformD3D12() override;
};
```

Source: `editor/shader/shader_baker/shader_baker_export_plugin_platform_d3d12.h` — ShaderBakerExportPluginPlatformD3D12 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/shader/shader_baker/SCsub` — vulkan, d3d12, and metal conditional source selection
- Code: `editor/shader/shader_baker/shader_baker_export_plugin_platform_d3d12.h:35` — ShaderBakerExportPluginPlatformD3D12
- Code: `editor/shader/shader_baker/shader_baker_export_plugin_platform_metal.h:35` — ShaderBakerExportPluginPlatformMetal
- Code: `editor/shader/shader_baker/shader_baker_export_plugin_platform_vulkan.h:35` — ShaderBakerExportPluginPlatformVulkan

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Metal-cpp object bridge](#topic-metal-cpp-object-bridge)
- [SCons module build configuration](#topic-module-build-configuration)

<a id="topic-post-import-processing"></a>
## 250. Post-import processing

**Scope:** First-party

**Builds on:** [Scene importing](#topic-scene-importing).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation separates parsing a source format from editor-controlled transformations of its imported result.

Post-import processing operates on the imported scene after format conversion and accepts options for node, mesh, material, animation, and skeleton categories.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class PostImportPluginSkeletonRestFixer : public EditorScenePostImportPlugin {
	GDCLASS(PostImportPluginSkeletonRestFixer, EditorScenePostImportPlugin);

public:
	virtual void get_internal_import_options(InternalImportCategory p_category, List<ResourceImporter::ImportOption> *r_options) override;
	virtual Variant get_internal_option_visibility(InternalImportCategory p_category, const String &p_scene_import_type, const String &p_option, const HashMap<StringName, Variant> &p_options) const override;
	virtual void internal_process(InternalImportCategory p_category, Node *p_base_scene, Node *p_node, Ref<Resource> p_resource, const Dictionary &p_options) override;
};
```

Source: `editor/import/3d/post_import_plugin_skeleton_rest_fixer.h` — PostImportPluginSkeletonRestFixer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/import/3d/resource_importer_scene.h:87` — EditorScenePostImport
- Code: `editor/import/3d/resource_importer_scene.h:153` — EditorScenePostImportPlugin::InternalImportCategory
- Code: `editor/import/3d/post_import_plugin_skeleton_rest_fixer.h:35` — PostImportPluginSkeletonRestFixer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene graph and lifecycle](#topic-scene-tree)
- [Skeletons, animation, and ragdolls](#topic-skeletal-ragdoll)

<a id="topic-resource-dependency-resolution"></a>
## 251. Resource dependency resolution

**Scope:** First-party

**Builds on:** [Project filesystem index](#topic-filesystem-project-index).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Stored dependency values can carry type and UID information that is resolved before presentation.

Dependency resolution uses file-index entries and ResourceUID mappings to expose each indexed resource's current dependency paths.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
Vector<String> EditorFileSystemDirectory::get_file_deps(int p_idx) const {
	ERR_FAIL_INDEX_V(p_idx, files.size(), Vector<String>());
	Vector<String> deps;

	for (int i = 0; i < files[p_idx]->deps.size(); i++) {
		String dep = files[p_idx]->deps[i];
		int sep_idx = dep.find("::"); //may contain type information, unwanted
		if (sep_idx != -1) {
			dep = dep.substr(0, sep_idx);
		}
		ResourceUID::ID uid = ResourceUID::get_singleton()->text_to_id(dep);
		if (uid != ResourceUID::INVALID_ID) {
			//return proper dependency resource from uid
```

Source: `editor/file_system/editor_file_system.cpp` — EditorFileSystemDirectory::get_file_deps (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorFileSystemDirectory::FileInfo

Evidence:
- Code: `editor/file_system/editor_file_system.cpp:143` — EditorFileSystemDirectory::get_file_deps
- Code: `editor/file_system/editor_file_system.cpp:287` — EditorFileSystem::_scan_for_uid_directory

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resource and server identifiers](#topic-resource-identifiers)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-scene-tree-and-signal-connections"></a>
## 252. Scene tree and signal connections

**Scope:** First-party

**Builds on:** [2D and 3D scene authoring](#topic-scene-authoring).

### Start with the idea

A game scene is easier to manage as a hierarchy than as one flat list. Parent-child structure lets movement, visibility, and lifetime flow through related objects without every object needing to know about every other object.

The source treats structural scene navigation and signal wiring as editor services.

Scene authoring is accompanied by a SceneTreeEditor and a ConnectionsDock that inspect nodes, signals, methods, connection arguments, and bound values.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class SceneTreeEditor : public Control {
	GDCLASS(SceneTreeEditor, Control);

	EditorSelection *editor_selection = nullptr;

	enum SceneTreeEditorButton {
		BUTTON_SUBSCENE = 0,
		BUTTON_VISIBILITY = 1,
		BUTTON_SCRIPT = 2,
		BUTTON_LOCK = 3,
		BUTTON_GROUP = 4,
		BUTTON_WARNING = 5,
		BUTTON_SIGNALS = 6,
```

Source: `editor/scene/scene_tree_editor.h` — SceneTreeEditor : public Control (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/scene_tree_editor.h:44` — SceneTreeEditor : public Control
- Code: `editor/scene/connections_dialog.h` — ConnectDialog::ConnectionData and ConnectionsDock
- Code: `editor/scene/connections_dialog.cpp` — signal_name and signal_args metadata handling
- External (official, unverified (source anchor not found)): [Using signals — Godot Engine documentation](https://docs.godotengine.org/en/stable/getting_started/step_by_step/signals.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene graph and lifecycle](#topic-scene-tree)
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-undo-redo-history"></a>
## 253. Undo and redo history

**Scope:** First-party

**Builds on:** [Editor scene sessions](#topic-editor-scene-sessions).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The manager separates histories by integer ID and retains reference-counted objects required by undo and redo stacks.

Undo and redo assigns actions and saved versions to a scene session, allowing edits to identify a history and report unsaved state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
bool EditorData::is_scene_changed(int p_idx) {
	if (p_idx == -1) {
		p_idx = current_edited_scene;
	}
	ERR_FAIL_INDEX_V(p_idx, edited_scene.size(), false);

	uint64_t current_scene_version = undo_redo_manager->get_or_create_history(edited_scene[p_idx].history_id).undo_redo->get_version();
	bool is_changed = edited_scene[p_idx].last_checked_version != current_scene_version;
	edited_scene.write[p_idx].last_checked_version = current_scene_version;
	return is_changed;
}

int EditorData::get_scene_history_id_from_path(const String &p_path) const {
```

Source: `editor/editor_data.cpp` — EditorData::is_scene_changed (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorUndoRedoManager::History, EditedScene

Evidence:
- Code: `editor/editor_undo_redo_manager.h` — EditorUndoRedoManager::History
- Code: `editor/editor_data.cpp:463` — EditorData::is_scene_changed

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Editor scene sessions](#topic-editor-scene-sessions)

<a id="topic-wayland-window-backend"></a>
## 254. Wayland window backend

**Scope:** First-party

**Builds on:** [Linux/BSD desktop integration](#topic-linuxbsd-desktop-integration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Wayland state is split between a display server, a protocol thread, and an embedder.

Wayland window handling builds on Linux/BSD platform integration with a dedicated thread, client protocol objects, and window state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class DisplayServerWayland : public DisplayServer {
	GDSOFTCLASS(DisplayServerWayland, DisplayServer);

	struct WindowData {
		DisplayServerEnums::WindowID id = DisplayServerEnums::INVALID_WINDOW_ID;

		DisplayServerEnums::WindowID parent_id = DisplayServerEnums::INVALID_WINDOW_ID;

		// For popups.
		DisplayServerEnums::WindowID root_id = DisplayServerEnums::INVALID_WINDOW_ID;

		// For toplevels.
		List<DisplayServerEnums::WindowID> popup_stack;
```

Source: `platform/linuxbsd/wayland/display_server_wayland.h` — DisplayServerWayland (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WaylandThread.WindowState

Evidence:
- Code: `platform/linuxbsd/wayland/display_server_wayland.h:62` — DisplayServerWayland
- Code: `platform/linuxbsd/wayland/wayland_thread.h:103` — WaylandThread
- Code: `platform/linuxbsd/wayland/wayland_thread.h` — WaylandThread::WindowState
- Code: `platform/linuxbsd/wayland/wayland_embedder.h:87` — WaylandEmbedder

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Task scheduling and synchronization](#topic-task-scheduling)
- [Explicit DRM synchronization objects](#topic-explicit-drm-syncobj)
- [Linux/BSD desktop integration](#topic-linuxbsd-desktop-integration)

<a id="topic-web-javascript-bridge"></a>
## 255. Web JavaScript bridge

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the explicit engine-to-browser object boundary, distinct from browser input and audio adapters.

The JavaScript bridge exposes evaluation, interface lookup, callback creation, object creation, buffer conversion, downloads, PWA operations, and JavaScript object proxies.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
class JavaScriptObjectImpl : public JavaScriptObject {
	GDSOFTCLASS(JavaScriptObjectImpl, JavaScriptObject);

private:
	friend class JavaScriptBridge;

	int _js_id = 0;
	Callable _callable;

	WASM_EXPORT static int _variant2js(const void **p_args, int p_pos, godot_js_wrapper_ex *r_val, void **p_lock);
	WASM_EXPORT static void _free_lock(void **p_lock, int p_type);
	WASM_EXPORT static Variant _js2variant(int p_type, godot_js_wrapper_ex *p_val);
	WASM_EXPORT static void *_alloc_variants(int p_size);
```

Source: `platform/web/javascript_bridge_singleton.cpp` — JavaScriptObjectImpl (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JavaScriptObjectImpl

Evidence:
- Code: `platform/web/api/javascript_bridge_singleton.h:45` — JavaScriptBridge
- Code: `platform/web/javascript_bridge_singleton.cpp:70` — JavaScriptObjectImpl
- Code: `platform/web/js/libs/library_godot_javascript_singleton.js:31` — GodotJSWrapper

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry buffer storage](#topic-buffer-storage)
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [Intersection and occlusion filters](#topic-filter-callbacks)
- [Engine object model](#topic-object-model)
- [Shader interface metadata](#topic-shader-interface-metadata)

<a id="topic-sdl-iostreams"></a>
## 256. Abstract I/O streams

**Scope:** Vendored: sdl

**Builds on:** [Runtime property groups and hints](#topic-sdl-runtime-properties).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The stream implementation also provides fixed-width endian read and write helpers.

SDL I/O streams wrap byte operations for file and memory implementations and expose an SDL property group.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct SDL_IOStream
{
    SDL_IOStreamInterface iface;
    void *userdata;
    SDL_IOStatus status;
    SDL_PropertiesID props;
};

#ifdef SDL_PLATFORM_3DS
#include "n3ds/SDL_iostreamromfs.h"
#endif // SDL_PLATFORM_3DS

#ifdef SDL_PLATFORM_ANDROID
```

Source: `thirdparty/sdl/io/SDL_iostream.c` — SDL_IOStream (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_IOStream, SDL_PropertiesID

Evidence:
- Code: `thirdparty/sdl/io/SDL_iostream.c:48` — SDL_IOStream
- Code: `thirdparty/sdl/io/SDL_iostream.c:1508` — SDL_ReadU32LE
- Code: `thirdparty/sdl/include/SDL3/SDL_iostream.h:31` — SDL_IOStream

<a id="topic-alpha-plane-coding"></a>
## 257. Alpha-plane coding

**Scope:** Vendored: libwebp

**Builds on:** [Picture planes and pixel ownership](#topic-input-picture-planes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Alpha encoding has its own filter choice and can use lossless compression.

The alpha plane is filtered, optionally quantized, and encoded before its bytes are incorporated with the lossy image representation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
(uint32_t*)WebPSafeMalloc(xsize * 3, sizeof(*copy_buffer));
  const int limit_bits = VP8LNearLosslessBits(quality);
  assert(argb_dst != NULL);
  assert(limit_bits > 0);
  assert(limit_bits <= MAX_LIMIT_BITS);
  if (copy_buffer == NULL) {
    return 0;
  }
  // For small icon images, don't attempt to apply near-lossless compression.
  if ((xsize < MIN_DIM_FOR_NEAR_LOSSLESS &&
       ysize < MIN_DIM_FOR_NEAR_LOSSLESS) ||
      ysize < 3) {
    for (i = 0; i < ysize; ++i) {
      memcpy(argb_dst + i * xsize, picture->argb + i * picture->argb_stride,
```

Source: `thirdparty/libwebp/src/enc/near_lossless_enc.c` — VP8LNearLossless (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPPicture, WebPConfig

Evidence:
- Code: `thirdparty/libwebp/src/enc/alpha_enc.c` — VP8EncEncodeAlpha
- Code: `thirdparty/libwebp/src/enc/near_lossless_enc.c` — VP8LNearLossless
- Code: `thirdparty/libwebp/src/webp/encode.h` — WebPConfig.alpha_compression, WebPPicture.a

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Encoder configuration](#topic-encoder-configuration)
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Picture planes and pixel ownership](#topic-input-picture-planes)

<a id="topic-cff-font-subsetting"></a>
## 258. CFF font subsetting

**Scope:** Vendored: harfbuzz

**Builds on:** [Font subsetting](#topic-font-subsetting).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

CFF-specific code is distinct from generic table dispatch because it carries charstring and subroutine structures.

The font subsetting pipeline contains CFF1 and CFF2 plans, accelerators, serializers, and subroutine-subsetting helpers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
namespace OT {
struct cff2_subset_plan
{
  bool create (const OT::cff2::accelerator_subset_t &acc,
	      hb_subset_plan_t *plan)
  {
    /* make sure notdef is first */
    hb_codepoint_t old_glyph;
    if (!plan->old_gid_for_new_gid (0, &old_glyph) || (old_glyph != 0)) return false;

    num_glyphs = plan->num_output_glyphs ();
    orig_fdcount = acc.fdArray->count;

    drop_hints = plan->flags & HB_SUBSET_FLAGS_NO_HINTING;
```

Source: `thirdparty/harfbuzz/src/hb-subset-cff2.cc` — cff2_subset_plan (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_subset_plan_t, hb_subset_accelerator_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-cff-common.hh:541` — cff_subset_accelerator_t
- Code: `thirdparty/harfbuzz/src/hb-subset-cff2.cc:492` — cff2_subset_plan

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Font driver modules](#topic-font-driver-modules)
- [PostScript font services](#topic-postscript-font-services)
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)

<a id="topic-cpu-specialized-dsp"></a>
## 259. CPU-specialized DSP backends

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The scalar sources establish algorithms while architecture files express equivalent work with vector intrinsics or target assembly.

DSP function families have scalar, SSE2/SSE4.1/AVX2, NEON, MSA, and MIPS implementations that are compile-time selected for target capabilities.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
extern VP8CPUInfo VP8GetCPUInfo;
extern void VP8EncDspInitSSE2(void);
extern void VP8EncDspInitSSE41(void);
extern void VP8EncDspInitNEON(void);
extern void VP8EncDspInitMIPS32(void);
extern void VP8EncDspInitMIPSdspR2(void);
extern void VP8EncDspInitMSA(void);

WEBP_DSP_INIT_FUNC(VP8EncDspInit) {
  VP8DspInit();  // common inverse transforms
  InitTables();

  // default C implementations
#if !WEBP_NEON_OMIT_C_CODE
```

Source: `thirdparty/libwebp/src/dsp/enc.c` — VP8EncDspInit (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/libwebp/src/dsp/enc.c:750` — VP8EncDspInit
- Code: `thirdparty/libwebp/src/dsp/enc_sse2.c` — SSE2 transform and quantization kernels
- Code: `thirdparty/libwebp/src/dsp/enc_neon.c` — NEON transform and quantization kernels
- Code: `thirdparty/libwebp/src/dsp/enc_msa.c` — MSA transform and prediction kernels

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Codec SIMD specialization](#topic-codec-simd-specialization)
- [ISA and platform portability](#topic-isa-portability)
- [SIMD ray packets](#topic-simd-ray-packets)
- [Motion blur](#topic-motion-blur)
- [Main loop, OS, and time services](#topic-runtime-loop-time)

<a id="topic-catmull-clark-patch-construction"></a>
## 260. Catmull–Clark patch construction

**Scope:** Vendored: embree

**Builds on:** [Half-edge topology](#topic-half-edge-topology).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Catmull–Clark patch construction reads a half-edge topology and converts its limit data to patch geometry; half-edge topology is needed to explain this construction.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<typename Vertex, typename Vertex_t = Vertex>
    class __aligned(64) CatmullClarkPatchT
    {
    public:
    typedef CatmullClark1RingT<Vertex,Vertex_t> CatmullClark1Ring;
    typedef typename CatmullClark1Ring::Type Type;
    
    array_t<CatmullClark1RingT<Vertex,Vertex_t>,4> ring;
    
    public:
    __forceinline CatmullClarkPatchT () {}

    __forceinline CatmullClarkPatchT (const HalfEdge* first_half_edge, const char* vertices, size_t stride) {
      init(first_half_edge,vertices,stride);
```

Source: `thirdparty/embree/kernels/subdiv/catmullclark_patch.h` — CatmullClarkPatchT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

HalfEdge

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/catmullclark_patch.h:12` — CatmullClarkPatchT
- Code: `thirdparty/embree/kernels/subdiv/catmullclark_patch.h:278` — GeneralCatmullClarkPatchT
- Code: `thirdparty/embree/kernels/subdiv/gregory_patch.h:14` — GregoryPatchT

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Half-edge topology](#topic-half-edge-topology)
- [Geometry resources](#topic-geometry-resources)

<a id="topic-character-encoding-conversion"></a>
## 261. Character-encoding conversion

**Scope:** Vendored: icu4c

**Builds on:** [Unicode data and properties](#topic-unicode-data).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The partition includes UTF-7, UTF-8, UTF-16, UTF-32, ISO-2022, MBCS, and other converter implementations.

ICU models converters with shared static data, an implementation dispatch structure, contexts, and UTF-specific implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const UConverterSharedData _UTF8Data=
        UCNV_IMMUTABLE_SHARED_DATA_INITIALIZER(&_UTF8StaticData, &_UTF8Impl);

/* CESU-8 converter data ---------------------------------------------------- */

static const UConverterImpl _CESU8Impl={
    UCNV_CESU8,

    nullptr,
    nullptr,

    nullptr,
    nullptr,
```

Source: `thirdparty/icu4c/common/ucnv_u8.cpp` — _UTF8Data (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

UConverter, UConverterSharedData

Evidence:
- Code: `thirdparty/icu4c/common/ucnv_bld.h:93` — UConverterSharedData
- Code: `thirdparty/icu4c/common/ucnv_u8.cpp:898` — _UTF8Data

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU block texture conversion](#topic-block-texture-transcoding)
- [Dynamic runtime values](#topic-dynamic-variant)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Endian-safe binary I/O](#topic-endian-safe-binary-io)
- [JPEG decompression and coefficient transcoding](#topic-jpeg-decompression-transcoding)
- [Managed-native interop](#topic-managed-native-interop)
- [Post-import processing](#topic-post-import-processing)
- [Project source migration](#topic-project-upgrade)
- [Web JavaScript bridge](#topic-web-javascript-bridge)
- [YUV/RGB conversion and chroma processing](#topic-yuv-rgb-conversion)
- [OpenXR dispatch forwarding](#topic-openxr-dispatch)

<a id="topic-codec-simd-specialization"></a>
## 262. Codec SIMD specialization

**Scope:** Vendored: etcpak

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Conditional C preprocessing exposes AVX2, SSE4.1, and NEON codec declarations; C preprocessing is needed to explain which declarations are available.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
#ifdef __AVX2__
void DitherAvx2( uint8_t* data, __m128i px0, __m128i px1, __m128i px2, __m128i px3 );
#endif

#endif
```

Source: `thirdparty/etcpak/Dither.hpp` — DitherAvx2 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/etcpak/Dither.hpp:18` — DitherAvx2
- Code: `thirdparty/etcpak/Tables.hpp:30` — g_table_SIMD
- Code: `thirdparty/etcpak/Tables.hpp:44` — g_table128_NEON

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [CPU-specialized DSP backends](#topic-cpu-specialized-dsp)
- [Optional SIMD codec paths](#topic-simd-accelerated-codecs)

<a id="topic-collision-filtering"></a>
## 263. Collision filtering

**Scope:** Vendored: jolt_physics

**Builds on:** [Collision shapes](#topic-collision-shapes).

**This prepares you for:** [Broad-phase collision detection](#topic-broad-phase).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The filtering interfaces let a caller restrict which bodies or subshapes may interact.

Collision filtering applies object-layer, broad-phase-layer, group, body, and shape filters before candidate generation or collision queries.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// - Or if there's no filter for the first object, the second group filter says the objects can collide
class JPH_EXPORT CollisionGroup
{
	JPH_DECLARE_SERIALIZABLE_NON_VIRTUAL(JPH_EXPORT, CollisionGroup)

public:
	using GroupID			= uint32;
	using SubGroupID		= uint32;

	static const GroupID	cInvalidGroup = ~GroupID(0);
	static const SubGroupID	cInvalidSubGroup = ~SubGroupID(0);

	/// Default constructor
							CollisionGroup() = default;
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionGroup.h` — CollisionGroup (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/ObjectLayer.h` — ObjectLayerFilter and ObjectLayerPairFilter
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/ShapeFilter.h` — ShapeFilter and ReversedShapeFilter
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionGroup.h:19` — CollisionGroup

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)
- [Spatial predictive filters](#topic-spatial-predictive-filters)

<a id="topic-sdl-platform-backends"></a>
## 264. Compile-time platform backends

**Scope:** Vendored: sdl

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The source tree contains platform-specific input, loading, timing, synchronization, and device paths.

SDL uses compile-time conditions to select Linux, Windows, macOS, dummy, and other platform backend implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#ifndef SDL_build_config_h_
#define SDL_build_config_h_

#include <SDL3/SDL_platform_defines.h>

/**
 *  \file SDL_build_config.h
 *
 *  This is a set of defines to configure the SDL features
 */

/* Add any platform that doesn't build using the configure system. */
#if defined(SDL_PLATFORM_PRIVATE)
```

Source: `thirdparty/sdl/include/build_config/SDL_build_config.h` — SDL_build_config_h_ (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/sdl/include/build_config/SDL_build_config.h:22` — SDL_build_config_h_
- Code: `thirdparty/sdl/thread/pthread/SDL_systhread_c.h:25` — SYS_ThreadHandle
- Code: `thirdparty/sdl/thread/windows/SDL_systhread_c.h:28` — SYS_ThreadHandle

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Motion blur](#topic-motion-blur)
- [Main loop, OS, and time services](#topic-runtime-loop-time)
- [HID device enumeration and backends](#topic-hid-device-enumeration)
- [Profiling name interning](#topic-profiling-interning)
- [Rendering backend abstraction](#topic-rendering-backends)
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)
- [Vector font export](#topic-vector-font-export)

<a id="topic-compressed-font-stream-adapters"></a>
## 265. Compressed font stream adapters

**Scope:** Vendored: freetype

**Builds on:** [Font streams](#topic-font-streams).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Bzip2 adapter checks the header, initializes a decompressor, resets on backward seeks, and installs read and close callbacks.

Optional Bzip2 and LZW components wrap a compatible source stream and expose decompressed bytes through the same FT_Stream callbacks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
FT_EXPORT_DEF( FT_Error )
  FT_Stream_OpenBzip2( FT_Stream  stream,
                       FT_Stream  source )
  {
    FT_Error      error;
    FT_Memory     memory;
    FT_BZip2File  zip = NULL;


    if ( !stream || !source )
    {
      error = FT_THROW( Invalid_Stream_Handle );
      goto Exit;
    }
```

Source: `thirdparty/freetype/src/bzip2/ftbzip2.c` — FT_Stream_OpenBzip2 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:471` — FT_Stream_OpenBzip2
- Code: `thirdparty/freetype/src/lzw/ftlzw.c:342` — FT_Stream_OpenLZW
- External (official, verified): [BZIP2 Streams - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-bzip2.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-convex-collision"></a>
## 266. Convex support and penetration

**Scope:** Vendored: jolt_physics

**Builds on:** [Collision shapes](#topic-collision-shapes).

**This prepares you for:** [Narrow-phase collision queries](#topic-narrow-phase).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ConvexShape exposes support mappings that collision algorithms can query.

Convex collision represents collision shapes with support functions and uses GJK closest-point and EPA penetration calculations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// - The closest point indicates the position / direction of least penetration
class EPAPenetrationDepth
{
private:
	// Typedefs
	static constexpr int cMaxPoints = EPAConvexHullBuilder::cMaxPoints;
	static constexpr int cMaxPointsToIncludeOriginInHull = 32;
	static_assert(cMaxPointsToIncludeOriginInHull < cMaxPoints);

	using Triangle = EPAConvexHullBuilder::Triangle;
	using Points = EPAConvexHullBuilder::Points;

	/// The GJK algorithm, used to start the EPA algorithm
	GJKClosestPoint		mGJK;
```

Source: `thirdparty/jolt_physics/Jolt/Geometry/EPAPenetrationDepth.h` — EPAPenetrationDepth (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shape

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/ConvexShape.h` — ConvexShape::Support and GetSupportFunction
- Code: `thirdparty/jolt_physics/Jolt/Geometry/GJKClosestPoint.h:21` — GJKClosestPoint
- Code: `thirdparty/jolt_physics/Jolt/Geometry/EPAPenetrationDepth.h:36` — EPAPenetrationDepth

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [native 3D collision pipeline](#topic-physics-3d-collision-pipeline)

<a id="topic-gpu-memory-suballocation"></a>
## 267. D3D12 memory allocation

**Scope:** Vendored: d3d12ma

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The subsystem includes block metadata, pools, budgets, defragmentation, and virtual allocations.

D3D12MA's Allocator, Pool, and VirtualBlock APIs manage resource and virtual allocations using allocation callbacks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
You can dump internal state of the allocator to a string in JSON format using function D3D12MA::Allocator::BuildStatsString().
The result is guaranteed to be correct JSON.
It uses Windows Unicode (UTF-16) encoding.
Any strings provided by user (see D3D12MA::Allocation::SetName())
are copied as-is and properly escaped for JSON.
It must be freed using function D3D12MA::Allocator::FreeStatsString().

The format of this JSON string is not part of official documentation of the library,
but it will not change in backward-incompatible way without increasing library major version number
and appropriate mention in changelog.

The JSON string contains all the data that can be obtained using D3D12MA::Allocator::CalculateStatistics().
It can also contain detailed map of allocated memory blocks and their regions -
```

Source: `thirdparty/d3d12ma/D3D12MemAlloc.h` — D3D12MA::Allocator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/d3d12ma/D3D12MemAlloc.h:269` — D3D12MA::Allocator
- Code: `thirdparty/d3d12ma/D3D12MemAlloc.h:378` — D3D12MA::Pool
- Code: `thirdparty/d3d12ma/D3D12MemAlloc.h:1498` — D3D12MA::VirtualBlock

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-enet-wire-commands"></a>
## 268. ENet wire commands

**Scope:** Vendored: enet

**Builds on:** [ENet host and peer transport](#topic-enet-host-peer-transport).

**This prepares you for:** [Godot ENet socket adaptation](#topic-godot-enet-socket-adaptation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ENet wire commands select fixed protocol layout sizes and drive peer state transitions; host-peer transport is required to interpret the peer state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static const size_t commandSizes [ENET_PROTOCOL_COMMAND_COUNT] =
{
    0,
    sizeof (ENetProtocolAcknowledge),
    sizeof (ENetProtocolConnect),
    sizeof (ENetProtocolVerifyConnect),
    sizeof (ENetProtocolDisconnect),
    sizeof (ENetProtocolPing),
    sizeof (ENetProtocolSendReliable),
    sizeof (ENetProtocolSendUnreliable),
    sizeof (ENetProtocolSendFragment),
    sizeof (ENetProtocolSendUnsequenced),
    sizeof (ENetProtocolBandwidthLimit),
```

Source: `thirdparty/enet/protocol.c` — commandSizes (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ENetPeer

Evidence:
- Code: `thirdparty/enet/enet/protocol.h:41` — ENetProtocolCommand
- Code: `thirdparty/enet/protocol.c:12` — commandSizes
- Code: `thirdparty/enet/protocol.c:30` — enet_protocol_command_size

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [ENet transport and multiplayer adaptation](#topic-enet-transport-and-multiplayer)

<a id="topic-encoder-configuration"></a>
## 269. Encoder configuration

**Scope:** Vendored: libwebp

**Builds on:** [Picture planes and pixel ownership](#topic-input-picture-planes).

**This prepares you for:** [Lossy macroblock encoding](#topic-lossy-macroblock-encoding), [Worker-based parallelism](#topic-worker-parallelism).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Initialization and validation are explicit ABI-aware API steps.

WebPConfig selects lossy or lossless encoding and controls quality, effort, segmentation, filtering, alpha, thread, and memory choices for the configured picture.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
int WebPConfigInitInternal(WebPConfig* config,
                           WebPPreset preset, float quality, int version) {
  if (WEBP_ABI_IS_INCOMPATIBLE(version, WEBP_ENCODER_ABI_VERSION)) {
    return 0;   // caller/system version mismatch!
  }
  if (config == NULL) return 0;

  config->quality = quality;
  config->target_size = 0;
  config->target_PSNR = 0.;
  config->method = 4;
  config->sns_strength = 50;
  config->filter_strength = 60;   // mid-filtering
```

Source: `thirdparty/libwebp/src/enc/config_enc.c` — WebPConfigInitInternal (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPConfig, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — struct WebPConfig, WebPConfigInit, WebPValidateConfig
- Code: `thirdparty/libwebp/src/enc/config_enc.c:27` — WebPConfigInitInternal
- External (official, verified): [WebP API Documentation](https://developers.google.com/speed/webp/docs/api), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Alpha-plane coding](#topic-alpha-plane-coding)
- [Thread synchronization](#topic-concurrency)
- [Android export pipeline](#topic-android-export-pipeline)
- [Apple embedded packaging and signing](#topic-apple-embedded-packaging)
- [Audio buses, streams, and effects](#topic-audio-bus-processing)
- [Contextual completion contracts](#topic-completion-contexts)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Engine bootstrap](#topic-engine-bootstrap)
- [Frame timing and physics stepping](#topic-frame-timing)
- [High-level RPC routing](#topic-high-level-rpc)
- [Input action and shortcut configuration](#topic-input-action-configuration)
- [ISA and platform portability](#topic-isa-portability)
- [OpenXR action configuration](#topic-openxr-action-configuration)
- [PSA key lifecycle and storage](#topic-psa-key-lifecycle)
- [Textures, meshes, and materials](#topic-rendering-resources)
- [Runtime configuration](#topic-runtime-configuration)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [TLS connection and session state](#topic-tls-connection-state)
- [Render-pass configuration](#topic-vulkan-render-pass-configuration)
- [Swapchain presentation](#topic-vulkan-swapchain-presentation)
- [Task scheduling and synchronization](#topic-task-scheduling)

<a id="topic-endian-safe-binary-io"></a>
## 270. Endian-safe binary I/O

**Scope:** Vendored: libwebp

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation uses conditional endianness macros instead of assuming a host byte order.

Endian macros and memory conversion helpers normalize integer fields while container and bitstream code serializes binary bytes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// Create fourcc of the chunk from the chunk tag characters.
#define MKFOURCC(a, b, c, d) ((a) | (b) << 8 | (c) << 16 | (uint32_t)(d) << 24)

// VP8 related constants.
#define VP8_SIGNATURE 0x9d012a              // Signature in VP8 data.
#define VP8_MAX_PARTITION0_SIZE (1 << 19)   // max size of mode partition
#define VP8_MAX_PARTITION_SIZE  (1 << 24)   // max size for token partition
#define VP8_FRAME_HEADER_SIZE 10  // Size of the frame header within VP8 data.

// VP8L related constants.
#define VP8L_SIGNATURE_SIZE          1      // VP8L signature size.
#define VP8L_MAGIC_BYTE              0x2f   // VP8L signature byte.
#define VP8L_IMAGE_SIZE_BITS         14     // Number of bits used to store
                                            // width and height.
```

Source: `thirdparty/libwebp/src/webp/format_constants.h` — MKFOURCC (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPData

Evidence:
- Code: `thirdparty/libwebp/src/utils/endian_inl_utils.h` — HToLE32, HToLE16, BSwap32
- Code: `thirdparty/libwebp/src/webp/format_constants.h:18` — MKFOURCC
- Code: `thirdparty/libwebp/src/mux/muxinternal.c` — MuxAssemble

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [Entropy bitstreams](#topic-entropy-bitstreams)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-freetype-module-composition"></a>
## 271. FreeType module composition

**Scope:** Vendored: freetype

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The FreeType base and autofit modules combine included implementation units into compiled engine modules; C preprocessing is needed to explain the inclusion-based composition.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#define FT_MAKE_OPTION_SINGLE_OBJECT

#include "ft-hb.c"
#include "ft-hb-ft.c"
#include "afadjust.c"
#include "afblue.c"
#include "afcjk.c"
#include "afdummy.c"
#include "afglobal.c"
#include "afgsub.c"
#include "afhints.c"
#include "afindic.c"
#include "aflatin.c"
```

Source: `thirdparty/freetype/src/autofit/autofit.c` — FT_MAKE_OPTION_SINGLE_OBJECT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/freetype/src/base/ftbase.c:19` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/autofit/autofit.c:19` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/include/freetype/config/ftmodule.h:13` — FT_USE_MODULE

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Font driver modules](#topic-font-driver-modules)

<a id="topic-block-texture-transcoding"></a>
## 272. GPU block texture conversion

**Scope:** Vendored: basis_universal

**Builds on:** [Basis texture transcoding](#topic-basis-transcoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The transcoders support a range of destination block formats and uncompressed pixel formats.

Texture conversion code maps ETC1S, UASTC, and ASTC blocks to GPU block or uncompressed target formats, calculating output block counts and storage.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
bool basisu_lowlevel_etc1s_transcoder::transcode_slice(void* pDst_blocks, uint32_t num_blocks_x, uint32_t num_blocks_y, const uint8_t* pImage_data, uint32_t image_data_size, block_format fmt,
		uint32_t output_block_or_pixel_stride_in_bytes, bool bc1_allow_threecolor_blocks, const bool is_video, const bool is_alpha_slice, const uint32_t level_index, const uint32_t orig_width, const uint32_t orig_height, uint32_t output_row_pitch_in_blocks_or_pixels,
		basisu_transcoder_state* pState, bool transcode_alpha, void *pAlpha_blocks, uint32_t output_rows_in_pixels, uint32_t decode_flags)
	{
		// 'pDst_blocks' unused when disabling *all* hardware transcode options
		// (and 'bc1_allow_threecolor_blocks' when disabling DXT)
		BASISU_NOTE_UNUSED(pDst_blocks);
		BASISU_NOTE_UNUSED(bc1_allow_threecolor_blocks);
		BASISU_NOTE_UNUSED(transcode_alpha);
		BASISU_NOTE_UNUSED(pAlpha_blocks);

		assert(g_transcoder_initialized);
		if (!g_transcoder_initialized)
```

Source: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_lowlevel_etc1s_transcoder::transcode_slice (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp:8488` — basisu_lowlevel_etc1s_transcoder::transcode_slice
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp:10047` — basisu_lowlevel_uastc_ldr_4x4_transcoder::transcode_slice

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [ASTC block coding](#topic-astc-block-coding)
- [Block texture encoding](#topic-block-texture-encoding)
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Image codec integration](#topic-image-codec-registration)

<a id="topic-gamepad-mapping"></a>
## 273. Gamepad mapping and classification

**Scope:** Vendored: sdl

**Builds on:** [HID device enumeration and backends](#topic-hid-device-enumeration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

A built-in mapping database and controller-type rules are included.

Gamepad mapping classifies HID enumeration results and maps device controls to SDL gamepad axes and buttons.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// The SDL gamepad structure
struct SDL_Gamepad
{
    SDL_Joystick *joystick _guarded; // underlying joystick device
    int ref_count _guarded;

    const char *name _guarded;
    SDL_GamepadType type _guarded;
    SDL_GamepadFaceStyle face_style _guarded;
    GamepadMapping_t *mapping _guarded;
    int num_bindings _guarded;
    SDL_GamepadBinding *bindings _guarded;
    SDL_GamepadBinding **last_match_axis _guarded;
    Uint8 *last_hat_mask _guarded;
```

Source: `thirdparty/sdl/joystick/SDL_gamepad.c` — SDL_Gamepad (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/joystick/SDL_gamepad.c:59` — SDL_Gamepad
- Code: `thirdparty/sdl/joystick/SDL_gamepad_db.h:31` — s_GamepadMappings
- Code: `thirdparty/sdl/joystick/controller_type.c` — SDL_GetGamepadTypeFromVIDPID

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)

<a id="topic-geometry-families"></a>
## 274. Geometry families

**Scope:** Vendored: embree

**Builds on:** [Geometry resources](#topic-geometry-resources).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The checked configuration directly enables triangle geometry, while the indexed source also contains the other family implementations.

The source defines Geometry subclasses for triangle meshes, quad meshes, curves, line segments, points, grids, subdivision meshes, user geometry, and instances.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/*! represents an array of bicubic bezier curves */
  struct CurveGeometry : public Geometry
  {
    /*! type of this geometry */
    static const Geometry::GTypeMask geom_type = Geometry::MTY_CURVE4;

  public:
    
    /*! bezier curve construction */
    CurveGeometry (Device* device, Geometry::GType gtype);
    
  public:
    void setMask(unsigned mask);
    void setNumTimeSteps (unsigned int numTimeSteps);
```

Source: `thirdparty/embree/kernels/common/scene_curves.h` — CurveGeometry (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMesh, QuadMesh, CurveGeometry, LineSegments, Points, GridMesh, SubdivMesh, UserGeometry

Evidence:
- Code: `thirdparty/embree/kernels/common/scene_triangle_mesh.h:12` — TriangleMesh
- Code: `thirdparty/embree/kernels/common/scene_curves.h:19` — CurveGeometry
- Code: `thirdparty/embree/kernels/common/scene_grid_mesh.h:12` — GridMesh
- Code: `thirdparty/embree/kernels/config.h:13` — EMBREE_GEOMETRY_TRIANGLE

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision shapes](#topic-collision-shapes)
- [Geometry preprocessing](#topic-geometry-preprocessing)
- [Geometry resources](#topic-geometry-resources)
- [Halfedge topology](#topic-halfedge-topology)
- [Triangle provenance](#topic-mesh-provenance)
- [Motion-blur geometry](#topic-motion-blur-geometry)
- [Ray–primitive intersection](#topic-ray-primitive-intersection)
- [SIMD ray packets](#topic-simd-ray-packets)
- [Spatial indexing and ray queries](#topic-spatial-indexing)
- [Subdivision surface evaluation](#topic-subdivision-surface-evaluation)
- [Subgrid intersection](#topic-subgrid-intersection)
- [Triangle intersection algorithms](#topic-triangle-intersection-algorithms)

<a id="topic-geometry-preprocessing"></a>
## 275. Geometry preprocessing

**Scope:** Vendored: jolt_physics

**Builds on:** [Collision shapes](#topic-collision-shapes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The geometry utilities are used to prepare collision-friendly representations from source triangles and vertices.

Geometry preprocessing converts triangle lists to indexed geometry, builds convex hulls, and partitions triangles for spatial acceleration structures used by collision shapes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// A convex hull builder that tries to create hulls as accurately as possible. Used for offline processing.
class JPH_EXPORT ConvexHullBuilder : public NonCopyable
{
public:
	// Forward declare
	class Face;

	/// Class that holds the information of an edge
	class Edge : public NonCopyable
	{
	public:
		JPH_OVERRIDE_NEW_DELETE

		/// Constructor
```

Source: `thirdparty/jolt_physics/Jolt/Geometry/ConvexHullBuilder.h` — ConvexHullBuilder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shape

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Geometry/Indexify.h` — Indexify and Deindexify
- Code: `thirdparty/jolt_physics/Jolt/Geometry/ConvexHullBuilder.h:20` — ConvexHullBuilder
- Code: `thirdparty/jolt_physics/Jolt/TriangleSplitter/TriangleSplitter.h:13` — TriangleSplitter

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Geometry resources](#topic-geometry-resources)

<a id="topic-glyph-caching"></a>
## 276. Glyph and face caching

**Scope:** Vendored: freetype

**Builds on:** [Font driver modules](#topic-font-driver-modules).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Cache entries are organized around manager-owned caches, cache nodes, glyph families, and MRU ordering.

The cache manager holds faces made available by format drivers and caches character maps, glyph images, and small-bitmap nodes under resource limits.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
typedef struct  FTC_ManagerRec_
  {
    FT_Library          library;
    FT_Memory           memory;

    FTC_Node            nodes_list;
    FT_Offset           max_weight;
    FT_Offset           cur_weight;
    FT_UInt             num_nodes;

    FTC_Cache           caches[FTC_MAX_CACHES];
    FT_UInt             num_caches;
```

Source: `thirdparty/freetype/src/cache/ftcmanag.h` — FTC_Manager (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FTC_SNodeRec

Evidence:
- Code: `thirdparty/freetype/src/cache/ftcmanag.h:135` — FTC_Manager
- Code: `thirdparty/freetype/src/cache/ftcsbits.h:31` — FTC_SNodeRec_
- Code: `thirdparty/freetype/src/cache/ftcmru.h:84` — FTC_MruList
- External (official, verified): [Cache Sub-System - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-cache_subsystem.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-glyph-rasterization"></a>
## 277. Glyph rasterization

**Scope:** Vendored: freetype

**Builds on:** [Font driver modules](#topic-font-driver-modules).

**This prepares you for:** [Signed-distance-field glyph rendering](#topic-signed-distance-fields).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The raster and smooth components expose raster-function objects, while other components provide specialized renderers.

Renderer modules turn driver-loaded glyph outlines into monochrome, gray, SDF, or SVG-backed bitmap representations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
* The renderer can be initialized with a call to
   * `ft_standard_raster.raster_new'; a bitmap can be generated
   * with a call to `ft_standard_raster.raster_render'.
   *
   * See the comments and documentation in the file `ftimage.h' for more
   * details on how the raster works.
   *
   */


  /**************************************************************************
   *
   * This is a rewrite of the FreeType 1.x scan-line converter
   *
```

Source: `thirdparty/freetype/src/raster/ftraster.c` — ft_standard_raster (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SVG_RendererRec

Evidence:
- Code: `thirdparty/freetype/src/raster/ftraster.c:35` — ft_standard_raster
- Code: `thirdparty/freetype/src/smooth/ftgrays.c:2134` — ft_grays_raster
- Code: `thirdparty/freetype/src/svg/svgtypes.h:27` — SVG_RendererRec_
- External (official, verified): [Scanline Converter - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-raster.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Graphics pipeline configuration](#topic-vulkan-pipeline-configuration)
- [Signed-distance-field glyph rendering](#topic-signed-distance-fields)

<a id="topic-graphite-rule-execution"></a>
## 278. Graphite SILF rule execution

**Scope:** Vendored: graphite

**Builds on:** [Binary font-table access](#topic-font-table-access).

**This prepares you for:** [Graphite segment shaping](#topic-graphite-shaping).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Rules are loaded from Graphite font data, matched during passes, and evaluated by the machine implementation.

Graphite executes decoded SILF rule constraints and actions through a finite-state matcher and bytecode machine.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
class Machine::Code::decoder
{
public:
    struct limits;
    static const int NUMCONTEXTS = 256;

    decoder(limits & lims, Code &code, enum passtype pt) throw();

    bool        load(const byte * bc_begin, const byte * bc_end);
    void        apply_analysis(instr * const code, instr * code_end);
    byte        max_ref() { return _max_ref; }
    int         out_index() const { return _out_index; }
```

Source: `thirdparty/graphite/src/Code.cpp` — Machine::Code::decoder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Silf, Segment, Slot

Evidence:
- Code: `thirdparty/graphite/src/Silf.cpp` — graphite2::Silf
- Code: `thirdparty/graphite/src/Pass.cpp` — graphite2::Pass
- Code: `thirdparty/graphite/src/Code.cpp:57` — Machine::Code::decoder
- Code: `thirdparty/graphite/src/inc/Rule.h` — Rule, RuleEntry, State, and FiniteStateMachine
- External (official, unverified (no local verification cache)): [Graphite technical overview](https://graphite.sil.org/graphite_techAbout.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GDScript bytecode compilation and execution](#topic-gdscript-bytecode-runtime)

<a id="topic-halfedge-topology"></a>
## 279. Halfedge topology

**Scope:** Vendored: manifold

**Builds on:** [Manifold mesh interchange](#topic-manifold-mesh-data).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Topology is separate from vertex positions and property vertices.

The Manifold mesh kernel represents every triangle boundary with directed Halfedge records whose pairedHalfedge links encode adjacency.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
*/
struct Halfedge {
  int startVert, endVert;
  int pairedHalfedge;
  int propVert;
  bool IsForward() const { return startVert < endVert; }
  bool operator<(const Halfedge& other) const {
    return startVert == other.startVert ? endVert < other.endVert
                                        : startVert < other.startVert;
  }
};

struct Barycentric {
  int tri;
```

Source: `thirdparty/manifold/src/shared.h` — struct Halfedge (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Halfedge, MeshGLP

Evidence:
- Code: `thirdparty/manifold/src/shared.h:119` — struct Halfedge
- Code: `thirdparty/manifold/src/impl.cpp` — PrepHalfedges; Manifold::Impl halfedge_ construction

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Half-edge topology](#topic-half-edge-topology)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)

<a id="topic-isa-portability"></a>
## 280. ISA and platform portability

**Scope:** Vendored: embree

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The source contains SSE2, AVX, AVX2, AVX-512, ARM, and WebAssembly-oriented paths.

Preprocessor-selected ISA configuration chooses SIMD namespaces and headers, while platform shims adapt unavailable WebAssembly control-register operations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

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
- Code: `thirdparty/embree/common/sys/sysinfo.h` — ISA selection macros
- Code: `thirdparty/embree/common/simd/vint4_sse2.h:20` — vint<4>
- Code: `thirdparty/embree/common/simd/vint8_avx2.h:18` — vint<8>
- Code: `thirdparty/embree/common/simd/wasm/emulation.h:6` — _MM_SET_EXCEPTION_MASK

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [CPU-specialized DSP backends](#topic-cpu-specialized-dsp)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [GUI controls and layout](#topic-gui-control-layout)
- [Optional SIMD codec paths](#topic-simd-accelerated-codecs)
- [SIMD ray packets](#topic-simd-ray-packets)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-identifier-spoof-checking"></a>
## 281. Identifier spoof checking

**Scope:** Vendored: icu4c

**Builds on:** [Unicode data and properties](#topic-unicode-data).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation includes confusable data, script sets, check results, and UTF-16 and UTF-8 API paths.

Unicode property checks and configured allowed character or locale sets are held by SpoofImpl and exposed through USpoofChecker APIs.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct USpoofChecker;
/**
 * @stable ICU 4.2
 */
typedef struct USpoofChecker USpoofChecker; /**< typedef for C of USpoofChecker */

struct USpoofCheckResult;
/**
 * @see uspoof_openCheckResult
 * @stable ICU 58
 */
typedef struct USpoofCheckResult USpoofCheckResult;
```

Source: `thirdparty/icu4c/i18n/unicode/uspoof.h` — USpoofChecker (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SpoofImpl, USpoofChecker

Evidence:
- Code: `thirdparty/icu4c/i18n/uspoof_impl.h:55` — SpoofImpl
- Code: `thirdparty/icu4c/i18n/unicode/uspoof.h:54` — USpoofChecker

<a id="topic-image-quality-metrics"></a>
## 282. Image quality metrics

**Scope:** Vendored: libwebp

**Builds on:** [Picture planes and pixel ownership](#topic-input-picture-planes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The encoder exposes aggregate statistics and picture-distortion helpers for reporting and comparison.

PSNR, SSIM, local similarity, and squared-error routines assess encoded or reconstructed image planes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
int WebPPictureDistortion(const WebPPicture* src, const WebPPicture* ref,
                          int type, float results[5]) {
  int w, h, c;
  int ok = 0;
  WebPPicture p0, p1;
  double total_size = 0., total_distortion = 0.;
  if (src == NULL || ref == NULL ||
      src->width != ref->width || src->height != ref->height ||
      results == NULL) {
    return 0;
  }

  VP8SSIMDspInit();
```

Source: `thirdparty/libwebp/src/enc/picture_psnr_enc.c` — WebPPictureDistortion (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPAuxStats, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/dsp/ssim.c` — VP8SSIMGet, VP8SSIM
- Code: `thirdparty/libwebp/src/enc/picture_psnr_enc.c:180` — WebPPictureDistortion
- Code: `thirdparty/libwebp/src/webp/encode.h` — struct WebPAuxStats, WebPPictureDistortion

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Diagnostic expectations](#topic-diagnostic-expectations)
- [SFNT table loading](#topic-sfnt-tables)

<a id="topic-hit-results"></a>
## 283. Intersection results

**Scope:** Vendored: embree

**Builds on:** [Ray query state](#topic-ray-query).

**This prepares you for:** [Intersection and occlusion filters](#topic-filter-callbacks), [Ray packets and streams](#topic-packet-queries), [Primitive intersection](#topic-primitive-intersection).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The intersection distance is maintained through the ray state used by traversal and intersector epilogues.

An RTCRayHit combines a ray query with geometric-normal, barycentric-coordinate, primitive-ID, geometry-ID, and instance-stack hit results.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* Hit structure for a single ray */
struct RTC_ALIGN(16) RTCHit
{
  float Ng_x;          // x coordinate of geometry normal
  float Ng_y;          // y coordinate of geometry normal
  float Ng_z;          // z coordinate of geometry normal

  float u;             // barycentric u coordinate of hit
  float v;             // barycentric v coordinate of hit

  unsigned int primID; // primitive ID
  unsigned int geomID; // geometry ID
  unsigned int instID[RTC_MAX_INSTANCE_LEVEL_COUNT]; // instance ID
#if defined(RTC_GEOMETRY_INSTANCE_ARRAY)
```

Source: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCHit (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCHit, RTCRayHit

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:30` — RTCHit
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:48` — RTCRayHit
- Code: `thirdparty/embree/kernels/geometry/intersector_epilog.h:21` — Intersect1Epilog1

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry resources](#topic-geometry-resources)
- [Ray–primitive intersection](#topic-ray-primitive-intersection)
- [Triangle intersection algorithms](#topic-triangle-intersection-algorithms)
- [Instancing](#topic-instancing)

<a id="topic-jpeg-decompression-transcoding"></a>
## 284. JPEG decompression and coefficient transcoding

**Scope:** Vendored: libjpeg-turbo

**Builds on:** [Image decode pipelines](#topic-image-decode-pipeline).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The normal decompression path is modular; the coefficient path is explicitly documented as a transcoding mode.

The JPEG code implements an image decode pipeline that selects decompression modules, can merge chroma upsampling with YCC-to-RGB conversion, and can expose raw DCT coefficient arrays for transcoding.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
GLOBAL(void)
_jinit_merged_upsampler(j_decompress_ptr cinfo)
{
  my_merged_upsample_ptr upsample;

  if (cinfo->data_precision != BITS_IN_JSAMPLE)
    ERREXIT1(cinfo, JERR_BAD_PRECISION, cinfo->data_precision);

  upsample = (my_merged_upsample_ptr)
    (*cinfo->mem->alloc_small) ((j_common_ptr)cinfo, JPOOL_IMAGE,
                                sizeof(my_merged_upsampler));
  cinfo->upsample = (struct jpeg_upsampler *)upsample;
  upsample->pub.start_pass = start_pass_merged_upsample;
  upsample->pub.need_context_rows = FALSE;
```

Source: `thirdparty/libjpeg-turbo/src/jdmerge.c` — jinit_merged_upsampler (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JPEG Decompression Context, JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jdmerge.c` — jinit_merged_upsampler
- Code: `thirdparty/libjpeg-turbo/src/jdtrans.c:49` — jpeg_read_coefficients
- External (official, unverified (no local verification cache)): [libjpeg-turbo Documentation](https://libjpeg-turbo.org/Documentation/Documentation), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [HDR, JPEG, and KTX loading](#topic-image-format-loading)
- [Image decode pipelines](#topic-image-decode-pipeline)
- [Raster image input](#topic-raster-image-input)

<a id="topic-ktx-texture-container"></a>
## 285. KTX texture containers

**Scope:** Vendored: libktx

**Builds on:** [Texture format descriptions](#topic-texture-format-description).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

KTX1 and KTX2 are represented by related texture objects with internal constructors and type-specific vtables.

KTX texture containers carry a texture format description, per-level indexing, stream or memory-backed data, virtual dispatch tables, and KTX2 supercompression state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
KTX_error_code
ktxTexture2_transcodeUastc(ktxTexture2* This,
                           alpha_content_e alphaContent,
                           ktxTexture2* prototype,
                           ktx_transcode_fmt_e outputFormat,
                           ktx_transcode_flags transcodeFlags);

/**
 * @memberof ktxTexture2
 * @ingroup reader
 * @~English
 * @brief Transcode a KTX2 texture with BasisLZ/ETC1S or UASTC images.
 *
 * If the texture contains BasisLZ supercompressed images, Inflates them from
```

Source: `thirdparty/libktx/lib/basis_transcode.cpp` — ktxTexture2_transcodeUastc (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

KTX2 Texture, KTX2 Private Texture State, KTX Level Index Entry

Evidence:
- Code: `thirdparty/libktx/include/ktx.h:280` — ktxTexture
- Code: `thirdparty/libktx/lib/texture2.h:33` — ktxTexture2_private
- Code: `thirdparty/libktx/lib/basis_transcode.cpp:55` — ktxTexture2_transcodeUastc
- External (official, unverified (source anchor not found)): [KTX - GPU Texture Container Format](https://www.khronos.org/ktx/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [HDR, JPEG, and KTX loading](#topic-image-format-loading)
- [KTX2 container transcoding](#topic-ktx2-container-transcoding)
- [OpenXR dispatch forwarding](#topic-openxr-dispatch)
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-lossless-transform-pipeline"></a>
## 286. Lossless pixel transform pipeline

**Scope:** Vendored: libwebp

**Builds on:** [Picture planes and pixel ownership](#topic-input-picture-planes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Transform choice is analyzed from the input and recorded as lossless coding features.

The lossless encoder transforms ARGB pixels through predictor, subtract-green, cross-color, and color-indexing stages before entropy coding.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// qualities.
int VP8LResidualImage(int width, int height, int min_bits, int max_bits,
                      int low_effort, uint32_t* const argb,
                      uint32_t* const argb_scratch, uint32_t* const image,
                      int near_lossless_quality, int exact,
                      int used_subtract_green, const WebPPicture* const pic,
                      int percent_range, int* const percent,
                      int* const best_bits) {
  int percent_start = *percent;
  const int max_quantization = 1 << VP8LNearLosslessBits(near_lossless_quality);
  if (low_effort) {
    const int tiles_per_row = VP8LSubSampleSize(width, max_bits);
    const int tiles_per_col = VP8LSubSampleSize(height, max_bits);
    int i;
```

Source: `thirdparty/libwebp/src/enc/predictor_enc.c` — VP8LResidualImage (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPPicture, VP8Encoder

Evidence:
- Code: `thirdparty/libwebp/src/enc/vp8l_enc.c` — VP8LEncodeImage, ApplySubtractGreenTransform
- Code: `thirdparty/libwebp/src/dsp/lossless.c` — VP8LSubtractGreenFromBlueAndRed, VP8LTransformColorInverse
- Code: `thirdparty/libwebp/src/enc/predictor_enc.c:771` — VP8LResidualImage

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU command encoding](#topic-gpu-command-encoding)
- [Entropy bitstreams](#topic-entropy-bitstreams)
- [WebP chunk parsing, incremental decode, and animation](#topic-webp-riff-decoding)

<a id="topic-metal-cpp-object-bridge"></a>
## 287. Metal-cpp object bridge

**Scope:** Vendored: metal-cpp

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [GPU resources and pipelines](#topic-gpu-resources-pipelines).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The wrapper layer covers Metal, Metal 4, MetalFX, and QuartzCore headers.

The Metal portion models framework objects as C++ wrapper classes whose inline methods send selectors to the underlying object.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
_MTL_INLINE MTL4::BinaryFunction* MTL4::Archive::newBinaryFunction(const MTL4::BinaryFunctionDescriptor* descriptor, NS::Error** error)
{
    return Object::sendMessage<MTL4::BinaryFunction*>(this, _MTL_PRIVATE_SEL(newBinaryFunctionWithDescriptor_error_), descriptor, error);
}

_MTL_INLINE MTL::ComputePipelineState* MTL4::Archive::newComputePipelineState(const MTL4::ComputePipelineDescriptor* descriptor, NS::Error** error)
{
    return Object::sendMessage<MTL::ComputePipelineState*>(this, _MTL_PRIVATE_SEL(newComputePipelineStateWithDescriptor_error_), descriptor, error);
}

_MTL_INLINE MTL::ComputePipelineState* MTL4::Archive::newComputePipelineState(const MTL4::ComputePipelineDescriptor* descriptor, const MTL4::PipelineStageDynamicLinkingDescriptor* dynamicLinkingDescriptor, NS::Error** error)
{
    return Object::sendMessage<MTL::ComputePipelineState*>(this, _MTL_PRIVATE_SEL(newComputePipelineStateWithDescriptor_dynamicLinkingDescriptor_error_), descriptor, dynamicLinkingDescriptor, error);
```

Source: `thirdparty/metal-cpp/Metal/MTL4Archive.hpp` — MTL4::Archive::newBinaryFunction (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MTL4::Archive

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTL4Archive.hpp:65` — MTL4::Archive::newBinaryFunction
- External (official, verified): [Get started with Metal-cpp](https://developer.apple.com/metal/cpp/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Metal drawable presentation](#topic-metal-presentation)
- [MetalFX scaling and interpolation](#topic-metalfx-scaling)
- [Platform-selective shader baking](#topic-platform-selective-shader-baking)
- [Rendering backend abstraction](#topic-rendering-backends)
- [Engine object model](#topic-object-model)

<a id="topic-navmesh-contours-polygons"></a>
## 288. Navigation contours and polygons

**Scope:** Vendored: recastnavigation

**Builds on:** [Navigation heightfields](#topic-navmesh-heightfields).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Recast extracts contours from a compact heightfield, then builds polygon and detail meshes from those contours.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/// @see rcAllocContourSet, rcCompactHeightfield, rcContourSet, rcConfig
bool rcBuildContours(rcContext* ctx, const rcCompactHeightfield& chf,
					 const float maxError, const int maxEdgeLen,
					 rcContourSet& cset, const int buildFlags)
{
	rcAssert(ctx);
	
	const int w = chf.width;
	const int h = chf.height;
	const int borderSize = chf.borderSize;
	
	rcScopedTimer timer(ctx, RC_TIMER_BUILD_CONTOURS);
	
	rcVcopy(cset.bmin, chf.bmin);
```

Source: `thirdparty/recastnavigation/Recast/Source/RecastContour.cpp` — rcBuildContours (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

rcContourSet, rcPolyMesh, rcPolyMeshDetail

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Source/RecastContour.cpp:823` — rcBuildContours
- Code: `thirdparty/recastnavigation/Recast/Source/RecastMesh.cpp:989` — rcBuildPolyMesh

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Compact heightfield representation](#topic-compact-heightfield)

<a id="topic-navigation-regions"></a>
## 289. Navigation region construction

**Scope:** Vendored: recastnavigation

**Builds on:** [Compact heightfield representation](#topic-compact-heightfield).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation contains watershed, monotone, and layer-oriented region-building paths.

Region construction labels connected compact heightfield spans into navigation regions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/// @see rcCompactHeightfield, rcCompactSpan, rcBuildDistanceField, rcBuildRegionsMonotone, rcConfig
bool rcBuildRegionsMonotone(rcContext* ctx, rcCompactHeightfield& chf,
							const int borderSize, const int minRegionArea, const int mergeRegionArea)
{
	rcAssert(ctx);
	
	rcScopedTimer timer(ctx, RC_TIMER_BUILD_REGIONS);
	
	const int w = chf.width;
	const int h = chf.height;
	unsigned short id = 1;
	
	rcScopedDelete<unsigned short> srcReg((unsigned short*)rcAlloc(sizeof(unsigned short)*chf.spanCount, RC_ALLOC_TEMP));
	if (!srcReg)
```

Source: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp` — rcBuildRegions (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

rcCompactHeightfield

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp:1247` — rcBuildRegions
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp:1247` — rcBuildRegionsMonotone
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp:1661` — rcBuildLayerRegions

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Navigation agents and regions](#topic-navigation-agents)
- [Navigation maps and path queries](#topic-navigation)
- [Navigation map composition](#topic-navigation-map-composition)

<a id="topic-openexr-image-decoding"></a>
## 290. OpenEXR image decoding

**Scope:** Vendored: tinyexr

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied partition demonstrates integration of a single-header image decoder rather than a ThorVG loader implementation.

TinyEXR is compiled as a header implementation with zlib included first, and its public header defines EXR-oriented image API data and functions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
#define TINYEXR_IMPLEMENTATION
#include "tinyexr.h"
```

Source: `thirdparty/tinyexr/tinyexr.cc` — TINYEXR_IMPLEMENTATION (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EXRImage

Evidence:
- Code: `thirdparty/tinyexr/tinyexr.cc:9` — TINYEXR_IMPLEMENTATION
- Code: `thirdparty/tinyexr/tinyexr.h` — EXRImage, EXRHeader

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [zlib stream compression](#topic-zlib-stream-codec)

<a id="topic-opentype-table-routing"></a>
## 291. OpenType table routing

**Scope:** Vendored: harfbuzz

**Builds on:** [Font subsetting](#topic-font-subsetting).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Separate dispatchers cover color, layout, ordinary glyph-related, and variation tables.

Font subsetting routes recognized OpenType tags to typed table subsetters; color routing explicitly skips CBDT because CBLC handles it.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
bool _hb_subset_table_color		(hb_subset_plan_t *plan, hb_vector_t<char> &buf, hb_tag_t tag, bool *success)
{
#ifndef HB_NO_COLOR
  switch (tag)
  {
  case HB_TAG('s','b','i','x'): *success = _hb_subset_table<const OT::sbix> (plan, buf); return true;
  case HB_TAG('C','O','L','R'): *success = _hb_subset_table<const OT::COLR> (plan, buf); return true;
  case HB_TAG('C','P','A','L'): *success = _hb_subset_table<const OT::CPAL> (plan, buf); return true;
  case HB_TAG('C','B','L','C'): *success = _hb_subset_table<const OT::CBLC> (plan, buf); return true;
  case HB_TAG('C','B','D','T'): *success = true; return true; /* skip CBDT, handled by CBLC */
  }
#endif
  return false;
```

Source: `thirdparty/harfbuzz/src/hb-subset-table-color.cc` — _hb_subset_table_color (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-table-color.cc:8` — _hb_subset_table_color
- Code: `thirdparty/harfbuzz/src/hb-subset-table-layout.cc:9` — _hb_subset_table_layout

<a id="topic-openxr-runtime-loading"></a>
## 292. OpenXR runtime loading

**Scope:** Vendored: openxr

**Builds on:** [OpenXR structure chains](#topic-openxr-structure-chains).

**This prepares you for:** [OpenXR dispatch forwarding](#topic-openxr-dispatch).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The embedded loader reads runtime and API-layer manifests, validates `next` chains during instance creation, and creates a loader instance.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
// Responsible for finding and parsing Runtime-specific manifest files.
class RuntimeManifestFile : public ManifestFile {
   public:
    // Factory method
    static XrResult FindManifestFiles(const std::string &openxr_command,
                                      std::vector<std::unique_ptr<RuntimeManifestFile>> &manifest_files);

   private:
    RuntimeManifestFile(const std::string &filename, const std::string &library_path);
    static void CreateIfValid(const std::string &filename, std::vector<std::unique_ptr<RuntimeManifestFile>> &manifest_files);
    static void CreateIfValid(const Json::Value &root_node, const std::string &filename,
                              std::vector<std::unique_ptr<RuntimeManifestFile>> &manifest_files);
};
```

Source: `thirdparty/openxr/src/loader/manifest_file.hpp` — RuntimeManifestFile (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ManifestFile, LoaderInstance, XrInstanceCreateInfo

Evidence:
- Code: `thirdparty/openxr/src/loader/manifest_file.hpp:69` — RuntimeManifestFile
- Code: `thirdparty/openxr/src/loader/loader_instance.hpp` — LoaderInstance::CreateLoaderInstance

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Instancing](#topic-instancing)
- [Diagnostic expectations](#topic-diagnostic-expectations)
- [Dynamic runtime values](#topic-dynamic-variant)
- [Engine bootstrap](#topic-engine-bootstrap)
- [Engine object model](#topic-object-model)
- [OpenXR composition layers](#topic-openxr-composition-layers)
- [OpenXR extension wrappers](#topic-openxr-extension-wrappers)
- [OpenXR render models](#topic-openxr-render-models)
- [OpenXR runtime integration](#topic-openxr-runtime-integration)
- [Pseudo-random generation](#topic-random-generation)
- [Class metadata and reflection](#topic-reflection)
- [Runtime configuration](#topic-runtime-configuration)
- [Main loop, OS, and time services](#topic-runtime-loop-time)
- [Runtime class metadata](#topic-runtime-metadata)
- [Scene graph nodes](#topic-scene-graph)
- [Event queue and watches](#topic-sdl-event-queue)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)
- [Vulkan instance setup](#topic-vulkan-instance-setup)

<a id="topic-simd-accelerated-codecs"></a>
## 293. Optional SIMD codec paths

**Scope:** Vendored: libjpeg-turbo

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Selection occurs through feature macros, runtime checks where implemented, and function-pointer assignment.

The codec libraries use C conditional compilation to select architecture-specific NEON, SSE2, VSX, LSX, MMX, and other optimized routines, while retaining scalar fallbacks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
void
png_init_filter_functions_sse2(png_structp pp, unsigned int bpp)
{
   /* The techniques used to implement each of these filters in SSE operate on
    * one pixel at a time.
    * So they generally speed up 3bpp images about 3x, 4bpp images about 4x.
    * They can scale up to 6 and 8 bpp images and down to 2 bpp images,
    * but they'd not likely have any benefit for 1bpp images.
    * Most of these can be implemented using only MMX and 64-bit registers,
    * but they end up a bit slower than using the equally-ubiquitous SSE2.
   */
   png_debug(1, "in png_init_filter_functions_sse2");
   if (bpp == 3)
   {
```

Source: `thirdparty/libpng/intel/intel_init.c` — png_init_filter_functions_sse2 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jsimd.h:15` — WITH_SIMD
- Code: `thirdparty/libpng/intel/intel_init.c:19` — png_init_filter_functions_sse2
- Code: `thirdparty/libtheora/x86/x86enc.c:21` — oc_enc_accel_init_x86

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Codec SIMD specialization](#topic-codec-simd-specialization)
- [ISA and platform portability](#topic-isa-portability)
- [SIMD ray packets](#topic-simd-ray-packets)

<a id="topic-png-stream-transforms"></a>
## 294. PNG streaming I/O and row transforms

**Scope:** Vendored: libpng

**Builds on:** [Image decode pipelines](#topic-image-decode-pipeline).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

I/O customization and pixel transformations are separate from information-structure mutation.

The PNG code implements an image decode pipeline with replaceable read and write callbacks, push-mode input states, metadata validity flags, and row-level transformations such as BGR mapping and 16-bit byte swapping.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
{
      case PNG_READ_SIG_MODE:
      {
         png_push_read_sig(png_ptr, info_ptr);
         break;
      }

      case PNG_READ_CHUNK_MODE:
      {
         png_push_read_chunk(png_ptr, info_ptr);
         break;
      }

      case PNG_READ_IDAT_MODE:
```

Source: `thirdparty/libpng/pngpread.c` — PNG_READ_SIG_MODE (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PNG Information State

Evidence:
- Code: `thirdparty/libpng/pngrio.c:49` — png_default_read_data
- Code: `thirdparty/libpng/pngpread.c:18` — PNG_READ_SIG_MODE
- Code: `thirdparty/libpng/pngtrans.c:34` — png_set_swap
- External (official, unverified (source anchor not found)): [Portable Network Graphics (PNG) Specification (Third Edition)](https://www.w3.org/TR/png-3/), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [Image decode pipelines](#topic-image-decode-pipeline)
- [PNG image codec](#topic-png-image-codec)
- [Raster image input](#topic-raster-image-input)
- [Class metadata and reflection](#topic-reflection)

<a id="topic-network-data-exchange"></a>
## 295. Packets, HTTP, JSON, and JSON-RPC

**Scope:** First-party

**Builds on:** [Dynamic values and dictionaries](#topic-dynamic-values).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These APIs separate transport helpers from the high-level multiplayer interface.

Packet peers transmit raw bytes or Variant values, while JSON and JSON-RPC map values into external message documents.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<description>
				Given a Dictionary which takes the form of a JSON-RPC request: unpack the request and run it. Methods are resolved by looking at the field called "method" and looking for an equivalently named function in the JSONRPC object. If one is found that method is called.
				To add new supported methods extend the JSONRPC class and call [method process_action] on your subclass.
				[param action]: The action to be run, as a Dictionary in the form of a JSON-RPC request or notification.
			</description>
		</method>
		<method name="process_string">
			<return type="String" />
			<param index="0" name="action" type="String" />
			<description>
			</description>
		</method>
		<method name="set_method">
			<return type="void" />
```

Source: `doc/classes/JSONRPC.xml` — JSONRPC (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

JSON-RPC document, Variant

Evidence:
- Code: `doc/classes/PacketPeer.xml:2` — PacketPeer
- Code: `doc/classes/HTTPClient.xml` — HTTPClient.query_string_from_dict
- Code: `doc/classes/JSONRPC.xml:2` — JSONRPC

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic runtime values](#topic-dynamic-variant)
- [High-level RPC routing](#topic-high-level-rpc)
- [WebRTC data channels](#topic-webrtc-data-channels)
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-physics-queries"></a>
## 296. Physics queries and motion tests

**Scope:** First-party

**Builds on:** [Physics spaces, bodies, shapes, and joints](#topic-physics-spaces).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Direct-space-state calls provide ad hoc queries, while cast nodes keep query behavior in the scene tree.

Physics queries use physics spaces to submit configured point, ray, shape, and motion tests and receive collision data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="PhysicsDirectSpaceState2D" inherits="Object" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Provides direct access to a physics space in the [PhysicsServer2D].
	</brief_description>
	<description>
		Provides direct access to a physics space in the [PhysicsServer2D]. It's used mainly to do queries against objects and areas residing in a given space.
		[b]Note:[/b] This class is not meant to be instantiated directly. Use [member World2D.direct_space_state] to get the world's physics 2D space state.
	</description>
	<tutorials>
		<link title="Physics introduction">$DOCS_URL/tutorials/physics/physics_introduction.html</link>
		<link title="Ray-casting">$DOCS_URL/tutorials/physics/ray-casting.html</link>
	</tutorials>
	<methods>
```

Source: `doc/classes/PhysicsDirectSpaceState2D.xml` — PhysicsDirectSpaceState2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D

Evidence:
- Code: `doc/classes/PhysicsDirectSpaceState2D.xml:2` — PhysicsDirectSpaceState2D
- Code: `doc/classes/PhysicsShapeQueryParameters2D.xml:2` — PhysicsShapeQueryParameters2D
- Code: `doc/classes/PhysicsTestMotionResult2D.xml:2` — PhysicsTestMotionResult2D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision shapes](#topic-collision-shapes)
- [Motion blur](#topic-motion-blur)
- [3D physics query contracts](#topic-physics-space-queries)
- [Ray query state](#topic-ray-query)

<a id="topic-postscript-font-services"></a>
## 297. PostScript font services

**Scope:** Vendored: freetype

**Builds on:** [Font driver modules](#topic-font-driver-modules).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied sources include Type 1 and Type 2 decoding, CFF decoding, parser tables, and a separate hinter module.

PSAux supplies parsing, decoding, character maps, and hint services shared by CFF, CID, and Type 1 drivers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static
  const PSAux_Interface  psaux_interface =
  {
    &ps_table_funcs,
    &ps_parser_funcs,
    &t1_builder_funcs,
    &t1_decoder_funcs,
    t1_decrypt,
    cff_random,
    ps_decoder_init,
    t1_make_subfont,

    (const T1_CMap_ClassesRec*) &t1_cmap_classes,
```

Source: `thirdparty/freetype/src/psaux/psauxmod.c` — psaux_interface (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/freetype/src/psaux/psauxmod.c:150` — psaux_interface
- Code: `thirdparty/freetype/src/pshinter/pshrec.h:134` — PS_HintsRec
- External (official, unverified (source anchor not found)): [The Type 1 and CID Drivers - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-t1_cid_driver.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [CFF font subsetting](#topic-cff-font-subsetting)

<a id="topic-primitive-references"></a>
## 298. Primitive references

**Scope:** Vendored: embree

**Builds on:** [Geometry resources](#topic-geometry-resources).

**This prepares you for:** [BVH split heuristics](#topic-builder-heuristics), [BVH construction](#topic-bvh-construction).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Motion-blur variants carry linear bounds and time-segment information.

A PrimRef supplies bounds plus geometry and primitive identifiers so builders can partition scene primitives independently of their source geometry layout.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/*! A primitive reference stores the bounds of the primitive and its ID. */
  struct __aligned(32) PrimRef 
  {
    __forceinline PrimRef () {}

#if defined(__AVX__)
    __forceinline PrimRef(const PrimRef& v) { 
      vfloat8::store((float*)this,vfloat8::load((float*)&v));
    }
    __forceinline PrimRef& operator=(const PrimRef& v) { 
      vfloat8::store((float*)this,vfloat8::load((float*)&v)); return *this;
    }
#endif
```

Source: `thirdparty/embree/kernels/builders/primref.h` — PrimRef (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PrimRef, PrimRefMB, SubGridBuildData

Evidence:
- Code: `thirdparty/embree/kernels/builders/primref.h:11` — PrimRef
- Code: `thirdparty/embree/kernels/builders/primref_mb.h:15` — PrimRefMB
- Code: `thirdparty/embree/kernels/builders/primrefgen.cpp` — PrimRef construction

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Broad-phase collision detection](#topic-broad-phase)
- [Geometry resources](#topic-geometry-resources)
- [Instancing](#topic-instancing)
- [Motion blur](#topic-motion-blur)
- [Spatial indexing and ray queries](#topic-spatial-indexing)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-raster-font-rendering"></a>
## 299. Raster font rendering

**Scope:** Vendored: harfbuzz

**Builds on:** [Color-font paint processing](#topic-color-font-paint).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The raster implementation includes PNG-related image code as well as scan-conversion-oriented drawing code.

Color paint records are rasterized through image, draw, paint, clipping, edge, and glyph-extents structures.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/* hb_raster_image_t — pixel artifact */
struct hb_raster_image_t
{
  hb_object_header_t  header;

  hb_vector_t<uint8_t> buffer;
  hb_raster_extents_t  extents     = {};
  hb_raster_format_t   format      = HB_RASTER_FORMAT_A8;

  HB_INTERNAL static unsigned bytes_per_pixel (hb_raster_format_t format);
  HB_INTERNAL bool configure (hb_raster_format_t format, hb_raster_extents_t extents);
  HB_INTERNAL bool deserialize_from_png (hb_blob_t *png);
  HB_INTERNAL hb_blob_t *serialize_to_png_or_fail () const;
  HB_INTERNAL void clear ();
```

Source: `thirdparty/harfbuzz/src/hb-raster-image.hh` — hb_raster_image_t (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_raster_image_t, hb_raster_paint_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-raster-image.hh:37` — hb_raster_image_t
- Code: `thirdparty/harfbuzz/src/hb-raster-paint.hh:102` — hb_raster_paint_t

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Image decode pipelines](#topic-image-decode-pipeline)
- [Raster image input](#topic-raster-image-input)

<a id="topic-regex-jit-codegen"></a>
## 300. Regular-expression JIT code generation

**Scope:** Vendored: pcre2

**Builds on:** [Regular-expression compilation and matching](#topic-regex-compile-match).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

PCRE2 JIT compilation turns compiled pattern control flow into architecture-specific generated code using SLJIT labels and jumps.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
PCRE2_EXP_DEFN int PCRE2_CALL_CONVENTION
pcre2_jit_compile(pcre2_code *code, uint32_t options)
{
pcre2_real_code *re = (pcre2_real_code *)code;
#ifdef SUPPORT_JIT
void *exec_memory;
executable_functions *functions;
static int executable_allocator_is_working = -1;

if (executable_allocator_is_working == -1)
  {
  /* Checks whether the executable allocator is working. This check
     might run multiple times in multi-threaded environments, but the
     result should not be affected by it. */
```

Source: `thirdparty/pcre2/src/pcre2_jit_compile.c` — pcre2_jit_compile (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

sljit_compiler, sljit_jump, sljit_label

Evidence:
- Code: `thirdparty/pcre2/src/pcre2_jit_compile.c:14176` — pcre2_jit_compile
- Code: `thirdparty/pcre2/deps/sljit/sljit_src/sljitLir.h:484` — struct sljit_jump

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GUI controls and layout](#topic-gui-control-layout)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-sfnt-tables"></a>
## 301. SFNT table loading

**Scope:** Vendored: freetype

**Builds on:** [Font streams](#topic-font-streams).

**This prepares you for:** [OpenType and TrueTypeGX validation](#topic-font-table-validation), [TrueType GX variation data](#topic-variable-font-data).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The component is shared infrastructure for OpenType and TrueType-related font handling.

SFNT services read structured tables from font streams, including cmap, metrics, embedded bitmaps, color layers, SVG documents, and WOFF/WOFF2 data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static const TT_CMap_Class  tt_cmap_classes[] =
  {
#undef  TTCMAPCITEM
#define TTCMAPCITEM( a )  &a,
#include "ttcmapc.h"
    NULL,
  };


  /* parse the `cmap' table and build the corresponding TT_CMap objects */
  /* in the current face                                                */
  /*                                                                    */
  FT_LOCAL_DEF( FT_Error )
```

Source: `thirdparty/freetype/src/sfnt/ttcmap.c` — tt_cmap_classes (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

FT_Stream, PFR_FaceRec

Evidence:
- Code: `thirdparty/freetype/src/sfnt/ttload.h:43` — tt_face_load_font_dir
- Code: `thirdparty/freetype/src/sfnt/ttcmap.c:3767` — tt_cmap_classes
- Code: `thirdparty/freetype/src/sfnt/ttsvg.h:35` — tt_face_load_svg_doc
- External (official, verified): [TrueType Tables - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Binary font-table access](#topic-font-table-access)
- [Image quality metrics](#topic-image-quality-metrics)

<a id="topic-spirv-generation"></a>
## 302. SPIR-V generation

**Scope:** Vendored: glslang

**Builds on:** [Shader-language front end](#topic-shader-language-front-end).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The lowering code, builder, and IR classes together form the shader binary-generation path.

A traverser lowers the front end's typed intermediate tree to SPIR-V instructions and a module/function/block graph.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
spv::Builder::AccessChain::CoherentFlags TGlslangToSpvTraverser::TranslateCoherent(const glslang::TType& type)
{
    spv::Builder::AccessChain::CoherentFlags flags = {};
    flags.coherent = type.getQualifier().coherent;
    flags.devicecoherent = type.getQualifier().devicecoherent;
    flags.queuefamilycoherent = type.getQualifier().queuefamilycoherent;
    // shared variables are implicitly workgroupcoherent in GLSL.
    flags.workgroupcoherent = type.getQualifier().workgroupcoherent ||
                              type.getQualifier().storage == glslang::EvqShared;
    flags.subgroupcoherent = type.getQualifier().subgroupcoherent;
    flags.shadercallcoherent = type.getQualifier().shadercallcoherent;
    flags.volatil = type.getQualifier().volatil;
    flags.nontemporal = type.getQualifier().nontemporal;
```

Source: `thirdparty/glslang/SPIRV/GlslangToSpv.cpp` — TGlslangToSpvTraverser (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

spv::Module, spv::Function, spv::Block

Evidence:
- Code: `thirdparty/glslang/SPIRV/GlslangToSpv.cpp:124` — TGlslangToSpvTraverser
- Code: `thirdparty/glslang/SPIRV/SpvBuilder.cpp` — spv::Builder
- Code: `thirdparty/glslang/SPIRV/spvIR.h` — spv::Module
- External (official, verified): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Shader programs and material parameters](#topic-shader-programs)
- [Shader interface metadata](#topic-shader-interface-metadata)
- [Shader reflection](#topic-shader-reflection)
- [SPIR-V intermediate representation](#topic-spirv-intermediate-representation)
- [SPIR-V shader reflection](#topic-spirv-shader-reflection)
- [SPIR-V rewriting and optimization](#topic-spirv-rewriting)

<a id="topic-spirv-shader-reflection"></a>
## 303. SPIR-V shader reflection

**Scope:** Vendored: spirv-reflect

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Shader interface metadata](#topic-shader-interface-metadata).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The C header also provides a C++ ShaderModule wrapper when compiled as C++.

The reflection API creates a shader-module metadata view from SPIR-V code and exposes entry points, descriptors, interfaces, push constants, and specialization constants.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/*! @struct SpvReflectShaderModule

*/
typedef struct SpvReflectShaderModule {
  SpvReflectGenerator               generator;
  const char*                       entry_point_name;
  uint32_t                          entry_point_id;
  uint32_t                          entry_point_count;
  SpvReflectEntryPoint*             entry_points;
  SpvSourceLanguage                 source_language;
  uint32_t                          source_language_version;
  const char*                       source_file;
  const char*                       source_source;
```

Source: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectShaderModule (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SpvReflectShaderModule, SpvReflectEntryPoint

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:600` — SpvReflectShaderModule
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectCreateShaderModule, spvReflectDestroyShaderModule
- External (official, unverified (source anchor not found)): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Runtime class metadata](#topic-runtime-metadata)
- [Shader programs and material parameters](#topic-shader-programs)
- [Shader reflection](#topic-shader-reflection)
- [SPIR-V generation](#topic-spirv-generation)
- [SPIR-V intermediate representation](#topic-spirv-intermediate-representation)
- [SPIR-V rewriting and optimization](#topic-spirv-rewriting)

<a id="topic-shader-reflection"></a>
## 304. Shader reflection

**Scope:** Vendored: spirv-reflect

**Builds on:** [SPIR-V intermediate representation](#topic-spirv-intermediate-representation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

It is implemented independently from SPIRV-Cross under a C API.

The reflection API reads SPIR-V IR and returns module, entry-point, descriptor-binding, interface-variable, and push-constant metadata.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
const SpvReflectDescriptorBinding* spvReflectGetDescriptorBinding(const SpvReflectShaderModule* p_module, uint32_t binding_number,
                                                                  uint32_t set_number, SpvReflectResult* p_result) {
  const SpvReflectDescriptorBinding* p_descriptor = NULL;
  if (IsNotNull(p_module)) {
    for (uint32_t index = 0; index < p_module->descriptor_binding_count; ++index) {
      const SpvReflectDescriptorBinding* p_potential = &p_module->descriptor_bindings[index];
      if ((p_potential->binding == binding_number) && (p_potential->set == set_number)) {
        p_descriptor = p_potential;
        break;
      }
    }
  }
  if (IsNotNull(p_result)) {
```

Source: `thirdparty/spirv-reflect/spirv_reflect.c` — spvReflectGetDescriptorBinding (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SpvReflectShaderModule, SpvReflectDescriptorBinding

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:600` — SpvReflectShaderModule
- Code: `thirdparty/spirv-reflect/spirv_reflect.c:4976` — spvReflectGetDescriptorBinding
- Code: `thirdparty/spirv-reflect/spirv_reflect.c:5356` — spvReflectGetEntryPointPushConstantBlock
- External (official, unverified (source anchor not found)): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Runtime class metadata](#topic-runtime-metadata)
- [Shader interface metadata](#topic-shader-interface-metadata)
- [Shader programs and material parameters](#topic-shader-programs)
- [Shader interface mapping and reflection](#topic-shader-interface-analysis)
- [SPIR-V shader reflection](#topic-spirv-shader-reflection)
- [SPIR-V generation](#topic-spirv-generation)
- [SPIR-V rewriting and optimization](#topic-spirv-rewriting)

<a id="topic-skeletal-ik"></a>
## 305. Skeletal modifiers and inverse kinematics

**Scope:** First-party

**Builds on:** [Resources and asset lifecycle](#topic-resources).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The repository supplies general IK bases plus FABRIK and Jacobian solver variants.

IK skeleton modifiers process bone chains, targets, and joint-limitation resources to modify skeletal poses.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="IKModifier3D" inherits="SkeletonModifier3D" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		A node for inverse kinematics which may modify more than one bone.
	</brief_description>
	<description>
		Base class of [SkeletonModifier3D]s that has some joint lists and applies inverse kinematics. This class has some structs, enums, and helper methods which are useful to solve inverse kinematics.
	</description>
	<tutorials>
		<link title="Inverse Kinematics Returns to Godot 4.6 - IKModifier3D">https://godotengine.org/article/inverse-kinematics-returns-to-godot-4-6/#ikmodifier3d-and-7-child-classes</link>
	</tutorials>
	<methods>
		<method name="clear_settings">
			<return type="void" />
```

Source: `doc/classes/IKModifier3D.xml` — IKModifier3D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Resource

Evidence:
- Code: `doc/classes/IKModifier3D.xml:2` — IKModifier3D
- Code: `doc/classes/IterateIK3D.xml:2` — IterateIK3D
- Code: `doc/classes/JointLimitation3D.xml:2` — JointLimitation3D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resource-backed assets](#topic-resource-assets)
- [Skeletons, animation, and ragdolls](#topic-skeletal-ragdoll)

<a id="topic-spatial-predictive-filters"></a>
## 306. Spatial predictive filters

**Scope:** Vendored: libwebp

**Builds on:** [Picture planes and pixel ownership](#topic-input-picture-planes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Scalar, SSE2, NEON, MSA, and MIPS implementations cover the same filtering roles.

Spatial filters produce residuals from neighboring pixel-plane values using horizontal, vertical, gradient, and other predictor forms.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
static WEBP_INLINE int GradientPredictor_SSE2(uint8_t a, uint8_t b, uint8_t c) {
  const int g = a + b - c;
  return ((g & ~0xff) == 0) ? g : (g < 0) ? 0 : 255;  // clip to 8bit
}

static void GradientPredictDirect_SSE2(const uint8_t* const row,
                                       const uint8_t* const top,
                                       uint8_t* WEBP_RESTRICT const out,
                                       int length) {
  const int max_pos = length & ~7;
  int i;
  const __m128i zero = _mm_setzero_si128();
  for (i = 0; i < max_pos; i += 8) {
```

Source: `thirdparty/libwebp/src/dsp/filters_sse2.c` — GradientPredictor_SSE2 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/libwebp/src/dsp/filters.c` — GradientPredictor_C, WebPFiltersInit
- Code: `thirdparty/libwebp/src/utils/filters_utils.c:32` — WebPEstimateBestFilter
- Code: `thirdparty/libwebp/src/dsp/filters_sse2.c:129` — GradientPredictor_SSE2

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision filtering](#topic-collision-filtering)
- [Intersection and occlusion filters](#topic-filter-callbacks)

<a id="topic-subdivision-surface-evaluation"></a>
## 307. Subdivision surface evaluation

**Scope:** Vendored: embree

**Builds on:** [Parametric curve bases](#topic-curve-and-patch-bases).

**This prepares you for:** [Feature-adaptive tessellation](#topic-feature-adaptive-tessellation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Subdivision evaluation combines curve and patch bases with grid sampling; curve and patch bases are needed to explain the evaluator.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct EvalPatch 
    {
      /* creates EvalPatch from a CatmullClarkPatch */
      template<typename Allocator>
      __noinline static Ref create(const Allocator& alloc, const CatmullClarkPatch& patch) 
      {
        size_t ofs = 0, bytes = patch.bytes();
        void* ptr = alloc(bytes);
        patch.serialize(ptr,ofs);
        assert(ofs == bytes);
        return Ref(EVAL_PATCH, ptr);
      }
    };
```

Source: `thirdparty/embree/kernels/subdiv/patch.h` — EvalPatch (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SubGrid, GridMesh

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/patch.h:93` — EvalPatch
- Code: `thirdparty/embree/kernels/subdiv/patch_eval.h:14` — PatchEval
- Code: `thirdparty/embree/kernels/subdiv/patch_eval_grid.h:13` — PatchEvalGrid

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Half-edge topology](#topic-half-edge-topology)

<a id="topic-tls-connection-state"></a>
## 308. TLS connection and session state

**Scope:** Vendored: mbedtls

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [X.509 certificate processing](#topic-x509-certificate-processing).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Client, TLS 1.2, TLS 1.3, message, and common TLS units implement different portions of the protocol path.

The TLS state machine performs cryptographic operations while carrying configuration, negotiated session, handshake, record, and protocol state in mbedtls_ssl_context.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
MBEDTLS_CHECK_RETURN_CRITICAL
int mbedtls_ssl_write_client_hello(mbedtls_ssl_context *ssl);

#endif /* MBEDTLS_SSL_CLIENT_H */
```

Source: `thirdparty/mbedtls/library/ssl_client.h` — mbedtls_ssl_write_client_hello (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

mbedtls_ssl_session

Evidence:
- Code: `thirdparty/mbedtls/include/mbedtls/ssl.h` — struct mbedtls_ssl_session; struct mbedtls_ssl_config; struct mbedtls_ssl_context
- Code: `thirdparty/mbedtls/library/ssl_tls.c` — mbedtls_ssl_session_save; mbedtls_ssl_session_load
- Code: `thirdparty/mbedtls/library/ssl_client.h:16` — mbedtls_ssl_write_client_hello
- External (official, verified): [Mbed TLS documentation hub](https://mbed-tls.readthedocs.io/en/latest/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [Cryptographic resources and contexts](#topic-crypto-resources)
- [Networking and transport I/O](#topic-network-io)
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)
- [Byte streams and socket servers](#topic-stream-networking)
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-break-iteration"></a>
## 309. Text break iteration

**Scope:** Vendored: icu4c

**Builds on:** [Unicode data and properties](#topic-unicode-data).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The partition includes parsing, table building, runtime iteration, and dictionary or LSTM break-engine code.

ICU builds and executes rule-based break iterators, including dictionary-backed language break engines and compiled rule tables.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
*/
class LanguageBreakEngine : public UObject {
 public:

  /**
   * <p>Default constructor.</p>
   *
   */
  LanguageBreakEngine();

  /**
   * <p>Virtual destructor.</p>
   */
  virtual ~LanguageBreakEngine();
```

Source: `thirdparty/icu4c/common/brkeng.h` — LanguageBreakEngine (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RuleBasedBreakIterator, LanguageBreakEngine

Evidence:
- Code: `thirdparty/icu4c/common/unicode/rbbi.h:55` — RuleBasedBreakIterator
- Code: `thirdparty/icu4c/common/brkeng.h:28` — LanguageBreakEngine

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Dynamic arrays and dictionaries](#topic-variant-containers)

<a id="topic-theora-block-video-codec"></a>
## 310. Theora block video coding

**Scope:** Vendored: libtheora

**Builds on:** [Ogg pages, packets, and bit packing](#topic-ogg-pages-and-packets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Encoding and decoding are implemented separately but share transform, quantization, fragment, and state machinery.

The Theora codec consumes Ogg packets while unpacking quantization parameters, DCT token data, motion-compensated frame state, and optional accelerated transform routines.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
int th_decode_packetin(th_dec_ctx *_dec,const ogg_packet *_op,
 ogg_int64_t *_granpos){
  int ret;
  if(_dec==NULL||_op==NULL)return TH_EFAULT;
  /*A completely empty packet indicates a dropped frame and is treated exactly
     like an inter frame with no coded blocks.*/
  if(_op->bytes==0){
    _dec->state.frame_type=OC_INTER_FRAME;
    _dec->state.ntotal_coded_fragis=0;
  }
  else{
    oc_pack_readinit(&_dec->opb,_op->packet,_op->bytes);
    ret=oc_dec_frame_header_unpack(_dec);
```

Source: `thirdparty/libtheora/decode.c` — th_decode_packetin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Theora Stream Information, Ogg Packet

Evidence:
- Code: `thirdparty/libtheora/dequant.c:24` — oc_quant_params_unpack
- Code: `thirdparty/libtheora/decode.c:2740` — th_decode_packetin
- Code: `thirdparty/libtheora/mcenc.c:517` — oc_mcenc_search
- External (primary, unverified (source anchor not found)): [Theora Specification](https://www.theora.org/doc/Theora.pdf), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Instancing](#topic-instancing)
- [Motion blur](#topic-motion-blur)
- [Ogg packet sequences](#topic-ogg-packet-sequences)
- [Theora video streams](#topic-theora-video-streams)
- [Transform, quantization, and rate-distortion search](#topic-transform-quantization-rate-distortion)

<a id="topic-triangle-intersection-algorithms"></a>
## 311. Triangle intersection algorithms

**Scope:** Vendored: embree

**Builds on:** [Ray–primitive intersection](#topic-ray-primitive-intersection).

**This prepares you for:** [Motion-blur geometry](#topic-motion-blur-geometry), [SIMD ray packets](#topic-simd-ray-packets), [Subgrid intersection](#topic-subgrid-intersection).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Triangle intersection uses Möller–Trumbore, Plücker, and Woop kernels; ray-primitive intersection is required to interpret their hit candidates.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<int M, bool early_out = true>
    struct MoellerTrumboreIntersector1
    {
      __forceinline MoellerTrumboreIntersector1() {}

      __forceinline MoellerTrumboreIntersector1(const Ray& ray, const void* ptr) {}

      template<typename UVMapper>
      __forceinline bool intersect(const vbool<M>& valid0,
                                   Ray& ray,
                                   const Vec3vf<M>& tri_v0,
                                   const Vec3vf<M>& tri_e1,
                                   const Vec3vf<M>& tri_e2,
                                   const Vec3vf<M>& tri_Ng,
```

Source: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h` — MoellerTrumboreIntersector1 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h:62` — MoellerTrumboreIntersector1
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_pluecker.h:62` — PlueckerIntersector1
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_woop.h:77` — WoopIntersector1

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Geometry resources](#topic-geometry-resources)
- [Intersection results](#topic-hit-results)
- [Primitive intersection](#topic-primitive-intersection)
- [Ray query state](#topic-ray-query)

<a id="topic-upnp-port-mapping"></a>
## 312. UPnP port-mapping control

**Scope:** Vendored: miniupnpc

**Builds on:** [UPnP device discovery](#topic-upnp-device-discovery).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SOAP command helpers add, delete, query, and list port mappings on a discovered device.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
void
ParsePortListing(const char * buffer, int bufsize,
                 struct PortMappingParserData * pdata)
{
	struct xmlparser parser;

	memset(pdata, 0, sizeof(struct PortMappingParserData));
	/* init xmlparser */
	parser.xmlstart = buffer;
	parser.xmlsize = bufsize;
	parser.data = pdata;
	parser.starteltfunc = startelt;
	parser.endeltfunc = endelt;
	parser.datafunc = data;
```

Source: `thirdparty/miniupnpc/src/portlistingparse.c` — ParsePortListing (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PortMapping, PortMappingParserData

Evidence:
- Code: `thirdparty/miniupnpc/include/miniupnpc/upnpcommands.h:201` — UPNP_AddPortMapping
- Code: `thirdparty/miniupnpc/src/portlistingparse.c:149` — ParsePortListing

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [UPnP device discovery](#topic-upnp-device-discovery)

<a id="topic-unicode-normalization"></a>
## 313. Unicode normalization

**Scope:** Vendored: icu4c

**Builds on:** [Unicode data and properties](#topic-unicode-data).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code contains NFC, NFKC, FCD, and related Normalizer2 factory paths.

ICU implements Normalizer2 variants and loads normalization data into tries, extra mappings, and composition data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Normalizer2Impl;

class U_COMMON_API ReorderingBuffer : public UMemory {
public:
    /** Constructs only; init() should be called. */
    ReorderingBuffer(const Normalizer2Impl &ni, UnicodeString &dest) :
        impl(ni), str(dest),
        start(nullptr), reorderStart(nullptr), limit(nullptr),
        remainingCapacity(0), lastCC(0) {}
    /** Constructs, removes the string contents, and initializes for a small initial capacity. */
    ReorderingBuffer(const Normalizer2Impl &ni, UnicodeString &dest, UErrorCode &errorCode);
    ~ReorderingBuffer() {
        if (start != nullptr) {
```

Source: `thirdparty/icu4c/common/normalizer2impl.h` — Normalizer2Impl (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Normalizer2Impl, UCPTrie

Evidence:
- Code: `thirdparty/icu4c/common/normalizer2impl.h:137` — Normalizer2Impl
- Code: `thirdparty/icu4c/common/normalizer2.cpp` — Normalizer2Factory::getNFCInstance

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Unicode data and properties](#topic-unicode-data)

<a id="topic-variable-font-subsetting"></a>
## 314. Variable-font table subsetting

**Scope:** Vendored: harfbuzz

**Builds on:** [Font subsetting](#topic-font-subsetting).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Variation handling is controlled by table tags and user-axis locations in the subset plan.

The font subsetting pipeline dispatches HVAR, VVAR, gvar, fvar, avar, cvar, and mvar tables, with fvar and avar passed through when user axis locations are empty.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
bool _hb_subset_table_var		(hb_subset_plan_t *plan, hb_vector_t<char> &buf, hb_tag_t tag, bool *success)
{
#ifndef HB_NO_VAR
  switch (tag)
  {
  case HB_TAG('H','V','A','R'): *success = _hb_subset_table<const OT::HVAR> (plan, buf); return true;
  case HB_TAG('V','V','A','R'): *success = _hb_subset_table<const OT::VVAR> (plan, buf); return true;
  case HB_TAG('g','v','a','r'): *success = _hb_subset_table<const OT::gvar> (plan, buf); return true;
  case HB_TAG('f','v','a','r'):
    if (plan->user_axes_location.is_empty ())
      *success = _hb_subset_table_passthrough (plan, tag);
    else
      *success = _hb_subset_table<const OT::fvar> (plan, buf);
```

Source: `thirdparty/harfbuzz/src/hb-subset-table-var.cc` — _hb_subset_table_var (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-table-var.cc:10` — _hb_subset_table_var
- Code: `thirdparty/harfbuzz/src/hb-subset-plan-var.cc` — hb_subset_plan_var

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [TrueType GX variation data](#topic-variable-font-data)

<a id="topic-vector-font-export"></a>
## 315. Vector font export

**Scope:** Vendored: harfbuzz

**Builds on:** [Color-font paint processing](#topic-color-font-paint).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The vector path, draw, paint, PDF, and SVG files form a separate output family from raster paint.

Color paint records are emitted through vector drawing and paint backends that include PDF and SVG implementations.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/* Callback context + trampoline for hb_paint_sweep_gradient_tiles. */
struct hb_vector_svg_sweep_ctx_t
{
  hb_vector_buf_t *body;
  unsigned precision;
  float cx, cy, radius;
};

static void
hb_vector_svg_sweep_emit_patch (float a0, hb_color_t c0,
				float a1, hb_color_t c1,
				void *user_data)
{
  auto *ctx = (hb_vector_svg_sweep_ctx_t *) user_data;
```

Source: `thirdparty/harfbuzz/src/hb-vector-paint-svg.cc` — hb_vector_svg_sweep_ctx_t (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

hb_vector_paint_t, hb_vector_draw_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-vector-paint.hh:41` — hb_vector_paint_t
- Code: `thirdparty/harfbuzz/src/hb-vector-paint-svg.cc:247` — hb_vector_svg_sweep_ctx_t

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Compile-time platform backends](#topic-sdl-platform-backends)
- [Generic container infrastructure](#topic-generic-containers)
- [Multi-channel distance fields](#topic-multichannel-distance-fields)
- [Visual shader vector operations](#topic-visual-shader-vector-operations)

<a id="topic-extensions"></a>
## 316. Virtual implementation extensions

**Scope:** First-party

**Builds on:** [Runtime class metadata](#topic-runtime-metadata).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The contracts are exposed as subclass APIs such as PhysicsServer2DExtension and TextServerExtension.

Extension classes declare required or optional virtual callbacks so external implementations can replace physics, rendering, text, and scripting behavior.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="PhysicsServer2DExtension" inherits="PhysicsServer2D" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Provides virtual methods that can be overridden to create custom [PhysicsServer2D] implementations.
	</brief_description>
	<description>
		This class extends [PhysicsServer2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
		Intended for use with GDExtension to create custom implementations of [PhysicsServer2D].
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="_area_add_shape" qualifiers="virtual required">
			<return type="void" />
```

Source: `doc/classes/PhysicsServer2DExtension.xml` — PhysicsServer2DExtension (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RenderSceneBuffersConfiguration, RID

Evidence:
- Code: `doc/classes/PhysicsServer2DExtension.xml:2` — PhysicsServer2DExtension
- Code: `doc/classes/TextServerExtension.xml:2` — TextServerExtension
- Code: `doc/classes/ScriptLanguageExtension.xml:2` — ScriptLanguageExtension

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Extension API compatibility baselines](#topic-extension-api-compatibility)
- [Vulkan extension structure chains](#topic-extension-structure-chains)
- [GDExtension libraries](#topic-gdextension-libraries)
- [Native extension loading](#topic-native-extensions)
- [OpenXR composition layers](#topic-openxr-composition-layers)
- [OpenXR extension wrappers](#topic-openxr-extension-wrappers)
- [OpenXR render models](#topic-openxr-render-models)
- [Pluggable and wrapped server backends](#topic-pluggable-server-backends)
- [Resource formats and serialization](#topic-resource-formats)
- [Script languages and instances](#topic-script-hosting)
- [Vulkan instance setup](#topic-vulkan-instance-setup)

<a id="topic-vorbis-block-synthesis"></a>
## 317. Vorbis block analysis and synthesis

**Scope:** Vendored: libvorbis

**Builds on:** [Ogg pages, packets, and bit packing](#topic-ogg-pages-and-packets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation has separate analysis and synthesis entry points over the same block-oriented codec model.

Vorbis analysis and synthesis operate on codec blocks and Ogg packets, using mapping, floor, residue, codebook, window, and MDCT modules.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* decides between modes, dispatches to the appropriate mapping. */
int vorbis_analysis(vorbis_block *vb, ogg_packet *op){
  int ret,i;
  vorbis_block_internal *vbi=vb->internal;

  vb->glue_bits=0;
  vb->time_bits=0;
  vb->floor_bits=0;
  vb->res_bits=0;

  /* first things first.  Make sure encode is ready */
  for(i=0;i<PACKETBLOBS;i++)
    oggpack_reset(vbi->packetblob[i]);
```

Source: `thirdparty/libvorbis/analysis.c` — vorbis_analysis (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Vorbis Block, Ogg Packet

Evidence:
- Code: `thirdparty/libvorbis/analysis.c:29` — vorbis_analysis
- Code: `thirdparty/libvorbis/synthesis.c:25` — vorbis_synthesis
- Code: `thirdparty/libvorbis/mdct.c:396` — mdct_backward
- External (primary, unverified (source anchor not found)): [Vorbis I Specification](https://xiph.org/vorbis/doc/Vorbis_I_spec.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ogg packet sequences](#topic-ogg-packet-sequences)

<a id="topic-webp-riff-decoding"></a>
## 318. WebP chunk parsing, incremental decode, and animation

**Scope:** Vendored: libwebp

**Builds on:** [Image decode pipelines](#topic-image-decode-pipeline).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The decoder, demuxer, and animation decoder are distinct modules over shared WebP byte data.

The WebP code implements an image decode pipeline that recognizes VP8 and VP8L payloads, incrementally retains input, demuxes RIFF chunks into frames, and composites animation frames.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct WebPIDecoder {
  DecState state;         // current decoding state
  WebPDecParams params;   // Params to store output info
  int is_lossless;        // for down-casting 'dec'.
  void* dec;              // either a VP8Decoder or a VP8LDecoder instance
  VP8Io io;

  MemBuffer mem;          // input memory buffer.
  WebPDecBuffer output;   // output buffer (when no external one is supplied,
                          // or if the external one has slow-memory)
  WebPDecBuffer* final_output;  // Slow-memory output to copy to eventually.
  size_t chunk_size;      // Compressed VP8/VP8L size extracted from Header.
```

Source: `thirdparty/libwebp/src/dec/idec_dec.c` — WebPIDecoder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebP Decoder State

Evidence:
- Code: `thirdparty/libwebp/src/dec/webp_dec.c:154` — ParseOptionalChunks
- Code: `thirdparty/libwebp/src/dec/idec_dec.c:72` — WebPIDecoder
- Code: `thirdparty/libwebp/src/demux/demux.c:64` — WebPDemuxer
- External (official, unverified (source anchor not found)): [WebP Container Specification](https://developers.google.com/speed/webp/docs/riff_container), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [Lossless pixel transform pipeline](#topic-lossless-transform-pipeline)
- [WebP image I/O](#topic-webp-image-io)
- [Image decode pipelines](#topic-image-decode-pipeline)

<a id="topic-yuv-rgb-conversion"></a>
## 319. YUV/RGB conversion and chroma processing

**Scope:** Vendored: libwebp

**Builds on:** [Picture planes and pixel ownership](#topic-input-picture-planes).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Format packers also target RGBA, BGRA, RGB565, and RGBA4444 outputs.

Conversion kernels turn RGB or ARGB input into YUV planes, reconstruct RGB outputs, and upsample chroma with neighboring samples.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
int WebPPictureARGBToYUVADithered(WebPPicture* picture, WebPEncCSP colorspace,
                                  float dithering) {
  return PictureARGBToYUVA(picture, colorspace, dithering, 0);
}

int WebPPictureARGBToYUVA(WebPPicture* picture, WebPEncCSP colorspace) {
  return PictureARGBToYUVA(picture, colorspace, 0.f, 0);
}

int WebPPictureSharpARGBToYUVA(WebPPicture* picture) {
  return PictureARGBToYUVA(picture, WEBP_YUV420, 0.f, 1);
}
// for backward compatibility
```

Source: `thirdparty/libwebp/src/enc/picture_csp_enc.c` — WebPPictureARGBToYUVA (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebPPicture, WebPDecBuffer

Evidence:
- Code: `thirdparty/libwebp/src/dsp/yuv.h` — VP8YUVToR, VP8RGBToY
- Code: `thirdparty/libwebp/src/dsp/yuv.c` — WebPConvertARGBToYUV
- Code: `thirdparty/libwebp/src/dsp/upsampling.c` — WebPUpsamplersInit
- Code: `thirdparty/libwebp/src/enc/picture_csp_enc.c:657` — WebPPictureARGBToYUVA

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [Picture planes and pixel ownership](#topic-input-picture-planes)

<a id="topic-deferred-execution"></a>
## 320. Deferred calls and worker tasks

**Scope:** First-party

**Builds on:** [Dynamic invocation and signals](#topic-dynamic-invocation-signals).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

MessageQueue records Object/Callable work for later execution, while WorkerThreadPool represents task and group work for worker threads.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CallQueue {
	friend class MessageQueue;

public:
	enum {
		PAGE_SIZE_BYTES = 4096
	};

	struct Page {
		uint8_t data[PAGE_SIZE_BYTES];
	};

	// Needs to be public to be able to define it outside the class.
	// Needs to lock because there can be multiple of these allocators in several threads.
```

Source: `core/object/message_queue.h` — MessageQueue (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Object, Variant

Evidence:
- Code: `core/object/message_queue.h:42` — MessageQueue
- Code: `core/object/message_queue.h` — MessageQueue::Message
- Code: `core/object/worker_thread_pool.h` — WorkerThreadPool::Task
- Code: `core/templates/command_queue_mt.h:41` — CommandQueueMT

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)
- [Thread and synchronization abstractions](#topic-sdl-thread-abstractions)
- [Task scheduling and synchronization](#topic-task-scheduling)

<a id="topic-generic-containers"></a>
## 321. Generic container infrastructure

**Scope:** First-party

**Builds on:** [Dynamic arrays and dictionaries](#topic-variant-containers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The template layer supplies the C++ storage and lookup structures that underpin Array and Dictionary, including CowData, Vector, maps, sets, lists, spans, and RID owners.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template <typename T>
class CowData {
public:
	typedef int64_t Size;
	typedef uint64_t USize;
	static constexpr USize MAX_INT = INT64_MAX;

private:
	// Alignment:  ↓ max_align_t           ↓ USize          ↓ USize            ↓ MAX_ALIGN
	//             ┌────────────────────┬──┬───────────────┬──┬─────────────┬──┬───────────...
	//             │ SafeNumeric<USize> │░░│ USize         │░░│ USize       │░░│ T[]
	//             │ ref. count         │░░│ data capacity │░░│ data size   │░░│ data
	//             └────────────────────┴──┴───────────────┴──┴─────────────┴──┴───────────...
	// Offset:     ↑ REF_COUNT_OFFSET      ↑ CAPACITY_OFFSET  ↑ SIZE_OFFSET    ↑ DATA_OFFSET
```

Source: `core/templates/cowdata.h` — CowData (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Array, Dictionary

Evidence:
- Code: `core/templates/cowdata.h:60` — CowData
- Code: `core/templates/vector.h:34` — Vector
- Code: `core/templates/hash_map.h:42` — HashMap
- Code: `core/templates/rid_owner.h:49` — RID_Owner

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Resource and server identifiers](#topic-resource-identifiers)
- [Vector font export](#topic-vector-font-export)

<a id="topic-input-actions"></a>
## 322. Input events and actions

**Scope:** First-party

**Builds on:** [Scene graph and lifecycle](#topic-scene-tree).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The model covers physical keyboard, mouse, touch, joystick, gesture, and manually generated action events.

InputMap associates named actions with InputEvent instances, and nodes receive input-event callbacks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="InputEvent" inherits="Resource" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Abstract base class for input events.
	</brief_description>
	<description>
		Abstract base class of all types of input events. See [method Node._input].
	</description>
	<tutorials>
		<link title="Using InputEvent">$DOCS_URL/tutorials/inputs/inputevent.html</link>
		<link title="Viewport and canvas transforms">$DOCS_URL/tutorials/2d/2d_transforms.html</link>
		<link title="2D Dodge The Creeps Demo">https://godotengine.org/asset-library/asset/2712</link>
		<link title="3D Voxel Demo">https://godotengine.org/asset-library/asset/2755</link>
	</tutorials>
```

Source: `doc/classes/InputEvent.xml` — InputEvent (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

InputEvent

Evidence:
- Code: `core/input/input.h` — Input::parse_input_event
- Code: `doc/classes/InputEvent.xml:2` — InputEvent
- Code: `doc/classes/InputMap.xml` — InputMap::action_add_event

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Input events and actions](#topic-input-events-actions)
- [Platform display and window adaptation](#topic-platform-display-adaptation)

<a id="topic-native-extensions"></a>
## 323. Native extension loading

**Scope:** First-party

**Builds on:** [Class metadata and reflection](#topic-reflection), [Resources and asset lifecycle](#topic-resources).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This lets native libraries add engine-facing types without compiling them into the core binary.

A GDExtension is a Resource loaded through a loader; its registered classes become ClassDB-visible extension classes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GDExtensionMethodBind;

class GDExtension : public Resource {
	GDCLASS(GDExtension, Resource)

	friend class GDExtensionManager;

	Ref<GDExtensionLoader> loader;

	bool reloadable = false;

	struct Extension {
		ObjectGDExtension gdextension;
```

Source: `core/extension/gdextension.h` — GDExtension (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDExtension, ClassInfo

Evidence:
- Code: `core/extension/gdextension.h:39` — GDExtension
- Code: `core/object/class_db.h` — ClassDB::register_extension_class
- Code: `doc/classes/GDExtensionManager.xml:2` — GDExtensionManager

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Extension API compatibility baselines](#topic-extension-api-compatibility)
- [Virtual implementation extensions](#topic-extensions)
- [GDExtension libraries](#topic-gdextension-libraries)
- [Reflection metadata](#topic-reflection-metadata)
- [Runtime class metadata](#topic-runtime-metadata)

<a id="topic-undo-redo"></a>
## 324. Undo and redo actions

**Scope:** First-party

**Builds on:** [Dynamic invocation and signals](#topic-dynamic-invocation-signals).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

UndoRedo records Object/Callable operations as actions so they can be replayed in undo or redo direction.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class UndoRedo : public Object {
	GDCLASS(UndoRedo, Object);
	OBJ_SAVE_TYPE(UndoRedo);

public:
	enum MergeMode {
		MERGE_DISABLE,
		MERGE_ENDS,
		MERGE_ALL
	};

	typedef void (*CommitNotifyCallback)(void *p_ud, const String &p_name);
```

Source: `core/object/undo_redo.h` — UndoRedo (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Object, Variant

Evidence:
- Code: `core/object/undo_redo.h:36` — UndoRedo
- Code: `core/object/undo_redo.h` — UndoRedo::Operation
- Code: `core/object/undo_redo.h` — UndoRedo::Action

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)

<a id="topic-two-dimensional-content"></a>
## 325. 2D shapes, tiles, and paths

**Scope:** First-party

**Builds on:** [Resource-backed assets](#topic-resource-assets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The 2D resource partition is conditionally compiled for physics and navigation features.

Resource-owned Shape2D subclasses, TileSet data, NavigationPolygon data, and PolygonPathFinder provide 2D geometry and tile-oriented content.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class NavigationPolygon : public Resource {
	GDCLASS(NavigationPolygon, Resource);
	RWLock rwlock;

	Vector<Vector2> vertices;
	Vector<Vector<int>> polygons;
	Vector<Vector<Vector2>> outlines;

	mutable Rect2 item_rect;
	mutable bool rect_cache_dirty = true;

	Mutex navigation_mesh_generation;
	// Navigation mesh
```

Source: `scene/resources/2d/navigation_polygon.h` — class NavigationPolygon : public Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TileMapPattern

Evidence:
- Code: `scene/resources/2d/shape_2d.h:35` — class Shape2D : public Resource
- Code: `scene/resources/2d/tile_set.h` — class TileMapPattern and class TileSet
- Code: `scene/resources/2d/navigation_polygon.h:36` — class NavigationPolygon : public Resource
- Code: `scene/resources/2d/polygon_path_finder.h:35` — class PolygonPathFinder : public Resource

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry resources](#topic-geometry-resources)
- [Physics shapes, objects, and collision reports](#topic-physics-collision)
- [Resources and asset lifecycle](#topic-resources)
- [Tile libraries, cells, and patterns](#topic-tile-libraries)

<a id="topic-skeleton-modification"></a>
## 326. 2D skeleton modification stacks

**Scope:** First-party

**Builds on:** [Resource-backed assets](#topic-resource-assets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied stack includes CCDIK, FABRIK, LookAt, TwoBoneIK, jiggle, and physical-bone variants.

SkeletonModification2D resources define individual 2D skeleton operations, and SkeletonModificationStack2D holds modifications for a Skeleton2D.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class SkeletonModification2D : public Resource {
	GDCLASS(SkeletonModification2D, Resource);
	friend class Skeleton2D;
	friend class Bone2D;

protected:
	static void _bind_methods();

	SkeletonModificationStack2D *stack = nullptr;
	int execution_mode = 0; // 0 = process

	bool enabled = true;
	bool is_setup = false;
```

Source: `scene/resources/2d/skeleton/skeleton_modification_2d.h` — class SkeletonModification2D : public Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/resources/2d/skeleton/skeleton_modification_2d.h:41` — class SkeletonModification2D : public Resource
- Code: `scene/resources/2d/skeleton/skeleton_modification_stack_2d.h:43` — class SkeletonModificationStack2D : public Resource
- Code: `scene/resources/2d/skeleton/SCsub` — registered skeleton modification source files

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Skeletons, animation, and ragdolls](#topic-skeletal-ragdoll)
- [Skeleton modification and physical-bone simulation](#topic-skeleton-modifiers)

<a id="topic-three-dimensional-content"></a>
## 327. 3D shapes, worlds, and skins

**Scope:** First-party

**Builds on:** [Resource-backed assets](#topic-resource-assets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The 3D resource partition is selected only when 3D is enabled by the build configuration.

Resource-owned Shape3D subclasses, World3D, Skin, and mesh import data represent 3D collision, world, skeletal, and mesh content.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class ImporterMesh : public Resource {
	GDCLASS(ImporterMesh, Resource)

	struct Surface {
		Mesh::PrimitiveType primitive;
		Array arrays;
		struct BlendShape {
			Array arrays;
		};
		Vector<BlendShape> blend_shape_data;
		struct LOD {
			Vector<int> indices;
			float distance = 0.0f;
```

Source: `scene/resources/3d/importer_mesh.h` — class ImporterMesh : public Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/resources/3d/shape_3d.h:38` — class Shape3D : public Resource
- Code: `scene/resources/3d/world_3d.h:46` — class World3D : public Resource
- Code: `scene/resources/3d/skin.h` — class Skin : public Resource and Skin::Bind
- Code: `scene/resources/3d/importer_mesh.h:47` — class ImporterMesh : public Resource

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-audio-bus-processing"></a>
## 328. Audio buses, streams, and effects

**Scope:** First-party

**Builds on:** [Resource-backed assets](#topic-resource-assets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Effect classes create reference-counted processing instances.

AudioServer owns buses with channels and effects, while Resource-derived streams, effects, and bus layouts define reusable audio configuration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class AudioStream : public Resource {
	GDCLASS(AudioStream, Resource);
	OBJ_SAVE_TYPE(AudioStream); // Saves derived classes with common type so they can be interchanged.

	enum {
		MAX_TAGGED_OFFSETS = 8
	};

	uint64_t tagged_frame = 0;
	uint64_t offset_count = 0;
	float tagged_offsets[MAX_TAGGED_OFFSETS];

protected:
```

Source: `servers/audio/audio_stream.h` — class AudioStream : public Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

AudioBusLayout

Evidence:
- Code: `servers/audio/audio_server.h` — AudioServer::Bus, AudioServer::Channel, AudioServer::Effect, and AudioBusLayout
- Code: `servers/audio/audio_stream.h:157` — class AudioStream : public Resource
- Code: `servers/audio/audio_effect.h` — AudioEffectInstance : public RefCounted and AudioEffect : public Resource

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-frame-timing"></a>
## 329. Frame timing and physics stepping

**Scope:** First-party

**Builds on:** [Engine bootstrap](#topic-engine-bootstrap).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The timer synchronization layer separates frame-duration handling from main startup.

Frame timing calculates a process delta and bounded physics-step count after engine bootstrap determines the active physics tick configuration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
// advance physics clock by p_process_step, return appropriate number of steps to simulate
MainFrameTime MainTimerSync::advance_core(double p_physics_step, int p_physics_ticks_per_second, double p_process_step) {
	MainFrameTime ret;

	ret.process_step = p_process_step;

	// simple determination of number of physics iteration
	time_accum += ret.process_step;
	ret.physics_steps = std::floor(time_accum * p_physics_ticks_per_second);

	int min_typical_steps = typical_physics_steps[0];
	int max_typical_steps = min_typical_steps + 1;

	// given the past recorded steps and typical steps to match, calculate bounds for this
```

Source: `main/main_timer_sync.cpp` — MainTimerSync::advance (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MainFrameTime

Evidence:
- Code: `main/main_timer_sync.h:35` — MainFrameTime
- Code: `main/main_timer_sync.h:43` — MainTimerSync
- Code: `main/main_timer_sync.cpp:539` — MainTimerSync::advance

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)

<a id="topic-rendering-assets"></a>
## 330. Meshes, materials, textures, and instancing

**Scope:** First-party

**Builds on:** [Resources and asset lifecycle](#topic-resources), [Scene graph and lifecycle](#topic-scene-tree).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The resource layer holds drawable data while scene nodes select where and how it is rendered.

Mesh resources provide surfaces; geometry nodes instance them, Materials control shading, and MultiMesh batches many instances.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="MultiMesh" inherits="Resource" api_type="core" keywords="batch" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Provides high-performance drawing of a mesh multiple times using GPU instancing.
	</brief_description>
	<description>
		MultiMesh provides low-level mesh instancing. Drawing thousands of [MeshInstance3D] nodes can be slow, since each object is submitted to the GPU then drawn individually.
		MultiMesh is much faster as it can draw thousands of instances with a single draw call, resulting in less API overhead.
		As a drawback, if the instances are too far away from each other, performance may be reduced as every single instance will always render (they are spatially indexed as one, for the whole object).
		Since instances may have any behavior, the AABB used for visibility must be provided by the user.
		[b]Note:[/b] A MultiMesh is a single object, therefore the same maximum lights per object restriction applies. This means, that once the maximum lights are consumed by one or more instances, the rest of the MultiMesh instances will [b]not[/b] receive any lighting.
		[b]Note:[/b] Blend Shapes will be ignored if used in a MultiMesh.
	</description>
	<tutorials>
```

Source: `doc/classes/MultiMesh.xml` — MultiMesh (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Mesh, Material

Evidence:
- Code: `scene/resources/mesh.h` — Mesh::get_surface_count
- Code: `scene/resources/mesh.h` — Mesh::surface_set_material
- Code: `doc/classes/MultiMesh.xml:2` — MultiMesh

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry resources](#topic-geometry-resources)
- [GUI controls and layout](#topic-gui-control-layout)
- [Instancing](#topic-instancing)
- [COLLADA scene interchange](#topic-collada-import)
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Halfedge topology](#topic-halfedge-topology)
- [Navigation mesh construction](#topic-navigation-mesh-construction)
- [Post-import processing](#topic-post-import-processing)
- [Textures, meshes, and materials](#topic-rendering-resources)
- [Resource-specific editing](#topic-resource-specific-editors)
- [Spatial indexing and ray queries](#topic-spatial-indexing)
- [3D shapes, worlds, and skins](#topic-three-dimensional-content)
- [WebRTC multiplayer mesh](#topic-webrtc-multiplayer-mesh)
- [Resource-backed assets](#topic-resource-assets)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-packed-scenes"></a>
## 331. Packed scene serialization

**Scope:** First-party

**Builds on:** [Resources and asset lifecycle](#topic-resources), [Scene graph and lifecycle](#topic-scene-tree).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SceneState separates reusable name, Variant, path, node, and connection tables from the PackedScene resource wrapper.

A PackedScene stores a resource-backed SceneState whose node entries, property values, and connection entries reconstruct a node hierarchy.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
Array setup_resources_in_array(Array &array_to_scan, const SceneState::NodeData &n, HashMap<Node *, HashMap<Ref<Resource>, Ref<Resource>>> &p_resources_local_to_scenes, Node *node, const StringName sname, int i, Node **ret_nodes, SceneState::GenEditState p_edit_state) const;
	Dictionary setup_resources_in_dictionary(Dictionary &p_dictionary_to_scan, const SceneState::NodeData &p_n, HashMap<Node *, HashMap<Ref<Resource>, Ref<Resource>>> &p_resources_local_to_scenes, Node *p_node, const StringName p_sname, int p_i, Node **p_ret_nodes, SceneState::GenEditState p_edit_state) const;
	Variant make_local_resource(Variant &value, const SceneState::NodeData &p_node_data, HashMap<Node *, HashMap<Ref<Resource>, Ref<Resource>>> &p_resources_local_to_scenes, Node *p_node, const StringName p_sname, int p_i, Node **p_ret_nodes, SceneState::GenEditState p_edit_state) const;
	bool has_local_resource(const Array &p_array) const;

	Ref<SceneState> get_base_scene_state() const;

	void update_instance_resource(String p_path, Ref<PackedScene> p_packed_scene);

	//unbuild API

	int get_node_count() const;
	StringName get_node_type(int p_idx) const;
```

Source: `scene/resources/packed_scene.h` — SceneState::NodeData (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PackedScene, SceneState

Evidence:
- Code: `scene/resources/packed_scene.h:167` — SceneState::NodeData
- Code: `scene/resources/packed_scene.h` — SceneState::ConnectionData
- Code: `scene/resources/packed_scene.h:36` — PackedScene

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [BVH traversal](#topic-bvh-traversal)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Scene state inspection](#topic-scene-state)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-performance-monitors"></a>
## 332. Performance monitors

**Scope:** First-party

**Builds on:** [Engine bootstrap](#topic-engine-bootstrap).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The `Performance` object and `MonitorCall` class are in the main runtime partition.

Performance monitors collect engine counters and execute registered monitor calls after engine bootstrap.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/**************************************************************************/
/*  performance.cpp                                                       */
/**************************************************************************/
/*                         This file is part of:                          */
/*                             GODOT ENGINE                               */
/*                        https://godotengine.org                         */
/**************************************************************************/
/* Copyright (c) 2014-present Godot Engine contributors (see AUTHORS.md). */
/* Copyright (c) 2007-2014 Juan Linietsky, Ariel Manzur.                  */
/*                                                                        */
/* Permission is hereby granted, free of charge, to any person obtaining  */
/* a copy of this software and associated documentation files (the        */
/* "Software"), to deal in the Software without restriction, including    */
/* without limitation the rights to use, copy, modify, merge, publish,    */
```

Source: `main/performance.cpp` — Performance (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `main/performance.h:43` — Performance
- Code: `main/performance.h:158` — MonitorCall
- Code: `main/performance.cpp:59` — Performance

<a id="topic-text-interface"></a>
## 333. Text layout and editing

**Scope:** First-party

**Builds on:** [GUI controls and layout](#topic-gui-controls).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

TextEdit, LineEdit, Label, and RichTextLabel each consume shaped-text resources in this partition.

GUI text controls hold shaped line and paragraph data, use text-server glyph and selection queries, and track IME text and selection state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class TextParagraph : public RefCounted {
	GDCLASS(TextParagraph, RefCounted);
	_THREAD_SAFE_CLASS_

private:
	mutable RID dropcap_rid;
	mutable int dropcap_lines = 0;
	Rect2 dropcap_margins;

	RID rid;
	mutable LocalVector<RID> lines_rid;

	mutable bool lines_dirty = true;
```

Source: `scene/resources/text_paragraph.h` — class TextParagraph : public RefCounted (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/gui/text_edit.h` — TextEdit::Text and TextEdit::Line
- Code: `scene/gui/text_edit.cpp` — TextEdit IME and shaped_text_get_selection calls
- Code: `scene/gui/rich_text_label.h` — class RichTextLabel and RichTextLabel::Line
- Code: `scene/resources/text_paragraph.h:40` — class TextParagraph : public RefCounted

<a id="topic-rendering-resources"></a>
## 334. Textures, meshes, and materials

**Scope:** First-party

**Builds on:** [Resource-backed assets](#topic-resource-assets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The inspected source includes array meshes, primitive meshes, compressed and image textures, and several material families.

Resource-owned Mesh, Material, Texture, Environment, and Sky classes carry rendering-facing content and configuration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class Environment : public Resource {
	GDCLASS(Environment, Resource);

public:
	enum BGMode {
		BG_CLEAR_COLOR,
		BG_COLOR,
		BG_SKY,
		BG_CANVAS,
		BG_KEEP,
		BG_CAMERA_FEED,
		BG_MAX
	};
```

Source: `scene/resources/environment.h` — class Environment : public Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/resources/mesh.h` — class Mesh and class ArrayMesh
- Code: `scene/resources/material.h` — class Material, ShaderMaterial, and BaseMaterial3D
- Code: `scene/resources/texture.h` — Texture, Texture2D, TextureLayered, and Texture3D
- Code: `scene/resources/environment.h:38` — class Environment : public Resource

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor and project settings](#topic-editor-and-project-settings)
- [Encoder configuration](#topic-encoder-configuration)
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-themes-and-style-boxes"></a>
## 335. Themes and style boxes

**Scope:** First-party

**Builds on:** [Resource-backed assets](#topic-resource-assets).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Default-theme construction also appears in the inspected theme source.

Resource-backed Theme data is resolved through ThemeDB, ThemeContext, and ThemeOwner, while StyleBox subclasses supply control styling.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class StyleBox : public Resource {
	GDCLASS(StyleBox, Resource);
	RES_BASE_EXTENSION("stylebox");
	OBJ_SAVE_TYPE(StyleBox);

	float content_margin[4];

protected:
	static void _bind_methods();
	virtual float get_style_margin(Side p_side) const { return 0; }

	GDVIRTUAL2C_REQUIRED(_draw, RID, Rect2)
	GDVIRTUAL1RC(Rect2, _get_draw_rect, Rect2)
```

Source: `scene/resources/style_box.h` — class StyleBox : public Resource (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `scene/resources/theme.h:38` — class Theme : public Resource
- Code: `scene/resources/style_box.h:38` — class StyleBox : public Resource
- Code: `scene/theme/theme_db.h` — class ThemeDB and class ThemeContext
- Code: `scene/theme/default_theme.cpp` — default control colors and StyleBoxFlat construction

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor theming](#topic-editor-theming)
- [GUI controls and layout](#topic-gui-control-layout)
- [Resources and asset lifecycle](#topic-resources)
- [UI themes and style boxes](#topic-ui-theming)
- [Control-tree user interfaces](#topic-ui-composition)

<a id="topic-scene-contexts"></a>
## 336. Scene-aware test context

**Scope:** First-party

**Builds on:** [Contextual completion contracts](#topic-completion-contexts).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Completion cases select this context through a configuration `scene` path.

A scene input gives a completion context containing nodes, attached scripts, resources, and unique names.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
[node name="UniqueA" type="Node" parent="UniqueNames"]
unique_name_in_owner = true
script = ExtResource("1_ldc4g")

[node name="A" type="Node" parent="."]
script = ExtResource("1_ldc4g")
```

Source: `modules/gdscript/tests/scripts/completion/get_node/get_node.tscn` — [node name="UniqueA" type="Node" parent="UniqueNames"] (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Scene Fixture, Completion Test Configuration

Evidence:
- Code: `modules/gdscript/tests/scripts/completion/get_node/get_node.tscn:14` — [node name="UniqueA" type="Node" parent="UniqueNames"]
- Code: `modules/gdscript/tests/scripts/completion/get_node/literal_scene/dollar_class_scene.cfg` — [input] scene

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resource-backed assets](#topic-resource-assets)
- [Scene construction and commit](#topic-scene-commit)
- [COLLADA scene interchange](#topic-collada-import)
- [Editing selection history](#topic-editing-selection-history)
- [Editor authoring workspaces](#topic-editor-authoring-workspaces)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Editor plugin state and custom types](#topic-editor-plugin-state)
- [Editor scene sessions](#topic-editor-scene-sessions)
- [Engine bootstrap](#topic-engine-bootstrap)
- [GUI controls and layout](#topic-gui-control-layout)
- [Instancing](#topic-instancing)
- [OpenXR composition layers](#topic-openxr-composition-layers)
- [OpenXR render models](#topic-openxr-render-models)
- [Packed scene serialization](#topic-packed-scenes)
- [Post-import processing](#topic-post-import-processing)
- [Primitive references](#topic-primitive-references)
- [Project catalog and selection](#topic-project-catalog)
- [Property-list and scene-property helpers](#topic-property-introspection)
- [Raycast occlusion culling](#topic-raycast-occlusion-culling)
- [Replicated property synchronization](#topic-replicated-property-synchronization)
- [Resources and asset lifecycle](#topic-resources)
- [Runtime configuration](#topic-runtime-configuration)
- [2D and 3D scene authoring](#topic-scene-authoring)
- [Scene render data and buffers](#topic-scene-data-and-buffers)
- [Scene graph nodes](#topic-scene-graph)
- [Scene importing](#topic-scene-importing)
- [Scene state inspection](#topic-scene-state)
- [Scene graph and lifecycle](#topic-scene-tree)
- [Scene tree and signal connections](#topic-scene-tree-and-signal-connections)
- [SceneTree execution](#topic-scene-tree-execution)
- [Screen-space and environment effects](#topic-screen-and-environment-effects)
- [3D physics nodes](#topic-three-dimensional-physics)
- [Tile authoring](#topic-tile-authoring)
- [2D physics nodes](#topic-two-dimensional-physics)
- [Undo and redo history](#topic-undo-redo-history)

<a id="topic-scripting"></a>
## 337. Script resources and instances

**Scope:** First-party

**Builds on:** [Class metadata and reflection](#topic-reflection), [Resources and asset lifecycle](#topic-resources).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The native module represents script validity, reload state, members, and compiler/analyzer relationships.

GDScript is a Script Resource whose compiled members and functions supply script instances to Object-derived engine objects.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class GDScriptNativeClass : public RefCounted {
	GDCLASS(GDScriptNativeClass, RefCounted);

	StringName name;

protected:
	bool _get(const StringName &p_name, Variant &r_ret) const;
	static void _bind_methods();

public:
	_FORCE_INLINE_ const StringName &get_name() const { return name; }
	Variant _new();
	Object *instantiate();
```

Source: `modules/gdscript/gdscript.h` — GDScript (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

GDScript, Node

Evidence:
- Code: `modules/gdscript/gdscript.h:57` — GDScript
- Code: `modules/gdscript/gdscript.h` — GDScript::MemberInfo

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)
- [Contextual completion contracts](#topic-completion-contexts)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Project filesystem index](#topic-filesystem-project-index)
- [Fixture contracts](#topic-fixture-contracts)
- [GDScript bytecode compilation and execution](#topic-gdscript-bytecode-runtime)
- [GDScript editor services](#topic-gdscript-editor-services)
- [GDScript static analysis](#topic-gdscript-static-analysis)
- [managed C# script bridge and source generation](#topic-managed-csharp-scripting)
- [Script languages and instances](#topic-script-hosting)
- [Managed script registration metadata](#topic-script-registration-metadata)
- [Unicode data and properties](#topic-unicode-data)
- [Visual Shader source partition](#topic-visual-shader-module-build)

<a id="topic-webrtc-multiplayer-mesh"></a>
## 338. WebRTC multiplayer mesh

**Scope:** First-party

**Builds on:** [WebRTC data channels](#topic-webrtc-data-channels), [WebRTC peer connections](#topic-webrtc-peer-connections).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ConnectedPeer records support the multiplayer abstraction's peer tracking.

A WebRTCMultiplayerPeer keeps a mesh of WebRTC peer connections and their data channels.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class WebRTCMultiplayerPeer : public MultiplayerPeer {
	GDCLASS(WebRTCMultiplayerPeer, MultiplayerPeer);

protected:
	static void _bind_methods();

private:
	enum {
		CH_RELIABLE = 0,
		CH_ORDERED = 1,
		CH_UNRELIABLE = 2,
		CH_RESERVED_MAX = 3
	};
```

Source: `modules/webrtc/webrtc_multiplayer_peer.h` — WebRTCMultiplayerPeer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

WebRTCMultiplayerPeer

Evidence:
- Code: `modules/webrtc/webrtc_multiplayer_peer.h:37` — WebRTCMultiplayerPeer
- Code: `modules/webrtc/webrtc_multiplayer_peer.h:58` — ConnectedPeer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)

<a id="topic-android-platform-host"></a>
## 339. Android platform host

**Scope:** First-party

**Builds on:** [Engine bootstrap](#topic-engine-bootstrap).

**This prepares you for:** [Android JNI interoperation](#topic-android-jni-interop), [Android rendering and input](#topic-android-render-input).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Android partition has both native C++ platform classes and JVM-facing host classes.

Android platform hosting starts and manages the native engine from Android activity and library code, continuing the engine bootstrap on the mobile host.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```kotlin
*/
class Godot private constructor(val context: Context) {

	companion object {
		private val TAG = Godot::class.java.simpleName

		@Volatile private var INSTANCE: Godot? = null

		@JvmStatic
		fun getInstance(context: Context): Godot {
			return INSTANCE ?: synchronized(this) {
				INSTANCE ?: Godot(context.applicationContext).also { INSTANCE = it }
			}
		}
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt` — Godot (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/os_android.h:43` — OS_Android
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt:2` — Godot
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotActivity.kt:2` — GodotActivity

<a id="topic-editor-plugin-extension"></a>
## 340. Editor plug-in extension

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

EditorPlugin is the integration surface used by the editor’s built-in tool plug-ins as well as external extensions.

Plugins attach behavior through polymorphic C++ hook interfaces and may register import, scene-format, post-import, inspector, gizmo, debugger, and resource-conversion plugins.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorPlugin : public Node {
	GDCLASS(EditorPlugin, Node);
	friend class EditorData;

	bool input_event_forwarding_always_enabled = false;
	bool force_draw_over_forwarding_enabled = false;

	String last_main_screen_name;
	String plugin_version;

#ifndef DISABLE_DEPRECATED
	static inline HashMap<Control *, EditorDock *> legacy_docks;
```

Source: `editor/plugins/editor_plugin.h` — EditorPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/plugins/editor_plugin.h:59` — EditorPlugin
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin::add_scene_format_importer_plugin
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin::add_inspector_plugin

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [3D gizmo authoring](#topic-3d-gizmo-authoring)
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [Android export pipeline](#topic-android-export-pipeline)
- [Apple embedded packaging and signing](#topic-apple-embedded-packaging)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Editor plugin state and custom types](#topic-editor-plugin-state)
- [Editor scene sessions](#topic-editor-scene-sessions)
- [Custom inspector property editors](#topic-inspector-property-editors)
- [Platform-selective shader baking](#topic-platform-selective-shader-baking)
- [Reflective property inspection](#topic-property-inspection)
- [Resources and asset lifecycle](#topic-resources)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-editor-plugin-lifecycle"></a>
## 341. Editor plugin lifecycle

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Custom inspector property editors](#topic-inspector-property-editors).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation favors many focused plugin subclasses over one monolithic editor tool.

C++ editor-plugin specializations package feature-specific editor behavior behind EditorPlugin-derived classes, including scene, script, and shader tools.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class MeshLibraryEditorPlugin : public EditorPlugin {
	GDCLASS(MeshLibraryEditorPlugin, EditorPlugin);

	static inline MeshLibraryEditorPlugin *singleton = nullptr;

	MeshLibraryEditor *mesh_library_editor = nullptr;

public:
	_FORCE_INLINE_ static MeshLibraryEditorPlugin *get_singleton() { return singleton; }

	virtual void edit(Object *p_node) override;
	virtual bool handles(Object *p_node) const override;
	virtual void make_visible(bool p_visible) override;
```

Source: `editor/scene/3d/mesh_library_editor_plugin.h` — MeshLibraryEditorPlugin : public EditorPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/3d/mesh_library_editor_plugin.h:126` — MeshLibraryEditorPlugin : public EditorPlugin
- Code: `editor/shader/shader_editor_plugin.h:48` — ShaderEditorPlugin : public EditorPlugin
- External (official, unverified (source anchor not found)): [EditorPlugin — Godot Engine documentation](https://docs.godotengine.org/en/latest/classes/class_editorplugin.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor plugins and customization hooks](#topic-editor-extensibility)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Android export pipeline](#topic-android-export-pipeline)
- [Apple embedded packaging and signing](#topic-apple-embedded-packaging)
- [Editor plugin state and custom types](#topic-editor-plugin-state)
- [Editor scene sessions](#topic-editor-scene-sessions)
- [Platform-selective shader baking](#topic-platform-selective-shader-baking)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Script resources and instances](#topic-scripting)

<a id="topic-export-orchestration"></a>
## 342. Export orchestration

**Scope:** First-party

**Builds on:** [Export presets](#topic-export-presets).

**This prepares you for:** [Target-platform export](#topic-target-platform-export).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

EditorExport is the registry joining target configuration with exporter extensions.

Export orchestration owns export presets, target platforms, and plugins, and maps a target platform to its runnable preset.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
EditorExport *EditorExport::singleton = nullptr;

void EditorExport::_save() {
	Ref<ConfigFile> config;
	Ref<ConfigFile> credentials;
	config.instantiate();
	credentials.instantiate();

	for (const KeyValue<Ref<EditorExportPlatform>, Ref<EditorExportPreset>> &E : runnable_presets) {
		config->set_value(RUNNABLE_SECTION_NAME, E.key->get_name(), E.value->get_name());
	}

	for (int i = 0; i < export_presets.size(); i++) {
```

Source: `editor/export/editor_export.cpp` — EditorExport (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorExport, EditorExportPreset

Evidence:
- Code: `editor/export/editor_export.h:38` — EditorExport
- Code: `editor/export/editor_export.cpp:38` — EditorExport

<a id="topic-astc-block-coding"></a>
## 343. ASTC block coding

**Scope:** Vendored: basis_universal

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ASTC logic is used both by the Basis transcoder and its HDR intermediate paths.

ASTC helpers represent physical and logical blocks with endpoint ranges, weight grids, partitions, and bit-level packing.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
{
	struct astc_blk
	{
		uint8_t m_vals[16];
	};

	// ASTC_HDR_MAX_VAL is the maximum color component value that can be encoded.
	// If the input has values higher than this, they need to be linearly scaled so all values are between [0,ASTC_HDR_MAX_VAL], and the linear scaling inverted in the shader.
	const float ASTC_HDR_MAX_VAL = 65216.0f; // actually MAX_QLOG12_VAL

	// Maximum usable QLOG encodings, and their floating point equivalent values, that don't result in NaN/Inf's.
	const uint32_t MAX_QLOG7 = 123;
	//const float MAX_QLOG7_VAL = 55296.0f;
```

Source: `thirdparty/basis_universal/transcoder/basisu_astc_hdr_core.h` — astc_blk (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_astc_helpers.h` — astc_helpers::astc_block
- Code: `thirdparty/basis_universal/transcoder/basisu_astc_helpers.h` — astc_helpers::log_astc_block
- Code: `thirdparty/basis_universal/transcoder/basisu_astc_hdr_core.h:7` — astc_blk

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU block texture conversion](#topic-block-texture-transcoding)
- [Image codec integration](#topic-image-codec-registration)

<a id="topic-bvh-construction"></a>
## 344. BVH construction

**Scope:** Vendored: embree

**Builds on:** [Primitive references](#topic-primitive-references).

**This prepares you for:** [BVH split heuristics](#topic-builder-heuristics), [BVH traversal](#topic-bvh-traversal), [Motion blur](#topic-motion-blur), [Scene construction and commit](#topic-scene-commit).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The generic builder evaluates leaf and split costs, while factory and specialized builders select node and primitive representations.

BVH construction partitions PrimRef records into nodes and leaves through configurable builder callbacks and settings.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
typedef HeuristicArrayBinningSAH<PrimRef,NUM_OBJECT_BINS> Heuristic;
      typedef GeneralBVHBuilder::BuildRecordT<Set,typename Heuristic::Split> BuildRecord;
      typedef GeneralBVHBuilder::Settings Settings;

      /*! special builder that propagates reduction over the tree */
      template<
      typename ReductionTy,
        typename CreateAllocFunc,
        typename CreateNodeFunc,
        typename UpdateNodeFunc,
        typename CreateLeafFunc,
        typename ProgressMonitor>

        static ReductionTy build(CreateAllocFunc createAlloc,
```

Source: `thirdparty/embree/kernels/builders/bvh_builder_sah.h` — GeneralBVHBuilder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BVHN, GeneralBVHBuilder, BVHNodeRecord

Evidence:
- Code: `thirdparty/embree/kernels/builders/bvh_builder_sah.h:20` — GeneralBVHBuilder
- Code: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build
- Code: `thirdparty/embree/kernels/bvh/bvh_node_ref.h:18` — BVHNodeRecord

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Allocation and reference ownership](#topic-allocation)
- [native 2D collision pipeline](#topic-physics-2d-collision-pipeline)
- [Spatial indexing and ray queries](#topic-spatial-indexing)

<a id="topic-basis-container-layout"></a>
## 345. Basis file layout

**Scope:** Vendored: basis_universal

**Builds on:** [Basis texture transcoding](#topic-basis-transcoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The transcoder reads a header followed by an array of slice descriptions addressed through header offsets.

A .basis file header and slice descriptors locate compressed slices and identify their texture form before Basis Universal decodes them.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
bool basisu_transcoder::get_file_info(const void* pData, uint32_t data_size, basisu_file_info& file_info) const
	{
		if (!validate_file_checksums(pData, data_size, false))
		{
			BASISU_DEVEL_ERROR("basisu_transcoder::get_file_info: validate_file_checksums failed\n");
			return false;
		}

		const basis_file_header* pHeader = static_cast<const basis_file_header*>(pData);
		const basis_slice_desc* pSlice_descs = reinterpret_cast<const basis_slice_desc*>(static_cast<const uint8_t*>(pData) + pHeader->m_slice_desc_file_ofs);

		file_info.m_version = pHeader->m_ver;
```

Source: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_transcoder::get_file_info (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BasisFileHeader, BasisSliceDescriptor

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_file_headers.h:99` — basis_file_header
- Code: `thirdparty/basis_universal/transcoder/basisu_file_headers.h:32` — basis_slice_desc
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp:12204` — basisu_transcoder::get_file_info

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Texture compression pipeline](#topic-texture-compression-pipeline)

<a id="topic-prefix-code-decoding"></a>
## 346. Bitstream and Huffman decoding

**Scope:** Vendored: basis_universal

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [Brotli stream decompression](#topic-brotli-stream-decompression).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This is the shared compression idiom visible in Basis internal decoding and Brotli decoder code.

Both Basis and Brotli implement bit readers and Huffman decoding tables to recover symbols from compressed bit streams.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
uint32_t BrotliBuildHuffmanTable(HuffmanCode* root_table,
                                 int root_bits,
                                 const uint16_t* const symbol_lists,
                                 uint16_t* count) {
  HuffmanCode code;       /* current table entry */
  HuffmanCode* table;     /* next available space in table */
  int len;                /* current code length */
  int symbol;             /* symbol index in original or sorted table */
  brotli_reg_t key;       /* prefix code */
  brotli_reg_t key_step;  /* prefix code addend */
  brotli_reg_t sub_key;   /* 2nd level table prefix code */
  brotli_reg_t sub_key_step;  /* 2nd level table prefix code addend */
  int step;               /* step size to replicate values in current table */
```

Source: `thirdparty/brotli/dec/huffman.c` — BrotliBuildHuffmanTable (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder_internal.h:146` — huffman_decoding_table
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder_internal.h:148` — bitwise_decoder
- Code: `thirdparty/brotli/dec/huffman.c:169` — BrotliBuildHuffmanTable

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Brotli stream decompression](#topic-brotli-stream-decompression)
- [Histograms and Huffman codes](#topic-histograms-and-huffman-codes)

<a id="topic-broad-phase"></a>
## 347. Broad-phase collision detection

**Scope:** Vendored: jolt_physics

**Builds on:** [Collision filtering](#topic-collision-filtering).

**This prepares you for:** [Narrow-phase collision queries](#topic-narrow-phase).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code provides brute-force and quadtree broad-phase implementations.

Broad-phase collision detection uses Body world-space AABox bounds and collision filtering to generate candidate body pairs and answer coarse queries.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// Used to do coarse collision detection operations to quickly prune out bodies that will not collide.
class JPH_EXPORT BroadPhase : public BroadPhaseQuery
{
public:
	/// Initialize the broadphase.
	/// @param inBodyManager The body manager singleton
	/// @param inLayerInterface Interface that maps object layers to broadphase layers.
	/// Note that the broadphase takes a pointer to the data inside inObjectToBroadPhaseLayer so this object should remain static.
	virtual void		Init(BodyManager *inBodyManager, const BroadPhaseLayerInterface &inLayerInterface);

	/// Should be called after many objects have been inserted to make the broadphase more efficient, usually done on startup only
	virtual void		Optimize()															{ /* Optionally overridden by implementation */ }

	/// Must be called just before updating the broadphase when none of the body mutexes are locked
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/BroadPhase/BroadPhase.h` — BroadPhase (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/BroadPhase/BroadPhase.h:7` — BroadPhase
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/BroadPhase/BroadPhaseQuadTree.h:14` — BroadPhaseQuadTree
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/BroadPhase/QuadTree.h:20` — QuadTree

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Primitive references](#topic-primitive-references)

<a id="topic-ui-composition"></a>
## 348. Control-tree user interfaces

**Scope:** First-party

**Builds on:** [Scene graph and lifecycle](#topic-scene-tree).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The same tree model supports widgets, dialogs, menus, graph editors, and editor controls.

Control-derived nodes compose user interfaces, and Container nodes automatically arrange their Control children.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<brief_description>
		Base class for all GUI containers.
	</brief_description>
	<description>
		Base class for all GUI containers. A [Container] automatically arranges its child controls in a certain way. This class can be inherited to make custom container types.
	</description>
	<tutorials>
		<link title="Using Containers">$DOCS_URL/tutorials/ui/gui_containers.html</link>
	</tutorials>
	<methods>
		<method name="_get_allowed_size_flags_horizontal" qualifiers="virtual const">
			<return type="PackedInt32Array" />
			<description>
				Implement to return a list of allowed horizontal [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
```

Source: `doc/classes/Container.xml` — Container (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Node

Evidence:
- Code: `doc/classes/Control.xml:2` — Control
- Code: `doc/classes/Container.xml:2` — Container
- Code: `doc/classes/GraphEdit.xml:2` — GraphEdit

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GUI controls and layout](#topic-gui-control-layout)
- [Parametric curve bases](#topic-curve-and-patch-bases)
- [Endian-safe binary I/O](#topic-endian-safe-binary-io)
- [GDScript bytecode compilation and execution](#topic-gdscript-bytecode-runtime)
- [GUI controls and layout](#topic-gui-controls)
- [ISA and platform portability](#topic-isa-portability)
- [Display, camera, video, and movie services](#topic-platform-presentation)
- [Regular-expression JIT code generation](#topic-regex-jit-codegen)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Themes and style boxes](#topic-themes-and-style-boxes)
- [UI themes and style boxes](#topic-ui-theming)
- [Version-control integration](#topic-version-control-integration)
- [WebSocket framing and event contexts](#topic-websocket-frame-events)

<a id="topic-feature-adaptive-tessellation"></a>
## 349. Feature-adaptive tessellation

**Scope:** Vendored: embree

**Builds on:** [Subdivision surface evaluation](#topic-subdivision-surface-evaluation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Feature-adaptive tessellation recursively partitions parameter ranges for patch evaluation; subdivision evaluation is needed to explain the patch evaluation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<typename Vertex, typename Vertex_t = Vertex>
      struct FeatureAdaptiveEval
      {
      public:
        
        typedef PatchT<Vertex,Vertex_t> Patch;
        typedef typename Patch::Ref Ref;
        typedef GeneralCatmullClarkPatchT<Vertex,Vertex_t> GeneralCatmullClarkPatch;
        typedef CatmullClark1RingT<Vertex,Vertex_t> CatmullClarkRing;
        typedef CatmullClarkPatchT<Vertex,Vertex_t> CatmullClarkPatch;
        typedef BSplinePatchT<Vertex,Vertex_t> BSplinePatch;
        typedef BezierPatchT<Vertex,Vertex_t> BezierPatch;
        typedef GregoryPatchT<Vertex,Vertex_t> GregoryPatch;
        typedef BilinearPatchT<Vertex,Vertex_t> BilinearPatch;
```

Source: `thirdparty/embree/kernels/subdiv/feature_adaptive_eval.h` — FeatureAdaptiveEval (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SubGrid

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/feature_adaptive_eval.h:13` — FeatureAdaptiveEval
- Code: `thirdparty/embree/kernels/subdiv/feature_adaptive_eval_grid.h:16` — FeatureAdaptiveEvalGrid
- Code: `thirdparty/embree/kernels/subdiv/tessellation.h:8` — tessellation

<a id="topic-gpu-resources-pipelines"></a>
## 350. GPU resources and pipelines

**Scope:** Vendored: metal-cpp

**Builds on:** [Metal-cpp object bridge](#topic-metal-cpp-object-bridge).

**This prepares you for:** [GPU command encoding](#topic-gpu-command-encoding), [MetalFX scaling and interpolation](#topic-metalfx-scaling).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Descriptors configure GPU resources and pipeline state before a device produces allocations, functions, or pipeline state objects.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
_MTL_INLINE MTL4::PipelineDescriptor* MTL4::PipelineDescriptor::alloc()
{
    return NS::Object::alloc<MTL4::PipelineDescriptor>(_MTL_PRIVATE_CLS(MTL4PipelineDescriptor));
}

_MTL_INLINE MTL4::PipelineDescriptor* MTL4::PipelineDescriptor::init()
{
    return NS::Object::init<MTL4::PipelineDescriptor>();
}

_MTL_INLINE NS::String* MTL4::PipelineDescriptor::label() const
{
    return Object::sendMessage<NS::String*>(this, _MTL_PRIVATE_SEL(label));
```

Source: `thirdparty/metal-cpp/Metal/MTL4PipelineState.hpp` — MTL4::PipelineDescriptor (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MTL::Buffer, MTL::Texture, MTL4::PipelineDescriptor

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTLDevice.hpp:236` — MTL::Device
- Code: `thirdparty/metal-cpp/Metal/MTL4PipelineState.hpp:122` — MTL4::PipelineDescriptor

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Android Gradle assembly](#topic-android-gradle-assembly)
- [APK signing and verification](#topic-apk-signing)
- [CFF font subsetting](#topic-cff-font-subsetting)
- [JPEG decompression and coefficient transcoding](#topic-jpeg-decompression-transcoding)
- [PNG streaming I/O and row transforms](#topic-png-stream-transforms)
- [Shader interface mapping and reflection](#topic-shader-interface-analysis)
- [Variable-font table subsetting](#topic-variable-font-subsetting)
- [WebP chunk parsing, incremental decode, and animation](#topic-webp-riff-decoding)
- [Resource-backed assets](#topic-resource-assets)

<a id="topic-godot-enet-socket-adaptation"></a>
## 351. Godot ENet socket adaptation

**Scope:** Vendored: enet

**Builds on:** [ENet wire commands](#topic-enet-wire-commands).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Godot ENet socket adapter supplies UDP and DTLS socket classes to the transport; ENet wire commands provide this adapter's transport-facing purpose.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/// Abstract ENet interface for UDP/DTLS.
class ENetGodotSocket {
public:
	virtual Error bind(IPAddress p_ip, uint16_t p_port) = 0;
	virtual Error get_socket_address(IPAddress *r_ip, uint16_t *r_port) = 0;
	virtual Error sendto(const uint8_t *p_buffer, int p_len, int &r_sent, IPAddress p_ip, uint16_t p_port) = 0;
	virtual Error recvfrom(uint8_t *p_buffer, int p_len, int &r_read, IPAddress &r_ip, uint16_t &r_port) = 0;
	virtual int set_option(ENetSocketOption p_option, int p_value) = 0;
	virtual void close() = 0;
	virtual void set_refuse_new_connections(bool p_enable) {} /* Only used by dtls server */
	virtual bool can_upgrade() { return false; } /* Only true in ENetUDP */
	virtual void disconnect_peer(ENetPeer *p_peer) {}
	virtual void bind_peer(ENetPeer *p_peer) {}
	virtual ~ENetGodotSocket() {}
```

Source: `thirdparty/enet/enet_godot.cpp` — ENetGodotSocket (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/enet/enet_godot.cpp:47` — ENetGodotSocket
- Code: `thirdparty/enet/enet_godot.cpp:56` — ENetUDP
- Code: `thirdparty/enet/enet_godot.cpp:62` — ENetDTLSClient
- Code: `thirdparty/enet/enet_godot.cpp:63` — ENetDTLSServer

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [ENet host and peer transport](#topic-enet-host-peer-transport)
- [ENet transport and multiplayer adaptation](#topic-enet-transport-and-multiplayer)
- [Networking and transport I/O](#topic-network-io)
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-graphite-shaping"></a>
## 352. Graphite segment shaping

**Scope:** Vendored: graphite

**Builds on:** [Binary font-table access](#topic-font-table-access), [Graphite SILF rule execution](#topic-graphite-rule-execution).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

A face selects the applicable SILF behavior, a font supplies scaled metrics, and a segment exposes shaped slots.

Graphite uses decoded font tables and rule passes to turn Unicode into a Segment whose linked Slots carry glyph and placement state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
struct gr_segment : public graphite2::Segment {};
```

Source: `thirdparty/graphite/src/inc/Segment.h` — graphite2::Segment (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Face, Font, Segment, Slot, FeatureVal

Evidence:
- Code: `thirdparty/graphite/src/inc/Segment.h:213` — graphite2::Segment
- Code: `thirdparty/graphite/src/Segment.cpp` — graphite2::Segment construction
- Code: `thirdparty/graphite/src/gr_segment.cpp` — gr_seg_cinfo, gr_seg_first_slot, and gr_seg_last_slot
- Code: `thirdparty/graphite/src/gr_slot.cpp` — gr_slot_next_in_segment and attachment traversal
- External (official, unverified (no local verification cache)): [Graphite technical overview](https://graphite.sil.org/graphite_techAbout.html), accessed 2026-07-15

<a id="topic-filter-callbacks"></a>
## 353. Intersection and occlusion filters

**Scope:** Vendored: embree

**Builds on:** [Intersection results](#topic-hit-results).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

User geometry also exposes externally supplied intersect and occluded function argument structures.

Geometry-installed and query-context filters process candidate hit results before they update an RTCRayHit or report occlusion.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* Arguments for RTCFilterFunctionN */
struct RTCFilterFunctionNArguments
{
  int* valid;
  void* geometryUserPtr;
  struct RTCRayQueryContext* context;
  struct RTCRayN* ray;
  struct RTCHitN* hit;
  unsigned int N;
};

/* Filter callback function */
typedef void (*RTCFilterFunctionN)(const struct RTCFilterFunctionNArguments* args);
```

Source: `thirdparty/embree/include/embree4/rtcore_common.h` — RTCFilterFunctionNArguments (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCFilterFunctionNArguments, RTCIntersectFunctionNArguments, RTCOccludedFunctionNArguments

Evidence:
- Code: `thirdparty/embree/kernels/geometry/filter.h:15` — runIntersectionFilter1Helper
- Code: `thirdparty/embree/kernels/common/accelset.h:14` — IntersectFunctionNArguments
- Code: `thirdparty/embree/include/embree4/rtcore_common.h:313` — RTCFilterFunctionNArguments

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Font streams](#topic-font-streams)
- [Web JavaScript bridge](#topic-web-javascript-bridge)
- [WebSocket framing and event contexts](#topic-websocket-frame-events)
- [Geometry resources](#topic-geometry-resources)
- [Primitive intersection](#topic-primitive-intersection)
- [Spatial predictive filters](#topic-spatial-predictive-filters)

<a id="topic-ktx2-container-transcoding"></a>
## 354. KTX2 container transcoding

**Scope:** Vendored: basis_universal

**Builds on:** [Basis texture transcoding](#topic-basis-transcoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

KTX2 support is an optional container path alongside direct .basis transcoding.

The KTX2 transcoder retains the source data, parses the KTX2 header and per-level index, then transcodes its texture levels.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// Thread-specific ETC1S/supercompressed UASTC transcoder state. (If you're not doing multithreading transcoding you can ignore this.)
	struct ktx2_transcoder_state
	{
		basist::basisu_transcoder_state m_transcoder_state;
		basisu::uint8_vec m_level_uncomp_data;
		int m_uncomp_data_level_index;

		void clear()
		{
			m_transcoder_state.clear();
			m_level_uncomp_data.clear();
			m_uncomp_data_level_index = -1;
		}
	};
```

Source: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — ktx2_transcoder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Ktx2Header, Ktx2LevelIndex, Ktx2TranscoderState

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:951` — ktx2_transcoder
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:721` — ktx2_header
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:741` — ktx2_level_index

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [KTX texture containers](#topic-ktx-texture-container)
- [Texture compression pipeline](#topic-texture-compression-pipeline)

<a id="topic-lossy-macroblock-encoding"></a>
## 355. Lossy macroblock encoding

**Scope:** Vendored: libwebp

**Builds on:** [Encoder configuration](#topic-encoder-configuration), [Picture planes and pixel ownership](#topic-input-picture-planes).

**This prepares you for:** [Transform, quantization, and rate-distortion search](#topic-transform-quantization-rate-distortion).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Analysis, iteration, quantization, filtering, and syntax emission are separate source modules in the lossy path.

The lossy encoder walks the configured picture planes by macroblock, choosing prediction modes and producing residual tokens for the VP8 stream.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// main entry point
int VP8EncAnalyze(VP8Encoder* const enc) {
  int ok = 1;
  const int do_segments =
      enc->config->emulate_jpeg_size ||   // We need the complexity evaluation.
      (enc->segment_hdr.num_segments > 1) ||
      (enc->method <= 1);  // for method 0 - 1, we need preds[] to be filled.
  if (do_segments) {
    const int last_row = enc->mb_h;
    const int total_mb = last_row * enc->mb_w;
#ifdef WEBP_USE_THREAD
    // We give a little more than a half work to the main thread.
    const int split_row = (9 * last_row + 15) >> 4;
    const int kMinSplitRow = 2;  // minimal rows needed for mt to be worth it
```

Source: `thirdparty/libwebp/src/enc/analysis_enc.c` — VP8EncAnalyze (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VP8Encoder, WebPPicture, WebPConfig

Evidence:
- Code: `thirdparty/libwebp/src/enc/analysis_enc.c:426` — VP8EncAnalyze
- Code: `thirdparty/libwebp/src/enc/iterator_enc.c:79` — VP8IteratorInit
- Code: `thirdparty/libwebp/src/enc/frame_enc.c:788` — VP8EncTokenLoop
- Code: `thirdparty/libwebp/src/enc/vp8i_enc.h:139` — struct VP8Encoder

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [GPU command encoding](#topic-gpu-command-encoding)
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-motion-blur-geometry"></a>
## 356. Motion-blur geometry

**Scope:** Vendored: embree

**Builds on:** [Triangle intersection algorithms](#topic-triangle-intersection-algorithms).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Motion-blur triangle code computes vertices at ray time before triangle intersection, so triangle intersection is required to interpret its candidates.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<int N, int K, bool filter>
    struct SubGridMBIntersectorKPluecker
    {
      typedef SubGridMBQBVHN<N> Primitive;
      typedef SubGridQuadMIntersectorKPluecker<4,K,filter> Precalculations;

      static __forceinline void intersect(const vbool<K>& valid_i, Precalculations& pre, RayHitK<K>& ray, RayQueryContext* context, const SubGrid& subgrid)
      {
        size_t m_valid = movemask(valid_i);
        while(m_valid)
        {
          size_t ID = bscf(m_valid);
          intersect(pre,ray,ID,context,subgrid);
        }
```

Source: `thirdparty/embree/kernels/geometry/subgrid_mb_intersector.h` — SubGridMBIntersectorKPluecker (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/trianglev_mb_intersector.h:15` — TriangleMvMBIntersector1Moeller
- Code: `thirdparty/embree/kernels/geometry/trianglev_mb_intersector.h:149` — TriangleMvMBIntersectorKPluecker
- Code: `thirdparty/embree/kernels/geometry/subgrid_mb_intersector.h:100` — SubGridMBIntersectorKPluecker

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)
- [Motion blur](#topic-motion-blur)
- [Ray query state](#topic-ray-query)
- [Main loop, OS, and time services](#topic-runtime-loop-time)

<a id="topic-font-table-validation"></a>
## 357. OpenType and TrueTypeGX validation

**Scope:** Vendored: freetype

**Builds on:** [SFNT table loading](#topic-sfnt-tables).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The OpenType module covers BASE, GDEF, GPOS, GSUB, JSTF, and MATH; the GX module contains validators for multiple AAT table families.

OpenType and TrueTypeGX validators check structured tables after SFNT parsing so higher layers can consume checked offsets and indices.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#define FT_MAKE_OPTION_SINGLE_OBJECT

#include "gxvbsln.c"
#include "gxvcommn.c"
#include "gxvfeat.c"
#include "gxvjust.c"
#include "gxvkern.c"
#include "gxvlcar.c"
#include "gxvmod.c"
#include "gxvmort.c"
#include "gxvmort0.c"
#include "gxvmort1.c"
#include "gxvmort2.c"
```

Source: `thirdparty/freetype/src/gxvalid/gxvalid.c` — FT_MAKE_OPTION_SINGLE_OBJECT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/freetype/src/otvalid/otvalid.c:19` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/gxvalid/gxvalid.c:20` — FT_MAKE_OPTION_SINGLE_OBJECT
- External (official, unverified (source anchor not found)): [OpenType Validation - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-ot_validation.html), accessed 2026-07-15
- External (official, verified): [TrueTypeGX/AAT Validation - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-gx_validation.html), accessed 2026-07-16

<a id="topic-openxr-dispatch"></a>
## 358. OpenXR dispatch forwarding

**Scope:** Vendored: openxr

**Builds on:** [OpenXR runtime loading](#topic-openxr-runtime-loading).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Generated loader entry points forward API calls through a dispatch table belonging to the selected runtime.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
extern "C" LOADER_EXPORT XRAPI_ATTR XrResult XRAPI_CALL xrWaitFrame(
    XrSession                                   session,
    const XrFrameWaitInfo*                      frameWaitInfo,
    XrFrameState*                               frameState) XRLOADER_ABI_TRY {
    LoaderInstance* loader_instance;
    XrResult result = ActiveLoaderInstance::Get(&loader_instance, "xrWaitFrame");
    if (XR_SUCCEEDED(result)) {
        result = loader_instance->DispatchTable()->WaitFrame(session, frameWaitInfo, frameState);
    }
    return result;
}
XRLOADER_ABI_CATCH_FALLBACK
```

Source: `thirdparty/openxr/src/loader/xr_generated_loader.cpp` — xrWaitFrame (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

XrGeneratedDispatchTableCore

Evidence:
- Code: `thirdparty/openxr/src/xr_generated_dispatch_table_core.h:41` — XrGeneratedDispatchTableCore
- Code: `thirdparty/openxr/src/loader/xr_generated_loader.cpp:415` — xrWaitFrame

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [Collision shapes](#topic-collision-shapes)
- [KTX texture containers](#topic-ktx-texture-container)
- [Temporal upscaling dispatch](#topic-temporal-upscaling-dispatch)

<a id="topic-physics-collision"></a>
## 359. Physics shapes, objects, and collision reports

**Scope:** First-party

**Builds on:** [Resources and asset lifecycle](#topic-resources), [Scene graph and lifecycle](#topic-scene-tree).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation provides parallel 2D and 3D object, shape, joint, and movement-report APIs.

2D and 3D CollisionObject nodes own Shape resources through shape owners; body movement returns KinematicCollision results.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<description>
		Abstract base class for 2D physics objects. [CollisionObject2D] can hold any number of [Shape2D]s for collision. Each shape must be assigned to a [i]shape owner[/i]. Shape owners are not nodes and do not appear in the editor, but are accessible through code using the [code]shape_owner_*[/code] methods.
		[b]Note:[/b] Only collisions between objects within the same canvas ([Viewport] canvas or [CanvasLayer]) are supported. The behavior of collisions between objects in different canvases is undefined.
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="_input_event" qualifiers="virtual">
			<return type="void" />
			<param index="0" name="viewport" type="Viewport" />
			<param index="1" name="event" type="InputEvent" />
			<param index="2" name="shape_idx" type="int" />
			<description>
				Detects unhandled mouse and touch [InputEvent]s through [param event] when they happen while hovering over the object. Gesture events are not detected. [param viewport] is the [Viewport] that the event originated in (for viewports other than the main one to be detected, [member Viewport.physics_object_picking] needs to be set to [code]true[/code]). [param shape_idx] is the index of the detected shape from [PhysicsServer2D].
```

Source: `doc/classes/CollisionObject2D.xml` — CollisionObject2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shape2D, CollisionShape2D, KinematicCollision2D

Evidence:
- Code: `doc/classes/CollisionObject2D.xml:2` — CollisionObject2D
- Code: `doc/classes/CollisionShape2D.xml:2` — CollisionShape2D
- Code: `doc/classes/KinematicCollision2D.xml:2` — KinematicCollision2D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision shapes](#topic-collision-shapes)
- [2D shapes, tiles, and paths](#topic-two-dimensional-content)
- [Resource-backed assets](#topic-resource-assets)

<a id="topic-primitive-intersection"></a>
## 360. Primitive intersection

**Scope:** Vendored: embree

**Builds on:** [Geometry resources](#topic-geometry-resources), [Intersection results](#topic-hit-results), [Ray query state](#topic-ray-query).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The indexed code includes triangle, quad, curve, line, cone, cylinder, disc, grid, object, and instance intersectors.

Primitive intersectors test a ray query against Geometry representations and update RTCRayHit through common intersection or occlusion epilogues.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<int M>
      struct ConeCurveIntersector1
      {
        typedef CurvePrecalculations1 Precalculations;
        
        struct ray_tfar {
          Ray& ray;
          __forceinline ray_tfar(Ray& ray) : ray(ray) {}
          __forceinline vfloat<M> operator() () const { return ray.tfar; };
        };

        template<typename Epilog>
        static __forceinline bool intersect(const vbool<M>& valid_i,
                                            Ray& ray,
```

Source: `thirdparty/embree/kernels/geometry/coneline_intersector.h` — ConeCurveIntersector1 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMIntersectorKMoeller, QuadMIntersector1MoellerTrumbore, ConeCurveIntersector1

Evidence:
- Code: `thirdparty/embree/kernels/geometry/quad_intersector_moeller.h:115` — QuadMIntersector1MoellerTrumbore
- Code: `thirdparty/embree/kernels/geometry/curve_intersector_virtual.h:38` — VirtualCurveIntersector
- Code: `thirdparty/embree/kernels/geometry/coneline_intersector.h:149` — ConeCurveIntersector1
- Code: `thirdparty/embree/kernels/geometry/intersector_epilog.h:21` — Intersect1Epilog1

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Intersection and occlusion filters](#topic-filter-callbacks)
- [Ray–primitive intersection](#topic-ray-primitive-intersection)
- [Ray query state](#topic-ray-query)
- [Raycast occlusion culling](#topic-raycast-occlusion-culling)
- [Triangle intersection algorithms](#topic-triangle-intersection-algorithms)

<a id="topic-serialization"></a>
## 361. RTTI-based serialization

**Scope:** Vendored: jolt_physics

**Place in the path:** Start here for this branch of the guide.

**This prepares you for:** [State recording and validation](#topic-state-recording).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The same registration pattern is used by settings, materials, constraints, skeletons, and curves.

RTTI-based serialization registers attributes for object types and writes or restores binary state through StreamIn and StreamOut.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// Attributes are members of classes that need to be serialized.
class SerializableAttribute
{
public:
	///@ Serialization functions
	using pGetMemberPrimitiveType = const RTTI * (*)();
	using pIsType = bool (*)(int inArrayDepth, EOSDataType inDataType, const char *inClassName);
	using pReadData = bool (*)(IObjectStreamIn &ioStream, void *inObject);
	using pWriteData = void (*)(IObjectStreamOut &ioStream, const void *inObject);
	using pWriteDataType = void (*)(IObjectStreamOut &ioStream);

	/// Constructor
								SerializableAttribute(const char *inName, uint inMemberOffset, pGetMemberPrimitiveType inGetMemberPrimitiveType, pIsType inIsType, pReadData inReadData, pWriteData inWriteData, pWriteDataType inWriteDataType) : mName(inName), mMemberOffset(inMemberOffset), mGetMemberPrimitiveType(inGetMemberPrimitiveType), mIsType(inIsType), mReadData(inReadData), mWriteData(inWriteData), mWriteDataType(inWriteDataType) { }
```

Source: `thirdparty/jolt_physics/Jolt/ObjectStream/SerializableAttribute.h` — SerializableAttribute (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ConstraintSettings, PhysicsMaterialSimple, Skeleton

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/ObjectStream/SerializableAttribute.h:36` — SerializableAttribute
- Code: `thirdparty/jolt_physics/Jolt/Core/RTTI.h` — RTTI::GetAttribute
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/PhysicsMaterialSimple.cpp` — PhysicsMaterialSimple::SaveBinaryState and RestoreBinaryState

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Engine object model](#topic-object-model)

<a id="topic-raster-image-input"></a>
## 362. Raster image input

**Scope:** Vendored: basis_universal

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Basis encoder dependency includes separate PNG and JPEG decoding implementations.

Raster image input adapters decode PNG and JPEG bytes into image buffers for texture-processing callers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// Call the get_total_bytes_read() method to determine the actual size of the JPEG stream after successful decoding.
	class jpeg_decoder_stream
	{
	public:
		jpeg_decoder_stream() { }
		virtual ~jpeg_decoder_stream() { }

		// The read() method is called when the internal input buffer is empty.
		// Parameters:
		// pBuf - input buffer
		// max_bytes_to_read - maximum bytes that can be written to pBuf
		// pEOF_flag - set this to true if at end of stream (no more bytes remaining)
		// Returns -1 on error, otherwise return the number of bytes actually written to the buffer (which may be 0).
		// Notes: This method will be called in a loop until you set *pEOF_flag to true or the internal buffer is full.
```

Source: `thirdparty/basis_universal/encoder/jpgd.h` — jpeg_decoder (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PngInfo

Evidence:
- Code: `thirdparty/basis_universal/encoder/pvpngreader.h` — pv_png::get_png_info
- Code: `thirdparty/basis_universal/encoder/pvpngreader.h` — pv_png::load_png
- Code: `thirdparty/basis_universal/encoder/jpgd.h:34` — jpeg_decoder

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [HDR, JPEG, and KTX loading](#topic-image-format-loading)
- [JPEG decompression and coefficient transcoding](#topic-jpeg-decompression-transcoding)
- [PNG streaming I/O and row transforms](#topic-png-stream-transforms)
- [Raster font rendering](#topic-raster-font-rendering)
- [Image decode pipelines](#topic-image-decode-pipeline)
- [PNG image codec](#topic-png-image-codec)

<a id="topic-packet-queries"></a>
## 363. Ray packets and streams

**Scope:** Vendored: embree

**Builds on:** [Intersection results](#topic-hit-results), [Ray query state](#topic-ray-query).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Packet traversal and packet primitive intersectors reuse the same query concepts across multiple lanes.

Packet execution packs several ray queries and their hit results into 4-, 8-, 16-, or N-wide layouts, with array-of-structures and structure-of-arrays stream readers.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* Ray structure for a packet of 4 rays */
struct RTC_ALIGN(16) RTCRay4
{
  float org_x[4];
  float org_y[4];
  float org_z[4];
  float tnear[4];

  float dir_x[4];
  float dir_y[4];
  float dir_z[4];
  float time[4];

  float tfar[4];
```

Source: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCRay4 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCRay4, RTCRayHit16, RTCRayNt, RayStreamSOA

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:55` — RTCRay4
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:268` — RTCRayHitNt
- Code: `thirdparty/embree/kernels/common/ray.h:474` — RayStreamSOA
- Code: `thirdparty/embree/kernels/bvh/node_intersector_packet.h:17` — TravRayK

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Compressed font stream adapters](#topic-compressed-font-stream-adapters)
- [Font streams](#topic-font-streams)
- [KTX texture containers](#topic-ktx-texture-container)
- [Lossy macroblock encoding](#topic-lossy-macroblock-encoding)
- [MP3 audio resources](#topic-mp3-audio-resources)
- [Packets, HTTP, JSON, and JSON-RPC](#topic-network-data-exchange)
- [Networking and transport I/O](#topic-network-io)
- [HTTP and multiplayer](#topic-networking)
- [Ogg packet sequences](#topic-ogg-packet-sequences)
- [Ogg pages, packets, and bit packing](#topic-ogg-pages-and-packets)
- [Ogg Vorbis audio streams](#topic-ogg-vorbis-audio-streams)
- [State recording and validation](#topic-state-recording)
- [Byte streams and socket servers](#topic-stream-networking)
- [Theora video streams](#topic-theora-video-streams)
- [Variant text parsing and writing](#topic-variant-text-serialization)
- [WebSocket peers](#topic-websocket-peers)
- [Zstandard stream compression](#topic-zstandard-stream-codec)
- [Dynamic arrays and dictionaries](#topic-variant-containers)

<a id="topic-resource-identifiers"></a>
## 364. Resource and server identifiers

**Scope:** First-party

**Builds on:** [Resource formats and serialization](#topic-resource-formats).

**This prepares you for:** [RenderingDevice descriptors and handles](#topic-rendering-device-resources).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

The two identifiers intentionally serve different lifetimes: project references versus live server objects.

ResourceUID maps project resource identifiers to paths, while an RID is an opaque session-only handle for a low-level server resource.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="RID" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		A handle for a [Resource]'s unique identifier.
	</brief_description>
	<description>
		The RID [Variant] type is used to access a low-level resource by its unique ID. RIDs are opaque, which means they do not grant access to the resource by themselves. They are used by the low-level server classes, such as [DisplayServer], [RenderingServer], [TextServer], etc.
		A low-level resource may correspond to a high-level [Resource], such as [Texture] or [Mesh].
		[b]Note:[/b] RIDs are only useful during the current session. It won't correspond to a similar resource if sent over a network, or loaded from a file at a later time.
		[b]Note:[/b] In a boolean context, an RID will evaluate to [code]false[/code] if it has the invalid ID [code]0[/code]. Otherwise, an RID will always evaluate to [code]true[/code]. This is equivalent to calling [method is_valid].
	</description>
	<tutorials>
	</tutorials>
	<constructors>
```

Source: `doc/classes/RID.xml` — RID (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ResourceUID, RID

Evidence:
- Code: `doc/classes/ResourceUID.xml:2` — ResourceUID
- Code: `doc/classes/RID.xml:2` — RID

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Generic container infrastructure](#topic-generic-containers)
- [Resource dependency resolution](#topic-resource-dependency-resolution)
- [Scene render data and buffers](#topic-scene-data-and-buffers)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-simd-ray-packets"></a>
## 365. SIMD ray packets

**Scope:** Vendored: embree

**Builds on:** [Triangle intersection algorithms](#topic-triangle-intersection-algorithms).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SIMD ray packets evaluate width-parameterized ray and triangle lanes, so triangle intersection is required to interpret lane-validity masks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<int M, int K>
    struct SphereIntersectorK
    {
      typedef CurvePrecalculationsK<K> Precalculations;

      template<typename Epilog>
      static __forceinline bool intersect(const vbool<M>& valid_i,
                                          RayK<K>& ray, size_t k,
                                          RayQueryContext* context,
                                          const Points* geom,
                                          const Precalculations& pre,
                                          const Vec4vf<M>& v0i,
                                          const Epilog& epilog)
      {
```

Source: `thirdparty/embree/kernels/geometry/sphere_intersector.h` — SphereIntersectorK (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h:247` — MoellerTrumboreIntersectorK
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_pluecker.h:237` — PlueckerIntersectorK
- Code: `thirdparty/embree/kernels/geometry/sphere_intersector.h:158` — SphereIntersectorK

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [CPU-specialized DSP backends](#topic-cpu-specialized-dsp)
- [Geometry families](#topic-geometry-families)
- [Ray query state](#topic-ray-query)
- [Optional SIMD codec paths](#topic-simd-accelerated-codecs)
- [ISA and platform portability](#topic-isa-portability)

<a id="topic-scene-state"></a>
## 366. Scene state inspection

**Scope:** First-party

**Builds on:** [Resource formats and serialization](#topic-resource-formats).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

It is an inspection representation rather than a mutable scene graph.

SceneState exposes a packed scene's resources, nodes, exported properties, overrides, and built-in scripts without instantiating the scene.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="SceneState" inherits="RefCounted" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Provides access to a scene file's information.
	</brief_description>
	<description>
		Maintains a list of resources, nodes, exported and overridden properties, and built-in scripts associated with a scene. They cannot be modified from a [SceneState], only accessed. Useful for peeking into what a [PackedScene] contains without instantiating it.
		This class cannot be instantiated directly, it is retrieved for a given scene as the result of [method PackedScene.get_state].
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="get_base_scene_state" qualifiers="const">
			<return type="SceneState" />
```

Source: `doc/classes/SceneState.xml` — SceneState (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SceneState, Resource

Evidence:
- Code: `doc/classes/SceneState.xml:2` — SceneState

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Packed scene serialization](#topic-packed-scenes)
- [Resource-backed assets](#topic-resource-assets)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)
- [Runtime property groups and hints](#topic-sdl-runtime-properties)

<a id="topic-shader-interface-analysis"></a>
## 367. Shader interface mapping and reflection

**Scope:** Vendored: glslang

**Builds on:** [Shader source compilation](#topic-shader-source-compilation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Program-facing reflection records are derived from shader declarations and linker objects.

IO mapping and reflection use intermediate-tree traversal to expose a compiled shader interface as uniform, buffer, atomic-counter, and pipeline input/output entries.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
class TReflectionTraverser : public TIntermTraverser {
public:
    TReflectionTraverser(const TIntermediate& i, TReflection& r) :
	                     TIntermTraverser(), intermediate(i), reflection(r), updateStageMasks(true) { }

    virtual bool visitBinary(TVisit, TIntermBinary* node);
    virtual void visitSymbol(TIntermSymbol* base);

    // Add a simple reference to a uniform variable to the uniform database, no dereference involved.
    // However, no dereference doesn't mean simple... it could be a complex aggregate.
    void addUniform(const TIntermSymbol& base)
    {
        if (processedDerefs.find(&base) == processedDerefs.end()) {
```

Source: `thirdparty/glslang/glslang/MachineIndependent/reflection.cpp` — TReflectionTraverser (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TProgram, TObjectReflection

Evidence:
- Code: `thirdparty/glslang/glslang/MachineIndependent/iomapper.h` — TDefaultIoResolverBase and TGlslIoMapper
- Code: `thirdparty/glslang/glslang/MachineIndependent/reflection.h:56` — TReflection
- Code: `thirdparty/glslang/glslang/MachineIndependent/reflection.cpp:78` — TReflectionTraverser
- External (primary, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.html), accessed 2026-07-15
- External (primary, unverified (source anchor not found)): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry buffer storage](#topic-buffer-storage)
- [BVH traversal](#topic-bvh-traversal)
- [GPU resources and pipelines](#topic-gpu-resources-pipelines)
- [Runtime class metadata](#topic-runtime-metadata)
- [Shader interface metadata](#topic-shader-interface-metadata)
- [Shader reflection](#topic-shader-reflection)

<a id="topic-shader-interface-metadata"></a>
## 368. Shader interface metadata

**Scope:** Vendored: spirv-reflect

**Builds on:** [SPIR-V shader reflection](#topic-spirv-shader-reflection).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The module and entry-point structures use counts plus arrays or pointers to express the reflected relationships.

Reflection exposes descriptor sets, interface variables, push-constant blocks, and specialization constants from one reflected shader module; it depends on SPIR-V shader reflection because it is contained in the reflected shader module.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/*! @struct SpvReflectEntryPoint

 */
typedef struct SpvReflectEntryPoint {
  const char*                       name;
  uint32_t                          id;

  SpvExecutionModel                 spirv_execution_model;
  SpvReflectShaderStageFlagBits     shader_stage;

  uint32_t                          input_variable_count;
  SpvReflectInterfaceVariable**     input_variables;
  uint32_t                          output_variable_count;
```

Source: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectEntryPoint (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SpvReflectDescriptorSet, SpvReflectDescriptorBinding, SpvReflectInterfaceVariable, SpvReflectBlockVariable, SpvReflectSpecializationConstant

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectDescriptorSet, SpvReflectDescriptorBinding
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectInterfaceVariable, SpvReflectBlockVariable, SpvReflectSpecializationConstant
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:530` — SpvReflectEntryPoint

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Runtime class metadata](#topic-runtime-metadata)
- [Shader interface mapping and reflection](#topic-shader-interface-analysis)
- [Metal drawable presentation](#topic-metal-presentation)
- [OpenXR extension wrappers](#topic-openxr-extension-wrappers)
- [Shader reflection](#topic-shader-reflection)
- [Web JavaScript bridge](#topic-web-javascript-bridge)
- [WebXR session bridge](#topic-webxr-session-bridge)
- [zlib stream compression](#topic-zlib-stream-codec)
- [Shader programs and material parameters](#topic-shader-programs)
- [SPIR-V generation](#topic-spirv-generation)
- [SPIR-V intermediate representation](#topic-spirv-intermediate-representation)
- [SPIR-V rewriting and optimization](#topic-spirv-rewriting)

<a id="topic-signed-distance-fields"></a>
## 369. Signed-distance-field glyph rendering

**Scope:** Vendored: freetype

**Builds on:** [Glyph rasterization](#topic-glyph-rasterization).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Common SDF support includes distance normalization and shared renderer properties.

The SDF and BSDF renderers use rasterization to populate signed-distance output bitmaps with configurable parameters.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#define FT_MAKE_OPTION_SINGLE_OBJECT

#include "ftsdfrend.c"
#include "ftsdfcommon.c"
#include "ftbsdf.c"
#include "ftsdf.c"


/* END */
```

Source: `thirdparty/freetype/src/sdf/sdf.c` — FT_MAKE_OPTION_SINGLE_OBJECT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/freetype/src/sdf/sdf.c:21` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/sdf/ftsdfcommon.c` — ft_sdf_format
- External (official, verified): [Scanline Converter - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-raster.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Glyph rasterization](#topic-glyph-rasterization)

<a id="topic-subgrid-intersection"></a>
## 370. Subgrid intersection

**Scope:** Vendored: embree

**Builds on:** [Triangle intersection algorithms](#topic-triangle-intersection-algorithms).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Subgrid intersection uses grid identifiers to load a quad neighborhood and applies triangle intersection algorithms to its triangles; triangle intersection algorithms are needed to interpret those tests.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/* Stores M quads from an indexed face set */
      struct SubGrid
      {
        /* Virtual interface to query information about the quad type */
        struct Type : public PrimitiveType
        {
          const char* name() const;
          size_t sizeActive(const char* This) const;
          size_t sizeTotal(const char* This) const;
          size_t getBytes(const char* This) const;
        };
        static Type type;

      public:
```

Source: `thirdparty/embree/kernels/geometry/subgrid.h` — SubGrid (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SubGrid, GridMesh

Evidence:
- Code: `thirdparty/embree/kernels/geometry/subgrid.h:13` — SubGrid
- Code: `thirdparty/embree/kernels/geometry/subgrid_intersector.h:21` — SubGridIntersector1Moeller
- Code: `thirdparty/embree/kernels/geometry/subgrid_intersector.h:124` — SubGridIntersector1Pluecker

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry families](#topic-geometry-families)

<a id="topic-textures-and-placeholders"></a>
## 371. Texture resources and fallback placeholders

**Scope:** First-party

**Builds on:** [Resource formats and serialization](#topic-resource-formats).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Layered and 3D textures require consistently sized image data across their layers.

Texture resources represent 2D, 3D, layered, and RenderingDevice-backed image data; placeholder texture classes preserve limited shape information when source texture data is unavailable or omitted.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="PlaceholderTexture2D" inherits="Texture2D" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Placeholder class for a 2-dimensional texture.
	</brief_description>
	<description>
		This class is used when loading a project that uses a [Texture2D] subclass in 2 conditions:
		- When running the project exported in dedicated server mode, only the texture's dimensions are kept (as they may be relied upon for gameplay purposes or positioning of other elements). This allows reducing the exported PCK's size significantly.
		- When this subclass is missing due to using a different engine version or build (e.g. modules disabled).
		[b]Note:[/b] This is not intended to be used as an actual texture for rendering. It is not guaranteed to work like one in shaders or materials (for example when calculating UV).
	</description>
	<tutorials>
	</tutorials>
	<members>
```

Source: `doc/classes/PlaceholderTexture2D.xml` — PlaceholderTexture2D (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RDTextureFormat, RID

Evidence:
- Code: `doc/classes/Texture2D.xml:2` — Texture2D
- Code: `doc/classes/TextureLayered.xml:2` — TextureLayered
- Code: `doc/classes/Texture2DRD.xml:2` — Texture2DRD
- Code: `doc/classes/PlaceholderTexture2D.xml:2` — PlaceholderTexture2D

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Collision shapes](#topic-collision-shapes)
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [RenderingDevice descriptors and handles](#topic-rendering-device-resources)
- [Resource-backed assets](#topic-resource-assets)

<a id="topic-tile-libraries"></a>
## 372. Tile libraries, cells, and patterns

**Scope:** First-party

**Builds on:** [Resource formats and serialization](#topic-resource-formats).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Tile sources can expose atlas-backed or scene-backed tiles.

A TileSet Resource owns tile sources, identifies tiles by source, atlas-coordinate, and alternative IDs, and supplies tile data to TileMapLayer cells and TileMapPattern copies.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="TileData" inherits="Object" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Settings for a single tile in a [TileSet].
	</brief_description>
	<description>
		[TileData] object represents a single tile in a [TileSet]. It is usually edited using the tileset editor, but it can be modified at runtime using [method TileMapLayer._tile_data_runtime_update].
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="add_collision_polygon">
			<return type="void" />
			<param index="0" name="layer_id" type="int" />
```

Source: `doc/classes/TileData.xml` — TileData (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

TileSet, TileData

Evidence:
- Code: `doc/classes/TileSet.xml:2` — TileSet
- Code: `doc/classes/TileData.xml:2` — TileData
- Code: `doc/classes/TileMapPattern.xml:2` — TileMapPattern

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Resources and asset lifecycle](#topic-resources)
- [Tile authoring](#topic-tile-authoring)
- [Tile map layer data](#topic-tilemap-layer-data)
- [2D shapes, tiles, and paths](#topic-two-dimensional-content)

<a id="topic-variable-font-data"></a>
## 373. TrueType GX variation data

**Scope:** Vendored: freetype

**Builds on:** [SFNT table loading](#topic-sfnt-tables).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The variation loader is compiled with the TrueType driver and models avar correspondence data.

TrueType GX variation loading maps a face's variation tables and coordinates into variable-font state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#ifdef TT_CONFIG_OPTION_GX_VAR_SUPPORT


#define FT_Stream_FTell( stream )                         \
          (FT_ULong)( (stream)->cursor - (stream)->base )
#define FT_Stream_SeekSet( stream, off )                               \
          (stream)->cursor =                                           \
            ( (off) < (FT_ULong)( (stream)->limit - (stream)->base ) ) \
                        ? (stream)->base + (off)                       \
                        : (stream)->limit


  /* some macros we need */
```

Source: `thirdparty/freetype/src/truetype/ttgxvar.c` — TT_CONFIG_OPTION_GX_VAR_SUPPORT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/freetype/src/truetype/ttgxvar.c:60` — TT_CONFIG_OPTION_GX_VAR_SUPPORT
- Code: `thirdparty/freetype/src/truetype/ttgxvar.h:41` — GX_AVarCorrespondenceRec_
- External (official, verified): [OpenType Font Variations, TrueType GX, and Adobe MM Fonts - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Variable-font table subsetting](#topic-variable-font-subsetting)

<a id="topic-worker-parallelism"></a>
## 374. Worker-based parallelism

**Scope:** Vendored: libwebp

**Builds on:** [Encoder configuration](#topic-encoder-configuration).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Thread behavior is abstracted behind a library interface instead of direct thread calls in encoder modules.

The encoding implementation obtains a WebPWorkerInterface and conditionally splits analysis or lossless work when thread-level configuration permits it.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
WebPPicture picture_side;
  const WebPWorkerInterface* const worker_interface = WebPGetWorkerInterface();
  int ok_main;

  if (enc_main == NULL || !VP8LBitWriterInit(&bw_side, 0)) {
    VP8LEncoderDelete(enc_main);
    return WebPEncodingSetError(picture, VP8_ENC_ERROR_OUT_OF_MEMORY);
  }

  // Avoid "garbage value" error from Clang's static analysis tool.
  if (!WebPPictureInit(&picture_side)) {
    goto Error;
  }
```

Source: `thirdparty/libwebp/src/enc/vp8l_enc.c` — WebPGetWorkerInterface (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VP8Encoder

Evidence:
- Code: `thirdparty/libwebp/src/utils/thread_utils.h:71` — WebPWorkerInterface
- Code: `thirdparty/libwebp/src/utils/thread_utils.c:215` — WebPGetWorkerInterface
- Code: `thirdparty/libwebp/src/enc/analysis_enc.c` — WebPGetWorkerInterface, do_mt
- Code: `thirdparty/libwebp/src/enc/vp8l_enc.c:1705` — WebPGetWorkerInterface

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Thread synchronization](#topic-concurrency)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Task scheduling and synchronization](#topic-task-scheduling)

<a id="topic-x509-certificate-processing"></a>
## 375. X.509 certificate processing

**Scope:** Vendored: mbedtls

**Builds on:** [TLS connection and session state](#topic-tls-connection-state).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation is split into certificate, request, revocation-list, OID, and writing translation units.

TLS credential processing parses, writes, and verifies certificates, certificate requests, revocation lists, names, and related ASN.1 data.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
*/
int mbedtls_x509_crt_verify_with_profile(mbedtls_x509_crt *crt,
                                         mbedtls_x509_crt *trust_ca,
                                         mbedtls_x509_crl *ca_crl,
                                         const mbedtls_x509_crt_profile *profile,
                                         const char *cn, uint32_t *flags,
                                         int (*f_vrfy)(void *, mbedtls_x509_crt *, int, uint32_t *),
                                         void *p_vrfy)
{
    return x509_crt_verify_restartable_ca_cb(crt, trust_ca, ca_crl,
                                             NULL, NULL,
                                             profile, cn, flags,
                                             f_vrfy, p_vrfy, NULL);
}
```

Source: `thirdparty/mbedtls/library/x509_crt.c` — mbedtls_x509_crt_verify_with_profile (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `thirdparty/mbedtls/include/mbedtls/x509_crt.h` — mbedtls_x509_crt API
- Code: `thirdparty/mbedtls/library/x509_crt.c:3147` — mbedtls_x509_crt_verify_with_profile
- Code: `thirdparty/mbedtls/library/x509_csr.c` — X.509 CSR parsing implementation
- Code: `thirdparty/mbedtls/library/x509_oid.c` — X.509 OID lookup implementation
- External (official, verified): [Mbed TLS documentation hub](https://mbed-tls.readthedocs.io/en/latest/), accessed 2026-07-16

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Mbed TLS-backed crypto and transport](#topic-tls-crypto-backend)

<a id="topic-managed-native-interop"></a>
## 376. Managed-native interop

**Scope:** First-party

**Place in the path:** Start here for this branch of the guide.

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The bridge has a size-checked unmanaged callback table and native runtime interop table with matching ordering requirements.

Unsafe C# connects managed values and callbacks across a fixed native/managed ABI using variant conversion, GC handles, and function-pointer callbacks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

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
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/Marshaling.cs:17` — Marshaling
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs:8` — ManagedCallbacks
- Code: `modules/mono/glue/runtime_interop.cpp:1863` — godotsharp::get_runtime_interop_funcs

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Character-encoding conversion](#topic-character-encoding-conversion)
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Dynamic runtime values](#topic-dynamic-variant)

<a id="topic-android-jni-interop"></a>
## 377. Android JNI interoperation

**Scope:** First-party

**Builds on:** [Android platform host](#topic-android-platform-host).

**This prepares you for:** [Android plugins and signals](#topic-android-plugin-signals), [Android storage bridge](#topic-android-storage-bridge).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Native wrappers, Java bridge classes, and callable adapters form the implemented boundary.

JNI interoperation converts values and directs callbacks across the Android platform host’s Java/native boundary using C++ Variant marshalling.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
/**************************************************************************/
/*  callable_jni.cpp                                                      */
/**************************************************************************/
/*                         This file is part of:                          */
/*                             GODOT ENGINE                               */
/*                        https://godotengine.org                         */
/**************************************************************************/
/* Copyright (c) 2014-present Godot Engine contributors (see AUTHORS.md). */
/* Copyright (c) 2007-2014 Juan Linietsky, Ariel Manzur.                  */
/*                                                                        */
/* Permission is hereby granted, free of charge, to any person obtaining  */
/* a copy of this software and associated documentation files (the        */
/* "Software"), to deal in the Software without restriction, including    */
/* without limitation the rights to use, copy, modify, merge, publish,    */
```

Source: `platform/android/variant/callable_jni.cpp` — callable_jni (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Callable

Evidence:
- Code: `platform/android/java_class_wrapper.h` — JavaClassWrapper
- Code: `platform/android/java_class_wrapper.cpp` — JavaClassWrapper::_jobject_to_variant
- Code: `platform/android/variant/callable_jni.cpp:2` — callable_jni
- External (official, unverified (source anchor not found)): [JNI tips](https://developer.android.com/ndk/guides/jni-tips), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Dynamic values and dictionaries](#topic-dynamic-values)
- [Dynamic runtime values](#topic-dynamic-variant)

<a id="topic-android-render-input"></a>
## 378. Android rendering and input

**Scope:** First-party

**Builds on:** [Android platform host](#topic-android-platform-host).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation includes separate GL and Vulkan render-view classes plus a shared Android input handler.

Android rendering views and native rendering drivers carry surface and input work from the Android platform host.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```java
*/
class GodotGLRenderView extends GLSurfaceView implements GodotRenderView {
	private final Godot godot;
	private final GodotInputHandler inputHandler;
	private final GodotRenderer godotRenderer;
	private final SparseArray<PointerIcon> customPointerIcons = new SparseArray<>();

	public GodotGLRenderView(Godot godot, GodotInputHandler inputHandler, XRMode xrMode, boolean useDebugOpengl, boolean shouldBeTranslucent) {
		super(godot.getContext());

		this.godot = godot;
		this.inputHandler = inputHandler;
		this.godotRenderer = new GodotRenderer();
		setPointerIcon(PointerIcon.getSystemIcon(getContext(), PointerIcon.TYPE_DEFAULT));
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotGLRenderView.java` — GodotGLRenderView (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotGLRenderView.java:2` — GodotGLRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotVulkanRenderView.java:2` — GodotVulkanRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/input/GodotInputHandler.java:2` — GodotInputHandler
- Code: `platform/android/rendering_context_driver_vulkan_android.h:39` — RenderingContextDriverVulkanAndroid

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Rendering backend abstraction](#topic-rendering-backends)

<a id="topic-inspector-property-editors"></a>
## 379. Custom inspector property editors

**Scope:** First-party

**Builds on:** [Editor plugin lifecycle](#topic-editor-plugin-lifecycle).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

This path is used for gradients, materials, textures, particle-process ranges, GUI controls, fonts, and style boxes.

The implementation uses editor-plugin lifecycle integration to replace or extend property editing with specialized EditorInspectorPlugin and EditorProperty classes.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorInspectorPluginGradient : public EditorInspectorPlugin {
	GDCLASS(EditorInspectorPluginGradient, EditorInspectorPlugin);

public:
	virtual bool can_handle(Object *p_object) override;
	virtual void parse_begin(Object *p_object) override;
};

class GradientEditorPlugin : public EditorPlugin {
	GDCLASS(GradientEditorPlugin, EditorPlugin);

public:
	virtual String get_plugin_name() const override { return "Gradient"; }
```

Source: `editor/scene/gradient_editor_plugin.h` — EditorInspectorPluginGradient : public EditorInspectorPlugin (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `editor/scene/gradient_editor_plugin.h:129` — EditorInspectorPluginGradient : public EditorInspectorPlugin
- Code: `editor/scene/particle_process_material_editor_plugin.h:43` — ParticleProcessMaterialMinMaxPropertyEditor : public EditorProperty
- Code: `editor/scene/gui/font_config_plugin.h` — EditorInspectorPluginFontVariation and EditorPropertyOTVariation
- External (official, unverified (source anchor not found)): [Inspector plugins — Godot Engine documentation](https://docs.godotengine.org/en/stable/tutorials/plugins/editor/inspector_plugins.html), accessed 2026-07-15

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Editor plugins and customization hooks](#topic-editor-extensibility)
- [Editor plug-in extension](#topic-editor-plugin-extension)

<a id="topic-target-platform-export"></a>
## 380. Target-platform export

**Scope:** First-party

**Builds on:** [Export orchestration](#topic-export-orchestration).

**This prepares you for:** [Apple embedded packaging and signing](#topic-apple-embedded-packaging).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The base platform type supplies common packaging facilities while derived platform types supply target behavior.

A target platform implementation supplies target-specific validation, option, run, and project, pack, or ZIP export behavior to export orchestration.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class EditorExportPlatform : public RefCounted {
	GDCLASS(EditorExportPlatform, RefCounted);

protected:
	static void _bind_methods();

public:
	typedef Error (*EditorExportSaveFunction)(const Ref<EditorExportPreset> &p_preset, void *p_userdata, const String &p_path, const Vector<uint8_t> &p_data, int p_file, int p_total, const Vector<String> &p_enc_in_filters, const Vector<String> &p_enc_ex_filters, const Vector<uint8_t> &p_key, uint64_t p_seed, bool p_delta);
	typedef Error (*EditorExportRemoveFunction)(const Ref<EditorExportPreset> &p_preset, void *p_userdata, const String &p_path);
	typedef Error (*EditorExportSaveSharedObject)(const Ref<EditorExportPreset> &p_preset, void *p_userdata, const SharedObject &p_so);

	enum DebugFlags {
		DEBUG_FLAG_DUMB_CLIENT = 1,
```

Source: `editor/export/editor_export_platform.h` — EditorExportPlatform (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorExportPlatform, EditorExportPlatform::ExportMessage

Evidence:
- Code: `editor/export/editor_export_platform.h:51` — EditorExportPlatform
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::save_pack
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::save_zip

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Running projects and instances](#topic-run-management)
- [Platform export packaging](#topic-platform-exports)
- [ZIP archive I/O](#topic-zip-archive-io)

<a id="topic-builder-heuristics"></a>
## 381. BVH split heuristics

**Scope:** Vendored: embree

**Builds on:** [BVH construction](#topic-bvh-construction), [Primitive references](#topic-primitive-references).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The supplied builders calculate costs from bounds, primitive counts, block size, traversal cost, and intersection cost.

SAH, Morton, spatial, temporal, strand, open-merge, and motion-blur heuristics choose how PrimRef records are divided while constructing a BVH.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/*! maps bounding box to morton code */
      struct MortonCodeMapping
      {
        static const size_t LATTICE_BITS_PER_DIM = 10;
        static const size_t LATTICE_SIZE_PER_DIM = size_t(1) << LATTICE_BITS_PER_DIM;

        vfloat4 base;
        vfloat4 scale;

        __forceinline MortonCodeMapping(const BBox3fa& bounds)
        {
          base  = (vfloat4)bounds.lower;
          const vfloat4 diag  = (vfloat4)bounds.upper - (vfloat4)bounds.lower;
          scale = select(diag > vfloat4(1E-19f), rcp(diag) * vfloat4(LATTICE_SIZE_PER_DIM * 0.99f),vfloat4(0.0f));
```

Source: `thirdparty/embree/kernels/builders/bvh_builder_morton.h` — MortonCodeMapping (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BinMapping, MortonCodeMapping, SpatialBinMapping, HeuristicMBlurTemporalSplit

Evidence:
- Code: `thirdparty/embree/kernels/builders/heuristic_binning.h:17` — BinMapping
- Code: `thirdparty/embree/kernels/builders/bvh_builder_morton.h:71` — MortonCodeMapping
- Code: `thirdparty/embree/kernels/builders/heuristic_spatial.h:17` — SpatialBinMapping
- Code: `thirdparty/embree/kernels/builders/heuristic_timesplit_array.h:17` — HeuristicMBlurTemporalSplit

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Motion blur](#topic-motion-blur)
- [Spatial indexing and ray queries](#topic-spatial-indexing)

<a id="topic-bvh-traversal"></a>
## 382. BVH traversal

**Scope:** Vendored: embree

**Builds on:** [BVH construction](#topic-bvh-construction), [Ray query state](#topic-ray-query).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code includes scalar, packet, hybrid, frustum, point-query, quantized-node, and oriented-node traversal paths.

Traversal tests the committed BVH with a ray query, orders candidate nodes, and calls primitive intersectors for reached leaves.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
template<int N, int types, bool robust, typename PrimitiveIntersector1>
    class BVHNIntersector1
    {
      /* shortcuts for frequently used types */
      typedef typename PrimitiveIntersector1::Precalculations Precalculations;
      typedef typename PrimitiveIntersector1::Primitive Primitive;
      typedef BVHN<N> BVH;
      typedef typename BVH::NodeRef NodeRef;
      typedef typename BVH::AABBNode AABBNode;
      typedef typename BVH::AABBNodeMB4D AABBNodeMB4D;

      static const size_t stackSize = 1+(N-1)*BVH::maxDepth+3; // +3 due to 16-wide store

    public:
```

Source: `thirdparty/embree/kernels/bvh/bvh_intersector1.h` — BVHNIntersector1 (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BVHNIntersector1, BVHNIntersectorKHybrid, TravRay

Evidence:
- Code: `thirdparty/embree/kernels/bvh/bvh_intersector1.h:16` — BVHNIntersector1
- Code: `thirdparty/embree/kernels/bvh/bvh_intersector_hybrid.h:20` — BVHNIntersectorKHybrid
- Code: `thirdparty/embree/kernels/bvh/node_intersector1.h:152` — TravRay

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Canvas and viewport presentation](#topic-canvas-and-viewports)
- [COLLADA scene interchange](#topic-collada-import)
- [Expression parsing and evaluation](#topic-expression-evaluation)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [HTTP and multiplayer](#topic-networking)
- [ObjectDB snapshots](#topic-objectdb-snapshots)
- [Packed scene serialization](#topic-packed-scenes)
- [Post-import processing](#topic-post-import-processing)
- [Ray query state](#topic-ray-query)
- [Scene graph nodes](#topic-scene-graph)
- [Scene importing](#topic-scene-importing)
- [Scene graph and lifecycle](#topic-scene-tree)
- [SceneTree execution](#topic-scene-tree-execution)
- [Shader interface mapping and reflection](#topic-shader-interface-analysis)
- [Visual Shader source partition](#topic-visual-shader-module-build)
- [Visual shader nodes](#topic-visual-shader-nodes)
- [Geometry resources](#topic-geometry-resources)
- [Spatial indexing and ray queries](#topic-spatial-indexing)

<a id="topic-brotli-stream-decompression"></a>
## 383. Brotli stream decompression

**Scope:** Vendored: brotli

**Builds on:** [Bitstream and Huffman decoding](#topic-prefix-code-decoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Its public decoding path is stateful and can return output incrementally.

The Brotli decoder consumes a guarded byte reader, decodes Huffman and prefix-coded streams, and exposes output from a decoder state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
const uint8_t* BrotliDecoderTakeOutput(BrotliDecoderState* s, size_t* size) {
  uint8_t* result = 0;
  size_t available_out = *size ? *size : 1u << 24;
  size_t requested_out = available_out;
  BrotliDecoderErrorCode status;
  if ((s->ringbuffer == 0) || ((int)s->error_code < 0)) {
    *size = 0;
    return 0;
  }
  WrapRingBuffer(s);
  status = WriteRingBuffer(s, &available_out, &result, 0, BROTLI_TRUE);
  /* Either WriteRingBuffer returns those "success" codes... */
  if (status == BROTLI_DECODER_SUCCESS ||
```

Source: `thirdparty/brotli/dec/decode.c` — BrotliDecoderTakeOutput (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

BrotliDecoderState

Evidence:
- Code: `thirdparty/brotli/dec/bit_reader.h:42` — BrotliBitReader
- Code: `thirdparty/brotli/dec/decode.c:2882` — BrotliDecoderTakeOutput
- Code: `thirdparty/brotli/dec/state.h:250` — BrotliDecoderStateStruct

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Image decode pipelines](#topic-image-decode-pipeline)
- [Bitstream and Huffman decoding](#topic-prefix-code-decoding)
- [Histograms and Huffman codes](#topic-histograms-and-huffman-codes)

<a id="topic-gpu-command-encoding"></a>
## 384. GPU command encoding

**Scope:** Vendored: metal-cpp

**Builds on:** [GPU resources and pipelines](#topic-gpu-resources-pipelines).

**This prepares you for:** [Metal drawable presentation](#topic-metal-presentation).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Command buffers create render, compute, blit, and resource-state encoders that record work against configured GPU resources.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void                            setErrorOptions(MTL::CommandBufferErrorOption errorOptions);

    void                            setLogState(const MTL::LogState* logState);

    void                            setRetainedReferences(bool retainedReferences);
};
class CommandBufferEncoderInfo : public NS::Referencing<CommandBufferEncoderInfo>
{
public:
    NS::Array*               debugSignposts() const;

    CommandEncoderErrorState errorState() const;
```

Source: `thirdparty/metal-cpp/Metal/MTLCommandBuffer.hpp` — MTL::CommandBuffer (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MTL::CommandBuffer, MTL::RenderCommandEncoder, MTL::ComputeCommandEncoder

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTLCommandBuffer.hpp:259` — MTL::CommandBuffer
- Code: `thirdparty/metal-cpp/Metal/MTLComputeCommandEncoder.hpp:131` — MTL::ComputeCommandEncoder

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Lossless pixel transform pipeline](#topic-lossless-transform-pipeline)
- [Lossy macroblock encoding](#topic-lossy-macroblock-encoding)
- [Resource-backed assets](#topic-resource-assets)
- [Resources and asset lifecycle](#topic-resources)

<a id="topic-metalfx-scaling"></a>
## 385. MetalFX scaling and interpolation

**Scope:** Vendored: metal-cpp

**Builds on:** [GPU resources and pipelines](#topic-gpu-resources-pipelines).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

MetalFX descriptors create spatial, temporal, temporal-denoised, and frame-interpolation effect instances from a Metal device.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
_MTLFX_INLINE MTLFX::FrameInterpolator* MTLFX::FrameInterpolatorDescriptor::newFrameInterpolator( const MTL::Device* device ) const
{
    return NS::Object::sendMessage< MTLFX::FrameInterpolator* >( this, _MTLFX_PRIVATE_SEL( newFrameInterpolatorWithDevice_ ), device );
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------------

_MTLFX_INLINE MTL4FX::FrameInterpolator* MTLFX::FrameInterpolatorDescriptor::newFrameInterpolator( const MTL::Device* device, const MTL4::Compiler* compiler ) const
{
    return NS::Object::sendMessage< MTL4FX::FrameInterpolator* >( this, _MTLFX_PRIVATE_SEL( newFrameInterpolatorWithDevice_compiler_ ), device, compiler );
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------------
```

Source: `thirdparty/metal-cpp/MetalFX/MTLFXFrameInterpolator.hpp` — MTLFX::FrameInterpolatorDescriptor::newFrameInterpolator (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

MTLFX::SpatialScaler, MTLFX::TemporalScaler, MTLFX::FrameInterpolator

Evidence:
- Code: `thirdparty/metal-cpp/MetalFX/MTLFXSpatialScaler.hpp:231` — MTLFX::SpatialScalerDescriptor::newSpatialScaler
- Code: `thirdparty/metal-cpp/MetalFX/MTLFXFrameInterpolator.hpp:325` — MTLFX::FrameInterpolatorDescriptor::newFrameInterpolator

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Device runtime](#topic-device-runtime)
- [Metal-cpp object bridge](#topic-metal-cpp-object-bridge)
- [Transform interpolation](#topic-transform-interpolation)

<a id="topic-motion-blur"></a>
## 386. Motion blur

**Scope:** Vendored: embree

**Builds on:** [BVH construction](#topic-bvh-construction), [Geometry resources](#topic-geometry-resources), [Ray query state](#topic-ray-query).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The implementation has linear-bound primitive references, motion-blur builders, time-range node forms, and time-segment geometry access.

Motion-blur Geometry and BVH nodes interpolate bounds at ray-query time.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/*! A primitive reference stores the bounds of the primitive and its ID. */
  struct PrimRefMB
  {
    typedef LBBox3fa BBox;

    __forceinline PrimRefMB () {}

    __forceinline PrimRefMB (const LBBox3fa& lbounds_i, unsigned int activeTimeSegments, BBox1f time_range, unsigned int totalTimeSegments, unsigned int geomID, unsigned int primID)
      : lbounds((LBBox3fx)lbounds_i), time_range(time_range)
    {
      assert(activeTimeSegments > 0);
      lbounds.bounds0.lower.a = geomID;
      lbounds.bounds0.upper.a = primID;
      lbounds.bounds1.lower.a = activeTimeSegments;
```

Source: `thirdparty/embree/kernels/builders/primref_mb.h` — PrimRefMB (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

PrimRefMB, AABBNodeMB_t, AABBNodeMB4D_t

Evidence:
- Code: `thirdparty/embree/kernels/builders/primref_mb.h:15` — PrimRefMB
- Code: `thirdparty/embree/kernels/bvh/bvh_node_aabb_mb.h:12` — AABBNodeMB_t
- Code: `thirdparty/embree/kernels/bvh/bvh_node_aabb_mb4d.h:12` — AABBNodeMB4D_t
- Code: `thirdparty/embree/kernels/common/default.h` — timeSegment

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Android Gradle assembly](#topic-android-gradle-assembly)
- [BVH split heuristics](#topic-builder-heuristics)
- [Constraint solving](#topic-constraints)
- [CPU-specialized DSP backends](#topic-cpu-specialized-dsp)
- [Motion-blur geometry](#topic-motion-blur-geometry)
- [Physics queries and motion tests](#topic-physics-queries)
- [2D and 3D physics queries](#topic-physics-server-queries)
- [3D physics query contracts](#topic-physics-space-queries)
- [Project catalog and selection](#topic-project-catalog)
- [Ray query state](#topic-ray-query)
- [Rigid-body motion](#topic-rigid-body-motion)
- [Main loop, OS, and time services](#topic-runtime-loop-time)
- [Compile-time platform backends](#topic-sdl-platform-backends)
- [Theora block video coding](#topic-theora-block-video-codec)
- [Wraparound network time](#topic-wraparound-network-time)
- [Primitive references](#topic-primitive-references)
- [Spatial indexing and ray queries](#topic-spatial-indexing)

<a id="topic-narrow-phase"></a>
## 387. Narrow-phase collision queries

**Scope:** Vendored: jolt_physics

**Builds on:** [Broad-phase collision detection](#topic-broad-phase), [Collision shapes](#topic-collision-shapes), [Convex support and penetration](#topic-convex-collision).

**This prepares you for:** [Contact manifolds and warm-starting](#topic-contact-management), [Vehicle dynamics](#topic-vehicle-dynamics).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Results are emitted through specialized collision collectors and dispatch handlers.

Narrow-phase collision queries test collision shapes for rays, points, overlaps, and casts after broad-phase candidates have been selected.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
// A collision collector that flips the collision results
	class ReversedCollector : public CollideShapeCollector
	{
	public:
		explicit				ReversedCollector(CollideShapeCollector &ioCollector) :
			CollideShapeCollector(ioCollector),
			mCollector(ioCollector)
		{
		}

		virtual void			AddHit(const CollideShapeResult &inResult) override
		{
			// Add the reversed hit
			mCollector.AddHit(inResult.Reversed());
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionDispatch.cpp` — ReversedCollector (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Shape, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/NarrowPhaseQuery.h:21` — NarrowPhaseQuery
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionCollector.h:44` — CollisionCollector
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionDispatch.cpp:37` — ReversedCollector

<a id="topic-rendering-device-resources"></a>
## 388. RenderingDevice descriptors and handles

**Scope:** First-party

**Builds on:** [Resource and server identifiers](#topic-resource-identifiers).

**This prepares you for:** [Renderer-owned resource storage](#topic-renderer-storage), [Scene render data and buffers](#topic-scene-data-and-buffers), [Shader programs and material parameters](#topic-shader-programs), [Viewport render-frame data](#topic-viewport-render-data).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

RD-prefixed descriptor objects carry configuration into the low-level rendering API.

RenderingDevice creates and consumes RID handles for buffers, textures, shaders, uniforms, pipelines, framebuffers, and acceleration structures.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```text
<?xml version="1.0" encoding="UTF-8" ?>
<class name="RDTextureFormat" inherits="RefCounted" api_type="core" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="../class.xsd">
	<brief_description>
		Texture format (used by [RenderingDevice]).
	</brief_description>
	<description>
		This object is used by [RenderingDevice].
	</description>
	<tutorials>
	</tutorials>
	<methods>
		<method name="add_shareable_format">
			<return type="void" />
			<param index="0" name="format" type="int" enum="RenderingDevice.DataFormat" />
```

Source: `doc/classes/RDTextureFormat.xml` — RDTextureFormat (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RDTextureFormat, RDUniform, RDAccelerationStructureGeometry

Evidence:
- Code: `doc/classes/RenderingDevice.xml:2` — RenderingDevice
- Code: `doc/classes/RDUniform.xml:2` — RDUniform
- Code: `doc/classes/RDTextureFormat.xml:2` — RDTextureFormat

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Texture resources and fallback placeholders](#topic-textures-and-placeholders)

<a id="topic-scene-commit"></a>
## 389. Scene construction and commit

**Scope:** Vendored: embree

**Builds on:** [BVH construction](#topic-bvh-construction), [Device runtime](#topic-device-runtime), [Geometry resources](#topic-geometry-resources).

**This prepares you for:** [Instancing](#topic-instancing).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The scene layer also checks whether attached geometry has changed before marking itself modified.

A Scene is created from a Device, retains Geometry instances, and commits a BVH used by query calls.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void Scene::checkIfModifiedAndSet ()
{
  if (isModified ()) return;

  auto geometryIsModified = [this](size_t geomID)->bool {
    return isGeometryModified(geomID);
  };

  if (parallel_any_of (size_t(0), geometries.size (), geometryIsModified)) {
    setModified ();
  }
}
```

Source: `thirdparty/embree/kernels/common/scene_verify.cpp` — Scene::checkIfModifiedAndSet (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RTCScene, Scene

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h:12` — RTCScene
- Code: `thirdparty/embree/kernels/common/scene.h:38` — Scene
- Code: `thirdparty/embree/kernels/common/scene_verify.cpp:11` — Scene::checkIfModifiedAndSet

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [COLLADA scene interchange](#topic-collada-import)
- [Editing selection history](#topic-editing-selection-history)
- [Editor authoring workspaces](#topic-editor-authoring-workspaces)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Editor plugin state and custom types](#topic-editor-plugin-state)
- [Editor scene sessions](#topic-editor-scene-sessions)
- [Engine bootstrap](#topic-engine-bootstrap)
- [Explicit DRM synchronization objects](#topic-explicit-drm-syncobj)
- [GUI controls and layout](#topic-gui-control-layout)
- [OpenXR composition layers](#topic-openxr-composition-layers)
- [OpenXR render models](#topic-openxr-render-models)
- [Post-import processing](#topic-post-import-processing)
- [Primitive references](#topic-primitive-references)
- [Project catalog and selection](#topic-project-catalog)
- [Property-list and scene-property helpers](#topic-property-introspection)
- [Raycast occlusion culling](#topic-raycast-occlusion-culling)
- [Replicated property synchronization](#topic-replicated-property-synchronization)
- [Resource-backed assets](#topic-resource-assets)
- [Resources and asset lifecycle](#topic-resources)
- [Runtime configuration](#topic-runtime-configuration)
- [2D and 3D scene authoring](#topic-scene-authoring)
- [Scene-aware test context](#topic-scene-contexts)
- [Scene render data and buffers](#topic-scene-data-and-buffers)
- [Scene graph nodes](#topic-scene-graph)
- [Scene importing](#topic-scene-importing)
- [Scene state inspection](#topic-scene-state)
- [Scene graph and lifecycle](#topic-scene-tree)
- [Scene tree and signal connections](#topic-scene-tree-and-signal-connections)
- [SceneTree execution](#topic-scene-tree-execution)
- [Screen-space and environment effects](#topic-screen-and-environment-effects)
- [3D physics nodes](#topic-three-dimensional-physics)
- [Tile authoring](#topic-tile-authoring)
- [2D physics nodes](#topic-two-dimensional-physics)
- [Undo and redo history](#topic-undo-redo-history)
- [Version-control integration](#topic-version-control-integration)
- [Spatial indexing and ray queries](#topic-spatial-indexing)

<a id="topic-transform-quantization-rate-distortion"></a>
## 390. Transform, quantization, and rate-distortion search

**Scope:** Vendored: libwebp

**Builds on:** [Lossy macroblock encoding](#topic-lossy-macroblock-encoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The code includes both direct quantization and trellis-based rate-distortion choices.

The macroblock path transforms residuals, quantizes them with VP8Matrix parameters, evaluates rate-distortion costs, and reconstructs blocks.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
// passing zillions of params.
typedef struct VP8Residual VP8Residual;
struct VP8Residual {
  int first;
  int last;
  const int16_t* coeffs;

  int coeff_type;
  ProbaArray*   prob;
  StatsArray*   stats;
  CostArrayPtr  costs;
};

void VP8InitResidual(int first, int coeff_type,
```

Source: `thirdparty/libwebp/src/enc/cost_enc.h` — struct VP8Residual (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

VP8Encoder

Evidence:
- Code: `thirdparty/libwebp/src/enc/quant_enc.c` — TrellisQuantizeBlock, VP8SetSegmentParams
- Code: `thirdparty/libwebp/src/enc/cost_enc.h:31` — struct VP8Residual
- Code: `thirdparty/libwebp/src/dsp/enc.c` — VP8FTransform, VP8EncQuantizeBlock

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Theora block video coding](#topic-theora-block-video-codec)

<a id="topic-renderer-storage"></a>
## 391. Renderer-owned resource storage

**Scope:** First-party

**Builds on:** [RenderingDevice descriptors and handles](#topic-rendering-device-resources).

### Start with the idea

Games reuse the same data—textures, scenes, scripts, and sounds—in many places. A resource gives that data an identity and a lifecycle so the engine can load, share, save, and release it consistently.

Storage isolates resource lifetime and backend-specific data from scene traversal and draw submission.

The RD storage services own backend representations of lights, materials, meshes, particles, textures, and GPU resources addressed by rendering IDs.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class LightStorage : public RendererLightStorage {
public:
	enum ShadowAtlastQuadrant : uint32_t {
		QUADRANT_SHIFT = 27,
		OMNI_LIGHT_FLAG = 1 << 26,
		SHADOW_INDEX_MASK = OMNI_LIGHT_FLAG - 1,
		SHADOW_INVALID = 0xFFFFFFFF
	};

private:
	static LightStorage *singleton;
	uint32_t max_cluster_elements = 512;
```

Source: `servers/rendering/renderer_rd/storage_rd/light_storage.h` — LightStorage (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `servers/rendering/renderer_rd/storage_rd/light_storage.h:47` — LightStorage
- Code: `servers/rendering/renderer_rd/storage_rd/material_storage.h:47` — MaterialStorage
- Code: `servers/rendering/renderer_rd/storage_rd/mesh_storage.h:44` — MeshStorage
- Code: `servers/rendering/renderer_rd/storage_rd/texture_storage.h:48` — TextureStorage

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [CPU and GPU particle systems](#topic-particle-systems)
- [Resource-backed assets](#topic-resource-assets)

<a id="topic-scene-data-and-buffers"></a>
## 392. Scene render data and buffers

**Scope:** First-party

**Builds on:** [RenderingDevice descriptors and handles](#topic-rendering-device-resources).

**This prepares you for:** [Screen-space and environment effects](#topic-screen-and-environment-effects).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The renderer receives a compact frame-oriented view rather than owning scene objects directly.

RenderDataRD gathers visible scene-instance pointers and RID lists for lights, probes, decals, lightmaps, and fog volumes, while RenderSceneBuffersRD supplies named GPU scene textures.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class RenderDataRD : public RenderData {
	GDCLASS(RenderDataRD, RenderData);

public:
	// Access methods to expose data externally
	virtual Ref<RenderSceneBuffers> get_render_scene_buffers() const override { return render_buffers; }
	virtual RenderSceneData *get_render_scene_data() const override { return scene_data; }

	virtual RID get_environment() const override { return environment; }
	virtual RID get_camera_attributes() const override { return camera_attributes; }

	// Members are publicly accessible within the render engine.
	Ref<RenderSceneBuffersRD> render_buffers;
```

Source: `servers/rendering/renderer_rd/storage_rd/render_data_rd.h` — RenderDataRD (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

RenderDataRD

Evidence:
- Code: `servers/rendering/renderer_rd/storage_rd/render_data_rd.h:38` — RenderDataRD
- Code: `servers/rendering/renderer_rd/storage_rd/render_scene_buffers_rd.h:66` — RenderSceneBuffersRD

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Instancing](#topic-instancing)
- [Resource and server identifiers](#topic-resource-identifiers)
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-android-plugin-signals"></a>
## 393. Android plugins and signals

**Scope:** First-party

**Builds on:** [Android JNI interoperation](#topic-android-jni-interop).

**This prepares you for:** [Android instrumented integration tests](#topic-android-instrumented-tests).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The test app registers plugin classes in its manifest and GDScript tests retrieve and connect to plugin signals.

Android plugins are discovered from application metadata and expose annotated methods and signals through JNI interoperation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```java
*/
public final class GodotPluginRegistry {
	private static final String TAG = GodotPluginRegistry.class.getSimpleName();

	/**
	 * Prefix used for version 1 of the Godot plugin, mostly compatible with Godot 3.x
	 */
	private static final String GODOT_PLUGIN_V1_NAME_PREFIX = "org.godotengine.plugin.v1.";
	/**
	 * Prefix used for version 2 of the Godot plugin, compatible with Godot 4.2+
	 */
	private static final String GODOT_PLUGIN_V2_NAME_PREFIX = "org.godotengine.plugin.v2.";

	private static GodotPluginRegistry instance;
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/plugin/GodotPluginRegistry.java` — GodotPluginRegistry (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

SignalInfo

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/plugin/GodotPluginRegistry.java:2` — GodotPluginRegistry
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/plugin/SignalInfo.java:2` — SignalInfo
- Code: `platform/android/java/app/src/instrumented/AndroidManifest.xml:9` — org.godotengine.plugin.v2.SignalTestPlugin
- Code: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd:18` — test_signal_connection

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Class metadata and reflection](#topic-reflection)
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-android-storage-bridge"></a>
## 394. Android storage bridge

**Scope:** First-party

**Builds on:** [Android JNI interoperation](#topic-android-jni-interop).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Java/Kotlin implementation has separate access types rather than a single Android storage path.

Android file and directory access handlers implement asset, filesystem, MediaStore, and SAF paths through JNI interoperation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```kotlin
*/
class DirectoryAccessHandler(context: Context) {

	companion object {
		private val TAG = DirectoryAccessHandler::class.java.simpleName

		internal const val INVALID_DIR_ID = -1
		internal const val STARTING_DIR_ID = 1
	}

	private enum class AccessType(val nativeValue: Int) {
		ACCESS_RESOURCES(0),

		/**
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/directory/DirectoryAccessHandler.kt` — DirectoryAccessHandler (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/file/FileAccessHandler.kt:2` — FileAccessHandler
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/file/MediaStoreData.kt:2` — MediaStoreData
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/file/SAFData.kt:2` — SAFData
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/directory/DirectoryAccessHandler.kt:2` — DirectoryAccessHandler

<a id="topic-apple-embedded-packaging"></a>
## 395. Apple embedded packaging and signing

**Scope:** First-party

**Builds on:** [Target-platform export](#topic-target-platform-export).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Apple exporter extends the common platform exporter and uses dedicated signing and plugin-configuration types.

Apple embedded packaging adds Xcode project data, assets, Apple plugin configuration, and code-signing structures to a target-platform export.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class CodeSignCodeResources {
public:
	enum class CRMatch {
		CR_MATCH_NO,
		CR_MATCH_YES,
		CR_MATCH_NESTED,
		CR_MATCH_OPTIONAL,
	};

private:
	struct CRFile {
		String name;
		String hash;
```

Source: `editor/export/codesign.h` — CodeSign (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

EditorExportPlatform, EditorExportPlatform::ExportMessage

Evidence:
- Code: `editor/export/editor_export_platform_apple_embedded.h:60` — EditorExportPlatformAppleEmbedded
- Code: `editor/export/plugin_config_apple_embedded.h:53` — PluginConfigAppleEmbedded
- Code: `editor/export/codesign.h:354` — CodeSign

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Apple embedded hosting](#topic-apple-embedded-hosting)
- [Editor and project settings](#topic-editor-and-project-settings)
- [Editor plug-in extension](#topic-editor-plugin-extension)
- [Editor plugin lifecycle](#topic-editor-plugin-lifecycle)
- [Encoder configuration](#topic-encoder-configuration)

<a id="topic-contact-management"></a>
## 396. Contact manifolds and warm-starting

**Scope:** Vendored: jolt_physics

**Builds on:** [Narrow-phase collision queries](#topic-narrow-phase).

**This prepares you for:** [Constraint solving](#topic-constraints).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The contact manager owns the cache used to retain collision-solver information.

Contact management converts narrow-phase collision results into contact constraints and caches manifolds, body pairs, and contact points between updates.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
const ContactConstraintManager::MKeyValue *ContactConstraintManager::ManifoldCache::Find(const SubShapeIDPair &inKey, uint64 inKeyHash) const
{
	JPH_ASSERT(mIsFinalized);
	return mCachedManifolds.Find(inKey, inKeyHash);
}

ContactConstraintManager::MKeyValue *ContactConstraintManager::ManifoldCache::Create(ContactAllocator &ioContactAllocator, const SubShapeIDPair &inKey, uint64 inKeyHash, int inNumContactPoints)
{
	JPH_ASSERT(!mIsFinalized);
	MKeyValue *kv = mCachedManifolds.Create(ioContactAllocator, inKey, inKeyHash, CachedManifold::sGetRequiredExtraSize(inNumContactPoints));
	if (kv == nullptr)
	{
		ioContactAllocator.mErrors |= EPhysicsUpdateError::ManifoldCacheFull;
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Constraints/ContactConstraintManager.cpp` — ContactConstraintManager::ManifoldCache::Find (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CachedManifold, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/ContactConstraintManager.h` — CachedContactPoint, CachedManifold, CachedBodyPair, ManifoldCache
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/ContactConstraintManager.cpp:333` — ContactConstraintManager::ManifoldCache::Find

<a id="topic-instancing"></a>
## 397. Instancing

**Scope:** Vendored: embree

**Builds on:** [Geometry resources](#topic-geometry-resources), [Scene construction and commit](#topic-scene-commit).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The configured API sets RTC_MAX_INSTANCE_LEVEL_COUNT to one.

Instance and InstanceArray are Geometry types attached to a Scene that contribute transformed primitive references and maintain instance-query context state.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
#define RTC_MAX_INSTANCE_LEVEL_COUNT 1

// #cmakedefine EMBREE_GEOMETRY_INSTANCE_ARRAY
#if defined(EMBREE_GEOMETRY_INSTANCE_ARRAY)
  #define RTC_GEOMETRY_INSTANCE_ARRAY
#endif

// #cmakedefine01 EMBREE_SYCL_GEOMETRY_CALLBACK

#define EMBREE_MIN_WIDTH 0
#define RTC_MIN_WIDTH EMBREE_MIN_WIDTH

#if !defined(EMBREE_STATIC_LIB)
```

Source: `thirdparty/embree/include/embree4/rtcore_config.h` — RTC_MAX_INSTANCE_LEVEL_COUNT (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Instance, InstanceArray, InstancePrimitive

Evidence:
- Code: `thirdparty/embree/kernels/common/scene_instance.h:14` — Instance
- Code: `thirdparty/embree/kernels/common/scene_instance_array.h:14` — InstanceArray
- Code: `thirdparty/embree/kernels/common/instance_stack.h` — instance stack
- Code: `thirdparty/embree/include/embree4/rtcore_config.h:16` — RTC_MAX_INSTANCE_LEVEL_COUNT

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Geometry resources](#topic-geometry-resources)
- [glTF material and texture conversion](#topic-gltf-materials-textures)
- [Intersection results](#topic-hit-results)
- [OpenXR runtime loading](#topic-openxr-runtime-loading)
- [Meshes, materials, textures, and instancing](#topic-rendering-assets)
- [Rigid-body motion](#topic-rigid-body-motion)
- [Scene render data and buffers](#topic-scene-data-and-buffers)
- [Theora block video coding](#topic-theora-block-video-codec)
- [Vulkan instance setup](#topic-vulkan-instance-setup)
- [Primitive references](#topic-primitive-references)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-metal-presentation"></a>
## 398. Metal drawable presentation

**Scope:** Vendored: metal-cpp

**Builds on:** [GPU command encoding](#topic-gpu-command-encoding).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

QuartzCore Metal layers produce drawables that implement the Metal drawable interface for presentation.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
_CA_INLINE CA::MetalLayer* CA::MetalDrawable::layer() const
{
    return Object::sendMessage<MetalLayer*>(this, _CA_PRIVATE_SEL(layer));
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------------

_CA_INLINE MTL::Texture* CA::MetalDrawable::texture() const
{
    return Object::sendMessage<MTL::Texture*>(this, _CA_PRIVATE_SEL(texture));
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------------
```

Source: `thirdparty/metal-cpp/QuartzCore/CAMetalDrawable.hpp` — CA::MetalDrawable (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

CA::MetalLayer, CA::MetalDrawable

Evidence:
- Code: `thirdparty/metal-cpp/QuartzCore/CAMetalLayer.hpp:145` — CA::MetalLayer::nextDrawable
- Code: `thirdparty/metal-cpp/QuartzCore/CAMetalDrawable.hpp:45` — CA::MetalDrawable

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Metal-cpp object bridge](#topic-metal-cpp-object-bridge)
- [Shader interface metadata](#topic-shader-interface-metadata)
- [Swapchain presentation](#topic-vulkan-swapchain-presentation)

<a id="topic-screen-and-environment-effects"></a>
## 399. Screen-space and environment effects

**Scope:** First-party

**Builds on:** [Scene render data and buffers](#topic-scene-data-and-buffers).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

These classes are effect-oriented services layered over scene buffers and shader pipelines.

RD effects and environment services allocate and process scene buffers for luminance, temporal anti-aliasing, reflections, ambient effects, fog, global illumination, sky, tone mapping, and scaling.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
class SSEffects {
private:
	static SSEffects *singleton;

public:
	static SSEffects *get_singleton() { return singleton; }

	SSEffects();
	~SSEffects();

	/* Last Frame */

	void allocate_last_frame_buffer(Ref<RenderSceneBuffersRD> p_render_buffers, bool p_use_ssil, bool p_use_ssr);
```

Source: `servers/rendering/renderer_rd/effects/ss_effects.h` — SSEffects (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `servers/rendering/renderer_rd/effects/luminance.h` — Luminance, LuminanceBuffers
- Code: `servers/rendering/renderer_rd/effects/taa.h:38` — TAA
- Code: `servers/rendering/renderer_rd/effects/ss_effects.h:76` — SSEffects
- Code: `servers/rendering/renderer_rd/environment/fog.h` — Fog, VolumetricFog
- Code: `servers/rendering/renderer_rd/environment/gi.h` — GI, SDFGI
- Code: `servers/rendering/renderer_rd/environment/sky.h:49` — SkyRD

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Scene construction and commit](#topic-scene-commit)
- [Scene-aware test context](#topic-scene-contexts)

<a id="topic-android-instrumented-tests"></a>
## 400. Android instrumented integration tests

**Scope:** First-party

**Builds on:** [Android plugins and signals](#topic-android-plugin-signals).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

The Android test source includes Kotlin test code, Java/Kotlin plugin fixtures, and a Godot project with GDScript tests.

Android instrumentation tests launch an app project that exercises plugins, signals, storage, and JavaClassWrapper conversions.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```kotlin
@RunWith(AndroidJUnit4::class)
class GodotAppTest {

	companion object {
		private val TAG = GodotAppTest::class.java.simpleName

		private const val GODOT_APP_LAUNCHER_CLASS_NAME = "com.godot.game.GodotAppLauncher"
		private const val GODOT_APP_CLASS_NAME = "com.godot.game.GodotApp"

		private val TEST_COMMAND_LINE_PARAMS = arrayOf("This is a test")
	}

	private fun getTestPlugin(): GodotAppInstrumentedTestPlugin? {
		return GodotPluginRegistry.getPluginRegistry()
```

Source: `platform/android/java/app/src/androidTestInstrumented/java/com/godot/game/GodotAppTest.kt` — GodotAppTest (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

Evidence:
- Code: `platform/android/java/app/src/androidTestInstrumented/java/com/godot/game/GodotAppTest.kt:2` — GodotAppTest
- Code: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd:19` — test_signal_emission
- Code: `platform/android/java/app/src/instrumented/assets/test/file_access/file_access_tests.gd:1` — FileAccessTests
- Code: `platform/android/java/app/src/instrumented/assets/test/javaclasswrapper/java_class_wrapper_tests.gd:1` — JavaClassWrapperTests

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Signal awaitability](#topic-signal-awaitability)

<a id="topic-constraints"></a>
## 401. Constraint solving

**Scope:** Vendored: jolt_physics

**Builds on:** [Contact manifolds and warm-starting](#topic-contact-management), [Rigid-body motion](#topic-rigid-body-motion).

**This prepares you for:** [Skeletons, animation, and ragdolls](#topic-skeletal-ragdoll), [State recording and validation](#topic-state-recording), [Vehicle dynamics](#topic-vehicle-dynamics).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

ConstraintSettings are serializable configuration objects; runtime Constraint objects participate in solver islands.

Constraint solving applies impulses to rigid-body motion and reuses contact manifolds while supporting point, distance, hinge, slider, fixed, gear, pulley, cone, six-degree-of-freedom, path, and swing-twist constraints.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void TwoBodyConstraint::BuildIslands(uint32 inConstraintIndex, IslandBuilder &ioBuilder, BodyManager &inBodyManager)
{
#ifdef JPH_ENABLE_ASSERTS
	// Validates that a body that is sleeping has zero velocity.
	mBody1->ValidateMotion();
	mBody2->ValidateMotion();
#endif

	bool body1_dynamic = mBody1->IsDynamic();
	bool body2_dynamic = mBody2->IsDynamic();

	// Activate bodies
	BodyID body_ids[2];
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Constraints/TwoBodyConstraint.cpp` — TwoBodyConstraint::BuildIslands (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

ConstraintSettings, TwoBodyConstraint

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/Constraint.h` — ConstraintSettings and Constraint
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/TwoBodyConstraint.cpp:23` — TwoBodyConstraint::BuildIslands
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/SixDOFConstraint.h` — SixDOFConstraintSettings and SixDOFConstraint

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Motion blur](#topic-motion-blur)

<a id="topic-skeletal-ragdoll"></a>
## 402. Skeletons, animation, and ragdolls

**Scope:** Vendored: jolt_physics

**Builds on:** [Constraint solving](#topic-constraints).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

SkeletonPose, SkeletalAnimation, SkeletonMapper, and Ragdoll form the animation-to-physics integration surface.

Skeletons retain named joints and parent relationships, animations produce joint states, and ragdolls connect those joints to Body parts and two-body constraints.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// Resource for a skinned animation
class JPH_EXPORT SkeletalAnimation : public RefTarget<SkeletalAnimation>
{
	JPH_DECLARE_SERIALIZABLE_NON_VIRTUAL(JPH_EXPORT, SkeletalAnimation)

public:
	/// Contains the current state of a joint, a local space transformation relative to its parent joint
	class JointState
	{
		JPH_DECLARE_SERIALIZABLE_NON_VIRTUAL(JPH_EXPORT, JointState)

	public:
		/// Convert from a local space matrix
		void							FromMatrix(Mat44Arg inMatrix);
```

Source: `thirdparty/jolt_physics/Jolt/Skeleton/SkeletalAnimation.h` — SkeletalAnimation (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Skeleton, Joint, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Skeleton/Skeleton.cpp` — Skeleton::CalculateParentJointIndices and AreJointsCorrectlyOrdered
- Code: `thirdparty/jolt_physics/Jolt/Skeleton/SkeletalAnimation.h:18` — SkeletalAnimation
- Code: `thirdparty/jolt_physics/Jolt/Physics/Ragdoll/Ragdoll.h` — RagdollSettings and Ragdoll

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [3D gizmo authoring](#topic-3d-gizmo-authoring)
- [glTF scene-node hierarchy](#topic-gltf-node-hierarchy)
- [Post-import processing](#topic-post-import-processing)
- [Skeletal modifiers and inverse kinematics](#topic-skeletal-ik)
- [2D skeleton modification stacks](#topic-skeleton-modification)
- [Skeleton modification and physical-bone simulation](#topic-skeleton-modifiers)

<a id="topic-state-recording"></a>
## 403. State recording and validation

**Scope:** Vendored: jolt_physics

**Builds on:** [Constraint solving](#topic-constraints), [RTTI-based serialization](#topic-serialization).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

StateRecorderImpl supports rewinding, clearing, reading, writing, and validation diagnostics.

State recording saves and validates Body and Constraint state through a stream that can compare replayed bytes with current values.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```c
/// A simple class that logs the state of the simulation. The resulting text file can be used to diff between platforms and find issues in determinism.
class DeterminismLog
{
private:
	JPH_INLINE uint32		Convert(float inValue) const
	{
		return *(uint32 *)&inValue;
	}

	JPH_INLINE uint64		Convert(double inValue) const
	{
		return *(uint64 *)&inValue;
	}
```

Source: `thirdparty/jolt_physics/Jolt/Physics/DeterminismLog.h` — DeterminismLog (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body, TwoBodyConstraint

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/StateRecorder.h:77` — StateRecorder
- Code: `thirdparty/jolt_physics/Jolt/Physics/StateRecorderImpl.cpp` — StateRecorderImpl::ReadBytes, WriteBytes, Rewind, Clear
- Code: `thirdparty/jolt_physics/Jolt/Physics/DeterminismLog.h:21` — DeterminismLog

### Related topics

These links come from shared vocabulary only (synthesized), so they suggest related reading rather than a prerequisite order.
- [Ray packets and streams](#topic-packet-queries)

<a id="topic-vehicle-dynamics"></a>
## 404. Vehicle dynamics

**Scope:** Vendored: jolt_physics

**Builds on:** [Constraint solving](#topic-constraints), [Narrow-phase collision queries](#topic-narrow-phase), [Rigid-body motion](#topic-rigid-body-motion).

### Start with the idea

Start with the practical job this code performs. Keep the names in the source as labels for that job; the example below shows where the behavior becomes concrete.

Wheeled, motorcycle, and tracked controllers are separate implementations over the vehicle constraint.

Vehicle dynamics couples a VehicleConstraint with a rigid Body, wheel collision tests, suspension, controllers, engine, transmission, differentials, tracks, and anti-roll bars.

### How the implementation applies it

The source example below is the concrete entry point. Read it once for the data it receives, then again for the branch, call, or stored result that changes the program state.

### Walk through a small source example

```cpp
void VehicleTransmissionSettings::SaveBinaryState(StreamOut &inStream) const
{
	inStream.Write(mMode);
	inStream.Write(mGearRatios);
	inStream.Write(mReverseGearRatios);
	inStream.Write(mSwitchTime);
	inStream.Write(mClutchReleaseTime);
	inStream.Write(mSwitchLatency);
	inStream.Write(mShiftUpRPM);
	inStream.Write(mShiftDownRPM);
	inStream.Write(mClutchStrength);
}
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleTransmission.cpp` — VehicleTransmissionSettings::SaveBinaryState (excerpt selected locally from grounded evidence)

**What to notice:** This is real code from the cited location. Read it in three passes: identify the input or stored state, find the decision or operation that changes it, then connect that result to the purpose described at the start of the section.

### Concrete structures to recognize

Body, VehicleTransmissionSettings

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleConstraint.h:65` — VehicleConstraint
- Code: `thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleCollisionTester.h` — VehicleCollisionTester, VehicleCollisionTesterRay, VehicleCollisionTesterCastSphere, VehicleCollisionTesterCastCylinder
- Code: `thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleTransmission.cpp:25` — VehicleTransmissionSettings::SaveBinaryState
<!-- rope-ladder:end document -->
