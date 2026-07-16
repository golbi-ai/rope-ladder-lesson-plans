<!-- rope-ladder:begin document 03545a73bf27538698effd26c1c24a0dfdce3c3edd3e8734f5daa08bd265ee7c -->
# Lesson plan

Each unit stays within one prerequisite depth and is sized to keep a focused teaching-assistant reflection.

## Unit 1 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-dynamic-variant-resources-cpp-templates-traits`
Display slug: `unit-001-dynamic-runtime-values`

### Domain: Dynamic runtime values — First-party

This is the central interchange value across the object, container, parser, and binding code.

Variant is the tagged runtime-value representation used for construction, conversion, operators, calls, indexing, and member access.

Read more: [Dynamic runtime values](references/domain-concepts.md#topic-dynamic-variant)

Use these entity examples: Variant.

### Domain: Resources and asset lifecycle — First-party

Resources are shared data objects used by scenes, scripts, meshes, shapes, materials, and extensions.

A Resource is a RefCounted asset value with a path, name, scene-local option, and loader or saver lifecycle.

Read more: [Resources and asset lifecycle](references/domain-concepts.md#topic-resources)

Use these entity examples: Resource, Mesh, Material, GDExtension.

### Implementation: C++ templates and traits — First-party

Generic core containers and binding helpers are primarily header-defined template code.

The implementation uses C++ templates and trait specializations to adapt behavior to types, including zero construction, hashing, tuple recursion, and argument conversion.

Read more: [C++ templates and traits](references/language-concepts.md#topic-cpp-templates-traits)

Apply it through: AABB.

## Unit 2 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-geometry-transforms-cpp-engine-polymorphism`
Display slug: `unit-002-geometry-and-transform-values`

### Domain: Geometry and transform values — First-party

These are the coordinate and bounds values shared by geometric algorithms and dynamic argument conversion.

The math layer defines 2D and 3D vectors, rectangles, boxes, planes, bases, quaternions, projections, and transforms as reusable value types.

Read more: [Geometry and transform values](references/domain-concepts.md#topic-geometry-transforms)

Use these entity examples: AABB, Transform3D.

### Implementation: C++ inheritance, virtual interfaces, and Ref ownership — First-party

Core abstractions such as resources, input events, and resource format loaders are C++ class hierarchies.

C++ classes use inheritance, virtual methods, Ref ownership, and static registries to implement core services.

Read more: [C++ inheritance, virtual interfaces, and Ref ownership](references/language-concepts.md#topic-cpp-engine-polymorphism)

Apply it through: Resource, InputEvent, GDExtension.

## Unit 3 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-object-identity-lifetime-cpp-native-classes`
Display slug: `unit-003-object-identity-and-lifetime`

### Domain: Object identity and lifetime — First-party

Object supplies engine object identity and ObjectDB support, while RefCounted and Ref provide reference-managed object use.

Read more: [Object identity and lifetime](references/domain-concepts.md#topic-object-identity-lifetime)

Use these entity examples: Object, ObjectID.

### Implementation: C++ native class hierarchy — First-party

The engine’s native implementation expresses its core object, scene, resource, extension, and scripting types as C++ classes, then exposes selected APIs through ClassDB.

Native components use C++ classes, inheritance, and virtual functions to implement engine types such as Resource, Node, and GDScript.

Read more: [C++ native class hierarchy](references/language-concepts.md#topic-cpp-native-classes)

Apply it through: Resource, Node, PackedScene, GDScript.

## Unit 4 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-string-names-paths-cpp-class-registration`
Display slug: `unit-004-strings-interned-names-and-node-paths`

### Domain: Strings, interned names, and node paths — First-party

The string layer implements Unicode strings, hash-backed interned StringName values, path and subpath storage in NodePath, and fuzzy-search helpers.

Read more: [Strings, interned names, and node paths](references/domain-concepts.md#topic-string-names-paths)

Use these entity examples: StringName, NodePath.

### Implementation: C++ runtime class registration — First-party

Godot's native API surface is assembled through ClassDB rather than by a separate hand-written runtime schema.

C++ templates and classes implement the runtime registry that records native class identity, inheritance, construction, methods, properties, signals, and constants.

Read more: [C++ runtime class registration](references/language-concepts.md#topic-cpp-class-registration)

Apply it through: Resource, RID.

## Unit 5 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-object-model-resource-loading-python-build-code-generation`
Display slug: `unit-005-engine-object-model`

### Domain: Engine object model — First-party

This is the base model that the scene, resource, script, and editor APIs build upon.

Object-derived types provide the common runtime identity and behavior base, while RefCounted supplies reference-counted lifetime for resource-like values.

Read more: [Engine object model](references/domain-concepts.md#topic-object-model)

Use these entity examples: Node, Resource, ClassInfo.

### Domain: Resource loading and caching — First-party

Loading is a registry-based service rather than a single file-format implementation.

ResourceLoader selects registered ResourceFormatLoader instances, applies cache modes, reports dependencies, and tracks threaded load tasks.

Read more: [Resource loading and caching](references/domain-concepts.md#topic-resource-loading)

Use these entity examples: Resource, ResourceLoader::ThreadLoadTask.

### Implementation: Python build code generation — First-party

This is build-time support rather than the engine's C++ runtime.

The Python build helper defines generation and build-entry functions for virtual-method source generation.

Read more: [Python build code generation](references/language-concepts.md#topic-python-build-code-generation)

## Unit 6 — Engine services

Shared subsystem: Engine services.

Lesson ID: `unit-dynamic-values-input-events-actions-cpp-class-inheritance`
Display slug: `unit-006-c-class-inheritance-for-backend-contracts`

### Domain: Dynamic values and dictionaries — First-party

It is the common value carrier used by reflection, scene serialization, scripting, and several networking APIs.

Variant is the cross-cutting tagged value representation for scalar, math, object, callable, signal, dictionary, array, and packed-array values.

Read more: [Dynamic values and dictionaries](references/domain-concepts.md#topic-dynamic-values)

Use these entity examples: Variant, JSON-RPC document.

### Domain: Input events and actions — First-party

The implementation separates concrete device events from named action bindings.

InputEvent resource subclasses represent key, mouse, joypad, touch, gesture, MIDI, shortcut, and action input; InputMap associates named actions with InputEvent lists and deadzones.

Read more: [Input events and actions](references/domain-concepts.md#topic-input-events-actions)

Use these entity examples: InputEvent, InputMap::Action.

### Implementation: C++ class inheritance for backend contracts — First-party

The code uses public derived classes to specialize a common engine-facing interface.

C++ class inheritance expresses shared renderer and server contracts and allows concrete dummy, extension, wrapper, clustered, and mobile implementations.

Read more: [C++ class inheritance for backend contracts](references/language-concepts.md#topic-cpp-class-inheritance)

## Unit 7 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-runtime-metadata-crypto-resources-cpp-templates-and-views`
Display slug: `unit-007-runtime-class-metadata`

### Domain: Runtime class metadata — First-party

This metadata is the bridge between native engine classes and the exposed API.

The runtime records classes, inheritance, methods, properties, signals, and instantiation capability for engine objects.

Read more: [Runtime class metadata](references/domain-concepts.md#topic-runtime-metadata)

Use these entity examples: RID, Resource.

### Domain: Cryptographic resources and contexts — First-party

The crypto area combines persisted key or certificate resources with operational contexts.

CryptoKey and X509Certificate are Resource types, while AES, hashing, HMAC, TLS options, and crypto operations are exposed through dedicated contexts and services.

Read more: [Cryptographic resources and contexts](references/domain-concepts.md#topic-crypto-resources)

Use these entity examples: CryptoKey, X509Certificate.

### Implementation: C++ templates and non-owning views — First-party

The rendering code combines engine containers with pointer-and-length style views for recorded GPU work.

Template-based containers and views expose typed data without copying it, including VectorView access to command and resource arrays.

Read more: [C++ templates and non-owning views](references/language-concepts.md#topic-cpp-templates-and-views)

Apply it through: RecordedCommand.

## Unit 8 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-debugger-transport-memory-allocation-glsl-compute-workgroups`
Display slug: `unit-008-debugger-capture-and-transport`

### Domain: Debugger capture and transport — First-party

Remote debugging is modeled as a service with message and capture state rather than as a platform adapter.

The debugger subsystem has engine captures and profilers, local and remote debugger implementations, and a socket-backed remote debugger peer.

Read more: [Debugger capture and transport](references/domain-concepts.md#topic-debugger-transport)

### Domain: Engine allocation helpers — First-party

Memory helpers define default allocators, typed allocators, allocation result wrappers, and global nil support.

Read more: [Engine allocation helpers](references/domain-concepts.md#topic-memory-allocation)

### Implementation: GLSL compute workgroups and synchronization — First-party

The supplied shaders use compute passes for velocity, luminance, resolve, roughness, hierarchy, skinning, and postprocessing work.

Compute shaders use declared resource interfaces, workgroup identifiers, bounds checks, shared arrays, and memory barriers for image-processing operations.

Read more: [GLSL compute workgroups and synchronization](references/language-concepts.md#topic-glsl-compute-workgroups)

## Unit 9 — Rendering & hardware

Shared subsystem: Rendering & hardware.

Lesson ID: `unit-gdextension-libraries-cpp-polymorphic-backends`
Display slug: `unit-009-c-polymorphic-backend-classes`

### Domain: GDExtension libraries — First-party

The extension runtime exposes a generated interface while retaining library and class lifecycle state in the engine.

GDExtension is a Resource that holds a loader, registered extension classes, initialization state, and library open, initialize, deinitialize, and close operations.

Read more: [GDExtension libraries](references/domain-concepts.md#topic-gdextension-libraries)

Use these entity examples: GDExtension.

### Implementation: C++ polymorphic backend classes — First-party

The repository uses public inheritance across audio, filesystem, networking, rendering-context, rendering-device, and editor plugin implementations.

C++ backend classes inherit engine interfaces so a rendering or platform service can be selected through a common base type.

Read more: [C++ polymorphic backend classes](references/language-concepts.md#topic-cpp-polymorphic-backends)

Apply it through: RenderingDeviceDriverVulkan::BufferInfo.

## Unit 10 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-input-midi-runtime-loop-time-cpp-object-hierarchies`
Display slug: `unit-010-c-class-inheritance`

### Domain: Keyboard and MIDI parsing — First-party

Keyboard utilities map key codes and names, while MIDIDriver::Parser interprets status and data bytes into MIDI messages.

Read more: [Keyboard and MIDI parsing](references/domain-concepts.md#topic-input-midi)

### Domain: Main loop, OS, and time services — First-party

OS hosts platform runtime services, MainLoop defines the loop-facing object type, and Time exposes date and time behavior.

Read more: [Main loop, OS, and time services](references/domain-concepts.md#topic-runtime-loop-time)

### Implementation: C++ class inheritance — First-party

The source uses base-class declarations to organize runtime, GUI, and asset types.

C++ class declarations define the repository's object families through public base classes, including Node-derived runtime objects, Resource-derived assets, and Viewport-derived windows.

Read more: [C++ class inheritance](references/language-concepts.md#topic-cpp-object-hierarchies)

Apply it through: Node, PackedScene.

## Unit 11 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-network-io-project-settings-cpp-const-borrowing`
Display slug: `unit-011-c-const-references-and-pointers`

### Domain: Networking and transport I/O — First-party

Protocol-facing abstractions are organized in core I/O, with Web-specific implementations for browser constraints.

The core exposes HTTP clients, stream peers, packet peers, TCP and UDP servers, DTLS and TLS peers, IP resolution, and Unix-domain socket support.

Read more: [Networking and transport I/O](references/domain-concepts.md#topic-network-io)

Use these entity examples: IPAddress.

### Domain: Project settings and feature overrides — First-party

This is the central configuration store for project-level state.

ProjectSettings stores named Variant values with persistence metadata, feature overrides, autoload definitions, resource paths, and a change version.

Read more: [Project settings and feature overrides](references/domain-concepts.md#topic-project-settings)

Use these entity examples: ProjectSettings, ProjectSettings::VariantContainer.

### Implementation: C++ const references and pointers — First-party

This is a pervasive access pattern in rendering, scene, GUI, and resource code.

C++ const references and pointers expose existing repository objects and buffers without copying them, such as shader parameter names, vectors, arrays, and geometry data.

Read more: [C++ const references and pointers](references/language-concepts.md#topic-cpp-const-borrowing)

Apply it through: Animation, GraphEditConnection.

## Unit 12 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-random-generation-synchronization-primitives-cpp-generic-containers`
Display slug: `unit-012-c-templates-and-typed-containers`

### Domain: Pseudo-random generation — First-party

RandomPCG supplies pseudo-random generation and RandomNumberGenerator exposes a reference-counted runtime wrapper.

Read more: [Pseudo-random generation](references/domain-concepts.md#topic-random-generation)

### Domain: Threads and synchronization — First-party

The OS layer defines Thread, mutexes, condition variables, read-write locks, semaphores, spin locks, and safe binary mutexes.

Read more: [Threads and synchronization](references/domain-concepts.md#topic-synchronization-primitives)

### Implementation: C++ templates and typed containers — First-party

Godot-specific Vector and Ref types are used as C++ template applications throughout this partition.

C++ template applications express typed containers and handles, including Vector collections of Ref-managed connections and parameter-pack-sized argument arrays.

Read more: [C++ templates and typed containers](references/language-concepts.md#topic-cpp-generic-containers)

Apply it through: GraphEditConnection, SceneState.

## Unit 13 — Startup & frame loop

Shared subsystem: Startup & frame loop.

Lesson ID: `unit-runtime-configuration-physics-space-queries-python-build-scripts`
Display slug: `unit-013-runtime-configuration`

### Domain: Runtime configuration — First-party

The main executable and a Godot project file both provide concrete configuration inputs.

Runtime configuration reads project settings such as the main scene and boot-splash options; the application project file also declares its main scene.

Read more: [Runtime configuration](references/domain-concepts.md#topic-runtime-configuration)

### Domain: 3D physics query contracts — First-party

The public query objects retain typed parameters while server-side direct-space operations consume them.

Physics-server types define ray, point, shape, and motion parameter and result records; RefCounted query objects expose those contracts to callers.

Read more: [3D physics query contracts](references/domain-concepts.md#topic-physics-space-queries)

Use these entity examples: PhysicsRayQueryParameters3D.

### Implementation: Python build scripts — First-party

The inventory contains Python builders beside SCsub files that import them.

Python build scripts import builder modules and define builder functions used by SCons-facing scene-theme scripts.

Read more: [Python build scripts](references/language-concepts.md#topic-python-build-scripts)

## Unit 14 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-xr-tracking-cpp-interface-polymorphism`
Display slug: `unit-014-c-interface-polymorphism`

### Domain: XR tracking and poses — First-party

Trackers classify sources and poses carry the spatial data exposed for those sources.

XRServer coordinates interfaces and tracker types, while positional, body, controller, face, and hand trackers publish XR pose state.

Read more: [XR tracking and poses](references/domain-concepts.md#topic-xr-tracking)

Use these entity examples: XRTracker, XRPose.

### Implementation: C++ interface polymorphism — First-party

The public extension classes mirror native polymorphic interfaces and distinguish required callbacks from optional ones.

C++ base classes and virtual callbacks define extension contracts whose subclasses provide physics, rendering, stream, and text implementations.

Read more: [C++ interface polymorphism](references/language-concepts.md#topic-cpp-interface-polymorphism)

Apply it through: RenderSceneBuffersConfiguration, RID.

## Unit 15 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-native-platform-adapters-png-image-codec-cpp-classes-inheritance`
Display slug: `unit-015-c-classes-and-inheritance`

### Domain: Native platform adapters — First-party

Unix and Windows have parallel implementations for several services, while Apple has additional embedded implementations.

Platform adapters implement filesystem, sockets, IP resolution, operating-system services, and threads behind engine interfaces selected by the driver build.

Read more: [Native platform adapters](references/domain-concepts.md#topic-native-platform-adapters)

### Domain: PNG image codec — First-party

PNG support is isolated as a driver with loader and saver classes.

The PNG driver supplies image loading and resource saving implementations and can build bundled libpng sources.

Read more: [PNG image codec](references/domain-concepts.md#topic-png-image-codec)

### Implementation: C++ classes and inheritance — First-party

Class inheritance provides the common contracts on which the editor’s services and UI controls are built.

The editor is structured from C++ classes that derive from engine base types such as Node, Container, ScrollContainer, ResourceImporter, and RefCounted.

Read more: [C++ classes and inheritance](references/language-concepts.md#topic-cpp-classes-inheritance)

## Unit 16 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-rendering-backends-scene-graph-cpp-classes-and-inheritance`
Display slug: `unit-016-scene-graph-nodes`

### Domain: Rendering backend abstraction — First-party

The source partition contains separate backend implementations rather than one graphics API implementation.

Rendering backends specialize common context and device-driver abstractions for Vulkan, GLES3, Direct3D 12, and Metal.

Read more: [Rendering backend abstraction](references/domain-concepts.md#topic-rendering-backends)

Use these entity examples: RenderingDeviceDriverVulkan::BufferInfo, RenderingDeviceDriverVulkan::CommandBufferInfo.

### Domain: Scene graph nodes — First-party

2D nodes inherit from CanvasItem through Node2D, while 3D nodes inherit from Node through Node3D.

Node2D and Node3D are scene graph node bases for specialized runtime behavior.

Read more: [Scene graph nodes](references/domain-concepts.md#topic-scene-graph)

Use these entity examples: Node2D, Node3D.

### Implementation: C++ classes and inheritance — First-party

Class declarations establish the editor subsystem types and their engine-facing base behavior.

The editor source declares engine services as C++ classes derived from Object, Resource, Control, and other engine base classes.

Read more: [C++ classes and inheritance](references/language-concepts.md#topic-cpp-classes-and-inheritance)

Apply it through: EditorSettings, VCS Diff File.

## Unit 17 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-canvas-and-viewports-animation-blending-cplusplus-enumerated-export-state`
Display slug: `unit-017-canvas-and-viewport-presentation`

### Domain: Canvas and viewport presentation — First-party

CanvasItem is the 2D visual branch; Window is a Viewport subclass.

CanvasItem, CanvasLayer, Viewport, and Window extend the Node tree with visual presentation and window-facing state.

Read more: [Canvas and viewport presentation](references/domain-concepts.md#topic-canvas-and-viewports)

Use these entity examples: Node.

### Domain: Animation blending and playback — First-party

Animation graph resources retain nodes, blend points, transitions, and node connections.

AnimationMixer, AnimationPlayer, AnimationTree, blend spaces, blend trees, and state machines compose animation playback.

Read more: [Animation blending and playback](references/domain-concepts.md#topic-animation-blending)

Use these entity examples: AnimationMixer, AnimationNodeBlendTree, AnimationNodeStateMachine.

### Implementation: C++ enumerated export state — First-party

Named integral categories keep export configuration and reporting states explicit at API boundaries.

C++ enumerations classify export messages, export filters, and file or script export modes.

Read more: [C++ enumerated export state](references/language-concepts.md#topic-cplusplus-enumerated-export-state)

Apply it through: EditorExportPreset, EditorExportPlatform::ExportMessage.

## Unit 18 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-property-introspection-scene-tree-execution-cplusplus-export-callbacks`
Display slug: `unit-018-c-export-callbacks`

### Domain: Property-list and scene-property helpers — First-party

These helpers bridge reflective property operations with scene state.

PropertyListHelper resolves slash-delimited property names to property records, while PropertyUtils compares properties and walks scene-state pack data.

Read more: [Property-list and scene-property helpers](references/domain-concepts.md#topic-property-introspection)

### Domain: SceneTree execution — First-party

This is the runtime coordinator above individual nodes.

A SceneTree schedules the Node hierarchy and maintains scene-level timer and tween processing collections.

Read more: [SceneTree execution](references/domain-concepts.md#topic-scene-tree-execution)

Use these entity examples: Node.

### Implementation: C++ export callbacks — First-party

EditorExportPlatform carries callback types so common packing code can delegate destination-specific work.

C++ function-pointer and callable declarations connect generic packaging code to save, removal, shared-object, and script callback implementations.

Read more: [C++ export callbacks](references/language-concepts.md#topic-cplusplus-export-callbacks)

Apply it through: EditorExportPlatform.

## Unit 19 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-physics-spaces-shader-programs-cplusplus-polymorphic-platform-adapters`
Display slug: `unit-019-physics-spaces-bodies-shapes-and-joints`

### Domain: Physics spaces, bodies, shapes, and joints — First-party

The implementation provides parallel 2D and 3D API families.

Each physics server models a self-contained space containing bodies, areas, joints, and shapes, and exposes APIs to create and configure them.

Read more: [Physics spaces, bodies, shapes, and joints](references/domain-concepts.md#topic-physics-spaces)

Use these entity examples: PhysicsMaterial, RID.

### Domain: Shader programs and material parameters — First-party

The high-level Shader API is distinct from RDShaderFile and RDShaderSPIRV, which are RenderingDevice-facing representations.

A Shader resource supplies custom shader source, ShaderInclude supplies reusable source fragments, and ShaderMaterial binds a Shader with parameter values.

Read more: [Shader programs and material parameters](references/domain-concepts.md#topic-shader-programs)

Use these entity examples: ShaderMaterial, RDPipelineSpecializationConstant.

### Implementation: C++ polymorphic platform adapters — First-party

The common export-platform base is extended by platform-specific exporters, including the Apple embedded exporter.

C++ class inheritance and virtual functions define the target-platform adapter contract.

Read more: [C++ polymorphic platform adapters](references/language-concepts.md#topic-cplusplus-polymorphic-platform-adapters)

Apply it through: EditorExportPlatform.

## Unit 20 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-stream-networking-csharp-attributes-reflection`
Display slug: `unit-020-c-attributes-and-reflection`

### Domain: Byte streams and socket servers — First-party

StreamPeerBuffer provides the same stream cursor model over an in-memory byte array.

StreamPeer exposes typed and raw byte-stream I/O, specialized peers implement TCP, TLS, compression, and Unix-domain sockets, and socket servers accept peer connections.

Read more: [Byte streams and socket servers](references/domain-concepts.md#topic-stream-networking)

Use these entity examples: StreamPeerBuffer, RID.

### Implementation: C# attributes and reflection — First-party

Godot declares attributes for scripts, exported members, signals, RPCs, tools, and serialization, then bridge code inspects attributes at runtime.

Attributes declare engine-facing metadata and reflection reads it to connect types, members, and script information.

Read more: [C# attributes and reflection](references/language-concepts.md#topic-csharp-attributes-reflection)

Apply it through: ScriptPathAttribute, AssemblyHasScriptsAttribute.

## Unit 21 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-class-reference-generation-cryptography-classes-inheritance-and-lookup`
Display slug: `unit-021-classes-inheritance-and-lookup`

### Domain: Class-reference generation — First-party

The Python tool owns the intermediate documentation model and output formatting.

Documentation generation parses class-reference XML into class, type, property, parameter, signal, method, and constant definitions before producing reStructuredText.

Read more: [Class-reference generation](references/domain-concepts.md#topic-class-reference-generation)

Use these entity examples: ClassDef, TypeName, ClassStatusProgress.

### Domain: Cryptographic keys, hashing, and TLS support — First-party

CryptoKey is a Resource, so key material participates in the engine resource lifecycle.

Crypto, CryptoKey, HashingContext, and HMACContext expose key generation, signatures, encryption, hashing, and message authentication APIs.

Read more: [Cryptographic keys, hashing, and TLS support](references/domain-concepts.md#topic-cryptography)

### Implementation: Classes, inheritance, and lookup — First-party

Fixtures test nested and external resolution separately from runtime inheritance behavior.

Classes supply nested types, inheritance, `super`, and lexical member lookup; preloaded classes also participate in lookup.

Read more: [Classes, inheritance, and lookup](references/language-concepts.md#topic-classes-inheritance-and-lookup)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 22 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-editor-extensibility-file-and-config-persistence-csharp-assembly-load-contexts`
Display slug: `unit-022-c-assemblyloadcontext-plugin-loading`

### Domain: Editor plugins and customization hooks — First-party

Editor-facing APIs are distinct from exported-project APIs and specialize controls, property editing, importing, and export behavior.

EditorPlugin subclasses register inspector, import, export, debugger, and resource-preview hooks in the editor.

Read more: [Editor plugins and customization hooks](references/domain-concepts.md#topic-editor-extensibility)

Use these entity examples: Resource.

### Domain: Files, directories, and configuration — First-party

This is direct application data persistence rather than the Resource import and loading path.

FileAccess and DirAccess read or write project and user paths, while ConfigFile stores section and key values.

Read more: [Files, directories, and configuration](references/domain-concepts.md#topic-file-and-config-persistence)

Use these entity examples: ConfigFile.

### Implementation: C# AssemblyLoadContext plugin loading — First-party

PluginLoadContext selects shared assemblies from the main context and resolves other dependencies through AssemblyDependencyResolver.

The editor plugin loader creates AssemblyLoadContext instances to resolve dependencies and support unloading.

Read more: [C# AssemblyLoadContext plugin loading](references/language-concepts.md#topic-csharp-assembly-load-contexts)

## Unit 23 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-navigation-csharp-partial-source-generation`
Display slug: `unit-023-c-partial-types-and-source-generator-tests`

### Domain: Navigation maps and path queries — First-party

Both 2D and 3D APIs expose equivalent navigation query objects and server services.

Navigation agents use server maps, regions, path-query parameters, and path-query results to follow a target position.

Read more: [Navigation maps and path queries](references/domain-concepts.md#topic-navigation)

Use these entity examples: NavigationPathQueryParameters2D, NavigationPathQueryResult2D.

### Implementation: C# partial types and source-generator tests — First-party

Partial declarations let generated files add members to a type declared in the sample or test source.

The .NET SDK test projects define partial C# types and verify generators that emit Godot-facing signal, property, method, serialization, and script-path code.

Read more: [C# partial types and source-generator tests](references/language-concepts.md#topic-csharp-partial-source-generation)

Apply it through: CSharpScript.

## Unit 24 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-skeleton-modifiers-cpp-binary-data-codecs`
Display slug: `unit-024-c-binary-data-codecs`

### Domain: Skeleton modification and physical-bone simulation — First-party

Physical-bone adapters transfer simulated transforms back to visual bones.

2D modifier stacks order Resource-based bone modifications, while SkeletonModifier3D nodes run after AnimationMixer playback and can implement IK, constraints, and skeleton physics.

Read more: [Skeleton modification and physical-bone simulation](references/domain-concepts.md#topic-skeleton-modifiers)

Use these entity examples: SkeletonModificationStack2D, SkeletonProfile.

### Implementation: C++ binary data codecs — First-party

The accessor implementation separates byte transport from conversion into numeric and vector values.

glTF decoding reads typed C++ containers of raw bytes through buffer-view offsets and strides, while encoding writes byte arrays into newly allocated buffer views.

Read more: [C++ binary data codecs](references/language-concepts.md#topic-cpp-binary-data-codecs)

Apply it through: GLTFBufferView, GLTFAccessor.

## Unit 25 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-text-and-localization-concurrency-callables-and-lambdas`
Display slug: `unit-025-callables-and-lambdas`

### Domain: Text backends and translation domains — First-party

TextLine and TextParagraph are TextServer abstractions for a single line or paragraph.

TextServerManager registers and selects TextServer implementations, while TranslationServer stores language data in named TranslationDomain collections.

Read more: [Text backends and translation domains](references/domain-concepts.md#topic-text-and-localization)

Use these entity examples: TranslationDomain, TextLine.

### Domain: Thread synchronization — First-party

The class documentation specifies lifetime constraints to avoid cleanup races and deadlocks.

Mutex synchronizes access to a critical section between Thread instances and is documented as a reentrant binary semaphore.

Read more: [Thread synchronization](references/domain-concepts.md#topic-concurrency)

### Implementation: Callables and lambdas — First-party

The tests use typed parameters and returns, direct `call()`, captures, binding, and callable introspection.

Callable values can reference functions, constructors, built-ins, and lambdas; lambda bodies capture their enclosing context.

Read more: [Callables and lambdas](references/language-concepts.md#topic-callables-and-lambdas)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 26 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-ui-theming-glsl-compute-shaders`
Display slug: `unit-026-glsl-compute-shaders`

### Domain: UI themes and style boxes — First-party

Panel, PanelContainer, PopupPanel, separators, and controls consume these style definitions.

Theme resources apply reusable colors, constants, fonts, icons, and StyleBoxes across Control and Window branches, while StyleBox defines a visual box treatment.

Read more: [UI themes and style boxes](references/domain-concepts.md#topic-ui-theming)

Use these entity examples: Theme, StyleBox.

### Implementation: GLSL compute shaders — First-party

These shaders are converted into C++ headers by the module build.

Betsy shader sources declare compute workgroup sizes, bind sampled and storage images, fetch texels by invocation ID, and store generated output.

Read more: [GLSL compute shaders](references/language-concepts.md#topic-glsl-compute-shaders)

Apply it through: Image.

## Unit 27 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-viewport-render-data-editor-scene-sessions-python-scons-configuration`
Display slug: `unit-027-editor-scene-sessions`

### Domain: Viewport render-frame data — First-party

These are internal rendering-server objects rather than general script-created scene objects.

RenderData and RenderSceneData exist for a viewport frame, while RenderSceneBuffersConfiguration configures buffers recreated when a viewport changes.

Read more: [Viewport render-frame data](references/domain-concepts.md#topic-viewport-render-data)

Use these entity examples: RenderSceneBuffersConfiguration, RenderData.

### Domain: Editor scene sessions — First-party

EditedScene is the durable editor-side record for an open scene rather than the scene tree itself.

EditorData represents each open editor scene with its root, project path, plugin state, selection, undo history ID, and timing/version metadata.

Read more: [Editor scene sessions](references/domain-concepts.md#topic-editor-scene-sessions)

Use these entity examples: EditedScene.

### Implementation: Python SCons configuration functions — First-party

Module configuration is executable build logic rather than declarative project metadata alone.

Python build files define can_build and configure functions and call SCons environment methods to select module sources and generated artifacts.

Read more: [Python SCons configuration functions](references/language-concepts.md#topic-python-scons-configuration)

## Unit 28 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-filesystem-project-index-scene-importing-types-inference-and-conversions`
Display slug: `unit-028-project-filesystem-index`

### Domain: Project filesystem index — First-party

The index separates directory topology from per-file metadata and supports rescanning.

The filesystem service scans project directories into a tree of file records with resource type, UID, import state, dependency list, and script-class metadata.

Read more: [Project filesystem index](references/domain-concepts.md#topic-filesystem-project-index)

Use these entity examples: EditorFileSystemDirectory, EditorFileSystemDirectory::FileInfo.

### Domain: Scene importing — First-party

This is the central pipeline for bringing 3D scene formats into the editor.

Scene importers expose extensions, options, flags, and an import operation that produces a Node-based scene.

Read more: [Scene importing](references/domain-concepts.md#topic-scene-importing)

Use these entity examples: ColladaParseState.

### Implementation: Types, inference, and conversions — First-party

This is the foundational type-system topic for the fixture corpus.

GDScript fixtures contrast explicit type hints, `:=` inference, `Variant` boundaries, null assignment, and conversion at typed assignments and returns.

Read more: [Types, inference, and conversions](references/language-concepts.md#topic-types-inference-and-conversions)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 29 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-scene-authoring-build-time-source-generation-cpp-const-reference-accessors`
Display slug: `unit-029-2d-and-3d-scene-authoring`

### Domain: 2D and 3D scene authoring — First-party

The partition implements distinct editing surfaces rather than a single dimension-agnostic viewport.

Scene authoring is split between CanvasItemEditor for 2D work and Node3DEditor with Node3DEditorViewport for 3D work.

Read more: [2D and 3D scene authoring](references/domain-concepts.md#topic-scene-authoring)

### Domain: Build-time source generation — First-party

These helpers convert build inputs into generated C++ declarations and byte data.

The Python build helpers generate C++ tables for documentation data, exporter registration, compressed documentation, and compressed translations.

Read more: [Build-time source generation](references/domain-concepts.md#topic-build-time-source-generation)

Use these entity examples: EditorTranslationList.

### Implementation: C++ const-reference accessors — First-party

Metadata and API accessors make read-only intent explicit in both parameter and return types.

C++ member functions expose lookup results through const pointers and accept immutable String references to avoid copying lookup inputs.

Read more: [C++ const-reference accessors](references/language-concepts.md#topic-cpp-const-reference-accessors)

Apply it through: OpenXRInteractionProfile.

## Unit 30 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-editor-and-project-settings-cpp-inheritance-interfaces`
Display slug: `unit-030-editor-and-project-settings`

### Domain: Editor and project settings — First-party

Editor preferences and project configuration are separate concerns in this partition.

EditorSettings persists editor-side setting containers and project metadata, while dedicated dialogs expose editor, project, autoload, action-map, and input-event configuration.

Read more: [Editor and project settings](references/domain-concepts.md#topic-editor-and-project-settings)

Use these entity examples: EditorSettings.

### Implementation: C++ inheritance interfaces — First-party

The partition consistently uses derived classes to plug implementations into engine-facing abstractions.

C++ classes specialize engine and OpenXR base interfaces through public inheritance.

Read more: [C++ inheritance interfaces](references/language-concepts.md#topic-cpp-inheritance-interfaces)

Apply it through: OpenXRStructureBase, ShapedTextDataAdvanced.

## Unit 31 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-editor-build-composition-cpp-resource-and-polymorphism`
Display slug: `unit-031-editor-build-composition`

### Domain: Editor build composition — First-party

Editor source aggregation is expressed as Python-based SCons configuration alongside the C++ implementation.

SCsub build scripts add C++ source files to the editor source set and invoke child SConscript files for nested editor subsystems.

Read more: [Editor build composition](references/domain-concepts.md#topic-editor-build-composition)

### Implementation: C++ resources and polymorphic engine objects — First-party

The implementation derives persisted configuration from Resource and runtime services from engine base classes such as MultiplayerAPI and NavigationServer2D or NavigationServer3D.

C++ engine modules define polymorphic Resource and server objects that own long-lived configuration and runtime state.

Read more: [C++ resources and polymorphic engine objects](references/language-concepts.md#topic-cpp-resource-and-polymorphism)

Apply it through: SceneReplicationConfig, OggPacketSequence, OpenXRActionMap.

## Unit 32 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-shader-editing-and-language-plugins-3d-gizmo-authoring-cpp-scope-locking`
Display slug: `unit-032-shader-editing-and-language-plugins`

### Domain: Shader editing and language plugins — First-party

Language selection is represented by EditorShaderLanguagePlugin instances rather than hard-coded solely in one text editor.

Shader editing uses a ShaderEditorPlugin, a text shader editor, syntax highlighting, shader-language plugins, creation dialogs, and shader-file editing.

Read more: [Shader editing and language plugins](references/domain-concepts.md#topic-shader-editing-and-language-plugins)

### Domain: 3D gizmo authoring — First-party

The editor distributes type-specific 3D visualization behavior across dedicated gizmo plug-ins.

Node3D editor gizmo plug-ins visualize and manipulate cameras, lights, meshes, occluders, particle emitters, physics objects, skeleton tools, and visibility volumes.

Read more: [3D gizmo authoring](references/domain-concepts.md#topic-3d-gizmo-authoring)

### Implementation: C++ scope-bound locking — First-party

The repeated local MutexLock declarations make the synchronization boundary visible at write sites.

The Jolt contact listener creates scope-bound C++ lock objects before mutating shared contact, manifold, and debug-contact collections.

Read more: [C++ scope-bound locking](references/language-concepts.md#topic-cpp-scope-locking)

Apply it through: JoltSpace3D.

## Unit 33 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-resource-previews-debug-adapter-protocol-cpp-templates-and-const-access`
Display slug: `unit-033-asynchronous-resource-previews`

### Domain: Asynchronous resource previews — First-party

Preview generation is a shared inspector-side service rather than a feature embedded in each resource editor.

Resource previews use generators selected by handles, queue source paths or in-memory resources, and cache generated full and small textures with metadata.

Read more: [Asynchronous resource previews](references/domain-concepts.md#topic-resource-previews)

Use these entity examples: PreviewRequest, ResourcePreviewCacheEntry.

### Domain: Debug Adapter Protocol integration — First-party

The protocol layer translates debugger and editor state into a language-agnostic wire format.

The editor's debug adapter serializes protocol content as JSON, manages protocol data types, and starts a local protocol server.

Read more: [Debug Adapter Protocol integration](references/domain-concepts.md#topic-debug-adapter-protocol)

Use these entity examples: DAP::Source, DAP::Breakpoint, DAP::Capabilities.

### Implementation: C++ templates and const access — First-party

The code uses Ref<T>, Vector<T>, LocalVector<T>, TypedArray<T>, and const references or pointers throughout module and scene implementations.

C++ templates and const-qualified values provide typed access to engine handles and collection views.

Read more: [C++ templates and const access](references/language-concepts.md#topic-cpp-templates-and-const-access)

Apply it through: VisualShaderGraph, NavigationAgent3D.

## Unit 34 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-editor-authoring-workspaces-editor-theming-cpp-templates-and-specialization`
Display slug: `unit-034-c-templates-and-specialization`

### Domain: Editor authoring workspaces — First-party

These are editor-facing authoring surfaces rather than runtime rendering backends.

The editor implements dock management, filesystem browsing, scene-tree editing, animation editing, asset-library browsing, audio editing, and debugger views as specialized controls and plugins.

Read more: [Editor authoring workspaces](references/domain-concepts.md#topic-editor-authoring-workspaces)

### Domain: Editor theming — First-party

Theme construction is centralized in EditorThemeManager while visual variants reside in separate builders.

Editor theming builds a Theme from ThemeConfiguration and settings, with distinct classic and modern builders plus font, icon, color-map, and scale support.

Read more: [Editor theming](references/domain-concepts.md#topic-editor-theming)

### Implementation: C++ templates and specialization — First-party

The reusable interpolation structure is specialized where quaternion behavior needs its own implementation.

The glTF importer uses C++ class templates for interpolation and supplies a quaternion specialization; accessor decoding also instantiates numeric decoding for concrete element types.

Read more: [C++ templates and specialization](references/language-concepts.md#topic-cpp-templates-and-specialization)

Apply it through: GLTFAnimation, GLTFAccessor.

## Unit 35 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-localization-and-template-generation-project-catalog-gdscript-typed-collections`
Display slug: `unit-035-gdscript-typed-collections-and-signature-compatibility`

### Domain: Localization and translation-template generation — First-party

Message context, singular text, plural text, comments, and source locations are explicitly handled during template generation.

Localization tooling parses translation inputs, exposes locale selection and preview, edits localization settings, and generates message maps and POT-style template output.

Read more: [Localization and translation-template generation](references/domain-concepts.md#topic-localization-and-template-generation)

Use these entity examples: Translation Message.

### Domain: Project catalog and selection — First-party

The project manager operates on structured records rather than treating projects as paths alone.

ProjectList models known projects with name, path, tags, main scene, version, unsupported features, last-edit time, and favorite state.

Read more: [Project catalog and selection](references/domain-concepts.md#topic-project-catalog)

Use these entity examples: ProjectCatalog, ProjectCatalogItem.

### Implementation: GDScript typed collections and signature compatibility — First-party

The analyzer test corpus specifies accepted and rejected typed containers, function variance, and override signatures.

This extends type declarations with typed Array and Dictionary element constraints, plus parameter and return compatibility checks for inherited functions.

Read more: [GDScript typed collections and signature compatibility](references/language-concepts.md#topic-gdscript-typed-collections)

Apply it through: GDScriptParser::Node.

## Unit 36 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-project-upgrade-property-inspection-iteration-protocols`
Display slug: `unit-036-iteration-protocols`

### Domain: Project source migration — First-party

Migration is source-aware instead of a blind string replacement.

ProjectConverter3To4 retains whether each source line is a comment while applying named rename, conversion, and check-only passes to project source.

Read more: [Project source migration](references/domain-concepts.md#topic-project-upgrade)

Use these entity examples: SourceLine.

### Domain: Reflective property inspection — First-party

Specialized controls are selected by property shape while inspector plug-ins can supply custom editors.

The inspector reads Object properties into EditorProperty controls and supports custom parsing, revert, copy, paste, keying, pinning, and favorites.

Read more: [Reflective property inspection](references/domain-concepts.md#topic-property-inspection)

### Implementation: Iteration protocols — First-party

Typed and `Variant` paths are exercised for built-in and user-defined iteration.

`for` loops cover ranges, collections, and custom objects that provide iterator hook functions.

Read more: [Iteration protocols](references/language-concepts.md#topic-iteration-protocols)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 37 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-resource-specific-editors-run-management-javascript-browser-ffi`
Display slug: `unit-037-javascript-browser-ffi`

### Domain: Resource-specific editing — First-party

The implementation chooses type-specific editing experiences for these resource families.

Dedicated editor controls and docks edit gradients, curves, materials, sprite frames, mesh libraries, textures, packed scenes, and resource preloaders.

Read more: [Resource-specific editing](references/domain-concepts.md#topic-resource-specific-editors)

### Domain: Running projects and instances — First-party

Run UI, process embedding, device selection, and multiple-instance arguments are implemented as related editor services.

The run subsystem starts editor launches, exposes native-device runs, supports configurable multiple instances, and hosts embedded game-view processes.

Read more: [Running projects and instances](references/domain-concepts.md#topic-run-management)

### Implementation: JavaScript browser FFI — First-party

The bridges use runtime allocation and parsing helpers, ID mapping, browser event callbacks, and Web API object methods.

JavaScript bridge libraries marshal browser objects, strings, buffers, and callbacks between web APIs and C++ platform code.

Read more: [JavaScript browser FFI](references/language-concepts.md#topic-javascript-browser-ffi)

Apply it through: WebRTCPeerConnection, WebRTCDataChannel, WebXRInterfaceJS.

## Unit 38 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-resource-importing-tile-authoring-pattern-matching`
Display slug: `unit-038-pattern-matching-and-guards`

### Domain: Standalone resource importing — First-party

This is separate from scene import and produces editor resource assets from individual files.

Resource-importer classes handle images, textures, texture atlases, SVG, fonts, WAV audio, CSV translations, shaders, and bitmaps.

Read more: [Standalone resource importing](references/domain-concepts.md#topic-resource-importing)

### Domain: Tile authoring — First-party

Tile editing is implemented as several coordinated editor tools rather than a single monolithic panel.

Tile tools edit atlas sources, per-tile data, terrain data, proxies, map layers, and scene-collection sources.

Read more: [Tile authoring](references/domain-concepts.md#topic-tile-authoring)

### Implementation: Pattern matching and guards — First-party

Parser and runtime fixtures separately cover structural matching syntax and guarded execution.

`match` supports literals, arrays, dictionaries, bindings, multiple patterns, wildcards, and guarded branches.

Read more: [Pattern matching and guards](references/language-concepts.md#topic-pattern-matching)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 39 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-version-control-integration-fixture-contracts-python-scons-module-configuration`
Display slug: `unit-039-fixture-contracts`

### Domain: Version-control integration — First-party

The interface separates VCS-provided data structures from the editor-facing plugin.

Version-control integration exchanges structured status, commit, file-diff, hunk, and line data through EditorVCSInterface and presents it through VersionControlEditorPlugin.

Read more: [Version-control integration](references/domain-concepts.md#topic-version-control-integration)

Use these entity examples: VCS Diff File, VCS Diff Hunk, VCS Diff Line.

### Domain: Fixture contracts — First-party

A source fixture provides the case; its paired stored artifact provides the expected observable result.

The suite stores GDScript source alongside expected outcome files and test assets, so behavior is expressed as a checked fixture contract.

Read more: [Fixture contracts](references/domain-concepts.md#topic-fixture-contracts)

Use these entity examples: Test Script Fixture, Expected Result Fixture.

### Implementation: Python SCons module configuration — First-party

These scripts are the build-time composition layer for the supplied modules.

Python SCons scripts declare module buildability, dependencies, cloned environments, and source-file collection.

Read more: [Python SCons module configuration](references/language-concepts.md#topic-python-scons-module-configuration)

## Unit 40 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-openxr-runtime-integration-navigation-map-composition-python-build-hooks`
Display slug: `unit-040-openxr-runtime-integration`

### Domain: OpenXR runtime integration — First-party

This is the integration boundary; later OpenXR topics specialize its action, extension, spatial, and rendering behavior.

The module places OpenXRInterface over OpenXRAPI so the engine can interact with an OpenXR runtime.

Read more: [OpenXR runtime integration](references/domain-concepts.md#topic-openxr-runtime-integration)

Use these entity examples: OpenXRActionMap, OpenXRFutureResult.

### Domain: Navigation map composition — First-party

Parallel 2D and 3D implementations expose map-owned region, link, agent, and obstacle collections and iteration snapshots.

Navigation maps collect regions, links, agents, and obstacles and build per-frame read iterations.

Read more: [Navigation map composition](references/domain-concepts.md#topic-navigation-map-composition)

### Implementation: Python build hooks — First-party

Top-level and module-local scripts define functions such as can_build and configure, while SCsub files add sources and web JavaScript libraries conditionally.

Python configuration scripts define module build hooks and configure build environments.

Read more: [Python build hooks](references/language-concepts.md#topic-python-build-hooks)

## Unit 41 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-scene-replication-configuration-webrtc-peer-connections-warnings-and-suppression`
Display slug: `unit-041-scene-replication-configuration`

### Domain: Scene replication configuration — First-party

Each configured property stores a spawn flag and a replication mode, and the configuration maintains derived lists for replication phases.

A saved SceneReplicationConfig separates NodePath properties into spawn, sync, and watch lists.

Read more: [Scene replication configuration](references/domain-concepts.md#topic-scene-replication-configuration)

Use these entity examples: SceneReplicationConfig, ReplicationProperty.

### Domain: WebRTC peer connections — First-party

The core class has extension and JavaScript-backed implementations.

A WebRTC peer connection exposes connection state, negotiation, ICE candidates, and data-channel creation to Godot.

Read more: [WebRTC peer connections](references/domain-concepts.md#topic-webrtc-peer-connections)

Use these entity examples: WebRTCPeerConnection.

### Implementation: Warnings and selective suppression — First-party

Warning text is itself captured as a tested output contract.

The analyzer reports unsafe, unused, shadowing, inference, conversion, and coroutine cases, while `@warning_ignore` selectively suppresses specified warnings.

Read more: [Warnings and selective suppression](references/language-concepts.md#topic-warnings-and-suppression)

Apply it through: Expected Result Fixture, Test Script Fixture.

## Unit 42 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-gdscript-static-analysis-javascript-web-runtime`
Display slug: `unit-042-javascript-browser-runtime-libraries`

### Domain: GDScript static analysis — First-party

The supplied analyzer tests exercise typed arrays and dictionaries, enum types, override signatures, casts, return types, and invalid assignments.

The analyzer consumes parsed script trees to resolve inheritance, infer and check data types, validate annotations, and diagnose invalid expressions.

Read more: [GDScript static analysis](references/domain-concepts.md#topic-gdscript-static-analysis)

Use these entity examples: GDScriptParser::Node, GDScript.

### Implementation: JavaScript browser runtime libraries — First-party

The Web platform places browser-facing implementation in JavaScript files while C++ classes call their exported bridge functions.

JavaScript libraries implement engine startup, preloading, runtime heap access, browser display and input callbacks, audio worklets, fetch, MIDI, and JavaScript object proxies.

Read more: [JavaScript browser runtime libraries](references/language-concepts.md#topic-javascript-web-runtime)

Apply it through: JavaScriptObjectImpl.

## Unit 43 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-script-registration-metadata-visual-shader-nodes-cpp-runtime-adapters`
Display slug: `unit-043-c-inheritance-and-reference-counted-adapters`

### Domain: Managed script registration metadata — First-party

The generator emits ScriptPathAttribute declarations and an AssemblyHasScriptsAttribute containing discovered script types.

C# script classes receive generated script-path and assembly script-type metadata for engine discovery.

Read more: [Managed script registration metadata](references/domain-concepts.md#topic-script-registration-metadata)

Use these entity examples: ScriptPathAttribute, AssemblyHasScriptsAttribute.

### Domain: Visual shader nodes — First-party

The module defines input, output, parameter, expression, varying, math, texture, and particle node families.

Node resources provide reusable shader operations within typed graphs of node records and connection records.

Read more: [Visual shader nodes](references/domain-concepts.md#topic-visual-shader-nodes)

Use these entity examples: VisualShaderNode.

### Implementation: C++ inheritance and reference-counted adapters — First-party

The Android wrapper types and display/OS classes use C++ classes and inheritance as the native adaptation mechanism.

C++ defines reference-counted Java adapter classes and platform subclasses that specialize common engine interfaces.

Read more: [C++ inheritance and reference-counted adapters](references/language-concepts.md#topic-cpp-runtime-adapters)

## Unit 44 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-websocket-peers-openxr-binding-modifiers-python-scons-build-hooks`
Display slug: `unit-044-websocket-peers`

### Domain: WebSocket peers — First-party

The browser build includes a JavaScript socket bridge, while native builds include a WSLay-based peer.

WebSocketPeer defines packet-oriented socket behavior and has native and browser-backed implementations.

Read more: [WebSocket peers](references/domain-concepts.md#topic-websocket-peers)

Use these entity examples: WebSocketPeer.

### Domain: Binding modifiers — First-party

The implementation has a base modifier plus profile-level and action-binding subclasses, including D-pad and analog-threshold modifiers.

Binding modifiers alter interaction-profile or per-action binding data before it is submitted to OpenXR.

Read more: [Binding modifiers](references/domain-concepts.md#topic-openxr-binding-modifiers)

Use these entity examples: OpenXRIPBinding.

### Implementation: Python SCons build hooks — First-party

The platform build layer uses Python functions called by SCons environments.

Python build scripts define platform detection, option and flag setup, configuration hooks, bundle generation, template assembly, and Emscripten helper functions.

Read more: [Python SCons build hooks](references/language-concepts.md#topic-python-scons-build-hooks)

## Unit 45 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-csharp-translation-extraction-dotnet-editor-build-integration-cpp-classes-and-ref-handles`
Display slug: `unit-045-c-translation-extraction`

### Domain: C# translation extraction — First-party

The parser builds parse options from project define constants and resolves invocation symbols through a semantic model.

The editor parses C# source into a Roslyn compilation and extracts constant translation-call arguments and nearby comments.

Read more: [C# translation extraction](references/domain-concepts.md#topic-csharp-translation-extraction)

### Domain: Editor build integration — First-party

The implementation includes project-file editing, a build logger that writes an issues CSV, build management, and problem views.

Editor tools generate .NET projects, invoke builds, and surface CSV diagnostics.

Read more: [Editor build integration](references/domain-concepts.md#topic-dotnet-editor-build-integration)

Use these entity examples: BuildDiagnostic.

### Implementation: C++ inheritance and engine reference handles — First-party

The supplied C++ code combines class hierarchies with engine-specific ownership wrappers in tests and mocks.

C++ test infrastructure uses public inheritance for test doubles and stores engine objects in Ref<T> handles created with memnew.

Read more: [C++ inheritance and engine reference handles](references/language-concepts.md#topic-cpp-classes-and-ref-handles)

Apply it through: Animation, AnimationTrack.

## Unit 46 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-fbx-gltf-scene-import-gdscript-editor-services-gdscript-language-server`
Display slug: `unit-046-fbx-scene-import-through-gltf-structures`

### Domain: FBX scene import through GLTF structures — First-party

The implementation maps FBX scene, node, mesh, material, texture, animation, skin, and light data into GLTF-oriented state.

FBXDocument and FBXState specialize GLTF document and state structures while importers expose UFBX and FBX2GLTF editor entry points.

Read more: [FBX scene import through GLTF structures](references/domain-concepts.md#topic-fbx-gltf-scene-import)

Use these entity examples: FBXDocument, FBXState.

### Domain: GDScript editor services — First-party

The services are built only for editor builds.

Editor tooling consumes GDScript parser output to generate documentation, color syntax, extract translations, and provide completion-related behavior.

Read more: [GDScript editor services](references/domain-concepts.md#topic-gdscript-editor-services)

Use these entity examples: GDScriptParser::Node, GDScriptTokenizer::Token.

### Domain: GDScript language-server support — First-party

LSP-shaped request and response structures are defined in godot_lsp.h and served by JSONRPC-based protocol classes.

The language server reuses parser-derived symbol data to manage documents, resolve symbols, publish diagnostics, provide links, and build workspace edits.

Read more: [GDScript language-server support](references/domain-concepts.md#topic-gdscript-language-server)

Use these entity examples: LSP::TextDocumentItem, LSP::DocumentSymbol, GDScriptWorkspace.

## Unit 47 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-image-format-loading-high-level-rpc-ide-messaging-protocol`
Display slug: `unit-047-hdr-jpeg-and-ktx-loading`

### Domain: HDR, JPEG, and KTX loading — First-party

Separate module adapters implement an HDR ImageFormatLoader, a libjpeg-turbo ImageFormatLoader, and a KTX ResourceFormatLoader.

Read more: [HDR, JPEG, and KTX loading](references/domain-concepts.md#topic-image-format-loading)

### Domain: High-level RPC routing — First-party

RPC configuration is cached per node and packet processing distinguishes remote-call, spawn, despawn, sync, and system commands.

SceneMultiplayer routes RPC calls, peer state, connection commands, and RPC configuration.

Read more: [High-level RPC routing](references/domain-concepts.md#topic-high-level-rpc)

### Domain: IDE messaging protocol — First-party

Peer, decoder, client, server, and CLI forwarding code share Message and MessageContent as the protocol envelope.

IDE clients and the editor exchange framed requests and responses after a handshake, with JSON bodies for typed requests.

Read more: [IDE messaging protocol](references/domain-concepts.md#topic-ide-messaging-protocol)

Use these entity examples: Message, GodotIdeMetadata.

## Unit 48 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-image-codec-registration-mp3-audio-resources-tls-crypto-backend`
Display slug: `unit-048-image-codec-integration`

### Domain: Image codec integration — First-party

The supplied modules cover BMP loading, DDS resource loading and saving, and several block-compression implementations.

Image codec modules provide loader, saver, compressor, and decompressor implementations against engine image and resource interfaces.

Read more: [Image codec integration](references/domain-concepts.md#topic-image-codec-registration)

Use these entity examples: Image, DDSFormatInfo.

### Domain: MP3 audio resources — First-party

AudioStreamMP3 and AudioStreamPlaybackMP3 form the runtime pair, while ResourceImporterMP3 handles editor import.

The MP3 module loads MP3 stream data, supplies resampled playback, and imports MP3 assets.

Read more: [MP3 audio resources](references/domain-concepts.md#topic-mp3-audio-resources)

### Domain: Mbed TLS-backed crypto and transport — First-party

The Mbed TLS module supplies Godot Crypto, certificate, TLS context, DTLS server, and TLS peer implementations.

Read more: [Mbed TLS-backed crypto and transport](references/domain-concepts.md#topic-tls-crypto-backend)

## Unit 49 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-navigation-mesh-construction-objectdb-snapshots-ogg-vorbis-audio-streams`
Display slug: `unit-049-navigation-mesh-construction`

### Domain: Navigation mesh construction — First-party

2D construction uses polygon geometry and 3D construction uses Recast-backed mesh generation when configured.

The 2D and 3D generators turn source geometry into navigation mesh data and connect region or link geometry into map iterations.

Read more: [Navigation mesh construction](references/domain-concepts.md#topic-navigation-mesh-construction)

### Domain: ObjectDB snapshots — First-party

The runtime collector transports snapshot data while editor classes turn it into inspectable data and controls.

The debug-only object database profiler collects serialized object snapshots and renders summary, class, node, object, and ref-counted views.

Read more: [ObjectDB snapshots](references/domain-concepts.md#topic-objectdb-snapshots)

### Domain: Ogg Vorbis audio streams — First-party

The module also provides an editor resource importer.

AudioStreamOggVorbis and its playback class provide an audio-stream resource and resampled playback implementation.

Read more: [Ogg Vorbis audio streams](references/domain-concepts.md#topic-ogg-vorbis-audio-streams)

Use these entity examples: AudioStreamOggVorbis, AudioStreamPlaybackOggVorbis.

## Unit 50 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-ogg-packet-sequences-openxr-action-configuration-procedural-noise-textures`
Display slug: `unit-050-ogg-packet-sequences`

### Domain: Ogg packet sequences — First-party

The module defines a Resource plus a ref-counted playback companion.

OggPacketSequence persists packet data, page granule positions, and sample-rate information for sequence playback.

Read more: [Ogg packet sequences](references/domain-concepts.md#topic-ogg-packet-sequences)

Use these entity examples: OggPacketSequence.

### Domain: OpenXR action configuration — First-party

The action-map subtree uses Resource-derived configuration objects for the OpenXR input graph.

OpenXRActionMap persists action sets, actions, interaction profiles, bindings, modifiers, and haptic feedback configuration.

Read more: [OpenXR action configuration](references/domain-concepts.md#topic-openxr-action-configuration)

Use these entity examples: OpenXRActionMap.

### Domain: Procedural noise textures — First-party

The module supplies an abstract Noise resource, FastNoiseLite implementation, texture resources, editor preview support, and tests.

Noise generators produce values and image data that NoiseTexture2D and NoiseTexture3D turn into textures.

Read more: [Procedural noise textures](references/domain-concepts.md#topic-procedural-noise-textures)

## Unit 51 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-raycast-occlusion-culling-regular-expression-results-signal-awaitability`
Display slug: `unit-051-raycast-occlusion-culling`

### Domain: Raycast occlusion culling — First-party

Embree-backed lightmap and static raycaster classes supplement the central occlusion-cull implementation.

The raycast module represents occluders, instances, scenarios, and raycast thread data to provide renderer-scene occlusion culling.

Read more: [Raycast occlusion culling](references/domain-concepts.md#topic-raycast-occlusion-culling)

Use these entity examples: RaycastOccluder.

### Domain: Regular-expression matching — First-party

RegExMatch is the cross-call result object for search operations rather than a local temporary.

The regex module exposes compiled regular expressions and RefCounted match results with ranges.

Read more: [Regular-expression matching](references/domain-concepts.md#topic-regular-expression-results)

Use these entity examples: RegExMatch.

### Domain: Signal awaitability — First-party

SignalAwaiter bridges signal completion back to managed await continuations.

A Signal combines an owner and a signal name and exposes a custom awaiter for asynchronous notification.

Read more: [Signal awaitability](references/domain-concepts.md#topic-signal-awaitability)

Use these entity examples: Variant, ManagedCallbacks.

## Unit 52 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-theora-video-streams-upnp-device-control-visual-shader-module-build`
Display slug: `unit-052-theora-video-streams`

### Domain: Theora video streams — First-party

The stream, loader, playback, and writer types separate runtime playback from editor export behavior.

The Theora module defines a video-stream resource, playback implementation, resource loader, and OGV movie writer.

Read more: [Theora video streams](references/domain-concepts.md#topic-theora-video-streams)

Use these entity examples: VideoStreamTheora.

### Domain: UPnP device control — First-party

UPNPDevice is the public device object; UPNPMiniUPNP and UPNPDeviceMiniUPNP provide the third-party-backed forms.

The UPnP module represents generic devices and MiniUPnP-backed specializations behind RefCounted APIs.

Read more: [UPnP device control](references/domain-concepts.md#topic-upnp-device-control)

Use these entity examples: UPNPDevice.

### Domain: Visual Shader source partition — First-party

This inventory establishes the module's source partition without relying on unlisted implementation bodies.

The Visual Shader build script compiles the module's C++ sources, its node sources, and editor sources when building the editor.

Read more: [Visual Shader source partition](references/domain-concepts.md#topic-visual-shader-module-build)

## Unit 53 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-webp-image-io-webxr-session-bridge-zip-archive-io`
Display slug: `unit-053-webp-image-i-o`

### Domain: WebP image I/O — First-party

It conditionally compiles bundled libwebp sources through its module build script.

The WebP module implements image loading, saving, and shared buffer helpers.

Read more: [WebP image I/O](references/domain-concepts.md#topic-webp-image-io)

Use these entity examples: ImageLoaderWebP, ResourceSaverWebP.

### Domain: WebXR session bridge — First-party

The JavaScript bridge manages sessions, reference spaces, frame data, views, and input-source data for the C++ implementation.

WebXR interface state receives browser session callbacks and input sources through web-platform bindings.

Read more: [WebXR session bridge](references/domain-concepts.md#topic-webxr-session-bridge)

Use these entity examples: WebXRInterfaceJS, WebXRInputSource.

### Domain: ZIP archive I/O — First-party

The module includes tests for compressed and uncompressed archives.

ZIPReader reads archives and ZIPPacker creates archives through reference-counted API objects.

Read more: [ZIP archive I/O](references/domain-concepts.md#topic-zip-archive-io)

Use these entity examples: ZIPReader, ZIPPacker.

## Unit 54 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-gltf-materials-textures-gltf-node-hierarchy-interactive-music-transitions`
Display slug: `unit-054-gltf-material-and-texture-conversion`

### Domain: glTF material and texture conversion — First-party

GLTFDocument maps glTF material dictionaries, PBR metallic-roughness values, textures, samplers, UV coordinates, color transforms, alpha mode, and texture-transform extensions to Godot material resources.

Read more: [glTF material and texture conversion](references/domain-concepts.md#topic-gltf-materials-textures)

Use these entity examples: GLTFState.

### Domain: glTF scene-node hierarchy — First-party

GLTFDocument reconstructs parent and child relationships among indexed GLTF node records and attaches meshes, cameras, lights, skins, and skeleton membership through node indexes.

Read more: [glTF scene-node hierarchy](references/domain-concepts.md#topic-gltf-node-hierarchy)

Use these entity examples: GLTFNode, GLTFState.

### Domain: interactive music transition tables — First-party

AudioStreamInteractive combines clips with a TransitionKey-indexed transition map so playback can select a rule for a clip-to-clip change.

Read more: [interactive music transition tables](references/domain-concepts.md#topic-interactive-music-transitions)

Use these entity examples: AudioStreamInteractive, AudioStreamInteractive Transition.

## Unit 55 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-mobile-vr-interface-physics-2d-collision-pipeline-physics-3d-collision-pipeline`
Display slug: `unit-055-mobile-xr-interface`

### Domain: mobile XR interface — First-party

MobileVRInterface is the module's concrete XRInterface implementation.

Read more: [mobile XR interface](references/domain-concepts.md#topic-mobile-vr-interface)

### Domain: native 2D collision pipeline — First-party

Physics spaces hold collision objects; the 2D BVH broad phase finds candidate collision objects and the solver dispatches shape-pair routines, including separating-axis tests.

Read more: [native 2D collision pipeline](references/domain-concepts.md#topic-physics-2d-collision-pipeline)

Use these entity examples: GodotSpace2D, GodotCollisionObject2D.

### Domain: native 3D collision pipeline — First-party

The native 3D server manages spaces, bodies, shapes, and joints; its collision code includes GJK/EPA support and SAT shape-pair routines.

Read more: [native 3D collision pipeline](references/domain-concepts.md#topic-physics-3d-collision-pipeline)

Use these entity examples: GodotSpace3D, GodotCollisionObject3D.

## Unit 56 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-android-export-pipeline-linuxbsd-desktop-integration-ios-embedded-adapter`
Display slug: `unit-056-android-export-pipeline`

### Domain: Android export pipeline — First-party

`EditorExportPlatformAndroid` is the editor-facing exporter and its helpers hold export-specific data structures.

Android export derives Gradle project files, manifest content, localized project names, icons, ABIs, and plugin configuration from project and export inputs.

Read more: [Android export pipeline](references/domain-concepts.md#topic-android-export-pipeline)

Use these entity examples: APKExportData, PluginConfigAndroid, LauncherIcon, CustomExportData.

### Domain: Linux/BSD desktop integration — First-party

The platform layer contains both generic Linux/BSD OS code and desktop-protocol integrations.

Linux/BSD platform code integrates OS services, portals, screensaver handling, TTS, X11, and Wayland display servers.

Read more: [Linux/BSD desktop integration](references/domain-concepts.md#topic-linuxbsd-desktop-integration)

### Domain: iOS embedded adapter — First-party

The iOS build lists device metrics, display, view, main, and OS Objective-C++ sources in one library.

iOS platform classes adapt the engine’s OS, display, and view responsibilities for an embedded Apple target.

Read more: [iOS embedded adapter](references/domain-concepts.md#topic-ios-embedded-adapter)

## Unit 57 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-native-desktop-services-platform-exports-android-remote-engine-service`
Display slug: `unit-057-native-desktop-services`

### Domain: Native desktop services — First-party

These services are separate native integrations rather than methods embedded solely in window management.

macOS and Windows platform code includes native-menu and text-to-speech adapters alongside their display servers.

Read more: [Native desktop services](references/domain-concepts.md#topic-native-desktop-services)

### Domain: Platform export packaging — First-party

Export code is platform-specific but follows EditorExportPlatform-derived services.

Platform export plug-ins implement macOS, Web, Windows, and visionOS export behavior; the Windows exporter also includes PE template modification structures.

Read more: [Platform export packaging](references/domain-concepts.md#topic-platform-exports)

### Domain: Remote Android engine service — First-party

The service defines numeric message codes and bundle keys, while a fragment binds to and receives remote engine state.

Remote engine service messages convey engine status, errors, surface packages, dimensions, host tokens, and command-line parameters between Android UI and service code.

Read more: [Remote Android engine service](references/domain-concepts.md#topic-android-remote-engine-service)

Use these entity examples: GodotService.EngineStatus.

## Unit 58 — Vendored: spirv-cross

Shared subsystem: Vendored: spirv-cross.

Lesson ID: `unit-extension-api-compatibility-spirv-intermediate-representation`
Display slug: `unit-058-spir-v-intermediate-representation`

### Domain: Extension API compatibility baselines — First-party

The baseline files record expected API changes such as added arguments, changed types, and removed APIs.

Version-pair directories retain expected extension-API validation differences, and the validator aggregates their relevant diagnostic lines.

Read more: [Extension API compatibility baselines](references/domain-concepts.md#topic-extension-api-compatibility)

### Domain: SPIR-V intermediate representation — Background concept

SPIRV-Cross stores these concepts in its common IR structures and compiler accessors.

The shader tooling parses SPIR-V into typed values, variables, blocks, functions, decorations, and entry-point metadata.

Read more: [SPIR-V intermediate representation](references/domain-concepts.md#topic-spirv-intermediate-representation)

Use these entity examples: SPIRType, SPIRVariable.

## Unit 59 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-resource-formats-reflection-cpp-variadic-binding`
Display slug: `unit-059-resource-formats-and-serialization`

### Domain: Resource formats and serialization — First-party

Format-specific code is separated from the common resource-loading service.

The resource-loading service uses binary, JSON, image, crypto, and extension resource-format loaders and savers.

Read more: [Resource formats and serialization](references/domain-concepts.md#topic-resource-formats)

Use these entity examples: Resource.

### Domain: Class metadata and reflection — First-party

Registration also records construction behavior and API exposure for native and extension classes.

ClassDB records Object-derived classes, properties, methods, signals, and constants for runtime lookup.

Read more: [Class metadata and reflection](references/domain-concepts.md#topic-reflection)

Use these entity examples: ClassInfo, Variant.

### Implementation: C++ variadic binding — First-party

This lets native method and signal calls share typed, compile-time-generated argument marshalling.

The binding code uses variadic templates and parameter packs to move argument lists into Variant pointer arrays.

Read more: [C++ variadic binding](references/language-concepts.md#topic-cpp-variadic-binding)

Apply it through: Variant, Object.

## Unit 60 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-dynamic-invocation-signals-variant-containers-cpp-copy-on-write-containers`
Display slug: `unit-060-dynamic-invocation-and-signals`

### Domain: Dynamic invocation and signals — First-party

The object and callable layers call Object methods and signal handlers with Variant argument arrays, including bound and unbound callable forms.

Read more: [Dynamic invocation and signals](references/domain-concepts.md#topic-dynamic-invocation-signals)

Use these entity examples: Object, Variant.

### Domain: Dynamic arrays and dictionaries — First-party

Array and Dictionary store Variant values, while typed validators and typed wrappers constrain their element, key, or value types.

Read more: [Dynamic arrays and dictionaries](references/domain-concepts.md#topic-variant-containers)

Use these entity examples: Array, Dictionary, Variant.

### Implementation: C++ copy-on-write container storage — First-party

Capacity growth, exact allocation, reallocation, and copy-to-new-buffer operations are explicitly represented in CowData.

The container layer uses C++ container classes to separate CowData buffer allocation and copying from Vector-style value interfaces.

Read more: [C++ copy-on-write container storage](references/language-concepts.md#topic-cpp-copy-on-write-containers)

Apply it through: Array, Dictionary.

## Unit 61 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-reflection-metadata-variant-text-serialization-cpp-scoped-locks`
Display slug: `unit-061-reflection-metadata`

### Domain: Reflection metadata — First-party

ClassDB, GDType, MethodInfo, and PropertyInfo describe Object classes, inheritance, methods, properties, constants, enums, and signals.

Read more: [Reflection metadata](references/domain-concepts.md#topic-reflection-metadata)

Use these entity examples: Object, GDType, MethodInfo, PropertyInfo.

### Domain: Variant text parsing and writing — First-party

VariantParser and VariantWriter serialize Variant values as String text through stream, token, tag, and resource-parser abstractions.

Read more: [Variant text parsing and writing](references/domain-concepts.md#topic-variant-text-serialization)

Use these entity examples: Variant, Array, Dictionary.

### Implementation: C++ scoped lock guards — First-party

Synchronization wrappers are declared in the OS layer and reused by object and reflection code.

The implementation uses a scoped C++ MutexLock object and class-specific lock helpers around shared runtime state.

Read more: [C++ scoped lock guards](references/language-concepts.md#topic-cpp-scoped-locks)

## Unit 62 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-expression-evaluation-mesh-geometry-algorithms-cpp-tagged-storage`
Display slug: `unit-062-c-tagged-storage-and-casts`

### Domain: Expression parsing and evaluation — First-party

Expression defines token and expression-node types and evaluates expression nodes against Variant inputs and built-in functions.

Read more: [Expression parsing and evaluation](references/domain-concepts.md#topic-expression-evaluation)

Use these entity examples: Variant.

### Domain: Mesh geometry algorithms — First-party

The implementation uses geometry primitives to build Delaunay triangulations, convex hulls, quick hulls, and polygon triangulations.

Read more: [Mesh geometry algorithms](references/domain-concepts.md#topic-mesh-geometry-algorithms)

Use these entity examples: TriangleMesh, AABB.

### Implementation: C++ tagged storage and casts — First-party

Variant separates the active type tag from storage-handling code for each supported runtime type.

The implementation represents runtime Variant values with a C++ type tag and accesses payloads through explicit casts and type-specific accessors.

Read more: [C++ tagged storage and casts](references/language-concepts.md#topic-cpp-tagged-storage)

Apply it through: Variant, Transform3D.

## Unit 63 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-packed-resource-files-profiling-interning-cpp-runtime-casts`
Display slug: `unit-063-c-static-casts-across-track-types`

### Domain: Packed resource files — First-party

PackedData is the virtual-file index behind pack-backed file access.

Packed data records pack paths, offsets, sizes, hashes, encryption flags, bundle flags, and delta flags, then supplies files to the resource-loading service through PackSource.

Read more: [Packed resource files](references/domain-concepts.md#topic-packed-resource-files)

Use these entity examples: PackedData::PackedFile.

### Domain: Profiling name interning — First-party

The profiling implementation interns StringName metadata and source-location data for tracing backends.

Read more: [Profiling name interning](references/domain-concepts.md#topic-profiling-interning)

Use these entity examples: StringName.

### Implementation: C++ static casts across track types — First-party

Animation evaluation repeats this pattern for several concrete track classes.

The implementation uses static_cast to interpret a base Track pointer as a derived Track object after a track kind has been selected.

Read more: [C++ static casts across track types](references/language-concepts.md#topic-cpp-runtime-casts)

Apply it through: Animation.

## Unit 64 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-script-hosting-spatial-indexing-godot-shader-language`
Display slug: `unit-064-godot-shader-source-composition`

### Domain: Script languages and instances — First-party

The scripting layer attaches Object-facing script instances, manages registered script languages, and provides extension-backed script implementations.

Read more: [Script languages and instances](references/domain-concepts.md#topic-script-hosting)

Use these entity examples: Object, Variant.

### Domain: Spatial indexing and ray queries — First-party

Spatial trees use geometry bounds to maintain items, cull them, refit nodes, and accelerate triangle-mesh and static-ray queries.

Read more: [Spatial indexing and ray queries](references/domain-concepts.md#topic-spatial-indexing)

Use these entity examples: TriangleMesh, AABB.

### Implementation: Godot shader source composition — First-party

The repository distinguishes high-level .gdshader source from RenderingDevice-native GLSL and SPIR-V resources.

Godot's shader language writes Shader resource source and can compose reusable ShaderInclude fragments with a preprocessor include directive.

Read more: [Godot shader source composition](references/language-concepts.md#topic-godot-shader-language)

Apply it through: ShaderMaterial, RDPipelineSpecializationConstant.

## Unit 65 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-transform-interpolation-localization-cpp-plugin-specialization`
Display slug: `unit-065-c-plugin-specialization`

### Domain: Transform interpolation — First-party

TransformInterpolator and InterpolatedProperty calculate intermediate geometry transforms and values.

Read more: [Transform interpolation](references/domain-concepts.md#topic-transform-interpolation)

Use these entity examples: Transform3D.

### Domain: Translation and locale selection — First-party

Translation catalogs map StringName keys to messages, TranslationDomain selects applicable catalogs, and TranslationServer manages locale, fallback, and plural-rule services.

Read more: [Translation and locale selection](references/domain-concepts.md#topic-localization)

Use these entity examples: Translation, TranslationDomain, StringName.

### Implementation: C++ plugin specialization — First-party

The pattern is repeated across scene, script, shader, resource, and version-control integrations.

C++ classes specialize EditorPlugin for feature-specific integrations; this requires C++ classes as the base abstraction mechanism.

Read more: [C++ plugin specialization](references/language-concepts.md#topic-cpp-plugin-specialization)

## Unit 66 — Startup & frame loop

Shared subsystem: Startup & frame loop.

Lesson ID: `unit-engine-bootstrap-navigation-queries-cpp-runtime-polymorphism`
Display slug: `unit-066-engine-bootstrap`

### Domain: Engine bootstrap — First-party

`Main` is the common startup layer before a selected platform host runs the engine.

Engine bootstrap consumes runtime configuration to initialize execution and select the configured main scene.

Read more: [Engine bootstrap](references/domain-concepts.md#topic-engine-bootstrap)

### Domain: 2D and 3D navigation queries — First-party

2D uses Vector2 paths and 3D uses Vector3 paths.

RefCounted navigation query parameter and result objects exchange path points, path types, and path-owner identifiers with the 2D and 3D navigation server APIs.

Read more: [2D and 3D navigation queries](references/domain-concepts.md#topic-navigation-queries)

Use these entity examples: NavigationPathQueryResult2D.

### Implementation: C++ virtual interfaces — First-party

This is the main extension idiom in the inspected editor code.

Godot uses polymorphic base interfaces so importers, inspector plug-ins, preview generators, and editor plug-ins can supply type-specific behavior through virtual methods.

Read more: [C++ virtual interfaces](references/language-concepts.md#topic-cpp-runtime-polymorphism)

## Unit 67 — Engine services

Shared subsystem: Engine services.

Lesson ID: `unit-physics-server-queries-pluggable-server-backends-cpp-godot-binding-macros`
Display slug: `unit-067-2d-and-3d-physics-queries`

### Domain: 2D and 3D physics queries — First-party

The 2D and 3D partitions mirror this server-plus-query arrangement.

Physics servers expose direct body and space state APIs, while RefCounted query parameter and result objects package ray, point, shape, and motion requests.

Read more: [2D and 3D physics queries](references/domain-concepts.md#topic-physics-server-queries)

### Domain: Pluggable and wrapped server backends — First-party

The same server-facing abstractions support fallback, extension-provided, and wrapped implementations.

C++ inheritance supplies extension, dummy, and multithread-wrapped implementations behind the physics, text, XR, and rendering server interfaces.

Read more: [Pluggable and wrapped server backends](references/domain-concepts.md#topic-pluggable-server-backends)

### Implementation: C++ engine binding macros — First-party

These macros are a Godot-specific layer over C++ declarations used throughout editor extension points.

Godot's C++ classes declare engine binding and script-overridable callbacks through macros such as GDCLASS and GDVIRTUAL.

Read more: [C++ engine binding macros](references/language-concepts.md#topic-cpp-godot-binding-macros)

## Unit 68 — Rendering & hardware

Shared subsystem: Rendering & hardware.

Lesson ID: `unit-apple-embedded-hosting-python-build-source-generation`
Display slug: `unit-068-apple-embedded-hosting`

### Domain: Apple embedded hosting — First-party

This layer packages the engine for Apple embedded targets.

Apple embedded hosting uses SwiftUI and Objective-C++ alongside display, keyboard, view-controller, text-to-speech, and Vulkan-context platform adapters.

Read more: [Apple embedded hosting](references/domain-concepts.md#topic-apple-embedded-hosting)

### Implementation: Python build source generation — First-party

The functions produce text and byte-oriented C++ output rather than editor runtime behavior.

Python build helpers perform build-time source generation by transforming inputs into generated C++ source and invoking external translation tooling when available.

Read more: [Python build source generation](references/language-concepts.md#topic-python-build-source-generation)

Apply it through: EditorTranslationList.

## Unit 69 — Rendering & hardware

Shared subsystem: Rendering & hardware.

Lesson ID: `unit-audio-backend-adapters-midi-input-adapters-python-build-actions`
Display slug: `unit-069-audio-backend-adapters`

### Domain: Audio backend adapters — First-party

The implementations derive from the common AudioDriver interface.

Audio output is provided by platform adapters selected with the driver build for ALSA, PulseAudio, CoreAudio, WASAPI, and XAudio2.

Read more: [Audio backend adapters](references/domain-concepts.md#topic-audio-backend-adapters)

### Domain: MIDI input adapters — First-party

Each named implementation derives from the common MIDIDriver interface.

MIDI input is supplied by platform adapters for ALSA MIDI, CoreMIDI, and WinMM.

Read more: [MIDI input adapters](references/domain-concepts.md#topic-midi-input-adapters)

### Implementation: Python build-action functions — First-party

Python is used for editor build support, not the editor’s primary runtime implementation.

Python defines the icon-build action that receives target, source, and environment arguments, while SCsub scripts invoke Python-based build-environment methods.

Read more: [Python build-action functions](references/language-concepts.md#topic-python-build-actions)

## Unit 70 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-resource-assets-csharp-source-generation`
Display slug: `unit-070-resource-backed-assets`

### Domain: Resource-backed assets — First-party

The resource family is the common content representation used throughout the inspected scene partition.

Resource-derived classes represent reusable scene content such as textures, meshes, materials, themes, animations, shapes, and packed scenes.

Read more: [Resource-backed assets](references/domain-concepts.md#topic-resource-assets)

Use these entity examples: PackedScene, Animation, TileMapPattern.

### Implementation: C# partial types and source generation — First-party

Godot's analyzers and generators inspect C# syntax and semantic symbols, report invalid declarations, and add generated source to the compilation.

Compilation produces attribute-bearing partial declarations for script paths and assembly script types.

Read more: [C# partial types and source generation](references/language-concepts.md#topic-csharp-source-generation)

Apply it through: ScriptPathAttribute, AssemblyHasScriptsAttribute, ManagedCallbacks.

## Unit 71 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-scene-tree-gui-controls-annotations-static-state-and-lifecycle`
Display slug: `unit-071-scene-graph-and-lifecycle`

### Domain: Scene graph and lifecycle — First-party

Nodes carry parent, owner, children, name, tree, processing, and RPC configuration state.

A scene tree arranges Object-derived Node instances into a parent-child hierarchy that SceneTree owns and updates.

Read more: [Scene graph and lifecycle](references/domain-concepts.md#topic-scene-tree)

Use these entity examples: Node, SceneState.

### Domain: GUI controls and layout — First-party

The partition includes both leaf controls and container controls.

Control-derived GUI classes implement selection, input, scrolling, split layout, menus, and graph editing within the CanvasItem tree.

Read more: [GUI controls and layout](references/domain-concepts.md#topic-gui-controls)

Use these entity examples: GraphEditConnection.

### Implementation: Annotations, static state, and lifecycle — First-party

The suite covers annotation validity, tool-mode requirements, static initialization, and ready-time values.

Annotations modify class declarations and members, while static state and `@onready` participate in initialization and lifecycle behavior.

Read more: [Annotations, static state, and lifecycle](references/language-concepts.md#topic-annotations-static-state-and-lifecycle)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 72 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-two-dimensional-physics-three-dimensional-physics-properties-and-accessors`
Display slug: `unit-072-2d-physics-nodes`

### Domain: 2D physics nodes — First-party

The hierarchy includes static, rigid, character, animatable, and joint node types.

2D physics scene graph nodes provide collision objects, bodies, joints, casts, and collision results.

Read more: [2D physics nodes](references/domain-concepts.md#topic-two-dimensional-physics)

Use these entity examples: CollisionObject2D, RigidBody2D, KinematicCollision2D.

### Domain: 3D physics nodes — First-party

The hierarchy includes static, rigid, character, physical-bone, soft-body, and vehicle implementations.

3D physics scene graph nodes provide collision objects, bodies, joints, casts, soft bodies, and vehicles.

Read more: [3D physics nodes](references/domain-concepts.md#topic-three-dimensional-physics)

Use these entity examples: CollisionObject3D, RigidBody3D, SoftBody3D.

### Implementation: Properties and accessors — First-party

Tests cover direct assignment, backing variables, getter/setter syntax, and compound access chains.

Class properties route reads and writes through inline or named accessors, including static and chained access.

Read more: [Properties and accessors](references/language-concepts.md#topic-properties-and-accessors)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 73 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-particle-systems-platform-presentation-signals-and-await`
Display slug: `unit-073-cpu-and-gpu-particle-systems`

### Domain: CPU and GPU particle systems — First-party

The module also includes 3D GPU collision and attractor node families.

2D and 3D particle nodes provide separate CPU and GPU particle-system implementations.

Read more: [CPU and GPU particle systems](references/domain-concepts.md#topic-particle-systems)

Use these entity examples: CPUParticles2D, GPUParticles3D.

### Domain: Display, camera, video, and movie services — First-party

A headless DisplayServer implementation is also present.

DisplayServer abstracts display functionality, CameraServer manages CameraFeed objects, VideoStreamPlayer presents video in a Control, and MovieWriter emits movie output.

Read more: [Display, camera, video, and movie services](references/domain-concepts.md#topic-platform-presentation)

### Implementation: Signals and await — First-party

Fixtures distinguish meaningful awaits from redundant ones and cover signal payload shapes.

Signal values and callable functions support waiting for emissions or coroutine completion with `await`.

Read more: [Signals and await](references/language-concepts.md#topic-signals-and-await)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 74 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-gui-control-layout-networking-typed-collections`
Display slug: `unit-074-gui-controls-and-layout`

### Domain: GUI controls and layout — First-party

The implementation includes buttons, text and code editing, dialogs, file selection, color picking, graph editing, and flow layout.

Control-derived widgets and Container-derived layouts implement the scene GUI layer.

Read more: [GUI controls and layout](references/domain-concepts.md#topic-gui-control-layout)

Use these entity examples: Control, Container, CodeEdit.

### Domain: HTTP and multiplayer — First-party

The two APIs appear as separate scene-main facilities.

HTTPRequest is a Node API for HTTP work, while MultiplayerAPI and MultiplayerPeer define reference-counted multiplayer and packet-peer interfaces.

Read more: [HTTP and multiplayer](references/domain-concepts.md#topic-networking)

### Implementation: Typed arrays and dictionaries — First-party

Collection shape is tested both during analysis and during runtime assignment or argument passing.

Typed arrays and dictionaries apply static typing to element and key/value shapes; the fixture suite exercises declarations, inference, casts, defaults, transfers, and invalid runtime assignments.

Read more: [Typed arrays and dictionaries](references/language-concepts.md#topic-typed-collections)

Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 75 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-navigation-agents-cpp-jni-variant-marshalling`
Display slug: `unit-075-c-jni-variant-marshalling`

### Domain: Navigation agents and regions — First-party

Separate 2D and 3D implementations retain paths, path metadata, links, regions, and obstacles.

Navigation agents consume navigation-result paths, while regions, links, and obstacles describe navigation-world participation.

Read more: [Navigation agents and regions](references/domain-concepts.md#topic-navigation-agents)

Use these entity examples: NavigationAgent2D, NavigationAgent3D, NavigationRegion3D.

### Implementation: C++ JNI Variant marshalling — First-party

The native Android wrapper identifies Java values through JNI calls and produces Godot `Variant` representations.

C++ JNI Variant marshalling builds on RefCounted adapter classes to convert Java primitive wrappers, strings, arrays, objects, and callables into engine Variant values.

Read more: [C++ JNI Variant marshalling](references/language-concepts.md#topic-cpp-jni-variant-marshalling)

Apply it through: Callable.

## Unit 76 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-skeletal-modifiers-and-ik-c-dynamic-library-wrappers`
Display slug: `unit-076-c-dynamic-library-function-wrappers`

### Domain: Skeletal modifiers and inverse kinematics — First-party

The module contains multiple IK algorithms and skeletal runtime modifiers.

SkeletonModifier3D subclasses apply bone constraints, IK solvers, retargeting, spring-bone simulation, and XR tracker data.

Read more: [Skeletal modifiers and inverse kinematics](references/domain-concepts.md#topic-skeletal-modifiers-and-ik)

Use these entity examples: SkeletonModifier3D, ChainIK3D, SpringBoneSimulator3D.

### Implementation: C dynamic-library function wrappers — First-party

The generated wrapper headers declare replacement function-pointer symbols and identify the wrapped shared-library sonames.

Generated C function-pointer wrappers dynamically route DBus, Fontconfig, Speech Dispatcher, and Wayland/libdecor calls for Linux/BSD desktop integration.

Read more: [C dynamic-library function wrappers](references/language-concepts.md#topic-c-dynamic-library-wrappers)

## Unit 77 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-tilemap-layer-data-physics-queries-objective-cpp-apple-adapters`
Display slug: `unit-077-physics-queries-and-motion-tests`

### Domain: Tile map layer data — First-party

A TileMap node coordinates TileMapLayer instances.

TileMapLayer stores cell data keyed by grid coordinates and derives rendering, physics, navigation, and debug quadrants from that data.

Read more: [Tile map layer data](references/domain-concepts.md#topic-tilemap-layer-data)

Use these entity examples: TileMapLayer, TileMapLayerCellData.

### Domain: Physics queries and motion tests — First-party

Direct-space-state calls provide ad hoc queries, while cast nodes keep query behavior in the scene tree.

Physics queries use physics spaces to submit configured point, ray, shape, and motion tests and receive collision data.

Read more: [Physics queries and motion tests](references/domain-concepts.md#topic-physics-queries)

Use these entity examples: PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D.

### Implementation: Objective-C++ Apple platform adapters — First-party

macOS and visionOS implementation files use .mm sources where C++ engine classes meet Apple platform APIs.

Objective-C++ implementation files combine C++ platform adapters with Cocoa-style objects, display coordinates, events, menus, views, and text-to-speech calls.

Read more: [Objective-C++ Apple platform adapters](references/language-concepts.md#topic-objective-cpp-apple-adapters)

## Unit 78 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-extensions-network-data-exchange-objective-cpp-ios-adapters`
Display slug: `unit-078-virtual-implementation-extensions`

### Domain: Virtual implementation extensions — First-party

The contracts are exposed as subclass APIs such as PhysicsServer2DExtension and TextServerExtension.

Extension classes declare required or optional virtual callbacks so external implementations can replace physics, rendering, text, and scripting behavior.

Read more: [Virtual implementation extensions](references/domain-concepts.md#topic-extensions)

Use these entity examples: RenderSceneBuffersConfiguration, RID.

### Domain: Packets, HTTP, JSON, and JSON-RPC — First-party

These APIs separate transport helpers from the high-level multiplayer interface.

Packet peers transmit raw bytes or Variant values, while JSON and JSON-RPC map values into external message documents.

Read more: [Packets, HTTP, JSON, and JSON-RPC](references/domain-concepts.md#topic-network-data-exchange)

Use these entity examples: JSON-RPC document, Variant.

### Implementation: Objective-C++ iOS adapters — First-party

The iOS SCons build compiles `.mm` sources for device metrics, display, view, main, and OS behavior.

Objective-C++ iOS adapters implement an embedded platform layer for the iOS embedded adapter.

Read more: [Objective-C++ iOS adapters](references/language-concepts.md#topic-objective-cpp-ios-adapters)

## Unit 79 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-skeletal-ik-export-presets-collada-import`
Display slug: `unit-079-export-presets`

### Domain: Skeletal modifiers and inverse kinematics — First-party

The repository supplies general IK bases plus FABRIK and Jacobian solver variants.

IK skeleton modifiers process bone chains, targets, and joint-limitation resources to modify skeletal poses.

Read more: [Skeletal modifiers and inverse kinematics](references/domain-concepts.md#topic-skeletal-ik)

Use these entity examples: Resource.

### Domain: Export presets — First-party

Preset configuration is target-platform-specific while retaining resource-selection policy.

An export preset selects indexed project files, applies include, exclude, and customized-file rules, and records output path, features, patches, and encryption options.

Read more: [Export presets](references/domain-concepts.md#topic-export-presets)

Use these entity examples: EditorExportPreset.

### Domain: COLLADA scene interchange — First-party

COLLADA is represented as an intermediate parser state before editor scene construction.

The COLLADA parser supplies a scene importer with image, material, effect, mesh, skin, morph, node, visual-scene, and animation data collected in Collada::State.

Read more: [COLLADA scene interchange](references/domain-concepts.md#topic-collada-import)

Use these entity examples: ColladaParseState.

## Unit 80 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-editing-selection-history-editor-plugin-state-input-action-configuration`
Display slug: `unit-080-editing-selection-history`

### Domain: Editing selection history — First-party

The history stores object paths and the currently selected point in that path history.

Selection history gives an editor scene session back-and-forward navigation across selected objects and nested-resource properties.

Read more: [Editing selection history](references/domain-concepts.md#topic-editing-selection-history)

Use these entity examples: EditedScene.

### Domain: Editor plugin state and custom types — First-party

EditorData selects a main editor plugin for an object and retains plugin-provided type metadata.

Editor plugins can handle objects, register custom types, and serialize plugin state into the active scene's editor state.

Read more: [Editor plugin state and custom types](references/domain-concepts.md#topic-editor-plugin-state)

Use these entity examples: EditedScene.

### Domain: Input action and shortcut configuration — First-party

The implementation supports keyboard, mouse, and joypad event capture in a shared line-edit control.

Input action configuration depends on editor and project settings because ActionMapEditor, event search, event capture, and input-event configuration edit action-event mappings.

Read more: [Input action and shortcut configuration](references/domain-concepts.md#topic-input-action-configuration)

## Unit 81 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-platform-selective-shader-baking-post-import-processing-resource-dependency-resolution`
Display slug: `unit-081-platform-selective-shader-baking`

### Domain: Platform-selective shader baking — First-party

The platform split occurs in build manifests and is reflected in three platform adapter classes.

Platform-selective shader baking depends on shader editing and compiles distinct Vulkan, D3D12, or Metal export-plugin sources when matching SCons environment flags are enabled.

Read more: [Platform-selective shader baking](references/domain-concepts.md#topic-platform-selective-shader-baking)

### Domain: Post-import processing — First-party

The implementation separates parsing a source format from editor-controlled transformations of its imported result.

Post-import processing operates on the imported scene after format conversion and accepts options for node, mesh, material, animation, and skeleton categories.

Read more: [Post-import processing](references/domain-concepts.md#topic-post-import-processing)

### Domain: Resource dependency resolution — First-party

Stored dependency values can carry type and UID information that is resolved before presentation.

Dependency resolution uses file-index entries and ResourceUID mappings to expose each indexed resource's current dependency paths.

Read more: [Resource dependency resolution](references/domain-concepts.md#topic-resource-dependency-resolution)

Use these entity examples: EditorFileSystemDirectory::FileInfo.

## Unit 82 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-scene-tree-and-signal-connections-undo-redo-history-completion-contexts`
Display slug: `unit-082-contextual-completion-contracts`

### Domain: Scene tree and signal connections — First-party

The source treats structural scene navigation and signal wiring as editor services.

Scene authoring is accompanied by a SceneTreeEditor and a ConnectionsDock that inspect nodes, signals, methods, connection arguments, and bound values.

Read more: [Scene tree and signal connections](references/domain-concepts.md#topic-scene-tree-and-signal-connections)

### Domain: Undo and redo history — First-party

The manager separates histories by integer ID and retains reference-counted objects required by undo and redo stacks.

Undo and redo assigns actions and saved versions to a scene session, allowing edits to identify a history and report unsaved state.

Read more: [Undo and redo history](references/domain-concepts.md#topic-undo-redo-history)

Use these entity examples: EditorUndoRedoManager::History, EditedScene.

### Domain: Contextual completion contracts — First-party

The configuration records assertions over displayed or inserted completion text.

Completion fixture contracts pair a cursor-bearing script with configuration rules that include or exclude candidates.

Read more: [Contextual completion contracts](references/domain-concepts.md#topic-completion-contexts)

Use these entity examples: Completion Test Configuration, Test Script Fixture.

## Unit 83 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-webrtc-data-channels-diagnostic-expectations-enet-transport-and-multiplayer`
Display slug: `unit-083-webrtc-data-channels`

### Domain: WebRTC data channels — First-party

The module supplies a base channel plus extension and JavaScript implementations.

A WebRTC data channel carries PacketPeer data through a WebRTC peer connection.

Read more: [WebRTC data channels](references/domain-concepts.md#topic-webrtc-data-channels)

Use these entity examples: WebRTCDataChannel.

### Domain: Diagnostic expectations — First-party

This makes diagnostic behavior part of the tested interface.

Diagnostic fixture contracts preserve parser errors, analyzer warnings, and runtime errors as expected textual results.

Read more: [Diagnostic expectations](references/domain-concepts.md#topic-diagnostic-expectations)

Use these entity examples: Expected Result Fixture.

### Domain: ENet transport and multiplayer adaptation — First-party

The connection implementation also invokes the engine Compression API for payload data.

ENetConnection and ENetPacketPeer wrap connection and peer behavior, while ENetMultiplayerPeer adapts the transport to MultiplayerPeer.

Read more: [ENet transport and multiplayer adaptation](references/domain-concepts.md#topic-enet-transport-and-multiplayer)

Use these entity examples: ENetConnection, ENetPacketPeer, ENetMultiplayerPeer.

## Unit 84 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-gdscript-bytecode-runtime-gpu-texture-compression-godot-member-exposure`
Display slug: `unit-084-gdscript-bytecode-compilation-and-execution`

### Domain: GDScript bytecode compilation and execution — First-party

GDScriptFunction stores execution metadata, while GDScriptInstance binds a script to an engine object.

The compiler lowers analyzed script trees through code-generation classes, and the VM executes functions using validated calls, getters, setters, and container operations.

Read more: [GDScript bytecode compilation and execution](references/domain-concepts.md#topic-gdscript-bytecode-runtime)

Use these entity examples: GDScript, GDScriptInstance, GDScriptFunction.

### Domain: GPU block-compression dispatch — First-party

The module packages BC1, BC4, BC6H, alpha-stitching, and RGB-to-RGBA shader sources.

The Betsy compressor relies on GLSL compute shaders that fetch source texels or blocks and write compressed results to storage images.

Read more: [GPU block-compression dispatch](references/domain-concepts.md#topic-gpu-texture-compression)

Use these entity examples: Image.

### Domain: Godot member exposure — First-party

The source generators inspect compatible C# members and create the metadata queried by ScriptManagerBridge.

Generator output registers exported members, signals, and RPC metadata as engine-visible method and property descriptions.

Read more: [Godot member exposure](references/domain-concepts.md#topic-godot-member-exposure)

Use these entity examples: Variant.

## Unit 85 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-lsp-semantic-fixtures-navigation-avoidance-navigation-path-queries`
Display slug: `unit-085-language-server-semantic-fixtures`

### Domain: Language-server semantic fixtures — First-party

These inputs are distinct from completion assertions because they model symbols and source ranges.

Language-server fixture contracts provide nested declarations, local scopes, documentation comments, and lambdas for semantic-editor tests.

Read more: [Language-server semantic fixtures](references/domain-concepts.md#topic-lsp-semantic-fixtures)

Use these entity examples: Test Script Fixture.

### Domain: Navigation avoidance — First-party

Both dimensions model agent positions and velocities plus obstacle geometry, and their maps expose agent and obstacle collections.

Agents and obstacles are map members that provide avoidance data alongside pathfinding data.

Read more: [Navigation avoidance](references/domain-concepts.md#topic-navigation-avoidance)

### Domain: Navigation path queries — First-party

The query implementations maintain begin and end polygons, search data, path points, and path-return constraints.

Path queries consume a map read iteration, select polygons, and construct a path corridor in 2D or 3D.

Read more: [Navigation path queries](references/domain-concepts.md#topic-navigation-path-queries)

## Unit 86 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-openxr-composition-layers-openxr-extension-wrappers-openxr-render-models`
Display slug: `unit-086-openxr-composition-layers`

### Domain: OpenXR composition layers — First-party

The scene layer has quad, cylinder, and equirect subclasses, while the extension owns swapchain and composition-layer state.

Composition-layer scene nodes and an extension submit specialized layer descriptions to the OpenXR runtime.

Read more: [OpenXR composition layers](references/domain-concepts.md#topic-openxr-composition-layers)

### Domain: OpenXR extension wrappers — First-party

The supplied extensions cover platform graphics bindings, controller profiles, composition layers, hand tracking, render models, futures, and spatial capabilities.

Extension wrappers isolate optional OpenXR runtime features behind a common base interface.

Read more: [OpenXR extension wrappers](references/domain-concepts.md#topic-openxr-extension-wrappers)

Use these entity examples: OpenXRFutureResult, OpenXRRenderModelData.

### Domain: OpenXR render models — First-party

Render-model data is separate from interaction-profile suggestions, allowing runtime-provided device models to drive presentation.

The render-model extension tracks runtime render models, and scene nodes display models or manage their activation.

Read more: [OpenXR render models](references/domain-concepts.md#topic-openxr-render-models)

Use these entity examples: OpenXRRenderModelData.

## Unit 87 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-plugin-and-assembly-reload-replicated-property-synchronization-replicated-spawning`
Display slug: `unit-087-plugin-and-assembly-reload`

### Domain: Plugin and assembly reload — First-party

The loader separates shared assemblies from plugin-specific dependency resolution and records a weak reference for load-context tracking.

Editor plug-ins load into AssemblyLoadContext instances that resolve dependencies and may be collectible.

Read more: [Plugin and assembly reload](references/domain-concepts.md#topic-plugin-and-assembly-reload)

### Domain: Replicated property synchronization — First-party

The synchronizer builds watchers from configured paths and participates in replication start and stop through object configuration calls.

Synchronizers observe configured properties and forward sync property lists to scene replication.

Read more: [Replicated property synchronization](references/domain-concepts.md#topic-replicated-property-synchronization)

Use these entity examples: SceneReplicationConfig.

### Domain: Replicated spawning — First-party

Spawner state records tracked nodes and their spawn data, while the replication interface sends spawn and despawn commands.

The spawner tracks spawnable scenes and sends configured spawn property lists while creating or removing nodes.

Read more: [Replicated spawning](references/domain-concepts.md#topic-replicated-spawning)

Use these entity examples: SceneReplicationConfig.

## Unit 88 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-module-build-configuration-visual-shader-vector-operations-websocket-multiplayer`
Display slug: `unit-088-scons-module-build-configuration`

### Domain: SCons module build configuration — First-party

Python SCons scripts determine whether modules can build and add C++ source files, generated shader headers, or third-party sources to the build environment.

Read more: [SCons module build configuration](references/domain-concepts.md#topic-module-build-configuration)

### Domain: Visual shader vector operations — First-party

The node set includes operator, function, distance, refraction, compose, and decompose variants.

Vector-operation nodes apply arithmetic and vector functions as operations within a visual shader graph.

Read more: [Visual shader vector operations](references/domain-concepts.md#topic-visual-shader-vector-operations)

Use these entity examples: VisualShaderNodeVectorOp.

### Domain: WebSocket multiplayer — First-party

Its packet and pending-peer records are the cross-cutting state used by the multiplayer transport.

WebSocketMultiplayerPeer tracks WebSocket peers plus pending peers and packets.

Read more: [WebSocket multiplayer](references/domain-concepts.md#topic-websocket-multiplayer)

Use these entity examples: WebSocketMultiplayerPeer, WebSocketMultiplayerPacket.

## Unit 89 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-gltf-accessors-managed-csharp-scripting-android-gradle-assembly`
Display slug: `unit-089-android-gradle-assembly`

### Domain: glTF binary accessor decoding and encoding — First-party

Mesh attributes, indices, animation samples, and other numeric glTF values pass through this layer.

GLTFAccessor converts typed values to and from GLTFState’s indexed buffer views, including sparse data when the sparse representation is smaller than dense data.

Read more: [glTF binary accessor decoding and encoding](references/domain-concepts.md#topic-gltf-accessors)

Use these entity examples: GLTFAccessor, GLTFBufferView, GLTFState.

### Domain: managed C# script bridge and source generation — First-party

The Mono module represents managed scripts through CSharpScript and CSharpInstance, while its SDK tests use partial C# types and source generators for properties, methods, signals, serialization, and script paths.

Read more: [managed C# script bridge and source generation](references/domain-concepts.md#topic-managed-csharp-scripting)

Use these entity examples: CSharpScript.

### Domain: Android Gradle assembly — First-party

The Android build root explicitly includes all five module names.

Gradle settings and build logic assemble app, library, editor, native-source-index, and install-time asset-pack modules for the Android export pipeline.

Read more: [Android Gradle assembly](references/domain-concepts.md#topic-android-gradle-assembly)

## Unit 90 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-apk-signing-browser-runtime-adapters-platform-display-adaptation`
Display slug: `unit-090-apk-signing-and-verification`

### Domain: APK signing and verification — First-party

The Android editor has a Kotlin utility that imports `ApkSigner` and `ApkVerifier`.

APK signing and verification operate on Android export pipeline artifacts through the bundled apksig implementation.

Read more: [APK signing and verification](references/domain-concepts.md#topic-apk-signing)

### Domain: Browser runtime adapters — First-party

The browser implementation is split between C++ platform objects and JavaScript libraries compiled into the Web build.

Web audio, display, input, fetch, MIDI, filesystem, and WebGL adapters call JavaScript libraries and exchange data with C++ Web platform classes.

Read more: [Browser runtime adapters](references/domain-concepts.md#topic-browser-runtime-adapters)

Use these entity examples: JavaScriptObjectImpl.

### Domain: Platform display and window adaptation — First-party

The same engine-facing display role is implemented separately for X11, macOS, Windows, and Web.

DisplayServer implementations adapt platform windows and events; the Web adapter constructs InputEvent instances and forwards them to Input.

Read more: [Platform display and window adaptation](references/domain-concepts.md#topic-platform-display-adaptation)

Use these entity examples: InputEvent.

## Unit 91 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-wayland-window-backend-web-javascript-bridge`
Display slug: `unit-091-wayland-window-backend`

### Domain: Wayland window backend — First-party

Wayland state is split between a display server, a protocol thread, and an embedder.

Wayland window handling builds on Linux/BSD platform integration with a dedicated thread, client protocol objects, and window state.

Read more: [Wayland window backend](references/domain-concepts.md#topic-wayland-window-backend)

Use these entity examples: WaylandThread.WindowState.

### Domain: Web JavaScript bridge — First-party

This is the explicit engine-to-browser object boundary, distinct from browser input and audio adapters.

The JavaScript bridge exposes evaluation, interface lookup, callback creation, object creation, buffer conversion, downloads, PWA operations, and JavaScript object proxies.

Read more: [Web JavaScript bridge](references/domain-concepts.md#topic-web-javascript-bridge)

Use these entity examples: JavaScriptObjectImpl.

## Unit 92 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-native-extensions-deferred-execution-cpp-specialized-marshalling`
Display slug: `unit-092-c-specialized-argument-marshalling`

### Domain: Native extension loading — First-party

This lets native libraries add engine-facing types without compiling them into the core binary.

A GDExtension is a Resource loaded through a loader; its registered classes become ClassDB-visible extension classes.

Read more: [Native extension loading](references/domain-concepts.md#topic-native-extensions)

Use these entity examples: GDExtension, ClassInfo.

### Domain: Deferred calls and worker tasks — First-party

MessageQueue records Object/Callable work for later execution, while WorkerThreadPool represents task and group work for worker threads.

Read more: [Deferred calls and worker tasks](references/domain-concepts.md#topic-deferred-execution)

Use these entity examples: Object, Variant.

### Implementation: C++ specialized argument marshalling — First-party

The specializations select direct, converting, reference, vector, pointer, object, reference-counted, and typed-container paths.

The binding layer specializes PtrToArg through C++ templates to marshal bound method types and return values.

Read more: [C++ specialized argument marshalling](references/language-concepts.md#topic-cpp-specialized-marshalling)

Apply it through: Variant, PropertyInfo.

## Unit 93 — Core runtime

Shared subsystem: Core runtime.

Lesson ID: `unit-generic-containers-cpp-template-binding`
Display slug: `unit-093-c-template-based-method-binding`

### Domain: Generic container infrastructure — First-party

The template layer supplies the C++ storage and lookup structures that underpin Array and Dictionary, including CowData, Vector, maps, sets, lists, spans, and RID owners.

Read more: [Generic container infrastructure](references/domain-concepts.md#topic-generic-containers)

Use these entity examples: Array, Dictionary.

### Implementation: C++ template-based method binding — First-party

This keeps native method registration type-aware while producing the reflective MethodBind representation used by script-facing APIs.

ClassDB uses C++ templates, `if constexpr`, and member-function traits to generate binders for native methods.

Read more: [C++ template-based method binding](references/language-concepts.md#topic-cpp-template-binding)

Apply it through: ClassInfo, Variant.

## Unit 94 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-input-actions-gdscript-query-objects`
Display slug: `unit-094-gdscript-query-object-api-use`

### Domain: Input events and actions — First-party

The model covers physical keyboard, mouse, touch, joystick, gesture, and manually generated action events.

InputMap associates named actions with InputEvent instances, and nodes receive input-event callbacks.

Read more: [Input events and actions](references/domain-concepts.md#topic-input-actions)

Use these entity examples: InputEvent.

### Implementation: GDScript query-object API use — First-party

The documented API uses object construction and member access rather than an inline positional-argument ray query.

GDScript can configure a physics query object, pass it to a physics space, and receive the collision result through chained method calls.

Read more: [GDScript query-object API use](references/language-concepts.md#topic-gdscript-query-objects)

Apply it through: PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D.

## Unit 95 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-undo-redo-csharp-unsafe-interop`
Display slug: `unit-095-c-unsafe-interop-and-function-pointers`

### Domain: Undo and redo actions — First-party

UndoRedo records Object/Callable operations as actions so they can be replayed in undo or redo direction.

Read more: [Undo and redo actions](references/domain-concepts.md#topic-undo-redo)

Use these entity examples: Object, Variant.

### Implementation: C# unsafe interop and function pointers — First-party

GodotSharp validates unmanaged table size before reinterpreting its address, marshals engine values through native layouts, and uses generated callback tables for native-to-managed calls.

Unsafe C# stores pointers and function pointers in ABI structs and invokes generated partial native calls.

Read more: [C# unsafe interop and function pointers](references/language-concepts.md#topic-csharp-unsafe-interop)

Apply it through: ManagedCallbacks, Variant.

## Unit 96 — Startup & frame loop

Shared subsystem: Startup & frame loop.

Lesson ID: `unit-frame-timing-performance-monitors-cmake-native-source-index`
Display slug: `unit-096-frame-timing-and-physics-stepping`

### Domain: Frame timing and physics stepping — First-party

The timer synchronization layer separates frame-duration handling from main startup.

Frame timing calculates a process delta and bounded physics-step count after engine bootstrap determines the active physics tick configuration.

Read more: [Frame timing and physics stepping](references/domain-concepts.md#topic-frame-timing)

Use these entity examples: MainFrameTime.

### Domain: Performance monitors — First-party

The `Performance` object and `MonitorCall` class are in the main runtime partition.

Performance monitors collect engine counters and execute registered monitor calls after engine bootstrap.

Read more: [Performance monitors](references/domain-concepts.md#topic-performance-monitors)

### Implementation: CMake native source indexing — First-party

The CMake file explicitly labels itself non-functional and configures it for Android Studio editor support.

CMake source-index configuration exposes the native C/C++ source tree to Android Studio as part of Android Gradle assembly.

Read more: [CMake native source indexing](references/language-concepts.md#topic-cmake-native-source-index)

## Unit 97 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-audio-bus-processing-packed-scenes-groovy-gradle-build-logic`
Display slug: `unit-097-packed-scene-serialization`

### Domain: Audio buses, streams, and effects — First-party

Effect classes create reference-counted processing instances.

AudioServer owns buses with channels and effects, while Resource-derived streams, effects, and bus layouts define reusable audio configuration.

Read more: [Audio buses, streams, and effects](references/domain-concepts.md#topic-audio-bus-processing)

Use these entity examples: AudioBusLayout.

### Domain: Packed scene serialization — First-party

SceneState separates reusable name, Variant, path, node, and connection tables from the PackedScene resource wrapper.

A PackedScene stores a resource-backed SceneState whose node entries, property values, and connection entries reconstruct a node hierarchy.

Read more: [Packed scene serialization](references/domain-concepts.md#topic-packed-scenes)

Use these entity examples: PackedScene, SceneState.

### Implementation: Groovy Gradle build logic — First-party

The Android Gradle root, app, library, and editor build files contain executable Groovy configuration.

Groovy Gradle scripts generate Android build tasks and select variants as part of Android Gradle assembly.

Read more: [Groovy Gradle build logic](references/language-concepts.md#topic-groovy-gradle-build-logic)

## Unit 98 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-two-dimensional-content-skeleton-modification-three-dimensional-content`
Display slug: `unit-098-2d-shapes-tiles-and-paths`

### Domain: 2D shapes, tiles, and paths — First-party

The 2D resource partition is conditionally compiled for physics and navigation features.

Resource-owned Shape2D subclasses, TileSet data, NavigationPolygon data, and PolygonPathFinder provide 2D geometry and tile-oriented content.

Read more: [2D shapes, tiles, and paths](references/domain-concepts.md#topic-two-dimensional-content)

Use these entity examples: TileMapPattern.

### Domain: 2D skeleton modification stacks — First-party

The supplied stack includes CCDIK, FABRIK, LookAt, TwoBoneIK, jiggle, and physical-bone variants.

SkeletonModification2D resources define individual 2D skeleton operations, and SkeletonModificationStack2D holds modifications for a Skeleton2D.

Read more: [2D skeleton modification stacks](references/domain-concepts.md#topic-skeleton-modification)

### Domain: 3D shapes, worlds, and skins — First-party

The 3D resource partition is selected only when 3D is enabled by the build configuration.

Resource-owned Shape3D subclasses, World3D, Skin, and mesh import data represent 3D collision, world, skeletal, and mesh content.

Read more: [3D shapes, worlds, and skins](references/domain-concepts.md#topic-three-dimensional-content)

## Unit 99 — Scene runtime

Shared subsystem: Scene runtime.

Lesson ID: `unit-text-interface-rendering-resources-themes-and-style-boxes`
Display slug: `unit-099-text-layout-and-editing`

### Domain: Text layout and editing — First-party

TextEdit, LineEdit, Label, and RichTextLabel each consume shaped-text resources in this partition.

GUI text controls hold shaped line and paragraph data, use text-server glyph and selection queries, and track IME text and selection state.

Read more: [Text layout and editing](references/domain-concepts.md#topic-text-interface)

### Domain: Textures, meshes, and materials — First-party

The inspected source includes array meshes, primitive meshes, compressed and image textures, and several material families.

Resource-owned Mesh, Material, Texture, Environment, and Sky classes carry rendering-facing content and configuration.

Read more: [Textures, meshes, and materials](references/domain-concepts.md#topic-rendering-resources)

### Domain: Themes and style boxes — First-party

Default-theme construction also appears in the inspected theme source.

Resource-backed Theme data is resolved through ThemeDB, ThemeContext, and ThemeOwner, while StyleBox subclasses supply control styling.

Read more: [Themes and style boxes](references/domain-concepts.md#topic-themes-and-style-boxes)

## Unit 100 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-rendering-assets-physics-collision-resource-identifiers`
Display slug: `unit-100-meshes-materials-textures-and-instancing`

### Domain: Meshes, materials, textures, and instancing — First-party

The resource layer holds drawable data while scene nodes select where and how it is rendered.

Mesh resources provide surfaces; geometry nodes instance them, Materials control shading, and MultiMesh batches many instances.

Read more: [Meshes, materials, textures, and instancing](references/domain-concepts.md#topic-rendering-assets)

Use these entity examples: Mesh, Material.

### Domain: Physics shapes, objects, and collision reports — First-party

The implementation provides parallel 2D and 3D object, shape, joint, and movement-report APIs.

2D and 3D CollisionObject nodes own Shape resources through shape owners; body movement returns KinematicCollision results.

Read more: [Physics shapes, objects, and collision reports](references/domain-concepts.md#topic-physics-collision)

Use these entity examples: Shape2D, CollisionShape2D, KinematicCollision2D.

### Domain: Resource and server identifiers — First-party

The two identifiers intentionally serve different lifetimes: project references versus live server objects.

ResourceUID maps project resource identifiers to paths, while an RID is an opaque session-only handle for a low-level server resource.

Read more: [Resource and server identifiers](references/domain-concepts.md#topic-resource-identifiers)

Use these entity examples: ResourceUID, RID.

## Unit 101 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-ui-composition-scene-state-textures-and-placeholders`
Display slug: `unit-101-control-tree-user-interfaces`

### Domain: Control-tree user interfaces — First-party

The same tree model supports widgets, dialogs, menus, graph editors, and editor controls.

Control-derived nodes compose user interfaces, and Container nodes automatically arrange their Control children.

Read more: [Control-tree user interfaces](references/domain-concepts.md#topic-ui-composition)

Use these entity examples: Node.

### Domain: Scene state inspection — First-party

It is an inspection representation rather than a mutable scene graph.

SceneState exposes a packed scene's resources, nodes, exported properties, overrides, and built-in scripts without instantiating the scene.

Read more: [Scene state inspection](references/domain-concepts.md#topic-scene-state)

Use these entity examples: SceneState, Resource.

### Domain: Texture resources and fallback placeholders — First-party

Layered and 3D textures require consistently sized image data across their layers.

Texture resources represent 2D, 3D, layered, and RenderingDevice-backed image data; placeholder texture classes preserve limited shape information when source texture data is unavailable or omitted.

Read more: [Texture resources and fallback placeholders](references/domain-concepts.md#topic-textures-and-placeholders)

Use these entity examples: RDTextureFormat, RID.

## Unit 102 — Editor & export

Shared subsystem: Editor & export.

Lesson ID: `unit-tile-libraries-editor-plugin-lifecycle-export-orchestration`
Display slug: `unit-102-editor-plugin-lifecycle`

### Domain: Tile libraries, cells, and patterns — First-party

Tile sources can expose atlas-backed or scene-backed tiles.

A TileSet Resource owns tile sources, identifies tiles by source, atlas-coordinate, and alternative IDs, and supplies tile data to TileMapLayer cells and TileMapPattern copies.

Read more: [Tile libraries, cells, and patterns](references/domain-concepts.md#topic-tile-libraries)

Use these entity examples: TileSet, TileData.

### Domain: Editor plugin lifecycle — First-party

The implementation favors many focused plugin subclasses over one monolithic editor tool.

C++ editor-plugin specializations package feature-specific editor behavior behind EditorPlugin-derived classes, including scene, script, and shader tools.

Read more: [Editor plugin lifecycle](references/domain-concepts.md#topic-editor-plugin-lifecycle)

### Domain: Export orchestration — First-party

EditorExport is the registry joining target configuration with exporter extensions.

Export orchestration owns export presets, target platforms, and plugins, and maps a target platform to its runnable preset.

Read more: [Export orchestration](references/domain-concepts.md#topic-export-orchestration)

Use these entity examples: EditorExport, EditorExportPreset.

## Unit 103 — Modules & extensions

Shared subsystem: Modules & extensions.

Lesson ID: `unit-editor-plugin-extension-scripting-scene-contexts`
Display slug: `unit-103-script-resources-and-instances`

### Domain: Editor plug-in extension — First-party

EditorPlugin is the integration surface used by the editor’s built-in tool plug-ins as well as external extensions.

Plugins attach behavior through polymorphic C++ hook interfaces and may register import, scene-format, post-import, inspector, gizmo, debugger, and resource-conversion plugins.

Read more: [Editor plug-in extension](references/domain-concepts.md#topic-editor-plugin-extension)

### Domain: Script resources and instances — First-party

The native module represents script validity, reload state, members, and compiler/analyzer relationships.

GDScript is a Script Resource whose compiled members and functions supply script instances to Object-derived engine objects.

Read more: [Script resources and instances](references/domain-concepts.md#topic-scripting)

Use these entity examples: GDScript, Node.

### Domain: Scene-aware test context — First-party

Completion cases select this context through a configuration `scene` path.

A scene input gives a completion context containing nodes, attached scripts, resources, and unique names.

Read more: [Scene-aware test context](references/domain-concepts.md#topic-scene-contexts)

Use these entity examples: Scene Fixture, Completion Test Configuration.

## Unit 104 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-webrtc-multiplayer-mesh-android-platform-host`
Display slug: `unit-104-android-platform-host`

### Domain: WebRTC multiplayer mesh — First-party

ConnectedPeer records support the multiplayer abstraction's peer tracking.

A WebRTCMultiplayerPeer keeps a mesh of WebRTC peer connections and their data channels.

Read more: [WebRTC multiplayer mesh](references/domain-concepts.md#topic-webrtc-multiplayer-mesh)

Use these entity examples: WebRTCMultiplayerPeer.

### Domain: Android platform host — First-party

The Android partition has both native C++ platform classes and JVM-facing host classes.

Android platform hosting starts and manages the native engine from Android activity and library code, continuing the engine bootstrap on the mobile host.

Read more: [Android platform host](references/domain-concepts.md#topic-android-platform-host)

## Unit 105 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-rendering-device-resources-target-platform-export-cpp-frame-scheduling`
Display slug: `unit-105-renderingdevice-descriptors-and-handles`

### Domain: RenderingDevice descriptors and handles — First-party

RD-prefixed descriptor objects carry configuration into the low-level rendering API.

RenderingDevice creates and consumes RID handles for buffers, textures, shaders, uniforms, pipelines, framebuffers, and acceleration structures.

Read more: [RenderingDevice descriptors and handles](references/domain-concepts.md#topic-rendering-device-resources)

Use these entity examples: RDTextureFormat, RDUniform, RDAccelerationStructureGeometry.

### Domain: Target-platform export — First-party

The base platform type supplies common packaging facilities while derived platform types supply target behavior.

A target platform implementation supplies target-specific validation, option, run, and project, pack, or ZIP export behavior to export orchestration.

Read more: [Target-platform export](references/domain-concepts.md#topic-target-platform-export)

Use these entity examples: EditorExportPlatform, EditorExportPlatform::ExportMessage.

### Implementation: C++ frame-time arithmetic — First-party

`MainTimerSync` contains the timing policy; `DeltaSmoother` is a separate timing-related class in the same subsystem.

C++ frame timing turns accumulated time and a physics tick rate into bounded physics-step counts for frame timing.

Read more: [C++ frame-time arithmetic](references/language-concepts.md#topic-cpp-frame-scheduling)

Apply it through: MainFrameTime.

## Unit 106 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-inspector-property-editors-managed-native-interop-gdscript-declarations`
Display slug: `unit-106-gdscript-declarations-and-typing`

### Domain: Custom inspector property editors — First-party

This path is used for gradients, materials, textures, particle-process ranges, GUI controls, fonts, and style boxes.

The implementation uses editor-plugin lifecycle integration to replace or extend property editing with specialized EditorInspectorPlugin and EditorProperty classes.

Read more: [Custom inspector property editors](references/domain-concepts.md#topic-inspector-property-editors)

### Domain: Managed-native interop — First-party

The bridge has a size-checked unmanaged callback table and native runtime interop table with matching ordering requirements.

Unsafe C# connects managed values and callbacks across a fixed native/managed ABI using variant conversion, GC handles, and function-pointer callbacks.

Read more: [Managed-native interop](references/domain-concepts.md#topic-managed-native-interop)

Use these entity examples: ManagedCallbacks, Variant.

### Implementation: GDScript declarations and typing — First-party

GDScript is implemented in a dedicated module that models compiled scripts, member information, analysis, and compilation.

GDScript source uses `extends`, `class_name`, typed variables, functions, and annotations; the engine represents each script as a GDScript Resource.

Read more: [GDScript declarations and typing](references/language-concepts.md#topic-gdscript-declarations)

Apply it through: GDScript, Variant.

## Unit 107 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-android-jni-interop-android-render-input-nodepaths-and-indexed-access`
Display slug: `unit-107-android-jni-interoperation`

### Domain: Android JNI interoperation — First-party

Native wrappers, Java bridge classes, and callable adapters form the implemented boundary.

JNI interoperation converts values and directs callbacks across the Android platform host’s Java/native boundary using C++ Variant marshalling.

Read more: [Android JNI interoperation](references/domain-concepts.md#topic-android-jni-interop)

Use these entity examples: Callable.

### Domain: Android rendering and input — First-party

The implementation includes separate GL and Vulkan render-view classes plus a shared Android input handler.

Android rendering views and native rendering drivers carry surface and input work from the Android platform host.

Read more: [Android rendering and input](references/domain-concepts.md#topic-android-render-input)

### Implementation: Node paths and indexed access — First-party

Completion expectations vary with the source expression and with the node and script information available from a scene.

The fixtures use `$`, `%`, and indexed self access to resolve Node members inside a scene-aware test context.

Read more: [Node paths and indexed access](references/language-concepts.md#topic-nodepaths-and-indexed-access)

Apply it through: Scene Fixture, Completion Test Configuration.

## Unit 108 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-java-android-host-api`
Display slug: `unit-108-java-android-host-apis`

### Implementation: Java Android host APIs — First-party

Java classes supply the Android-facing half of rendering, input, plugins, and native-library integration.

Java uses class inheritance and interface implementation to attach GL/Vulkan views, plugin registration, native library calls, and input handling to the Android platform host.

Read more: [Java Android host APIs](references/language-concepts.md#topic-java-android-host-api)

## Unit 109 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-kotlin-runtime-coordination`
Display slug: `unit-109-kotlin-runtime-and-service-coordination`

### Implementation: Kotlin runtime and service coordination — First-party

Kotlin is used for the main `Godot` host, activities, services, remote fragments, storage handlers, and editor code.

Kotlin coordinates the Android platform host through nullable runtime state, Android host construction, and integer-coded remote-service messages.

Read more: [Kotlin runtime and service coordination](references/language-concepts.md#topic-kotlin-runtime-coordination)

Apply it through: Callable, GodotService.EngineStatus.

## Unit 110 — Documentation & reference

Shared subsystem: Documentation & reference.

Lesson ID: `unit-scene-data-and-buffers-renderer-storage-gdscript-signals-callables`
Display slug: `unit-110-gdscript-signals-and-callbacks`

### Domain: Scene render data and buffers — First-party

The renderer receives a compact frame-oriented view rather than owning scene objects directly.

RenderDataRD gathers visible scene-instance pointers and RID lists for lights, probes, decals, lightmaps, and fog volumes, while RenderSceneBuffersRD supplies named GPU scene textures.

Read more: [Scene render data and buffers](references/domain-concepts.md#topic-scene-data-and-buffers)

Use these entity examples: RenderDataRD.

### Domain: Renderer-owned resource storage — First-party

Storage isolates resource lifetime and backend-specific data from scene traversal and draw submission.

The RD storage services own backend representations of lights, materials, meshes, particles, textures, and GPU resources addressed by rendering IDs.

Read more: [Renderer-owned resource storage](references/domain-concepts.md#topic-renderer-storage)

### Implementation: GDScript signals and callbacks — First-party

This is the idiomatic script-facing bridge to Godot’s signal and input-dispatch systems.

GDScript connects a Signal to a Callable and uses callback functions such as `_input` in Object-derived nodes.

Read more: [GDScript signals and callbacks](references/language-concepts.md#topic-gdscript-signals-callables)

Apply it through: Node, Variant.

## Unit 111 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-apple-embedded-packaging-android-plugin-signals-android-storage-bridge`
Display slug: `unit-111-android-plugins-and-signals`

### Domain: Apple embedded packaging and signing — First-party

The Apple exporter extends the common platform exporter and uses dedicated signing and plugin-configuration types.

Apple embedded packaging adds Xcode project data, assets, Apple plugin configuration, and code-signing structures to a target-platform export.

Read more: [Apple embedded packaging and signing](references/domain-concepts.md#topic-apple-embedded-packaging)

Use these entity examples: EditorExportPlatform, EditorExportPlatform::ExportMessage.

### Domain: Android plugins and signals — First-party

The test app registers plugin classes in its manifest and GDScript tests retrieve and connect to plugin signals.

Android plugins are discovered from application metadata and expose annotated methods and signals through JNI interoperation.

Read more: [Android plugins and signals](references/domain-concepts.md#topic-android-plugin-signals)

Use these entity examples: SignalInfo.

### Domain: Android storage bridge — First-party

The Java/Kotlin implementation has separate access types rather than a single Android storage path.

Android file and directory access handlers implement asset, filesystem, MediaStore, and SAF paths through JNI interoperation.

Read more: [Android storage bridge](references/domain-concepts.md#topic-android-storage-bridge)

## Unit 112 — Platforms & delivery

Shared subsystem: Platforms & delivery.

Lesson ID: `unit-screen-and-environment-effects-android-instrumented-tests-gdscript-plugin-integration`
Display slug: `unit-112-android-instrumented-integration-tests`

### Domain: Screen-space and environment effects — First-party

These classes are effect-oriented services layered over scene buffers and shader pipelines.

RD effects and environment services allocate and process scene buffers for luminance, temporal anti-aliasing, reflections, ambient effects, fog, global illumination, sky, tone mapping, and scaling.

Read more: [Screen-space and environment effects](references/domain-concepts.md#topic-screen-and-environment-effects)

### Domain: Android instrumented integration tests — First-party

The Android test source includes Kotlin test code, Java/Kotlin plugin fixtures, and a Godot project with GDScript tests.

Android instrumentation tests launch an app project that exercises plugins, signals, storage, and JavaClassWrapper conversions.

Read more: [Android instrumented integration tests](references/domain-concepts.md#topic-android-instrumented-tests)

### Implementation: GDScript Android plugin integration — First-party

The instrumented Godot project is a runnable scene with test scripts for Android plugins, storage, and Java class wrapping.

GDScript test code retrieves JNISingleton objects, calls JavaClassWrapper, and connects signals through Android plugin signals.

Read more: [GDScript Android plugin integration](references/language-concepts.md#topic-gdscript-plugin-integration)

Apply it through: SignalInfo.

## Unit 113 — Vendored: amd-fsr2

Shared subsystem: Vendored: amd-fsr2.

Lesson ID: `unit-temporal-upscaling-dispatch`
Display slug: `unit-113-temporal-upscaling-dispatch`

### Domain: Temporal upscaling dispatch — Vendored: amd-fsr2

The implementation contains temporal resource selection and named shader-pass sources.

Vendored FSR2 dispatch code selects alternating frame resources, handles accumulation reset, and issues compute work based on render and display dimensions.

Read more: [Temporal upscaling dispatch](references/domain-concepts.md#topic-temporal-upscaling-dispatch)

## Unit 114 — Vendored: basis_universal

Shared subsystem: Vendored: basis_universal.

Lesson ID: `unit-basis-transcoding-cpp-byte-exact-binary-parsing`
Display slug: `unit-114-c-byte-exact-binary-parsing`

### Domain: Basis texture transcoding — Vendored: basis_universal

The implementation exposes a container-level transcoder plus specialized ETC1S, UASTC, and ASTC paths.

The Basis Universal subsystem decodes ETC1S and UASTC texture slices into selected texture formats through high-level and low-level transcoders.

Read more: [Basis texture transcoding](references/domain-concepts.md#topic-basis-transcoding)

Use these entity examples: BasisFileHeader, BasisSliceDescriptor.

### Implementation: C++ byte-exact binary parsing — Vendored: basis_universal

The texture containers deliberately operate on byte-addressed buffers and packed integer fields.

This C++ code reads byte streams through packed fields, pointer casts, and bit readers; this is necessary to map serialized texture headers to in-memory values.

Read more: [C++ byte-exact binary parsing](references/language-concepts.md#topic-cpp-byte-exact-binary-parsing)

Apply it through: BasisFileHeader, BasisSliceDescriptor, Ktx2Header, Ktx2LevelIndex.

## Unit 115 — Vendored: basis_universal

Shared subsystem: Vendored: basis_universal.

Lesson ID: `unit-basis-container-layout-cpp-abstraction-and-polymorphism`
Display slug: `unit-115-c-types-encapsulation-and-inheritance`

### Domain: Basis file layout — Vendored: basis_universal

The transcoder reads a header followed by an array of slice descriptions addressed through header offsets.

A .basis file header and slice descriptors locate compressed slices and identify their texture form before Basis Universal decodes them.

Read more: [Basis file layout](references/domain-concepts.md#topic-basis-container-layout)

Use these entity examples: BasisFileHeader, BasisSliceDescriptor.

### Implementation: C++ types, encapsulation, and inheritance — Vendored: basis_universal

The partition uses classes as stable subsystem boundaries and exposes operations through member functions.

This C++ code models decoders, geometry engines, and allocators as named types with inheritance and encapsulated state.

Read more: [C++ types, encapsulation, and inheritance](references/language-concepts.md#topic-cpp-abstraction-and-polymorphism)

Apply it through: Ktx2TranscoderState, PolyPathD.

## Unit 116 — Vendored: basis_universal

Shared subsystem: Vendored: basis_universal.

Lesson ID: `unit-prefix-code-decoding-ktx2-container-transcoding-cpp-templates-and-generic-containers`
Display slug: `unit-116-bitstream-and-huffman-decoding`

### Domain: Bitstream and Huffman decoding — Vendored: basis_universal

This is the shared compression idiom visible in Basis internal decoding and Brotli decoder code.

Both Basis and Brotli implement bit readers and Huffman decoding tables to recover symbols from compressed bit streams.

Read more: [Bitstream and Huffman decoding](references/domain-concepts.md#topic-prefix-code-decoding)

### Domain: KTX2 container transcoding — Vendored: basis_universal

KTX2 support is an optional container path alongside direct .basis transcoding.

The KTX2 transcoder retains the source data, parses the KTX2 header and per-level index, then transcodes its texture levels.

Read more: [KTX2 container transcoding](references/domain-concepts.md#topic-ktx2-container-transcoding)

Use these entity examples: Ktx2Header, Ktx2LevelIndex, Ktx2TranscoderState.

### Implementation: C++ templates and generic containers — Vendored: basis_universal

Basis containers, Clipper points and paths, and Embree math types use type parameters rather than duplicate implementations per element type.

This C++ code parameterizes collection and math utilities by type, so their typed operations can be reused.

Read more: [C++ templates and generic containers](references/language-concepts.md#topic-cpp-templates-and-generic-containers)

Apply it through: Ktx2LevelIndex, PolyPathD.

## Unit 117 — Vendored: basis_universal

Shared subsystem: Vendored: basis_universal.

Lesson ID: `unit-astc-block-coding-block-texture-transcoding-raster-image-input`
Display slug: `unit-117-astc-block-coding`

### Domain: ASTC block coding — Vendored: basis_universal

ASTC logic is used both by the Basis transcoder and its HDR intermediate paths.

ASTC helpers represent physical and logical blocks with endpoint ranges, weight grids, partitions, and bit-level packing.

Read more: [ASTC block coding](references/domain-concepts.md#topic-astc-block-coding)

### Domain: GPU block texture conversion — Vendored: basis_universal

The transcoders support a range of destination block formats and uncompressed pixel formats.

Texture conversion code maps ETC1S, UASTC, and ASTC blocks to GPU block or uncompressed target formats, calculating output block counts and storage.

Read more: [GPU block texture conversion](references/domain-concepts.md#topic-block-texture-transcoding)

### Domain: Raster image input — Vendored: basis_universal

The Basis encoder dependency includes separate PNG and JPEG decoding implementations.

Raster image input adapters decode PNG and JPEG bytes into image buffers for texture-processing callers.

Read more: [Raster image input](references/domain-concepts.md#topic-raster-image-input)

Use these entity examples: PngInfo.

## Unit 118 — Vendored: basis_universal

Shared subsystem: Vendored: basis_universal.

Lesson ID: `unit-texture-compression-pipeline`
Display slug: `unit-118-texture-compression-pipeline`

### Domain: Texture compression pipeline — Vendored: basis_universal

The supplied implementation includes ETC-family, UASTC, ASTC, GPU texture, and container-output code.

Vendored Basis Universal code separates frontend block/codebook work, backend encoding, and Basis or KTX2 output creation.

Read more: [Texture compression pipeline](references/domain-concepts.md#topic-texture-compression-pipeline)

## Unit 119 — Vendored: brotli

Shared subsystem: Vendored: brotli.

Lesson ID: `unit-brotli-stream-decompression-c-stateful-streaming-api`
Display slug: `unit-119-brotli-stream-decompression`

### Domain: Brotli stream decompression — Vendored: brotli

Its public decoding path is stateful and can return output incrementally.

The Brotli decoder consumes a guarded byte reader, decodes Huffman and prefix-coded streams, and exposes output from a decoder state.

Read more: [Brotli stream decompression](references/domain-concepts.md#topic-brotli-stream-decompression)

Use these entity examples: BrotliDecoderState.

### Implementation: C stateful streaming APIs — Vendored: brotli

Brotli exposes decoder state through C interfaces while storing its operational fields in an internal structure.

C-facing code represents decoder progress as pointer-rich state and accepts caller-supplied allocation callbacks.

Read more: [C stateful streaming APIs](references/language-concepts.md#topic-c-stateful-streaming-api)

Apply it through: BrotliDecoderState.

## Unit 120 — Vendored: clipper2

Shared subsystem: Vendored: clipper2.

Lesson ID: `unit-polygon-clipping`
Display slug: `unit-120-polygon-clipping-and-result-trees`

### Domain: Polygon clipping and result trees — Vendored: clipper2

Integer and scaled floating-point front ends share the ClipperBase execution machinery.

Clipper accepts subject and clip paths, runs a specified clip and fill rule, and builds closed or open paths or a hierarchy.

Read more: [Polygon clipping and result trees](references/domain-concepts.md#topic-polygon-clipping)

Use these entity examples: PolyPathD.

## Unit 121 — Vendored: cvtt

Shared subsystem: Vendored: cvtt.

Lesson ID: `unit-block-texture-encoding-cpp-simd-intrinsics`
Display slug: `unit-121-c-simd-wrappers-and-intrinsics`

### Domain: Block texture encoding — Vendored: cvtt

This is a separate encoder subsystem from Basis transcoding.

CVTT compression input uses pixel blocks, options, encoding plans, and format-specific BC6H, BC7, ETC, and S3TC computation.

Read more: [Block texture encoding](references/domain-concepts.md#topic-block-texture-encoding)

### Implementation: C++ SIMD wrappers and intrinsics — Vendored: cvtt

Embree wraps x86 and ARM SIMD operations, while CVTT selects SSE2 support at preprocessing time.

C++ vector wrappers invoke architecture intrinsics to operate on multiple lane values.

Read more: [C++ SIMD wrappers and intrinsics](references/language-concepts.md#topic-cpp-simd-intrinsics)

## Unit 122 — Vendored: d3d12ma

Shared subsystem: Vendored: d3d12ma.

Lesson ID: `unit-gpu-memory-suballocation`
Display slug: `unit-122-d3d12-memory-allocation`

### Domain: D3D12 memory allocation — Vendored: d3d12ma

The subsystem includes block metadata, pools, budgets, defragmentation, and virtual allocations.

D3D12MA's Allocator, Pool, and VirtualBlock APIs manage resource and virtual allocations using allocation callbacks.

Read more: [D3D12 memory allocation](references/domain-concepts.md#topic-gpu-memory-suballocation)

## Unit 123 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-bvh-construction-geometry-resources-cpp-class-templates`
Display slug: `unit-123-bvh-construction`

### Domain: BVH construction — Vendored: embree

The generic builder evaluates leaf and split costs, while factory and specialized builders select node and primitive representations.

BVH construction partitions PrimRef records into nodes and leaves through configurable builder callbacks and settings.

Read more: [BVH construction](references/domain-concepts.md#topic-bvh-construction)

Use these entity examples: BVHN, GeneralBVHBuilder, BVHNodeRecord.

### Domain: Geometry resources — Vendored: embree

Geometry is the scene content abstraction that builders and intersectors address by geometry ID.

A Geometry represents typed primitive data, supplies bounds, and provides the common base used by mesh, curve, point, line, grid, user, and instance implementations.

Read more: [Geometry resources](references/domain-concepts.md#topic-geometry-resources)

Use these entity examples: RTCGeometry, Geometry.

### Implementation: C++ class templates and specialization — Vendored: embree

Embree passes the lane-width parameter through vector types such as Vec3vf<K> and specializes subgrid intersectors for width four.

C++ class templates parameterize Embree vector and intersector code by lane width, and a fixed-width specialization selects a width-specific implementation; C++ classes are needed to explain class templates.

Read more: [C++ class templates and specialization](references/language-concepts.md#topic-cpp-class-templates)

Apply it through: TriangleMi, SubGrid.

## Unit 124 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-ray-query-cxx-wide-simd`
Display slug: `unit-124-ray-query-state`

### Domain: Ray query state — Vendored: embree

The API provides scalar rays plus 4-, 8-, 16-, and compile-time-N packet forms.

A ray query carries origin, direction, near and far distances, time, mask, ID, and flags for intersection or occlusion traversal.

Read more: [Ray query state](references/domain-concepts.md#topic-ray-query)

Use these entity examples: RTCRay, RTCRayHit.

### Implementation: C++ SIMD wrapper specialization — Vendored: embree

The same vector-oriented algorithms are instantiated against widths such as four, eight, and sixteen lanes, with platform-specific definitions behind the wrapper names.

C++ templates and wrapper types represent fixed-width integer, unsigned-integer, float, mask, and long-vector lanes while ISA-specific headers implement their operations.

Read more: [C++ SIMD wrapper specialization](references/language-concepts.md#topic-cxx-wide-simd)

Apply it through: vint4, vint8, vuint16.

## Unit 125 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-hit-results-cxx-closures`
Display slug: `unit-125-intersection-results`

### Domain: Intersection results — Vendored: embree

The intersection distance is maintained through the ray state used by traversal and intersector epilogues.

An RTCRayHit combines a ray query with geometric-normal, barycentric-coordinate, primitive-ID, geometry-ID, and instance-stack hit results.

Read more: [Intersection results](references/domain-concepts.md#topic-hit-results)

Use these entity examples: RTCHit, RTCRayHit.

### Implementation: C++ lambdas and callback objects — Vendored: embree

This lets generic builder code operate on caller-selected node and leaf representations without hard-coding a single output format.

C++ lambdas and callable objects bind BVH builders to allocation, node creation, leaf creation, progress, and completion behavior.

Read more: [C++ lambdas and callback objects](references/language-concepts.md#topic-cxx-closures)

Apply it through: ProgressMonitorClosure, GeneralBVHBuilder, BVHNBuilderVirtual.

## Unit 126 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-scene-commit-triangle-intersection-algorithms-cxx-c-abi`
Display slug: `unit-126-scene-construction-and-commit`

### Domain: Scene construction and commit — Vendored: embree

The scene layer also checks whether attached geometry has changed before marking itself modified.

A Scene is created from a Device, retains Geometry instances, and commits a BVH used by query calls.

Read more: [Scene construction and commit](references/domain-concepts.md#topic-scene-commit)

Use these entity examples: RTCScene, Scene.

### Domain: Triangle intersection algorithms — Vendored: embree

Triangle intersection uses Möller–Trumbore, Plücker, and Woop kernels; ray-primitive intersection is required to interpret their hit candidates.

Read more: [Triangle intersection algorithms](references/domain-concepts.md#topic-triangle-intersection-algorithms)

Use these entity examples: TriangleMi.

### Implementation: C++ linkage and opaque API handles — Vendored: embree

Implementation classes remain internal while API calls exchange opaque handle types and plain structures such as RTCRayHit.

C++ exposes RTCDevice, RTCScene, RTCGeometry, and RTCBuffer as opaque pointer handles through a C-linkage public boundary.

Read more: [C++ linkage and opaque API handles](references/language-concepts.md#topic-cxx-c-abi)

Apply it through: RTCDevice, RTCScene, RTCGeometry, RTCBuffer.

## Unit 127 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-motion-blur-cxx-conditional-compilation`
Display slug: `unit-127-motion-blur`

### Domain: Motion blur — Vendored: embree

The implementation has linear-bound primitive references, motion-blur builders, time-range node forms, and time-segment geometry access.

Motion-blur Geometry and BVH nodes interpolate bounds at ray-query time.

Read more: [Motion blur](references/domain-concepts.md#topic-motion-blur)

Use these entity examples: PrimRefMB, AABBNodeMB_t, AABBNodeMB4D_t.

### Implementation: C++ preprocessor configuration — Vendored: embree

The configuration is compile-time: the resulting build exposes only code paths enabled by its macro definitions.

C++ preprocessing selects tasking backends, supported geometry features, ISA namespaces, platform headers, and fallback shims before compilation.

Read more: [C++ preprocessor configuration](references/language-concepts.md#topic-cxx-conditional-compilation)

Apply it through: TaskScheduler, Geometry, vint8.

## Unit 128 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-primitive-intersection-cxx-templates`
Display slug: `unit-128-primitive-intersection`

### Domain: Primitive intersection — Vendored: embree

The indexed code includes triangle, quad, curve, line, cone, cylinder, disc, grid, object, and instance intersectors.

Primitive intersectors test a ray query against Geometry representations and update RTCRayHit through common intersection or occlusion epilogues.

Read more: [Primitive intersection](references/domain-concepts.md#topic-primitive-intersection)

Use these entity examples: TriangleMIntersectorKMoeller, QuadMIntersector1MoellerTrumbore, ConeCurveIntersector1.

### Implementation: C++ templates and specialization — Vendored: embree

The implementation uses type and non-type template parameters instead of a single runtime-polymorphic kernel for all widths and primitive forms.

C++ templates parameterize BVH branching, packet width, vector width, primitive layout, and builder behavior at compile time.

Read more: [C++ templates and specialization](references/language-concepts.md#topic-cxx-templates)

Apply it through: BVHN, RTCRayNt, MortonBuilder.

## Unit 129 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-primitive-references-cxx-atomics`
Display slug: `unit-129-primitive-references`

### Domain: Primitive references — Vendored: embree

Motion-blur variants carry linear bounds and time-segment information.

A PrimRef supplies bounds plus geometry and primitive identifiers so builders can partition scene primitives independently of their source geometry layout.

Read more: [Primitive references](references/domain-concepts.md#topic-primitive-references)

Use these entity examples: PrimRef, PrimRefMB, SubGridBuildData.

### Implementation: C++ atomic synchronization — Vendored: embree

The implementation wraps std::atomic and uses it directly in barrier and allocator-adjacent synchronization paths.

C++ atomic counters and compare-exchange loops coordinate active barriers and conditional minimum or maximum updates across workers.

Read more: [C++ atomic synchronization](references/language-concepts.md#topic-cxx-atomics)

Apply it through: atomic, BarrierActive, BarrierActiveAutoReset.

## Unit 130 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-builder-heuristics-bvh-traversal-cxx-class-hierarchy`
Display slug: `unit-130-bvh-split-heuristics`

### Domain: BVH split heuristics — Vendored: embree

The supplied builders calculate costs from bounds, primitive counts, block size, traversal cost, and intersection cost.

SAH, Morton, spatial, temporal, strand, open-merge, and motion-blur heuristics choose how PrimRef records are divided while constructing a BVH.

Read more: [BVH split heuristics](references/domain-concepts.md#topic-builder-heuristics)

Use these entity examples: BinMapping, MortonCodeMapping, SpatialBinMapping, HeuristicMBlurTemporalSplit.

### Domain: BVH traversal — Vendored: embree

The code includes scalar, packet, hybrid, frustum, point-query, quantized-node, and oriented-node traversal paths.

Traversal tests the committed BVH with a ray query, orders candidate nodes, and calls primitive intersectors for reached leaves.

Read more: [BVH traversal](references/domain-concepts.md#topic-bvh-traversal)

Use these entity examples: BVHNIntersector1, BVHNIntersectorKHybrid, TravRay.

### Implementation: C++ class hierarchies and reference bases — Vendored: embree

The code combines reusable base classes with specialized derived classes rather than representing all runtime objects as unrelated records.

C++ class inheritance connects shared reference ownership to acceleration, scene, geometry, builder, and scheduler abstractions.

Read more: [C++ class hierarchies and reference bases](references/language-concepts.md#topic-cxx-class-hierarchy)

Apply it through: RefCount, AccelData, Geometry, Builder.

## Unit 131 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-device-runtime-instancing-packet-queries`
Display slug: `unit-131-device-runtime`

### Domain: Device runtime — Vendored: embree

It is the API-level runtime owner behind scene and buffer creation.

A Device configures Embree execution, exposes device and thread error messages, and limits tasking resources.

Read more: [Device runtime](references/domain-concepts.md#topic-device-runtime)

Use these entity examples: RTCDevice, Device.

### Domain: Instancing — Vendored: embree

The configured API sets RTC_MAX_INSTANCE_LEVEL_COUNT to one.

Instance and InstanceArray are Geometry types attached to a Scene that contribute transformed primitive references and maintain instance-query context state.

Read more: [Instancing](references/domain-concepts.md#topic-instancing)

Use these entity examples: Instance, InstanceArray, InstancePrimitive.

### Domain: Ray packets and streams — Vendored: embree

Packet traversal and packet primitive intersectors reuse the same query concepts across multiple lanes.

Packet execution packs several ray queries and their hit results into 4-, 8-, 16-, or N-wide layouts, with array-of-structures and structure-of-arrays stream readers.

Read more: [Ray packets and streams](references/domain-concepts.md#topic-packet-queries)

Use these entity examples: RTCRay4, RTCRayHit16, RTCRayNt, RayStreamSOA.

## Unit 132 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-simd-ray-packets-subdivision-surface-evaluation-catmull-clark-patch-construction`
Display slug: `unit-132-simd-ray-packets`

### Domain: SIMD ray packets — Vendored: embree

SIMD ray packets evaluate width-parameterized ray and triangle lanes, so triangle intersection is required to interpret lane-validity masks.

Read more: [SIMD ray packets](references/domain-concepts.md#topic-simd-ray-packets)

Use these entity examples: TriangleMi.

### Domain: Subdivision surface evaluation — Vendored: embree

Subdivision evaluation combines curve and patch bases with grid sampling; curve and patch bases are needed to explain the evaluator.

Read more: [Subdivision surface evaluation](references/domain-concepts.md#topic-subdivision-surface-evaluation)

Use these entity examples: SubGrid, GridMesh.

### Domain: Catmull–Clark patch construction — Vendored: embree

Catmull–Clark patch construction reads a half-edge topology and converts its limit data to patch geometry; half-edge topology is needed to explain this construction.

Read more: [Catmull–Clark patch construction](references/domain-concepts.md#topic-catmull-clark-patch-construction)

Use these entity examples: HalfEdge.

## Unit 133 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-feature-adaptive-tessellation-geometry-families-half-edge-topology`
Display slug: `unit-133-feature-adaptive-tessellation`

### Domain: Feature-adaptive tessellation — Vendored: embree

Feature-adaptive tessellation recursively partitions parameter ranges for patch evaluation; subdivision evaluation is needed to explain the patch evaluation.

Read more: [Feature-adaptive tessellation](references/domain-concepts.md#topic-feature-adaptive-tessellation)

Use these entity examples: SubGrid.

### Domain: Geometry families — Vendored: embree

The checked configuration directly enables triangle geometry, while the indexed source also contains the other family implementations.

The source defines Geometry subclasses for triangle meshes, quad meshes, curves, line segments, points, grids, subdivision meshes, user geometry, and instances.

Read more: [Geometry families](references/domain-concepts.md#topic-geometry-families)

Use these entity examples: TriangleMesh, QuadMesh, CurveGeometry, LineSegments, Points, GridMesh, SubdivMesh, UserGeometry.

### Domain: Half-edge topology — Vendored: embree

Subdivision connectivity is represented by HalfEdge navigation and by Catmull–Clark rings that inspect neighboring edges, faces, and crease information.

Read more: [Half-edge topology](references/domain-concepts.md#topic-half-edge-topology)

Use these entity examples: HalfEdge.

## Unit 134 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-isa-portability-filter-callbacks-motion-blur-geometry`
Display slug: `unit-134-isa-and-platform-portability`

### Domain: ISA and platform portability — Vendored: embree

The source contains SSE2, AVX, AVX2, AVX-512, ARM, and WebAssembly-oriented paths.

Preprocessor-selected ISA configuration chooses SIMD namespaces and headers, while platform shims adapt unavailable WebAssembly control-register operations.

Read more: [ISA and platform portability](references/domain-concepts.md#topic-isa-portability)

Use these entity examples: vint4, vint8, vuint16.

### Domain: Intersection and occlusion filters — Vendored: embree

User geometry also exposes externally supplied intersect and occluded function argument structures.

Geometry-installed and query-context filters process candidate hit results before they update an RTCRayHit or report occlusion.

Read more: [Intersection and occlusion filters](references/domain-concepts.md#topic-filter-callbacks)

Use these entity examples: RTCFilterFunctionNArguments, RTCIntersectFunctionNArguments, RTCOccludedFunctionNArguments.

### Domain: Motion-blur geometry — Vendored: embree

Motion-blur triangle code computes vertices at ray time before triangle intersection, so triangle intersection is required to interpret its candidates.

Read more: [Motion-blur geometry](references/domain-concepts.md#topic-motion-blur-geometry)

Use these entity examples: TriangleMi.

## Unit 135 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-curve-and-patch-bases-ray-primitive-intersection-subgrid-intersection`
Display slug: `unit-135-parametric-curve-bases`

### Domain: Parametric curve bases — Vendored: embree

These basis implementations are reused by the subdivision-side curve and patch evaluators.

Bezier, B-spline, and Catmull-Rom basis evaluators generate positions and derivatives from four control points.

Read more: [Parametric curve bases](references/domain-concepts.md#topic-curve-and-patch-bases)

### Domain: Ray–primitive intersection — Vendored: embree

Embree provides separate kernels for several primitive families while retaining the same hit-candidate flow.

Ray-primitive intersection evaluates rays against triangle, sphere, and rounded-line geometry and forwards valid hit candidates to an epilog.

Read more: [Ray–primitive intersection](references/domain-concepts.md#topic-ray-primitive-intersection)

Use these entity examples: SubGrid, TriangleMi.

### Domain: Subgrid intersection — Vendored: embree

Subgrid intersection uses grid identifiers to load a quad neighborhood and applies triangle intersection algorithms to its triangles; triangle intersection algorithms are needed to interpret those tests.

Read more: [Subgrid intersection](references/domain-concepts.md#topic-subgrid-intersection)

Use these entity examples: SubGrid, GridMesh.

## Unit 136 — Vendored: embree

Shared subsystem: Vendored: embree.

Lesson ID: `unit-allocation-buffer-storage-task-scheduling`
Display slug: `unit-136-allocation-and-reference-ownership`

### Domain: Allocation and reference ownership — Vendored: embree

Builders receive allocator callbacks, while shared runtime objects inherit or use RefCount-based ownership.

Reference-counted objects, FastAllocator blocks, cached allocation, aligned allocation, and monitored allocation support resource and BVH memory management.

Read more: [Allocation and reference ownership](references/domain-concepts.md#topic-allocation)

Use these entity examples: RefCount, Ref, FastAllocator, CachedAllocator.

### Domain: Geometry buffer storage — Vendored: embree

The public API also defines opaque RTCBuffer handles and buffer-type slots for indices, vertices, attributes, normals, transforms, and related data.

A Buffer is a reference-counted storage abstraction, while BufferView and RawBufferView expose views used to bind geometry data.

Read more: [Geometry buffer storage](references/domain-concepts.md#topic-buffer-storage)

Use these entity examples: RTCBuffer, Buffer.

### Domain: Task scheduling and synchronization — Vendored: embree

The scheduler can be selected between internal, TBB, and PPL implementations through configuration macros.

The internal scheduler represents work as tasks, queues, threads, and thread pools, and provides barriers and atomic helpers for synchronization.

Read more: [Task scheduling and synchronization](references/domain-concepts.md#topic-task-scheduling)

Use these entity examples: TaskScheduler, Task, TaskQueue, ThreadPool, BarrierActive.

## Unit 137 — Vendored: enet

Shared subsystem: Vendored: enet.

Lesson ID: `unit-enet-wire-commands-c-function-pointer-callbacks`
Display slug: `unit-137-enet-wire-commands`

### Domain: ENet wire commands — Vendored: enet

ENet wire commands select fixed protocol layout sizes and drive peer state transitions; host-peer transport is required to interpret the peer state.

Read more: [ENet wire commands](references/domain-concepts.md#topic-enet-wire-commands)

Use these entity examples: ENetPeer.

### Implementation: C function-pointer callbacks — Vendored: enet

The ENet callback record is initialized with standard allocation functions and can be replaced through its initialization API.

C function-pointer fields let ENet obtain allocator, deallocator, and out-of-memory handlers through a cross-cutting callbacks record.

Read more: [C function-pointer callbacks](references/language-concepts.md#topic-c-function-pointer-callbacks)

Apply it through: ENetCallbacks.

## Unit 138 — Vendored: enet

Shared subsystem: Vendored: enet.

Lesson ID: `unit-godot-enet-socket-adaptation-enet-host-peer-transport-wraparound-network-time`
Display slug: `unit-138-godot-enet-socket-adaptation`

### Domain: Godot ENet socket adaptation — Vendored: enet

The Godot ENet socket adapter supplies UDP and DTLS socket classes to the transport; ENet wire commands provide this adapter's transport-facing purpose.

Read more: [Godot ENet socket adaptation](references/domain-concepts.md#topic-godot-enet-socket-adaptation)

### Domain: ENet host and peer transport — Vendored: enet

ENet creates a host with a bounded peer allocation and manages peer state changes when connecting or disconnecting.

Read more: [ENet host and peer transport](references/domain-concepts.md#topic-enet-host-peer-transport)

Use these entity examples: ENetHost, ENetPeer.

### Domain: Wraparound network time — Vendored: enet

ENet compares and subtracts time values with an overflow threshold so ordering and differences remain defined across its configured time wraparound.

Read more: [Wraparound network time](references/domain-concepts.md#topic-wraparound-network-time)

## Unit 139 — Vendored: etcpak

Shared subsystem: Vendored: etcpak.

Lesson ID: `unit-codec-simd-specialization-texture-block-codecs`
Display slug: `unit-139-codec-simd-specialization`

### Domain: Codec SIMD specialization — Vendored: etcpak

Conditional C preprocessing exposes AVX2, SSE4.1, and NEON codec declarations; C preprocessing is needed to explain which declarations are available.

Read more: [Codec SIMD specialization](references/domain-concepts.md#topic-codec-simd-specialization)

### Domain: Texture block codecs — Vendored: etcpak

etcpak exposes ETC/EAC and BC block compression plus block-decoding entry points that operate on typed source and destination buffers.

Read more: [Texture block codecs](references/domain-concepts.md#topic-texture-block-codecs)

## Unit 140 — Vendored: freetype

Shared subsystem: Vendored: freetype.

Lesson ID: `unit-font-driver-modules-sfnt-tables-c-preprocessor-composition`
Display slug: `unit-140-font-driver-modules`

### Domain: Font driver modules — Vendored: freetype

Each format family has a driver class and generally separates driver, object, loader, and glyph-loading concerns.

Format-specific driver classes expose CFF, CID, PCF, PFR, Type 1, Type 42, Windows FNT, and TrueType implementations to the FreeType module layer.

Read more: [Font driver modules](references/domain-concepts.md#topic-font-driver-modules)

Use these entity examples: PFR_FaceRec.

### Domain: SFNT table loading — Vendored: freetype

The component is shared infrastructure for OpenType and TrueType-related font handling.

SFNT services read structured tables from font streams, including cmap, metrics, embedded bitmaps, color layers, SVG documents, and WOFF/WOFF2 data.

Read more: [SFNT table loading](references/domain-concepts.md#topic-sfnt-tables)

Use these entity examples: FT_Stream, PFR_FaceRec.

### Implementation: C preprocessor composition — Vendored: freetype

FreeType uses this source-level composition pattern for both its base and autofit module entry translation units.

C preprocessing defines FT_MAKE_OPTION_SINGLE_OBJECT and then includes component implementation files to form a single-object library unit.

Read more: [C preprocessor composition](references/language-concepts.md#topic-c-preprocessor-composition)

## Unit 141 — Vendored: freetype

Shared subsystem: Vendored: freetype.

Lesson ID: `unit-c-structures-and-pointers`
Display slug: `unit-141-c-structures-and-pointer-handles`

### Implementation: C structures and pointer handles — Vendored: freetype

FreeType records commonly embed a base record and keep state behind pointer typedefs.

C structures and pointers package mutable subsystem state and expose opaque handle types.

Read more: [C structures and pointer handles](references/language-concepts.md#topic-c-structures-and-pointers)

Apply it through: FT_BZip2FileRec, PFR_FaceRec, SVG_RendererRec.

## Unit 142 — Vendored: freetype

Shared subsystem: Vendored: freetype.

Lesson ID: `unit-font-streams-c-explicit-resource-lifecycles`
Display slug: `unit-142-font-streams`

### Domain: Font streams — Vendored: freetype

This is the byte-access boundary used by the supplied compression and font-reading code.

Font input is represented by an FT_Stream that can deliver bytes through an in-memory base or a read callback.

Read more: [Font streams](references/domain-concepts.md#topic-font-streams)

Use these entity examples: FT_Stream, FT_BZip2FileRec.

### Implementation: C explicit resource lifecycles — Vendored: freetype

The Bzip2 adapter names initialization, reset, finalization, I/O, and close functions separately and wires them into FT_Stream.

C functions establish, reset, and release C struct state explicitly through paired lifecycle operations.

Read more: [C explicit resource lifecycles](references/language-concepts.md#topic-c-explicit-resource-lifecycles)

Apply it through: FT_Stream, FT_BZip2FileRec.

## Unit 143 — Vendored: freetype

Shared subsystem: Vendored: freetype.

Lesson ID: `unit-glyph-rasterization-c-function-pointer-interfaces`
Display slug: `unit-143-glyph-rasterization`

### Domain: Glyph rasterization — Vendored: freetype

The raster and smooth components expose raster-function objects, while other components provide specialized renderers.

Renderer modules turn driver-loaded glyph outlines into monochrome, gray, SDF, or SVG-backed bitmap representations.

Read more: [Glyph rasterization](references/domain-concepts.md#topic-glyph-rasterization)

Use these entity examples: SVG_RendererRec.

### Implementation: C function-pointer interfaces — Vendored: freetype

The implementation assigns allocation and I/O functions into external library or stream interface records.

C function pointers attach behavior to struct-and-pointer handles without requiring a language-level object system.

Read more: [C function-pointer interfaces](references/language-concepts.md#topic-c-function-pointer-interfaces)

Apply it through: FT_Stream, FT_BZip2FileRec.

## Unit 144 — Vendored: freetype

Shared subsystem: Vendored: freetype.

Lesson ID: `unit-compressed-font-stream-adapters-freetype-module-composition-glyph-caching`
Display slug: `unit-144-compressed-font-stream-adapters`

### Domain: Compressed font stream adapters — Vendored: freetype

The Bzip2 adapter checks the header, initializes a decompressor, resets on backward seeks, and installs read and close callbacks.

Optional Bzip2 and LZW components wrap a compatible source stream and expose decompressed bytes through the same FT_Stream callbacks.

Read more: [Compressed font stream adapters](references/domain-concepts.md#topic-compressed-font-stream-adapters)

Use these entity examples: FT_Stream, FT_BZip2FileRec.

### Domain: FreeType module composition — Vendored: freetype

The FreeType base and autofit modules combine included implementation units into compiled engine modules; C preprocessing is needed to explain the inclusion-based composition.

Read more: [FreeType module composition](references/domain-concepts.md#topic-freetype-module-composition)

### Domain: Glyph and face caching — Vendored: freetype

Cache entries are organized around manager-owned caches, cache nodes, glyph families, and MRU ordering.

The cache manager holds faces made available by format drivers and caches character maps, glyph images, and small-bitmap nodes under resource limits.

Read more: [Glyph and face caching](references/domain-concepts.md#topic-glyph-caching)

Use these entity examples: FTC_SNodeRec.

## Unit 145 — Vendored: freetype

Shared subsystem: Vendored: freetype.

Lesson ID: `unit-font-table-validation-postscript-font-services-signed-distance-fields`
Display slug: `unit-145-opentype-and-truetypegx-validation`

### Domain: OpenType and TrueTypeGX validation — Vendored: freetype

The OpenType module covers BASE, GDEF, GPOS, GSUB, JSTF, and MATH; the GX module contains validators for multiple AAT table families.

OpenType and TrueTypeGX validators check structured tables after SFNT parsing so higher layers can consume checked offsets and indices.

Read more: [OpenType and TrueTypeGX validation](references/domain-concepts.md#topic-font-table-validation)

### Domain: PostScript font services — Vendored: freetype

The supplied sources include Type 1 and Type 2 decoding, CFF decoding, parser tables, and a separate hinter module.

PSAux supplies parsing, decoding, character maps, and hint services shared by CFF, CID, and Type 1 drivers.

Read more: [PostScript font services](references/domain-concepts.md#topic-postscript-font-services)

### Domain: Signed-distance-field glyph rendering — Vendored: freetype

Common SDF support includes distance normalization and shared renderer properties.

The SDF and BSDF renderers use rasterization to populate signed-distance output bitmaps with configurable parameters.

Read more: [Signed-distance-field glyph rendering](references/domain-concepts.md#topic-signed-distance-fields)

## Unit 146 — Vendored: freetype

Shared subsystem: Vendored: freetype.

Lesson ID: `unit-variable-font-data`
Display slug: `unit-146-truetype-gx-variation-data`

### Domain: TrueType GX variation data — Vendored: freetype

The variation loader is compiled with the TrueType driver and models avar correspondence data.

TrueType GX variation loading maps a face's variation tables and coordinates into variable-font state.

Read more: [TrueType GX variation data](references/domain-concepts.md#topic-variable-font-data)

## Unit 147 — Vendored: gamepadmotionhelpers

Shared subsystem: Vendored: gamepadmotionhelpers.

Lesson ID: `unit-gamepad-motion-fusion-cpp-class-state-and-composition`
Display slug: `unit-147-c-class-state-and-composition`

### Domain: Gamepad motion fusion and calibration — Vendored: gamepadmotionhelpers

The implementation uses vector and quaternion operations, smoothing, stillness checks, and calibration confidence settings.

GamepadMotion maintains orientation, gravity, processed acceleration, and automatic gyro-calibration state from gyro and accelerometer samples.

Read more: [Gamepad motion fusion and calibration](references/domain-concepts.md#topic-gamepad-motion-fusion)

### Implementation: C++ class state and composition — Vendored: gamepadmotionhelpers

The GamepadMotion helper is implemented in a C++ header containing class and struct definitions plus their behavior.

C++ classes group motion settings, calibration, vector, quaternion, and motion-update behavior behind named types.

Read more: [C++ class state and composition](references/language-concepts.md#topic-cpp-class-state-and-composition)

## Unit 148 — Vendored: gamepadmotionhelpers

Shared subsystem: Vendored: gamepadmotionhelpers.

Lesson ID: `unit-cpp-numeric-value-operations`
Display slug: `unit-148-c-numeric-value-operations`

### Implementation: C++ numeric value operations — Vendored: gamepadmotionhelpers

The motion helper uses float-valued vectors and quaternions together with standard-library clamps and maxima.

C++ member-based value types calculate vector lengths, normalized directions, quaternion orientation, and time-scaled motion updates.

Read more: [C++ numeric value operations](references/language-concepts.md#topic-cpp-numeric-value-operations)

## Unit 149 — Vendored: glad

Shared subsystem: Vendored: glad.

Lesson ID: `unit-graphics-procedure-loading`
Display slug: `unit-149-graphics-procedure-loading`

### Domain: Graphics procedure loading — Vendored: glad

The source is a vendored generated loader rather than a rendering implementation.

Generated glad loaders resolve OpenGL, EGL, and GLX procedure entry points and track available versions and extensions.

Read more: [Graphics procedure loading](references/domain-concepts.md#topic-graphics-procedure-loading)

## Unit 150 — Vendored: glslang

Shared subsystem: Vendored: glslang.

Lesson ID: `unit-shader-interface-analysis-spirv-generation-cpp-visitor-traversal`
Display slug: `unit-150-c-visitor-style-intermediate-tree-traversal`

### Domain: Shader interface mapping and reflection — Vendored: glslang

Program-facing reflection records are derived from shader declarations and linker objects.

IO mapping and reflection use intermediate-tree traversal to expose a compiled shader interface as uniform, buffer, atomic-counter, and pipeline input/output entries.

Read more: [Shader interface mapping and reflection](references/domain-concepts.md#topic-shader-interface-analysis)

Use these entity examples: TProgram, TObjectReflection.

### Domain: SPIR-V generation — Vendored: glslang

The lowering code, builder, and IR classes together form the shader binary-generation path.

A traverser lowers the front end's typed intermediate tree to SPIR-V instructions and a module/function/block graph.

Read more: [SPIR-V generation](references/domain-concepts.md#topic-spirv-generation)

Use these entity examples: spv::Module, spv::Function, spv::Block.

### Implementation: C++ visitor-style intermediate-tree traversal — Vendored: glslang

glslang uses specialized TIntermTraverser subclasses for reflection, textual output, live-variable gathering, link validation, and limit checking.

C++ inheritance provides traversal specializations that analyze intermediate shader-tree nodes.

Read more: [C++ visitor-style intermediate-tree traversal](references/language-concepts.md#topic-cpp-visitor-traversal)

Apply it through: TObjectReflection.

## Unit 151 — Vendored: glslang

Shared subsystem: Vendored: glslang.

Lesson ID: `unit-shader-source-compilation-shader-language-front-end`
Display slug: `unit-151-shader-source-compilation`

### Domain: Shader source compilation — Vendored: glslang

The public API separates an individual shader object from a program that links shader objects.

TShader and TProgram compile shader source after preprocessed tokens into stage intermediates and program-level diagnostics.

Read more: [Shader source compilation](references/domain-concepts.md#topic-shader-source-compilation)

Use these entity examples: TShader, TProgram.

### Domain: Shader-language front end — Vendored: glslang

This implementation provides the internal representation consumed by the SPIR-V lowering code.

The vendored glslang front end models shader types, symbols, source locations, parse context, and a typed intermediate tree.

Read more: [Shader-language front end](references/domain-concepts.md#topic-shader-language-front-end)

Use these entity examples: spv::Module, spv::Function, spv::Block.

## Unit 152 — Vendored: graphite

Shared subsystem: Vendored: graphite.

Lesson ID: `unit-font-table-access-graphite-rule-execution-cpp-c-linkage-adapters`
Display slug: `unit-152-binary-font-table-access`

### Domain: Binary font-table access — Vendored: graphite

The code directly models table records, offsets, and character-to-glyph lookup structures.

Graphite and HarfBuzz read binary font tables, including cmap mappings, to obtain glyph data and layout behavior.

Read more: [Binary font-table access](references/domain-concepts.md#topic-font-table-access)

Use these entity examples: Face, hb_blob_t, hb_face_t.

### Domain: Graphite SILF rule execution — Vendored: graphite

Rules are loaded from Graphite font data, matched during passes, and evaluated by the machine implementation.

Graphite executes decoded SILF rule constraints and actions through a finite-state matcher and bytecode machine.

Read more: [Graphite SILF rule execution](references/domain-concepts.md#topic-graphite-rule-execution)

Use these entity examples: Silf, Segment, Slot.

### Implementation: C++ C-linkage adapters — Vendored: graphite

Graphite exposes C-callable functions that delegate to C++ objects such as CharInfo, Face, Segment, and Slot.

C++ classes are wrapped behind C linkage functions and derived opaque API types, so C callers use handles rather than C++ class declarations.

Read more: [C++ C-linkage adapters](references/language-concepts.md#topic-cpp-c-linkage-adapters)

Apply it through: Face, Font, Segment, Slot.

## Unit 153 — Vendored: graphite

Shared subsystem: Vendored: graphite.

Lesson ID: `unit-graphite-shaping-cpp-resource-lifetime`
Display slug: `unit-153-graphite-segment-shaping`

### Domain: Graphite segment shaping — Vendored: graphite

A face selects the applicable SILF behavior, a font supplies scaled metrics, and a segment exposes shaped slots.

Graphite uses decoded font tables and rule passes to turn Unicode into a Segment whose linked Slots carry glyph and placement state.

Read more: [Graphite segment shaping](references/domain-concepts.md#topic-graphite-shaping)

Use these entity examples: Face, Font, Segment, Slot, FeatureVal.

### Implementation: C++ constructor and destructor resource lifetime — Vendored: graphite

The Font constructor establishes borrowed references and per-glyph cache storage; its destructor frees that storage.

A C++ class constructor and destructor control the Font's allocation and release of its cached advance array.

Read more: [C++ constructor and destructor resource lifetime](references/language-concepts.md#topic-cpp-resource-lifetime)

Apply it through: Font.

## Unit 154 — Vendored: graphite

Shared subsystem: Vendored: graphite.

Lesson ID: `unit-cpp-templates-for-binary-data`
Display slug: `unit-154-c-templates-for-binary-data-operations`

### Implementation: C++ templates for binary data operations — Vendored: graphite

The libraries use templates to retain typed interfaces while parsing binary font tables and copying fixed-size byte sequences.

C++ templates specialize reusable byte operations for a compile-time size and typed OpenType offsets.

Read more: [C++ templates for binary data operations](references/language-concepts.md#topic-cpp-templates-for-binary-data)

Apply it through: hb_blob_t, hb_face_t.

## Unit 155 — Vendored: harfbuzz

Shared subsystem: Vendored: harfbuzz.

Lesson ID: `unit-unicode-data-font-subsetting-cpp-preprocessor-configuration`
Display slug: `unit-155-unicode-data-and-properties`

### Domain: Unicode data and properties — Vendored: harfbuzz

Generated tables are paired with implementations that expose properties and Unicode processing APIs.

The included code stores generated Unicode script, emoji, normalization, bidi, case, and general-character-property data for lookup-oriented services.

Read more: [Unicode data and properties](references/domain-concepts.md#topic-unicode-data)

Use these entity examples: UCPTrie, hb_unicode_funcs_t.

### Domain: Font subsetting — Vendored: harfbuzz

The partition contains input, planning, table-specific processing, serialization, and CFF-specific subset code.

HarfBuzz represents caller selections as a subset input and derives an in-memory plan for rewriting a font.

Read more: [Font subsetting](references/domain-concepts.md#topic-font-subsetting)

Use these entity examples: hb_subset_input_t, hb_subset_plan_t.

### Implementation: C++ preprocessor configuration — Vendored: harfbuzz

HarfBuzz uses feature macros such as HB_NO_COLOR and include guards throughout this partition.

Preprocessor conditions select optional HarfBuzz subsystems and header inclusion boundaries before C++ compilation.

Read more: [C++ preprocessor configuration](references/language-concepts.md#topic-cpp-preprocessor-configuration)

Apply it through: hb_subset_plan_t.

## Unit 156 — Vendored: harfbuzz

Shared subsystem: Vendored: harfbuzz.

Lesson ID: `unit-color-font-paint-cff-font-subsetting-cpp-static-generated-data`
Display slug: `unit-156-color-font-paint-processing`

### Domain: Color-font paint processing — Vendored: harfbuzz

The paint layer is independent of a particular raster or vector output format.

HarfBuzz supplies paint functions plus bounded and extents contexts to evaluate color-font paint operations.

Read more: [Color-font paint processing](references/domain-concepts.md#topic-color-font-paint)

Use these entity examples: hb_paint_funcs_t, hb_paint_extents_context_t.

### Domain: CFF font subsetting — Vendored: harfbuzz

CFF-specific code is distinct from generic table dispatch because it carries charstring and subroutine structures.

The font subsetting pipeline contains CFF1 and CFF2 plans, accelerators, serializers, and subroutine-subsetting helpers.

Read more: [CFF font subsetting](references/domain-concepts.md#topic-cff-font-subsetting)

Use these entity examples: hb_subset_plan_t, hb_subset_accelerator_t.

### Implementation: C++ static generated data — Vendored: harfbuzz

The repository contains generated HarfBuzz UCD tables and generated ICU normalization, bidi, case, and character-property tables.

Static const arrays and values embed generated Unicode and normalization data directly in translation units.

Read more: [C++ static generated data](references/language-concepts.md#topic-cpp-static-generated-data)

Apply it through: UCPTrie.

## Unit 157 — Vendored: harfbuzz

Shared subsystem: Vendored: harfbuzz.

Lesson ID: `unit-opentype-table-routing-raster-font-rendering-cpp-templates`
Display slug: `unit-157-opentype-table-routing`

### Domain: OpenType table routing — Vendored: harfbuzz

Separate dispatchers cover color, layout, ordinary glyph-related, and variation tables.

Font subsetting routes recognized OpenType tags to typed table subsetters; color routing explicitly skips CBDT because CBLC handles it.

Read more: [OpenType table routing](references/domain-concepts.md#topic-opentype-table-routing)

Use these entity examples: hb_subset_plan_t.

### Domain: Raster font rendering — Vendored: harfbuzz

The raster implementation includes PNG-related image code as well as scan-conversion-oriented drawing code.

Color paint records are rasterized through image, draw, paint, clipping, edge, and glyph-extents structures.

Read more: [Raster font rendering](references/domain-concepts.md#topic-raster-font-rendering)

Use these entity examples: hb_raster_image_t, hb_raster_paint_t.

### Implementation: C++ templates — Vendored: harfbuzz

The code uses template parameters for both general utility code and statically selected font-table implementations.

Templates parameterize data structures and algorithms, including string code-point appenders and typed OpenType table subset calls.

Read more: [C++ templates](references/language-concepts.md#topic-cpp-templates)

Apply it through: UCPTrie, hb_subset_plan_t.

## Unit 158 — Vendored: harfbuzz

Shared subsystem: Vendored: harfbuzz.

Lesson ID: `unit-variable-font-subsetting-vector-font-export-text-shaping-plans`
Display slug: `unit-158-variable-font-table-subsetting`

### Domain: Variable-font table subsetting — Vendored: harfbuzz

Variation handling is controlled by table tags and user-axis locations in the subset plan.

The font subsetting pipeline dispatches HVAR, VVAR, gvar, fvar, avar, cvar, and mvar tables, with fvar and avar passed through when user axis locations are empty.

Read more: [Variable-font table subsetting](references/domain-concepts.md#topic-variable-font-subsetting)

Use these entity examples: hb_subset_plan_t.

### Domain: Vector font export — Vendored: harfbuzz

The vector path, draw, paint, PDF, and SVG files form a separate output family from raster paint.

Color paint records are emitted through vector drawing and paint backends that include PDF and SVG implementations.

Read more: [Vector font export](references/domain-concepts.md#topic-vector-font-export)

Use these entity examples: hb_vector_paint_t, hb_vector_draw_t.

### Domain: Text shaping plans — Vendored: harfbuzz

The public shaping entry points accept features and an optional shaper list, while plan creation accepts matching configuration.

A shaping-plan key records segment properties, user features, coordinates, and shaper selection so shaping can use a plan object.

Read more: [Text shaping plans](references/domain-concepts.md#topic-text-shaping-plans)

Use these entity examples: hb_shape_plan_t, hb_shape_plan_key_t.

## Unit 159 — Vendored: icu4c

Shared subsystem: Vendored: icu4c.

Lesson ID: `unit-character-encoding-conversion-identifier-spoof-checking-cpp-classes`
Display slug: `unit-159-c-classes-and-inheritance`

### Domain: Character-encoding conversion — Vendored: icu4c

The partition includes UTF-7, UTF-8, UTF-16, UTF-32, ISO-2022, MBCS, and other converter implementations.

ICU models converters with shared static data, an implementation dispatch structure, contexts, and UTF-specific implementations.

Read more: [Character-encoding conversion](references/domain-concepts.md#topic-character-encoding-conversion)

Use these entity examples: UConverter, UConverterSharedData.

### Domain: Identifier spoof checking — Vendored: icu4c

The implementation includes confusable data, script sets, check results, and UTF-16 and UTF-8 API paths.

Unicode property checks and configured allowed character or locale sets are held by SpoofImpl and exposed through USpoofChecker APIs.

Read more: [Identifier spoof checking](references/domain-concepts.md#topic-identifier-spoof-checking)

Use these entity examples: SpoofImpl, USpoofChecker.

### Implementation: C++ classes and inheritance — Vendored: icu4c

Class declarations organize most ICU and Jolt stateful services; HarfBuzz more often uses structs.

C++ classes define reusable state and behavior, and the implementation derives UnicodeString from Replaceable and many ICU services from UObject or UMemory.

Read more: [C++ classes and inheritance](references/language-concepts.md#topic-cpp-classes)

Apply it through: Locale, AABBTreeBuilder.

## Unit 160 — Vendored: icu4c

Shared subsystem: Vendored: icu4c.

Lesson ID: `unit-break-iteration-unicode-normalization-cpp-overloading`
Display slug: `unit-160-c-member-function-overloading`

### Domain: Text break iteration — Vendored: icu4c

The partition includes parsing, table building, runtime iteration, and dictionary or LSTM break-engine code.

ICU builds and executes rule-based break iterators, including dictionary-backed language break engines and compiled rule tables.

Read more: [Text break iteration](references/domain-concepts.md#topic-break-iteration)

Use these entity examples: RuleBasedBreakIterator, LanguageBreakEngine.

### Domain: Unicode normalization — Vendored: icu4c

The code contains NFC, NFKC, FCD, and related Normalizer2 factory paths.

ICU implements Normalizer2 variants and loads normalization data into tries, extra mappings, and composition data.

Read more: [Unicode normalization](references/domain-concepts.md#topic-unicode-normalization)

Use these entity examples: Normalizer2Impl, UCPTrie.

### Implementation: C++ member-function overloading — Vendored: icu4c

The overload set preserves one conceptual operation while accepting multiple input representations.

Overloaded member functions provide related operations for different argument forms, such as LocaleMatcher matching a single locale, an iterator, or a list string.

Read more: [C++ member-function overloading](references/language-concepts.md#topic-cpp-overloading)

Apply it through: LocaleMatcher, Locale.

## Unit 161 — Vendored: icu4c

Shared subsystem: Vendored: icu4c.

Lesson ID: `unit-locale-resolution-resource-bundle-data-cpp-virtual-dispatch`
Display slug: `unit-161-c-virtual-dispatch`

### Domain: Locale resolution — Vendored: icu4c

The implementation exposes both object-oriented LocaleMatcher APIs and lower-level locale parsing code.

ICU parses locale identifiers and matches desired locales against supported locales using LSR values, likely-subtag data, and locale-distance data.

Read more: [Locale resolution](references/domain-concepts.md#topic-locale-resolution)

Use these entity examples: Locale, LSR, LocaleMatcher.

### Domain: Resource-bundle data — Vendored: icu4c

The code has both C-facing UResourceBundle state and C++ ResourceBundle wrappers.

ICU resource bundles expose strings, binary values, integer vectors, keys, names, and locales over loaded resource data.

Read more: [Resource-bundle data](references/domain-concepts.md#topic-resource-bundle-data)

Use these entity examples: UResourceBundle, ResourceDataValue.

### Implementation: C++ virtual dispatch — Vendored: icu4c

This is used where ICU needs an interface with multiple concrete implementations.

Virtual dispatch uses derived classes to substitute behavior, as locale iterators override next and ICU break engines derive from a common engine type.

Read more: [C++ virtual dispatch](references/language-concepts.md#topic-cpp-virtual-dispatch)

Apply it through: LocaleMatcher, RuleBasedBreakIterator.

## Unit 162 — Vendored: icu4c

Shared subsystem: Vendored: icu4c.

Lesson ID: `unit-cpp-object-representation-casts-cpp-function-pointer-interfaces`
Display slug: `unit-162-c-object-representation-casts`

### Implementation: C++ object-representation casts — Vendored: icu4c

These operations are concentrated in code that consumes serialized binary layouts.

Pointer reinterpretation reads packed data through typed views, particularly when ICU validates or loads binary resource and Unicode data.

Read more: [C++ object-representation casts](references/language-concepts.md#topic-cpp-object-representation-casts)

Apply it through: UResourceBundle, UCPTrie.

### Implementation: C-compatible function-pointer interfaces — Vendored: icu4c

This is a data-driven dispatch mechanism used alongside C++ class polymorphism.

C-compatible function-pointer interfaces let converter implementations select operations through a shared implementation record.

Read more: [C-compatible function-pointer interfaces](references/language-concepts.md#topic-cpp-function-pointer-interfaces)

Apply it through: UConverter, UConverterSharedData.

## Unit 163 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-constraints-narrow-phase-cxx-reflection-macros`
Display slug: `unit-163-constraint-solving`

### Domain: Constraint solving — Vendored: jolt_physics

ConstraintSettings are serializable configuration objects; runtime Constraint objects participate in solver islands.

Constraint solving applies impulses to rigid-body motion and reuses contact manifolds while supporting point, distance, hinge, slider, fixed, gear, pulley, cone, six-degree-of-freedom, path, and swing-twist constraints.

Read more: [Constraint solving](references/domain-concepts.md#topic-constraints)

Use these entity examples: ConstraintSettings, TwoBodyConstraint.

### Domain: Narrow-phase collision queries — Vendored: jolt_physics

Results are emitted through specialized collision collectors and dispatch handlers.

Narrow-phase collision queries test collision shapes for rays, points, overlaps, and casts after broad-phase candidates have been selected.

Read more: [Narrow-phase collision queries](references/domain-concepts.md#topic-narrow-phase)

Use these entity examples: Shape, Body.

### Implementation: Macro-based RTTI registration — Vendored: jolt_physics

The JPH_IMPLEMENT_SERIALIZABLE and JPH_ADD_ATTRIBUTE macro families are the repository's reflection convention.

Macro-based RTTI registration attaches class metadata and serializable member attributes to types that participate in dynamic lookup and object streaming.

Read more: [Macro-based RTTI registration](references/language-concepts.md#topic-cxx-reflection-macros)

Apply it through: ConstraintSettings, PhysicsMaterialSimple.

## Unit 164 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-collision-shapes-vehicle-dynamics-cxx-object-model`
Display slug: `unit-164-collision-shapes`

### Domain: Collision shapes — Vendored: jolt_physics

The Shape hierarchy is the common geometry contract for bodies and collision queries.

Collision shapes define the geometric representation used for ray casts, shape casts, point tests, overlap tests, triangle extraction, and collision dispatch.

Read more: [Collision shapes](references/domain-concepts.md#topic-collision-shapes)

Use these entity examples: Shape, ShapeSettings.

### Domain: Vehicle dynamics — Vendored: jolt_physics

Wheeled, motorcycle, and tracked controllers are separate implementations over the vehicle constraint.

Vehicle dynamics couples a VehicleConstraint with a rigid Body, wheel collision tests, suspension, controllers, engine, transmission, differentials, tracks, and anti-roll bars.

Read more: [Vehicle dynamics](references/domain-concepts.md#topic-vehicle-dynamics)

Use these entity examples: Body, VehicleTransmissionSettings.

### Implementation: C++ classes, inheritance, and polymorphic interfaces — Vendored: jolt_physics

The implementation favors abstract base types with specialized derived classes.

C++ classes and inheritance define extensible interfaces for shapes, constraints, collision queries, job systems, renderers, and body access.

Read more: [C++ classes, inheritance, and polymorphic interfaces](references/language-concepts.md#topic-cxx-object-model)

Apply it through: Shape, Body.

## Unit 165 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-broad-phase-collision-filtering-cxx-preprocessor-configuration`
Display slug: `unit-165-broad-phase-collision-detection`

### Domain: Broad-phase collision detection — Vendored: jolt_physics

The code provides brute-force and quadtree broad-phase implementations.

Broad-phase collision detection uses Body world-space AABox bounds and collision filtering to generate candidate body pairs and answer coarse queries.

Read more: [Broad-phase collision detection](references/domain-concepts.md#topic-broad-phase)

Use these entity examples: Body.

### Domain: Collision filtering — Vendored: jolt_physics

The filtering interfaces let a caller restrict which bodies or subshapes may interact.

Collision filtering applies object-layer, broad-phase-layer, group, body, and shape filters before candidate generation or collision queries.

Read more: [Collision filtering](references/domain-concepts.md#topic-collision-filtering)

Use these entity examples: Body.

### Implementation: Preprocessor-selected platform configuration — Vendored: jolt_physics

Jolt centralizes many feature flags in Core.h and applies branches in platform-specific implementation headers.

Preprocessor configuration selects processor architecture, platform APIs, floating-point precision, instruction-set support, diagnostics, and optional subsystems.

Read more: [Preprocessor-selected platform configuration](references/language-concepts.md#topic-cxx-preprocessor-configuration)

Apply it through: Body.

## Unit 166 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-contact-management-convex-collision-cxx-raii-reference-ownership`
Display slug: `unit-166-contact-manifolds-and-warm-starting`

### Domain: Contact manifolds and warm-starting — Vendored: jolt_physics

The contact manager owns the cache used to retain collision-solver information.

Contact management converts narrow-phase collision results into contact constraints and caches manifolds, body pairs, and contact points between updates.

Read more: [Contact manifolds and warm-starting](references/domain-concepts.md#topic-contact-management)

Use these entity examples: CachedManifold, Body.

### Domain: Convex support and penetration — Vendored: jolt_physics

ConvexShape exposes support mappings that collision algorithms can query.

Convex collision represents collision shapes with support functions and uses GJK closest-point and EPA penetration calculations.

Read more: [Convex support and penetration](references/domain-concepts.md#topic-convex-collision)

Use these entity examples: Shape.

### Implementation: RAII, non-copyable resources, and intrusive references — Vendored: jolt_physics

The library explicitly marks many resources NonCopyable and wraps longer-lived objects in intrusive reference types.

RAII-bound C++ objects manage locks and scopes, while Ref and RefTarget keep shared simulation objects alive without ordinary copying.

Read more: [RAII, non-copyable resources, and intrusive references](references/language-concepts.md#topic-cxx-raii-reference-ownership)

Apply it through: Shape, PhysicsScene.

## Unit 167 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-serialization-rigid-body-motion-cxx-simd-alignment`
Display slug: `unit-167-rtti-based-serialization`

### Domain: RTTI-based serialization — Vendored: jolt_physics

The same registration pattern is used by settings, materials, constraints, skeletons, and curves.

RTTI-based serialization registers attributes for object types and writes or restores binary state through StreamIn and StreamOut.

Read more: [RTTI-based serialization](references/domain-concepts.md#topic-serialization)

Use these entity examples: ConstraintSettings, PhysicsMaterialSimple, Skeleton.

### Domain: Rigid-body motion — Vendored: jolt_physics

MotionProperties is the mutable dynamic state associated with movable rigid bodies.

Rigid-body motion uses Body transform state plus MotionProperties to model static, kinematic, and dynamic behavior, permitted degrees of freedom, mass, inertia, and velocities.

Read more: [Rigid-body motion](references/domain-concepts.md#topic-rigid-body-motion)

Use these entity examples: Body.

### Implementation: SIMD wrappers and alignment — Vendored: jolt_physics

Public math types hide platform intrinsics behind Vec3, Vec4, UVec4, matrices, and box helpers.

SIMD vector types use aligned storage and compile-time CPU branches to implement vector, matrix, bounding-box, and geometry operations across SSE, NEON, and RISC-V vector paths.

Read more: [SIMD wrappers and alignment](references/language-concepts.md#topic-cxx-simd-alignment)

Apply it through: Body.

## Unit 168 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-state-recording-geometry-preprocessing-cxx-stream-serialization`
Display slug: `unit-168-state-recording-and-validation`

### Domain: State recording and validation — Vendored: jolt_physics

StateRecorderImpl supports rewinding, clearing, reading, writing, and validation diagnostics.

State recording saves and validates Body and Constraint state through a stream that can compare replayed bytes with current values.

Read more: [State recording and validation](references/domain-concepts.md#topic-state-recording)

Use these entity examples: Body, TwoBodyConstraint.

### Domain: Geometry preprocessing — Vendored: jolt_physics

The geometry utilities are used to prepare collision-friendly representations from source triangles and vertices.

Geometry preprocessing converts triangle lists to indexed geometry, builds convex hulls, and partitions triangles for spatial acceleration structures used by collision shapes.

Read more: [Geometry preprocessing](references/domain-concepts.md#topic-geometry-preprocessing)

Use these entity examples: Shape.

### Implementation: Stream-oriented binary serialization — Vendored: jolt_physics

Several types pair SaveBinaryState and RestoreBinaryState methods with their RTTI registrations.

StreamIn and StreamOut implement binary persistence using RTTI hashes for dynamically typed objects and explicit field writes for concrete state.

Read more: [Stream-oriented binary serialization](references/language-concepts.md#topic-cxx-stream-serialization)

Apply it through: ConstraintSettings, PhysicsMaterialSimple, Skeleton.

## Unit 169 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-skeletal-ragdoll-aabb-tree-construction-cxx-lambdas-standard-algorithms`
Display slug: `unit-169-skeletons-animation-and-ragdolls`

### Domain: Skeletons, animation, and ragdolls — Vendored: jolt_physics

SkeletonPose, SkeletalAnimation, SkeletonMapper, and Ragdoll form the animation-to-physics integration surface.

Skeletons retain named joints and parent relationships, animations produce joint states, and ragdolls connect those joints to Body parts and two-body constraints.

Read more: [Skeletons, animation, and ragdolls](references/domain-concepts.md#topic-skeletal-ragdoll)

Use these entity examples: Skeleton, Joint, Body.

### Domain: AABB tree construction — Vendored: jolt_physics

This is a geometry-acceleration structure within the supplied Jolt partition.

Jolt's AABBTreeBuilder stores nodes and indexed triangles and accepts a maximum-triangles-per-leaf setting while building its tree.

Read more: [AABB tree construction](references/domain-concepts.md#topic-aabb-tree-construction)

Use these entity examples: AABBTreeBuilder.

### Implementation: Lambdas and standard algorithms — Vendored: jolt_physics

The implementation uses lambdas for ordered curve lookup and for constructing static sampled geometry.

Lambdas provide inline comparison and initialization behavior when standard algorithms or static data construction need local callable logic.

Read more: [Lambdas and standard algorithms](references/language-concepts.md#topic-cxx-lambdas-standard-algorithms)

Apply it through: LinearCurve.

## Unit 170 — Vendored: jolt_physics

Shared subsystem: Vendored: jolt_physics.

Lesson ID: `unit-cxx-enums-bitflags`
Display slug: `unit-170-scoped-enums-and-bit-flags`

### Implementation: Scoped enums and bit flags — Vendored: jolt_physics

The code uses compact underlying integer types for many externally meaningful simulation modes.

Scoped enum class values model motion, body, collision, update-error, and configuration states, while selected enums provide bitwise combinations.

Read more: [Scoped enums and bit flags](references/language-concepts.md#topic-cxx-enums-bitflags)

Apply it through: Body, ConstraintSettings.

## Unit 171 — Vendored: libjpeg-turbo

Shared subsystem: Vendored: libjpeg-turbo.

Lesson ID: `unit-image-decode-pipeline-c-aggregate-callback-modules`
Display slug: `unit-171-image-decode-pipelines`

### Domain: Image decode pipelines — Vendored: libjpeg-turbo

This is the common lifecycle behind the JPEG, PNG, and WebP implementations in this partition.

The vendored image libraries parse encoded input, keep decoder state, and emit raster rows or planes; JPEG has scanline and upsampling modules, PNG applies row transforms, and WebP routes decoded output through VP8 I/O.

Read more: [Image decode pipelines](references/domain-concepts.md#topic-image-decode-pipeline)

Use these entity examples: JPEG Decompression Context, PNG Information State, WebP Decoder State.

### Implementation: C aggregate state and callback modules — Vendored: libjpeg-turbo

The implementation uses public context structures plus private extension structures instead of language-level objects.

C aggregate types and typed function pointers define stateful processing modules, letting JPEG upsamplers retain shared buffers and dispatch a per-component routine.

Read more: [C aggregate state and callback modules](references/language-concepts.md#topic-c-aggregate-callback-modules)

Apply it through: JPEG Decompression Context, JPEG Upsampler State.

## Unit 172 — Vendored: libjpeg-turbo

Shared subsystem: Vendored: libjpeg-turbo.

Lesson ID: `unit-jpeg-decompression-transcoding-c-platform-selection`
Display slug: `unit-172-c-preprocessor-platform-and-precision-selection`

### Domain: JPEG decompression and coefficient transcoding — Vendored: libjpeg-turbo

The normal decompression path is modular; the coefficient path is explicitly documented as a transcoding mode.

The JPEG code implements an image decode pipeline that selects decompression modules, can merge chroma upsampling with YCC-to-RGB conversion, and can expose raw DCT coefficient arrays for transcoding.

Read more: [JPEG decompression and coefficient transcoding](references/domain-concepts.md#topic-jpeg-decompression-transcoding)

Use these entity examples: JPEG Decompression Context, JPEG Upsampler State.

### Implementation: C preprocessor platform and precision selection — Vendored: libjpeg-turbo

The same source tree supports multiple sample precisions, compilers, and CPU feature sets through macro-controlled compilation.

C conditional compilation selects precision-dependent names, optional compiler intrinsics, and architecture-specific implementations before the codec modules are compiled.

Read more: [C preprocessor platform and precision selection](references/language-concepts.md#topic-c-platform-selection)

Apply it through: JPEG Upsampler State.

## Unit 173 — Vendored: libjpeg-turbo

Shared subsystem: Vendored: libjpeg-turbo.

Lesson ID: `unit-simd-accelerated-codecs`
Display slug: `unit-173-optional-simd-codec-paths`

### Domain: Optional SIMD codec paths — Vendored: libjpeg-turbo

Selection occurs through feature macros, runtime checks where implemented, and function-pointer assignment.

The codec libraries use C conditional compilation to select architecture-specific NEON, SSE2, VSX, LSX, MMX, and other optimized routines, while retaining scalar fallbacks.

Read more: [Optional SIMD codec paths](references/domain-concepts.md#topic-simd-accelerated-codecs)

Use these entity examples: JPEG Upsampler State.

## Unit 174 — Vendored: libktx

Shared subsystem: Vendored: libktx.

Lesson ID: `unit-ktx-texture-container-cpp-basis-transcoding`
Display slug: `unit-174-c-basis-transcoding-behind-a-c-compatible-api`

### Domain: KTX texture containers — Vendored: libktx

KTX1 and KTX2 are represented by related texture objects with internal constructors and type-specific vtables.

KTX texture containers carry a texture format description, per-level indexing, stream or memory-backed data, virtual dispatch tables, and KTX2 supercompression state.

Read more: [KTX texture containers](references/domain-concepts.md#topic-ktx-texture-container)

Use these entity examples: KTX2 Texture, KTX2 Private Texture State, KTX Level Index Entry.

### Implementation: C++ Basis transcoding behind a C-compatible API — Vendored: libktx

This is the material C++ portion of the supplied KTX implementation; most neighboring public and internal KTX APIs are C declarations.

The KTX Basis transcode implementation uses C++ references and standard-library state vectors, while its headers retain a C-compatible interface for the enclosing C API.

Read more: [C++ Basis transcoding behind a C-compatible API](references/language-concepts.md#topic-cpp-basis-transcoding)

Apply it through: KTX2 Texture, KTX2 Private Texture State.

## Unit 175 — Vendored: libktx

Shared subsystem: Vendored: libktx.

Lesson ID: `unit-texture-format-description`
Display slug: `unit-175-texture-format-descriptions`

### Domain: Texture format descriptions — Vendored: libktx

The descriptor is an exchanged binary description of texel layout rather than an opaque GPU-native layout.

Khronos Data Format descriptors represent image-format layout using header-word and sample-word accessors, while helper code creates, interprets, queries, and prints those descriptors.

Read more: [Texture format descriptions](references/domain-concepts.md#topic-texture-format-description)

Use these entity examples: KTX Level Index Entry.

## Unit 176 — Vendored: libogg

Shared subsystem: Vendored: libogg.

Lesson ID: `unit-ogg-pages-and-packets`
Display slug: `unit-176-ogg-pages-packets-and-bit-packing`

### Domain: Ogg pages, packets, and bit packing — Vendored: libogg

This is the shared container layer used by the supplied Vorbis and Theora implementations.

The Ogg implementation packs variable-sized words into octet streams and maintains stream state that frames packets into page headers and bodies.

Read more: [Ogg pages, packets, and bit packing](references/domain-concepts.md#topic-ogg-pages-and-packets)

Use these entity examples: Ogg Stream State, Ogg Packet.

## Unit 177 — Vendored: libpng

Shared subsystem: Vendored: libpng.

Lesson ID: `unit-png-stream-transforms`
Display slug: `unit-177-png-streaming-i-o-and-row-transforms`

### Domain: PNG streaming I/O and row transforms — Vendored: libpng

I/O customization and pixel transformations are separate from information-structure mutation.

The PNG code implements an image decode pipeline with replaceable read and write callbacks, push-mode input states, metadata validity flags, and row-level transformations such as BGR mapping and 16-bit byte swapping.

Read more: [PNG streaming I/O and row transforms](references/domain-concepts.md#topic-png-stream-transforms)

Use these entity examples: PNG Information State.

## Unit 178 — Vendored: libtheora

Shared subsystem: Vendored: libtheora.

Lesson ID: `unit-theora-block-video-codec`
Display slug: `unit-178-theora-block-video-coding`

### Domain: Theora block video coding — Vendored: libtheora

Encoding and decoding are implemented separately but share transform, quantization, fragment, and state machinery.

The Theora codec consumes Ogg packets while unpacking quantization parameters, DCT token data, motion-compensated frame state, and optional accelerated transform routines.

Read more: [Theora block video coding](references/domain-concepts.md#topic-theora-block-video-codec)

Use these entity examples: Theora Stream Information, Ogg Packet.

## Unit 179 — Vendored: libvorbis

Shared subsystem: Vendored: libvorbis.

Lesson ID: `unit-vorbis-block-synthesis`
Display slug: `unit-179-vorbis-block-analysis-and-synthesis`

### Domain: Vorbis block analysis and synthesis — Vendored: libvorbis

The implementation has separate analysis and synthesis entry points over the same block-oriented codec model.

Vorbis analysis and synthesis operate on codec blocks and Ogg packets, using mapping, floor, residue, codebook, window, and MDCT modules.

Read more: [Vorbis block analysis and synthesis](references/domain-concepts.md#topic-vorbis-block-synthesis)

Use these entity examples: Vorbis Block, Ogg Packet.

## Unit 180 — Vendored: libwebp

Shared subsystem: Vendored: libwebp.

Lesson ID: `unit-input-picture-planes-encoder-configuration-c-preprocessor-portability`
Display slug: `unit-180-picture-planes-and-pixel-ownership`

### Domain: Picture planes and pixel ownership — Vendored: libwebp

The selected representation determines which pixel fields the encoder treats as native input.

WebPPicture represents one image as either ARGB pixels or YUV(A) planes with dimensions, strides, optional alpha, and private allocation pointers.

Read more: [Picture planes and pixel ownership](references/domain-concepts.md#topic-input-picture-planes)

Use these entity examples: WebPPicture.

### Domain: Encoder configuration — Vendored: libwebp

Initialization and validation are explicit ABI-aware API steps.

WebPConfig selects lossy or lossless encoding and controls quality, effort, segmentation, filtering, alpha, thread, and memory choices for the configured picture.

Read more: [Encoder configuration](references/domain-concepts.md#topic-encoder-configuration)

Use these entity examples: WebPConfig, WebPPicture.

### Implementation: C preprocessor portability — Vendored: libwebp

This is the portability layer that lets scalar and architecture-specific sources coexist.

Conditional compilation and macros select target-dependent code, inline helpers, endianness behavior, and instruction-family implementations at compile time.

Read more: [C preprocessor portability](references/language-concepts.md#topic-c-preprocessor-portability)

## Unit 181 — Vendored: libwebp

Shared subsystem: Vendored: libwebp.

Lesson ID: `unit-lossy-macroblock-encoding-endian-safe-binary-io-c-integer-bitwise-packing`
Display slug: `unit-181-lossy-macroblock-encoding`

### Domain: Lossy macroblock encoding — Vendored: libwebp

Analysis, iteration, quantization, filtering, and syntax emission are separate source modules in the lossy path.

The lossy encoder walks the configured picture planes by macroblock, choosing prediction modes and producing residual tokens for the VP8 stream.

Read more: [Lossy macroblock encoding](references/domain-concepts.md#topic-lossy-macroblock-encoding)

Use these entity examples: VP8Encoder, WebPPicture, WebPConfig.

### Domain: Endian-safe binary I/O — Vendored: libwebp

The implementation uses conditional endianness macros instead of assuming a host byte order.

Endian macros and memory conversion helpers normalize integer fields while container and bitstream code serializes binary bytes.

Read more: [Endian-safe binary I/O](references/domain-concepts.md#topic-endian-safe-binary-io)

Use these entity examples: WebPData.

### Implementation: C fixed-width integers and bitwise packing — Vendored: libwebp

The code uses uint8_t through uint64_t to state intended storage widths in data and bitstream paths.

Fixed-width integer values hold packed ARGB samples, codec fields, chunk tags, and arithmetic intermediates; shifts, masks, and bitwise operators extract or assemble channel values.

Read more: [C fixed-width integers and bitwise packing](references/language-concepts.md#topic-c-integer-bitwise-packing)

Apply it through: WebPData, WebPPicture.

## Unit 182 — Vendored: libwebp

Shared subsystem: Vendored: libwebp.

Lesson ID: `unit-alpha-plane-coding-cpu-specialized-dsp-c-pointers-arrays-and-strides`
Display slug: `unit-182-alpha-plane-coding`

### Domain: Alpha-plane coding — Vendored: libwebp

Alpha encoding has its own filter choice and can use lossless compression.

The alpha plane is filtered, optionally quantized, and encoded before its bytes are incorporated with the lossy image representation.

Read more: [Alpha-plane coding](references/domain-concepts.md#topic-alpha-plane-coding)

Use these entity examples: WebPPicture, WebPConfig.

### Domain: CPU-specialized DSP backends — Vendored: libwebp

The scalar sources establish algorithms while architecture files express equivalent work with vector intrinsics or target assembly.

DSP function families have scalar, SSE2/SSE4.1/AVX2, NEON, MSA, and MIPS implementations that are compile-time selected for target capabilities.

Read more: [CPU-specialized DSP backends](references/domain-concepts.md#topic-cpu-specialized-dsp)

### Implementation: C pointers, arrays, and strides — Vendored: libwebp

Most image operations accept borrowed pointer ranges and compute row addresses instead of owning per-pixel objects.

C aggregates are reached through pointers and stepped across image memory with width- and stride-based pointer arithmetic for rows and planes.

Read more: [C pointers, arrays, and strides](references/language-concepts.md#topic-c-pointers-arrays-and-strides)

Apply it through: WebPPicture, WebPRescaler, WebPDecBuffer.

## Unit 183 — Vendored: libwebp

Shared subsystem: Vendored: libwebp.

Lesson ID: `unit-image-quality-metrics-lossless-transform-pipeline-c-resource-lifecycles`
Display slug: `unit-183-c-resource-lifecycles-and-ownership`

### Domain: Image quality metrics — Vendored: libwebp

The encoder exposes aggregate statistics and picture-distortion helpers for reporting and comparison.

PSNR, SSIM, local similarity, and squared-error routines assess encoded or reconstructed image planes.

Read more: [Image quality metrics](references/domain-concepts.md#topic-image-quality-metrics)

Use these entity examples: WebPAuxStats, WebPPicture.

### Domain: Lossless pixel transform pipeline — Vendored: libwebp

Transform choice is analyzed from the input and recorded as lossless coding features.

The lossless encoder transforms ARGB pixels through predictor, subtract-green, cross-color, and color-indexing stages before entropy coding.

Read more: [Lossless pixel transform pipeline](references/domain-concepts.md#topic-lossless-transform-pipeline)

Use these entity examples: WebPPicture, VP8Encoder.

### Implementation: C resource lifecycles and ownership — Vendored: libwebp

Resource retirement is explicit; structure storage itself is generally caller-owned.

Pointer-bearing structures use explicit init, allocate, clear, free, copy, and view functions so the caller can distinguish owned image buffers from borrowed views.

Read more: [C resource lifecycles and ownership](references/language-concepts.md#topic-c-resource-lifecycles)

Apply it through: WebPPicture, WebPMemoryWriter, WebPDecBuffer.

## Unit 184 — Vendored: libwebp

Shared subsystem: Vendored: libwebp.

Lesson ID: `unit-spatial-predictive-filters-transform-quantization-rate-distortion-c-abi-versioned-initialization`
Display slug: `unit-184-spatial-predictive-filters`

### Domain: Spatial predictive filters — Vendored: libwebp

Scalar, SSE2, NEON, MSA, and MIPS implementations cover the same filtering roles.

Spatial filters produce residuals from neighboring pixel-plane values using horizontal, vertical, gradient, and other predictor forms.

Read more: [Spatial predictive filters](references/domain-concepts.md#topic-spatial-predictive-filters)

### Domain: Transform, quantization, and rate-distortion search — Vendored: libwebp

The code includes both direct quantization and trellis-based rate-distortion choices.

The macroblock path transforms residuals, quantizes them with VP8Matrix parameters, evaluates rate-distortion costs, and reconstructs blocks.

Read more: [Transform, quantization, and rate-distortion search](references/domain-concepts.md#topic-transform-quantization-rate-distortion)

Use these entity examples: VP8Encoder.

### Implementation: C ABI-versioned initialization — Vendored: libwebp

This protects the public structure contract against caller and library version mismatch.

Public structures are initialized through inline wrappers that pass a library ABI version to internal initialization functions before callers use the fields.

Read more: [C ABI-versioned initialization](references/language-concepts.md#topic-c-abi-versioned-initialization)

Apply it through: WebPConfig, WebPPicture.

## Unit 185 — Vendored: libwebp

Shared subsystem: Vendored: libwebp.

Lesson ID: `unit-webp-riff-decoding-worker-parallelism-yuv-rgb-conversion`
Display slug: `unit-185-webp-chunk-parsing-incremental-decode-and-animation`

### Domain: WebP chunk parsing, incremental decode, and animation — Vendored: libwebp

The decoder, demuxer, and animation decoder are distinct modules over shared WebP byte data.

The WebP code implements an image decode pipeline that recognizes VP8 and VP8L payloads, incrementally retains input, demuxes RIFF chunks into frames, and composites animation frames.

Read more: [WebP chunk parsing, incremental decode, and animation](references/domain-concepts.md#topic-webp-riff-decoding)

Use these entity examples: WebP Decoder State.

### Domain: Worker-based parallelism — Vendored: libwebp

Thread behavior is abstracted behind a library interface instead of direct thread calls in encoder modules.

The encoding implementation obtains a WebPWorkerInterface and conditionally splits analysis or lossless work when thread-level configuration permits it.

Read more: [Worker-based parallelism](references/domain-concepts.md#topic-worker-parallelism)

Use these entity examples: VP8Encoder.

### Domain: YUV/RGB conversion and chroma processing — Vendored: libwebp

Format packers also target RGBA, BGRA, RGB565, and RGBA4444 outputs.

Conversion kernels turn RGB or ARGB input into YUV planes, reconstruct RGB outputs, and upsample chroma with neighboring samples.

Read more: [YUV/RGB conversion and chroma processing](references/domain-concepts.md#topic-yuv-rgb-conversion)

Use these entity examples: WebPPicture, WebPDecBuffer.

## Unit 186 — Vendored: libwebp

Shared subsystem: Vendored: libwebp.

Lesson ID: `unit-entropy-bitstreams-histograms-and-huffman-codes`
Display slug: `unit-186-entropy-bitstreams`

### Domain: Entropy bitstreams — Vendored: libwebp

The same utility area contains both writing and parsing primitives.

VP8 and VP8L writers serialize codec decisions into range-coded or raw bit-level byte streams while readers maintain cursor, range, and value state.

Read more: [Entropy bitstreams](references/domain-concepts.md#topic-entropy-bitstreams)

### Domain: Histograms and Huffman codes — Vendored: libwebp

The lossless encoder can use multiple histogram groups over image tiles.

Histograms built from backward-reference symbols are clustered and converted into Huffman code lengths, codes, and coded tree headers.

Read more: [Histograms and Huffman codes](references/domain-concepts.md#topic-histograms-and-huffman-codes)

Use these entity examples: VP8LBackwardRefs.

## Unit 187 — Vendored: manifold

Shared subsystem: Vendored: manifold.

Lesson ID: `unit-halfedge-topology-manifold-mesh-data-cpp-ownership-and-polymorphism`
Display slug: `unit-187-halfedge-topology`

### Domain: Halfedge topology — Vendored: manifold

Topology is separate from vertex positions and property vertices.

The Manifold mesh kernel represents every triangle boundary with directed Halfedge records whose pairedHalfedge links encode adjacency.

Read more: [Halfedge topology](references/domain-concepts.md#topic-halfedge-topology)

Use these entity examples: Halfedge, MeshGLP.

### Domain: Manifold mesh interchange — Vendored: manifold

This is the external mesh-shaped data carrier at the geometry boundary.

MeshGLP carries indexed triangles, per-vertex properties, merge information, and run metadata into Manifold construction.

Read more: [Manifold mesh interchange](references/domain-concepts.md#topic-manifold-mesh-data)

Use these entity examples: MeshGLP.

### Implementation: C++ ownership and polymorphic trees — Vendored: manifold

This is the ownership structure behind lazily composed CSG values.

C++ class templates support CSG nodes through shared ownership, enable_shared_from_this, derived leaf and operation classes, and the Manifold implementation boundary.

Read more: [C++ ownership and polymorphic trees](references/language-concepts.md#topic-cpp-ownership-and-polymorphism)

Apply it through: Manifold.

## Unit 188 — Vendored: manifold

Shared subsystem: Vendored: manifold.

Lesson ID: `unit-mesh-provenance`
Display slug: `unit-188-triangle-provenance`

### Domain: Triangle provenance — Vendored: manifold

The relation data is used during boolean-result construction and property transfer.

CSG output keeps triangle provenance through TriRef records and MeshRelationD triRef collections so result triangles can be related to source meshes.

Read more: [Triangle provenance](references/domain-concepts.md#topic-mesh-provenance)

Use these entity examples: TriRef, MeshRelationD.

## Unit 189 — Vendored: mbedtls

Shared subsystem: Vendored: mbedtls.

Lesson ID: `unit-tls-connection-state-c-stateful-struct-apis`
Display slug: `unit-189-tls-connection-and-session-state`

### Domain: TLS connection and session state — Vendored: mbedtls

Client, TLS 1.2, TLS 1.3, message, and common TLS units implement different portions of the protocol path.

The TLS state machine performs cryptographic operations while carrying configuration, negotiated session, handshake, record, and protocol state in mbedtls_ssl_context.

Read more: [TLS connection and session state](references/domain-concepts.md#topic-tls-connection-state)

Use these entity examples: mbedtls_ssl_session.

### Implementation: C stateful struct APIs — Vendored: mbedtls

PSA, Mbed TLS TLS, and message-processing code all use this idiom.

C APIs expose mutable state through struct pointers, initialize it explicitly, and communicate failures through status-returning functions and caller-owned buffers.

Read more: [C stateful struct APIs](references/language-concepts.md#topic-c-stateful-struct-apis)

Apply it through: psa_key_attributes_t, mbedtls_ssl_session.

## Unit 190 — Vendored: mbedtls

Shared subsystem: Vendored: mbedtls.

Lesson ID: `unit-x509-certificate-processing-psa-key-lifecycle`
Display slug: `unit-190-x-509-certificate-processing`

### Domain: X.509 certificate processing — Vendored: mbedtls

The implementation is split into certificate, request, revocation-list, OID, and writing translation units.

TLS credential processing parses, writes, and verifies certificates, certificate requests, revocation lists, names, and related ASN.1 data.

Read more: [X.509 certificate processing](references/domain-concepts.md#topic-x509-certificate-processing)

### Domain: PSA key lifecycle and storage — Vendored: mbedtls

The implementation distinguishes volatile identifier ranges and contains dedicated storage and slot-management code.

PSA key records retain configuration-selected algorithms, type, lifetime, identifier, and policy attributes from initialization or import through slot management and optional storage.

Read more: [PSA key lifecycle and storage](references/domain-concepts.md#topic-psa-key-lifecycle)

Use these entity examples: psa_key_attributes_t.

## Unit 191 — Vendored: meshoptimizer

Shared subsystem: Vendored: meshoptimizer.

Lesson ID: `unit-cpp-pluggable-allocation`
Display slug: `unit-191-c-pluggable-allocation`

### Implementation: C++ pluggable allocation — Vendored: meshoptimizer

The hook is explicit rather than relying on a global replacement of operator new.

meshoptimizer stores allocation and deallocation function pointers in static allocator storage so callers can replace the allocation policy.

Read more: [C++ pluggable allocation](references/language-concepts.md#topic-cpp-pluggable-allocation)

Apply it through: meshopt_Meshlet.

## Unit 192 — Vendored: metal-cpp

Shared subsystem: Vendored: metal-cpp.

Lesson ID: `unit-gpu-resources-pipelines-cpp-native-wrappers`
Display slug: `unit-192-c-inline-native-wrappers`

### Domain: GPU resources and pipelines — Vendored: metal-cpp

Descriptors configure GPU resources and pipeline state before a device produces allocations, functions, or pipeline state objects.

Read more: [GPU resources and pipelines](references/domain-concepts.md#topic-gpu-resources-pipelines)

Use these entity examples: MTL::Buffer, MTL::Texture, MTL4::PipelineDescriptor.

### Implementation: C++ inline native wrappers — Vendored: metal-cpp

Metal-cpp declares lightweight C++ classes and forwards calls through typed `Object::sendMessage` invocations.

C++ wrapper classes use inheritance and inline member functions to expose native framework methods.

Read more: [C++ inline native wrappers](references/language-concepts.md#topic-cpp-native-wrappers)

Apply it through: MTL4::Archive, MTL4::BinaryFunction.

## Unit 193 — Vendored: metal-cpp

Shared subsystem: Vendored: metal-cpp.

Lesson ID: `unit-gpu-command-encoding-metal-cpp-object-bridge-cpp-runtime-symbol-loading`
Display slug: `unit-193-gpu-command-encoding`

### Domain: GPU command encoding — Vendored: metal-cpp

Command buffers create render, compute, blit, and resource-state encoders that record work against configured GPU resources.

Read more: [GPU command encoding](references/domain-concepts.md#topic-gpu-command-encoding)

Use these entity examples: MTL::CommandBuffer, MTL::RenderCommandEncoder, MTL::ComputeCommandEncoder.

### Domain: Metal-cpp object bridge — Vendored: metal-cpp

The wrapper layer covers Metal, Metal 4, MetalFX, and QuartzCore headers.

The Metal portion models framework objects as C++ wrapper classes whose inline methods send selectors to the underlying object.

Read more: [Metal-cpp object bridge](references/domain-concepts.md#topic-metal-cpp-object-bridge)

Use these entity examples: MTL4::Archive.

### Implementation: C++ runtime symbol loading — Vendored: metal-cpp

Metal, MetalFX, and QuartzCore private headers use this pattern to bind symbols at runtime.

C++ wrapper types resolve framework symbols dynamically through `dlsym`-based helper templates.

Read more: [C++ runtime symbol loading](references/language-concepts.md#topic-cpp-runtime-symbol-loading)

Apply it through: MTL4::Archive.

## Unit 194 — Vendored: metal-cpp

Shared subsystem: Vendored: metal-cpp.

Lesson ID: `unit-metal-presentation-metalfx-scaling`
Display slug: `unit-194-metal-drawable-presentation`

### Domain: Metal drawable presentation — Vendored: metal-cpp

QuartzCore Metal layers produce drawables that implement the Metal drawable interface for presentation.

Read more: [Metal drawable presentation](references/domain-concepts.md#topic-metal-presentation)

Use these entity examples: CA::MetalLayer, CA::MetalDrawable.

### Domain: MetalFX scaling and interpolation — Vendored: metal-cpp

MetalFX descriptors create spatial, temporal, temporal-denoised, and frame-interpolation effect instances from a Metal device.

Read more: [MetalFX scaling and interpolation](references/domain-concepts.md#topic-metalfx-scaling)

Use these entity examples: MTLFX::SpatialScaler, MTLFX::TemporalScaler, MTLFX::FrameInterpolator.

## Unit 195 — Vendored: miniupnpc

Shared subsystem: Vendored: miniupnpc.

Lesson ID: `unit-upnp-device-discovery-upnp-port-mapping-c-macro-codecs`
Display slug: `unit-195-upnp-device-discovery`

### Domain: UPnP device discovery — Vendored: miniupnpc

MiniUPnPc discovers UPnP devices, represents them as a linked device list, and parses an IGD description into URLs and service data.

Read more: [UPnP device discovery](references/domain-concepts.md#topic-upnp-device-discovery)

Use these entity examples: UPNPDev, UPNPUrls, IGDdatas.

### Domain: UPnP port-mapping control — Vendored: miniupnpc

SOAP command helpers add, delete, query, and list port mappings on a discovered device.

Read more: [UPnP port-mapping control](references/domain-concepts.md#topic-upnp-port-mapping)

Use these entity examples: PortMapping, PortMappingParserData.

### Implementation: C macro-based binary decoding — Vendored: miniupnpc

The MiniUPnP decoder offers pointer, callback-reader, and bounds-checked variants of the same 7-bit encoding loop.

C macros implement variable-length decoding by shifting accumulated values and testing continuation bits.

Read more: [C macro-based binary decoding](references/language-concepts.md#topic-c-macro-codecs)

## Unit 196 — Vendored: miniupnpc

Shared subsystem: Vendored: miniupnpc.

Lesson ID: `unit-c-parser-state`
Display slug: `unit-196-c-parser-state-machines`

### Implementation: C parser state machines — Vendored: miniupnpc

The XML parser state stores the start, end, and current input positions for the parsing routines.

C structs retain XML cursor state and parser result heads while miniUPnPc parsing functions build protocol records.

Read more: [C parser state machines](references/language-concepts.md#topic-c-parser-state)

Apply it through: NameValueParserData, PortMappingParserData.

## Unit 197 — Vendored: minizip

Shared subsystem: Vendored: minizip.

Lesson ID: `unit-zip64-archive-io`
Display slug: `unit-197-zip64-archive-i-o`

### Domain: ZIP64 archive I/O — Vendored: minizip

MiniZip reads and writes ZIP archives while its I/O API abstracts file opening, seeking, telling, and allocation callbacks.

Read more: [ZIP64 archive I/O](references/domain-concepts.md#topic-zip64-archive-io)

Use these entity examples: zlib_filefunc64_32_def.

## Unit 198 — Vendored: misc

Shared subsystem: Vendored: misc.

Lesson ID: `unit-native-support-algorithms`
Display slug: `unit-198-native-support-algorithms`

### Domain: Native support algorithms — Vendored: misc

The support subtree implements deterministic PCG random state evolution alongside compression, noise, color, audio, and packing algorithms.

Read more: [Native support algorithms](references/domain-concepts.md#topic-native-support-algorithms)

Use these entity examples: pcg32_random_t.

## Unit 199 — Vendored: msdfgen

Shared subsystem: Vendored: msdfgen.

Lesson ID: `unit-multichannel-distance-fields`
Display slug: `unit-199-multi-channel-distance-fields`

### Domain: Multi-channel distance fields — Vendored: msdfgen

MSDFgen represents vector shapes, colors their edges, and generates distance-field pixels for one, three, or four channels.

Read more: [Multi-channel distance fields](references/domain-concepts.md#topic-multichannel-distance-fields)

Use these entity examples: msdfgen::Shape, msdfgen::BitmapSection<float, 3>.

## Unit 200 — Vendored: openxr

Shared subsystem: Vendored: openxr.

Lesson ID: `unit-openxr-runtime-loading-openxr-dispatch-cpp-polymorphic-ownership`
Display slug: `unit-200-openxr-runtime-loading`

### Domain: OpenXR runtime loading — Vendored: openxr

The embedded loader reads runtime and API-layer manifests, validates `next` chains during instance creation, and creates a loader instance.

Read more: [OpenXR runtime loading](references/domain-concepts.md#topic-openxr-runtime-loading)

Use these entity examples: ManifestFile, LoaderInstance, XrInstanceCreateInfo.

### Domain: OpenXR dispatch forwarding — Vendored: openxr

Generated loader entry points forward API calls through a dispatch table belonging to the selected runtime.

Read more: [OpenXR dispatch forwarding](references/domain-concepts.md#topic-openxr-dispatch)

Use these entity examples: XrGeneratedDispatchTableCore.

### Implementation: C++ polymorphism and ownership — Vendored: openxr

The OpenXR loader exposes abstract recorder types while factory functions return owning smart pointers.

C++ classes use public inheritance and `std::unique_ptr` to retain polymorphic loader services and dispatch state.

Read more: [C++ polymorphism and ownership](references/language-concepts.md#topic-cpp-polymorphic-ownership)

Apply it through: LoaderInstance, XrGeneratedDispatchTableCore.

## Unit 201 — Vendored: openxr

Shared subsystem: Vendored: openxr.

Lesson ID: `unit-openxr-structure-chains-c-abi-bridging`
Display slug: `unit-201-openxr-structure-chains`

### Domain: OpenXR structure chains — Vendored: openxr

OpenXR input and output structures carry a typed structure identifier and a `next` pointer for extensible structure chains.

Read more: [OpenXR structure chains](references/domain-concepts.md#topic-openxr-structure-chains)

Use these entity examples: XrInstanceCreateInfo, XrBaseInStructure.

### Implementation: C ABI bridging — Vendored: openxr

Calling-convention macros are selected per platform and are used in both exported functions and function-pointer declarations.

OpenXR headers use macros and `extern "C"` guards so C++ callers expose C-compatible API declarations and function-pointer types.

Read more: [C ABI bridging](references/language-concepts.md#topic-c-abi-bridging)

Apply it through: XrInstanceCreateInfo.

## Unit 202 — Vendored: openxr

Shared subsystem: Vendored: openxr.

Lesson ID: `unit-cpp-exception-abi-boundaries`
Display slug: `unit-202-c-exception-containment-at-abi-boundaries`

### Implementation: C++ exception containment at ABI boundaries — Vendored: openxr

The macros are attached to exported loader functions so exceptions do not escape the loader ABI.

OpenXR loader macros use C++ exceptions internally but convert `std::bad_alloc` and other `std::exception` failures into OpenXR result codes.

Read more: [C++ exception containment at ABI boundaries](references/language-concepts.md#topic-cpp-exception-abi-boundaries)

Apply it through: LoaderInstance.

## Unit 203 — Vendored: pcre2

Shared subsystem: Vendored: pcre2.

Lesson ID: `unit-regex-jit-codegen-regex-compile-match`
Display slug: `unit-203-regular-expression-jit-code-generation`

### Domain: Regular-expression JIT code generation — Vendored: pcre2

PCRE2 JIT compilation turns compiled pattern control flow into architecture-specific generated code using SLJIT labels and jumps.

Read more: [Regular-expression JIT code generation](references/domain-concepts.md#topic-regex-jit-codegen)

Use these entity examples: sljit_compiler, sljit_jump, sljit_label.

### Domain: Regular-expression compilation and matching — Vendored: pcre2

PCRE2 compiles patterns into code objects and matches subjects with traditional and DFA matching engines.

Read more: [Regular-expression compilation and matching](references/domain-concepts.md#topic-regex-compile-match)

Use these entity examples: pcre2_code, pcre2_match_data.

## Unit 204 — Vendored: re-spirv

Shared subsystem: Vendored: re-spirv.

Lesson ID: `unit-spirv-rewriting`
Display slug: `unit-204-spir-v-rewriting-and-optimization`

### Domain: SPIR-V rewriting and optimization — Vendored: re-spirv

re-spirv parses a word-addressed shader into instructions, functions, blocks, results, decorations, and specializations before optimization.

Read more: [SPIR-V rewriting and optimization](references/domain-concepts.md#topic-spirv-rewriting)

Use these entity examples: Shader, Instruction, Optimizer.

## Unit 205 — Vendored: recastnavigation

Shared subsystem: Vendored: recastnavigation.

Lesson ID: `unit-compact-heightfield-navmesh-contours-polygons-navmesh-heightfields`
Display slug: `unit-205-compact-heightfield-representation`

### Domain: Compact heightfield representation — Vendored: recastnavigation

Region algorithms index cells and spans through the heightfield's width and height.

A compact heightfield stores grid dimensions, compact cells, spans, and per-span area data for navigation processing.

Read more: [Compact heightfield representation](references/domain-concepts.md#topic-compact-heightfield)

Use these entity examples: rcCompactHeightfield.

### Domain: Navigation contours and polygons — Vendored: recastnavigation

Recast extracts contours from a compact heightfield, then builds polygon and detail meshes from those contours.

Read more: [Navigation contours and polygons](references/domain-concepts.md#topic-navmesh-contours-polygons)

Use these entity examples: rcContourSet, rcPolyMesh, rcPolyMeshDetail.

### Domain: Navigation heightfields — Vendored: recastnavigation

Recast rasterizes input triangles into a heightfield and converts its spans into a compact heightfield with adjacency.

Read more: [Navigation heightfields](references/domain-concepts.md#topic-navmesh-heightfields)

Use these entity examples: rcHeightfield, rcCompactHeightfield, rcCompactSpan.

## Unit 206 — Vendored: recastnavigation

Shared subsystem: Vendored: recastnavigation.

Lesson ID: `unit-navigation-regions`
Display slug: `unit-206-navigation-region-construction`

### Domain: Navigation region construction — Vendored: recastnavigation

The implementation contains watershed, monotone, and layer-oriented region-building paths.

Region construction labels connected compact heightfield spans into navigation regions.

Read more: [Navigation region construction](references/domain-concepts.md#topic-navigation-regions)

Use these entity examples: rcCompactHeightfield.

## Unit 207 — Vendored: sdl

Shared subsystem: Vendored: sdl.

Lesson ID: `unit-sdl-iostreams-c-structs-pointers`
Display slug: `unit-207-c-structs-and-pointer-linked-state`

### Domain: Abstract I/O streams — Vendored: sdl

The stream implementation also provides fixed-width endian read and write helpers.

SDL I/O streams wrap byte operations for file and memory implementations and expose an SDL property group.

Read more: [Abstract I/O streams](references/domain-concepts.md#topic-sdl-iostreams)

Use these entity examples: SDL_IOStream, SDL_PropertiesID.

### Implementation: C structs and pointer-linked state — Vendored: sdl

SDL uses named and anonymous structs together with pointers for queues, device records, property storage, and platform handles.

C structs and pointers represent explicit runtime state, ownership links, and opaque-handle implementation data.

Read more: [C structs and pointer-linked state](references/language-concepts.md#topic-c-structs-pointers)

Apply it through: SDL_EventEntry, SDL_hid_device_info.

## Unit 208 — Vendored: sdl

Shared subsystem: Vendored: sdl.

Lesson ID: `unit-sdl-platform-backends-gamepad-mapping-c-function-pointer-tables`
Display slug: `unit-208-c-function-pointer-tables`

### Domain: Compile-time platform backends — Vendored: sdl

The source tree contains platform-specific input, loading, timing, synchronization, and device paths.

SDL uses compile-time conditions to select Linux, Windows, macOS, dummy, and other platform backend implementations.

Read more: [Compile-time platform backends](references/domain-concepts.md#topic-sdl-platform-backends)

### Domain: Gamepad mapping and classification — Vendored: sdl

A built-in mapping database and controller-type rules are included.

Gamepad mapping classifies HID enumeration results and maps device controls to SDL gamepad axes and buttons.

Read more: [Gamepad mapping and classification](references/domain-concepts.md#topic-gamepad-mapping)

Use these entity examples: SDL_hid_device_info.

### Implementation: C function-pointer tables — Vendored: sdl

SDL declares native-library symbols and device-driver operations as typed function pointers.

C function-pointer tables attach implementation callbacks to C struct state for dynamically selected system and device services.

Read more: [C function-pointer tables](references/language-concepts.md#topic-c-function-pointer-tables)

Apply it through: SDL_hid_device_info.

## Unit 209 — Vendored: sdl

Shared subsystem: Vendored: sdl.

Lesson ID: `unit-hid-device-enumeration-sdl-runtime-properties-c-preprocessor-platform-selection`
Display slug: `unit-209-c-preprocessor-platform-selection`

### Domain: HID device enumeration and backends — Vendored: sdl

Linux, macOS, Windows, libusb, and SDL joystick HIDAPI code are present in this partition.

SDL enumerates HID devices into linked device-information records and routes opened devices through platform and driver backends.

Read more: [HID device enumeration and backends](references/domain-concepts.md#topic-hid-device-enumeration)

Use these entity examples: SDL_hid_device_info.

### Domain: Runtime property groups and hints — Vendored: sdl

The implementation uses properties in core services and stream metadata.

SDL runtime property groups associate named values with an ID, while hints retrieve configuration from property and environment sources.

Read more: [Runtime property groups and hints](references/domain-concepts.md#topic-sdl-runtime-properties)

Use these entity examples: SDL_PropertiesID.

### Implementation: C preprocessor platform selection — Vendored: sdl

SDL's source files and headers select platform code with feature macros and use extern C guards around C-facing declarations.

C preprocessor conditions select platform implementations and preserve a C linkage boundary for C++ compilation units.

Read more: [C preprocessor platform selection](references/language-concepts.md#topic-c-preprocessor-platform-selection)

## Unit 210 — Vendored: sdl

Shared subsystem: Vendored: sdl.

Lesson ID: `unit-sdl-event-queue-portable-math-fallbacks-c-concurrent-state`
Display slug: `unit-210-c-managed-concurrent-state`

### Domain: Event queue and watches — Vendored: sdl

Queued entries can carry temporary memory associated with an event payload.

The event runtime stores SDL_Event values in a mutex-protected linked queue with an atomic count and event-watch support.

Read more: [Event queue and watches](references/domain-concepts.md#topic-sdl-event-queue)

Use these entity examples: SDL_Event, SDL_EventEntry.

### Domain: Portable mathematical fallbacks — Vendored: sdl

The bundled code includes trigonometric, logarithmic, exponential, power, and floating-point classification operations.

SDL dispatches mathematical functions to platform library functions when present or to bundled fdlibm-derived implementations otherwise.

Read more: [Portable mathematical fallbacks](references/domain-concepts.md#topic-portable-math-fallbacks)

### Implementation: C-managed concurrent state — Vendored: sdl

SDL wraps synchronization behind its own mutex and atomic APIs while its C code explicitly maintains the protected data.

C structs hold mutexes, atomic counts, and linked entries when SDL manages concurrent event and device state.

Read more: [C-managed concurrent state](references/language-concepts.md#topic-c-concurrent-state)

Apply it through: SDL_EventEntry.

## Unit 211 — Vendored: sdl

Shared subsystem: Vendored: sdl.

Lesson ID: `unit-sdl-thread-abstractions`
Display slug: `unit-211-thread-and-synchronization-abstractions`

### Domain: Thread and synchronization abstractions — Vendored: sdl

The event queue and device code can use these facilities rather than directly naming every platform primitive.

SDL implements mutex, semaphore, condition, read/write lock, thread, and TLS abstractions with generic, pthread, and Windows backends.

Read more: [Thread and synchronization abstractions](references/domain-concepts.md#topic-sdl-thread-abstractions)

## Unit 212 — Vendored: spirv-reflect

Shared subsystem: Vendored: spirv-reflect.

Lesson ID: `unit-spirv-shader-reflection-c-abi-manual-lifetime`
Display slug: `unit-212-c-abi-structures-and-manual-lifetime`

### Domain: SPIR-V shader reflection — Vendored: spirv-reflect

The C header also provides a C++ ShaderModule wrapper when compiled as C++.

The reflection API creates a shader-module metadata view from SPIR-V code and exposes entry points, descriptors, interfaces, push constants, and specialization constants.

Read more: [SPIR-V shader reflection](references/domain-concepts.md#topic-spirv-shader-reflection)

Use these entity examples: SpvReflectShaderModule, SpvReflectEntryPoint.

### Implementation: C ABI structures and manual lifetime — Vendored: spirv-reflect

The header is usable from C and wraps its declarations in extern "C" when compiled as C++.

The SPIR-V reflection interface exchanges C structs and pointers through create, enumerate, query, change, and explicit destroy functions.

Read more: [C ABI structures and manual lifetime](references/language-concepts.md#topic-c-abi-manual-lifetime)

Apply it through: SpvReflectShaderModule, SpvReflectDescriptorSet, SpvReflectDescriptorBinding.

## Unit 213 — Vendored: spirv-reflect

Shared subsystem: Vendored: spirv-reflect.

Lesson ID: `unit-shader-interface-metadata-c-preprocessor-configuration`
Display slug: `unit-213-shader-interface-metadata`

### Domain: Shader interface metadata — Vendored: spirv-reflect

The module and entry-point structures use counts plus arrays or pointers to express the reflected relationships.

Reflection exposes descriptor sets, interface variables, push-constant blocks, and specialization constants from one reflected shader module; it depends on SPIR-V shader reflection because it is contained in the reflected shader module.

Read more: [Shader interface metadata](references/domain-concepts.md#topic-shader-interface-metadata)

Use these entity examples: SpvReflectDescriptorSet, SpvReflectDescriptorBinding, SpvReflectInterfaceVariable, SpvReflectBlockVariable, SpvReflectSpecializationConstant.

### Implementation: C preprocessor feature and platform configuration — Vendored: spirv-reflect

The inspected code also uses feature macros to compile TinyEXR's header implementation and platform branches in Volk.

Configuration headers use macros and conditional inclusion to enable ThorVG capabilities, select platform-dependent threading, and choose system or bundled SPIR-V headers.

Read more: [C preprocessor feature and platform configuration](references/language-concepts.md#topic-c-preprocessor-configuration)

## Unit 214 — Vendored: spirv-reflect

Shared subsystem: Vendored: spirv-reflect.

Lesson ID: `unit-shader-reflection`
Display slug: `unit-214-shader-reflection`

### Domain: Shader reflection — Vendored: spirv-reflect

It is implemented independently from SPIRV-Cross under a C API.

The reflection API reads SPIR-V IR and returns module, entry-point, descriptor-binding, interface-variable, and push-constant metadata.

Read more: [Shader reflection](references/domain-concepts.md#topic-shader-reflection)

Use these entity examples: SpvReflectShaderModule, SpvReflectDescriptorBinding.

## Unit 215 — Vendored: tinyexr

Shared subsystem: Vendored: tinyexr.

Lesson ID: `unit-openexr-image-decoding`
Display slug: `unit-215-openexr-image-decoding`

### Domain: OpenEXR image decoding — Vendored: tinyexr

The supplied partition demonstrates integration of a single-header image decoder rather than a ThorVG loader implementation.

TinyEXR is compiled as a header implementation with zlib included first, and its public header defines EXR-oriented image API data and functions.

Read more: [OpenEXR image decoding](references/domain-concepts.md#topic-openexr-image-decoding)

Use these entity examples: EXRImage.

## Unit 216 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-godot-thirdparty-adaptation-vulkan-pipeline-configuration-cpp-typed-api-records`
Display slug: `unit-216-c-typed-api-records`

### Domain: Godot-specific third-party adaptation — Vendored: vulkan

These files document local deltas from upstream vendored code.

The repository applies VMA allocator integration patches, redirects Vulkan includes to Godot's Vulkan header, adds lazy-allocation statistics, and configures portability macros for bundled dependencies.

Read more: [Godot-specific third-party adaptation](references/domain-concepts.md#topic-godot-thirdparty-adaptation)

### Domain: Graphics pipeline configuration — Vendored: vulkan

The header represents pipeline setup as a composed configuration record rather than a local rendering algorithm.

`GraphicsPipelineCreateInfo` groups shader stages with separate vertex-input, assembly, viewport, rasterization, multisample, depth-stencil, color-blend, and dynamic-state records.

Read more: [Graphics pipeline configuration](references/domain-concepts.md#topic-vulkan-pipeline-configuration)

Use these entity examples: GraphicsPipelineCreateInfo.

### Implementation: C++ typed API records — Vendored: vulkan

The C++ binding exposes API data as named records instead of raw untyped parameter lists.

Generated C++ structs use typed pointers and brace defaults to represent Vulkan creation inputs and optional state records.

Read more: [C++ typed API records](references/language-concepts.md#topic-cpp-typed-api-records)

Apply it through: GraphicsPipelineCreateInfo, PresentInfoKHR.

## Unit 217 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-vulkan-render-pass-configuration-vulkan-swapchain-presentation-c-abi-record-and-dispatch`
Display slug: `unit-217-c-abi-records-and-dispatch-signatures`

### Domain: Render-pass configuration — Vendored: vulkan

The supplied C++ binding also declares the newer `RenderPassCreateInfo2` form.

`RenderPassCreateInfo` combines attachment descriptions, subpass descriptions, and subpass dependencies into one externally exchanged rendering configuration.

Read more: [Render-pass configuration](references/domain-concepts.md#topic-vulkan-render-pass-configuration)

Use these entity examples: RenderPassCreateInfo.

### Domain: Swapchain presentation — Vendored: vulkan

Presentation is represented by configuration and submission records in the generated API layer.

The binding pairs swapchain configuration with a presentation request whose wait semaphores, swapchains, and image indices are array inputs.

Read more: [Swapchain presentation](references/domain-concepts.md#topic-vulkan-swapchain-presentation)

Use these entity examples: SwapchainCreateInfoKHR, PresentInfoKHR, SwapchainKHR.

### Implementation: C ABI records and dispatch signatures — Vendored: vulkan

The record is supplied by a caller; the dispatch signature names instance, creation input, allocation callbacks, and an output surface pointer.

`VkViSurfaceCreateInfoNN` and `PFN_vkCreateViSurfaceNN` express ABI-shaped call parameters through a tagged C record and a typed creation-function pointer.

Read more: [C ABI records and dispatch signatures](references/language-concepts.md#topic-c-abi-record-and-dispatch)

Apply it through: VkViSurfaceCreateInfoNN, VkSurfaceKHR.

## Unit 218 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-vi-native-surface-creation-vulkan-extensible-records-cpp-flag-stringification`
Display slug: `unit-218-c-flag-stringification`

### Domain: VI native surface creation — Vendored: vulkan

This is a platform-specific Vulkan surface extension declaration.

The `VK_NN_vi_surface` header declares input for creating a `VkSurfaceKHR` from a Nintendo VI window handle.

Read more: [VI native surface creation](references/domain-concepts.md#topic-vi-native-surface-creation)

Use these entity examples: VkViSurfaceCreateInfoNN, VkSurfaceKHR.

### Domain: Vulkan extensible records — Vendored: vulkan

The convention appears across the generated binding's externally exchanged configuration records.

Many Vulkan input records carry a `pNext` pointer, while the C VI surface record also carries an `sType` discriminator, forming an extensible record convention.

Read more: [Vulkan extensible records](references/domain-concepts.md#topic-vulkan-extensible-records)

Use these entity examples: VkViSurfaceCreateInfoNN, GraphicsPipelineCreateInfo.

### Implementation: C++ flag stringification — Vendored: vulkan

It is a diagnostic representation helper for the generated C++ Vulkan binding.

The helper converts typed `FormatFeatureFlags` values into a string by testing individual flag bits and appending their names.

Read more: [C++ flag stringification](references/language-concepts.md#topic-cpp-flag-stringification)

## Unit 219 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-extension-structure-chains-vulkan-instance-setup-cpp-c-type-mapping`
Display slug: `unit-219-c-mapping-of-c-api-types`

### Domain: Vulkan extension structure chains — Vendored: vulkan

The supplied platform records use the common Vulkan extensible-structure shape.

The structs expose `sType` and `pNext` fields so extension records can be carried through Vulkan creation and query calls.

Read more: [Vulkan extension structure chains](references/domain-concepts.md#topic-extension-structure-chains)

Use these entity examples: VkMetalSurfaceCreateInfoEXT, VkExternalFormatQNX.

### Domain: Vulkan instance setup — Vendored: vulkan

This is the outer configuration record for the Vulkan API surface represented by the supplied headers.

The binding models instance setup as an `InstanceCreateInfo` record containing optional application metadata plus enabled layer and extension name arrays.

Read more: [Vulkan instance setup](references/domain-concepts.md#topic-vulkan-instance-setup)

Use these entity examples: InstanceCreateInfo, ApplicationInfo.

### Implementation: C++ mapping of C API types — Vendored: vulkan

The generated mapping provides a type-level bridge between the C Vulkan ABI names and the C++ API layer.

`CppType<Vk...>` specializations associate C Vulkan type names with C++ wrapper types.

Read more: [C++ mapping of C API types](references/language-concepts.md#topic-cpp-c-type-mapping)

Apply it through: GraphicsPipelineCreateInfo.

## Unit 220 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-vulkan-video-session-configuration-cpp-static-abi-contracts`
Display slug: `unit-220-c-static-abi-contracts`

### Domain: Vulkan video session configuration — Vendored: vulkan

This partition declares video API data contracts; it does not show a video-processing implementation.

The generated records model a video session through a video profile and standard-header-version information, with separate records for decode and encode operations.

Read more: [Vulkan video session configuration](references/domain-concepts.md#topic-vulkan-video-session-configuration)

Use these entity examples: VideoSessionCreateInfoKHR.

### Implementation: C++ static ABI contracts — Vendored: vulkan

The static-assertion header checks wrapper size compatibility with C Vulkan handles and selected type traits.

The binding uses a macro-selected assertion facility to verify wrapper ABI layout and copy or move properties at compile time.

Read more: [C++ static ABI contracts](references/language-concepts.md#topic-cpp-static-abi-contracts)

## Unit 221 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-c-compatible-header-guards`
Display slug: `unit-221-c-compatible-guarded-headers`

### Implementation: C-compatible guarded headers — Vendored: vulkan

This is the boundary mechanism for the platform extension's C-facing declarations.

The VI declaration file uses preprocessor guards and conditionally enters `extern "C"` when compiled as C++, so one declaration set can be included by C and C++ translation units.

Read more: [C-compatible guarded headers](references/language-concepts.md#topic-c-compatible-header-guards)

Apply it through: VkViSurfaceCreateInfoNN.

## Unit 222 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-c-abi-data-structures`
Display slug: `unit-222-c-abi-structures-and-opaque-state`

### Implementation: C ABI structures and opaque state — Vendored: vulkan

Vulkan descriptors expose typed fields, while zlib hides its internal state behind a pointer.

C declarations use tagged structures, pointer members, and opaque implementation pointers to express caller-owned descriptors and library-managed state across API boundaries.

Read more: [C ABI structures and opaque state](references/language-concepts.md#topic-c-abi-data-structures)

Apply it through: VkWaylandSurfaceCreateInfoKHR, z_stream.

## Unit 223 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-c-function-pointer-apis`
Display slug: `unit-223-c-function-pointer-api-declarations`

### Implementation: C function-pointer API declarations — Vendored: vulkan

The generated Vulkan headers provide both PFN typedefs and optionally visible direct prototypes.

C function-pointer declarations expose addressable Vulkan entry points and callback-based frame/event integration without requiring a concrete implementation type.

Read more: [C function-pointer API declarations](references/language-concepts.md#topic-c-function-pointer-apis)

Apply it through: VkWaylandSurfaceCreateInfoKHR, VkSurfaceKHR.

## Unit 224 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-cpp-module-interface-cpp-single-header-implementation`
Display slug: `unit-224-c-module-interface-partition`

### Implementation: C++ module interface partition — Vendored: vulkan

The file is a C++ module interface artifact for the Vulkan video API namespace.

`vulkan_video.cppm` declares the exported `vulkan_hpp:video` module partition and an exported video namespace.

Read more: [C++ module interface partition](references/language-concepts.md#topic-cpp-module-interface)

Apply it through: VideoSessionCreateInfoKHR.

### Implementation: C++ single-header implementation selection — Vendored: vulkan

The .cpp file is intentionally minimal because VMA bodies are emitted from the included header.

VMA uses preprocessor configuration to place its header implementation in one C++ translation unit while selecting a non-MSVC debug macro before inclusion.

Read more: [C++ single-header implementation selection](references/language-concepts.md#topic-cpp-single-header-implementation)

Apply it through: VmaAllocatorCreateInfo.

## Unit 225 — Vendored: vulkan

Shared subsystem: Vendored: vulkan.

Lesson ID: `unit-cpp-module-import`
Display slug: `unit-225-c-standard-library-module-import`

### Implementation: C++ standard-library module import — Vendored: vulkan

This is a conditional module-aware inclusion path in the generated C++ headers.

The Vulkan-Hpp façade contains an import std declaration and other headers test VULKAN_HPP_CXX_MODULE before falling back to textual includes.

Read more: [C++ standard-library module import](references/language-concepts.md#topic-cpp-module-import)

## Unit 226 — Vendored: wayland-protocols

Shared subsystem: Vendored: wayland-protocols.

Lesson ID: `unit-explicit-drm-syncobj`
Display slug: `unit-226-explicit-drm-synchronization-objects`

### Domain: Explicit DRM synchronization objects — Vendored: wayland-protocols

The schema models synchronization as externally exchanged timeline objects and 64-bit point values.

Wayland protocol objects import DRM synchronization timelines and attach acquire and release timeline points to a surface commit.

Read more: [Explicit DRM synchronization objects](references/domain-concepts.md#topic-explicit-drm-syncobj)

Use these entity examples: wp_linux_drm_syncobj_timeline_v1.

## Unit 227 — Vendored: wslay

Shared subsystem: Vendored: wslay.

Lesson ID: `unit-websocket-frame-events`
Display slug: `unit-227-websocket-framing-and-event-contexts`

### Domain: WebSocket framing and event contexts — Vendored: wslay

The implementation is organized around incremental frame reads and writes plus event callbacks.

Wslay separates frame-level I/O from event-level message handling through frame contexts, event contexts, callback tables, message sources, and send/control queues.

Read more: [WebSocket framing and event contexts](references/domain-concepts.md#topic-websocket-frame-events)

Use these entity examples: wslay_event_context.

## Unit 228 — Vendored: zlib

Shared subsystem: Vendored: zlib.

Lesson ID: `unit-zlib-stream-codec`
Display slug: `unit-228-zlib-stream-compression`

### Domain: zlib stream compression — Vendored: zlib

Checksum, compression, and decompression code are separate source modules behind zlib.h.

The bundled zlib sources implement Adler-32 and CRC-32 checksums, DEFLATE compression trees, and inflate state for a public streaming compression interface.

Read more: [zlib stream compression](references/domain-concepts.md#topic-zlib-stream-codec)

Use these entity examples: z_stream.

## Unit 229 — Vendored: zstd

Shared subsystem: Vendored: zstd.

Lesson ID: `unit-zstandard-stream-codec`
Display slug: `unit-229-zstandard-stream-compression`

### Domain: Zstandard stream compression — Vendored: zstd

The implementation is divided into common, compress, and decompress subtrees.

The bundled code provides Zstandard stream compression contexts, decompression contexts, dictionaries, entropy tables, hash-based match finders, and optional worker pools.

Read more: [Zstandard stream compression](references/domain-concepts.md#topic-zstandard-stream-codec)

## Prerequisite graph

### domain
- `3d-gizmo-authoring` enables `editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `allocation` enables `buffer-storage` (synthesized lexical, confidence 0.25).
- `allocation` enables `class-reference-generation` (synthesized lexical, confidence 0.25).
- `allocation` enables `godot-thirdparty-adaptation` (synthesized lexical, confidence 0.25).
- `allocation` enables `gpu-memory-suballocation` (synthesized lexical, confidence 0.25).
- `allocation` enables `histograms-and-huffman-codes` (synthesized lexical, confidence 0.25).
- `allocation` enables `networking` (synthesized lexical, confidence 0.25).
- `allocation` enables `object-identity-lifetime` (synthesized lexical, confidence 0.25).
- `allocation` enables `object-model` (synthesized lexical, confidence 0.25).
- `allocation` enables `random-generation` (synthesized lexical, confidence 0.25).
- `allocation` enables `zip-archive-io` (synthesized lexical, confidence 0.25).
- `alpha-plane-coding` enables `encoder-configuration` (synthesized lexical, confidence 0.25).
- `alpha-plane-coding` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `alpha-plane-coding` enables `input-picture-planes` (synthesized lexical, confidence 0.25).
- `android-export-pipeline` enables `android-gradle-assembly` (grounded, confidence 0.90).
- `android-export-pipeline` enables `apk-signing` (grounded, confidence 0.90).
- `android-jni-interop` enables `android-plugin-signals` (grounded, confidence 0.90).
- `android-jni-interop` enables `android-storage-bridge` (grounded, confidence 0.90).
- `android-platform-host` enables `android-jni-interop` (grounded, confidence 0.90).
- `android-platform-host` enables `android-render-input` (grounded, confidence 0.90).
- `android-plugin-signals` enables `android-instrumented-tests` (grounded, confidence 0.90).
- `apple-embedded-hosting` enables `apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `astc-block-coding` enables `block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `basis-transcoding` enables `basis-container-layout` (grounded, confidence 0.90).
- `basis-transcoding` enables `block-texture-transcoding` (grounded, confidence 0.90).
- `basis-transcoding` enables `ktx2-container-transcoding` (grounded, confidence 0.90).
- `block-texture-transcoding` enables `astc-block-coding` (synthesized lexical, confidence 0.25).
- `block-texture-transcoding` enables `block-texture-encoding` (synthesized lexical, confidence 0.25).
- `broad-phase` enables `narrow-phase` (grounded, confidence 0.90).
- `brotli-stream-decompression` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `brotli-stream-decompression` enables `prefix-code-decoding` (synthesized lexical, confidence 0.25).
- `buffer-storage` enables `gltf-accessors` (synthesized lexical, confidence 0.25).
- `buffer-storage` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `buffer-storage` enables `web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `buffer-storage` enables `webp-image-io` (synthesized lexical, confidence 0.25).
- `bvh-construction` enables `allocation` (synthesized lexical, confidence 0.25).
- `bvh-construction` enables `builder-heuristics` (grounded, confidence 0.90).
- `bvh-construction` enables `bvh-traversal` (grounded, confidence 0.90).
- `bvh-construction` enables `motion-blur` (grounded, confidence 0.90).
- `bvh-construction` enables `physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `bvh-construction` enables `scene-commit` (grounded, confidence 0.90).
- `bvh-traversal` enables `canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `collada-import` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `expression-evaluation` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `networking` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `packed-scenes` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `ray-query` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `scene-importing` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `scene-tree` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `scene-tree-execution` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `canvas-and-viewports` enables `gui-controls` (grounded, confidence 0.90).
- `canvas-and-viewports` enables `platform-presentation` (grounded, confidence 0.90).
- `catmull-clark-patch-construction` enables `half-edge-topology` (synthesized lexical, confidence 0.25).
- `cff-font-subsetting` enables `font-driver-modules` (synthesized lexical, confidence 0.25).
- `cff-font-subsetting` enables `postscript-font-services` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `dynamic-variant` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `endian-safe-binary-io` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `managed-native-interop` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `project-upgrade` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `character-encoding-conversion` enables `yuv-rgb-conversion` (synthesized lexical, confidence 0.25).
- `codec-simd-specialization` enables `cpu-specialized-dsp` (synthesized lexical, confidence 0.25).
- `codec-simd-specialization` enables `simd-accelerated-codecs` (synthesized lexical, confidence 0.25).
- `collision-filtering` enables `broad-phase` (grounded, confidence 0.90).
- `collision-shapes` enables `collision-filtering` (grounded, confidence 0.90).
- `collision-shapes` enables `convex-collision` (grounded, confidence 0.90).
- `collision-shapes` enables `geometry-preprocessing` (grounded, confidence 0.90).
- `collision-shapes` enables `narrow-phase` (grounded, confidence 0.90).
- `collision-shapes` enables `physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `collision-shapes` enables `physics-3d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `collision-shapes` enables `physics-collision` (synthesized lexical, confidence 0.25).
- `collision-shapes` enables `physics-queries` (synthesized lexical, confidence 0.25).
- `collision-shapes` enables `physics-server-queries` (synthesized lexical, confidence 0.25).
- `collision-shapes` enables `physics-space-queries` (synthesized lexical, confidence 0.25).
- `collision-shapes` enables `textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `color-font-paint` enables `raster-font-rendering` (grounded, confidence 0.90).
- `color-font-paint` enables `vector-font-export` (grounded, confidence 0.90).
- `compact-heightfield` enables `navigation-regions` (grounded, confidence 0.90).
- `compact-heightfield` enables `navmesh-contours-polygons` (synthesized lexical, confidence 0.25).
- `compact-heightfield` enables `navmesh-heightfields` (synthesized lexical, confidence 0.25).
- `completion-contexts` enables `gdscript-editor-services` (synthesized lexical, confidence 0.25).
- `completion-contexts` enables `scene-contexts` (grounded, confidence 0.90).
- `concurrency` enables `device-runtime` (synthesized lexical, confidence 0.25).
- `concurrency` enables `encoder-configuration` (synthesized lexical, confidence 0.25).
- `concurrency` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `concurrency` enables `sdl-event-queue` (synthesized lexical, confidence 0.25).
- `concurrency` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `concurrency` enables `synchronization-primitives` (synthesized lexical, confidence 0.25).
- `concurrency` enables `task-scheduling` (synthesized lexical, confidence 0.25).
- `concurrency` enables `wayland-window-backend` (synthesized lexical, confidence 0.25).
- `concurrency` enables `worker-parallelism` (synthesized lexical, confidence 0.25).
- `constraints` enables `skeletal-ragdoll` (grounded, confidence 0.90).
- `constraints` enables `state-recording` (grounded, confidence 0.90).
- `constraints` enables `vehicle-dynamics` (grounded, confidence 0.90).
- `contact-management` enables `constraints` (grounded, confidence 0.90).
- `convex-collision` enables `narrow-phase` (grounded, confidence 0.90).
- `convex-collision` enables `physics-3d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `cpu-specialized-dsp` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `cpu-specialized-dsp` enables `simd-ray-packets` (synthesized lexical, confidence 0.25).
- `crypto-resources` enables `cryptography` (synthesized lexical, confidence 0.25).
- `cryptography` enables `crypto-resources` (synthesized lexical, confidence 0.25).
- `cryptography` enables `resource-formats` (synthesized lexical, confidence 0.25).
- `cryptography` enables `tls-crypto-backend` (synthesized lexical, confidence 0.25).
- `curve-and-patch-bases` enables `subdivision-surface-evaluation` (grounded, confidence 0.90).
- `device-runtime` enables `diagnostic-expectations` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `dynamic-variant` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `engine-bootstrap` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `gamepad-mapping` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `gpu-resources-pipelines` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `hid-device-enumeration` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `metalfx-scaling` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `object-model` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `openxr-dispatch` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `openxr-render-models` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `openxr-runtime-integration` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `openxr-runtime-loading` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `random-generation` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `reflection` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `rendering-backends` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `run-management` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `runtime-configuration` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `runtime-loop-time` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `runtime-metadata` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `scene-commit` (grounded, confidence 0.90).
- `device-runtime` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `sdl-event-queue` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `device-runtime` enables `upnp-port-mapping` (synthesized lexical, confidence 0.25).
- `diagnostic-expectations` enables `device-runtime` (synthesized lexical, confidence 0.25).
- `diagnostic-expectations` enables `image-quality-metrics` (synthesized lexical, confidence 0.25).
- `dynamic-invocation-signals` enables `class-reference-generation` (synthesized lexical, confidence 0.25).
- `dynamic-invocation-signals` enables `deferred-execution` (grounded, confidence 0.90).
- `dynamic-invocation-signals` enables `dynamic-values` (synthesized lexical, confidence 0.25).
- `dynamic-invocation-signals` enables `signal-awaitability` (synthesized lexical, confidence 0.25).
- `dynamic-invocation-signals` enables `undo-redo` (grounded, confidence 0.90).
- `dynamic-values` enables `android-jni-interop` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `break-iteration` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `dynamic-invocation-signals` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `dynamic-variant` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `expression-evaluation` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `generic-containers` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `managed-native-interop` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `network-data-exchange` (grounded, confidence 0.90).
- `dynamic-values` enables `project-settings` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `variant-containers` (synthesized lexical, confidence 0.25).
- `dynamic-values` enables `variant-text-serialization` (synthesized lexical, confidence 0.25).
- `dynamic-variant` enables `android-jni-interop` (synthesized lexical, confidence 0.25).
- `dynamic-variant` enables `dynamic-invocation-signals` (grounded, confidence 0.90).
- `dynamic-variant` enables `dynamic-values` (synthesized lexical, confidence 0.25).
- `dynamic-variant` enables `expression-evaluation` (grounded, confidence 0.90).
- `dynamic-variant` enables `managed-native-interop` (synthesized lexical, confidence 0.25).
- `dynamic-variant` enables `network-data-exchange` (synthesized lexical, confidence 0.25).
- `dynamic-variant` enables `project-settings` (synthesized lexical, confidence 0.25).
- `dynamic-variant` enables `variant-containers` (grounded, confidence 0.90).
- `dynamic-variant` enables `variant-text-serialization` (grounded, confidence 0.90).
- `editor-and-project-settings` enables `android-export-pipeline` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `audio-bus-processing` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `completion-contexts` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `engine-bootstrap` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `frame-timing` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `high-level-rpc` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `input-action-configuration` (grounded, confidence 0.90).
- `editor-and-project-settings` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `openxr-action-configuration` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `psa-key-lifecycle` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `rendering-resources` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `runtime-configuration` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `tls-connection-state` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `vulkan-render-pass-configuration` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `editor-and-project-settings` enables `worker-parallelism` (synthesized lexical, confidence 0.25).
- `editor-extensibility` enables `editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `editor-extensibility` enables `inspector-property-editors` (synthesized lexical, confidence 0.25).
- `editor-plugin-extension` enables `android-export-pipeline` (synthesized lexical, confidence 0.25).
- `editor-plugin-extension` enables `apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `editor-plugin-extension` enables `editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `editor-plugin-extension` enables `editor-plugin-state` (synthesized lexical, confidence 0.25).
- `editor-plugin-extension` enables `editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `editor-plugin-extension` enables `inspector-property-editors` (synthesized lexical, confidence 0.25).
- `editor-plugin-extension` enables `platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `editor-plugin-lifecycle` enables `android-export-pipeline` (synthesized lexical, confidence 0.25).
- `editor-plugin-lifecycle` enables `apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `editor-plugin-lifecycle` enables `editor-plugin-state` (synthesized lexical, confidence 0.25).
- `editor-plugin-lifecycle` enables `editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `editor-plugin-lifecycle` enables `inspector-property-editors` (grounded, confidence 0.90).
- `editor-plugin-lifecycle` enables `platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `editor-scene-sessions` enables `editing-selection-history` (grounded, confidence 0.90).
- `editor-scene-sessions` enables `editor-plugin-state` (grounded, confidence 0.90).
- `editor-scene-sessions` enables `undo-redo-history` (grounded, confidence 0.90).
- `editor-theming` enables `themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `editor-theming` enables `ui-theming` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `android-export-pipeline` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `audio-bus-processing` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `completion-contexts` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `editor-and-project-settings` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `engine-bootstrap` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `frame-timing` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `high-level-rpc` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `input-action-configuration` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `lossy-macroblock-encoding` (grounded, confidence 0.90).
- `encoder-configuration` enables `openxr-action-configuration` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `psa-key-lifecycle` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `rendering-resources` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `runtime-configuration` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `tls-connection-state` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `vulkan-render-pass-configuration` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `encoder-configuration` enables `worker-parallelism` (grounded, confidence 0.90).
- `enet-host-peer-transport` enables `enet-wire-commands` (grounded, confidence 0.90).
- `enet-host-peer-transport` enables `godot-enet-socket-adaptation` (synthesized lexical, confidence 0.25).
- `enet-host-peer-transport` enables `wraparound-network-time` (synthesized lexical, confidence 0.25).
- `enet-transport-and-multiplayer` enables `enet-host-peer-transport` (synthesized lexical, confidence 0.25).
- `enet-transport-and-multiplayer` enables `enet-wire-commands` (synthesized lexical, confidence 0.25).
- `enet-transport-and-multiplayer` enables `godot-enet-socket-adaptation` (synthesized lexical, confidence 0.25).
- `enet-transport-and-multiplayer` enables `networking` (synthesized lexical, confidence 0.25).
- `enet-transport-and-multiplayer` enables `wraparound-network-time` (synthesized lexical, confidence 0.25).
- `enet-wire-commands` enables `godot-enet-socket-adaptation` (grounded, confidence 0.90).
- `engine-bootstrap` enables `android-platform-host` (grounded, confidence 0.90).
- `engine-bootstrap` enables `frame-timing` (grounded, confidence 0.90).
- `engine-bootstrap` enables `performance-monitors` (grounded, confidence 0.90).
- `entropy-bitstreams` enables `endian-safe-binary-io` (synthesized lexical, confidence 0.25).
- `export-orchestration` enables `target-platform-export` (grounded, confidence 0.90).
- `export-presets` enables `export-orchestration` (grounded, confidence 0.90).
- `extension-api-compatibility` enables `gdextension-libraries` (synthesized lexical, confidence 0.25).
- `extension-api-compatibility` enables `native-extensions` (synthesized lexical, confidence 0.25).
- `extensions` enables `extension-api-compatibility` (synthesized lexical, confidence 0.25).
- `extensions` enables `extension-structure-chains` (synthesized lexical, confidence 0.25).
- `extensions` enables `gdextension-libraries` (synthesized lexical, confidence 0.25).
- `extensions` enables `native-extensions` (synthesized lexical, confidence 0.25).
- `extensions` enables `openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `extensions` enables `openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `extensions` enables `openxr-render-models` (synthesized lexical, confidence 0.25).
- `extensions` enables `pluggable-server-backends` (synthesized lexical, confidence 0.25).
- `extensions` enables `resource-formats` (synthesized lexical, confidence 0.25).
- `extensions` enables `script-hosting` (synthesized lexical, confidence 0.25).
- `extensions` enables `vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `fbx-gltf-scene-import` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `fbx-gltf-scene-import` enables `gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `filesystem-project-index` enables `export-presets` (grounded, confidence 0.90).
- `filesystem-project-index` enables `resource-dependency-resolution` (grounded, confidence 0.90).
- `filter-callbacks` enables `font-streams` (synthesized lexical, confidence 0.25).
- `filter-callbacks` enables `web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `filter-callbacks` enables `websocket-frame-events` (synthesized lexical, confidence 0.25).
- `fixture-contracts` enables `completion-contexts` (grounded, confidence 0.90).
- `fixture-contracts` enables `diagnostic-expectations` (grounded, confidence 0.90).
- `fixture-contracts` enables `lsp-semantic-fixtures` (grounded, confidence 0.90).
- `font-driver-modules` enables `glyph-caching` (grounded, confidence 0.90).
- `font-driver-modules` enables `glyph-rasterization` (grounded, confidence 0.90).
- `font-driver-modules` enables `postscript-font-services` (grounded, confidence 0.90).
- `font-streams` enables `compressed-font-stream-adapters` (grounded, confidence 0.90).
- `font-streams` enables `sfnt-tables` (grounded, confidence 0.90).
- `font-subsetting` enables `cff-font-subsetting` (grounded, confidence 0.90).
- `font-subsetting` enables `opentype-table-routing` (grounded, confidence 0.90).
- `font-subsetting` enables `variable-font-subsetting` (grounded, confidence 0.90).
- `font-table-access` enables `graphite-rule-execution` (grounded, confidence 0.90).
- `font-table-access` enables `graphite-shaping` (grounded, confidence 0.90).
- `font-table-access` enables `sfnt-tables` (synthesized lexical, confidence 0.25).
- `freetype-module-composition` enables `font-driver-modules` (synthesized lexical, confidence 0.25).
- `gdextension-libraries` enables `native-extensions` (synthesized lexical, confidence 0.25).
- `gdscript-bytecode-runtime` enables `graphite-rule-execution` (synthesized lexical, confidence 0.25).
- `gdscript-editor-services` enables `shader-editing-and-language-plugins` (synthesized lexical, confidence 0.25).
- `gdscript-static-analysis` enables `gdscript-bytecode-runtime` (grounded, confidence 0.90).
- `geometry-families` enables `collision-shapes` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `geometry-preprocessing` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `geometry-resources` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `halfedge-topology` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `mesh-provenance` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `simd-ray-packets` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `spatial-indexing` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `subdivision-surface-evaluation` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `subgrid-intersection` (synthesized lexical, confidence 0.25).
- `geometry-families` enables `triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `buffer-storage` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `bvh-traversal` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `catmull-clark-patch-construction` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `filter-callbacks` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `geometry-families` (grounded, confidence 0.90).
- `geometry-resources` enables `geometry-preprocessing` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `hit-results` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `instancing` (grounded, confidence 0.90).
- `geometry-resources` enables `mesh-geometry-algorithms` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `motion-blur` (grounded, confidence 0.90).
- `geometry-resources` enables `navigation-mesh-construction` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `primitive-intersection` (grounded, confidence 0.90).
- `geometry-resources` enables `primitive-references` (grounded, confidence 0.90).
- `geometry-resources` enables `ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `rendering-assets` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `scene-commit` (grounded, confidence 0.90).
- `geometry-resources` enables `spatial-indexing` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `transform-interpolation` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `geometry-resources` enables `two-dimensional-content` (synthesized lexical, confidence 0.25).
- `geometry-transforms` enables `mesh-geometry-algorithms` (grounded, confidence 0.90).
- `geometry-transforms` enables `spatial-indexing` (grounded, confidence 0.90).
- `geometry-transforms` enables `transform-interpolation` (grounded, confidence 0.90).
- `gltf-materials-textures` enables `basis-container-layout` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `basis-transcoding` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `ktx-texture-container` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `ktx2-container-transcoding` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `raster-image-input` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `rendering-resources` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `resource-importing` (synthesized lexical, confidence 0.25).
- `gltf-materials-textures` enables `textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `collada-import` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `expression-evaluation` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `networking` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `packed-scenes` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `scene-importing` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `scene-tree` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `scene-tree-execution` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `gltf-node-hierarchy` enables `visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `glyph-rasterization` enables `signed-distance-fields` (grounded, confidence 0.90).
- `glyph-rasterization` enables `vulkan-pipeline-configuration` (synthesized lexical, confidence 0.25).
- `godot-enet-socket-adaptation` enables `network-io` (synthesized lexical, confidence 0.25).
- `godot-enet-socket-adaptation` enables `tls-crypto-backend` (synthesized lexical, confidence 0.25).
- `gpu-command-encoding` enables `lossless-transform-pipeline` (synthesized lexical, confidence 0.25).
- `gpu-command-encoding` enables `lossy-macroblock-encoding` (synthesized lexical, confidence 0.25).
- `gpu-command-encoding` enables `metal-presentation` (grounded, confidence 0.90).
- `gpu-resources-pipelines` enables `android-gradle-assembly` (synthesized lexical, confidence 0.25).
- `gpu-resources-pipelines` enables `apk-signing` (synthesized lexical, confidence 0.25).
- `gpu-resources-pipelines` enables `cff-font-subsetting` (synthesized lexical, confidence 0.25).
- `gpu-resources-pipelines` enables `gpu-command-encoding` (grounded, confidence 0.90).
- `gpu-resources-pipelines` enables `jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `gpu-resources-pipelines` enables `metalfx-scaling` (grounded, confidence 0.90).
- `gpu-resources-pipelines` enables `png-stream-transforms` (synthesized lexical, confidence 0.25).
- `gpu-resources-pipelines` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `gpu-resources-pipelines` enables `variable-font-subsetting` (synthesized lexical, confidence 0.25).
- `gpu-resources-pipelines` enables `webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `graphite-rule-execution` enables `graphite-shaping` (grounded, confidence 0.90).
- `gui-control-layout` enables `curve-and-patch-bases` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `gui-controls` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `platform-presentation` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `regex-jit-codegen` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `rendering-assets` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `ui-composition` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `ui-theming` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `version-control-integration` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `websocket-frame-events` (synthesized lexical, confidence 0.25).
- `gui-controls` enables `gui-control-layout` (synthesized lexical, confidence 0.25).
- `gui-controls` enables `text-interface` (grounded, confidence 0.90).
- `half-edge-topology` enables `catmull-clark-patch-construction` (grounded, confidence 0.90).
- `halfedge-topology` enables `half-edge-topology` (synthesized lexical, confidence 0.25).
- `hid-device-enumeration` enables `gamepad-mapping` (grounded, confidence 0.90).
- `high-level-rpc` enables `godot-member-exposure` (synthesized lexical, confidence 0.25).
- `high-level-rpc` enables `network-data-exchange` (synthesized lexical, confidence 0.25).
- `histograms-and-huffman-codes` enables `brotli-stream-decompression` (synthesized lexical, confidence 0.25).
- `histograms-and-huffman-codes` enables `prefix-code-decoding` (synthesized lexical, confidence 0.25).
- `hit-results` enables `filter-callbacks` (grounded, confidence 0.90).
- `hit-results` enables `packet-queries` (grounded, confidence 0.90).
- `hit-results` enables `primitive-intersection` (grounded, confidence 0.90).
- `hit-results` enables `ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `hit-results` enables `triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `image-codec-registration` enables `astc-block-coding` (synthesized lexical, confidence 0.25).
- `image-codec-registration` enables `block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `image-decode-pipeline` enables `brotli-stream-decompression` (synthesized lexical, confidence 0.25).
- `image-decode-pipeline` enables `jpeg-decompression-transcoding` (grounded, confidence 0.90).
- `image-decode-pipeline` enables `png-stream-transforms` (grounded, confidence 0.90).
- `image-decode-pipeline` enables `webp-riff-decoding` (grounded, confidence 0.90).
- `image-format-loading` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `image-format-loading` enables `jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `image-format-loading` enables `ktx-texture-container` (synthesized lexical, confidence 0.25).
- `image-format-loading` enables `raster-image-input` (synthesized lexical, confidence 0.25).
- `image-quality-metrics` enables `sfnt-tables` (synthesized lexical, confidence 0.25).
- `input-actions` enables `input-events-actions` (synthesized lexical, confidence 0.25).
- `input-actions` enables `platform-display-adaptation` (synthesized lexical, confidence 0.25).
- `input-events-actions` enables `input-actions` (synthesized lexical, confidence 0.25).
- `input-events-actions` enables `platform-display-adaptation` (grounded, confidence 0.90).
- `input-midi` enables `apple-embedded-hosting` (synthesized lexical, confidence 0.25).
- `input-midi` enables `browser-runtime-adapters` (synthesized lexical, confidence 0.25).
- `input-midi` enables `input-events-actions` (synthesized lexical, confidence 0.25).
- `input-midi` enables `midi-input-adapters` (synthesized lexical, confidence 0.25).
- `input-picture-planes` enables `alpha-plane-coding` (grounded, confidence 0.90).
- `input-picture-planes` enables `encoder-configuration` (grounded, confidence 0.90).
- `input-picture-planes` enables `geometry-transforms` (synthesized lexical, confidence 0.25).
- `input-picture-planes` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `input-picture-planes` enables `image-quality-metrics` (grounded, confidence 0.90).
- `input-picture-planes` enables `lossless-transform-pipeline` (grounded, confidence 0.90).
- `input-picture-planes` enables `lossy-macroblock-encoding` (grounded, confidence 0.90).
- `input-picture-planes` enables `spatial-predictive-filters` (grounded, confidence 0.90).
- `input-picture-planes` enables `yuv-rgb-conversion` (grounded, confidence 0.90).
- `instancing` enables `geometry-resources` (synthesized lexical, confidence 0.25).
- `instancing` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `instancing` enables `hit-results` (synthesized lexical, confidence 0.25).
- `instancing` enables `openxr-runtime-loading` (synthesized lexical, confidence 0.25).
- `instancing` enables `rendering-assets` (synthesized lexical, confidence 0.25).
- `instancing` enables `rigid-body-motion` (synthesized lexical, confidence 0.25).
- `instancing` enables `scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `instancing` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `instancing` enables `vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `jpeg-decompression-transcoding` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `jpeg-decompression-transcoding` enables `raster-image-input` (synthesized lexical, confidence 0.25).
- `ktx-texture-container` enables `image-format-loading` (synthesized lexical, confidence 0.25).
- `ktx2-container-transcoding` enables `ktx-texture-container` (synthesized lexical, confidence 0.25).
- `ktx2-container-transcoding` enables `texture-compression-pipeline` (synthesized lexical, confidence 0.25).
- `linuxbsd-desktop-integration` enables `wayland-window-backend` (grounded, confidence 0.90).
- `localization` enables `text-and-localization` (synthesized lexical, confidence 0.25).
- `lossless-transform-pipeline` enables `entropy-bitstreams` (synthesized lexical, confidence 0.25).
- `lossless-transform-pipeline` enables `webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `lossy-macroblock-encoding` enables `transform-quantization-rate-distortion` (grounded, confidence 0.90).
- `lsp-semantic-fixtures` enables `gdscript-language-server` (synthesized lexical, confidence 0.25).
- `manifold-mesh-data` enables `halfedge-topology` (grounded, confidence 0.90).
- `metal-cpp-object-bridge` enables `gpu-resources-pipelines` (grounded, confidence 0.90).
- `metal-cpp-object-bridge` enables `metal-presentation` (synthesized lexical, confidence 0.25).
- `metal-cpp-object-bridge` enables `metalfx-scaling` (synthesized lexical, confidence 0.25).
- `metal-cpp-object-bridge` enables `platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `metal-cpp-object-bridge` enables `rendering-backends` (synthesized lexical, confidence 0.25).
- `module-build-configuration` enables `editor-build-composition` (synthesized lexical, confidence 0.25).
- `module-build-configuration` enables `platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `android-gradle-assembly` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `builder-heuristics` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `constraints` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `cpu-specialized-dsp` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `physics-queries` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `physics-server-queries` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `physics-space-queries` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `project-catalog` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `ray-query` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `rigid-body-motion` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `runtime-loop-time` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `sdl-platform-backends` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `wraparound-network-time` (synthesized lexical, confidence 0.25).
- `motion-blur-geometry` enables `motion-blur` (synthesized lexical, confidence 0.25).
- `narrow-phase` enables `contact-management` (grounded, confidence 0.90).
- `narrow-phase` enables `vehicle-dynamics` (grounded, confidence 0.90).
- `native-extensions` enables `gdextension-libraries` (synthesized lexical, confidence 0.25).
- `native-platform-adapters` enables `apple-embedded-hosting` (grounded, confidence 0.90).
- `native-platform-adapters` enables `audio-backend-adapters` (grounded, confidence 0.90).
- `native-platform-adapters` enables `midi-input-adapters` (grounded, confidence 0.90).
- `navigation-agents` enables `compact-heightfield` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `editing-selection-history` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `half-edge-topology` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `navigation` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `navigation-map-composition` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `navigation-mesh-construction` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `navigation-queries` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `navigation-regions` (synthesized lexical, confidence 0.25).
- `navigation-agents` enables `tilemap-layer-data` (synthesized lexical, confidence 0.25).
- `navigation-map-composition` enables `navigation-avoidance` (grounded, confidence 0.90).
- `navigation-map-composition` enables `navigation-path-queries` (grounded, confidence 0.90).
- `navigation-regions` enables `navigation` (synthesized lexical, confidence 0.25).
- `navigation-regions` enables `navigation-agents` (synthesized lexical, confidence 0.25).
- `navigation-regions` enables `navigation-map-composition` (synthesized lexical, confidence 0.25).
- `navmesh-heightfields` enables `navmesh-contours-polygons` (grounded, confidence 0.90).
- `network-data-exchange` enables `webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `network-io` enables `webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `dynamic-invocation-signals` (grounded, confidence 0.90).
- `object-identity-lifetime` enables `navigation-queries` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `object-model` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `physics-server-queries` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `physics-space-queries` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `reflection-metadata` (grounded, confidence 0.90).
- `object-identity-lifetime` enables `regular-expression-results` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `resources` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `script-hosting` (grounded, confidence 0.90).
- `object-identity-lifetime` enables `upnp-device-control` (synthesized lexical, confidence 0.25).
- `object-model` enables `collision-filtering` (synthesized lexical, confidence 0.25).
- `object-model` enables `deferred-execution` (synthesized lexical, confidence 0.25).
- `object-model` enables `dynamic-invocation-signals` (synthesized lexical, confidence 0.25).
- `object-model` enables `dynamic-values` (synthesized lexical, confidence 0.25).
- `object-model` enables `metal-cpp-object-bridge` (synthesized lexical, confidence 0.25).
- `object-model` enables `navigation-queries` (synthesized lexical, confidence 0.25).
- `object-model` enables `object-identity-lifetime` (synthesized lexical, confidence 0.25).
- `object-model` enables `objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `object-model` enables `physics-server-queries` (synthesized lexical, confidence 0.25).
- `object-model` enables `physics-space-queries` (synthesized lexical, confidence 0.25).
- `object-model` enables `property-inspection` (synthesized lexical, confidence 0.25).
- `object-model` enables `reflection` (grounded, confidence 0.90).
- `object-model` enables `reflection-metadata` (synthesized lexical, confidence 0.25).
- `object-model` enables `regular-expression-results` (synthesized lexical, confidence 0.25).
- `object-model` enables `resources` (synthesized lexical, confidence 0.25).
- `object-model` enables `runtime-loop-time` (synthesized lexical, confidence 0.25).
- `object-model` enables `scene-tree` (grounded, confidence 0.90).
- `object-model` enables `script-hosting` (synthesized lexical, confidence 0.25).
- `object-model` enables `scripting` (synthesized lexical, confidence 0.25).
- `object-model` enables `serialization` (synthesized lexical, confidence 0.25).
- `object-model` enables `text-shaping-plans` (synthesized lexical, confidence 0.25).
- `object-model` enables `undo-redo` (synthesized lexical, confidence 0.25).
- `object-model` enables `upnp-device-control` (synthesized lexical, confidence 0.25).
- `object-model` enables `web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `ogg-packet-sequences` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `ogg-packet-sequences` enables `vorbis-block-synthesis` (synthesized lexical, confidence 0.25).
- `ogg-pages-and-packets` enables `theora-block-video-codec` (grounded, confidence 0.90).
- `ogg-pages-and-packets` enables `vorbis-block-synthesis` (grounded, confidence 0.90).
- `openxr-dispatch` enables `character-encoding-conversion` (synthesized lexical, confidence 0.25).
- `openxr-dispatch` enables `collision-shapes` (synthesized lexical, confidence 0.25).
- `openxr-dispatch` enables `ktx-texture-container` (synthesized lexical, confidence 0.25).
- `openxr-dispatch` enables `temporal-upscaling-dispatch` (synthesized lexical, confidence 0.25).
- `openxr-runtime-integration` enables `openxr-binding-modifiers` (synthesized lexical, confidence 0.25).
- `openxr-runtime-integration` enables `openxr-composition-layers` (grounded, confidence 0.90).
- `openxr-runtime-integration` enables `openxr-extension-wrappers` (grounded, confidence 0.90).
- `openxr-runtime-integration` enables `openxr-render-models` (grounded, confidence 0.90).
- `openxr-runtime-integration` enables `openxr-structure-chains` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `diagnostic-expectations` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `dynamic-variant` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `engine-bootstrap` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `object-model` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `openxr-dispatch` (grounded, confidence 0.90).
- `openxr-runtime-loading` enables `openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `openxr-render-models` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `openxr-runtime-integration` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `random-generation` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `reflection` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `runtime-configuration` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `runtime-loop-time` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `runtime-metadata` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `sdl-event-queue` (synthesized lexical, confidence 0.25).
- `openxr-runtime-loading` enables `sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `openxr-structure-chains` enables `openxr-runtime-loading` (grounded, confidence 0.90).
- `packed-scenes` enables `scene-state` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `compressed-font-stream-adapters` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `font-streams` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `ktx-texture-container` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `lossy-macroblock-encoding` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `mp3-audio-resources` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `network-data-exchange` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `network-io` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `networking` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `ogg-packet-sequences` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `ogg-pages-and-packets` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `ogg-vorbis-audio-streams` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `state-recording` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `stream-networking` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `theora-video-streams` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `variant-text-serialization` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `websocket-peers` (synthesized lexical, confidence 0.25).
- `packet-queries` enables `zstandard-stream-codec` (synthesized lexical, confidence 0.25).
- `particle-systems` enables `renderer-storage` (synthesized lexical, confidence 0.25).
- `physics-2d-collision-pipeline` enables `physics-3d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `physics-2d-collision-pipeline` enables `two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `physics-3d-collision-pipeline` enables `convex-collision` (synthesized lexical, confidence 0.25).
- `physics-3d-collision-pipeline` enables `three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `physics-collision` enables `two-dimensional-content` (synthesized lexical, confidence 0.25).
- `physics-space-queries` enables `physics-queries` (synthesized lexical, confidence 0.25).
- `physics-spaces` enables `physics-queries` (grounded, confidence 0.90).
- `platform-display-adaptation` enables `platform-presentation` (synthesized lexical, confidence 0.25).
- `png-stream-transforms` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `png-stream-transforms` enables `png-image-codec` (synthesized lexical, confidence 0.25).
- `png-stream-transforms` enables `raster-image-input` (synthesized lexical, confidence 0.25).
- `prefix-code-decoding` enables `brotli-stream-decompression` (grounded, confidence 0.90).
- `prefix-code-decoding` enables `histograms-and-huffman-codes` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `filter-callbacks` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `ray-query` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `broad-phase` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `builder-heuristics` (grounded, confidence 0.90).
- `primitive-references` enables `bvh-construction` (grounded, confidence 0.90).
- `primitive-references` enables `geometry-resources` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `instancing` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `motion-blur` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `spatial-indexing` (synthesized lexical, confidence 0.25).
- `project-settings` enables `editor-and-project-settings` (synthesized lexical, confidence 0.25).
- `property-inspection` enables `editor-extensibility` (synthesized lexical, confidence 0.25).
- `property-inspection` enables `editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `random-generation` enables `native-support-algorithms` (synthesized lexical, confidence 0.25).
- `raster-font-rendering` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `raster-font-rendering` enables `raster-image-input` (synthesized lexical, confidence 0.25).
- `raster-image-input` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `raster-image-input` enables `jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `raster-image-input` enables `png-image-codec` (synthesized lexical, confidence 0.25).
- `raster-image-input` enables `png-stream-transforms` (synthesized lexical, confidence 0.25).
- `ray-primitive-intersection` enables `triangle-intersection-algorithms` (grounded, confidence 0.90).
- `ray-query` enables `bvh-traversal` (grounded, confidence 0.90).
- `ray-query` enables `collision-shapes` (synthesized lexical, confidence 0.25).
- `ray-query` enables `hit-results` (grounded, confidence 0.90).
- `ray-query` enables `motion-blur` (grounded, confidence 0.90).
- `ray-query` enables `motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `ray-query` enables `packet-queries` (grounded, confidence 0.90).
- `ray-query` enables `physics-queries` (synthesized lexical, confidence 0.25).
- `ray-query` enables `physics-server-queries` (synthesized lexical, confidence 0.25).
- `ray-query` enables `physics-space-queries` (synthesized lexical, confidence 0.25).
- `ray-query` enables `primitive-intersection` (grounded, confidence 0.90).
- `ray-query` enables `ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `ray-query` enables `simd-ray-packets` (synthesized lexical, confidence 0.25).
- `ray-query` enables `spatial-indexing` (synthesized lexical, confidence 0.25).
- `ray-query` enables `triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `reflection` enables `android-plugin-signals` (synthesized lexical, confidence 0.25).
- `reflection` enables `editor-and-project-settings` (synthesized lexical, confidence 0.25).
- `reflection` enables `editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `reflection` enables `filesystem-project-index` (synthesized lexical, confidence 0.25).
- `reflection` enables `godot-member-exposure` (synthesized lexical, confidence 0.25).
- `reflection` enables `manifold-mesh-data` (synthesized lexical, confidence 0.25).
- `reflection` enables `native-extensions` (grounded, confidence 0.90).
- `reflection` enables `png-stream-transforms` (synthesized lexical, confidence 0.25).
- `reflection` enables `profiling-interning` (synthesized lexical, confidence 0.25).
- `reflection` enables `project-settings` (synthesized lexical, confidence 0.25).
- `reflection` enables `reflection-metadata` (synthesized lexical, confidence 0.25).
- `reflection` enables `resource-previews` (synthesized lexical, confidence 0.25).
- `reflection` enables `script-registration-metadata` (synthesized lexical, confidence 0.25).
- `reflection` enables `scripting` (grounded, confidence 0.90).
- `reflection` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `reflection` enables `spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `reflection` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `reflection` enables `vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `reflection-metadata` enables `native-extensions` (synthesized lexical, confidence 0.25).
- `reflection-metadata` enables `reflection` (synthesized lexical, confidence 0.25).
- `regex-compile-match` enables `regex-jit-codegen` (grounded, confidence 0.90).
- `rendering-assets` enables `collada-import` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `geometry-resources` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `halfedge-topology` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `navigation-mesh-construction` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `rendering-resources` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `resource-specific-editors` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `spatial-indexing` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `three-dimensional-content` (synthesized lexical, confidence 0.25).
- `rendering-assets` enables `webrtc-multiplayer-mesh` (synthesized lexical, confidence 0.25).
- `rendering-backends` enables `android-render-input` (synthesized lexical, confidence 0.25).
- `rendering-device-resources` enables `renderer-storage` (grounded, confidence 0.90).
- `rendering-device-resources` enables `scene-data-and-buffers` (grounded, confidence 0.90).
- `rendering-device-resources` enables `shader-programs` (asserted, confidence 0.30).
- `rendering-device-resources` enables `textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `rendering-device-resources` enables `viewport-render-data` (asserted, confidence 0.30).
- `resource-assets` enables `audio-bus-processing` (grounded, confidence 0.90).
- `resource-assets` enables `device-runtime` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `gpu-command-encoding` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `gpu-resources-pipelines` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `physics-collision` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `renderer-storage` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `rendering-assets` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `rendering-resources` (grounded, confidence 0.90).
- `resource-assets` enables `resource-previews` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `scene-contexts` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `scene-state` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `skeletal-ik` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `skeleton-modification` (grounded, confidence 0.90).
- `resource-assets` enables `temporal-upscaling-dispatch` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `themes-and-style-boxes` (grounded, confidence 0.90).
- `resource-assets` enables `three-dimensional-content` (grounded, confidence 0.90).
- `resource-assets` enables `two-dimensional-content` (grounded, confidence 0.90).
- `resource-assets` enables `ui-theming` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `resource-formats` enables `resource-identifiers` (grounded, confidence 0.90).
- `resource-formats` enables `scene-state` (grounded, confidence 0.90).
- `resource-formats` enables `textures-and-placeholders` (grounded, confidence 0.90).
- `resource-formats` enables `tile-libraries` (grounded, confidence 0.90).
- `resource-identifiers` enables `generic-containers` (synthesized lexical, confidence 0.25).
- `resource-identifiers` enables `rendering-device-resources` (grounded, confidence 0.90).
- `resource-identifiers` enables `resource-dependency-resolution` (synthesized lexical, confidence 0.25).
- `resource-identifiers` enables `scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `resource-loading` enables `image-format-loading` (synthesized lexical, confidence 0.25).
- `resource-loading` enables `packed-resource-files` (grounded, confidence 0.90).
- `resource-loading` enables `resource-formats` (grounded, confidence 0.90).
- `resource-previews` enables `editor-extensibility` (synthesized lexical, confidence 0.25).
- `resource-previews` enables `localization-and-template-generation` (synthesized lexical, confidence 0.25).
- `resources` enables `allocation` (synthesized lexical, confidence 0.25).
- `resources` enables `audio-bus-processing` (synthesized lexical, confidence 0.25).
- `resources` enables `crypto-resources` (synthesized lexical, confidence 0.25).
- `resources` enables `editing-selection-history` (synthesized lexical, confidence 0.25).
- `resources` enables `editor-extensibility` (synthesized lexical, confidence 0.25).
- `resources` enables `editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `resources` enables `filesystem-project-index` (synthesized lexical, confidence 0.25).
- `resources` enables `gdextension-libraries` (synthesized lexical, confidence 0.25).
- `resources` enables `glyph-caching` (synthesized lexical, confidence 0.25).
- `resources` enables `gpu-command-encoding` (synthesized lexical, confidence 0.25).
- `resources` enables `gpu-memory-suballocation` (synthesized lexical, confidence 0.25).
- `resources` enables `image-codec-registration` (synthesized lexical, confidence 0.25).
- `resources` enables `input-events-actions` (synthesized lexical, confidence 0.25).
- `resources` enables `native-extensions` (grounded, confidence 0.90).
- `resources` enables `object-model` (synthesized lexical, confidence 0.25).
- `resources` enables `ogg-vorbis-audio-streams` (synthesized lexical, confidence 0.25).
- `resources` enables `packed-resource-files` (synthesized lexical, confidence 0.25).
- `resources` enables `packed-scenes` (grounded, confidence 0.90).
- `resources` enables `physics-collision` (grounded, confidence 0.90).
- `resources` enables `png-image-codec` (synthesized lexical, confidence 0.25).
- `resources` enables `project-settings` (synthesized lexical, confidence 0.25).
- `resources` enables `rendering-assets` (grounded, confidence 0.90).
- `resources` enables `rendering-resources` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-assets` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-bundle-data` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-dependency-resolution` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-formats` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-identifiers` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-importing` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-loading` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-previews` (synthesized lexical, confidence 0.25).
- `resources` enables `resource-specific-editors` (synthesized lexical, confidence 0.25).
- `resources` enables `scripting` (grounded, confidence 0.90).
- `resources` enables `shader-programs` (synthesized lexical, confidence 0.25).
- `resources` enables `skeletal-ik` (grounded, confidence 0.90).
- `resources` enables `skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `resources` enables `themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `resources` enables `theora-video-streams` (synthesized lexical, confidence 0.25).
- `resources` enables `three-dimensional-content` (synthesized lexical, confidence 0.25).
- `resources` enables `tile-libraries` (synthesized lexical, confidence 0.25).
- `resources` enables `two-dimensional-content` (synthesized lexical, confidence 0.25).
- `resources` enables `variant-text-serialization` (synthesized lexical, confidence 0.25).
- `rigid-body-motion` enables `constraints` (grounded, confidence 0.90).
- `rigid-body-motion` enables `vehicle-dynamics` (grounded, confidence 0.90).
- `run-management` enables `manifold-mesh-data` (synthesized lexical, confidence 0.25).
- `run-management` enables `skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `run-management` enables `target-platform-export` (synthesized lexical, confidence 0.25).
- `runtime-configuration` enables `engine-bootstrap` (grounded, confidence 0.90).
- `runtime-configuration` enables `input-action-configuration` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `android-gradle-assembly` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `cpu-specialized-dsp` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `motion-blur` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `project-catalog` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `ray-query` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `sdl-platform-backends` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `wraparound-network-time` (synthesized lexical, confidence 0.25).
- `runtime-metadata` enables `extensions` (grounded, confidence 0.90).
- `runtime-metadata` enables `native-extensions` (synthesized lexical, confidence 0.25).
- `runtime-metadata` enables `reflection` (synthesized lexical, confidence 0.25).
- `runtime-metadata` enables `reflection-metadata` (synthesized lexical, confidence 0.25).
- `runtime-metadata` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `runtime-metadata` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `runtime-metadata` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `runtime-metadata` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `scene-authoring` enables `scene-tree-and-signal-connections` (grounded, confidence 0.90).
- `scene-commit` enables `collada-import` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `editing-selection-history` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `editor-authoring-workspaces` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `editor-plugin-state` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `engine-bootstrap` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `explicit-drm-syncobj` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `gui-control-layout` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `instancing` (grounded, confidence 0.90).
- `scene-commit` enables `openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `openxr-render-models` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `primitive-references` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `project-catalog` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `property-introspection` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `resource-assets` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `resources` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `runtime-configuration` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-authoring` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-contexts` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-importing` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-state` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-tree` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-tree-execution` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `screen-and-environment-effects` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `tile-authoring` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `undo-redo-history` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `version-control-integration` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `collada-import` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `editing-selection-history` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `editor-authoring-workspaces` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `editor-plugin-state` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `engine-bootstrap` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `gui-control-layout` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `instancing` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `openxr-render-models` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `packed-scenes` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `primitive-references` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `project-catalog` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `property-introspection` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `resource-assets` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `resources` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `runtime-configuration` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-authoring` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-commit` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-importing` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-state` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-tree` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `scene-tree-execution` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `screen-and-environment-effects` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `tile-authoring` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `scene-contexts` enables `undo-redo-history` (synthesized lexical, confidence 0.25).
- `scene-data-and-buffers` enables `screen-and-environment-effects` (grounded, confidence 0.90).
- `scene-graph` enables `gui-control-layout` (grounded, confidence 0.90).
- `scene-graph` enables `navigation-agents` (grounded, confidence 0.90).
- `scene-graph` enables `particle-systems` (grounded, confidence 0.90).
- `scene-graph` enables `skeletal-modifiers-and-ik` (grounded, confidence 0.90).
- `scene-graph` enables `three-dimensional-physics` (grounded, confidence 0.90).
- `scene-graph` enables `tilemap-layer-data` (grounded, confidence 0.90).
- `scene-graph` enables `two-dimensional-physics` (grounded, confidence 0.90).
- `scene-importing` enables `collada-import` (grounded, confidence 0.90).
- `scene-importing` enables `post-import-processing` (grounded, confidence 0.90).
- `scene-replication-configuration` enables `replicated-property-synchronization` (grounded, confidence 0.90).
- `scene-replication-configuration` enables `replicated-spawning` (grounded, confidence 0.90).
- `scene-state` enables `packed-scenes` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `collada-import` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `expression-evaluation` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `input-actions` (grounded, confidence 0.90).
- `scene-tree` enables `networking` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `packed-scenes` (grounded, confidence 0.90).
- `scene-tree` enables `physics-collision` (grounded, confidence 0.90).
- `scene-tree` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `rendering-assets` (grounded, confidence 0.90).
- `scene-tree` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `scene-importing` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `scene-tree-execution` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `ui-composition` (grounded, confidence 0.90).
- `scene-tree` enables `visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `scene-tree-and-signal-connections` enables `scene-tree` (synthesized lexical, confidence 0.25).
- `scripting` enables `completion-contexts` (synthesized lexical, confidence 0.25).
- `scripting` enables `editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `scripting` enables `filesystem-project-index` (synthesized lexical, confidence 0.25).
- `scripting` enables `fixture-contracts` (synthesized lexical, confidence 0.25).
- `scripting` enables `gdscript-bytecode-runtime` (synthesized lexical, confidence 0.25).
- `scripting` enables `gdscript-editor-services` (synthesized lexical, confidence 0.25).
- `scripting` enables `gdscript-static-analysis` (synthesized lexical, confidence 0.25).
- `scripting` enables `managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `scripting` enables `script-hosting` (synthesized lexical, confidence 0.25).
- `scripting` enables `script-registration-metadata` (synthesized lexical, confidence 0.25).
- `scripting` enables `unicode-data` (synthesized lexical, confidence 0.25).
- `scripting` enables `visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `sdl-event-queue` enables `platform-display-adaptation` (synthesized lexical, confidence 0.25).
- `sdl-platform-backends` enables `hid-device-enumeration` (synthesized lexical, confidence 0.25).
- `sdl-platform-backends` enables `profiling-interning` (synthesized lexical, confidence 0.25).
- `sdl-platform-backends` enables `rendering-backends` (synthesized lexical, confidence 0.25).
- `sdl-platform-backends` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `sdl-platform-backends` enables `vector-font-export` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `editing-selection-history` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `manifold-mesh-data` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `property-inspection` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `property-introspection` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `reflection` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `reflection-metadata` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `runtime-metadata` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `scene-replication-configuration` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `scene-state` (synthesized lexical, confidence 0.25).
- `sdl-runtime-properties` enables `sdl-iostreams` (grounded, confidence 0.90).
- `sdl-runtime-properties` enables `text-shaping-plans` (synthesized lexical, confidence 0.25).
- `sdl-thread-abstractions` enables `deferred-execution` (synthesized lexical, confidence 0.25).
- `sdl-thread-abstractions` enables `native-platform-adapters` (synthesized lexical, confidence 0.25).
- `sdl-thread-abstractions` enables `task-scheduling` (synthesized lexical, confidence 0.25).
- `serialization` enables `state-recording` (grounded, confidence 0.90).
- `sfnt-tables` enables `font-table-validation` (grounded, confidence 0.90).
- `sfnt-tables` enables `variable-font-data` (grounded, confidence 0.90).
- `shader-editing-and-language-plugins` enables `platform-selective-shader-baking` (grounded, confidence 0.90).
- `shader-interface-analysis` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `shader-interface-metadata` enables `metal-presentation` (synthesized lexical, confidence 0.25).
- `shader-interface-metadata` enables `openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `shader-interface-metadata` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `shader-interface-metadata` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `shader-interface-metadata` enables `web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `shader-interface-metadata` enables `webxr-session-bridge` (synthesized lexical, confidence 0.25).
- `shader-interface-metadata` enables `zlib-stream-codec` (synthesized lexical, confidence 0.25).
- `shader-language-front-end` enables `spirv-generation` (grounded, confidence 0.90).
- `shader-programs` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `shader-programs` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `shader-programs` enables `spirv-generation` (synthesized lexical, confidence 0.25).
- `shader-programs` enables `spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `shader-programs` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `shader-reflection` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `shader-reflection` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `shader-reflection` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `shader-source-compilation` enables `shader-interface-analysis` (grounded, confidence 0.90).
- `signal-awaitability` enables `android-instrumented-tests` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `android-plugin-signals` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `godot-member-exposure` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `reflection` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `reflection-metadata` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `runtime-metadata` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `signed-distance-fields` enables `glyph-rasterization` (synthesized lexical, confidence 0.25).
- `simd-accelerated-codecs` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `simd-accelerated-codecs` enables `simd-ray-packets` (synthesized lexical, confidence 0.25).
- `simd-ray-packets` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `3d-gizmo-authoring` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `skeletal-ik` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `skeleton-modification` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `skeleton-modifiers` enables `skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
- `skeleton-modifiers` enables `skeleton-modification` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `allocation` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `builder-heuristics` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `bvh-construction` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `bvh-traversal` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `motion-blur` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `scene-commit` (synthesized lexical, confidence 0.25).
- `spatial-predictive-filters` enables `collision-filtering` (synthesized lexical, confidence 0.25).
- `spatial-predictive-filters` enables `filter-callbacks` (synthesized lexical, confidence 0.25).
- `spirv-generation` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `spirv-generation` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-generation` enables `spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `spirv-generation` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-intermediate-representation` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `spirv-intermediate-representation` enables `shader-reflection` (grounded, confidence 0.90).
- `spirv-intermediate-representation` enables `spirv-generation` (synthesized lexical, confidence 0.25).
- `spirv-intermediate-representation` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `spirv-generation` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-shader-reflection` enables `shader-interface-metadata` (grounded, confidence 0.90).
- `spirv-shader-reflection` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-shader-reflection` enables `spirv-generation` (synthesized lexical, confidence 0.25).
- `spirv-shader-reflection` enables `spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `string-names-paths` enables `localization` (grounded, confidence 0.90).
- `string-names-paths` enables `profiling-interning` (grounded, confidence 0.90).
- `string-names-paths` enables `scene-replication-configuration` (synthesized lexical, confidence 0.25).
- `string-names-paths` enables `variant-text-serialization` (grounded, confidence 0.90).
- `subdivision-surface-evaluation` enables `feature-adaptive-tessellation` (grounded, confidence 0.90).
- `subdivision-surface-evaluation` enables `geometry-families` (synthesized lexical, confidence 0.25).
- `subdivision-surface-evaluation` enables `half-edge-topology` (synthesized lexical, confidence 0.25).
- `synchronization-primitives` enables `concurrency` (synthesized lexical, confidence 0.25).
- `synchronization-primitives` enables `sdl-event-queue` (synthesized lexical, confidence 0.25).
- `synchronization-primitives` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `target-platform-export` enables `apple-embedded-packaging` (grounded, confidence 0.90).
- `target-platform-export` enables `platform-exports` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `concurrency` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `deferred-execution` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `device-runtime` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `encoder-configuration` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `synchronization-primitives` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `wayland-window-backend` (synthesized lexical, confidence 0.25).
- `task-scheduling` enables `worker-parallelism` (synthesized lexical, confidence 0.25).
- `text-and-localization` enables `localization` (synthesized lexical, confidence 0.25).
- `texture-block-codecs` enables `block-texture-encoding` (synthesized lexical, confidence 0.25).
- `texture-compression-pipeline` enables `basis-container-layout` (synthesized lexical, confidence 0.25).
- `texture-compression-pipeline` enables `basis-transcoding` (synthesized lexical, confidence 0.25).
- `texture-format-description` enables `ktx-texture-container` (grounded, confidence 0.90).
- `themes-and-style-boxes` enables `editor-theming` (synthesized lexical, confidence 0.25).
- `themes-and-style-boxes` enables `ui-theming` (synthesized lexical, confidence 0.25).
- `theora-block-video-codec` enables `theora-video-streams` (synthesized lexical, confidence 0.25).
- `theora-video-streams` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `tile-authoring` enables `tile-libraries` (synthesized lexical, confidence 0.25).
- `tile-libraries` enables `tilemap-layer-data` (synthesized lexical, confidence 0.25).
- `tile-libraries` enables `two-dimensional-content` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `crypto-resources` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `network-io` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `stream-networking` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `tls-crypto-backend` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `x509-certificate-processing` (grounded, confidence 0.90).
- `tls-crypto-backend` enables `crypto-resources` (synthesized lexical, confidence 0.25).
- `tls-crypto-backend` enables `network-io` (synthesized lexical, confidence 0.25).
- `tls-crypto-backend` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `tls-crypto-backend` enables `stream-networking` (synthesized lexical, confidence 0.25).
- `tls-crypto-backend` enables `tls-connection-state` (synthesized lexical, confidence 0.25).
- `tls-crypto-backend` enables `x509-certificate-processing` (synthesized lexical, confidence 0.25).
- `transform-interpolation` enables `metalfx-scaling` (synthesized lexical, confidence 0.25).
- `transform-quantization-rate-distortion` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `triangle-intersection-algorithms` enables `motion-blur-geometry` (grounded, confidence 0.90).
- `triangle-intersection-algorithms` enables `simd-ray-packets` (grounded, confidence 0.90).
- `triangle-intersection-algorithms` enables `subgrid-intersection` (grounded, confidence 0.90).
- `ui-composition` enables `curve-and-patch-bases` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `endian-safe-binary-io` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `gdscript-bytecode-runtime` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `gui-control-layout` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `gui-controls` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `platform-presentation` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `regex-jit-codegen` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `rendering-assets` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `ui-theming` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `version-control-integration` (synthesized lexical, confidence 0.25).
- `ui-composition` enables `websocket-frame-events` (synthesized lexical, confidence 0.25).
- `ui-theming` enables `editor-theming` (synthesized lexical, confidence 0.25).
- `ui-theming` enables `themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `undo-redo-history` enables `editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `unicode-data` enables `break-iteration` (grounded, confidence 0.90).
- `unicode-data` enables `character-encoding-conversion` (grounded, confidence 0.90).
- `unicode-data` enables `identifier-spoof-checking` (grounded, confidence 0.90).
- `unicode-data` enables `unicode-normalization` (grounded, confidence 0.90).
- `unicode-normalization` enables `unicode-data` (synthesized lexical, confidence 0.25).
- `upnp-device-control` enables `upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `upnp-device-discovery` enables `upnp-port-mapping` (grounded, confidence 0.90).
- `upnp-port-mapping` enables `upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `variable-font-subsetting` enables `variable-font-data` (synthesized lexical, confidence 0.25).
- `variant-containers` enables `break-iteration` (synthesized lexical, confidence 0.25).
- `variant-containers` enables `dynamic-values` (synthesized lexical, confidence 0.25).
- `variant-containers` enables `generic-containers` (grounded, confidence 0.90).
- `variant-containers` enables `packet-queries` (synthesized lexical, confidence 0.25).
- `variant-containers` enables `vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `vector-font-export` enables `generic-containers` (synthesized lexical, confidence 0.25).
- `vector-font-export` enables `multichannel-distance-fields` (synthesized lexical, confidence 0.25).
- `vector-font-export` enables `visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `vi-native-surface-creation` enables `vulkan-extensible-records` (synthesized lexical, confidence 0.25).
- `visual-shader-module-build` enables `visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `visual-shader-nodes` enables `visual-shader-vector-operations` (grounded, confidence 0.90).
- `vulkan-instance-setup` enables `openxr-runtime-loading` (synthesized lexical, confidence 0.25).
- `vulkan-swapchain-presentation` enables `canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `vulkan-swapchain-presentation` enables `metal-presentation` (synthesized lexical, confidence 0.25).
- `wayland-window-backend` enables `explicit-drm-syncobj` (synthesized lexical, confidence 0.25).
- `wayland-window-backend` enables `linuxbsd-desktop-integration` (synthesized lexical, confidence 0.25).
- `webp-image-io` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `webp-image-io` enables `webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `webp-riff-decoding` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `webp-riff-decoding` enables `webp-image-io` (synthesized lexical, confidence 0.25).
- `webrtc-data-channels` enables `webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `webrtc-peer-connections` enables `webrtc-data-channels` (grounded, confidence 0.90).
- `webrtc-peer-connections` enables `webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `websocket-frame-events` enables `websocket-multiplayer` (synthesized lexical, confidence 0.25).
- `websocket-multiplayer` enables `networking` (synthesized lexical, confidence 0.25).
- `websocket-peers` enables `websocket-multiplayer` (grounded, confidence 0.90).
- `xr-tracking` enables `pluggable-server-backends` (synthesized lexical, confidence 0.25).
- `xr-tracking` enables `skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
- `yuv-rgb-conversion` enables `input-picture-planes` (synthesized lexical, confidence 0.25).
- `zip-archive-io` enables `target-platform-export` (synthesized lexical, confidence 0.25).
- `zip-archive-io` enables `zip64-archive-io` (synthesized lexical, confidence 0.25).
- `zlib-stream-codec` enables `openexr-image-decoding` (synthesized lexical, confidence 0.25).

### implementation language
- `annotations-static-state-and-lifecycle` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cpp-runtime-casts` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cpp-static-generated-data` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `properties-and-accessors` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `typed-collections` (synthesized lexical, confidence 0.25).
- `c-abi-data-structures` enables `c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-abi-versioned-initialization` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-concurrent-state` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-dynamic-library-wrappers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-explicit-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-function-pointer-tables` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-macro-codecs` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-parser-state` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-platform-selection` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-pointers-arrays-and-strides` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-preprocessor-composition` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-stateful-streaming-api` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-structs-pointers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-classes` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-atomics` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-c-abi` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `c-abi-manual-lifetime` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `c-abi-versioned-initialization` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `c-abi-versioned-initialization` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `c-abi-versioned-initialization` enables `csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `c-aggregate-callback-modules` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `c-aggregate-callback-modules` enables `c-explicit-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `c-aggregate-callback-modules` enables `c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `c-aggregate-callback-modules` enables `c-function-pointer-tables` (synthesized lexical, confidence 0.25).
- `c-aggregate-callback-modules` enables `c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `c-aggregate-callback-modules` enables `cpp-basis-transcoding` (grounded, confidence 0.90).
- `c-compatible-header-guards` enables `c-abi-record-and-dispatch` (grounded, confidence 0.90).
- `c-function-pointer-apis` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `c-function-pointer-callbacks` enables `c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `c-function-pointer-callbacks` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `c-function-pointer-callbacks` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `c-function-pointer-interfaces` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `c-function-pointer-interfaces` enables `c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `c-function-pointer-interfaces` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `c-function-pointer-interfaces` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `c-function-pointer-tables` enables `c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `c-function-pointer-tables` enables `c-stateful-streaming-api` (synthesized lexical, confidence 0.25).
- `c-function-pointer-tables` enables `cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `c-function-pointer-tables` enables `cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `c-function-pointer-tables` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `c-function-pointer-tables` enables `javascript-web-runtime` (synthesized lexical, confidence 0.25).
- `c-integer-bitwise-packing` enables `cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `c-macro-codecs` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `c-macro-codecs` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `c-macro-codecs` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `c-platform-selection` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-abi-manual-lifetime` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-resource-lifecycles` (grounded, confidence 0.90).
- `c-pointers-arrays-and-strides` enables `c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-structs-pointers` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `c-preprocessor-composition` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `c-preprocessor-composition` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `c-preprocessor-composition` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-composition` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `c-preprocessor-composition` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-composition` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `c-preprocessor-configuration` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `c-preprocessor-platform-selection` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `c-preprocessor-platform-selection` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-platform-selection` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `c-preprocessor-platform-selection` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-platform-selection` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `c-preprocessor-portability` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `c-preprocessor-portability` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `c-preprocessor-portability` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-portability` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `c-preprocessor-portability` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `c-preprocessor-portability` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `c-resource-lifecycles` enables `c-structs-pointers` (synthesized lexical, confidence 0.25).
- `c-resource-lifecycles` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `c-resource-lifecycles` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `c-resource-lifecycles` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `c-stateful-struct-apis` enables `callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `c-stateful-struct-apis` enables `nodepaths-and-indexed-access` (synthesized lexical, confidence 0.25).
- `c-structs-pointers` enables `c-concurrent-state` (grounded, confidence 0.90).
- `c-structs-pointers` enables `c-explicit-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `c-structs-pointers` enables `c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `c-structs-pointers` enables `c-function-pointer-tables` (grounded, confidence 0.90).
- `c-structs-pointers` enables `c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-dynamic-library-wrappers` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-explicit-resource-lifecycles` (grounded, confidence 0.90).
- `c-structures-and-pointers` enables `c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-function-pointer-interfaces` (grounded, confidence 0.90).
- `c-structures-and-pointers` enables `c-function-pointer-tables` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-pointers-arrays-and-strides` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-stateful-streaming-api` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `cpp-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `cpp-object-representation-casts` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `cpp-runtime-casts` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `c-structures-and-pointers` enables `cxx-c-abi` (synthesized lexical, confidence 0.25).
- `callables-and-lambdas` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `callables-and-lambdas` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `callables-and-lambdas` enables `cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `callables-and-lambdas` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `callables-and-lambdas` enables `signals-and-await` (grounded, confidence 0.90).
- `classes-inheritance-and-lookup` enables `annotations-static-state-and-lifecycle` (grounded, confidence 0.90).
- `classes-inheritance-and-lookup` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `gdscript-declarations` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `properties-and-accessors` (grounded, confidence 0.90).
- `cplusplus-enumerated-export-state` enables `cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `cplusplus-export-callbacks` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `cplusplus-export-callbacks` enables `c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `cplusplus-export-callbacks` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-byte-exact-binary-parsing` (grounded, confidence 0.90).
- `cpp-abstraction-and-polymorphism` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-simd-intrinsics` (grounded, confidence 0.90).
- `cpp-abstraction-and-polymorphism` enables `cpp-templates-and-generic-containers` (grounded, confidence 0.90).
- `cpp-abstraction-and-polymorphism` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `gdscript-declarations` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `properties-and-accessors` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-classes` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-atomics` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-c-abi` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `cpp-basis-transcoding` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-classes` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-atomics` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-c-abi` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-numeric-value-operations` (grounded, confidence 0.90).
- `cpp-class-state-and-composition` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `gdscript-declarations` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-class-state-and-composition` enables `properties-and-accessors` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-templates` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-overloading` (grounded, confidence 0.90).
- `cpp-classes` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-virtual-dispatch` (grounded, confidence 0.90).
- `cpp-classes` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `gdscript-declarations` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `properties-and-accessors` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cpp-c-linkage-adapters` (grounded, confidence 0.90).
- `cpp-classes-and-inheritance` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cpp-plugin-specialization` (grounded, confidence 0.90).
- `cpp-classes-and-inheritance` enables `cpp-resource-lifetime` (grounded, confidence 0.90).
- `cpp-classes-and-inheritance` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cpp-visitor-traversal` (grounded, confidence 0.90).
- `cpp-classes-and-inheritance` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-inheritance` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-classes` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-atomics` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-c-abi` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-class-templates` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-copy-on-write-containers` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-godot-binding-macros` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-runtime-polymorphism` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-scoped-locks` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-const-borrowing` enables `cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `cpp-const-borrowing` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-engine-polymorphism` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-engine-polymorphism` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-engine-polymorphism` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-engine-polymorphism` enables `objective-cpp-apple-adapters` (grounded, confidence 0.90).
- `cpp-function-pointer-interfaces` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-runtime-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-virtual-dispatch` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `cpp-module-import` enables `c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `cpp-module-import` enables `c-platform-selection` (synthesized lexical, confidence 0.25).
- `cpp-module-import` enables `cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-module-import` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-classes` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-template-binding` (grounded, confidence 0.90).
- `cpp-native-classes` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-atomics` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-c-abi` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `cpp-native-classes` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `cpp-native-wrappers` enables `cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `cpp-native-wrappers` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-native-wrappers` enables `cpp-polymorphic-ownership` (grounded, confidence 0.90).
- `cpp-native-wrappers` enables `cpp-runtime-symbol-loading` (grounded, confidence 0.90).
- `cpp-native-wrappers` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `cpp-native-wrappers` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-runtime-casts` (grounded, confidence 0.90).
- `cpp-object-hierarchies` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-ownership-and-polymorphism` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-pluggable-allocation` enables `c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `cpp-polymorphic-backends` enables `cpp-virtual-dispatch` (synthesized lexical, confidence 0.25).
- `cpp-polymorphic-backends` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `cpp-static-abi-contracts` (grounded, confidence 0.90).
- `cpp-preprocessor-configuration` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `cpp-resource-and-polymorphism` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cpp-resource-lifetime` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-runtime-adapters` enables `cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-runtime-adapters` enables `cpp-jni-variant-marshalling` (grounded, confidence 0.90).
- `cpp-runtime-polymorphism` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-runtime-polymorphism` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cpp-scope-locking` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-scoped-locks` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-static-generated-data` enables `cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-templates` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-generic-containers` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-generic-containers` enables `cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-generic-containers` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-specialization` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-specialization` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-specialization` enables `cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-specialization` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-views` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-specialized-marshalling` (grounded, confidence 0.90).
- `cpp-templates-traits` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-variadic-binding` (grounded, confidence 0.90).
- `cpp-templates-traits` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-typed-api-records` enables `cpp-c-type-mapping` (grounded, confidence 0.90).
- `cpp-typed-api-records` enables `cpp-flag-stringification` (grounded, confidence 0.90).
- `csharp-attributes-reflection` enables `c-abi-manual-lifetime` (synthesized lexical, confidence 0.25).
- `csharp-attributes-reflection` enables `csharp-source-generation` (grounded, confidence 0.90).
- `csharp-attributes-reflection` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `csharp-partial-source-generation` enables `csharp-source-generation` (synthesized lexical, confidence 0.25).
- `csharp-partial-source-generation` enables `csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `csharp-source-generation` enables `csharp-unsafe-interop` (grounded, confidence 0.90).
- `csharp-source-generation` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `csharp-source-generation` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `csharp-unsafe-interop` enables `c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `csharp-unsafe-interop` enables `c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `csharp-unsafe-interop` enables `cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `cxx-atomics` enables `c-concurrent-state` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-structs-pointers` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `gdscript-declarations` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `properties-and-accessors` (synthesized lexical, confidence 0.25).
- `cxx-closures` enables `c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `cxx-closures` enables `callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `cxx-closures` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `cxx-closures` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `cxx-lambdas-standard-algorithms` enables `callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-class-registration` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cxx-raii-reference-ownership` (grounded, confidence 0.90).
- `cxx-object-model` enables `cxx-reflection-macros` (grounded, confidence 0.90).
- `cxx-object-model` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `c-platform-selection` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `c-preprocessor-portability` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `cxx-simd-alignment` (grounded, confidence 0.90).
- `cxx-preprocessor-configuration` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `cxx-reflection-macros` enables `cxx-stream-serialization` (grounded, confidence 0.90).
- `cxx-templates` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cxx-templates` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cxx-templates` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cxx-templates` enables `cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `cxx-templates` enables `cxx-wide-simd` (grounded, confidence 0.90).
- `cxx-templates` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `cxx-wide-simd` enables `c-platform-selection` (synthesized lexical, confidence 0.25).
- `cxx-wide-simd` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cxx-wide-simd` enables `cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `gdscript-declarations` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `gdscript-declarations` enables `gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `gdscript-declarations` enables `gdscript-query-objects` (synthesized lexical, confidence 0.25).
- `gdscript-declarations` enables `gdscript-signals-callables` (grounded, confidence 0.90).
- `gdscript-declarations` enables `types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `c-abi-manual-lifetime` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `cpp-native-classes` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `gdscript-declarations` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `signals-and-await` (synthesized lexical, confidence 0.25).
- `gdscript-typed-collections` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `godot-shader-language` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `godot-shader-language` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `godot-shader-language` enables `glsl-compute-shaders` (synthesized lexical, confidence 0.25).
- `iteration-protocols` enables `cpp-overloading` (synthesized lexical, confidence 0.25).
- `javascript-browser-ffi` enables `javascript-web-runtime` (synthesized lexical, confidence 0.25).
- `javascript-web-runtime` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `objective-cpp-apple-adapters` enables `objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `objective-cpp-ios-adapters` enables `objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `properties-and-accessors` enables `csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-scons-configuration` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-build-actions` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-scons-configuration` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-actions` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-scons-configuration` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-build-actions` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-scons-configuration` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-scons-build-hooks` enables `gdscript-query-objects` (synthesized lexical, confidence 0.25).
- `python-scons-build-hooks` enables `python-build-hooks` (synthesized lexical, confidence 0.25).
- `python-scons-build-hooks` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-scons-build-hooks` enables `python-scons-configuration` (synthesized lexical, confidence 0.25).
- `python-scons-build-hooks` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-actions` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-hooks` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-scons-module-configuration` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-scons-module-configuration` enables `python-scons-configuration` (synthesized lexical, confidence 0.25).
- `signals-and-await` enables `csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `signals-and-await` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `signals-and-await` enables `warnings-and-suppression` (synthesized lexical, confidence 0.25).
- `typed-collections` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `typed-collections` enables `gdscript-typed-collections` (synthesized lexical, confidence 0.25).
- `types-inference-and-conversions` enables `cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `types-inference-and-conversions` enables `cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `types-inference-and-conversions` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `types-inference-and-conversions` enables `typed-collections` (grounded, confidence 0.90).
- `types-inference-and-conversions` enables `warnings-and-suppression` (asserted, confidence 0.30).

### curriculum
- `domain/3d-gizmo-authoring` enables `domain/editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/buffer-storage` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/class-reference-generation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/godot-thirdparty-adaptation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/gpu-memory-suballocation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/histograms-and-huffman-codes` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/object-identity-lifetime` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/object-model` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/random-generation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/zip-archive-io` (synthesized lexical, confidence 0.25).
- `domain/alpha-plane-coding` enables `domain/encoder-configuration` (synthesized lexical, confidence 0.25).
- `domain/alpha-plane-coding` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `domain/alpha-plane-coding` enables `domain/input-picture-planes` (synthesized lexical, confidence 0.25).
- `domain/android-export-pipeline` enables `domain/android-gradle-assembly` (grounded, confidence 0.90).
- `domain/android-export-pipeline` enables `domain/apk-signing` (grounded, confidence 0.90).
- `domain/android-gradle-assembly` enables `language/cmake-native-source-index` (grounded, confidence 0.90).
- `domain/android-gradle-assembly` enables `language/groovy-gradle-build-logic` (grounded, confidence 0.90).
- `domain/android-jni-interop` enables `domain/android-plugin-signals` (grounded, confidence 0.90).
- `domain/android-jni-interop` enables `domain/android-storage-bridge` (grounded, confidence 0.90).
- `domain/android-platform-host` enables `domain/android-jni-interop` (grounded, confidence 0.90).
- `domain/android-platform-host` enables `domain/android-render-input` (grounded, confidence 0.90).
- `domain/android-platform-host` enables `language/java-android-host-api` (grounded, confidence 0.90).
- `domain/android-platform-host` enables `language/kotlin-runtime-coordination` (grounded, confidence 0.90).
- `domain/android-plugin-signals` enables `domain/android-instrumented-tests` (grounded, confidence 0.90).
- `domain/android-plugin-signals` enables `language/gdscript-plugin-integration` (grounded, confidence 0.90).
- `domain/apple-embedded-hosting` enables `domain/apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `domain/astc-block-coding` enables `domain/block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `domain/basis-transcoding` enables `domain/basis-container-layout` (grounded, confidence 0.90).
- `domain/basis-transcoding` enables `domain/block-texture-transcoding` (grounded, confidence 0.90).
- `domain/basis-transcoding` enables `domain/ktx2-container-transcoding` (grounded, confidence 0.90).
- `domain/block-texture-transcoding` enables `domain/astc-block-coding` (synthesized lexical, confidence 0.25).
- `domain/block-texture-transcoding` enables `domain/block-texture-encoding` (synthesized lexical, confidence 0.25).
- `domain/broad-phase` enables `domain/narrow-phase` (grounded, confidence 0.90).
- `domain/brotli-stream-decompression` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/brotli-stream-decompression` enables `domain/prefix-code-decoding` (synthesized lexical, confidence 0.25).
- `domain/buffer-storage` enables `domain/gltf-accessors` (synthesized lexical, confidence 0.25).
- `domain/buffer-storage` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/buffer-storage` enables `domain/web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `domain/buffer-storage` enables `domain/webp-image-io` (synthesized lexical, confidence 0.25).
- `domain/build-time-source-generation` enables `language/python-build-source-generation` (grounded, confidence 0.90).
- `domain/bvh-construction` enables `domain/allocation` (synthesized lexical, confidence 0.25).
- `domain/bvh-construction` enables `domain/builder-heuristics` (grounded, confidence 0.90).
- `domain/bvh-construction` enables `domain/bvh-traversal` (grounded, confidence 0.90).
- `domain/bvh-construction` enables `domain/motion-blur` (grounded, confidence 0.90).
- `domain/bvh-construction` enables `domain/physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `domain/bvh-construction` enables `domain/scene-commit` (grounded, confidence 0.90).
- `domain/bvh-construction` enables `language/cxx-closures` (grounded, confidence 0.90).
- `domain/bvh-traversal` enables `domain/canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/expression-evaluation` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/packed-scenes` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/ray-query` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/scene-importing` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/scene-tree` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/scene-tree-execution` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `domain/canvas-and-viewports` enables `domain/gui-controls` (grounded, confidence 0.90).
- `domain/canvas-and-viewports` enables `domain/platform-presentation` (grounded, confidence 0.90).
- `domain/catmull-clark-patch-construction` enables `domain/half-edge-topology` (synthesized lexical, confidence 0.25).
- `domain/cff-font-subsetting` enables `domain/font-driver-modules` (synthesized lexical, confidence 0.25).
- `domain/cff-font-subsetting` enables `domain/postscript-font-services` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/dynamic-variant` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/endian-safe-binary-io` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/managed-native-interop` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/project-upgrade` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `domain/character-encoding-conversion` enables `domain/yuv-rgb-conversion` (synthesized lexical, confidence 0.25).
- `domain/codec-simd-specialization` enables `domain/cpu-specialized-dsp` (synthesized lexical, confidence 0.25).
- `domain/codec-simd-specialization` enables `domain/simd-accelerated-codecs` (synthesized lexical, confidence 0.25).
- `domain/collision-filtering` enables `domain/broad-phase` (grounded, confidence 0.90).
- `domain/collision-shapes` enables `domain/collision-filtering` (grounded, confidence 0.90).
- `domain/collision-shapes` enables `domain/convex-collision` (grounded, confidence 0.90).
- `domain/collision-shapes` enables `domain/geometry-preprocessing` (grounded, confidence 0.90).
- `domain/collision-shapes` enables `domain/narrow-phase` (grounded, confidence 0.90).
- `domain/collision-shapes` enables `domain/physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `domain/collision-shapes` enables `domain/physics-3d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `domain/collision-shapes` enables `domain/physics-collision` (synthesized lexical, confidence 0.25).
- `domain/collision-shapes` enables `domain/physics-queries` (synthesized lexical, confidence 0.25).
- `domain/collision-shapes` enables `domain/physics-server-queries` (synthesized lexical, confidence 0.25).
- `domain/collision-shapes` enables `domain/physics-space-queries` (synthesized lexical, confidence 0.25).
- `domain/collision-shapes` enables `domain/textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `domain/color-font-paint` enables `domain/raster-font-rendering` (grounded, confidence 0.90).
- `domain/color-font-paint` enables `domain/vector-font-export` (grounded, confidence 0.90).
- `domain/compact-heightfield` enables `domain/navigation-regions` (grounded, confidence 0.90).
- `domain/compact-heightfield` enables `domain/navmesh-contours-polygons` (synthesized lexical, confidence 0.25).
- `domain/compact-heightfield` enables `domain/navmesh-heightfields` (synthesized lexical, confidence 0.25).
- `domain/completion-contexts` enables `domain/gdscript-editor-services` (synthesized lexical, confidence 0.25).
- `domain/completion-contexts` enables `domain/scene-contexts` (grounded, confidence 0.90).
- `domain/concurrency` enables `domain/device-runtime` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/encoder-configuration` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/sdl-event-queue` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/synchronization-primitives` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/task-scheduling` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/wayland-window-backend` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/worker-parallelism` (synthesized lexical, confidence 0.25).
- `domain/constraints` enables `domain/skeletal-ragdoll` (grounded, confidence 0.90).
- `domain/constraints` enables `domain/state-recording` (grounded, confidence 0.90).
- `domain/constraints` enables `domain/vehicle-dynamics` (grounded, confidence 0.90).
- `domain/contact-management` enables `domain/constraints` (grounded, confidence 0.90).
- `domain/convex-collision` enables `domain/narrow-phase` (grounded, confidence 0.90).
- `domain/convex-collision` enables `domain/physics-3d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `domain/cpu-specialized-dsp` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/cpu-specialized-dsp` enables `domain/simd-ray-packets` (synthesized lexical, confidence 0.25).
- `domain/crypto-resources` enables `domain/cryptography` (synthesized lexical, confidence 0.25).
- `domain/cryptography` enables `domain/crypto-resources` (synthesized lexical, confidence 0.25).
- `domain/cryptography` enables `domain/resource-formats` (synthesized lexical, confidence 0.25).
- `domain/cryptography` enables `domain/tls-crypto-backend` (synthesized lexical, confidence 0.25).
- `domain/curve-and-patch-bases` enables `domain/subdivision-surface-evaluation` (grounded, confidence 0.90).
- `domain/device-runtime` enables `domain/diagnostic-expectations` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/dynamic-variant` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/engine-bootstrap` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/gamepad-mapping` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/gpu-resources-pipelines` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/hid-device-enumeration` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/metalfx-scaling` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/object-model` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/openxr-dispatch` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/openxr-render-models` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/openxr-runtime-integration` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/openxr-runtime-loading` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/random-generation` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/reflection` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/rendering-backends` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/run-management` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/runtime-configuration` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/runtime-loop-time` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/runtime-metadata` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/scene-commit` (grounded, confidence 0.90).
- `domain/device-runtime` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/sdl-event-queue` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `domain/upnp-port-mapping` (synthesized lexical, confidence 0.25).
- `domain/device-runtime` enables `language/cxx-c-abi` (grounded, confidence 0.90).
- `domain/diagnostic-expectations` enables `domain/device-runtime` (synthesized lexical, confidence 0.25).
- `domain/diagnostic-expectations` enables `domain/image-quality-metrics` (synthesized lexical, confidence 0.25).
- `domain/diagnostic-expectations` enables `language/warnings-and-suppression` (asserted, confidence 0.30).
- `domain/dynamic-invocation-signals` enables `domain/class-reference-generation` (synthesized lexical, confidence 0.25).
- `domain/dynamic-invocation-signals` enables `domain/deferred-execution` (grounded, confidence 0.90).
- `domain/dynamic-invocation-signals` enables `domain/dynamic-values` (synthesized lexical, confidence 0.25).
- `domain/dynamic-invocation-signals` enables `domain/signal-awaitability` (synthesized lexical, confidence 0.25).
- `domain/dynamic-invocation-signals` enables `domain/undo-redo` (grounded, confidence 0.90).
- `domain/dynamic-values` enables `domain/android-jni-interop` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/break-iteration` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/dynamic-invocation-signals` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/dynamic-variant` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/expression-evaluation` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/generic-containers` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/managed-native-interop` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/network-data-exchange` (grounded, confidence 0.90).
- `domain/dynamic-values` enables `domain/project-settings` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/variant-containers` (synthesized lexical, confidence 0.25).
- `domain/dynamic-values` enables `domain/variant-text-serialization` (synthesized lexical, confidence 0.25).
- `domain/dynamic-variant` enables `domain/android-jni-interop` (synthesized lexical, confidence 0.25).
- `domain/dynamic-variant` enables `domain/dynamic-invocation-signals` (grounded, confidence 0.90).
- `domain/dynamic-variant` enables `domain/dynamic-values` (synthesized lexical, confidence 0.25).
- `domain/dynamic-variant` enables `domain/expression-evaluation` (grounded, confidence 0.90).
- `domain/dynamic-variant` enables `domain/managed-native-interop` (synthesized lexical, confidence 0.25).
- `domain/dynamic-variant` enables `domain/network-data-exchange` (synthesized lexical, confidence 0.25).
- `domain/dynamic-variant` enables `domain/project-settings` (synthesized lexical, confidence 0.25).
- `domain/dynamic-variant` enables `domain/variant-containers` (grounded, confidence 0.90).
- `domain/dynamic-variant` enables `domain/variant-text-serialization` (grounded, confidence 0.90).
- `domain/dynamic-variant` enables `language/cpp-tagged-storage` (grounded, confidence 0.90).
- `domain/dynamic-variant` enables `language/cpp-variadic-binding` (grounded, confidence 0.90).
- `domain/editor-and-project-settings` enables `domain/android-export-pipeline` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/audio-bus-processing` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/completion-contexts` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/engine-bootstrap` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/frame-timing` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/high-level-rpc` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/input-action-configuration` (grounded, confidence 0.90).
- `domain/editor-and-project-settings` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/openxr-action-configuration` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/psa-key-lifecycle` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/rendering-resources` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/runtime-configuration` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/tls-connection-state` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/vulkan-render-pass-configuration` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `domain/editor-and-project-settings` enables `domain/worker-parallelism` (synthesized lexical, confidence 0.25).
- `domain/editor-build-composition` enables `language/python-build-actions` (grounded, confidence 0.90).
- `domain/editor-extensibility` enables `domain/editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `domain/editor-extensibility` enables `domain/inspector-property-editors` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-extension` enables `domain/android-export-pipeline` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-extension` enables `domain/apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-extension` enables `domain/editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-extension` enables `domain/editor-plugin-state` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-extension` enables `domain/editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-extension` enables `domain/inspector-property-editors` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-extension` enables `domain/platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-lifecycle` enables `domain/android-export-pipeline` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-lifecycle` enables `domain/apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-lifecycle` enables `domain/editor-plugin-state` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-lifecycle` enables `domain/editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `domain/editor-plugin-lifecycle` enables `domain/inspector-property-editors` (grounded, confidence 0.90).
- `domain/editor-plugin-lifecycle` enables `domain/platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `domain/editor-scene-sessions` enables `domain/editing-selection-history` (grounded, confidence 0.90).
- `domain/editor-scene-sessions` enables `domain/editor-plugin-state` (grounded, confidence 0.90).
- `domain/editor-scene-sessions` enables `domain/undo-redo-history` (grounded, confidence 0.90).
- `domain/editor-theming` enables `domain/themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `domain/editor-theming` enables `domain/ui-theming` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/android-export-pipeline` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/apple-embedded-packaging` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/audio-bus-processing` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/completion-contexts` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/editor-and-project-settings` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/engine-bootstrap` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/frame-timing` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/high-level-rpc` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/input-action-configuration` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/lossy-macroblock-encoding` (grounded, confidence 0.90).
- `domain/encoder-configuration` enables `domain/openxr-action-configuration` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/psa-key-lifecycle` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/rendering-resources` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/runtime-configuration` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/tls-connection-state` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/vulkan-render-pass-configuration` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `domain/encoder-configuration` enables `domain/worker-parallelism` (grounded, confidence 0.90).
- `domain/enet-host-peer-transport` enables `domain/enet-wire-commands` (grounded, confidence 0.90).
- `domain/enet-host-peer-transport` enables `domain/godot-enet-socket-adaptation` (synthesized lexical, confidence 0.25).
- `domain/enet-host-peer-transport` enables `domain/wraparound-network-time` (synthesized lexical, confidence 0.25).
- `domain/enet-transport-and-multiplayer` enables `domain/enet-host-peer-transport` (synthesized lexical, confidence 0.25).
- `domain/enet-transport-and-multiplayer` enables `domain/enet-wire-commands` (synthesized lexical, confidence 0.25).
- `domain/enet-transport-and-multiplayer` enables `domain/godot-enet-socket-adaptation` (synthesized lexical, confidence 0.25).
- `domain/enet-transport-and-multiplayer` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/enet-transport-and-multiplayer` enables `domain/wraparound-network-time` (synthesized lexical, confidence 0.25).
- `domain/enet-wire-commands` enables `domain/godot-enet-socket-adaptation` (grounded, confidence 0.90).
- `domain/engine-bootstrap` enables `domain/android-platform-host` (grounded, confidence 0.90).
- `domain/engine-bootstrap` enables `domain/frame-timing` (grounded, confidence 0.90).
- `domain/engine-bootstrap` enables `domain/performance-monitors` (grounded, confidence 0.90).
- `domain/entropy-bitstreams` enables `domain/endian-safe-binary-io` (synthesized lexical, confidence 0.25).
- `domain/export-orchestration` enables `domain/target-platform-export` (grounded, confidence 0.90).
- `domain/export-presets` enables `domain/export-orchestration` (grounded, confidence 0.90).
- `domain/extension-api-compatibility` enables `domain/gdextension-libraries` (synthesized lexical, confidence 0.25).
- `domain/extension-api-compatibility` enables `domain/native-extensions` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/extension-api-compatibility` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/extension-structure-chains` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/gdextension-libraries` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/native-extensions` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/openxr-render-models` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/pluggable-server-backends` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/resource-formats` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/script-hosting` (synthesized lexical, confidence 0.25).
- `domain/extensions` enables `domain/vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `domain/fbx-gltf-scene-import` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `domain/fbx-gltf-scene-import` enables `domain/gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `domain/filesystem-project-index` enables `domain/export-presets` (grounded, confidence 0.90).
- `domain/filesystem-project-index` enables `domain/resource-dependency-resolution` (grounded, confidence 0.90).
- `domain/filter-callbacks` enables `domain/font-streams` (synthesized lexical, confidence 0.25).
- `domain/filter-callbacks` enables `domain/web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `domain/filter-callbacks` enables `domain/websocket-frame-events` (synthesized lexical, confidence 0.25).
- `domain/fixture-contracts` enables `domain/completion-contexts` (grounded, confidence 0.90).
- `domain/fixture-contracts` enables `domain/diagnostic-expectations` (grounded, confidence 0.90).
- `domain/fixture-contracts` enables `domain/lsp-semantic-fixtures` (grounded, confidence 0.90).
- `domain/font-driver-modules` enables `domain/glyph-caching` (grounded, confidence 0.90).
- `domain/font-driver-modules` enables `domain/glyph-rasterization` (grounded, confidence 0.90).
- `domain/font-driver-modules` enables `domain/postscript-font-services` (grounded, confidence 0.90).
- `domain/font-streams` enables `domain/compressed-font-stream-adapters` (grounded, confidence 0.90).
- `domain/font-streams` enables `domain/sfnt-tables` (grounded, confidence 0.90).
- `domain/font-subsetting` enables `domain/cff-font-subsetting` (grounded, confidence 0.90).
- `domain/font-subsetting` enables `domain/opentype-table-routing` (grounded, confidence 0.90).
- `domain/font-subsetting` enables `domain/variable-font-subsetting` (grounded, confidence 0.90).
- `domain/font-table-access` enables `domain/graphite-rule-execution` (grounded, confidence 0.90).
- `domain/font-table-access` enables `domain/graphite-shaping` (grounded, confidence 0.90).
- `domain/font-table-access` enables `domain/sfnt-tables` (synthesized lexical, confidence 0.25).
- `domain/frame-timing` enables `language/cpp-frame-scheduling` (grounded, confidence 0.90).
- `domain/freetype-module-composition` enables `domain/font-driver-modules` (synthesized lexical, confidence 0.25).
- `domain/gdextension-libraries` enables `domain/native-extensions` (synthesized lexical, confidence 0.25).
- `domain/gdscript-bytecode-runtime` enables `domain/graphite-rule-execution` (synthesized lexical, confidence 0.25).
- `domain/gdscript-editor-services` enables `domain/shader-editing-and-language-plugins` (synthesized lexical, confidence 0.25).
- `domain/gdscript-static-analysis` enables `domain/gdscript-bytecode-runtime` (grounded, confidence 0.90).
- `domain/geometry-families` enables `domain/collision-shapes` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/geometry-preprocessing` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/geometry-resources` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/halfedge-topology` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/mesh-provenance` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/simd-ray-packets` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/spatial-indexing` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/subdivision-surface-evaluation` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/subgrid-intersection` (synthesized lexical, confidence 0.25).
- `domain/geometry-families` enables `domain/triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/buffer-storage` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/bvh-traversal` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/catmull-clark-patch-construction` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/filter-callbacks` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/geometry-families` (grounded, confidence 0.90).
- `domain/geometry-resources` enables `domain/geometry-preprocessing` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/hit-results` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/instancing` (grounded, confidence 0.90).
- `domain/geometry-resources` enables `domain/mesh-geometry-algorithms` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/motion-blur` (grounded, confidence 0.90).
- `domain/geometry-resources` enables `domain/navigation-mesh-construction` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/primitive-intersection` (grounded, confidence 0.90).
- `domain/geometry-resources` enables `domain/primitive-references` (grounded, confidence 0.90).
- `domain/geometry-resources` enables `domain/ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/rendering-assets` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/scene-commit` (grounded, confidence 0.90).
- `domain/geometry-resources` enables `domain/spatial-indexing` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/transform-interpolation` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `domain/geometry-resources` enables `domain/two-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/geometry-transforms` enables `domain/mesh-geometry-algorithms` (grounded, confidence 0.90).
- `domain/geometry-transforms` enables `domain/spatial-indexing` (grounded, confidence 0.90).
- `domain/geometry-transforms` enables `domain/transform-interpolation` (grounded, confidence 0.90).
- `domain/gltf-materials-textures` enables `domain/basis-container-layout` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/basis-transcoding` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/ktx-texture-container` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/ktx2-container-transcoding` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/raster-image-input` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/rendering-resources` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/resource-importing` (synthesized lexical, confidence 0.25).
- `domain/gltf-materials-textures` enables `domain/textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/expression-evaluation` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/packed-scenes` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/scene-importing` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/scene-tree` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/scene-tree-execution` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `domain/gltf-node-hierarchy` enables `domain/visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `domain/glyph-rasterization` enables `domain/signed-distance-fields` (grounded, confidence 0.90).
- `domain/glyph-rasterization` enables `domain/vulkan-pipeline-configuration` (synthesized lexical, confidence 0.25).
- `domain/godot-enet-socket-adaptation` enables `domain/network-io` (synthesized lexical, confidence 0.25).
- `domain/godot-enet-socket-adaptation` enables `domain/tls-crypto-backend` (synthesized lexical, confidence 0.25).
- `domain/gpu-command-encoding` enables `domain/lossless-transform-pipeline` (synthesized lexical, confidence 0.25).
- `domain/gpu-command-encoding` enables `domain/lossy-macroblock-encoding` (synthesized lexical, confidence 0.25).
- `domain/gpu-command-encoding` enables `domain/metal-presentation` (grounded, confidence 0.90).
- `domain/gpu-resources-pipelines` enables `domain/android-gradle-assembly` (synthesized lexical, confidence 0.25).
- `domain/gpu-resources-pipelines` enables `domain/apk-signing` (synthesized lexical, confidence 0.25).
- `domain/gpu-resources-pipelines` enables `domain/cff-font-subsetting` (synthesized lexical, confidence 0.25).
- `domain/gpu-resources-pipelines` enables `domain/gpu-command-encoding` (grounded, confidence 0.90).
- `domain/gpu-resources-pipelines` enables `domain/jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `domain/gpu-resources-pipelines` enables `domain/metalfx-scaling` (grounded, confidence 0.90).
- `domain/gpu-resources-pipelines` enables `domain/png-stream-transforms` (synthesized lexical, confidence 0.25).
- `domain/gpu-resources-pipelines` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/gpu-resources-pipelines` enables `domain/variable-font-subsetting` (synthesized lexical, confidence 0.25).
- `domain/gpu-resources-pipelines` enables `domain/webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `domain/graphite-rule-execution` enables `domain/graphite-shaping` (grounded, confidence 0.90).
- `domain/gui-control-layout` enables `domain/curve-and-patch-bases` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/gui-controls` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/platform-presentation` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/regex-jit-codegen` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/rendering-assets` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/ui-composition` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/ui-theming` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/version-control-integration` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/websocket-frame-events` (synthesized lexical, confidence 0.25).
- `domain/gui-controls` enables `domain/gui-control-layout` (synthesized lexical, confidence 0.25).
- `domain/gui-controls` enables `domain/text-interface` (grounded, confidence 0.90).
- `domain/half-edge-topology` enables `domain/catmull-clark-patch-construction` (grounded, confidence 0.90).
- `domain/halfedge-topology` enables `domain/half-edge-topology` (synthesized lexical, confidence 0.25).
- `domain/hid-device-enumeration` enables `domain/gamepad-mapping` (grounded, confidence 0.90).
- `domain/high-level-rpc` enables `domain/godot-member-exposure` (synthesized lexical, confidence 0.25).
- `domain/high-level-rpc` enables `domain/network-data-exchange` (synthesized lexical, confidence 0.25).
- `domain/histograms-and-huffman-codes` enables `domain/brotli-stream-decompression` (synthesized lexical, confidence 0.25).
- `domain/histograms-and-huffman-codes` enables `domain/prefix-code-decoding` (synthesized lexical, confidence 0.25).
- `domain/hit-results` enables `domain/filter-callbacks` (grounded, confidence 0.90).
- `domain/hit-results` enables `domain/packet-queries` (grounded, confidence 0.90).
- `domain/hit-results` enables `domain/primitive-intersection` (grounded, confidence 0.90).
- `domain/hit-results` enables `domain/ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `domain/hit-results` enables `domain/triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `domain/image-codec-registration` enables `domain/astc-block-coding` (synthesized lexical, confidence 0.25).
- `domain/image-codec-registration` enables `domain/block-texture-transcoding` (synthesized lexical, confidence 0.25).
- `domain/image-decode-pipeline` enables `domain/brotli-stream-decompression` (synthesized lexical, confidence 0.25).
- `domain/image-decode-pipeline` enables `domain/jpeg-decompression-transcoding` (grounded, confidence 0.90).
- `domain/image-decode-pipeline` enables `domain/png-stream-transforms` (grounded, confidence 0.90).
- `domain/image-decode-pipeline` enables `domain/webp-riff-decoding` (grounded, confidence 0.90).
- `domain/image-format-loading` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/image-format-loading` enables `domain/jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `domain/image-format-loading` enables `domain/ktx-texture-container` (synthesized lexical, confidence 0.25).
- `domain/image-format-loading` enables `domain/raster-image-input` (synthesized lexical, confidence 0.25).
- `domain/image-quality-metrics` enables `domain/sfnt-tables` (synthesized lexical, confidence 0.25).
- `domain/input-actions` enables `domain/input-events-actions` (synthesized lexical, confidence 0.25).
- `domain/input-actions` enables `domain/platform-display-adaptation` (synthesized lexical, confidence 0.25).
- `domain/input-events-actions` enables `domain/input-actions` (synthesized lexical, confidence 0.25).
- `domain/input-events-actions` enables `domain/platform-display-adaptation` (grounded, confidence 0.90).
- `domain/input-midi` enables `domain/apple-embedded-hosting` (synthesized lexical, confidence 0.25).
- `domain/input-midi` enables `domain/browser-runtime-adapters` (synthesized lexical, confidence 0.25).
- `domain/input-midi` enables `domain/input-events-actions` (synthesized lexical, confidence 0.25).
- `domain/input-midi` enables `domain/midi-input-adapters` (synthesized lexical, confidence 0.25).
- `domain/input-picture-planes` enables `domain/alpha-plane-coding` (grounded, confidence 0.90).
- `domain/input-picture-planes` enables `domain/encoder-configuration` (grounded, confidence 0.90).
- `domain/input-picture-planes` enables `domain/geometry-transforms` (synthesized lexical, confidence 0.25).
- `domain/input-picture-planes` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/input-picture-planes` enables `domain/image-quality-metrics` (grounded, confidence 0.90).
- `domain/input-picture-planes` enables `domain/lossless-transform-pipeline` (grounded, confidence 0.90).
- `domain/input-picture-planes` enables `domain/lossy-macroblock-encoding` (grounded, confidence 0.90).
- `domain/input-picture-planes` enables `domain/spatial-predictive-filters` (grounded, confidence 0.90).
- `domain/input-picture-planes` enables `domain/yuv-rgb-conversion` (grounded, confidence 0.90).
- `domain/instancing` enables `domain/geometry-resources` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/hit-results` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/openxr-runtime-loading` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/rendering-assets` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/rigid-body-motion` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `domain/ios-embedded-adapter` enables `language/objective-cpp-ios-adapters` (grounded, confidence 0.90).
- `domain/jpeg-decompression-transcoding` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/jpeg-decompression-transcoding` enables `domain/raster-image-input` (synthesized lexical, confidence 0.25).
- `domain/ktx-texture-container` enables `domain/image-format-loading` (synthesized lexical, confidence 0.25).
- `domain/ktx2-container-transcoding` enables `domain/ktx-texture-container` (synthesized lexical, confidence 0.25).
- `domain/ktx2-container-transcoding` enables `domain/texture-compression-pipeline` (synthesized lexical, confidence 0.25).
- `domain/linuxbsd-desktop-integration` enables `domain/wayland-window-backend` (grounded, confidence 0.90).
- `domain/linuxbsd-desktop-integration` enables `language/c-dynamic-library-wrappers` (grounded, confidence 0.90).
- `domain/localization` enables `domain/text-and-localization` (synthesized lexical, confidence 0.25).
- `domain/lossless-transform-pipeline` enables `domain/entropy-bitstreams` (synthesized lexical, confidence 0.25).
- `domain/lossless-transform-pipeline` enables `domain/webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `domain/lossy-macroblock-encoding` enables `domain/transform-quantization-rate-distortion` (grounded, confidence 0.90).
- `domain/lsp-semantic-fixtures` enables `domain/gdscript-language-server` (synthesized lexical, confidence 0.25).
- `domain/manifold-mesh-data` enables `domain/halfedge-topology` (grounded, confidence 0.90).
- `domain/metal-cpp-object-bridge` enables `domain/gpu-resources-pipelines` (grounded, confidence 0.90).
- `domain/metal-cpp-object-bridge` enables `domain/metal-presentation` (synthesized lexical, confidence 0.25).
- `domain/metal-cpp-object-bridge` enables `domain/metalfx-scaling` (synthesized lexical, confidence 0.25).
- `domain/metal-cpp-object-bridge` enables `domain/platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `domain/metal-cpp-object-bridge` enables `domain/rendering-backends` (synthesized lexical, confidence 0.25).
- `domain/module-build-configuration` enables `domain/editor-build-composition` (synthesized lexical, confidence 0.25).
- `domain/module-build-configuration` enables `domain/platform-selective-shader-baking` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/android-gradle-assembly` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/builder-heuristics` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/constraints` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/cpu-specialized-dsp` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/physics-queries` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/physics-server-queries` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/physics-space-queries` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/project-catalog` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/ray-query` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/rigid-body-motion` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/runtime-loop-time` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/sdl-platform-backends` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/wraparound-network-time` (synthesized lexical, confidence 0.25).
- `domain/motion-blur-geometry` enables `domain/motion-blur` (synthesized lexical, confidence 0.25).
- `domain/narrow-phase` enables `domain/contact-management` (grounded, confidence 0.90).
- `domain/narrow-phase` enables `domain/vehicle-dynamics` (grounded, confidence 0.90).
- `domain/native-extensions` enables `domain/gdextension-libraries` (synthesized lexical, confidence 0.25).
- `domain/native-platform-adapters` enables `domain/apple-embedded-hosting` (grounded, confidence 0.90).
- `domain/native-platform-adapters` enables `domain/audio-backend-adapters` (grounded, confidence 0.90).
- `domain/native-platform-adapters` enables `domain/midi-input-adapters` (grounded, confidence 0.90).
- `domain/navigation-agents` enables `domain/compact-heightfield` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/editing-selection-history` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/half-edge-topology` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/navigation` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/navigation-map-composition` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/navigation-mesh-construction` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/navigation-queries` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/navigation-regions` (synthesized lexical, confidence 0.25).
- `domain/navigation-agents` enables `domain/tilemap-layer-data` (synthesized lexical, confidence 0.25).
- `domain/navigation-map-composition` enables `domain/navigation-avoidance` (grounded, confidence 0.90).
- `domain/navigation-map-composition` enables `domain/navigation-path-queries` (grounded, confidence 0.90).
- `domain/navigation-regions` enables `domain/navigation` (synthesized lexical, confidence 0.25).
- `domain/navigation-regions` enables `domain/navigation-agents` (synthesized lexical, confidence 0.25).
- `domain/navigation-regions` enables `domain/navigation-map-composition` (synthesized lexical, confidence 0.25).
- `domain/navmesh-heightfields` enables `domain/navmesh-contours-polygons` (grounded, confidence 0.90).
- `domain/network-data-exchange` enables `domain/webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `domain/network-io` enables `domain/webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/dynamic-invocation-signals` (grounded, confidence 0.90).
- `domain/object-identity-lifetime` enables `domain/navigation-queries` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/object-model` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/physics-server-queries` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/physics-space-queries` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/reflection-metadata` (grounded, confidence 0.90).
- `domain/object-identity-lifetime` enables `domain/regular-expression-results` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/resources` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/script-hosting` (grounded, confidence 0.90).
- `domain/object-identity-lifetime` enables `domain/upnp-device-control` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/collision-filtering` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/deferred-execution` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/dynamic-invocation-signals` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/dynamic-values` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/metal-cpp-object-bridge` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/navigation-queries` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/object-identity-lifetime` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/physics-server-queries` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/physics-space-queries` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/property-inspection` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/reflection` (grounded, confidence 0.90).
- `domain/object-model` enables `domain/reflection-metadata` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/regular-expression-results` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/resources` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/runtime-loop-time` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/scene-tree` (grounded, confidence 0.90).
- `domain/object-model` enables `domain/script-hosting` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/scripting` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/serialization` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/text-shaping-plans` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/undo-redo` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/upnp-device-control` (synthesized lexical, confidence 0.25).
- `domain/object-model` enables `domain/web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `domain/ogg-packet-sequences` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/ogg-packet-sequences` enables `domain/vorbis-block-synthesis` (synthesized lexical, confidence 0.25).
- `domain/ogg-pages-and-packets` enables `domain/theora-block-video-codec` (grounded, confidence 0.90).
- `domain/ogg-pages-and-packets` enables `domain/vorbis-block-synthesis` (grounded, confidence 0.90).
- `domain/openxr-dispatch` enables `domain/character-encoding-conversion` (synthesized lexical, confidence 0.25).
- `domain/openxr-dispatch` enables `domain/collision-shapes` (synthesized lexical, confidence 0.25).
- `domain/openxr-dispatch` enables `domain/ktx-texture-container` (synthesized lexical, confidence 0.25).
- `domain/openxr-dispatch` enables `domain/temporal-upscaling-dispatch` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-integration` enables `domain/openxr-binding-modifiers` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-integration` enables `domain/openxr-composition-layers` (grounded, confidence 0.90).
- `domain/openxr-runtime-integration` enables `domain/openxr-extension-wrappers` (grounded, confidence 0.90).
- `domain/openxr-runtime-integration` enables `domain/openxr-render-models` (grounded, confidence 0.90).
- `domain/openxr-runtime-integration` enables `domain/openxr-structure-chains` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/diagnostic-expectations` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/dynamic-variant` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/engine-bootstrap` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/object-model` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/openxr-dispatch` (grounded, confidence 0.90).
- `domain/openxr-runtime-loading` enables `domain/openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/openxr-render-models` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/openxr-runtime-integration` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/random-generation` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/reflection` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/runtime-configuration` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/runtime-loop-time` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/runtime-metadata` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/sdl-event-queue` (synthesized lexical, confidence 0.25).
- `domain/openxr-runtime-loading` enables `domain/sdl-runtime-properties` (synthesized lexical, confidence 0.25).
- `domain/openxr-structure-chains` enables `domain/openxr-runtime-loading` (grounded, confidence 0.90).
- `domain/packed-scenes` enables `domain/scene-state` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/compressed-font-stream-adapters` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/font-streams` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/ktx-texture-container` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/lossy-macroblock-encoding` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/mp3-audio-resources` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/network-data-exchange` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/network-io` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/ogg-packet-sequences` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/ogg-pages-and-packets` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/ogg-vorbis-audio-streams` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/state-recording` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/stream-networking` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/theora-video-streams` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/variant-text-serialization` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/websocket-peers` (synthesized lexical, confidence 0.25).
- `domain/packet-queries` enables `domain/zstandard-stream-codec` (synthesized lexical, confidence 0.25).
- `domain/particle-systems` enables `domain/renderer-storage` (synthesized lexical, confidence 0.25).
- `domain/physics-2d-collision-pipeline` enables `domain/physics-3d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `domain/physics-2d-collision-pipeline` enables `domain/two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/physics-3d-collision-pipeline` enables `domain/convex-collision` (synthesized lexical, confidence 0.25).
- `domain/physics-3d-collision-pipeline` enables `domain/three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/physics-collision` enables `domain/two-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/physics-queries` enables `language/gdscript-query-objects` (grounded, confidence 0.90).
- `domain/physics-space-queries` enables `domain/physics-queries` (synthesized lexical, confidence 0.25).
- `domain/physics-spaces` enables `domain/physics-queries` (grounded, confidence 0.90).
- `domain/platform-display-adaptation` enables `domain/platform-presentation` (synthesized lexical, confidence 0.25).
- `domain/png-stream-transforms` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/png-stream-transforms` enables `domain/png-image-codec` (synthesized lexical, confidence 0.25).
- `domain/png-stream-transforms` enables `domain/raster-image-input` (synthesized lexical, confidence 0.25).
- `domain/prefix-code-decoding` enables `domain/brotli-stream-decompression` (grounded, confidence 0.90).
- `domain/prefix-code-decoding` enables `domain/histograms-and-huffman-codes` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/filter-callbacks` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/ray-query` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/broad-phase` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/builder-heuristics` (grounded, confidence 0.90).
- `domain/primitive-references` enables `domain/bvh-construction` (grounded, confidence 0.90).
- `domain/primitive-references` enables `domain/geometry-resources` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/instancing` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/motion-blur` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/spatial-indexing` (synthesized lexical, confidence 0.25).
- `domain/project-settings` enables `domain/editor-and-project-settings` (synthesized lexical, confidence 0.25).
- `domain/property-inspection` enables `domain/editor-extensibility` (synthesized lexical, confidence 0.25).
- `domain/property-inspection` enables `domain/editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `domain/random-generation` enables `domain/native-support-algorithms` (synthesized lexical, confidence 0.25).
- `domain/raster-font-rendering` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/raster-font-rendering` enables `domain/raster-image-input` (synthesized lexical, confidence 0.25).
- `domain/raster-image-input` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/raster-image-input` enables `domain/jpeg-decompression-transcoding` (synthesized lexical, confidence 0.25).
- `domain/raster-image-input` enables `domain/png-image-codec` (synthesized lexical, confidence 0.25).
- `domain/raster-image-input` enables `domain/png-stream-transforms` (synthesized lexical, confidence 0.25).
- `domain/ray-primitive-intersection` enables `domain/triangle-intersection-algorithms` (grounded, confidence 0.90).
- `domain/ray-query` enables `domain/bvh-traversal` (grounded, confidence 0.90).
- `domain/ray-query` enables `domain/collision-shapes` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/hit-results` (grounded, confidence 0.90).
- `domain/ray-query` enables `domain/motion-blur` (grounded, confidence 0.90).
- `domain/ray-query` enables `domain/motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/packet-queries` (grounded, confidence 0.90).
- `domain/ray-query` enables `domain/physics-queries` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/physics-server-queries` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/physics-space-queries` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/primitive-intersection` (grounded, confidence 0.90).
- `domain/ray-query` enables `domain/ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/simd-ray-packets` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/spatial-indexing` (synthesized lexical, confidence 0.25).
- `domain/ray-query` enables `domain/triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/android-plugin-signals` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/editor-and-project-settings` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/filesystem-project-index` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/godot-member-exposure` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/manifold-mesh-data` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/native-extensions` (grounded, confidence 0.90).
- `domain/reflection` enables `domain/png-stream-transforms` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/profiling-interning` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/project-settings` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/reflection-metadata` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/resource-previews` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/script-registration-metadata` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/scripting` (grounded, confidence 0.90).
- `domain/reflection` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `language/cpp-template-binding` (grounded, confidence 0.90).
- `domain/reflection-metadata` enables `domain/native-extensions` (synthesized lexical, confidence 0.25).
- `domain/reflection-metadata` enables `domain/reflection` (synthesized lexical, confidence 0.25).
- `domain/reflection-metadata` enables `language/cpp-specialized-marshalling` (grounded, confidence 0.90).
- `domain/regex-compile-match` enables `domain/regex-jit-codegen` (grounded, confidence 0.90).
- `domain/rendering-assets` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/geometry-resources` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/halfedge-topology` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/navigation-mesh-construction` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/rendering-resources` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/resource-specific-editors` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/spatial-indexing` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/three-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/rendering-assets` enables `domain/webrtc-multiplayer-mesh` (synthesized lexical, confidence 0.25).
- `domain/rendering-backends` enables `domain/android-render-input` (synthesized lexical, confidence 0.25).
- `domain/rendering-device-resources` enables `domain/renderer-storage` (grounded, confidence 0.90).
- `domain/rendering-device-resources` enables `domain/scene-data-and-buffers` (grounded, confidence 0.90).
- `domain/rendering-device-resources` enables `domain/shader-programs` (asserted, confidence 0.30).
- `domain/rendering-device-resources` enables `domain/textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `domain/rendering-device-resources` enables `domain/viewport-render-data` (asserted, confidence 0.30).
- `domain/resource-assets` enables `domain/audio-bus-processing` (grounded, confidence 0.90).
- `domain/resource-assets` enables `domain/device-runtime` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/gpu-command-encoding` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/gpu-resources-pipelines` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/physics-collision` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/renderer-storage` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/rendering-assets` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/rendering-resources` (grounded, confidence 0.90).
- `domain/resource-assets` enables `domain/resource-previews` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/scene-contexts` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/scene-state` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/skeletal-ik` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/skeleton-modification` (grounded, confidence 0.90).
- `domain/resource-assets` enables `domain/temporal-upscaling-dispatch` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/textures-and-placeholders` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/themes-and-style-boxes` (grounded, confidence 0.90).
- `domain/resource-assets` enables `domain/three-dimensional-content` (grounded, confidence 0.90).
- `domain/resource-assets` enables `domain/two-dimensional-content` (grounded, confidence 0.90).
- `domain/resource-assets` enables `domain/ui-theming` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `domain/resource-formats` enables `domain/resource-identifiers` (grounded, confidence 0.90).
- `domain/resource-formats` enables `domain/scene-state` (grounded, confidence 0.90).
- `domain/resource-formats` enables `domain/textures-and-placeholders` (grounded, confidence 0.90).
- `domain/resource-formats` enables `domain/tile-libraries` (grounded, confidence 0.90).
- `domain/resource-identifiers` enables `domain/generic-containers` (synthesized lexical, confidence 0.25).
- `domain/resource-identifiers` enables `domain/rendering-device-resources` (grounded, confidence 0.90).
- `domain/resource-identifiers` enables `domain/resource-dependency-resolution` (synthesized lexical, confidence 0.25).
- `domain/resource-identifiers` enables `domain/scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `domain/resource-loading` enables `domain/image-format-loading` (synthesized lexical, confidence 0.25).
- `domain/resource-loading` enables `domain/packed-resource-files` (grounded, confidence 0.90).
- `domain/resource-loading` enables `domain/resource-formats` (grounded, confidence 0.90).
- `domain/resource-previews` enables `domain/editor-extensibility` (synthesized lexical, confidence 0.25).
- `domain/resource-previews` enables `domain/localization-and-template-generation` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/allocation` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/audio-bus-processing` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/crypto-resources` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/editing-selection-history` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/editor-extensibility` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/filesystem-project-index` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/gdextension-libraries` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/glyph-caching` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/gpu-command-encoding` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/gpu-memory-suballocation` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/image-codec-registration` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/input-events-actions` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/native-extensions` (grounded, confidence 0.90).
- `domain/resources` enables `domain/object-model` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/ogg-vorbis-audio-streams` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/packed-resource-files` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/packed-scenes` (grounded, confidence 0.90).
- `domain/resources` enables `domain/physics-collision` (grounded, confidence 0.90).
- `domain/resources` enables `domain/png-image-codec` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/project-settings` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/rendering-assets` (grounded, confidence 0.90).
- `domain/resources` enables `domain/rendering-resources` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-assets` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-bundle-data` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-dependency-resolution` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-formats` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-identifiers` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-importing` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-loading` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-previews` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/resource-specific-editors` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/scripting` (grounded, confidence 0.90).
- `domain/resources` enables `domain/shader-programs` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/skeletal-ik` (grounded, confidence 0.90).
- `domain/resources` enables `domain/skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/theora-video-streams` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/three-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/tile-libraries` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/two-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/resources` enables `domain/variant-text-serialization` (synthesized lexical, confidence 0.25).
- `domain/rigid-body-motion` enables `domain/constraints` (grounded, confidence 0.90).
- `domain/rigid-body-motion` enables `domain/vehicle-dynamics` (grounded, confidence 0.90).
- `domain/run-management` enables `domain/manifold-mesh-data` (synthesized lexical, confidence 0.25).
- `domain/run-management` enables `domain/skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `domain/run-management` enables `domain/target-platform-export` (synthesized lexical, confidence 0.25).
- `domain/runtime-configuration` enables `domain/engine-bootstrap` (grounded, confidence 0.90).
- `domain/runtime-configuration` enables `domain/input-action-configuration` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/android-gradle-assembly` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/cpu-specialized-dsp` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/motion-blur` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/project-catalog` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/ray-query` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/sdl-platform-backends` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/wraparound-network-time` (synthesized lexical, confidence 0.25).
- `domain/runtime-metadata` enables `domain/extensions` (grounded, confidence 0.90).
- `domain/runtime-metadata` enables `domain/native-extensions` (synthesized lexical, confidence 0.25).
- `domain/runtime-metadata` enables `domain/reflection` (synthesized lexical, confidence 0.25).
- `domain/runtime-metadata` enables `domain/reflection-metadata` (synthesized lexical, confidence 0.25).
- `domain/runtime-metadata` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/runtime-metadata` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/runtime-metadata` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/runtime-metadata` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/scene-authoring` enables `domain/scene-tree-and-signal-connections` (grounded, confidence 0.90).
- `domain/scene-commit` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/editing-selection-history` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/editor-authoring-workspaces` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/editor-plugin-state` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/engine-bootstrap` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/explicit-drm-syncobj` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/gui-control-layout` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/instancing` (grounded, confidence 0.90).
- `domain/scene-commit` enables `domain/openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/openxr-render-models` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/primitive-references` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/project-catalog` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/property-introspection` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/resource-assets` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/resources` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/runtime-configuration` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-authoring` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-contexts` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-importing` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-state` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-tree` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-tree-execution` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/screen-and-environment-effects` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/tile-authoring` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/undo-redo-history` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/version-control-integration` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/editing-selection-history` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/editor-authoring-workspaces` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/editor-plugin-state` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/engine-bootstrap` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/gui-control-layout` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/instancing` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/openxr-composition-layers` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/openxr-render-models` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/packed-scenes` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/primitive-references` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/project-catalog` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/property-introspection` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/resource-assets` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/resources` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/runtime-configuration` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-authoring` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-commit` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-importing` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-state` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-tree` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/scene-tree-execution` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/screen-and-environment-effects` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/tile-authoring` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/two-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `domain/undo-redo-history` (synthesized lexical, confidence 0.25).
- `domain/scene-contexts` enables `language/nodepaths-and-indexed-access` (grounded, confidence 0.90).
- `domain/scene-data-and-buffers` enables `domain/screen-and-environment-effects` (grounded, confidence 0.90).
- `domain/scene-graph` enables `domain/gui-control-layout` (grounded, confidence 0.90).
- `domain/scene-graph` enables `domain/navigation-agents` (grounded, confidence 0.90).
- `domain/scene-graph` enables `domain/particle-systems` (grounded, confidence 0.90).
- `domain/scene-graph` enables `domain/skeletal-modifiers-and-ik` (grounded, confidence 0.90).
- `domain/scene-graph` enables `domain/three-dimensional-physics` (grounded, confidence 0.90).
- `domain/scene-graph` enables `domain/tilemap-layer-data` (grounded, confidence 0.90).
- `domain/scene-graph` enables `domain/two-dimensional-physics` (grounded, confidence 0.90).
- `domain/scene-importing` enables `domain/collada-import` (grounded, confidence 0.90).
- `domain/scene-importing` enables `domain/post-import-processing` (grounded, confidence 0.90).
- `domain/scene-replication-configuration` enables `domain/replicated-property-synchronization` (grounded, confidence 0.90).
- `domain/scene-replication-configuration` enables `domain/replicated-spawning` (grounded, confidence 0.90).
- `domain/scene-state` enables `domain/packed-scenes` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/expression-evaluation` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/input-actions` (grounded, confidence 0.90).
- `domain/scene-tree` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/objectdb-snapshots` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/packed-scenes` (grounded, confidence 0.90).
- `domain/scene-tree` enables `domain/physics-collision` (grounded, confidence 0.90).
- `domain/scene-tree` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/rendering-assets` (grounded, confidence 0.90).
- `domain/scene-tree` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/scene-importing` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/scene-tree-execution` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/ui-composition` (grounded, confidence 0.90).
- `domain/scene-tree` enables `domain/visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `language/gdscript-signals-callables` (grounded, confidence 0.90).
- `domain/scene-tree-and-signal-connections` enables `domain/scene-tree` (synthesized lexical, confidence 0.25).
- `domain/script-registration-metadata` enables `language/csharp-source-generation` (grounded, confidence 0.90).
- `domain/scripting` enables `domain/completion-contexts` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/editor-plugin-lifecycle` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/filesystem-project-index` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/fixture-contracts` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/gdscript-bytecode-runtime` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/gdscript-editor-services` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/gdscript-static-analysis` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/script-hosting` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/script-registration-metadata` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/unicode-data` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `domain/visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `domain/scripting` enables `language/gdscript-declarations` (grounded, confidence 0.90).
- `domain/sdl-event-queue` enables `domain/platform-display-adaptation` (synthesized lexical, confidence 0.25).
- `domain/sdl-platform-backends` enables `domain/hid-device-enumeration` (synthesized lexical, confidence 0.25).
- `domain/sdl-platform-backends` enables `domain/profiling-interning` (synthesized lexical, confidence 0.25).
- `domain/sdl-platform-backends` enables `domain/rendering-backends` (synthesized lexical, confidence 0.25).
- `domain/sdl-platform-backends` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/sdl-platform-backends` enables `domain/vector-font-export` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/editing-selection-history` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/manifold-mesh-data` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/property-inspection` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/property-introspection` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/reflection` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/reflection-metadata` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/runtime-metadata` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/scene-replication-configuration` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/scene-state` (synthesized lexical, confidence 0.25).
- `domain/sdl-runtime-properties` enables `domain/sdl-iostreams` (grounded, confidence 0.90).
- `domain/sdl-runtime-properties` enables `domain/text-shaping-plans` (synthesized lexical, confidence 0.25).
- `domain/sdl-thread-abstractions` enables `domain/deferred-execution` (synthesized lexical, confidence 0.25).
- `domain/sdl-thread-abstractions` enables `domain/native-platform-adapters` (synthesized lexical, confidence 0.25).
- `domain/sdl-thread-abstractions` enables `domain/task-scheduling` (synthesized lexical, confidence 0.25).
- `domain/serialization` enables `domain/state-recording` (grounded, confidence 0.90).
- `domain/sfnt-tables` enables `domain/font-table-validation` (grounded, confidence 0.90).
- `domain/sfnt-tables` enables `domain/variable-font-data` (grounded, confidence 0.90).
- `domain/shader-editing-and-language-plugins` enables `domain/platform-selective-shader-baking` (grounded, confidence 0.90).
- `domain/shader-interface-analysis` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/shader-interface-metadata` enables `domain/metal-presentation` (synthesized lexical, confidence 0.25).
- `domain/shader-interface-metadata` enables `domain/openxr-extension-wrappers` (synthesized lexical, confidence 0.25).
- `domain/shader-interface-metadata` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/shader-interface-metadata` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/shader-interface-metadata` enables `domain/web-javascript-bridge` (synthesized lexical, confidence 0.25).
- `domain/shader-interface-metadata` enables `domain/webxr-session-bridge` (synthesized lexical, confidence 0.25).
- `domain/shader-interface-metadata` enables `domain/zlib-stream-codec` (synthesized lexical, confidence 0.25).
- `domain/shader-language-front-end` enables `domain/spirv-generation` (grounded, confidence 0.90).
- `domain/shader-programs` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/shader-programs` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/shader-programs` enables `domain/spirv-generation` (synthesized lexical, confidence 0.25).
- `domain/shader-programs` enables `domain/spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `domain/shader-programs` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/shader-programs` enables `language/godot-shader-language` (grounded, confidence 0.90).
- `domain/shader-reflection` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/shader-reflection` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/shader-reflection` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/shader-source-compilation` enables `domain/shader-interface-analysis` (grounded, confidence 0.90).
- `domain/signal-awaitability` enables `domain/android-instrumented-tests` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/android-plugin-signals` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/godot-member-exposure` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/reflection` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/reflection-metadata` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/runtime-metadata` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `domain/signed-distance-fields` enables `domain/glyph-rasterization` (synthesized lexical, confidence 0.25).
- `domain/simd-accelerated-codecs` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/simd-accelerated-codecs` enables `domain/simd-ray-packets` (synthesized lexical, confidence 0.25).
- `domain/simd-ray-packets` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/3d-gizmo-authoring` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/gltf-node-hierarchy` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/skeletal-ik` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/skeleton-modification` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `domain/skeleton-modifiers` enables `domain/skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
- `domain/skeleton-modifiers` enables `domain/skeleton-modification` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/allocation` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/builder-heuristics` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/bvh-construction` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/bvh-traversal` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/motion-blur` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/scene-commit` (synthesized lexical, confidence 0.25).
- `domain/spatial-predictive-filters` enables `domain/collision-filtering` (synthesized lexical, confidence 0.25).
- `domain/spatial-predictive-filters` enables `domain/filter-callbacks` (synthesized lexical, confidence 0.25).
- `domain/spirv-generation` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/spirv-generation` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-generation` enables `domain/spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `domain/spirv-generation` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-intermediate-representation` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/spirv-intermediate-representation` enables `domain/shader-reflection` (grounded, confidence 0.90).
- `domain/spirv-intermediate-representation` enables `domain/spirv-generation` (synthesized lexical, confidence 0.25).
- `domain/spirv-intermediate-representation` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/spirv-generation` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-shader-reflection` enables `domain/shader-interface-metadata` (grounded, confidence 0.90).
- `domain/spirv-shader-reflection` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-shader-reflection` enables `domain/spirv-generation` (synthesized lexical, confidence 0.25).
- `domain/spirv-shader-reflection` enables `domain/spirv-intermediate-representation` (synthesized lexical, confidence 0.25).
- `domain/string-names-paths` enables `domain/localization` (grounded, confidence 0.90).
- `domain/string-names-paths` enables `domain/profiling-interning` (grounded, confidence 0.90).
- `domain/string-names-paths` enables `domain/scene-replication-configuration` (synthesized lexical, confidence 0.25).
- `domain/string-names-paths` enables `domain/variant-text-serialization` (grounded, confidence 0.90).
- `domain/subdivision-surface-evaluation` enables `domain/feature-adaptive-tessellation` (grounded, confidence 0.90).
- `domain/subdivision-surface-evaluation` enables `domain/geometry-families` (synthesized lexical, confidence 0.25).
- `domain/subdivision-surface-evaluation` enables `domain/half-edge-topology` (synthesized lexical, confidence 0.25).
- `domain/synchronization-primitives` enables `domain/concurrency` (synthesized lexical, confidence 0.25).
- `domain/synchronization-primitives` enables `domain/sdl-event-queue` (synthesized lexical, confidence 0.25).
- `domain/synchronization-primitives` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/target-platform-export` enables `domain/apple-embedded-packaging` (grounded, confidence 0.90).
- `domain/target-platform-export` enables `domain/platform-exports` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/concurrency` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/deferred-execution` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/device-runtime` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/encoder-configuration` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/synchronization-primitives` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/wayland-window-backend` (synthesized lexical, confidence 0.25).
- `domain/task-scheduling` enables `domain/worker-parallelism` (synthesized lexical, confidence 0.25).
- `domain/text-and-localization` enables `domain/localization` (synthesized lexical, confidence 0.25).
- `domain/texture-block-codecs` enables `domain/block-texture-encoding` (synthesized lexical, confidence 0.25).
- `domain/texture-compression-pipeline` enables `domain/basis-container-layout` (synthesized lexical, confidence 0.25).
- `domain/texture-compression-pipeline` enables `domain/basis-transcoding` (synthesized lexical, confidence 0.25).
- `domain/texture-format-description` enables `domain/ktx-texture-container` (grounded, confidence 0.90).
- `domain/themes-and-style-boxes` enables `domain/editor-theming` (synthesized lexical, confidence 0.25).
- `domain/themes-and-style-boxes` enables `domain/ui-theming` (synthesized lexical, confidence 0.25).
- `domain/theora-block-video-codec` enables `domain/theora-video-streams` (synthesized lexical, confidence 0.25).
- `domain/theora-video-streams` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/tile-authoring` enables `domain/tile-libraries` (synthesized lexical, confidence 0.25).
- `domain/tile-libraries` enables `domain/tilemap-layer-data` (synthesized lexical, confidence 0.25).
- `domain/tile-libraries` enables `domain/two-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/crypto-resources` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/network-io` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/stream-networking` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/tls-crypto-backend` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/x509-certificate-processing` (grounded, confidence 0.90).
- `domain/tls-crypto-backend` enables `domain/crypto-resources` (synthesized lexical, confidence 0.25).
- `domain/tls-crypto-backend` enables `domain/network-io` (synthesized lexical, confidence 0.25).
- `domain/tls-crypto-backend` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/tls-crypto-backend` enables `domain/stream-networking` (synthesized lexical, confidence 0.25).
- `domain/tls-crypto-backend` enables `domain/tls-connection-state` (synthesized lexical, confidence 0.25).
- `domain/tls-crypto-backend` enables `domain/x509-certificate-processing` (synthesized lexical, confidence 0.25).
- `domain/transform-interpolation` enables `domain/metalfx-scaling` (synthesized lexical, confidence 0.25).
- `domain/transform-quantization-rate-distortion` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/triangle-intersection-algorithms` enables `domain/motion-blur-geometry` (grounded, confidence 0.90).
- `domain/triangle-intersection-algorithms` enables `domain/simd-ray-packets` (grounded, confidence 0.90).
- `domain/triangle-intersection-algorithms` enables `domain/subgrid-intersection` (grounded, confidence 0.90).
- `domain/ui-composition` enables `domain/curve-and-patch-bases` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/endian-safe-binary-io` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/gdscript-bytecode-runtime` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/gui-control-layout` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/gui-controls` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/platform-presentation` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/regex-jit-codegen` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/rendering-assets` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/ui-theming` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/version-control-integration` (synthesized lexical, confidence 0.25).
- `domain/ui-composition` enables `domain/websocket-frame-events` (synthesized lexical, confidence 0.25).
- `domain/ui-theming` enables `domain/editor-theming` (synthesized lexical, confidence 0.25).
- `domain/ui-theming` enables `domain/themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `domain/undo-redo-history` enables `domain/editor-scene-sessions` (synthesized lexical, confidence 0.25).
- `domain/unicode-data` enables `domain/break-iteration` (grounded, confidence 0.90).
- `domain/unicode-data` enables `domain/character-encoding-conversion` (grounded, confidence 0.90).
- `domain/unicode-data` enables `domain/identifier-spoof-checking` (grounded, confidence 0.90).
- `domain/unicode-data` enables `domain/unicode-normalization` (grounded, confidence 0.90).
- `domain/unicode-normalization` enables `domain/unicode-data` (synthesized lexical, confidence 0.25).
- `domain/upnp-device-control` enables `domain/upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `domain/upnp-device-discovery` enables `domain/upnp-port-mapping` (grounded, confidence 0.90).
- `domain/upnp-port-mapping` enables `domain/upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `domain/variable-font-subsetting` enables `domain/variable-font-data` (synthesized lexical, confidence 0.25).
- `domain/variant-containers` enables `domain/break-iteration` (synthesized lexical, confidence 0.25).
- `domain/variant-containers` enables `domain/dynamic-values` (synthesized lexical, confidence 0.25).
- `domain/variant-containers` enables `domain/generic-containers` (grounded, confidence 0.90).
- `domain/variant-containers` enables `domain/packet-queries` (synthesized lexical, confidence 0.25).
- `domain/variant-containers` enables `domain/vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `domain/vector-font-export` enables `domain/generic-containers` (synthesized lexical, confidence 0.25).
- `domain/vector-font-export` enables `domain/multichannel-distance-fields` (synthesized lexical, confidence 0.25).
- `domain/vector-font-export` enables `domain/visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `domain/vi-native-surface-creation` enables `domain/vulkan-extensible-records` (synthesized lexical, confidence 0.25).
- `domain/visual-shader-module-build` enables `domain/visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `domain/visual-shader-nodes` enables `domain/visual-shader-vector-operations` (grounded, confidence 0.90).
- `domain/vulkan-instance-setup` enables `domain/openxr-runtime-loading` (synthesized lexical, confidence 0.25).
- `domain/vulkan-swapchain-presentation` enables `domain/canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `domain/vulkan-swapchain-presentation` enables `domain/metal-presentation` (synthesized lexical, confidence 0.25).
- `domain/wayland-window-backend` enables `domain/explicit-drm-syncobj` (synthesized lexical, confidence 0.25).
- `domain/wayland-window-backend` enables `domain/linuxbsd-desktop-integration` (synthesized lexical, confidence 0.25).
- `domain/webp-image-io` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/webp-image-io` enables `domain/webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `domain/webp-riff-decoding` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/webp-riff-decoding` enables `domain/webp-image-io` (synthesized lexical, confidence 0.25).
- `domain/webrtc-data-channels` enables `domain/webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `domain/webrtc-peer-connections` enables `domain/webrtc-data-channels` (grounded, confidence 0.90).
- `domain/webrtc-peer-connections` enables `domain/webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `domain/websocket-frame-events` enables `domain/websocket-multiplayer` (synthesized lexical, confidence 0.25).
- `domain/websocket-multiplayer` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/websocket-peers` enables `domain/websocket-multiplayer` (grounded, confidence 0.90).
- `domain/xr-tracking` enables `domain/pluggable-server-backends` (synthesized lexical, confidence 0.25).
- `domain/xr-tracking` enables `domain/skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
- `domain/yuv-rgb-conversion` enables `domain/input-picture-planes` (synthesized lexical, confidence 0.25).
- `domain/zip-archive-io` enables `domain/target-platform-export` (synthesized lexical, confidence 0.25).
- `domain/zip-archive-io` enables `domain/zip64-archive-io` (synthesized lexical, confidence 0.25).
- `domain/zlib-stream-codec` enables `domain/openexr-image-decoding` (synthesized lexical, confidence 0.25).
- `language/annotations-static-state-and-lifecycle` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/annotations-static-state-and-lifecycle` enables `language/cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `language/annotations-static-state-and-lifecycle` enables `language/cpp-runtime-casts` (synthesized lexical, confidence 0.25).
- `language/annotations-static-state-and-lifecycle` enables `language/cpp-static-generated-data` (synthesized lexical, confidence 0.25).
- `language/annotations-static-state-and-lifecycle` enables `language/cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `language/annotations-static-state-and-lifecycle` enables `language/properties-and-accessors` (synthesized lexical, confidence 0.25).
- `language/annotations-static-state-and-lifecycle` enables `language/typed-collections` (synthesized lexical, confidence 0.25).
- `language/c-abi-data-structures` enables `language/c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `domain/shader-interface-metadata` (grounded, confidence 0.90).
- `language/c-abi-manual-lifetime` enables `domain/spirv-shader-reflection` (grounded, confidence 0.90).
- `language/c-abi-manual-lifetime` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-abi-versioned-initialization` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-concurrent-state` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-dynamic-library-wrappers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-explicit-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-function-pointer-tables` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-macro-codecs` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-parser-state` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-platform-selection` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-pointers-arrays-and-strides` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-preprocessor-composition` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-stateful-streaming-api` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-structs-pointers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-classes` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-atomics` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-c-abi` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `language/c-abi-manual-lifetime` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/c-abi-versioned-initialization` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/c-abi-versioned-initialization` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/c-abi-versioned-initialization` enables `language/csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `language/c-aggregate-callback-modules` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/c-aggregate-callback-modules` enables `language/c-explicit-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `language/c-aggregate-callback-modules` enables `language/c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-aggregate-callback-modules` enables `language/c-function-pointer-tables` (synthesized lexical, confidence 0.25).
- `language/c-aggregate-callback-modules` enables `language/c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `language/c-aggregate-callback-modules` enables `language/cpp-basis-transcoding` (grounded, confidence 0.90).
- `language/c-compatible-header-guards` enables `language/c-abi-record-and-dispatch` (grounded, confidence 0.90).
- `language/c-function-pointer-apis` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-callbacks` enables `language/c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-callbacks` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-callbacks` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-interfaces` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-interfaces` enables `language/c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-interfaces` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-interfaces` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-tables` enables `language/c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-tables` enables `language/c-stateful-streaming-api` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-tables` enables `language/cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-tables` enables `language/cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-tables` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/c-function-pointer-tables` enables `language/javascript-web-runtime` (synthesized lexical, confidence 0.25).
- `language/c-integer-bitwise-packing` enables `domain/endian-safe-binary-io` (grounded, confidence 0.90).
- `language/c-integer-bitwise-packing` enables `language/cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `language/c-macro-codecs` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/c-macro-codecs` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/c-macro-codecs` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `domain/simd-accelerated-codecs` (grounded, confidence 0.90).
- `language/c-platform-selection` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/c-platform-selection` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-abi-manual-lifetime` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-resource-lifecycles` (grounded, confidence 0.90).
- `language/c-pointers-arrays-and-strides` enables `language/c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-structs-pointers` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-composition` enables `domain/codec-simd-specialization` (grounded, confidence 0.90).
- `language/c-preprocessor-composition` enables `domain/freetype-module-composition` (grounded, confidence 0.90).
- `language/c-preprocessor-composition` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-composition` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-composition` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-composition` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-composition` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-composition` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `domain/openexr-image-decoding` (grounded, confidence 0.90).
- `language/c-preprocessor-configuration` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-configuration` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-platform-selection` enables `domain/sdl-platform-backends` (grounded, confidence 0.90).
- `language/c-preprocessor-platform-selection` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-platform-selection` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-platform-selection` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-platform-selection` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-platform-selection` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-portability` enables `domain/cpu-specialized-dsp` (grounded, confidence 0.90).
- `language/c-preprocessor-portability` enables `domain/endian-safe-binary-io` (grounded, confidence 0.90).
- `language/c-preprocessor-portability` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-portability` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-portability` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-portability` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-portability` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/c-preprocessor-portability` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/c-resource-lifecycles` enables `language/c-structs-pointers` (synthesized lexical, confidence 0.25).
- `language/c-resource-lifecycles` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-resource-lifecycles` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/c-resource-lifecycles` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/c-stateful-streaming-api` enables `domain/brotli-stream-decompression` (grounded, confidence 0.90).
- `language/c-stateful-struct-apis` enables `domain/tls-connection-state` (grounded, confidence 0.90).
- `language/c-stateful-struct-apis` enables `language/callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `language/c-stateful-struct-apis` enables `language/nodepaths-and-indexed-access` (synthesized lexical, confidence 0.25).
- `language/c-structs-pointers` enables `language/c-concurrent-state` (grounded, confidence 0.90).
- `language/c-structs-pointers` enables `language/c-explicit-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `language/c-structs-pointers` enables `language/c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-structs-pointers` enables `language/c-function-pointer-tables` (grounded, confidence 0.90).
- `language/c-structs-pointers` enables `language/c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-dynamic-library-wrappers` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-explicit-resource-lifecycles` (grounded, confidence 0.90).
- `language/c-structures-and-pointers` enables `language/c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-function-pointer-interfaces` (grounded, confidence 0.90).
- `language/c-structures-and-pointers` enables `language/c-function-pointer-tables` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-pointers-arrays-and-strides` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-resource-lifecycles` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-stateful-streaming-api` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/cpp-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/cpp-object-representation-casts` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/cpp-runtime-casts` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/c-structures-and-pointers` enables `language/cxx-c-abi` (synthesized lexical, confidence 0.25).
- `language/callables-and-lambdas` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/callables-and-lambdas` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/callables-and-lambdas` enables `language/cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `language/callables-and-lambdas` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/callables-and-lambdas` enables `language/signals-and-await` (grounded, confidence 0.90).
- `language/classes-inheritance-and-lookup` enables `language/annotations-static-state-and-lifecycle` (grounded, confidence 0.90).
- `language/classes-inheritance-and-lookup` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/gdscript-declarations` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/properties-and-accessors` (grounded, confidence 0.90).
- `language/cplusplus-enumerated-export-state` enables `language/cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `language/cplusplus-export-callbacks` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/cplusplus-export-callbacks` enables `language/c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `language/cplusplus-export-callbacks` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `domain/gpu-memory-suballocation` (grounded, confidence 0.90).
- `language/cpp-abstraction-and-polymorphism` enables `language/annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-byte-exact-binary-parsing` (grounded, confidence 0.90).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-simd-intrinsics` (grounded, confidence 0.90).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-templates-and-generic-containers` (grounded, confidence 0.90).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/gdscript-declarations` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/properties-and-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-atomics` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-c-abi` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-basis-transcoding` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/cpp-binary-data-codecs` enables `domain/gltf-accessors` (grounded, confidence 0.90).
- `language/cpp-byte-exact-binary-parsing` enables `domain/astc-block-coding` (grounded, confidence 0.90).
- `language/cpp-byte-exact-binary-parsing` enables `domain/basis-container-layout` (grounded, confidence 0.90).
- `language/cpp-byte-exact-binary-parsing` enables `domain/ktx2-container-transcoding` (grounded, confidence 0.90).
- `language/cpp-byte-exact-binary-parsing` enables `domain/prefix-code-decoding` (grounded, confidence 0.90).
- `language/cpp-byte-exact-binary-parsing` enables `domain/raster-image-input` (grounded, confidence 0.90).
- `language/cpp-class-inheritance` enables `domain/pluggable-server-backends` (grounded, confidence 0.90).
- `language/cpp-class-inheritance` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-atomics` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-c-abi` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-numeric-value-operations` (grounded, confidence 0.90).
- `language/cpp-class-state-and-composition` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/gdscript-declarations` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-class-state-and-composition` enables `language/properties-and-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `domain/simd-ray-packets` (grounded, confidence 0.90).
- `language/cpp-class-templates` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-overloading` (grounded, confidence 0.90).
- `language/cpp-classes` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-virtual-dispatch` (grounded, confidence 0.90).
- `language/cpp-classes` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/gdscript-declarations` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/properties-and-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cpp-c-linkage-adapters` (grounded, confidence 0.90).
- `language/cpp-classes-and-inheritance` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cpp-plugin-specialization` (grounded, confidence 0.90).
- `language/cpp-classes-and-inheritance` enables `language/cpp-resource-lifetime` (grounded, confidence 0.90).
- `language/cpp-classes-and-inheritance` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cpp-visitor-traversal` (grounded, confidence 0.90).
- `language/cpp-classes-and-inheritance` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-inheritance` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-atomics` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-c-abi` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `domain/godot-enet-socket-adaptation` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-class-templates` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-copy-on-write-containers` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-godot-binding-macros` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-runtime-polymorphism` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-scoped-locks` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-const-borrowing` enables `language/cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `language/cpp-const-borrowing` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-engine-polymorphism` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-engine-polymorphism` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-engine-polymorphism` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-engine-polymorphism` enables `language/objective-cpp-apple-adapters` (grounded, confidence 0.90).
- `language/cpp-function-pointer-interfaces` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `domain/enet-transport-and-multiplayer` (grounded, confidence 0.90).
- `language/cpp-interface-polymorphism` enables `domain/extensions` (grounded, confidence 0.90).
- `language/cpp-interface-polymorphism` enables `domain/gdscript-language-server` (asserted, confidence 0.30).
- `language/cpp-interface-polymorphism` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-runtime-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-virtual-dispatch` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/cpp-jni-variant-marshalling` enables `domain/android-jni-interop` (grounded, confidence 0.90).
- `language/cpp-module-import` enables `language/c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `language/cpp-module-import` enables `language/c-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cpp-module-import` enables `language/cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-module-import` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-basis-transcoding` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-classes-and-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-exception-abi-boundaries` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-frame-scheduling` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-godot-binding-macros` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-interface-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-numeric-value-operations` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-plugin-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-polymorphic-backends` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-resource-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-runtime-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-scope-locking` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-template-binding` (grounded, confidence 0.90).
- `language/cpp-native-classes` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-templates-and-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-typed-api-records` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-atomics` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-c-abi` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-conditional-compilation` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-native-classes` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/cpp-native-wrappers` enables `domain/metal-cpp-object-bridge` (grounded, confidence 0.90).
- `language/cpp-native-wrappers` enables `language/cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `language/cpp-native-wrappers` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-native-wrappers` enables `language/cpp-polymorphic-ownership` (grounded, confidence 0.90).
- `language/cpp-native-wrappers` enables `language/cpp-runtime-symbol-loading` (grounded, confidence 0.90).
- `language/cpp-native-wrappers` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/cpp-native-wrappers` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cpp-class-state-and-composition` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-numeric-value-operations` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `domain/audio-bus-processing` (asserted, confidence 0.30).
- `language/cpp-object-hierarchies` enables `domain/navigation-queries` (grounded, confidence 0.90).
- `language/cpp-object-hierarchies` enables `domain/networking` (grounded, confidence 0.90).
- `language/cpp-object-hierarchies` enables `domain/physics-server-queries` (grounded, confidence 0.90).
- `language/cpp-object-hierarchies` enables `domain/resource-assets` (grounded, confidence 0.90).
- `language/cpp-object-hierarchies` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-runtime-casts` (grounded, confidence 0.90).
- `language/cpp-object-hierarchies` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-ownership-and-polymorphism` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-pluggable-allocation` enables `language/c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `language/cpp-plugin-specialization` enables `domain/editor-plugin-lifecycle` (grounded, confidence 0.90).
- `language/cpp-polymorphic-backends` enables `language/cpp-virtual-dispatch` (synthesized lexical, confidence 0.25).
- `language/cpp-polymorphic-backends` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/cpp-static-abi-contracts` (grounded, confidence 0.90).
- `language/cpp-preprocessor-configuration` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/cpp-resource-and-polymorphism` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cpp-resource-lifetime` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-runtime-adapters` enables `language/cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-runtime-adapters` enables `language/cpp-jni-variant-marshalling` (grounded, confidence 0.90).
- `language/cpp-runtime-polymorphism` enables `domain/editor-plugin-extension` (grounded, confidence 0.90).
- `language/cpp-runtime-polymorphism` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-runtime-polymorphism` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cpp-scope-locking` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-scoped-locks` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-static-generated-data` enables `language/cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-generic-containers` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-generic-containers` enables `language/cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-generic-containers` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-specialization` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-specialization` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-specialization` enables `language/cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-specialization` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-views` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-specialized-marshalling` (grounded, confidence 0.90).
- `language/cpp-templates-traits` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-variadic-binding` (grounded, confidence 0.90).
- `language/cpp-templates-traits` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-typed-api-records` enables `language/cpp-c-type-mapping` (grounded, confidence 0.90).
- `language/cpp-typed-api-records` enables `language/cpp-flag-stringification` (grounded, confidence 0.90).
- `language/cpp-visitor-traversal` enables `domain/shader-interface-analysis` (grounded, confidence 0.90).
- `language/csharp-assembly-load-contexts` enables `domain/plugin-and-assembly-reload` (grounded, confidence 0.90).
- `language/csharp-attributes-reflection` enables `domain/godot-member-exposure` (grounded, confidence 0.90).
- `language/csharp-attributes-reflection` enables `language/c-abi-manual-lifetime` (synthesized lexical, confidence 0.25).
- `language/csharp-attributes-reflection` enables `language/csharp-source-generation` (grounded, confidence 0.90).
- `language/csharp-attributes-reflection` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/csharp-partial-source-generation` enables `domain/managed-csharp-scripting` (grounded, confidence 0.90).
- `language/csharp-partial-source-generation` enables `language/csharp-source-generation` (synthesized lexical, confidence 0.25).
- `language/csharp-partial-source-generation` enables `language/csharp-unsafe-interop` (synthesized lexical, confidence 0.25).
- `language/csharp-source-generation` enables `language/csharp-unsafe-interop` (grounded, confidence 0.90).
- `language/csharp-source-generation` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/csharp-source-generation` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/csharp-unsafe-interop` enables `domain/managed-native-interop` (grounded, confidence 0.90).
- `language/csharp-unsafe-interop` enables `language/c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `language/csharp-unsafe-interop` enables `language/c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/csharp-unsafe-interop` enables `language/cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `language/cxx-atomics` enables `language/c-concurrent-state` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-structs-pointers` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-structures-and-pointers` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-object-hierarchies` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cxx-enums-bitflags` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/gdscript-declarations` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/properties-and-accessors` (synthesized lexical, confidence 0.25).
- `language/cxx-closures` enables `language/c-function-pointer-apis` (synthesized lexical, confidence 0.25).
- `language/cxx-closures` enables `language/callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `language/cxx-closures` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/cxx-closures` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `domain/isa-portability` (grounded, confidence 0.90).
- `language/cxx-conditional-compilation` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/cxx-lambdas-standard-algorithms` enables `language/callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-class-registration` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-classes-and-ref-handles` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cxx-raii-reference-ownership` (grounded, confidence 0.90).
- `language/cxx-object-model` enables `language/cxx-reflection-macros` (grounded, confidence 0.90).
- `language/cxx-object-model` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/c-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/c-preprocessor-portability` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/cxx-simd-alignment` (grounded, confidence 0.90).
- `language/cxx-preprocessor-configuration` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/cxx-reflection-macros` enables `domain/serialization` (grounded, confidence 0.90).
- `language/cxx-reflection-macros` enables `language/cxx-stream-serialization` (grounded, confidence 0.90).
- `language/cxx-templates` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cxx-templates` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cxx-templates` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cxx-templates` enables `language/cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `language/cxx-templates` enables `language/cxx-wide-simd` (grounded, confidence 0.90).
- `language/cxx-templates` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/cxx-wide-simd` enables `language/c-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cxx-wide-simd` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cxx-wide-simd` enables `language/cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `language/gdscript-declarations` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/gdscript-declarations` enables `language/gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `language/gdscript-declarations` enables `language/gdscript-query-objects` (synthesized lexical, confidence 0.25).
- `language/gdscript-declarations` enables `language/gdscript-signals-callables` (grounded, confidence 0.90).
- `language/gdscript-declarations` enables `language/types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/c-abi-manual-lifetime` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/cpp-native-classes` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/gdscript-declarations` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/callables-and-lambdas` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/signals-and-await` (synthesized lexical, confidence 0.25).
- `language/gdscript-typed-collections` enables `domain/gdscript-static-analysis` (asserted, confidence 0.30).
- `language/gdscript-typed-collections` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/glsl-compute-shaders` enables `domain/gpu-texture-compression` (grounded, confidence 0.90).
- `language/godot-shader-language` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/godot-shader-language` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/godot-shader-language` enables `language/glsl-compute-shaders` (synthesized lexical, confidence 0.25).
- `language/iteration-protocols` enables `language/cpp-overloading` (synthesized lexical, confidence 0.25).
- `language/javascript-browser-ffi` enables `language/javascript-web-runtime` (synthesized lexical, confidence 0.25).
- `language/javascript-web-runtime` enables `domain/browser-runtime-adapters` (grounded, confidence 0.90).
- `language/javascript-web-runtime` enables `domain/web-javascript-bridge` (grounded, confidence 0.90).
- `language/javascript-web-runtime` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/objective-cpp-apple-adapters` enables `language/objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `language/objective-cpp-ios-adapters` enables `language/objective-cpp-apple-adapters` (synthesized lexical, confidence 0.25).
- `language/properties-and-accessors` enables `language/csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-scons-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-build-actions` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-scons-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-actions` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-scons-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-build-actions` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-scons-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-scons-build-hooks` enables `language/gdscript-query-objects` (synthesized lexical, confidence 0.25).
- `language/python-scons-build-hooks` enables `language/python-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-scons-build-hooks` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-scons-build-hooks` enables `language/python-scons-configuration` (synthesized lexical, confidence 0.25).
- `language/python-scons-build-hooks` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `domain/module-build-configuration` (grounded, confidence 0.90).
- `language/python-scons-configuration` enables `language/python-build-actions` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-scons-module-configuration` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-scons-module-configuration` enables `language/python-scons-configuration` (synthesized lexical, confidence 0.25).
- `language/signals-and-await` enables `language/csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `language/signals-and-await` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/signals-and-await` enables `language/warnings-and-suppression` (synthesized lexical, confidence 0.25).
- `language/typed-collections` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/typed-collections` enables `language/gdscript-typed-collections` (synthesized lexical, confidence 0.25).
- `language/types-inference-and-conversions` enables `language/cpp-jni-variant-marshalling` (synthesized lexical, confidence 0.25).
- `language/types-inference-and-conversions` enables `language/cpp-tagged-storage` (synthesized lexical, confidence 0.25).
- `language/types-inference-and-conversions` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/types-inference-and-conversions` enables `language/typed-collections` (grounded, confidence 0.90).
- `language/types-inference-and-conversions` enables `language/warnings-and-suppression` (asserted, confidence 0.30).
<!-- rope-ladder:end document -->
