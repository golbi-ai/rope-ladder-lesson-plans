<!-- rope-ladder:begin document a7dc197bd2ae2c6517bcdc73f541407c0c54806d5e4d625acf17b5604b791581 -->
# Lesson plan

Each unit stays within one prerequisite depth and is sized to keep a focused teaching-assistant reflection.

## Unit 1

Lesson ID: `unit-aabb-tree-construction-android-remote-engine-service-c-integer-bitwise-packing`

### Domain: AABB tree construction

Jolt's AABBTreeBuilder stores nodes and indexed triangles and accepts a maximum-triangles-per-leaf setting while building its tree.
Use these entity examples: AABBTreeBuilder.

### Domain: Remote Android engine service

Remote engine service messages convey engine status, errors, surface packages, dimensions, host tokens, and command-line parameters between Android UI and service code.
Use these entity examples: GodotService.EngineStatus.

### Implementation: C fixed-width integers and bitwise packing

Fixed-width integer values hold packed ARGB samples, codec fields, chunk tags, and arithmetic intermediates; shifts, masks, and bitwise operators extract or assemble channel values.
Apply it through: WebPData, WebPPicture.

## Unit 2

Lesson ID: `unit-animation-blending-c-preprocessor-configuration`

### Domain: Animation blending and playback

AnimationMixer, AnimationPlayer, AnimationTree, blend spaces, blend trees, and state machines compose animation playback.
Use these entity examples: AnimationMixer, AnimationNodeBlendTree, AnimationNodeStateMachine.

### Implementation: C preprocessor feature and platform configuration

Configuration headers use macros and conditional inclusion to enable ThorVG capabilities, select platform-dependent threading, and choose system or bundled SPIR-V headers.

## Unit 3

Lesson ID: `unit-build-time-source-generation-color-font-paint-c-preprocessor-portability`

### Domain: Build-time source generation

The Python build helpers generate C++ tables for documentation data, exporter registration, compressed documentation, and compressed translations.
Use these entity examples: EditorTranslationList.

### Domain: Color-font paint processing

HarfBuzz supplies paint functions plus bounded and extents contexts to evaluate color-font paint operations.
Use these entity examples: hb_paint_funcs_t, hb_paint_extents_context_t.

### Implementation: C preprocessor portability

Conditional compilation and macros select target-dependent code, inline helpers, endianness behavior, and instruction-family implementations at compile time.

## Unit 4

Lesson ID: `unit-cryptography-csharp-translation-extraction-cpp-module-import`

### Domain: Cryptographic keys, hashing, and TLS support

Crypto, CryptoKey, HashingContext, and HMACContext expose key generation, signatures, encryption, hashing, and message authentication APIs.

### Domain: C# translation extraction

The editor parses C# source into a Roslyn compilation and extracts constant translation-call arguments and nearby comments.

### Implementation: C++ standard-library module import

The Vulkan-Hpp façade contains an import std declaration and other headers test VULKAN_HPP_CXX_MODULE before falling back to textual includes.

## Unit 5

Lesson ID: `unit-debug-adapter-protocol-debugger-transport-cpp-module-interface`

### Domain: Debug Adapter Protocol integration

The editor's debug adapter serializes protocol content as JSON, manages protocol data types, and starts a local protocol server.
Use these entity examples: DAP::Source, DAP::Breakpoint, DAP::Capabilities.

### Domain: Debugger capture and transport

The debugger subsystem has engine captures and profilers, local and remote debugger implementations, and a socket-backed remote debugger peer.

### Implementation: C++ module interface partition

`vulkan_video.cppm` declares the exported `vulkan_hpp:video` module partition and an exported video namespace.
Apply it through: VideoSessionCreateInfoKHR.

## Unit 6

Lesson ID: `unit-dotnet-editor-build-integration-editor-theming-csharp-assembly-load-contexts`

### Domain: Editor build integration

Editor tools generate .NET projects, invoke builds, and surface CSV diagnostics.
Use these entity examples: BuildDiagnostic.

### Domain: Editor theming

Editor theming builds a Theme from ThemeConfiguration and settings, with distinct classic and modern builders plus font, icon, color-map, and scale support.

### Implementation: C# AssemblyLoadContext plugin loading

The editor plugin loader creates AssemblyLoadContext instances to resolve dependencies and support unloading.

## Unit 7

Lesson ID: `unit-fbx-gltf-scene-import-csharp-attributes-reflection`

### Domain: FBX scene import through GLTF structures

FBXDocument and FBXState specialize GLTF document and state structures while importers expose UFBX and FBX2GLTF editor entry points.
Use these entity examples: FBXDocument, FBXState.

### Implementation: C# attributes and reflection

Attributes declare engine-facing metadata and reflection reads it to connect types, members, and script information.
Apply it through: ScriptPathAttribute, AssemblyHasScriptsAttribute.

## Unit 8

Lesson ID: `unit-file-and-config-persistence-font-subsetting-glsl-compute-workgroups`

### Domain: Files, directories, and configuration

FileAccess and DirAccess read or write project and user paths, while ConfigFile stores section and key values.
Use these entity examples: ConfigFile.

### Domain: Font subsetting

HarfBuzz represents caller selections as a subset input and derives an in-memory plan for rewriting a font.
Use these entity examples: hb_subset_input_t, hb_subset_plan_t.

### Implementation: GLSL compute workgroups and synchronization

Compute shaders use declared resource interfaces, workgroup identifiers, bounds checks, shared arrays, and memory barriers for image-processing operations.

## Unit 9

Lesson ID: `unit-font-table-access-gamepad-motion-fusion-iteration-protocols`

### Domain: Binary font-table access

Graphite and HarfBuzz read binary font tables, including cmap mappings, to obtain glyph data and layout behavior.
Use these entity examples: Face, hb_blob_t, hb_face_t.

### Domain: Gamepad motion fusion and calibration

GamepadMotion maintains orientation, gravity, processed acceleration, and automatic gyro-calibration state from gyro and accelerometer samples.

### Implementation: Iteration protocols

`for` loops cover ranges, collections, and custom objects that provide iterator hook functions.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 10

Lesson ID: `unit-geometry-resources-graphics-procedure-loading-pattern-matching`

### Domain: Geometry resources

A Geometry represents typed primitive data, supplies bounds, and provides the common base used by mesh, curve, point, line, grid, user, and instance implementations.
Use these entity examples: RTCGeometry, Geometry.

### Domain: Graphics procedure loading

Generated glad loaders resolve OpenGL, EGL, and GLX procedure entry points and track available versions and extensions.

### Implementation: Pattern matching and guards

`match` supports literals, arrays, dictionaries, bindings, multiple patterns, wildcards, and guarded branches.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 11

Lesson ID: `unit-ide-messaging-protocol-input-midi-python-scons-configuration`

### Domain: IDE messaging protocol

IDE clients and the editor exchange framed requests and responses after a handshake, with JSON bodies for typed requests.
Use these entity examples: Message, GodotIdeMetadata.

### Domain: Keyboard and MIDI parsing

Keyboard utilities map key codes and names, while MIDIDriver::Parser interprets status and data bytes into MIDI messages.

### Implementation: Python SCons configuration functions

Python build files define can_build and configure functions and call SCons environment methods to select module sources and generated artifacts.

## Unit 12

Lesson ID: `unit-input-picture-planes-interactive-music-transitions-ios-embedded-adapter`

### Domain: Picture planes and pixel ownership

WebPPicture represents one image as either ARGB pixels or YUV(A) planes with dimensions, strides, optional alpha, and private allocation pointers.
Use these entity examples: WebPPicture.

### Domain: interactive music transition tables

AudioStreamInteractive combines clips with a TransitionKey-indexed transition map so playback can select a rule for a clip-to-clip change.
Use these entity examples: AudioStreamInteractive, AudioStreamInteractive Transition.

### Domain: iOS embedded adapter

iOS platform classes adapt the engine’s OS, display, and view responsibilities for an embedded Apple target.

## Unit 13

Lesson ID: `unit-linuxbsd-desktop-integration-locale-resolution-memory-allocation`

### Domain: Linux/BSD desktop integration

Linux/BSD platform code integrates OS services, portals, screensaver handling, TTS, X11, and Wayland display servers.

### Domain: Locale resolution

ICU parses locale identifiers and matches desired locales against supported locales using LSR values, likely-subtag data, and locale-distance data.
Use these entity examples: Locale, LSR, LocaleMatcher.

### Domain: Engine allocation helpers

Memory helpers define default allocators, typed allocators, allocation result wrappers, and global nil support.

## Unit 14

Lesson ID: `unit-mobile-vr-interface-native-desktop-services-openxr-structure-chains`

### Domain: mobile XR interface

MobileVRInterface is the module's concrete XRInterface implementation.

### Domain: Native desktop services

macOS and Windows platform code includes native-menu and text-to-speech adapters alongside their display servers.

### Domain: OpenXR structure chains

OpenXR input and output structures carry a typed structure identifier and a `next` pointer for extensible structure chains.
Use these entity examples: XrInstanceCreateInfo, XrBaseInStructure.

## Unit 15

Lesson ID: `unit-physics-spaces-polygon-clipping-portable-math-fallbacks`

### Domain: Physics spaces, bodies, shapes, and joints

Each physics server models a self-contained space containing bodies, areas, joints, and shapes, and exposes APIs to create and configure them.
Use these entity examples: PhysicsMaterial, RID.

### Domain: Polygon clipping and result trees

Clipper accepts subject and clip paths, runs a specified clip and fill rule, and builds closed or open paths or a hierarchy.
Use these entity examples: PolyPathD.

### Domain: Portable mathematical fallbacks

SDL dispatches mathematical functions to platform library functions when present or to bundled fdlibm-derived implementations otherwise.

## Unit 16

Lesson ID: `unit-procedural-noise-textures-ray-query-regex-compile-match`

### Domain: Procedural noise textures

Noise generators produce values and image data that NoiseTexture2D and NoiseTexture3D turn into textures.

### Domain: Ray query state

A ray query carries origin, direction, near and far distances, time, mask, ID, and flags for intersection or occlusion traversal.
Use these entity examples: RTCRay, RTCRayHit.

### Domain: Regular-expression compilation and matching

PCRE2 compiles patterns into code objects and matches subjects with traditional and DFA matching engines.
Use these entity examples: pcre2_code, pcre2_match_data.

## Unit 17

Lesson ID: `unit-resources-shader-language-front-end-shader-source-compilation`

### Domain: Resources and asset lifecycle

A Resource is a RefCounted asset value with a path, name, scene-local option, and loader or saver lifecycle.
Use these entity examples: Resource, Mesh, Material, GDExtension.

### Domain: Shader-language front end

The vendored glslang front end models shader types, symbols, source locations, parse context, and a typed intermediate tree.
Use these entity examples: spv::Module, spv::Function, spv::Block.

### Domain: Shader source compilation

TShader and TProgram compile shader source after preprocessed tokens into stage intermediates and program-level diagnostics.
Use these entity examples: TShader, TProgram.

## Unit 18

Lesson ID: `unit-signal-awaitability-spirv-rewriting-string-names-paths`

### Domain: Signal awaitability

A Signal combines an owner and a signal name and exposes a custom awaiter for asynchronous notification.
Use these entity examples: Variant, ManagedCallbacks.

### Domain: SPIR-V rewriting and optimization

re-spirv parses a word-addressed shader into instructions, functions, blocks, results, decorations, and specializations before optimization.
Use these entity examples: Shader, Instruction, Optimizer.

### Domain: Strings, interned names, and node paths

The string layer implements Unicode strings, hash-backed interned StringName values, path and subpath storage in NodePath, and fuzzy-search helpers.
Use these entity examples: StringName, NodePath.

## Unit 19

Lesson ID: `unit-task-scheduling-text-and-localization-texture-block-codecs`

### Domain: Task scheduling and synchronization

The internal scheduler represents work as tasks, queues, threads, and thread pools, and provides barriers and atomic helpers for synchronization.
Use these entity examples: TaskScheduler, Task, TaskQueue, ThreadPool, BarrierActive.

### Domain: Text backends and translation domains

TextServerManager registers and selects TextServer implementations, while TranslationServer stores language data in named TranslationDomain collections.
Use these entity examples: TranslationDomain, TextLine.

### Domain: Texture block codecs

etcpak exposes ETC/EAC and BC block compression plus block-decoding entry points that operate on typed source and destination buffers.

## Unit 20

Lesson ID: `unit-texture-compression-pipeline-texture-format-description-vi-native-surface-creation`

### Domain: Texture compression pipeline

Vendored Basis Universal code separates frontend block/codebook work, backend encoding, and Basis or KTX2 output creation.

### Domain: Texture format descriptions

Khronos Data Format descriptors represent image-format layout using header-word and sample-word accessors, while helper code creates, interprets, queries, and prints those descriptors.
Use these entity examples: KTX Level Index Entry.

### Domain: VI native surface creation

The `VK_NN_vi_surface` header declares input for creating a `VkSurfaceKHR` from a Nintendo VI window handle.
Use these entity examples: VkViSurfaceCreateInfoNN, VkSurfaceKHR.

## Unit 21

Lesson ID: `unit-vulkan-video-session-configuration-webrtc-peer-connections-xr-tracking`

### Domain: Vulkan video session configuration

The generated records model a video session through a video profile and standard-header-version information, with separate records for decode and encode operations.
Use these entity examples: VideoSessionCreateInfoKHR.

### Domain: WebRTC peer connections

A WebRTC peer connection exposes connection state, negotiation, ICE candidates, and data-channel creation to Godot.
Use these entity examples: WebRTCPeerConnection.

### Domain: XR tracking and poses

XRServer coordinates interfaces and tracker types, while positional, body, controller, face, and hand trackers publish XR pose state.
Use these entity examples: XRTracker, XRPose.

## Unit 22

Lesson ID: `unit-alpha-plane-coding-c-abi-manual-lifetime`

### Domain: Alpha-plane coding

The alpha plane is filtered, optionally quantized, and encoded before its bytes are incorporated with the lossy image representation.
Use these entity examples: WebPPicture, WebPConfig.

### Implementation: C ABI structures and manual lifetime

The SPIR-V reflection interface exchanges C structs and pointers through create, enumerate, query, change, and explicit destroy functions.
Apply it through: SpvReflectShaderModule, SpvReflectDescriptorSet, SpvReflectDescriptorBinding.

## Unit 23

Lesson ID: `unit-concurrency-device-runtime-python-build-hooks`

### Domain: Thread synchronization

Mutex synchronizes access to a critical section between Thread instances and is documented as a reentrant binary semaphore.

### Domain: Device runtime

A Device configures Embree execution, exposes device and thread error messages, and limits tasking resources.
Use these entity examples: RTCDevice, Device.

### Implementation: Python build hooks

Python configuration scripts define module build hooks and configure build environments.

## Unit 24

Lesson ID: `unit-geometry-families-geometry-transforms-image-codec-registration`

### Domain: Geometry families

The source defines Geometry subclasses for triangle meshes, quad meshes, curves, line segments, points, grids, subdivision meshes, user geometry, and instances.
Use these entity examples: TriangleMesh, QuadMesh, CurveGeometry, LineSegments, Points, GridMesh, SubdivMesh, UserGeometry.

### Domain: Geometry and transform values

The math layer defines 2D and 3D vectors, rectangles, boxes, planes, bases, quaternions, projections, and transforms as reusable value types.
Use these entity examples: AABB, Transform3D.

### Domain: Image codec integration

Image codec modules provide loader, saver, compressor, and decompressor implementations against engine image and resource interfaces.
Use these entity examples: Image, DDSFormatInfo.

## Unit 25

Lesson ID: `unit-localization-module-build-configuration-opentype-table-routing`

### Domain: Translation and locale selection

Translation catalogs map StringName keys to messages, TranslationDomain selects applicable catalogs, and TranslationServer manages locale, fallback, and plural-rule services.
Use these entity examples: Translation, TranslationDomain, StringName.

### Domain: SCons module build configuration

Python SCons scripts determine whether modules can build and add C++ source files, generated shader headers, or third-party sources to the build environment.

### Domain: OpenType table routing

Font subsetting routes recognized OpenType tags to typed table subsetters; color routing explicitly skips CBDT because CBLC handles it.
Use these entity examples: hb_subset_plan_t.

## Unit 26

Lesson ID: `unit-plugin-and-assembly-reload-raster-font-rendering-resource-bundle-data`

### Domain: Plugin and assembly reload

Editor plug-ins load into AssemblyLoadContext instances that resolve dependencies and may be collectible.

### Domain: Raster font rendering

Color paint records are rasterized through image, draw, paint, clipping, edge, and glyph-extents structures.
Use these entity examples: hb_raster_image_t, hb_raster_paint_t.

### Domain: Resource-bundle data

ICU resource bundles expose strings, binary values, integer vectors, keys, names, and locales over loaded resource data.
Use these entity examples: UResourceBundle, ResourceDataValue.

## Unit 27

Lesson ID: `unit-resource-loading-spatial-predictive-filters-vulkan-extensible-records`

### Domain: Resource loading and caching

ResourceLoader selects registered ResourceFormatLoader instances, applies cache modes, reports dependencies, and tracks threaded load tasks.
Use these entity examples: Resource, ResourceLoader::ThreadLoadTask.

### Domain: Spatial predictive filters

Spatial filters produce residuals from neighboring pixel-plane values using horizontal, vertical, gradient, and other predictor forms.

### Domain: Vulkan extensible records

Many Vulkan input records carry a `pNext` pointer, while the C VI surface record also carries an `sType` discriminator, forming an extensible record convention.
Use these entity examples: VkViSurfaceCreateInfoNN, GraphicsPipelineCreateInfo.

## Unit 28

Lesson ID: `unit-editor-build-composition-encoder-configuration-c-abi-versioned-initialization`

### Domain: Editor build composition

SCsub build scripts add C++ source files to the editor source set and invoke child SConscript files for nested editor subsystems.

### Domain: Encoder configuration

WebPConfig selects lossy or lossless encoding and controls quality, effort, segmentation, filtering, alpha, thread, and memory choices for the configured picture.
Use these entity examples: WebPConfig, WebPPicture.

### Implementation: C ABI-versioned initialization

Public structures are initialized through inline wrappers that pass a library ABI version to internal initialization functions before callers use the fields.
Apply it through: WebPConfig, WebPPicture.

## Unit 29

Lesson ID: `unit-image-format-loading-mesh-geometry-algorithms-c-macro-codecs`

### Domain: HDR, JPEG, and KTX loading

Separate module adapters implement an HDR ImageFormatLoader, a libjpeg-turbo ImageFormatLoader, and a KTX ResourceFormatLoader.

### Domain: Mesh geometry algorithms

The implementation uses geometry primitives to build Delaunay triangulations, convex hulls, quick hulls, and polygon triangulations.
Use these entity examples: TriangleMesh, AABB.

### Implementation: C macro-based binary decoding

C macros implement variable-length decoding by shifting accumulated values and testing continuation bits.

## Unit 30

Lesson ID: `unit-mesh-provenance-openxr-runtime-loading-c-parser-state`

### Domain: Triangle provenance

CSG output keeps triangle provenance through TriRef records and MeshRelationD triRef collections so result triangles can be related to source meshes.
Use these entity examples: TriRef, MeshRelationD.

### Domain: OpenXR runtime loading

The embedded loader reads runtime and API-layer manifests, validates `next` chains during instance creation, and creates a loader instance.
Use these entity examples: ManifestFile, LoaderInstance, XrInstanceCreateInfo.

### Implementation: C parser state machines

C structs retain XML cursor state and parser result heads while miniUPnPc parsing functions build protocol records.
Apply it through: NameValueParserData, PortMappingParserData.

## Unit 31

Lesson ID: `unit-packed-resource-files-run-management-c-preprocessor-composition`

### Domain: Packed resource files

Packed data records pack paths, offsets, sizes, hashes, encryption flags, bundle flags, and delta flags, then supplies files to the resource-loading service through PackSource.
Use these entity examples: PackedData::PackedFile.

### Domain: Running projects and instances

The run subsystem starts editor launches, exposes native-device runs, supports configurable multiple instances, and hosts embedded game-view processes.

### Implementation: C preprocessor composition

C preprocessing defines FT_MAKE_OPTION_SINGLE_OBJECT and then includes component implementation files to form a single-object library unit.

## Unit 32

Lesson ID: `unit-c-structures-and-pointers`

### Implementation: C structures and pointer handles

C structures and pointers package mutable subsystem state and expose opaque handle types.
Apply it through: FT_BZip2FileRec, PFR_FaceRec, SVG_RendererRec.

## Unit 33

Lesson ID: `unit-synchronization-primitives-cpp-native-classes`

### Domain: Threads and synchronization

The OS layer defines Thread, mutexes, condition variables, read-write locks, semaphores, spin locks, and safe binary mutexes.

### Implementation: C++ native class hierarchy

Native components use C++ classes, inheritance, and virtual functions to implement engine types such as Resource, Node, and GDScript.
Apply it through: Resource, Node, PackedScene, GDScript.

## Unit 34

Lesson ID: `unit-transform-interpolation-wayland-window-backend`

### Domain: Transform interpolation

TransformInterpolator and InterpolatedProperty calculate intermediate geometry transforms and values.
Use these entity examples: Transform3D.

### Domain: Wayland window backend

Wayland window handling builds on Linux/BSD platform integration with a dedicated thread, client protocol objects, and window state.
Use these entity examples: WaylandThread.WindowState.

## Unit 35

Lesson ID: `unit-codec-simd-specialization-c-dynamic-library-wrappers`

### Domain: Codec SIMD specialization

Conditional C preprocessing exposes AVX2, SSE4.1, and NEON codec declarations; C preprocessing is needed to explain which declarations are available.

### Implementation: C dynamic-library function wrappers

Generated C function-pointer wrappers dynamically route DBus, Fontconfig, Speech Dispatcher, and Wayland/libdecor calls for Linux/BSD desktop integration.

## Unit 36

Lesson ID: `unit-editor-and-project-settings-freetype-module-composition-c-pointers-arrays-and-strides`

### Domain: Editor and project settings

EditorSettings persists editor-side setting containers and project metadata, while dedicated dialogs expose editor, project, autoload, action-map, and input-event configuration.
Use these entity examples: EditorSettings.

### Domain: FreeType module composition

The FreeType base and autofit modules combine included implementation units into compiled engine modules; C preprocessing is needed to explain the inclusion-based composition.

### Implementation: C pointers, arrays, and strides

C aggregates are reached through pointers and stepped across image memory with width- and stride-based pointer arithmetic for rows and planes.
Apply it through: WebPPicture, WebPRescaler, WebPDecBuffer.

## Unit 37

Lesson ID: `unit-object-model-openxr-dispatch-cpp-function-pointer-interfaces`

### Domain: Engine object model

Object-derived types provide the common runtime identity and behavior base, while RefCounted supplies reference-counted lifetime for resource-like values.
Use these entity examples: Node, Resource, ClassInfo.

### Domain: OpenXR dispatch forwarding

Generated loader entry points forward API calls through a dispatch table belonging to the selected runtime.
Use these entity examples: XrGeneratedDispatchTableCore.

### Implementation: C-compatible function-pointer interfaces

C-compatible function-pointer interfaces let converter implementations select operations through a shared implementation record.
Apply it through: UConverter, UConverterSharedData.

## Unit 38

Lesson ID: `unit-openxr-runtime-integration-sdl-event-queue-cpp-object-representation-casts`

### Domain: OpenXR runtime integration

The module places OpenXRInterface over OpenXRAPI so the engine can interact with an OpenXR runtime.
Use these entity examples: OpenXRActionMap, OpenXRFutureResult.

### Domain: Event queue and watches

The event runtime stores SDL_Event values in a mutex-protected linked queue with an atomic count and event-watch support.
Use these entity examples: SDL_Event, SDL_EventEntry.

### Implementation: C++ object-representation casts

Pointer reinterpretation reads packed data through typed views, particularly when ICU validates or loads binary resource and Unicode data.
Apply it through: UResourceBundle, UCPTrie.

## Unit 39

Lesson ID: `unit-python-build-actions`

### Implementation: Python build-action functions

Python defines the icon-build action that receives target, source, and environment arguments, while SCsub scripts invoke Python-based build-environment methods.

## Unit 40

Lesson ID: `unit-collision-shapes-c-aggregate-callback-modules`

### Domain: Collision shapes

Collision shapes define the geometric representation used for ray casts, shape casts, point tests, overlap tests, triangle extraction, and collision dispatch.
Use these entity examples: Shape, ShapeSettings.

### Implementation: C aggregate state and callback modules

C aggregate types and typed function pointers define stateful processing modules, letting JPEG upsamplers retain shared buffers and dispatch a per-component routine.
Apply it through: JPEG Decompression Context, JPEG Upsampler State.

## Unit 41

Lesson ID: `unit-dynamic-values-high-level-rpc-c-resource-lifecycles`

### Domain: Dynamic values and dictionaries

Variant is the cross-cutting tagged value representation for scalar, math, object, callable, signal, dictionary, array, and packed-array values.
Use these entity examples: Variant, JSON-RPC document.

### Domain: High-level RPC routing

SceneMultiplayer routes RPC calls, peer state, connection commands, and RPC configuration.

### Implementation: C resource lifecycles and ownership

Pointer-bearing structures use explicit init, allocate, clear, free, copy, and view functions so the caller can distinguish owned image buffers from borrowed views.
Apply it through: WebPPicture, WebPMemoryWriter, WebPDecBuffer.

## Unit 42

Lesson ID: `unit-openxr-action-configuration-openxr-binding-modifiers-psa-key-lifecycle`

### Domain: OpenXR action configuration

OpenXRActionMap persists action sets, actions, interaction profiles, bindings, modifiers, and haptic feedback configuration.
Use these entity examples: OpenXRActionMap.

### Domain: Binding modifiers

Binding modifiers alter interaction-profile or per-action binding data before it is submitted to OpenXR.
Use these entity examples: OpenXRIPBinding.

### Domain: PSA key lifecycle and storage

PSA key records retain configuration-selected algorithms, type, lifetime, identifier, and policy attributes from initialization or import through slot management and optional storage.
Use these entity examples: psa_key_attributes_t.

## Unit 43

Lesson ID: `unit-sdl-runtime-properties-vulkan-render-pass-configuration-worker-parallelism`

### Domain: Runtime property groups and hints

SDL runtime property groups associate named values with an ID, while hints retrieve configuration from property and environment sources.
Use these entity examples: SDL_PropertiesID.

### Domain: Render-pass configuration

`RenderPassCreateInfo` combines attachment descriptions, subpass descriptions, and subpass dependencies into one externally exchanged rendering configuration.
Use these entity examples: RenderPassCreateInfo.

### Domain: Worker-based parallelism

The encoding implementation obtains a WebPWorkerInterface and conditionally splits analysis or lossless work when thread-level configuration permits it.
Use these entity examples: VP8Encoder.

## Unit 44

Lesson ID: `unit-collision-filtering-cpp-basis-transcoding`

### Domain: Collision filtering

Collision filtering applies object-layer, broad-phase-layer, group, body, and shape filters before candidate generation or collision queries.
Use these entity examples: Body.

### Implementation: C++ Basis transcoding behind a C-compatible API

The KTX Basis transcode implementation uses C++ references and standard-library state vectors, while its headers retain a C-compatible interface for the enclosing C API.
Apply it through: KTX2 Texture, KTX2 Private Texture State.

## Unit 45

Lesson ID: `unit-convex-collision-geometry-preprocessing-property-inspection`

### Domain: Convex support and penetration

Convex collision represents collision shapes with support functions and uses GJK closest-point and EPA penetration calculations.
Use these entity examples: Shape.

### Domain: Geometry preprocessing

Geometry preprocessing converts triangle lists to indexed geometry, builds convex hulls, and partitions triangles for spatial acceleration structures used by collision shapes.
Use these entity examples: Shape.

### Domain: Reflective property inspection

The inspector reads Object properties into EditorProperty controls and supports custom parsing, revert, copy, paste, keying, pinning, and favorites.

## Unit 46

Lesson ID: `unit-runtime-metadata-scene-replication-configuration-sdl-iostreams`

### Domain: Runtime class metadata

The runtime records classes, inheritance, methods, properties, signals, and instantiation capability for engine objects.
Use these entity examples: RID, Resource.

### Domain: Scene replication configuration

A saved SceneReplicationConfig separates NodePath properties into spawn, sync, and watch lists.
Use these entity examples: SceneReplicationConfig, ReplicationProperty.

### Domain: Abstract I/O streams

SDL I/O streams wrap byte operations for file and memory implementations and expose an SDL property group.
Use these entity examples: SDL_IOStream, SDL_PropertiesID.

## Unit 47

Lesson ID: `unit-text-shaping-plans`

### Domain: Text shaping plans

A shaping-plan key records segment properties, user features, coordinates, and shaper selection so shaping can use a plan object.
Use these entity examples: hb_shape_plan_t, hb_shape_plan_key_t.

## Unit 48

Lesson ID: `unit-reflection-cpp-class-registration`

### Domain: Class metadata and reflection

ClassDB records Object-derived classes, properties, methods, signals, and constants for runtime lookup.
Use these entity examples: ClassInfo, Variant.

### Implementation: C++ runtime class registration

C++ templates and classes implement the runtime registry that records native class identity, inheritance, construction, methods, properties, signals, and constants.
Apply it through: Resource, RID.

## Unit 49

Lesson ID: `unit-replicated-spawning`

### Domain: Replicated spawning

The spawner tracks spawnable scenes and sends configured spawn property lists while creating or removing nodes.
Use these entity examples: SceneReplicationConfig.

## Unit 50

Lesson ID: `unit-godot-member-exposure-manifold-mesh-data-cpp-classes-and-ref-handles`

### Domain: Godot member exposure

Generator output registers exported members, signals, and RPC metadata as engine-visible method and property descriptions.
Use these entity examples: Variant.

### Domain: Manifold mesh interchange

MeshGLP carries indexed triangles, per-vertex properties, merge information, and run metadata into Manifold construction.
Use these entity examples: MeshGLP.

### Implementation: C++ inheritance and engine reference handles

C++ test infrastructure uses public inheritance for test doubles and stores engine objects in Ref<T> handles created with memnew.
Apply it through: Animation, AnimationTrack.

## Unit 51

Lesson ID: `unit-scripting`

### Domain: Script resources and instances

GDScript is a Script Resource whose compiled members and functions supply script instances to Object-derived engine objects.
Use these entity examples: GDScript, Node.

## Unit 52

Lesson ID: `unit-filesystem-project-index-fixture-contracts-cplusplus-enumerated-export-state`

### Domain: Project filesystem index

The filesystem service scans project directories into a tree of file records with resource type, UID, import state, dependency list, and script-class metadata.
Use these entity examples: EditorFileSystemDirectory, EditorFileSystemDirectory::FileInfo.

### Domain: Fixture contracts

The suite stores GDScript source alongside expected outcome files and test assets, so behavior is expressed as a checked fixture contract.
Use these entity examples: Test Script Fixture, Expected Result Fixture.

### Implementation: C++ enumerated export state

C++ enumerations classify export messages, export filters, and file or script export modes.
Apply it through: EditorExportPreset, EditorExportPlatform::ExportMessage.

## Unit 53

Lesson ID: `unit-script-registration-metadata-cpp-binary-data-codecs`

### Domain: Managed script registration metadata

C# script classes receive generated script-path and assembly script-type metadata for engine discovery.
Use these entity examples: ScriptPathAttribute, AssemblyHasScriptsAttribute.

### Implementation: C++ binary data codecs

glTF decoding reads typed C++ containers of raw bytes through buffer-view offsets and strides, while encoding writes byte arrays into newly allocated buffer views.
Apply it through: GLTFBufferView, GLTFAccessor.

## Unit 54

Lesson ID: `unit-unicode-data-cpp-class-state-and-composition`

### Domain: Unicode data and properties

The included code stores generated Unicode script, emoji, normalization, bidi, case, and general-character-property data for lookup-oriented services.
Use these entity examples: UCPTrie, hb_unicode_funcs_t.

### Implementation: C++ class state and composition

C++ classes group motion settings, calibration, vector, quaternion, and motion-update behavior behind named types.

## Unit 55

Lesson ID: `unit-cpp-classes-cpp-classes-and-inheritance`

### Implementation: C++ classes and inheritance

C++ classes define reusable state and behavior, and the implementation derives UnicodeString from Replaceable and many ICU services from UObject or UMemory.
Apply it through: Locale, AABBTreeBuilder.

### Implementation: C++ classes and inheritance

The editor source declares engine services as C++ classes derived from Object, Resource, Control, and other engine base classes.
Apply it through: EditorSettings, VCS Diff File.

## Unit 56

Lesson ID: `unit-cpp-exception-abi-boundaries`

### Implementation: C++ exception containment at ABI boundaries

OpenXR loader macros use C++ exceptions internally but convert `std::bad_alloc` and other `std::exception` failures into OpenXR result codes.
Apply it through: LoaderInstance.

## Unit 57

Lesson ID: `unit-cpp-polymorphic-backends`

### Implementation: C++ polymorphic backend classes

C++ backend classes inherit engine interfaces so a rendering or platform service can be selected through a common base type.
Apply it through: RenderingDeviceDriverVulkan::BufferInfo.

## Unit 58

Lesson ID: `unit-cpp-resource-and-polymorphism`

### Implementation: C++ resources and polymorphic engine objects

C++ engine modules define polymorphic Resource and server objects that own long-lived configuration and runtime state.
Apply it through: SceneReplicationConfig, OggPacketSequence, OpenXRActionMap.

## Unit 59

Lesson ID: `unit-cpp-runtime-adapters-cpp-scope-locking`

### Implementation: C++ inheritance and reference-counted adapters

C++ defines reference-counted Java adapter classes and platform subclasses that specialize common engine interfaces.

### Implementation: C++ scope-bound locking

The Jolt contact listener creates scope-bound C++ lock objects before mutating shared contact, manifold, and debug-contact collections.
Apply it through: JoltSpace3D.

## Unit 60

Lesson ID: `unit-cpp-typed-api-records`

### Implementation: C++ typed API records

Generated C++ structs use typed pointers and brace defaults to represent Vulkan creation inputs and optional state records.
Apply it through: GraphicsPipelineCreateInfo, PresentInfoKHR.

## Unit 61

Lesson ID: `unit-cxx-atomics`

### Implementation: C++ atomic synchronization

C++ atomic counters and compare-exchange loops coordinate active barriers and conditional minimum or maximum updates across workers.
Apply it through: atomic, BarrierActive, BarrierActiveAutoReset.

## Unit 62

Lesson ID: `unit-cxx-c-abi`

### Implementation: C++ linkage and opaque API handles

C++ exposes RTCDevice, RTCScene, RTCGeometry, and RTCBuffer as opaque pointer handles through a C-linkage public boundary.
Apply it through: RTCDevice, RTCScene, RTCGeometry, RTCBuffer.

## Unit 63

Lesson ID: `unit-cxx-conditional-compilation`

### Implementation: C++ preprocessor configuration

C++ preprocessing selects tasking backends, supported geometry features, ISA namespaces, platform headers, and fallback shims before compilation.
Apply it through: TaskScheduler, Geometry, vint8.

## Unit 64

Lesson ID: `unit-character-encoding-conversion-completion-contexts-c-abi-bridging`

### Domain: Character-encoding conversion

ICU models converters with shared static data, an implementation dispatch structure, contexts, and UTF-specific implementations.
Use these entity examples: UConverter, UConverterSharedData.

### Domain: Contextual completion contracts

Completion fixture contracts pair a cursor-bearing script with configuration rules that include or exclude candidates.
Use these entity examples: Completion Test Configuration, Test Script Fixture.

### Implementation: C ABI bridging

OpenXR headers use macros and `extern "C"` guards so C++ callers expose C-compatible API declarations and function-pointer types.
Apply it through: XrInstanceCreateInfo.

## Unit 65

Lesson ID: `unit-diagnostic-expectations-c-abi-data-structures`

### Domain: Diagnostic expectations

Diagnostic fixture contracts preserve parser errors, analyzer warnings, and runtime errors as expected textual results.
Use these entity examples: Expected Result Fixture.

### Implementation: C ABI structures and opaque state

C declarations use tagged structures, pointer members, and opaque implementation pointers to express caller-owned descriptors and library-managed state across API boundaries.
Apply it through: VkWaylandSurfaceCreateInfoKHR, z_stream.

## Unit 66

Lesson ID: `unit-export-presets-c-structs-pointers`

### Domain: Export presets

An export preset selects indexed project files, applies include, exclude, and customized-file rules, and records output path, features, patches, and encryption options.
Use these entity examples: EditorExportPreset.

### Implementation: C structs and pointer-linked state

C structs and pointers represent explicit runtime state, ownership links, and opaque-handle implementation data.
Apply it through: SDL_EventEntry, SDL_hid_device_info.

## Unit 67

Lesson ID: `unit-identifier-spoof-checking-lsp-semantic-fixtures-cpp-classes-inheritance`

### Domain: Identifier spoof checking

Unicode property checks and configured allowed character or locale sets are held by SpoofImpl and exposed through USpoofChecker APIs.
Use these entity examples: SpoofImpl, USpoofChecker.

### Domain: Language-server semantic fixtures

Language-server fixture contracts provide nested declarations, local scopes, documentation comments, and lambdas for semantic-editor tests.
Use these entity examples: Test Script Fixture.

### Implementation: C++ classes and inheritance

The editor is structured from C++ classes that derive from engine base types such as Node, Container, ScrollContainer, ResourceImporter, and RefCounted.

## Unit 68

Lesson ID: `unit-unicode-normalization-cpp-flag-stringification`

### Domain: Unicode normalization

ICU implements Normalizer2 variants and loads normalization data into tries, extra mappings, and composition data.
Use these entity examples: Normalizer2Impl, UCPTrie.

### Implementation: C++ flag stringification

The helper converts typed `FormatFeatureFlags` values into a string by testing individual flag bits and appending their names.

## Unit 69

Lesson ID: `unit-cpp-numeric-value-operations`

### Implementation: C++ numeric value operations

C++ member-based value types calculate vector lengths, normalized directions, quaternion orientation, and time-scaled motion updates.

## Unit 70

Lesson ID: `unit-cpp-overloading-cpp-plugin-specialization`

### Implementation: C++ member-function overloading

Overloaded member functions provide related operations for different argument forms, such as LocaleMatcher matching a single locale, an iterator, or a list string.
Apply it through: LocaleMatcher, Locale.

### Implementation: C++ plugin specialization

C++ classes specialize EditorPlugin for feature-specific integrations; this requires C++ classes as the base abstraction mechanism.

## Unit 71

Lesson ID: `unit-dynamic-variant-export-orchestration-c-concurrent-state`

### Domain: Dynamic runtime values

Variant is the tagged runtime-value representation used for construction, conversion, operators, calls, indexing, and member access.
Use these entity examples: Variant.

### Domain: Export orchestration

Export orchestration owns export presets, target platforms, and plugins, and maps a target platform to its runnable preset.
Use these entity examples: EditorExport, EditorExportPreset.

### Implementation: C-managed concurrent state

C structs hold mutexes, atomic counts, and linked entries when SDL manages concurrent event and device state.
Apply it through: SDL_EventEntry.

## Unit 72

Lesson ID: `unit-gdscript-editor-services-c-explicit-resource-lifecycles`

### Domain: GDScript editor services

Editor tooling consumes GDScript parser output to generate documentation, color syntax, extract translations, and provide completion-related behavior.
Use these entity examples: GDScriptParser::Node, GDScriptTokenizer::Token.

### Implementation: C explicit resource lifecycles

C functions establish, reset, and release C struct state explicitly through paired lifecycle operations.
Apply it through: FT_Stream, FT_BZip2FileRec.

## Unit 73

Lesson ID: `unit-image-quality-metrics-project-upgrade-c-function-pointer-tables`

### Domain: Image quality metrics

PSNR, SSIM, local similarity, and squared-error routines assess encoded or reconstructed image planes.
Use these entity examples: WebPAuxStats, WebPPicture.

### Domain: Project source migration

ProjectConverter3To4 retains whether each source line is a comment while applying named rename, conversion, and check-only passes to project source.
Use these entity examples: SourceLine.

### Implementation: C function-pointer tables

C function-pointer tables attach implementation callbacks to C struct state for dynamically selected system and device services.
Apply it through: SDL_hid_device_info.

## Unit 74

Lesson ID: `unit-scene-contexts-c-stateful-struct-apis`

### Domain: Scene-aware test context

A scene input gives a completion context containing nodes, attached scripts, resources, and unique names.
Use these entity examples: Scene Fixture, Completion Test Configuration.

### Implementation: C stateful struct APIs

C APIs expose mutable state through struct pointers, initialize it explicitly, and communicate failures through status-returning functions and caller-owned buffers.
Apply it through: psa_key_attributes_t, mbedtls_ssl_session.

## Unit 75

Lesson ID: `unit-yuv-rgb-conversion-cpp-copy-on-write-containers`

### Domain: YUV/RGB conversion and chroma processing

Conversion kernels turn RGB or ARGB input into YUV planes, reconstruct RGB outputs, and upsample chroma with neighboring samples.
Use these entity examples: WebPPicture, WebPDecBuffer.

### Implementation: C++ copy-on-write container storage

The container layer uses C++ container classes to separate CowData buffer allocation and copying from Vector-style value interfaces.
Apply it through: Array, Dictionary.

## Unit 76

Lesson ID: `unit-primitive-references-project-settings-c-stateful-streaming-api`

### Domain: Primitive references

A PrimRef supplies bounds plus geometry and primitive identifiers so builders can partition scene primitives independently of their source geometry layout.
Use these entity examples: PrimRef, PrimRefMB, SubGridBuildData.

### Domain: Project settings and feature overrides

ProjectSettings stores named Variant values with persistence metadata, feature overrides, autoload definitions, resource paths, and a change version.
Use these entity examples: ProjectSettings, ProjectSettings::VariantContainer.

### Implementation: C stateful streaming APIs

C-facing code represents decoder progress as pointer-rich state and accepts caller-supplied allocation callbacks.
Apply it through: BrotliDecoderState.

## Unit 77

Lesson ID: `unit-shader-editing-and-language-plugins-variant-containers-callables-and-lambdas`

### Domain: Shader editing and language plugins

Shader editing uses a ShaderEditorPlugin, a text shader editor, syntax highlighting, shader-language plugins, creation dialogs, and shader-file editing.

### Domain: Dynamic arrays and dictionaries

Array and Dictionary store Variant values, while typed validators and typed wrappers constrain their element, key, or value types.
Use these entity examples: Array, Dictionary, Variant.

### Implementation: Callables and lambdas

Callable values can reference functions, constructors, built-ins, and lambdas; lambda bodies capture their enclosing context.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 78

Lesson ID: `unit-cpp-godot-binding-macros`

### Implementation: C++ engine binding macros

Godot's C++ classes declare engine binding and script-overridable callbacks through macros such as GDCLASS and GDVIRTUAL.

## Unit 79

Lesson ID: `unit-cpp-interface-polymorphism`

### Implementation: C++ interface polymorphism

C++ base classes and virtual callbacks define extension contracts whose subclasses provide physics, rendering, stream, and text implementations.
Apply it through: RenderSceneBuffersConfiguration, RID.

## Unit 80

Lesson ID: `unit-javascript-web-runtime`

### Implementation: JavaScript browser runtime libraries

JavaScript libraries implement engine startup, preloading, runtime heap access, browser display and input callbacks, audio worklets, fetch, MIDI, and JavaScript object proxies.
Apply it through: JavaScriptObjectImpl.

## Unit 81

Lesson ID: `unit-nodepaths-and-indexed-access`

### Implementation: Node paths and indexed access

The fixtures use `$`, `%`, and indexed self access to resolve Node members inside a scene-aware test context.
Apply it through: Scene Fixture, Completion Test Configuration.

## Unit 82

Lesson ID: `unit-break-iteration-broad-phase-classes-inheritance-and-lookup`

### Domain: Text break iteration

ICU builds and executes rule-based break iterators, including dictionary-backed language break engines and compiled rule tables.
Use these entity examples: RuleBasedBreakIterator, LanguageBreakEngine.

### Domain: Broad-phase collision detection

Broad-phase collision detection uses Body world-space AABox bounds and collision filtering to generate candidate body pairs and answer coarse queries.
Use these entity examples: Body.

### Implementation: Classes, inheritance, and lookup

Classes supply nested types, inheritance, `super`, and lexical member lookup; preloaded classes also participate in lookup.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 83

Lesson ID: `unit-browser-runtime-adapters-bvh-construction-cpp-runtime-polymorphism`

### Domain: Browser runtime adapters

Web audio, display, input, fetch, MIDI, filesystem, and WebGL adapters call JavaScript libraries and exchange data with C++ Web platform classes.
Use these entity examples: JavaScriptObjectImpl.

### Domain: BVH construction

BVH construction partitions PrimRef records into nodes and leaves through configurable builder callbacks and settings.
Use these entity examples: BVHN, GeneralBVHBuilder, BVHNodeRecord.

### Implementation: C++ virtual interfaces

Godot uses polymorphic base interfaces so importers, inspector plug-ins, preview generators, and editor plug-ins can supply type-specific behavior through virtual methods.

## Unit 84

Lesson ID: `unit-enet-transport-and-multiplayer-extensions-cpp-virtual-dispatch`

### Domain: ENet transport and multiplayer adaptation

ENetConnection and ENetPacketPeer wrap connection and peer behavior, while ENetMultiplayerPeer adapts the transport to MultiplayerPeer.
Use these entity examples: ENetConnection, ENetPacketPeer, ENetMultiplayerPeer.

### Domain: Virtual implementation extensions

Extension classes declare required or optional virtual callbacks so external implementations can replace physics, rendering, text, and scripting behavior.
Use these entity examples: RenderSceneBuffersConfiguration, RID.

### Implementation: C++ virtual dispatch

Virtual dispatch uses derived classes to substitute behavior, as locale iterators override next and ICU break engines derive from a common engine type.
Apply it through: LocaleMatcher, RuleBasedBreakIterator.

## Unit 85

Lesson ID: `unit-gdscript-language-server-vulkan-swapchain-presentation-javascript-browser-ffi`

### Domain: GDScript language-server support

The language server reuses parser-derived symbol data to manage documents, resolve symbols, publish diagnostics, provide links, and build workspace edits.
Use these entity examples: LSP::TextDocumentItem, LSP::DocumentSymbol, GDScriptWorkspace.

### Domain: Swapchain presentation

The binding pairs swapchain configuration with a presentation request whose wait semaphores, swapchains, and image indices are array inputs.
Use these entity examples: SwapchainCreateInfoKHR, PresentInfoKHR, SwapchainKHR.

### Implementation: JavaScript browser FFI

JavaScript bridge libraries marshal browser objects, strings, buffers, and callbacks between web APIs and C++ platform code.
Apply it through: WebRTCPeerConnection, WebRTCDataChannel, WebXRInterfaceJS.

## Unit 86

Lesson ID: `unit-signals-and-await`

### Implementation: Signals and await

Signal values and callable functions support waiting for emissions or coroutine completion with `await`.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 87

Lesson ID: `unit-bvh-traversal-cpp-abstraction-and-polymorphism`

### Domain: BVH traversal

Traversal tests the committed BVH with a ray query, orders candidate nodes, and calls primitive intersectors for reached leaves.
Use these entity examples: BVHNIntersector1, BVHNIntersectorKHybrid, TravRay.

### Implementation: C++ types, encapsulation, and inheritance

This C++ code models decoders, geometry engines, and allocators as named types with inheritance and encapsulated state.
Apply it through: Ktx2TranscoderState, PolyPathD.

## Unit 88

Lesson ID: `unit-enet-host-peer-transport-extension-api-compatibility-extension-structure-chains`

### Domain: ENet host and peer transport

ENet creates a host with a bounded peer allocation and manages peer state changes when connecting or disconnecting.
Use these entity examples: ENetHost, ENetPeer.

### Domain: Extension API compatibility baselines

Version-pair directories retain expected extension-API validation differences, and the validator aggregates their relevant diagnostic lines.

### Domain: Vulkan extension structure chains

The structs expose `sType` and `pNext` fields so extension records can be carried through Vulkan creation and query calls.
Use these entity examples: VkMetalSurfaceCreateInfoEXT, VkExternalFormatQNX.

## Unit 89

Lesson ID: `unit-motion-blur-narrow-phase-resource-formats`

### Domain: Motion blur

Motion-blur Geometry and BVH nodes interpolate bounds at ray-query time.
Use these entity examples: PrimRefMB, AABBNodeMB_t, AABBNodeMB4D_t.

### Domain: Narrow-phase collision queries

Narrow-phase collision queries test collision shapes for rays, points, overlaps, and casts after broad-phase candidates have been selected.
Use these entity examples: Shape, Body.

### Domain: Resource formats and serialization

The resource-loading service uses binary, JSON, image, crypto, and extension resource-format loaders and savers.
Use these entity examples: Resource.

## Unit 90

Lesson ID: `unit-scene-commit`

### Domain: Scene construction and commit

A Scene is created from a Device, retains Geometry instances, and commits a BVH used by query calls.
Use these entity examples: RTCScene, Scene.

## Unit 91

Lesson ID: `unit-contact-management-cpp-byte-exact-binary-parsing`

### Domain: Contact manifolds and warm-starting

Contact management converts narrow-phase collision results into contact constraints and caches manifolds, body pairs, and contact points between updates.
Use these entity examples: CachedManifold, Body.

### Implementation: C++ byte-exact binary parsing

This C++ code reads byte streams through packed fields, pointer casts, and bit readers; this is necessary to map serialized texture headers to in-memory values.
Apply it through: BasisFileHeader, BasisSliceDescriptor, Ktx2Header, Ktx2LevelIndex.

## Unit 92

Lesson ID: `unit-editor-authoring-workspaces-enet-wire-commands-cpp-templates-and-generic-containers`

### Domain: Editor authoring workspaces

The editor implements dock management, filesystem browsing, scene-tree editing, animation editing, asset-library browsing, audio editing, and debugger views as specialized controls and plugins.

### Domain: ENet wire commands

ENet wire commands select fixed protocol layout sizes and drive peer state transitions; host-peer transport is required to interpret the peer state.
Use these entity examples: ENetPeer.

### Implementation: C++ templates and generic containers

This C++ code parameterizes collection and math utilities by type, so their typed operations can be reused.
Apply it through: Ktx2LevelIndex, PolyPathD.

## Unit 93

Lesson ID: `unit-explicit-drm-syncobj-gdextension-libraries-cxx-class-hierarchy`

### Domain: Explicit DRM synchronization objects

Wayland protocol objects import DRM synchronization timelines and attach acquire and release timeline points to a surface commit.
Use these entity examples: wp_linux_drm_syncobj_timeline_v1.

### Domain: GDExtension libraries

GDExtension is a Resource that holds a loader, registered extension classes, initialization state, and library open, initialize, deinitialize, and close operations.
Use these entity examples: GDExtension.

### Implementation: C++ class hierarchies and reference bases

C++ class inheritance connects shared reference ownership to acceleration, scene, geometry, builder, and scheduler abstractions.
Apply it through: RefCount, AccelData, Geometry, Builder.

## Unit 94

Lesson ID: `unit-gltf-node-hierarchy-instancing-openxr-composition-layers`

### Domain: glTF scene-node hierarchy

GLTFDocument reconstructs parent and child relationships among indexed GLTF node records and attaches meshes, cameras, lights, skins, and skeleton membership through node indexes.
Use these entity examples: GLTFNode, GLTFState.

### Domain: Instancing

Instance and InstanceArray are Geometry types attached to a Scene that contribute transformed primitive references and maintain instance-query context state.
Use these entity examples: Instance, InstanceArray, InstancePrimitive.

### Domain: OpenXR composition layers

Composition-layer scene nodes and an extension submit specialized layer descriptions to the OpenXR runtime.

## Unit 95

Lesson ID: `unit-openxr-render-models-property-introspection-replicated-property-synchronization`

### Domain: OpenXR render models

The render-model extension tracks runtime render models, and scene nodes display models or manage their activation.
Use these entity examples: OpenXRRenderModelData.

### Domain: Property-list and scene-property helpers

PropertyListHelper resolves slash-delimited property names to property records, while PropertyUtils compares properties and walks scene-state pack data.

### Domain: Replicated property synchronization

Synchronizers observe configured properties and forward sync property lists to scene replication.
Use these entity examples: SceneReplicationConfig.

## Unit 96

Lesson ID: `unit-resource-identifiers-runtime-configuration-runtime-loop-time`

### Domain: Resource and server identifiers

ResourceUID maps project resource identifiers to paths, while an RID is an opaque session-only handle for a low-level server resource.
Use these entity examples: ResourceUID, RID.

### Domain: Runtime configuration

Runtime configuration reads project settings such as the main scene and boot-splash options; the application project file also declares its main scene.

### Domain: Main loop, OS, and time services

OS hosts platform runtime services, MainLoop defines the loop-facing object type, and Time exposes date and time behavior.

## Unit 97

Lesson ID: `unit-scene-authoring-tile-authoring`

### Domain: 2D and 3D scene authoring

Scene authoring is split between CanvasItemEditor for 2D work and Node3DEditor with Node3DEditorViewport for 3D work.

### Domain: Tile authoring

Tile tools edit atlas sources, per-tile data, terrain data, proxies, map layers, and scene-collection sources.

## Unit 98

Lesson ID: `unit-astc-block-coding-cpu-specialized-dsp-annotations-static-state-and-lifecycle`

### Domain: ASTC block coding

ASTC helpers represent physical and logical blocks with endpoint ranges, weight grids, partitions, and bit-level packing.

### Domain: CPU-specialized DSP backends

DSP function families have scalar, SSE2/SSE4.1/AVX2, NEON, MSA, and MIPS implementations that are compile-time selected for target capabilities.

### Implementation: Annotations, static state, and lifecycle

Annotations modify class declarations and members, while static state and `@onready` participate in initialization and lifecycle behavior.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 99

Lesson ID: `unit-engine-bootstrap-godot-enet-socket-adaptation-cpp-c-linkage-adapters`

### Domain: Engine bootstrap

Engine bootstrap consumes runtime configuration to initialize execution and select the configured main scene.

### Domain: Godot ENet socket adaptation

The Godot ENet socket adapter supplies UDP and DTLS socket classes to the transport; ENet wire commands provide this adapter's transport-facing purpose.

### Implementation: C++ C-linkage adapters

C++ classes are wrapped behind C linkage functions and derived opaque API types, so C callers use handles rather than C++ class declarations.
Apply it through: Face, Font, Segment, Slot.

## Unit 100

Lesson ID: `unit-hit-results-input-action-configuration-cpp-object-hierarchies`

### Domain: Intersection results

An RTCRayHit combines a ray query with geometric-normal, barycentric-coordinate, primitive-ID, geometry-ID, and instance-stack hit results.
Use these entity examples: RTCHit, RTCRayHit.

### Domain: Input action and shortcut configuration

Input action configuration depends on editor and project settings because ActionMapEditor, event search, event capture, and input-event configuration edit action-event mappings.

### Implementation: C++ class inheritance

C++ class declarations define the repository's object families through public base classes, including Node-derived runtime objects, Resource-derived assets, and Viewport-derived windows.
Apply it through: Node, PackedScene.

## Unit 101

Lesson ID: `unit-project-catalog-rendering-device-resources-cpp-scoped-locks`

### Domain: Project catalog and selection

ProjectList models known projects with name, path, tags, main scene, version, unsupported features, last-edit time, and favorite state.
Use these entity examples: ProjectCatalog, ProjectCatalogItem.

### Domain: RenderingDevice descriptors and handles

RenderingDevice creates and consumes RID handles for buffers, textures, shaders, uniforms, pipelines, framebuffers, and acceleration structures.
Use these entity examples: RDTextureFormat, RDUniform, RDAccelerationStructureGeometry.

### Implementation: C++ scoped lock guards

The implementation uses a scoped C++ MutexLock object and class-specific lock helpers around shared runtime state.

## Unit 102

Lesson ID: `unit-resource-dependency-resolution-rigid-body-motion-cpp-templates-and-views`

### Domain: Resource dependency resolution

Dependency resolution uses file-index entries and ResourceUID mappings to expose each indexed resource's current dependency paths.
Use these entity examples: EditorFileSystemDirectory::FileInfo.

### Domain: Rigid-body motion

Rigid-body motion uses Body transform state plus MotionProperties to model static, kinematic, and dynamic behavior, permitted degrees of freedom, mass, inertia, and velocities.
Use these entity examples: Body.

### Implementation: C++ templates and non-owning views

Template-based containers and views expose typed data without copying it, including VectorView access to command and resource arrays.
Apply it through: RecordedCommand.

## Unit 103

Lesson ID: `unit-scene-tree-and-signal-connections-tile-libraries-cxx-enums-bitflags`

### Domain: Scene tree and signal connections

Scene authoring is accompanied by a SceneTreeEditor and a ConnectionsDock that inspect nodes, signals, methods, connection arguments, and bound values.

### Domain: Tile libraries, cells, and patterns

A TileSet Resource owns tile sources, identifies tiles by source, atlas-coordinate, and alternative IDs, and supplies tile data to TileMapLayer cells and TileMapPattern copies.
Use these entity examples: TileSet, TileData.

### Implementation: Scoped enums and bit flags

Scoped enum class values model motion, body, collision, update-error, and configuration states, while selected enums provide bitwise combinations.
Apply it through: Body, ConstraintSettings.

## Unit 104

Lesson ID: `unit-vulkan-instance-setup-wraparound-network-time`

### Domain: Vulkan instance setup

The binding models instance setup as an `InstanceCreateInfo` record containing optional application metadata plus enabled layer and extension name arrays.
Use these entity examples: InstanceCreateInfo, ApplicationInfo.

### Domain: Wraparound network time

ENet compares and subtracts time values with an overflow threshold so ordering and differences remain defined across its configured time wraparound.

## Unit 105

Lesson ID: `unit-android-platform-host-constraints-cpp-runtime-casts`

### Domain: Android platform host

Android platform hosting starts and manages the native engine from Android activity and library code, continuing the engine bootstrap on the mobile host.

### Domain: Constraint solving

Constraint solving applies impulses to rigid-body motion and reuses contact manifolds while supporting point, distance, hinge, slider, fixed, gear, pulley, cone, six-degree-of-freedom, path, and swing-twist constraints.
Use these entity examples: ConstraintSettings, TwoBodyConstraint.

### Implementation: C++ static casts across track types

The implementation uses static_cast to interpret a base Track pointer as a derived Track object after a track kind has been selected.
Apply it through: Animation.

## Unit 106

Lesson ID: `unit-frame-timing-packet-queries-cpp-static-generated-data`

### Domain: Frame timing and physics stepping

Frame timing calculates a process delta and bounded physics-step count after engine bootstrap determines the active physics tick configuration.
Use these entity examples: MainFrameTime.

### Domain: Ray packets and streams

Packet execution packs several ray queries and their hit results into 4-, 8-, 16-, or N-wide layouts, with array-of-structures and structure-of-arrays stream readers.
Use these entity examples: RTCRay4, RTCRayHit16, RTCRayNt, RayStreamSOA.

### Implementation: C++ static generated data

Static const arrays and values embed generated Unicode and normalization data directly in translation units.
Apply it through: UCPTrie.

## Unit 107

Lesson ID: `unit-performance-monitors-primitive-intersection-cpp-templates`

### Domain: Performance monitors

Performance monitors collect engine counters and execute registered monitor calls after engine bootstrap.

### Domain: Primitive intersection

Primitive intersectors test a ray query against Geometry representations and update RTCRayHit through common intersection or occlusion epilogues.
Use these entity examples: TriangleMIntersectorKMoeller, QuadMIntersector1MoellerTrumbore, ConeCurveIntersector1.

### Implementation: C++ templates

Templates parameterize data structures and algorithms, including string code-point appenders and typed OpenType table subset calls.
Apply it through: UCPTrie, hb_subset_plan_t.

## Unit 108

Lesson ID: `unit-resource-assets-scene-data-and-buffers-cxx-object-model`

### Domain: Resource-backed assets

Resource-derived classes represent reusable scene content such as textures, meshes, materials, themes, animations, shapes, and packed scenes.
Use these entity examples: PackedScene, Animation, TileMapPattern.

### Domain: Scene render data and buffers

RenderDataRD gathers visible scene-instance pointers and RID lists for lights, probes, decals, lightmaps, and fog volumes, while RenderSceneBuffersRD supplies named GPU scene textures.
Use these entity examples: RenderDataRD.

### Implementation: C++ classes, inheritance, and polymorphic interfaces

C++ classes and inheritance define extensible interfaces for shapes, constraints, collision queries, job systems, renderers, and body access.
Apply it through: Shape, Body.

## Unit 109

Lesson ID: `unit-scene-tree-shader-programs-properties-and-accessors`

### Domain: Scene graph and lifecycle

A scene tree arranges Object-derived Node instances into a parent-child hierarchy that SceneTree owns and updates.
Use these entity examples: Node, SceneState.

### Domain: Shader programs and material parameters

A Shader resource supplies custom shader source, ShaderInclude supplies reusable source fragments, and ShaderMaterial binds a Shader with parameter values.
Use these entity examples: ShaderMaterial, RDPipelineSpecializationConstant.

### Implementation: Properties and accessors

Class properties route reads and writes through inline or named accessors, including static and chained access.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 110

Lesson ID: `unit-tls-crypto-backend-viewport-render-data`

### Domain: Mbed TLS-backed crypto and transport

The Mbed TLS module supplies Godot Crypto, certificate, TLS context, DTLS server, and TLS peer implementations.

### Domain: Viewport render-frame data

RenderData and RenderSceneData exist for a viewport frame, while RenderSceneBuffersConfiguration configures buffers recreated when a viewport changes.
Use these entity examples: RenderSceneBuffersConfiguration, RenderData.

## Unit 111

Lesson ID: `unit-audio-bus-processing-canvas-and-viewports-cpp-frame-scheduling`

### Domain: Audio buses, streams, and effects

AudioServer owns buses with channels and effects, while Resource-derived streams, effects, and bus layouts define reusable audio configuration.
Use these entity examples: AudioBusLayout.

### Domain: Canvas and viewport presentation

CanvasItem, CanvasLayer, Viewport, and Window extend the Node tree with visual presentation and window-facing state.
Use these entity examples: Node.

### Implementation: C++ frame-time arithmetic

C++ frame timing turns accumulated time and a physics tick rate into bounded physics-step counts for frame timing.
Apply it through: MainFrameTime.

## Unit 112

Lesson ID: `unit-expression-evaluation-cpp-inheritance-interfaces`

### Domain: Expression parsing and evaluation

Expression defines token and expression-node types and evaluates expression nodes against Variant inputs and built-in functions.
Use these entity examples: Variant.

### Implementation: C++ inheritance interfaces

C++ classes specialize engine and OpenXR base interfaces through public inheritance.
Apply it through: OpenXRStructureBase, ShapedTextDataAdvanced.

## Unit 113

Lesson ID: `unit-filter-callbacks-csharp-partial-source-generation`

### Domain: Intersection and occlusion filters

Geometry-installed and query-context filters process candidate hit results before they update an RTCRayHit or report occlusion.
Use these entity examples: RTCFilterFunctionNArguments, RTCIntersectFunctionNArguments, RTCOccludedFunctionNArguments.

### Implementation: C# partial types and source-generator tests

The .NET SDK test projects define partial C# types and verify generators that emit Godot-facing signal, property, method, serialization, and script-path code.
Apply it through: CSharpScript.

## Unit 114

Lesson ID: `unit-input-actions-mp3-audio-resources-kotlin-runtime-coordination`

### Domain: Input events and actions

InputMap associates named actions with InputEvent instances, and nodes receive input-event callbacks.
Use these entity examples: InputEvent.

### Domain: MP3 audio resources

The MP3 module loads MP3 stream data, supplies resampled playback, and imports MP3 assets.

### Implementation: Kotlin runtime and service coordination

Kotlin coordinates the Android platform host through nullable runtime state, Android host construction, and integer-coded remote-service messages.
Apply it through: Callable, GodotService.EngineStatus.

## Unit 115

Lesson ID: `unit-network-data-exchange-objectdb-snapshots-ogg-packet-sequences`

### Domain: Packets, HTTP, JSON, and JSON-RPC

Packet peers transmit raw bytes or Variant values, while JSON and JSON-RPC map values into external message documents.
Use these entity examples: JSON-RPC document, Variant.

### Domain: ObjectDB snapshots

The debug-only object database profiler collects serialized object snapshots and renders summary, class, node, object, and ref-counted views.

### Domain: Ogg packet sequences

OggPacketSequence persists packet data, page granule positions, and sample-rate information for sequence playback.
Use these entity examples: OggPacketSequence.

## Unit 116

Lesson ID: `unit-ogg-pages-and-packets-ogg-vorbis-audio-streams-packed-scenes`

### Domain: Ogg pages, packets, and bit packing

The Ogg implementation packs variable-sized words into octet streams and maintains stream state that frames packets into page headers and bodies.
Use these entity examples: Ogg Stream State, Ogg Packet.

### Domain: Ogg Vorbis audio streams

AudioStreamOggVorbis and its playback class provide an audio-stream resource and resampled playback implementation.
Use these entity examples: AudioStreamOggVorbis, AudioStreamPlaybackOggVorbis.

### Domain: Packed scene serialization

A PackedScene stores a resource-backed SceneState whose node entries, property values, and connection entries reconstruct a node hierarchy.
Use these entity examples: PackedScene, SceneState.

## Unit 117

Lesson ID: `unit-physics-collision-ray-primitive-intersection-raycast-occlusion-culling`

### Domain: Physics shapes, objects, and collision reports

2D and 3D CollisionObject nodes own Shape resources through shape owners; body movement returns KinematicCollision results.
Use these entity examples: Shape2D, CollisionShape2D, KinematicCollision2D.

### Domain: Ray–primitive intersection

Ray-primitive intersection evaluates rays against triangle, sphere, and rounded-line geometry and forwards valid hit candidates to an epilog.
Use these entity examples: SubGrid, TriangleMi.

### Domain: Raycast occlusion culling

The raycast module represents occluders, instances, scenarios, and raycast thread data to provide renderer-scene occlusion culling.
Use these entity examples: RaycastOccluder.

## Unit 118

Lesson ID: `unit-resource-previews-scene-graph-scene-importing`

### Domain: Asynchronous resource previews

Resource previews use generators selected by handles, queue source paths or in-memory resources, and cache generated full and small textures with metadata.
Use these entity examples: PreviewRequest, ResourcePreviewCacheEntry.

### Domain: Scene graph nodes

Node2D and Node3D are scene graph node bases for specialized runtime behavior.
Use these entity examples: Node2D, Node3D.

### Domain: Scene importing

Scene importers expose extensions, options, flags, and an import operation that produces a Node-based scene.
Use these entity examples: ColladaParseState.

## Unit 119

Lesson ID: `unit-scene-tree-execution-screen-and-environment-effects-skeletal-ragdoll`

### Domain: SceneTree execution

A SceneTree schedules the Node hierarchy and maintains scene-level timer and tween processing collections.
Use these entity examples: Node.

### Domain: Screen-space and environment effects

RD effects and environment services allocate and process scene buffers for luminance, temporal anti-aliasing, reflections, ambient effects, fog, global illumination, sky, tone mapping, and scaling.

### Domain: Skeletons, animation, and ragdolls

Skeletons retain named joints and parent relationships, animations produce joint states, and ragdolls connect those joints to Body parts and two-body constraints.
Use these entity examples: Skeleton, Joint, Body.

## Unit 120

Lesson ID: `unit-spirv-generation-temporal-upscaling-dispatch-theora-video-streams`

### Domain: SPIR-V generation

A traverser lowers the front end's typed intermediate tree to SPIR-V instructions and a module/function/block graph.
Use these entity examples: spv::Module, spv::Function, spv::Block.

### Domain: Temporal upscaling dispatch

Vendored FSR2 dispatch code selects alternating frame resources, handles accumulation reset, and issues compute work based on render and display dimensions.

### Domain: Theora video streams

The Theora module defines a video-stream resource, playback implementation, resource loader, and OGV movie writer.
Use these entity examples: VideoStreamTheora.

## Unit 121

Lesson ID: `unit-tls-connection-state-ui-composition-variant-text-serialization`

### Domain: TLS connection and session state

The TLS state machine performs cryptographic operations while carrying configuration, negotiated session, handshake, record, and protocol state in mbedtls_ssl_context.
Use these entity examples: mbedtls_ssl_session.

### Domain: Control-tree user interfaces

Control-derived nodes compose user interfaces, and Container nodes automatically arrange their Control children.
Use these entity examples: Node.

### Domain: Variant text parsing and writing

VariantParser and VariantWriter serialize Variant values as String text through stream, token, tag, and resource-parser abstractions.
Use these entity examples: Variant, Array, Dictionary.

## Unit 122

Lesson ID: `unit-vehicle-dynamics-visual-shader-module-build-visual-shader-nodes`

### Domain: Vehicle dynamics

Vehicle dynamics couples a VehicleConstraint with a rigid Body, wheel collision tests, suspension, controllers, engine, transmission, differentials, tracks, and anti-roll bars.
Use these entity examples: Body, VehicleTransmissionSettings.

### Domain: Visual Shader source partition

The Visual Shader build script compiles the module's C++ sources, its node sources, and editor sources when building the editor.

### Domain: Visual shader nodes

Node resources provide reusable shader operations within typed graphs of node records and connection records.
Use these entity examples: VisualShaderNode.

## Unit 123

Lesson ID: `unit-websocket-peers-zstandard-stream-codec`

### Domain: WebSocket peers

WebSocketPeer defines packet-oriented socket behavior and has native and browser-backed implementations.
Use these entity examples: WebSocketPeer.

### Domain: Zstandard stream compression

The bundled code provides Zstandard stream compression contexts, decompression contexts, dictionaries, entropy tables, hash-based match finders, and optional worker pools.

## Unit 124

Lesson ID: `unit-3d-gizmo-authoring-cpp-native-wrappers`

### Domain: 3D gizmo authoring

Node3D editor gizmo plug-ins visualize and manipulate cameras, lights, meshes, occluders, particle emitters, physics objects, skeleton tools, and visibility volumes.

### Implementation: C++ inline native wrappers

C++ wrapper classes use inheritance and inline member functions to expose native framework methods.
Apply it through: MTL4::Archive, MTL4::BinaryFunction.

## Unit 125

Lesson ID: `unit-crypto-resources-csharp-source-generation`

### Domain: Cryptographic resources and contexts

CryptoKey and X509Certificate are Resource types, while AES, hashing, HMAC, TLS options, and crypto operations are exposed through dedicated contexts and services.
Use these entity examples: CryptoKey, X509Certificate.

### Implementation: C# partial types and source generation

Compilation produces attribute-bearing partial declarations for script paths and assembly script types.
Apply it through: ScriptPathAttribute, AssemblyHasScriptsAttribute, ManagedCallbacks.

## Unit 126

Lesson ID: `unit-editor-extensibility-font-streams-gui-control-layout`

### Domain: Editor plugins and customization hooks

EditorPlugin subclasses register inspector, import, export, debugger, and resource-preview hooks in the editor.
Use these entity examples: Resource.

### Domain: Font streams

Font input is represented by an FT_Stream that can deliver bytes through an in-memory base or a read callback.
Use these entity examples: FT_Stream, FT_BZip2FileRec.

### Domain: GUI controls and layout

Control-derived widgets and Container-derived layouts implement the scene GUI layer.
Use these entity examples: Control, Container, CodeEdit.

## Unit 127

Lesson ID: `unit-input-events-actions-localization-and-template-generation-managed-csharp-scripting`

### Domain: Input events and actions

InputEvent resource subclasses represent key, mouse, joypad, touch, gesture, MIDI, shortcut, and action input; InputMap associates named actions with InputEvent lists and deadzones.
Use these entity examples: InputEvent, InputMap::Action.

### Domain: Localization and translation-template generation

Localization tooling parses translation inputs, exposes locale selection and preview, edits localization settings, and generates message maps and POT-style template output.
Use these entity examples: Translation Message.

### Domain: managed C# script bridge and source generation

The Mono module represents managed scripts through CSharpScript and CSharpInstance, while its SDK tests use partial C# types and source generators for properties, methods, signals, serialization, and script paths.
Use these entity examples: CSharpScript.

## Unit 128

Lesson ID: `unit-navigation-agents-network-io-particle-systems`

### Domain: Navigation agents and regions

Navigation agents consume navigation-result paths, while regions, links, and obstacles describe navigation-world participation.
Use these entity examples: NavigationAgent2D, NavigationAgent3D, NavigationRegion3D.

### Domain: Networking and transport I/O

The core exposes HTTP clients, stream peers, packet peers, TCP and UDP servers, DTLS and TLS peers, IP resolution, and Unix-domain socket support.
Use these entity examples: IPAddress.

### Domain: CPU and GPU particle systems

2D and 3D particle nodes provide separate CPU and GPU particle-system implementations.
Use these entity examples: CPUParticles2D, GPUParticles3D.

## Unit 129

Lesson ID: `unit-scene-state-skeletal-ik-skeleton-modifiers`

### Domain: Scene state inspection

SceneState exposes a packed scene's resources, nodes, exported properties, overrides, and built-in scripts without instantiating the scene.
Use these entity examples: SceneState, Resource.

### Domain: Skeletal modifiers and inverse kinematics

IK skeleton modifiers process bone chains, targets, and joint-limitation resources to modify skeletal poses.
Use these entity examples: Resource.

### Domain: Skeleton modification and physical-bone simulation

2D modifier stacks order Resource-based bone modifications, while SkeletonModifier3D nodes run after AnimationMixer playback and can implement IK, constraints, and skeleton physics.
Use these entity examples: SkeletonModificationStack2D, SkeletonProfile.

## Unit 130

Lesson ID: `unit-spirv-shader-reflection-stream-networking-triangle-intersection-algorithms`

### Domain: SPIR-V shader reflection

The reflection API creates a shader-module metadata view from SPIR-V code and exposes entry points, descriptors, interfaces, push constants, and specialization constants.
Use these entity examples: SpvReflectShaderModule, SpvReflectEntryPoint.

### Domain: Byte streams and socket servers

StreamPeer exposes typed and raw byte-stream I/O, specialized peers implement TCP, TLS, compression, and Unix-domain sockets, and socket servers accept peer connections.
Use these entity examples: StreamPeerBuffer, RID.

### Domain: Triangle intersection algorithms

Triangle intersection uses Möller–Trumbore, Plücker, and Woop kernels; ray-primitive intersection is required to interpret their hit candidates.
Use these entity examples: TriangleMi.

## Unit 131

Lesson ID: `unit-two-dimensional-content-vorbis-block-synthesis-x509-certificate-processing`

### Domain: 2D shapes, tiles, and paths

Resource-owned Shape2D subclasses, TileSet data, NavigationPolygon data, and PolygonPathFinder provide 2D geometry and tile-oriented content.
Use these entity examples: TileMapPattern.

### Domain: Vorbis block analysis and synthesis

Vorbis analysis and synthesis operate on codec blocks and Ogg packets, using mapping, floor, residue, codebook, window, and MDCT modules.
Use these entity examples: Vorbis Block, Ogg Packet.

### Domain: X.509 certificate processing

TLS credential processing parses, writes, and verifies certificates, certificate requests, revocation lists, names, and related ASN.1 data.

## Unit 132

Lesson ID: `unit-compact-heightfield-compressed-font-stream-adapters-cpp-c-type-mapping`

### Domain: Compact heightfield representation

A compact heightfield stores grid dimensions, compact cells, spans, and per-span area data for navigation processing.
Use these entity examples: rcCompactHeightfield.

### Domain: Compressed font stream adapters

Optional Bzip2 and LZW components wrap a compatible source stream and expose decompressed bytes through the same FT_Stream callbacks.
Use these entity examples: FT_Stream, FT_BZip2FileRec.

### Implementation: C++ mapping of C API types

`CppType<Vk...>` specializations associate C Vulkan type names with C++ wrapper types.
Apply it through: GraphicsPipelineCreateInfo.

## Unit 133

Lesson ID: `unit-curve-and-patch-bases-editor-plugin-extension-cpp-class-inheritance`

### Domain: Parametric curve bases

Bezier, B-spline, and Catmull-Rom basis evaluators generate positions and derivatives from four control points.

### Domain: Editor plug-in extension

Plugins attach behavior through polymorphic C++ hook interfaces and may register import, scene-format, post-import, inspector, gizmo, debugger, and resource-conversion plugins.

### Implementation: C++ class inheritance for backend contracts

C++ class inheritance expresses shared renderer and server contracts and allows concrete dummy, extension, wrapper, clustered, and mobile implementations.

## Unit 134

Lesson ID: `unit-gui-controls-csharp-unsafe-interop`

### Domain: GUI controls and layout

Control-derived GUI classes implement selection, input, scrolling, split layout, menus, and graph editing within the CanvasItem tree.
Use these entity examples: GraphEditConnection.

### Implementation: C# unsafe interop and function pointers

Unsafe C# stores pointers and function pointers in ABI structs and invokes generated partial native calls.
Apply it through: ManagedCallbacks, Variant.

## Unit 135

Lesson ID: `unit-metal-cpp-object-bridge-motion-blur-geometry-python-build-code-generation`

### Domain: Metal-cpp object bridge

The Metal portion models framework objects as C++ wrapper classes whose inline methods send selectors to the underlying object.
Use these entity examples: MTL4::Archive.

### Domain: Motion-blur geometry

Motion-blur triangle code computes vertices at ray time before triangle intersection, so triangle intersection is required to interpret its candidates.
Use these entity examples: TriangleMi.

### Implementation: Python build code generation

The Python build helper defines generation and build-entry functions for virtual-method source generation.

## Unit 136

Lesson ID: `unit-platform-display-adaptation-regex-jit-codegen-renderer-storage`

### Domain: Platform display and window adaptation

DisplayServer implementations adapt platform windows and events; the Web adapter constructs InputEvent instances and forwards them to Input.
Use these entity examples: InputEvent.

### Domain: Regular-expression JIT code generation

PCRE2 JIT compilation turns compiled pattern control flow into architecture-specific generated code using SLJIT labels and jumps.
Use these entity examples: sljit_compiler, sljit_jump, sljit_label.

### Domain: Renderer-owned resource storage

The RD storage services own backend representations of lights, materials, meshes, particles, textures, and GPU resources addressed by rendering IDs.

## Unit 137

Lesson ID: `unit-rendering-assets-sfnt-tables-shader-interface-metadata`

### Domain: Meshes, materials, textures, and instancing

Mesh resources provide surfaces; geometry nodes instance them, Materials control shading, and MultiMesh batches many instances.
Use these entity examples: Mesh, Material.

### Domain: SFNT table loading

SFNT services read structured tables from font streams, including cmap, metrics, embedded bitmaps, color layers, SVG documents, and WOFF/WOFF2 data.
Use these entity examples: FT_Stream, PFR_FaceRec.

### Domain: Shader interface metadata

Reflection exposes descriptor sets, interface variables, push-constant blocks, and specialization constants from one reflected shader module; it depends on SPIR-V shader reflection because it is contained in the reflected shader module.
Use these entity examples: SpvReflectDescriptorSet, SpvReflectDescriptorBinding, SpvReflectInterfaceVariable, SpvReflectBlockVariable, SpvReflectSpecializationConstant.

## Unit 138

Lesson ID: `unit-skeletal-modifiers-and-ik-skeleton-modification-subgrid-intersection`

### Domain: Skeletal modifiers and inverse kinematics

SkeletonModifier3D subclasses apply bone constraints, IK solvers, retargeting, spring-bone simulation, and XR tracker data.
Use these entity examples: SkeletonModifier3D, ChainIK3D, SpringBoneSimulator3D.

### Domain: 2D skeleton modification stacks

SkeletonModification2D resources define individual 2D skeleton operations, and SkeletonModificationStack2D holds modifications for a Skeleton2D.

### Domain: Subgrid intersection

Subgrid intersection uses grid identifiers to load a quad neighborhood and applies triangle intersection algorithms to its triangles; triangle intersection algorithms are needed to interpret those tests.
Use these entity examples: SubGrid, GridMesh.

## Unit 139

Lesson ID: `unit-tilemap-layer-data-ui-theming-version-control-integration`

### Domain: Tile map layer data

TileMapLayer stores cell data keyed by grid coordinates and derives rendering, physics, navigation, and debug quadrants from that data.
Use these entity examples: TileMapLayer, TileMapLayerCellData.

### Domain: UI themes and style boxes

Theme resources apply reusable colors, constants, fonts, icons, and StyleBoxes across Control and Window branches, while StyleBox defines a visual box treatment.
Use these entity examples: Theme, StyleBox.

### Domain: Version-control integration

Version-control integration exchanges structured status, commit, file-diff, hunk, and line data through EditorVCSInterface and presents it through VersionControlEditorPlugin.
Use these entity examples: VCS Diff File, VCS Diff Hunk, VCS Diff Line.

## Unit 140

Lesson ID: `unit-webrtc-data-channels-websocket-frame-events`

### Domain: WebRTC data channels

A WebRTC data channel carries PacketPeer data through a WebRTC peer connection.
Use these entity examples: WebRTCDataChannel.

### Domain: WebSocket framing and event contexts

Wslay separates frame-level I/O from event-level message handling through frame contexts, event contexts, callback tables, message sources, and send/control queues.
Use these entity examples: wslay_event_context.

## Unit 141

Lesson ID: `unit-collada-import-c-function-pointer-interfaces`

### Domain: COLLADA scene interchange

The COLLADA parser supplies a scene importer with image, material, effect, mesh, skin, morph, node, visual-scene, and animation data collected in Collada::State.
Use these entity examples: ColladaParseState.

### Implementation: C function-pointer interfaces

C function pointers attach behavior to struct-and-pointer handles without requiring a language-level object system.
Apply it through: FT_Stream, FT_BZip2FileRec.

## Unit 142

Lesson ID: `unit-editor-plugin-lifecycle-font-table-validation-cplusplus-polymorphic-platform-adapters`

### Domain: Editor plugin lifecycle

C++ editor-plugin specializations package feature-specific editor behavior behind EditorPlugin-derived classes, including scene, script, and shader tools.

### Domain: OpenType and TrueTypeGX validation

OpenType and TrueTypeGX validators check structured tables after SFNT parsing so higher layers can consume checked offsets and indices.

### Implementation: C++ polymorphic platform adapters

C++ class inheritance and virtual functions define the target-platform adapter contract.
Apply it through: EditorExportPlatform.

## Unit 143

Lesson ID: `unit-gltf-materials-textures-cpp-engine-polymorphism`

### Domain: glTF material and texture conversion

GLTFDocument maps glTF material dictionaries, PBR metallic-roughness values, textures, samplers, UV coordinates, color transforms, alpha mode, and texture-transform extensions to Godot material resources.
Use these entity examples: GLTFState.

### Implementation: C++ inheritance, virtual interfaces, and Ref ownership

C++ classes use inheritance, virtual methods, Ref ownership, and static registries to implement core services.
Apply it through: Resource, InputEvent, GDExtension.

## Unit 144

Lesson ID: `unit-gpu-resources-pipelines-halfedge-topology-cpp-pluggable-allocation`

### Domain: GPU resources and pipelines

Descriptors configure GPU resources and pipeline state before a device produces allocations, functions, or pipeline state objects.
Use these entity examples: MTL::Buffer, MTL::Texture, MTL4::PipelineDescriptor.

### Domain: Halfedge topology

The Manifold mesh kernel represents every triangle boundary with directed Halfedge records whose pairedHalfedge links encode adjacency.
Use these entity examples: Halfedge, MeshGLP.

### Implementation: C++ pluggable allocation

meshoptimizer stores allocation and deallocation function pointers in static allocator storage so callers can replace the allocation policy.
Apply it through: meshopt_Meshlet.

## Unit 145

Lesson ID: `unit-managed-native-interop-navigation-mesh-construction-cpp-polymorphic-ownership`

### Domain: Managed-native interop

Unsafe C# connects managed values and callbacks across a fixed native/managed ABI using variant conversion, GC handles, and function-pointer callbacks.
Use these entity examples: ManagedCallbacks, Variant.

### Domain: Navigation mesh construction

The 2D and 3D generators turn source geometry into navigation mesh data and connect region or link geometry into map iterations.

### Implementation: C++ polymorphism and ownership

C++ classes use public inheritance and `std::unique_ptr` to retain polymorphic loader services and dispatch state.
Apply it through: LoaderInstance, XrGeneratedDispatchTableCore.

## Unit 146

Lesson ID: `unit-navigation-regions-java-android-host-api`

### Domain: Navigation region construction

Region construction labels connected compact heightfield spans into navigation regions.
Use these entity examples: rcCompactHeightfield.

### Implementation: Java Android host APIs

Java uses class inheritance and interface implementation to attach GL/Vulkan views, plugin registration, native library calls, and input handling to the Android platform host.

## Unit 147

Lesson ID: `unit-navmesh-heightfields-openxr-extension-wrappers-python-build-scripts`

### Domain: Navigation heightfields

Recast rasterizes input triangles into a heightfield and converts its spans into a compact heightfield with adjacency.
Use these entity examples: rcHeightfield, rcCompactHeightfield, rcCompactSpan.

### Domain: OpenXR extension wrappers

Extension wrappers isolate optional OpenXR runtime features behind a common base interface.
Use these entity examples: OpenXRFutureResult, OpenXRRenderModelData.

### Implementation: Python build scripts

Python build scripts import builder modules and define builder functions used by SCons-facing scene-theme scripts.

## Unit 148

Lesson ID: `unit-platform-presentation-pluggable-server-backends-post-import-processing`

### Domain: Display, camera, video, and movie services

DisplayServer abstracts display functionality, CameraServer manages CameraFeed objects, VideoStreamPlayer presents video in a Control, and MovieWriter emits movie output.

### Domain: Pluggable and wrapped server backends

C++ inheritance supplies extension, dummy, and multithread-wrapped implementations behind the physics, text, XR, and rendering server interfaces.

### Domain: Post-import processing

Post-import processing operates on the imported scene after format conversion and accepts options for node, mesh, material, animation, and skeleton categories.

## Unit 149

Lesson ID: `unit-resource-specific-editors-shader-reflection-spatial-indexing`

### Domain: Resource-specific editing

Dedicated editor controls and docks edit gradients, curves, materials, sprite frames, mesh libraries, textures, packed scenes, and resource preloaders.

### Domain: Shader reflection

The reflection API reads SPIR-V IR and returns module, entry-point, descriptor-binding, interface-variable, and push-constant metadata.
Use these entity examples: SpvReflectShaderModule, SpvReflectDescriptorBinding.

### Domain: Spatial indexing and ray queries

Spatial trees use geometry bounds to maintain items, cull them, refit nodes, and accelerate triangle-mesh and static-ray queries.
Use these entity examples: TriangleMesh, AABB.

## Unit 150

Lesson ID: `unit-subdivision-surface-evaluation-text-interface-themes-and-style-boxes`

### Domain: Subdivision surface evaluation

Subdivision evaluation combines curve and patch bases with grid sampling; curve and patch bases are needed to explain the evaluator.
Use these entity examples: SubGrid, GridMesh.

### Domain: Text layout and editing

GUI text controls hold shaped line and paragraph data, use text-server glyph and selection queries, and track IME text and selection state.

### Domain: Themes and style boxes

Resource-backed Theme data is resolved through ThemeDB, ThemeContext, and ThemeOwner, while StyleBox subclasses supply control styling.

## Unit 151

Lesson ID: `unit-three-dimensional-content-webrtc-multiplayer-mesh-websocket-multiplayer`

### Domain: 3D shapes, worlds, and skins

Resource-owned Shape3D subclasses, World3D, Skin, and mesh import data represent 3D collision, world, skeletal, and mesh content.

### Domain: WebRTC multiplayer mesh

A WebRTCMultiplayerPeer keeps a mesh of WebRTC peer connections and their data channels.
Use these entity examples: WebRTCMultiplayerPeer.

### Domain: WebSocket multiplayer

WebSocketMultiplayerPeer tracks WebSocket peers plus pending peers and packets.
Use these entity examples: WebSocketMultiplayerPeer, WebSocketMultiplayerPacket.

## Unit 152

Lesson ID: `unit-webxr-session-bridge-zlib-stream-codec`

### Domain: WebXR session bridge

WebXR interface state receives browser session callbacks and input sources through web-platform bindings.
Use these entity examples: WebXRInterfaceJS, WebXRInputSource.

### Domain: zlib stream compression

The bundled zlib sources implement Adler-32 and CRC-32 checksums, DEFLATE compression trees, and inflate state for a public streaming compression interface.
Use these entity examples: z_stream.

## Unit 153

Lesson ID: `unit-allocation-c-function-pointer-callbacks`

### Domain: Allocation and reference ownership

Reference-counted objects, FastAllocator blocks, cached allocation, aligned allocation, and monitored allocation support resource and BVH memory management.
Use these entity examples: RefCount, Ref, FastAllocator, CachedAllocator.

### Implementation: C function-pointer callbacks

C function-pointer fields let ENet obtain allocator, deallocator, and out-of-memory handlers through a cross-cutting callbacks record.
Apply it through: ENetCallbacks.

## Unit 154

Lesson ID: `unit-android-export-pipeline-basis-transcoding-cpp-generic-containers`

### Domain: Android export pipeline

Android export derives Gradle project files, manifest content, localized project names, icons, ABIs, and plugin configuration from project and export inputs.
Use these entity examples: APKExportData, PluginConfigAndroid, LauncherIcon, CustomExportData.

### Domain: Basis texture transcoding

The Basis Universal subsystem decodes ETC1S and UASTC texture slices into selected texture formats through high-level and low-level transcoders.
Use these entity examples: BasisFileHeader, BasisSliceDescriptor.

### Implementation: C++ templates and typed containers

C++ template applications express typed containers and handles, including Vector collections of Ref-managed connections and parameter-pack-sized argument arrays.
Apply it through: GraphEditConnection, SceneState.

## Unit 155

Lesson ID: `unit-builder-heuristics-cff-font-subsetting-objective-cpp-apple-adapters`

### Domain: BVH split heuristics

SAH, Morton, spatial, temporal, strand, open-merge, and motion-blur heuristics choose how PrimRef records are divided while constructing a BVH.
Use these entity examples: BinMapping, MortonCodeMapping, SpatialBinMapping, HeuristicMBlurTemporalSplit.

### Domain: CFF font subsetting

The font subsetting pipeline contains CFF1 and CFF2 plans, accelerators, serializers, and subroutine-subsetting helpers.
Use these entity examples: hb_subset_plan_t, hb_subset_accelerator_t.

### Implementation: Objective-C++ Apple platform adapters

Objective-C++ implementation files combine C++ platform adapters with Cocoa-style objects, display coordinates, events, menus, views, and text-to-speech calls.

## Unit 156

Lesson ID: `unit-editor-scene-sessions-python-build-source-generation`

### Domain: Editor scene sessions

EditorData represents each open editor scene with its root, project path, plugin state, selection, undo history ID, and timing/version metadata.
Use these entity examples: EditedScene.

### Implementation: Python build source generation

Python build helpers perform build-time source generation by transforming inputs into generated C++ source and invoking external translation tooling when available.
Apply it through: EditorTranslationList.

## Unit 157

Lesson ID: `unit-feature-adaptive-tessellation-gpu-command-encoding-python-scons-module-configuration`

### Domain: Feature-adaptive tessellation

Feature-adaptive tessellation recursively partitions parameter ranges for patch evaluation; subdivision evaluation is needed to explain the patch evaluation.
Use these entity examples: SubGrid.

### Domain: GPU command encoding

Command buffers create render, compute, blit, and resource-state encoders that record work against configured GPU resources.
Use these entity examples: MTL::CommandBuffer, MTL::RenderCommandEncoder, MTL::ComputeCommandEncoder.

### Implementation: Python SCons module configuration

Python SCons scripts declare module buildability, dependencies, cloned environments, and source-file collection.

## Unit 158

Lesson ID: `unit-half-edge-topology-inspector-property-editors-metalfx-scaling`

### Domain: Half-edge topology

Subdivision connectivity is represented by HalfEdge navigation and by Catmull–Clark rings that inspect neighboring edges, faces, and crease information.
Use these entity examples: HalfEdge.

### Domain: Custom inspector property editors

The implementation uses editor-plugin lifecycle integration to replace or extend property editing with specialized EditorInspectorPlugin and EditorProperty classes.

### Domain: MetalFX scaling and interpolation

MetalFX descriptors create spatial, temporal, temporal-denoised, and frame-interpolation effect instances from a Metal device.
Use these entity examples: MTLFX::SpatialScaler, MTLFX::TemporalScaler, MTLFX::FrameInterpolator.

## Unit 159

Lesson ID: `unit-navigation-navigation-map-composition-navmesh-contours-polygons`

### Domain: Navigation maps and path queries

Navigation agents use server maps, regions, path-query parameters, and path-query results to follow a target position.
Use these entity examples: NavigationPathQueryParameters2D, NavigationPathQueryResult2D.

### Domain: Navigation map composition

Navigation maps collect regions, links, agents, and obstacles and build per-frame read iterations.

### Domain: Navigation contours and polygons

Recast extracts contours from a compact heightfield, then builds polygon and detail meshes from those contours.
Use these entity examples: rcContourSet, rcPolyMesh, rcPolyMeshDetail.

## Unit 160

Lesson ID: `unit-openexr-image-decoding-physics-2d-collision-pipeline-platform-selective-shader-baking`

### Domain: OpenEXR image decoding

TinyEXR is compiled as a header implementation with zlib included first, and its public header defines EXR-oriented image API data and functions.
Use these entity examples: EXRImage.

### Domain: native 2D collision pipeline

Physics spaces hold collision objects; the 2D BVH broad phase finds candidate collision objects and the solver dispatches shape-pair routines, including separating-axis tests.
Use these entity examples: GodotSpace2D, GodotCollisionObject2D.

### Domain: Platform-selective shader baking

Platform-selective shader baking depends on shader editing and compiles distinct Vulkan, D3D12, or Metal export-plugin sources when matching SCons environment flags are enabled.

## Unit 161

Lesson ID: `unit-raster-image-input-rendering-resources-resource-importing`

### Domain: Raster image input

Raster image input adapters decode PNG and JPEG bytes into image buffers for texture-processing callers.
Use these entity examples: PngInfo.

### Domain: Textures, meshes, and materials

Resource-owned Mesh, Material, Texture, Environment, and Sky classes carry rendering-facing content and configuration.

### Domain: Standalone resource importing

Resource-importer classes handle images, textures, texture atlases, SVG, fonts, WAV audio, CSV translations, shaders, and bitmaps.

## Unit 162

Lesson ID: `unit-textures-and-placeholders-variable-font-subsetting`

### Domain: Texture resources and fallback placeholders

Texture resources represent 2D, 3D, layered, and RenderingDevice-backed image data; placeholder texture classes preserve limited shape information when source texture data is unavailable or omitted.
Use these entity examples: RDTextureFormat, RID.

### Domain: Variable-font table subsetting

The font subsetting pipeline dispatches HVAR, VVAR, gvar, fvar, avar, cvar, and mvar tables, with fvar and avar passed through when user axis locations are empty.
Use these entity examples: hb_subset_plan_t.

## Unit 163

Lesson ID: `unit-android-gradle-assembly-apk-signing-cpp-templates-and-const-access`

### Domain: Android Gradle assembly

Gradle settings and build logic assemble app, library, editor, native-source-index, and install-time asset-pack modules for the Android export pipeline.

### Domain: APK signing and verification

APK signing and verification operate on Android export pipeline artifacts through the bundled apksig implementation.

### Implementation: C++ templates and const access

C++ templates and const-qualified values provide typed access to engine handles and collection views.
Apply it through: VisualShaderGraph, NavigationAgent3D.

## Unit 164

Lesson ID: `unit-basis-container-layout-block-texture-transcoding-objective-cpp-ios-adapters`

### Domain: Basis file layout

A .basis file header and slice descriptors locate compressed slices and identify their texture form before Basis Universal decodes them.
Use these entity examples: BasisFileHeader, BasisSliceDescriptor.

### Domain: GPU block texture conversion

Texture conversion code maps ETC1S, UASTC, and ASTC blocks to GPU block or uncompressed target formats, calculating output block counts and storage.

### Implementation: Objective-C++ iOS adapters

Objective-C++ iOS adapters implement an embedded platform layer for the iOS embedded adapter.

## Unit 165

Lesson ID: `unit-buffer-storage-catmull-clark-patch-construction-editing-selection-history`

### Domain: Geometry buffer storage

A Buffer is a reference-counted storage abstraction, while BufferView and RawBufferView expose views used to bind geometry data.
Use these entity examples: RTCBuffer, Buffer.

### Domain: Catmull–Clark patch construction

Catmull–Clark patch construction reads a half-edge topology and converts its limit data to patch geometry; half-edge topology is needed to explain this construction.
Use these entity examples: HalfEdge.

### Domain: Editing selection history

Selection history gives an editor scene session back-and-forward navigation across selected objects and nested-resource properties.
Use these entity examples: EditedScene.

## Unit 166

Lesson ID: `unit-editor-plugin-state-font-driver-modules-godot-thirdparty-adaptation`

### Domain: Editor plugin state and custom types

Editor plugins can handle objects, register custom types, and serialize plugin state into the active scene's editor state.
Use these entity examples: EditedScene.

### Domain: Font driver modules

Format-specific driver classes expose CFF, CID, PCF, PFR, Type 1, Type 42, Windows FNT, and TrueType implementations to the FreeType module layer.
Use these entity examples: PFR_FaceRec.

### Domain: Godot-specific third-party adaptation

The repository applies VMA allocator integration patches, redirects Vulkan includes to Godot's Vulkan header, adds lazy-allocation statistics, and configures portability macros for bundled dependencies.

## Unit 167

Lesson ID: `unit-gpu-memory-suballocation-histograms-and-huffman-codes-ktx2-container-transcoding`

### Domain: D3D12 memory allocation

D3D12MA's Allocator, Pool, and VirtualBlock APIs manage resource and virtual allocations using allocation callbacks.

### Domain: Histograms and Huffman codes

Histograms built from backward-reference symbols are clustered and converted into Huffman code lengths, codes, and coded tree headers.
Use these entity examples: VP8LBackwardRefs.

### Domain: KTX2 container transcoding

The KTX2 transcoder retains the source data, parses the KTX2 header and per-level index, then transcodes its texture levels.
Use these entity examples: Ktx2Header, Ktx2LevelIndex, Ktx2TranscoderState.

## Unit 168

Lesson ID: `unit-lossless-transform-pipeline-lossy-macroblock-encoding-metal-presentation`

### Domain: Lossless pixel transform pipeline

The lossless encoder transforms ARGB pixels through predictor, subtract-green, cross-color, and color-indexing stages before entropy coding.
Use these entity examples: WebPPicture, VP8Encoder.

### Domain: Lossy macroblock encoding

The lossy encoder walks the configured picture planes by macroblock, choosing prediction modes and producing residual tokens for the VP8 stream.
Use these entity examples: VP8Encoder, WebPPicture, WebPConfig.

### Domain: Metal drawable presentation

QuartzCore Metal layers produce drawables that implement the Metal drawable interface for presentation.
Use these entity examples: CA::MetalLayer, CA::MetalDrawable.

## Unit 169

Lesson ID: `unit-navigation-avoidance-navigation-path-queries-networking`

### Domain: Navigation avoidance

Agents and obstacles are map members that provide avoidance data alongside pathfinding data.

### Domain: Navigation path queries

Path queries consume a map read iteration, select polygons, and construct a path corridor in 2D or 3D.

### Domain: HTTP and multiplayer

HTTPRequest is a Node API for HTTP work, while MultiplayerAPI and MultiplayerPeer define reference-counted multiplayer and packet-peer interfaces.

## Unit 170

Lesson ID: `unit-object-identity-lifetime-physics-3d-collision-pipeline-random-generation`

### Domain: Object identity and lifetime

Object supplies engine object identity and ObjectDB support, while RefCounted and Ref provide reference-managed object use.
Use these entity examples: Object, ObjectID.

### Domain: native 3D collision pipeline

The native 3D server manages spaces, bodies, shapes, and joints; its collision code includes GJK/EPA support and SAT shape-pair routines.
Use these entity examples: GodotSpace3D, GodotCollisionObject3D.

### Domain: Pseudo-random generation

RandomPCG supplies pseudo-random generation and RandomNumberGenerator exposes a reference-counted runtime wrapper.

## Unit 171

Lesson ID: `unit-two-dimensional-physics-undo-redo-history-variable-font-data`

### Domain: 2D physics nodes

2D physics scene graph nodes provide collision objects, bodies, joints, casts, and collision results.
Use these entity examples: CollisionObject2D, RigidBody2D, KinematicCollision2D.

### Domain: Undo and redo history

Undo and redo assigns actions and saved versions to a scene session, allowing edits to identify a history and report unsaved state.
Use these entity examples: EditorUndoRedoManager::History, EditedScene.

### Domain: TrueType GX variation data

TrueType GX variation loading maps a face's variation tables and coordinates into variable-font state.

## Unit 172

Lesson ID: `unit-zip-archive-io`

### Domain: ZIP archive I/O

ZIPReader reads archives and ZIPPacker creates archives through reference-counted API objects.
Use these entity examples: ZIPReader, ZIPPacker.

## Unit 173

Lesson ID: `unit-block-texture-encoding-dynamic-invocation-signals-cmake-native-source-index`

### Domain: Block texture encoding

CVTT compression input uses pixel blocks, options, encoding plans, and format-specific BC6H, BC7, ETC, and S3TC computation.

### Domain: Dynamic invocation and signals

The object and callable layers call Object methods and signal handlers with Variant argument arrays, including bound and unbound callable forms.
Use these entity examples: Object, Variant.

### Implementation: CMake native source indexing

CMake source-index configuration exposes the native C/C++ source tree to Android Studio as part of Android Gradle assembly.

## Unit 174

Lesson ID: `unit-entropy-bitstreams-gltf-accessors-cpp-templates-for-binary-data`

### Domain: Entropy bitstreams

VP8 and VP8L writers serialize codec decisions into range-coded or raw bit-level byte streams while readers maintain cursor, range, and value state.

### Domain: glTF binary accessor decoding and encoding

GLTFAccessor converts typed values to and from GLTFState’s indexed buffer views, including sparse data when the sparse representation is smaller than dense data.
Use these entity examples: GLTFAccessor, GLTFBufferView, GLTFState.

### Implementation: C++ templates for binary data operations

C++ templates specialize reusable byte operations for a compile-time size and typed OpenType offsets.
Apply it through: hb_blob_t, hb_face_t.

## Unit 175

Lesson ID: `unit-glyph-caching-glyph-rasterization-groovy-gradle-build-logic`

### Domain: Glyph and face caching

The cache manager holds faces made available by format drivers and caches character maps, glyph images, and small-bitmap nodes under resource limits.
Use these entity examples: FTC_SNodeRec.

### Domain: Glyph rasterization

Renderer modules turn driver-loaded glyph outlines into monochrome, gray, SDF, or SVG-backed bitmap representations.
Use these entity examples: SVG_RendererRec.

### Implementation: Groovy Gradle build logic

Groovy Gradle scripts generate Android build tasks and select variants as part of Android Gradle assembly.

## Unit 176

Lesson ID: `unit-ktx-texture-container-native-support-algorithms-navigation-queries`

### Domain: KTX texture containers

KTX texture containers carry a texture format description, per-level indexing, stream or memory-backed data, virtual dispatch tables, and KTX2 supercompression state.
Use these entity examples: KTX2 Texture, KTX2 Private Texture State, KTX Level Index Entry.

### Domain: Native support algorithms

The support subtree implements deterministic PCG random state evolution alongside compression, noise, color, audio, and packing algorithms.
Use these entity examples: pcg32_random_t.

### Domain: 2D and 3D navigation queries

RefCounted navigation query parameter and result objects exchange path points, path types, and path-owner identifiers with the 2D and 3D navigation server APIs.
Use these entity examples: NavigationPathQueryResult2D.

## Unit 177

Lesson ID: `unit-physics-server-queries-physics-space-queries-postscript-font-services`

### Domain: 2D and 3D physics queries

Physics servers expose direct body and space state APIs, while RefCounted query parameter and result objects package ray, point, shape, and motion requests.

### Domain: 3D physics query contracts

Physics-server types define ray, point, shape, and motion parameter and result records; RefCounted query objects expose those contracts to callers.
Use these entity examples: PhysicsRayQueryParameters3D.

### Domain: PostScript font services

PSAux supplies parsing, decoding, character maps, and hint services shared by CFF, CID, and Type 1 drivers.

## Unit 178

Lesson ID: `unit-prefix-code-decoding-reflection-metadata-regular-expression-results`

### Domain: Bitstream and Huffman decoding

Both Basis and Brotli implement bit readers and Huffman decoding tables to recover symbols from compressed bit streams.

### Domain: Reflection metadata

ClassDB, GDType, MethodInfo, and PropertyInfo describe Object classes, inheritance, methods, properties, constants, enums, and signals.
Use these entity examples: Object, GDType, MethodInfo, PropertyInfo.

### Domain: Regular-expression matching

The regex module exposes compiled regular expressions and RefCounted match results with ranges.
Use these entity examples: RegExMatch.

## Unit 179

Lesson ID: `unit-script-hosting-target-platform-export-three-dimensional-physics`

### Domain: Script languages and instances

The scripting layer attaches Object-facing script instances, manages registered script languages, and provides extension-backed script implementations.
Use these entity examples: Object, Variant.

### Domain: Target-platform export

A target platform implementation supplies target-specific validation, option, run, and project, pack, or ZIP export behavior to export orchestration.
Use these entity examples: EditorExportPlatform, EditorExportPlatform::ExportMessage.

### Domain: 3D physics nodes

3D physics scene graph nodes provide collision objects, bodies, joints, casts, soft bodies, and vehicles.
Use these entity examples: CollisionObject3D, RigidBody3D, SoftBody3D.

## Unit 180

Lesson ID: `unit-transform-quantization-rate-distortion-upnp-device-control-web-javascript-bridge`

### Domain: Transform, quantization, and rate-distortion search

The macroblock path transforms residuals, quantizes them with VP8Matrix parameters, evaluates rate-distortion costs, and reconstructs blocks.
Use these entity examples: VP8Encoder.

### Domain: UPnP device control

The UPnP module represents generic devices and MiniUPnP-backed specializations behind RefCounted APIs.
Use these entity examples: UPNPDevice.

### Domain: Web JavaScript bridge

The JavaScript bridge exposes evaluation, interface lookup, callback creation, object creation, buffer conversion, downloads, PWA operations, and JavaScript object proxies.
Use these entity examples: JavaScriptObjectImpl.

## Unit 181

Lesson ID: `unit-webp-image-io-zip64-archive-io`

### Domain: WebP image I/O

The WebP module implements image loading, saving, and shared buffer helpers.
Use these entity examples: ImageLoaderWebP, ResourceSaverWebP.

### Domain: ZIP64 archive I/O

MiniZip reads and writes ZIP archives while its I/O API abstracts file opening, seeking, telling, and allocation callbacks.
Use these entity examples: zlib_filefunc64_32_def.

## Unit 182

Lesson ID: `unit-brotli-stream-decompression-class-reference-generation-cpp-templates-traits`

### Domain: Brotli stream decompression

The Brotli decoder consumes a guarded byte reader, decodes Huffman and prefix-coded streams, and exposes output from a decoder state.
Use these entity examples: BrotliDecoderState.

### Domain: Class-reference generation

Documentation generation parses class-reference XML into class, type, property, parameter, signal, method, and constant definitions before producing reStructuredText.
Use these entity examples: ClassDef, TypeName, ClassStatusProgress.

### Implementation: C++ templates and traits

The implementation uses C++ templates and trait specializations to adapt behavior to types, including zero construction, hashing, tuple recursion, and argument conversion.
Apply it through: AABB.

## Unit 183

Lesson ID: `unit-endian-safe-binary-io-native-extensions-physics-queries`

### Domain: Endian-safe binary I/O

Endian macros and memory conversion helpers normalize integer fields while container and bitstream code serializes binary bytes.
Use these entity examples: WebPData.

### Domain: Native extension loading

A GDExtension is a Resource loaded through a loader; its registered classes become ClassDB-visible extension classes.
Use these entity examples: GDExtension, ClassInfo.

### Domain: Physics queries and motion tests

Physics queries use physics spaces to submit configured point, ray, shape, and motion tests and receive collision data.
Use these entity examples: PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D.

## Unit 184

Lesson ID: `unit-platform-exports-signed-distance-fields-theora-block-video-codec`

### Domain: Platform export packaging

Platform export plug-ins implement macOS, Web, Windows, and visionOS export behavior; the Windows exporter also includes PE template modification structures.

### Domain: Signed-distance-field glyph rendering

The SDF and BSDF renderers use rasterization to populate signed-distance output bitmaps with configurable parameters.

### Domain: Theora block video coding

The Theora codec consumes Ogg packets while unpacking quantization parameters, DCT token data, motion-compensated frame state, and optional accelerated transform routines.
Use these entity examples: Theora Stream Information, Ogg Packet.

## Unit 185

Lesson ID: `unit-undo-redo-upnp-device-discovery-vulkan-pipeline-configuration`

### Domain: Undo and redo actions

UndoRedo records Object/Callable operations as actions so they can be replayed in undo or redo direction.
Use these entity examples: Object, Variant.

### Domain: UPnP device discovery

MiniUPnPc discovers UPnP devices, represents them as a linked device list, and parses an IGD description into URLs and service data.
Use these entity examples: UPNPDev, UPNPUrls, IGDdatas.

### Domain: Graphics pipeline configuration

`GraphicsPipelineCreateInfo` groups shader stages with separate vertex-input, assembly, viewport, rasterization, multisample, depth-stencil, color-blend, and dynamic-state records.
Use these entity examples: GraphicsPipelineCreateInfo.

## Unit 186

Lesson ID: `unit-image-decode-pipeline-cpp-template-binding`

### Domain: Image decode pipelines

The vendored image libraries parse encoded input, keep decoder state, and emit raster rows or planes; JPEG has scanline and upsampling modules, PNG applies row transforms, and WebP routes decoded output through VP8 I/O.
Use these entity examples: JPEG Decompression Context, PNG Information State, WebP Decoder State.

### Implementation: C++ template-based method binding

ClassDB uses C++ templates, `if constexpr`, and member-function traits to generate binders for native methods.
Apply it through: ClassInfo, Variant.

## Unit 187

Lesson ID: `unit-upnp-port-mapping`

### Domain: UPnP port-mapping control

SOAP command helpers add, delete, query, and list port mappings on a discovered device.
Use these entity examples: PortMapping, PortMappingParserData.

## Unit 188

Lesson ID: `unit-jpeg-decompression-transcoding-cxx-templates`

### Domain: JPEG decompression and coefficient transcoding

The JPEG code implements an image decode pipeline that selects decompression modules, can merge chroma upsampling with YCC-to-RGB conversion, and can expose raw DCT coefficient arrays for transcoding.
Use these entity examples: JPEG Decompression Context, JPEG Upsampler State.

### Implementation: C++ templates and specialization

C++ templates parameterize BVH branching, packet width, vector width, primitive layout, and builder behavior at compile time.
Apply it through: BVHN, RTCRayNt, MortonBuilder.

## Unit 189

Lesson ID: `unit-png-stream-transforms-webp-riff-decoding`

### Domain: PNG streaming I/O and row transforms

The PNG code implements an image decode pipeline with replaceable read and write callbacks, push-mode input states, metadata validity flags, and row-level transformations such as BGR mapping and 16-bit byte swapping.
Use these entity examples: PNG Information State.

### Domain: WebP chunk parsing, incremental decode, and animation

The WebP code implements an image decode pipeline that recognizes VP8 and VP8L payloads, incrementally retains input, demuxes RIFF chunks into frames, and composites animation frames.
Use these entity examples: WebP Decoder State.

## Unit 190

Lesson ID: `unit-png-image-codec-cpp-class-templates`

### Domain: PNG image codec

The PNG driver supplies image loading and resource saving implementations and can build bundled libpng sources.

### Implementation: C++ class templates and specialization

C++ class templates parameterize Embree vector and intersector code by lane width, and a fixed-width specialization selects a width-specific implementation; C++ classes are needed to explain class templates.
Apply it through: TriangleMi, SubGrid.

## Unit 191

Lesson ID: `unit-cpp-ownership-and-polymorphism-cpp-runtime-symbol-loading`

### Implementation: C++ ownership and polymorphic trees

C++ class templates support CSG nodes through shared ownership, enable_shared_from_this, derived leaf and operation classes, and the Manifold implementation boundary.
Apply it through: Manifold.

### Implementation: C++ runtime symbol loading

C++ wrapper types resolve framework symbols dynamically through `dlsym`-based helper templates.
Apply it through: MTL4::Archive.

## Unit 192

Lesson ID: `unit-cpp-specialized-marshalling-cpp-templates-and-specialization`

### Implementation: C++ specialized argument marshalling

The binding layer specializes PtrToArg through C++ templates to marshal bound method types and return values.
Apply it through: Variant, PropertyInfo.

### Implementation: C++ templates and specialization

The glTF importer uses C++ class templates for interpolation and supplies a quaternion specialization; accessor decoding also instantiates numeric decoding for concrete element types.
Apply it through: GLTFAnimation, GLTFAccessor.

## Unit 193

Lesson ID: `unit-cxx-wide-simd`

### Implementation: C++ SIMD wrapper specialization

C++ templates and wrapper types represent fixed-width integer, unsigned-integer, float, mask, and long-vector lanes while ISA-specific headers implement their operations.
Apply it through: vint4, vint8, vuint16.

## Unit 194

Lesson ID: `unit-c-platform-selection`

### Implementation: C preprocessor platform and precision selection

C conditional compilation selects precision-dependent names, optional compiler intrinsics, and architecture-specific implementations before the codec modules are compiled.
Apply it through: JPEG Upsampler State.

## Unit 195

Lesson ID: `unit-cpp-simd-intrinsics-python-scons-build-hooks`

### Implementation: C++ SIMD wrappers and intrinsics

C++ vector wrappers invoke architecture intrinsics to operate on multiple lane values.

### Implementation: Python SCons build hooks

Python build scripts define platform detection, option and flag setup, configuration hooks, bundle generation, template assembly, and Emscripten helper functions.

## Unit 196

Lesson ID: `unit-simd-accelerated-codecs-cxx-preprocessor-configuration`

### Domain: Optional SIMD codec paths

The codec libraries use C conditional compilation to select architecture-specific NEON, SSE2, VSX, LSX, MMX, and other optimized routines, while retaining scalar fallbacks.
Use these entity examples: JPEG Upsampler State.

### Implementation: Preprocessor-selected platform configuration

Preprocessor configuration selects processor architecture, platform APIs, floating-point precision, instruction-set support, diagnostics, and optional subsystems.
Apply it through: Body.

## Unit 197

Lesson ID: `unit-cxx-reflection-macros`

### Implementation: Macro-based RTTI registration

Macro-based RTTI registration attaches class metadata and serializable member attributes to types that participate in dynamic lookup and object streaming.
Apply it through: ConstraintSettings, PhysicsMaterialSimple.

## Unit 198

Lesson ID: `unit-gdscript-query-objects`

### Implementation: GDScript query-object API use

GDScript can configure a physics query object, pass it to a physics space, and receive the collision result through chained method calls.
Apply it through: PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D.

## Unit 199

Lesson ID: `unit-serialization-simd-ray-packets-c-preprocessor-platform-selection`

### Domain: RTTI-based serialization

RTTI-based serialization registers attributes for object types and writes or restores binary state through StreamIn and StreamOut.
Use these entity examples: ConstraintSettings, PhysicsMaterialSimple, Skeleton.

### Domain: SIMD ray packets

SIMD ray packets evaluate width-parameterized ray and triangle lanes, so triangle intersection is required to interpret lane-validity masks.
Use these entity examples: TriangleMi.

### Implementation: C preprocessor platform selection

C preprocessor conditions select platform implementations and preserve a C linkage boundary for C++ compilation units.

## Unit 200

Lesson ID: `unit-cxx-simd-alignment-cxx-stream-serialization`

### Implementation: SIMD wrappers and alignment

SIMD vector types use aligned storage and compile-time CPU branches to implement vector, matrix, bounding-box, and geometry operations across SSE, NEON, and RISC-V vector paths.
Apply it through: Body.

### Implementation: Stream-oriented binary serialization

StreamIn and StreamOut implement binary persistence using RTTI hashes for dynamically typed objects and explicit field writes for concrete state.
Apply it through: ConstraintSettings, PhysicsMaterialSimple, Skeleton.

## Unit 201

Lesson ID: `unit-gdscript-declarations`

### Implementation: GDScript declarations and typing

GDScript source uses `extends`, `class_name`, typed variables, functions, and annotations; the engine represents each script as a GDScript Resource.
Apply it through: GDScript, Variant.

## Unit 202

Lesson ID: `unit-isa-portability-sdl-platform-backends-cpp-preprocessor-configuration`

### Domain: ISA and platform portability

Preprocessor-selected ISA configuration chooses SIMD namespaces and headers, while platform shims adapt unavailable WebAssembly control-register operations.
Use these entity examples: vint4, vint8, vuint16.

### Domain: Compile-time platform backends

SDL uses compile-time conditions to select Linux, Windows, macOS, dummy, and other platform backend implementations.

### Implementation: C++ preprocessor configuration

Preprocessor conditions select optional HarfBuzz subsystems and header inclusion boundaries before C++ compilation.
Apply it through: hb_subset_plan_t.

## Unit 203

Lesson ID: `unit-state-recording-gdscript-signals-callables`

### Domain: State recording and validation

State recording saves and validates Body and Constraint state through a stream that can compare replayed bytes with current values.
Use these entity examples: Body, TwoBodyConstraint.

### Implementation: GDScript signals and callbacks

GDScript connects a Signal to a Callable and uses callback functions such as `_input` in Object-derived nodes.
Apply it through: Node, Variant.

## Unit 204

Lesson ID: `unit-types-inference-and-conversions`

### Implementation: Types, inference, and conversions

GDScript fixtures contrast explicit type hints, `:=` inference, `Variant` boundaries, null assignment, and conversion at typed assignments and returns.
Apply it through: Test Script Fixture, Expected Result Fixture.

## Unit 205

Lesson ID: `unit-hid-device-enumeration-profiling-interning-c-compatible-header-guards`

### Domain: HID device enumeration and backends

SDL enumerates HID devices into linked device-information records and routes opened devices through platform and driver backends.
Use these entity examples: SDL_hid_device_info.

### Domain: Profiling name interning

The profiling implementation interns StringName metadata and source-location data for tracing backends.
Use these entity examples: StringName.

### Implementation: C-compatible guarded headers

The VI declaration file uses preprocessor guards and conditionally enters `extern "C"` when compiled as C++, so one declaration set can be included by C and C++ translation units.
Apply it through: VkViSurfaceCreateInfoNN.

## Unit 206

Lesson ID: `unit-rendering-backends-cpp-jni-variant-marshalling`

### Domain: Rendering backend abstraction

Rendering backends specialize common context and device-driver abstractions for Vulkan, GLES3, Direct3D 12, and Metal.
Use these entity examples: RenderingDeviceDriverVulkan::BufferInfo, RenderingDeviceDriverVulkan::CommandBufferInfo.

### Implementation: C++ JNI Variant marshalling

C++ JNI Variant marshalling builds on RefCounted adapter classes to convert Java primitive wrappers, strings, arrays, objects, and callables into engine Variant values.
Apply it through: Callable.

## Unit 207

Lesson ID: `unit-sdl-thread-abstractions-vector-font-export-cpp-single-header-implementation`

### Domain: Thread and synchronization abstractions

SDL implements mutex, semaphore, condition, read/write lock, thread, and TLS abstractions with generic, pthread, and Windows backends.

### Domain: Vector font export

Color paint records are emitted through vector drawing and paint backends that include PDF and SVG implementations.
Use these entity examples: hb_vector_paint_t, hb_vector_draw_t.

### Implementation: C++ single-header implementation selection

VMA uses preprocessor configuration to place its header implementation in one C++ translation unit while selecting a non-MSVC debug macro before inclusion.
Apply it through: VmaAllocatorCreateInfo.

## Unit 208

Lesson ID: `unit-cpp-static-abi-contracts-cpp-tagged-storage`

### Implementation: C++ static ABI contracts

The binding uses a macro-selected assertion facility to verify wrapper ABI layout and copy or move properties at compile time.

### Implementation: C++ tagged storage and casts

The implementation represents runtime Variant values with a C++ type tag and accesses payloads through explicit casts and type-specific accessors.
Apply it through: Variant, Transform3D.

## Unit 209

Lesson ID: `unit-cpp-variadic-binding`

### Implementation: C++ variadic binding

The binding code uses variadic templates and parameter packs to move argument lists into Variant pointer arrays.
Apply it through: Variant, Object.

## Unit 210

Lesson ID: `unit-cxx-closures`

### Implementation: C++ lambdas and callback objects

C++ lambdas and callable objects bind BVH builders to allocation, node creation, leaf creation, progress, and completion behavior.
Apply it through: ProgressMonitorClosure, GeneralBVHBuilder, BVHNBuilderVirtual.

## Unit 211

Lesson ID: `unit-cxx-lambdas-standard-algorithms-godot-shader-language`

### Implementation: Lambdas and standard algorithms

Lambdas provide inline comparison and initialization behavior when standard algorithms or static data construction need local callable logic.
Apply it through: LinearCurve.

### Implementation: Godot shader source composition

Godot's shader language writes Shader resource source and can compose reusable ShaderInclude fragments with a preprocessor include directive.
Apply it through: ShaderMaterial, RDPipelineSpecializationConstant.

## Unit 212

Lesson ID: `unit-typed-collections-warnings-and-suppression`

### Implementation: Typed arrays and dictionaries

Typed arrays and dictionaries apply static typing to element and key/value shapes; the fixture suite exercises declarations, inference, casts, defaults, transfers, and invalid runtime assignments.
Apply it through: Test Script Fixture, Expected Result Fixture.

### Implementation: Warnings and selective suppression

The analyzer reports unsafe, unused, shadowing, inference, conversion, and coroutine cases, while `@warning_ignore` selectively suppresses specified warnings.
Apply it through: Expected Result Fixture, Test Script Fixture.

## Unit 213

Lesson ID: `unit-android-jni-interop-android-render-input-cplusplus-export-callbacks`

### Domain: Android JNI interoperation

JNI interoperation converts values and directs callbacks across the Android platform host’s Java/native boundary using C++ Variant marshalling.
Use these entity examples: Callable.

### Domain: Android rendering and input

Android rendering views and native rendering drivers carry surface and input work from the Android platform host.

### Implementation: C++ export callbacks

C++ function-pointer and callable declarations connect generic packaging code to save, removal, shared-object, and script callback implementations.
Apply it through: EditorExportPlatform.

## Unit 214

Lesson ID: `unit-deferred-execution-gamepad-mapping-cpp-const-borrowing`

### Domain: Deferred calls and worker tasks

MessageQueue records Object/Callable work for later execution, while WorkerThreadPool represents task and group work for worker threads.
Use these entity examples: Object, Variant.

### Domain: Gamepad mapping and classification

Gamepad mapping classifies HID enumeration results and maps device controls to SDL gamepad axes and buttons.
Use these entity examples: SDL_hid_device_info.

### Implementation: C++ const references and pointers

C++ const references and pointers expose existing repository objects and buffers without copying them, such as shader parameter names, vectors, arrays, and geometry data.
Apply it through: Animation, GraphEditConnection.

## Unit 215

Lesson ID: `unit-generic-containers-multichannel-distance-fields-cpp-visitor-traversal`

### Domain: Generic container infrastructure

The template layer supplies the C++ storage and lookup structures that underpin Array and Dictionary, including CowData, Vector, maps, sets, lists, spans, and RID owners.
Use these entity examples: Array, Dictionary.

### Domain: Multi-channel distance fields

MSDFgen represents vector shapes, colors their edges, and generates distance-field pixels for one, three, or four channels.
Use these entity examples: msdfgen::Shape, msdfgen::BitmapSection<float, 3>.

### Implementation: C++ visitor-style intermediate-tree traversal

C++ inheritance provides traversal specializations that analyze intermediate shader-tree nodes.
Apply it through: TObjectReflection.

## Unit 216

Lesson ID: `unit-native-platform-adapters-visual-shader-vector-operations-gdscript-typed-collections`

### Domain: Native platform adapters

Platform adapters implement filesystem, sockets, IP resolution, operating-system services, and threads behind engine interfaces selected by the driver build.

### Domain: Visual shader vector operations

Vector-operation nodes apply arithmetic and vector functions as operations within a visual shader graph.
Use these entity examples: VisualShaderNodeVectorOp.

### Implementation: GDScript typed collections and signature compatibility

This extends type declarations with typed Array and Dictionary element constraints, plus parameter and return compatibility checks for inherited functions.
Apply it through: GDScriptParser::Node.

## Unit 217

Lesson ID: `unit-glsl-compute-shaders`

### Implementation: GLSL compute shaders

Betsy shader sources declare compute workgroup sizes, bind sampled and storage images, fetch texels by invocation ID, and store generated output.
Apply it through: Image.

## Unit 218

Lesson ID: `unit-android-plugin-signals-c-function-pointer-apis`

### Domain: Android plugins and signals

Android plugins are discovered from application metadata and expose annotated methods and signals through JNI interoperation.
Use these entity examples: SignalInfo.

### Implementation: C function-pointer API declarations

C function-pointer declarations expose addressable Vulkan entry points and callback-based frame/event integration without requiring a concrete implementation type.
Apply it through: VkWaylandSurfaceCreateInfoKHR, VkSurfaceKHR.

## Unit 219

Lesson ID: `unit-android-storage-bridge-apple-embedded-hosting-cpp-const-reference-accessors`

### Domain: Android storage bridge

Android file and directory access handlers implement asset, filesystem, MediaStore, and SAF paths through JNI interoperation.

### Domain: Apple embedded hosting

Apple embedded hosting uses SwiftUI and Objective-C++ alongside display, keyboard, view-controller, text-to-speech, and Vulkan-context platform adapters.

### Implementation: C++ const-reference accessors

C++ member functions expose lookup results through const pointers and accept immutable String references to avoid copying lookup inputs.
Apply it through: OpenXRInteractionProfile.

## Unit 220

Lesson ID: `unit-audio-backend-adapters-cpp-resource-lifetime`

### Domain: Audio backend adapters

Audio output is provided by platform adapters selected with the driver build for ALSA, PulseAudio, CoreAudio, WASAPI, and XAudio2.

### Implementation: C++ constructor and destructor resource lifetime

A C++ class constructor and destructor control the Font's allocation and release of its cached advance array.
Apply it through: Font.

## Unit 221

Lesson ID: `unit-gdscript-static-analysis-gpu-texture-compression-midi-input-adapters`

### Domain: GDScript static analysis

The analyzer consumes parsed script trees to resolve inheritance, infer and check data types, validate annotations, and diagnose invalid expressions.
Use these entity examples: GDScriptParser::Node, GDScript.

### Domain: GPU block-compression dispatch

The Betsy compressor relies on GLSL compute shaders that fetch source texels or blocks and write compressed results to storage images.
Use these entity examples: Image.

### Domain: MIDI input adapters

MIDI input is supplied by platform adapters for ALSA MIDI, CoreMIDI, and WinMM.

## Unit 222

Lesson ID: `unit-shader-interface-analysis`

### Domain: Shader interface mapping and reflection

IO mapping and reflection use intermediate-tree traversal to expose a compiled shader interface as uniform, buffer, atomic-counter, and pipeline input/output entries.
Use these entity examples: TProgram, TObjectReflection.

## Unit 223

Lesson ID: `unit-android-instrumented-tests-apple-embedded-packaging-c-abi-record-and-dispatch`

### Domain: Android instrumented integration tests

Android instrumentation tests launch an app project that exercises plugins, signals, storage, and JavaClassWrapper conversions.

### Domain: Apple embedded packaging and signing

Apple embedded packaging adds Xcode project data, assets, Apple plugin configuration, and code-signing structures to a target-platform export.
Use these entity examples: EditorExportPlatform, EditorExportPlatform::ExportMessage.

### Implementation: C ABI records and dispatch signatures

`VkViSurfaceCreateInfoNN` and `PFN_vkCreateViSurfaceNN` express ABI-shaped call parameters through a tagged C record and a typed creation-function pointer.
Apply it through: VkViSurfaceCreateInfoNN, VkSurfaceKHR.

## Unit 224

Lesson ID: `unit-gdscript-bytecode-runtime-cxx-raii-reference-ownership`

### Domain: GDScript bytecode compilation and execution

The compiler lowers analyzed script trees through code-generation classes, and the VM executes functions using validated calls, getters, setters, and container operations.
Use these entity examples: GDScript, GDScriptInstance, GDScriptFunction.

### Implementation: RAII, non-copyable resources, and intrusive references

RAII-bound C++ objects manage locks and scopes, while Ref and RefTarget keep shared simulation objects alive without ordinary copying.
Apply it through: Shape, PhysicsScene.

## Unit 225

Lesson ID: `unit-gdscript-plugin-integration`

### Implementation: GDScript Android plugin integration

GDScript test code retrieves JNISingleton objects, calls JavaClassWrapper, and connects signals through Android plugin signals.
Apply it through: SignalInfo.

## Unit 226

Lesson ID: `unit-graphite-rule-execution`

### Domain: Graphite SILF rule execution

Graphite executes decoded SILF rule constraints and actions through a finite-state matcher and bytecode machine.
Use these entity examples: Silf, Segment, Slot.

## Unit 227

Lesson ID: `unit-graphite-shaping`

### Domain: Graphite segment shaping

Graphite uses decoded font tables and rule passes to turn Unicode into a Segment whose linked Slots carry glyph and placement state.
Use these entity examples: Face, Font, Segment, Slot, FeatureVal.

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
- `allocation` enables `random-generation` (synthesized lexical, confidence 0.25).
- `allocation` enables `zip-archive-io` (synthesized lexical, confidence 0.25).
- `alpha-plane-coding` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
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
- `block-texture-transcoding` enables `block-texture-encoding` (synthesized lexical, confidence 0.25).
- `broad-phase` enables `narrow-phase` (grounded, confidence 0.90).
- `brotli-stream-decompression` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
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
- `bvh-traversal` enables `scene-graph` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `scene-importing` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `scene-tree` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `scene-tree-execution` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `bvh-traversal` enables `visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `canvas-and-viewports` enables `gui-controls` (grounded, confidence 0.90).
- `canvas-and-viewports` enables `platform-presentation` (grounded, confidence 0.90).
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
- `concurrency` enables `encoder-configuration` (synthesized lexical, confidence 0.25).
- `concurrency` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `concurrency` enables `sdl-event-queue` (synthesized lexical, confidence 0.25).
- `concurrency` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `concurrency` enables `synchronization-primitives` (synthesized lexical, confidence 0.25).
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
- `diagnostic-expectations` enables `image-quality-metrics` (synthesized lexical, confidence 0.25).
- `dynamic-invocation-signals` enables `class-reference-generation` (synthesized lexical, confidence 0.25).
- `dynamic-invocation-signals` enables `deferred-execution` (grounded, confidence 0.90).
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
- `gui-control-layout` enables `ui-theming` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `version-control-integration` (synthesized lexical, confidence 0.25).
- `gui-control-layout` enables `websocket-frame-events` (synthesized lexical, confidence 0.25).
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
- `instancing` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `instancing` enables `hit-results` (synthesized lexical, confidence 0.25).
- `instancing` enables `rendering-assets` (synthesized lexical, confidence 0.25).
- `instancing` enables `rigid-body-motion` (synthesized lexical, confidence 0.25).
- `instancing` enables `scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `instancing` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `instancing` enables `vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `ktx2-container-transcoding` enables `ktx-texture-container` (synthesized lexical, confidence 0.25).
- `linuxbsd-desktop-integration` enables `wayland-window-backend` (grounded, confidence 0.90).
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
- `motion-blur` enables `rigid-body-motion` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `runtime-loop-time` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `sdl-platform-backends` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `motion-blur` enables `wraparound-network-time` (synthesized lexical, confidence 0.25).
- `narrow-phase` enables `contact-management` (grounded, confidence 0.90).
- `narrow-phase` enables `vehicle-dynamics` (grounded, confidence 0.90).
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
- `navigation-regions` enables `navigation-map-composition` (synthesized lexical, confidence 0.25).
- `navmesh-heightfields` enables `navmesh-contours-polygons` (grounded, confidence 0.90).
- `network-data-exchange` enables `webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `network-io` enables `webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `dynamic-invocation-signals` (grounded, confidence 0.90).
- `object-identity-lifetime` enables `navigation-queries` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `physics-server-queries` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `physics-space-queries` (synthesized lexical, confidence 0.25).
- `object-identity-lifetime` enables `reflection-metadata` (grounded, confidence 0.90).
- `object-identity-lifetime` enables `regular-expression-results` (synthesized lexical, confidence 0.25).
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
- `physics-3d-collision-pipeline` enables `three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `physics-collision` enables `two-dimensional-content` (synthesized lexical, confidence 0.25).
- `physics-space-queries` enables `physics-queries` (synthesized lexical, confidence 0.25).
- `physics-spaces` enables `physics-queries` (grounded, confidence 0.90).
- `platform-display-adaptation` enables `platform-presentation` (synthesized lexical, confidence 0.25).
- `png-stream-transforms` enables `png-image-codec` (synthesized lexical, confidence 0.25).
- `prefix-code-decoding` enables `brotli-stream-decompression` (grounded, confidence 0.90).
- `primitive-intersection` enables `filter-callbacks` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `primitive-intersection` enables `triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `broad-phase` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `builder-heuristics` (grounded, confidence 0.90).
- `primitive-references` enables `bvh-construction` (grounded, confidence 0.90).
- `primitive-references` enables `instancing` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `motion-blur` (synthesized lexical, confidence 0.25).
- `primitive-references` enables `spatial-indexing` (synthesized lexical, confidence 0.25).
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
- `reflection` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `reflection` enables `vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `reflection-metadata` enables `native-extensions` (synthesized lexical, confidence 0.25).
- `regex-compile-match` enables `regex-jit-codegen` (grounded, confidence 0.90).
- `rendering-assets` enables `collada-import` (synthesized lexical, confidence 0.25).
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
- `resource-assets` enables `gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `gpu-command-encoding` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `gpu-resources-pipelines` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `physics-collision` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `renderer-storage` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `rendering-assets` (synthesized lexical, confidence 0.25).
- `resource-assets` enables `rendering-resources` (grounded, confidence 0.90).
- `resource-assets` enables `resource-previews` (synthesized lexical, confidence 0.25).
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
- `runtime-loop-time` enables `motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `runtime-loop-time` enables `project-catalog` (synthesized lexical, confidence 0.25).
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
- `scene-commit` enables `project-catalog` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `property-introspection` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `resource-assets` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `runtime-configuration` (synthesized lexical, confidence 0.25).
- `scene-commit` enables `scene-authoring` (synthesized lexical, confidence 0.25).
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
- `scene-tree` enables `canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `collada-import` (synthesized lexical, confidence 0.25).
- `scene-tree` enables `expression-evaluation` (synthesized lexical, confidence 0.25).
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
- `serialization` enables `state-recording` (grounded, confidence 0.90).
- `sfnt-tables` enables `font-table-validation` (grounded, confidence 0.90).
- `sfnt-tables` enables `variable-font-data` (grounded, confidence 0.90).
- `shader-editing-and-language-plugins` enables `platform-selective-shader-baking` (grounded, confidence 0.90).
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
- `shader-programs` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `shader-reflection` enables `shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `shader-source-compilation` enables `shader-interface-analysis` (grounded, confidence 0.90).
- `signal-awaitability` enables `android-instrumented-tests` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `android-plugin-signals` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `godot-member-exposure` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `reflection` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `reflection-metadata` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `runtime-metadata` (synthesized lexical, confidence 0.25).
- `signal-awaitability` enables `scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `simd-accelerated-codecs` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `simd-accelerated-codecs` enables `simd-ray-packets` (synthesized lexical, confidence 0.25).
- `simd-ray-packets` enables `isa-portability` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `3d-gizmo-authoring` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `post-import-processing` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `skeletal-ik` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `skeleton-modification` (synthesized lexical, confidence 0.25).
- `skeletal-ragdoll` enables `skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `skeleton-modifiers` enables `skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
- `skeleton-modifiers` enables `skeleton-modification` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `allocation` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `builder-heuristics` (synthesized lexical, confidence 0.25).
- `spatial-indexing` enables `physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `spatial-predictive-filters` enables `collision-filtering` (synthesized lexical, confidence 0.25).
- `spatial-predictive-filters` enables `filter-callbacks` (synthesized lexical, confidence 0.25).
- `spirv-generation` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `spirv-generation` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-generation` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `spirv-generation` (synthesized lexical, confidence 0.25).
- `spirv-rewriting` enables `spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `spirv-shader-reflection` enables `shader-interface-metadata` (grounded, confidence 0.90).
- `spirv-shader-reflection` enables `shader-reflection` (synthesized lexical, confidence 0.25).
- `string-names-paths` enables `localization` (grounded, confidence 0.90).
- `string-names-paths` enables `profiling-interning` (grounded, confidence 0.90).
- `string-names-paths` enables `scene-replication-configuration` (synthesized lexical, confidence 0.25).
- `string-names-paths` enables `variant-text-serialization` (grounded, confidence 0.90).
- `subdivision-surface-evaluation` enables `feature-adaptive-tessellation` (grounded, confidence 0.90).
- `subdivision-surface-evaluation` enables `half-edge-topology` (synthesized lexical, confidence 0.25).
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
- `theora-video-streams` enables `theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `tile-authoring` enables `tile-libraries` (synthesized lexical, confidence 0.25).
- `tile-libraries` enables `tilemap-layer-data` (synthesized lexical, confidence 0.25).
- `tile-libraries` enables `two-dimensional-content` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `crypto-resources` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `network-io` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `tls-connection-state` enables `stream-networking` (synthesized lexical, confidence 0.25).
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
- `ui-theming` enables `themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `unicode-data` enables `break-iteration` (grounded, confidence 0.90).
- `unicode-data` enables `character-encoding-conversion` (grounded, confidence 0.90).
- `unicode-data` enables `identifier-spoof-checking` (grounded, confidence 0.90).
- `unicode-data` enables `unicode-normalization` (grounded, confidence 0.90).
- `upnp-device-control` enables `upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `upnp-device-discovery` enables `upnp-port-mapping` (grounded, confidence 0.90).
- `variable-font-subsetting` enables `variable-font-data` (synthesized lexical, confidence 0.25).
- `variant-containers` enables `break-iteration` (synthesized lexical, confidence 0.25).
- `variant-containers` enables `generic-containers` (grounded, confidence 0.90).
- `variant-containers` enables `packet-queries` (synthesized lexical, confidence 0.25).
- `variant-containers` enables `vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `vector-font-export` enables `generic-containers` (synthesized lexical, confidence 0.25).
- `vector-font-export` enables `multichannel-distance-fields` (synthesized lexical, confidence 0.25).
- `vector-font-export` enables `visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `vi-native-surface-creation` enables `vulkan-extensible-records` (synthesized lexical, confidence 0.25).
- `visual-shader-module-build` enables `visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `visual-shader-nodes` enables `visual-shader-vector-operations` (grounded, confidence 0.90).
- `vulkan-swapchain-presentation` enables `canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `vulkan-swapchain-presentation` enables `metal-presentation` (synthesized lexical, confidence 0.25).
- `wayland-window-backend` enables `explicit-drm-syncobj` (synthesized lexical, confidence 0.25).
- `webp-image-io` enables `image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `webp-image-io` enables `webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `webrtc-data-channels` enables `webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `webrtc-peer-connections` enables `webrtc-data-channels` (grounded, confidence 0.90).
- `webrtc-peer-connections` enables `webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `websocket-frame-events` enables `websocket-multiplayer` (synthesized lexical, confidence 0.25).
- `websocket-multiplayer` enables `networking` (synthesized lexical, confidence 0.25).
- `websocket-peers` enables `websocket-multiplayer` (grounded, confidence 0.90).
- `xr-tracking` enables `pluggable-server-backends` (synthesized lexical, confidence 0.25).
- `xr-tracking` enables `skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
- `zip-archive-io` enables `target-platform-export` (synthesized lexical, confidence 0.25).
- `zip-archive-io` enables `zip64-archive-io` (synthesized lexical, confidence 0.25).
- `zlib-stream-codec` enables `openexr-image-decoding` (synthesized lexical, confidence 0.25).
- Removed cycle edge `undo-redo-history → editor-scene-sessions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `shader-interface-analysis → shader-interface-metadata`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `image-decode-pipeline → brotli-stream-decompression`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `jpeg-decompression-transcoding → image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `jpeg-decompression-transcoding → raster-image-input`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `png-stream-transforms → image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `png-stream-transforms → raster-image-input`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `webp-riff-decoding → image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `webp-riff-decoding → webp-image-io`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `brotli-stream-decompression → prefix-code-decoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `prefix-code-decoding → histograms-and-huffman-codes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `dynamic-invocation-signals → dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `subdivision-surface-evaluation → geometry-families`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `catmull-clark-patch-construction → half-edge-topology`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gui-controls → gui-control-layout`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `scene-state → packed-scenes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `rendering-assets → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `block-texture-transcoding → astc-block-coding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `ktx-texture-container → image-format-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `ktx2-container-transcoding → texture-compression-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `dynamic-invocation-signals → signal-awaitability`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `object-identity-lifetime → object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `shader-reflection → shader-interface-metadata`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `shader-reflection → spirv-shader-reflection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `spirv-shader-reflection → spirv-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `crypto-resources → cryptography`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `diagnostic-expectations → device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `dynamic-variant → dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `native-extensions → gdextension-libraries`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `reflection-metadata → reflection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `object-identity-lifetime → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `upnp-port-mapping → upnp-device-discovery`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `allocation → object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `spatial-indexing → bvh-construction`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `spatial-indexing → bvh-traversal`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `spatial-indexing → motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `skeletal-ragdoll → gltf-node-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `physics-3d-collision-pipeline → convex-collision`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `spatial-indexing → scene-commit`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `themes-and-style-boxes → editor-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `ui-theming → editor-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `themes-and-style-boxes → ui-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gui-control-layout → ui-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `navigation-regions → navigation-agents`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `scene-tree → gltf-node-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `input-events-actions → input-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `bvh-traversal → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `geometry-families → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `motion-blur-geometry → motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `theora-block-video-codec → theora-video-streams`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `primitive-intersection → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `instancing → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `project-settings → editor-and-project-settings`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `variant-containers → dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `signed-distance-fields → glyph-rasterization`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `instancing → openxr-runtime-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `vulkan-instance-setup → openxr-runtime-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `motion-blur → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `runtime-loop-time → motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `runtime-loop-time → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `concurrency → device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `sdl-thread-abstractions → task-scheduling`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `scene-commit → primitive-references`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `resource-assets → device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `resource-assets → scene-contexts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `scene-commit → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `scene-commit → scene-contexts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `primitive-references → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `scene-contexts → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `reflection → editor-and-project-settings`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `alpha-plane-coding → encoder-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `alpha-plane-coding → input-picture-planes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `yuv-rgb-conversion → input-picture-planes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `unicode-normalization → unicode-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `tls-connection-state → tls-crypto-backend`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `synchronization-primitives → concurrency`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `concurrency → task-scheduling`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `wayland-window-backend → linuxbsd-desktop-integration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `object-model → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `openxr-runtime-integration → openxr-structure-chains`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `localization → text-and-localization`: removed to preserve the stronger inferred prerequisite hierarchy.

### implementation language
- `annotations-static-state-and-lifecycle` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cpp-runtime-casts` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cpp-static-generated-data` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `properties-and-accessors` (synthesized lexical, confidence 0.25).
- `annotations-static-state-and-lifecycle` enables `typed-collections` (synthesized lexical, confidence 0.25).
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
- `c-pointers-arrays-and-strides` enables `c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-resource-lifecycles` (grounded, confidence 0.90).
- `c-pointers-arrays-and-strides` enables `c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `c-pointers-arrays-and-strides` enables `c-structs-pointers` (synthesized lexical, confidence 0.25).
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
- `classes-inheritance-and-lookup` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `classes-inheritance-and-lookup` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `cpp-abstraction-and-polymorphism` enables `annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-byte-exact-binary-parsing` (grounded, confidence 0.90).
- `cpp-abstraction-and-polymorphism` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-abstraction-and-polymorphism` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `cpp-class-inheritance` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-class-inheritance` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-class-registration` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
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
- `cpp-class-templates` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-class-templates` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
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
- `cpp-classes-and-ref-handles` enables `cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `cpp-classes-and-ref-handles` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
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
- `cpp-classes-inheritance` enables `cpp-class-templates` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-copy-on-write-containers` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-godot-binding-macros` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cpp-runtime-polymorphism` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-scoped-locks` (grounded, confidence 0.90).
- `cpp-classes-inheritance` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-classes-inheritance` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-const-borrowing` enables `cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-copy-on-write-containers` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-engine-polymorphism` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-engine-polymorphism` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-engine-polymorphism` enables `objective-cpp-apple-adapters` (grounded, confidence 0.90).
- `cpp-function-pointer-interfaces` enables `c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-generic-containers` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-inheritance-interfaces` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-interface-polymorphism` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `cpp-numeric-value-operations` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-copy-on-write-containers` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-numeric-value-operations` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cpp-runtime-casts` (grounded, confidence 0.90).
- `cpp-object-hierarchies` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `cxx-object-model` (synthesized lexical, confidence 0.25).
- `cpp-object-hierarchies` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cpp-ownership-and-polymorphism` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-pluggable-allocation` enables `c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `cpp-polymorphic-backends` enables `cpp-virtual-dispatch` (synthesized lexical, confidence 0.25).
- `cpp-polymorphic-backends` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cpp-preprocessor-configuration` enables `cpp-static-abi-contracts` (grounded, confidence 0.90).
- `cpp-preprocessor-configuration` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `cpp-resource-lifetime` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-runtime-adapters` enables `cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `cpp-runtime-adapters` enables `cpp-jni-variant-marshalling` (grounded, confidence 0.90).
- `cpp-runtime-polymorphism` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cpp-scope-locking` enables `cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `cpp-scoped-locks` enables `cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `cpp-static-generated-data` enables `cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-template-binding` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-const-access` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-generic-containers` enables `cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-generic-containers` enables `cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-generic-containers` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `cpp-templates-and-specialization` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
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
- `cpp-templates-for-binary-data` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cxx-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-for-binary-data` enables `cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-specialized-marshalling` (grounded, confidence 0.90).
- `cpp-templates-traits` enables `cpp-template-binding` (synthesized lexical, confidence 0.25).
- `cpp-templates-traits` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
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
- `csharp-unsafe-interop` enables `c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `csharp-unsafe-interop` enables `cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `cxx-atomics` enables `c-concurrent-state` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-abi-bridging` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `c-structs-pointers` (synthesized lexical, confidence 0.25).
- `cxx-c-abi` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cxx-class-hierarchy` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `cxx-closures` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `cxx-conditional-compilation` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `cxx-object-model` enables `cxx-raii-reference-ownership` (grounded, confidence 0.90).
- `cxx-object-model` enables `cxx-reflection-macros` (grounded, confidence 0.90).
- `cxx-object-model` enables `java-android-host-api` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `cxx-preprocessor-configuration` enables `cxx-simd-alignment` (grounded, confidence 0.90).
- `cxx-preprocessor-configuration` enables `godot-shader-language` (synthesized lexical, confidence 0.25).
- `cxx-reflection-macros` enables `cxx-stream-serialization` (grounded, confidence 0.90).
- `cxx-templates` enables `cpp-class-templates` (synthesized lexical, confidence 0.25).
- `cxx-templates` enables `cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `cxx-templates` enables `cxx-wide-simd` (grounded, confidence 0.90).
- `cxx-templates` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `cxx-wide-simd` enables `c-platform-selection` (synthesized lexical, confidence 0.25).
- `cxx-wide-simd` enables `cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `cxx-wide-simd` enables `cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `gdscript-declarations` enables `gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `gdscript-declarations` enables `gdscript-signals-callables` (grounded, confidence 0.90).
- `gdscript-declarations` enables `types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `gdscript-declarations` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `gdscript-query-objects` enables `types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `cxx-closures` (synthesized lexical, confidence 0.25).
- `gdscript-signals-callables` enables `cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `gdscript-typed-collections` enables `cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `godot-shader-language` enables `cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `godot-shader-language` enables `cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `godot-shader-language` enables `glsl-compute-shaders` (synthesized lexical, confidence 0.25).
- `iteration-protocols` enables `cpp-overloading` (synthesized lexical, confidence 0.25).
- `javascript-web-runtime` enables `javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `objective-cpp-apple-adapters` enables `objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `properties-and-accessors` enables `csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-actions` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-code-generation` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-actions` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-hooks` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-build-scripts` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `python-scons-build-hooks` enables `gdscript-query-objects` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-actions` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-code-generation` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-hooks` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-scripts` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-build-source-generation` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `python-scons-configuration` enables `python-scons-module-configuration` (synthesized lexical, confidence 0.25).
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
- Removed cycle edge `c-abi-data-structures → c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cplusplus-export-callbacks → gdscript-signals-callables`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-closures → callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gdscript-signals-callables → callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `csharp-unsafe-interop → c-aggregate-callback-modules`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gdscript-signals-callables → csharp-partial-source-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-closures → gdscript-signals-callables`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-lambdas-standard-algorithms → callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gdscript-signals-callables → signals-and-await`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gdscript-query-objects → c-abi-manual-lifetime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gdscript-query-objects → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gdscript-declarations → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `gdscript-declarations → gdscript-query-objects`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-code-generation → python-build-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-code-generation → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-scripts → python-build-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-scripts → python-build-code-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-scripts → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-scons-build-hooks → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-scons-build-hooks → python-build-scripts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-scons-build-hooks → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-scons-build-hooks → python-scons-module-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-scripts → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-scons-module-configuration → python-build-scripts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-scons-module-configuration → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-code-generation → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-preprocessor-configuration → c-preprocessor-platform-selection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-preprocessor-configuration → cxx-preprocessor-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-const-borrowing → cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `c-preprocessor-platform-selection → cxx-preprocessor-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-engine-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-generic-containers → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-templates → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-template-binding → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-templates → cpp-template-binding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-templates → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-and-const-access → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-templates → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-and-specialization → cpp-class-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-and-specialization → cpp-generic-containers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-and-views → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-and-specialization → cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-for-binary-data → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-templates → cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-traits → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-templates → cpp-templates-traits`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-templates → cxx-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-preprocessor-configuration → c-platform-selection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-preprocessor-configuration → c-preprocessor-portability`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-template-binding → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-template-binding → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-template-binding → cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-template-binding → cpp-templates-traits`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-templates → cpp-generic-containers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-templates → cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-generic-containers → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-and-const-access → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-for-binary-data → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-for-binary-data → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-traits → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-traits → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates-traits → cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `objective-cpp-ios-adapters → objective-cpp-apple-adapters`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-abstraction-and-polymorphism → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-inheritance-interfaces → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-inheritance-interfaces → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cpp-inheritance-interfaces`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cpp-native-wrappers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-object-model → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-object-model → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-inheritance → cxx-object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-abstraction-and-polymorphism → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-abstraction-and-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-inheritance-interfaces → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-inheritance-interfaces → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-inheritance-interfaces → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-inheritance-interfaces → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-object-model → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-object-model → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-inheritance-interfaces → cxx-object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-abstraction-and-polymorphism → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-object-hierarchies → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-object-hierarchies → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-object-hierarchies → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-object-hierarchies → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-object-hierarchies → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-class-hierarchy → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-class-hierarchy → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-class-hierarchy → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-class-hierarchy → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-class-hierarchy → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-object-hierarchies → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-object-model → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-object-model → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-templates → cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-registration → cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-state-and-composition → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-numeric-value-operations → cpp-class-state-and-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-copy-on-write-containers → cpp-class-state-and-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-copy-on-write-containers → cpp-numeric-value-operations`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-classes → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-classes-and-ref-handles → cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-classes-and-ref-handles → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `classes-inheritance-and-lookup → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `classes-inheritance-and-lookup → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `classes-inheritance-and-lookup → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-classes-inheritance → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-classes-inheritance → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-classes-inheritance → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-runtime-polymorphism → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-interface-polymorphism → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-interface-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-interface-polymorphism → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `javascript-browser-ffi → javascript-web-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-classes-and-ref-handles → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-resource-and-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `c-pointers-arrays-and-strides → c-abi-manual-lifetime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cxx-c-abi → c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-actions → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-actions → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-class-registration → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `cpp-basis-transcoding → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `c-pointers-arrays-and-strides → c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `python-build-hooks → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.

### curriculum
- `domain/3d-gizmo-authoring` enables `domain/editor-plugin-extension` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/buffer-storage` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/class-reference-generation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/godot-thirdparty-adaptation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/gpu-memory-suballocation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/histograms-and-huffman-codes` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/object-identity-lifetime` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/random-generation` (synthesized lexical, confidence 0.25).
- `domain/allocation` enables `domain/zip-archive-io` (synthesized lexical, confidence 0.25).
- `domain/alpha-plane-coding` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
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
- `domain/block-texture-transcoding` enables `domain/block-texture-encoding` (synthesized lexical, confidence 0.25).
- `domain/broad-phase` enables `domain/narrow-phase` (grounded, confidence 0.90).
- `domain/brotli-stream-decompression` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
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
- `domain/bvh-traversal` enables `domain/scene-graph` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/scene-importing` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/scene-tree` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/scene-tree-execution` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/visual-shader-module-build` (synthesized lexical, confidence 0.25).
- `domain/bvh-traversal` enables `domain/visual-shader-nodes` (synthesized lexical, confidence 0.25).
- `domain/canvas-and-viewports` enables `domain/gui-controls` (grounded, confidence 0.90).
- `domain/canvas-and-viewports` enables `domain/platform-presentation` (grounded, confidence 0.90).
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
- `domain/concurrency` enables `domain/encoder-configuration` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/sdl-event-queue` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/concurrency` enables `domain/synchronization-primitives` (synthesized lexical, confidence 0.25).
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
- `domain/diagnostic-expectations` enables `domain/image-quality-metrics` (synthesized lexical, confidence 0.25).
- `domain/diagnostic-expectations` enables `language/warnings-and-suppression` (asserted, confidence 0.30).
- `domain/dynamic-invocation-signals` enables `domain/class-reference-generation` (synthesized lexical, confidence 0.25).
- `domain/dynamic-invocation-signals` enables `domain/deferred-execution` (grounded, confidence 0.90).
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
- `domain/gui-control-layout` enables `domain/ui-theming` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/version-control-integration` (synthesized lexical, confidence 0.25).
- `domain/gui-control-layout` enables `domain/websocket-frame-events` (synthesized lexical, confidence 0.25).
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
- `domain/instancing` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/hit-results` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/rendering-assets` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/rigid-body-motion` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/scene-data-and-buffers` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/instancing` enables `domain/vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `domain/ios-embedded-adapter` enables `language/objective-cpp-ios-adapters` (grounded, confidence 0.90).
- `domain/ktx2-container-transcoding` enables `domain/ktx-texture-container` (synthesized lexical, confidence 0.25).
- `domain/linuxbsd-desktop-integration` enables `domain/wayland-window-backend` (grounded, confidence 0.90).
- `domain/linuxbsd-desktop-integration` enables `language/c-dynamic-library-wrappers` (grounded, confidence 0.90).
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
- `domain/motion-blur` enables `domain/rigid-body-motion` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/runtime-loop-time` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/sdl-platform-backends` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/motion-blur` enables `domain/wraparound-network-time` (synthesized lexical, confidence 0.25).
- `domain/narrow-phase` enables `domain/contact-management` (grounded, confidence 0.90).
- `domain/narrow-phase` enables `domain/vehicle-dynamics` (grounded, confidence 0.90).
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
- `domain/navigation-regions` enables `domain/navigation-map-composition` (synthesized lexical, confidence 0.25).
- `domain/navmesh-heightfields` enables `domain/navmesh-contours-polygons` (grounded, confidence 0.90).
- `domain/network-data-exchange` enables `domain/webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `domain/network-io` enables `domain/webrtc-data-channels` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/dynamic-invocation-signals` (grounded, confidence 0.90).
- `domain/object-identity-lifetime` enables `domain/navigation-queries` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/physics-server-queries` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/physics-space-queries` (synthesized lexical, confidence 0.25).
- `domain/object-identity-lifetime` enables `domain/reflection-metadata` (grounded, confidence 0.90).
- `domain/object-identity-lifetime` enables `domain/regular-expression-results` (synthesized lexical, confidence 0.25).
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
- `domain/physics-3d-collision-pipeline` enables `domain/three-dimensional-physics` (synthesized lexical, confidence 0.25).
- `domain/physics-collision` enables `domain/two-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/physics-queries` enables `language/gdscript-query-objects` (grounded, confidence 0.90).
- `domain/physics-space-queries` enables `domain/physics-queries` (synthesized lexical, confidence 0.25).
- `domain/physics-spaces` enables `domain/physics-queries` (grounded, confidence 0.90).
- `domain/platform-display-adaptation` enables `domain/platform-presentation` (synthesized lexical, confidence 0.25).
- `domain/png-stream-transforms` enables `domain/png-image-codec` (synthesized lexical, confidence 0.25).
- `domain/prefix-code-decoding` enables `domain/brotli-stream-decompression` (grounded, confidence 0.90).
- `domain/primitive-intersection` enables `domain/filter-callbacks` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/ray-primitive-intersection` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/primitive-intersection` enables `domain/triangle-intersection-algorithms` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/broad-phase` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/builder-heuristics` (grounded, confidence 0.90).
- `domain/primitive-references` enables `domain/bvh-construction` (grounded, confidence 0.90).
- `domain/primitive-references` enables `domain/instancing` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/motion-blur` (synthesized lexical, confidence 0.25).
- `domain/primitive-references` enables `domain/spatial-indexing` (synthesized lexical, confidence 0.25).
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
- `domain/reflection` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `domain/vulkan-instance-setup` (synthesized lexical, confidence 0.25).
- `domain/reflection` enables `language/cpp-template-binding` (grounded, confidence 0.90).
- `domain/reflection-metadata` enables `domain/native-extensions` (synthesized lexical, confidence 0.25).
- `domain/reflection-metadata` enables `language/cpp-specialized-marshalling` (grounded, confidence 0.90).
- `domain/regex-compile-match` enables `domain/regex-jit-codegen` (grounded, confidence 0.90).
- `domain/rendering-assets` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
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
- `domain/resource-assets` enables `domain/gltf-materials-textures` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/gpu-command-encoding` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/gpu-resources-pipelines` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/physics-collision` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/renderer-storage` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/rendering-assets` (synthesized lexical, confidence 0.25).
- `domain/resource-assets` enables `domain/rendering-resources` (grounded, confidence 0.90).
- `domain/resource-assets` enables `domain/resource-previews` (synthesized lexical, confidence 0.25).
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
- `domain/runtime-loop-time` enables `domain/motion-blur-geometry` (synthesized lexical, confidence 0.25).
- `domain/runtime-loop-time` enables `domain/project-catalog` (synthesized lexical, confidence 0.25).
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
- `domain/scene-commit` enables `domain/project-catalog` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/property-introspection` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/raycast-occlusion-culling` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/replicated-property-synchronization` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/resource-assets` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/runtime-configuration` (synthesized lexical, confidence 0.25).
- `domain/scene-commit` enables `domain/scene-authoring` (synthesized lexical, confidence 0.25).
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
- `domain/scene-tree` enables `domain/canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/collada-import` (synthesized lexical, confidence 0.25).
- `domain/scene-tree` enables `domain/expression-evaluation` (synthesized lexical, confidence 0.25).
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
- `domain/serialization` enables `domain/state-recording` (grounded, confidence 0.90).
- `domain/sfnt-tables` enables `domain/font-table-validation` (grounded, confidence 0.90).
- `domain/sfnt-tables` enables `domain/variable-font-data` (grounded, confidence 0.90).
- `domain/shader-editing-and-language-plugins` enables `domain/platform-selective-shader-baking` (grounded, confidence 0.90).
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
- `domain/shader-programs` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/shader-programs` enables `language/godot-shader-language` (grounded, confidence 0.90).
- `domain/shader-reflection` enables `domain/shader-interface-analysis` (synthesized lexical, confidence 0.25).
- `domain/shader-source-compilation` enables `domain/shader-interface-analysis` (grounded, confidence 0.90).
- `domain/signal-awaitability` enables `domain/android-instrumented-tests` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/android-plugin-signals` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/godot-member-exposure` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/managed-csharp-scripting` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/reflection` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/reflection-metadata` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/runtime-metadata` (synthesized lexical, confidence 0.25).
- `domain/signal-awaitability` enables `domain/scene-tree-and-signal-connections` (synthesized lexical, confidence 0.25).
- `domain/simd-accelerated-codecs` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/simd-accelerated-codecs` enables `domain/simd-ray-packets` (synthesized lexical, confidence 0.25).
- `domain/simd-ray-packets` enables `domain/isa-portability` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/3d-gizmo-authoring` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/post-import-processing` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/skeletal-ik` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/skeleton-modification` (synthesized lexical, confidence 0.25).
- `domain/skeletal-ragdoll` enables `domain/skeleton-modifiers` (synthesized lexical, confidence 0.25).
- `domain/skeleton-modifiers` enables `domain/skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
- `domain/skeleton-modifiers` enables `domain/skeleton-modification` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/allocation` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/builder-heuristics` (synthesized lexical, confidence 0.25).
- `domain/spatial-indexing` enables `domain/physics-2d-collision-pipeline` (synthesized lexical, confidence 0.25).
- `domain/spatial-predictive-filters` enables `domain/collision-filtering` (synthesized lexical, confidence 0.25).
- `domain/spatial-predictive-filters` enables `domain/filter-callbacks` (synthesized lexical, confidence 0.25).
- `domain/spirv-generation` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/spirv-generation` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-generation` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/shader-interface-metadata` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/spirv-generation` (synthesized lexical, confidence 0.25).
- `domain/spirv-rewriting` enables `domain/spirv-shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/spirv-shader-reflection` enables `domain/shader-interface-metadata` (grounded, confidence 0.90).
- `domain/spirv-shader-reflection` enables `domain/shader-reflection` (synthesized lexical, confidence 0.25).
- `domain/string-names-paths` enables `domain/localization` (grounded, confidence 0.90).
- `domain/string-names-paths` enables `domain/profiling-interning` (grounded, confidence 0.90).
- `domain/string-names-paths` enables `domain/scene-replication-configuration` (synthesized lexical, confidence 0.25).
- `domain/string-names-paths` enables `domain/variant-text-serialization` (grounded, confidence 0.90).
- `domain/subdivision-surface-evaluation` enables `domain/feature-adaptive-tessellation` (grounded, confidence 0.90).
- `domain/subdivision-surface-evaluation` enables `domain/half-edge-topology` (synthesized lexical, confidence 0.25).
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
- `domain/theora-video-streams` enables `domain/theora-block-video-codec` (synthesized lexical, confidence 0.25).
- `domain/tile-authoring` enables `domain/tile-libraries` (synthesized lexical, confidence 0.25).
- `domain/tile-libraries` enables `domain/tilemap-layer-data` (synthesized lexical, confidence 0.25).
- `domain/tile-libraries` enables `domain/two-dimensional-content` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/crypto-resources` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/network-io` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/sdl-thread-abstractions` (synthesized lexical, confidence 0.25).
- `domain/tls-connection-state` enables `domain/stream-networking` (synthesized lexical, confidence 0.25).
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
- `domain/ui-theming` enables `domain/themes-and-style-boxes` (synthesized lexical, confidence 0.25).
- `domain/unicode-data` enables `domain/break-iteration` (grounded, confidence 0.90).
- `domain/unicode-data` enables `domain/character-encoding-conversion` (grounded, confidence 0.90).
- `domain/unicode-data` enables `domain/identifier-spoof-checking` (grounded, confidence 0.90).
- `domain/unicode-data` enables `domain/unicode-normalization` (grounded, confidence 0.90).
- `domain/upnp-device-control` enables `domain/upnp-device-discovery` (synthesized lexical, confidence 0.25).
- `domain/upnp-device-discovery` enables `domain/upnp-port-mapping` (grounded, confidence 0.90).
- `domain/variable-font-subsetting` enables `domain/variable-font-data` (synthesized lexical, confidence 0.25).
- `domain/variant-containers` enables `domain/break-iteration` (synthesized lexical, confidence 0.25).
- `domain/variant-containers` enables `domain/generic-containers` (grounded, confidence 0.90).
- `domain/variant-containers` enables `domain/packet-queries` (synthesized lexical, confidence 0.25).
- `domain/variant-containers` enables `domain/vulkan-swapchain-presentation` (synthesized lexical, confidence 0.25).
- `domain/vector-font-export` enables `domain/generic-containers` (synthesized lexical, confidence 0.25).
- `domain/vector-font-export` enables `domain/multichannel-distance-fields` (synthesized lexical, confidence 0.25).
- `domain/vector-font-export` enables `domain/visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `domain/vi-native-surface-creation` enables `domain/vulkan-extensible-records` (synthesized lexical, confidence 0.25).
- `domain/visual-shader-module-build` enables `domain/visual-shader-vector-operations` (synthesized lexical, confidence 0.25).
- `domain/visual-shader-nodes` enables `domain/visual-shader-vector-operations` (grounded, confidence 0.90).
- `domain/vulkan-swapchain-presentation` enables `domain/canvas-and-viewports` (synthesized lexical, confidence 0.25).
- `domain/vulkan-swapchain-presentation` enables `domain/metal-presentation` (synthesized lexical, confidence 0.25).
- `domain/wayland-window-backend` enables `domain/explicit-drm-syncobj` (synthesized lexical, confidence 0.25).
- `domain/webp-image-io` enables `domain/image-decode-pipeline` (synthesized lexical, confidence 0.25).
- `domain/webp-image-io` enables `domain/webp-riff-decoding` (synthesized lexical, confidence 0.25).
- `domain/webrtc-data-channels` enables `domain/webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `domain/webrtc-peer-connections` enables `domain/webrtc-data-channels` (grounded, confidence 0.90).
- `domain/webrtc-peer-connections` enables `domain/webrtc-multiplayer-mesh` (grounded, confidence 0.90).
- `domain/websocket-frame-events` enables `domain/websocket-multiplayer` (synthesized lexical, confidence 0.25).
- `domain/websocket-multiplayer` enables `domain/networking` (synthesized lexical, confidence 0.25).
- `domain/websocket-peers` enables `domain/websocket-multiplayer` (grounded, confidence 0.90).
- `domain/xr-tracking` enables `domain/pluggable-server-backends` (synthesized lexical, confidence 0.25).
- `domain/xr-tracking` enables `domain/skeletal-modifiers-and-ik` (synthesized lexical, confidence 0.25).
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
- `language/c-pointers-arrays-and-strides` enables `language/c-aggregate-callback-modules` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-resource-lifecycles` (grounded, confidence 0.90).
- `language/c-pointers-arrays-and-strides` enables `language/c-stateful-struct-apis` (synthesized lexical, confidence 0.25).
- `language/c-pointers-arrays-and-strides` enables `language/c-structs-pointers` (synthesized lexical, confidence 0.25).
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
- `language/classes-inheritance-and-lookup` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/classes-inheritance-and-lookup` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `language/cpp-abstraction-and-polymorphism` enables `domain/gpu-memory-suballocation` (grounded, confidence 0.90).
- `language/cpp-abstraction-and-polymorphism` enables `language/annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-byte-exact-binary-parsing` (grounded, confidence 0.90).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-abstraction-and-polymorphism` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `language/cpp-class-inheritance` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-class-inheritance` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cmake-native-source-index` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cplusplus-enumerated-export-state` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-class-registration` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
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
- `language/cpp-class-templates` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-class-templates` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
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
- `language/cpp-classes-and-ref-handles` enables `language/cpp-binary-data-codecs` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-byte-exact-binary-parsing` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-c-type-mapping` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-and-ref-handles` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
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
- `language/cpp-classes-inheritance` enables `language/cpp-class-templates` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-copy-on-write-containers` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-godot-binding-macros` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cpp-runtime-polymorphism` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-scoped-locks` (grounded, confidence 0.90).
- `language/cpp-classes-inheritance` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cxx-class-hierarchy` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-classes-inheritance` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-const-borrowing` enables `language/cpp-const-reference-accessors` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-copy-on-write-containers` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-engine-polymorphism` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-engine-polymorphism` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-engine-polymorphism` enables `language/objective-cpp-apple-adapters` (grounded, confidence 0.90).
- `language/cpp-function-pointer-interfaces` enables `language/c-abi-record-and-dispatch` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-generic-containers` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-inheritance-interfaces` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `domain/enet-transport-and-multiplayer` (grounded, confidence 0.90).
- `language/cpp-interface-polymorphism` enables `domain/extensions` (grounded, confidence 0.90).
- `language/cpp-interface-polymorphism` enables `domain/gdscript-language-server` (asserted, confidence 0.30).
- `language/cpp-interface-polymorphism` enables `language/classes-inheritance-and-lookup` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-abstraction-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-interface-polymorphism` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `language/cpp-object-hierarchies` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cpp-runtime-casts` (grounded, confidence 0.90).
- `language/cpp-object-hierarchies` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/cxx-object-model` (synthesized lexical, confidence 0.25).
- `language/cpp-object-hierarchies` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cpp-ownership-and-polymorphism` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-pluggable-allocation` enables `language/c-function-pointer-callbacks` (synthesized lexical, confidence 0.25).
- `language/cpp-plugin-specialization` enables `domain/editor-plugin-lifecycle` (grounded, confidence 0.90).
- `language/cpp-polymorphic-backends` enables `language/cpp-virtual-dispatch` (synthesized lexical, confidence 0.25).
- `language/cpp-polymorphic-backends` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cpp-preprocessor-configuration` enables `language/cpp-static-abi-contracts` (grounded, confidence 0.90).
- `language/cpp-preprocessor-configuration` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/cpp-resource-lifetime` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-runtime-adapters` enables `language/cpp-classes-inheritance` (synthesized lexical, confidence 0.25).
- `language/cpp-runtime-adapters` enables `language/cpp-jni-variant-marshalling` (grounded, confidence 0.90).
- `language/cpp-runtime-polymorphism` enables `domain/editor-plugin-extension` (grounded, confidence 0.90).
- `language/cpp-runtime-polymorphism` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cpp-scope-locking` enables `language/cpp-scoped-locks` (synthesized lexical, confidence 0.25).
- `language/cpp-scoped-locks` enables `language/cxx-raii-reference-ownership` (synthesized lexical, confidence 0.25).
- `language/cpp-static-generated-data` enables `language/cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-template-binding` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-templates-and-const-access` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-templates-for-binary-data` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-const-access` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-generic-containers` enables `language/cpp-generic-containers` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-generic-containers` enables `language/cpp-templates-and-views` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-generic-containers` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-and-specialization` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
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
- `language/cpp-templates-for-binary-data` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-specialized-marshalling` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-templates-traits` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cpp-variadic-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cxx-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-for-binary-data` enables `language/cxx-wide-simd` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-ownership-and-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-runtime-symbol-loading` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-specialized-marshalling` (grounded, confidence 0.90).
- `language/cpp-templates-traits` enables `language/cpp-template-binding` (synthesized lexical, confidence 0.25).
- `language/cpp-templates-traits` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
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
- `language/csharp-unsafe-interop` enables `language/c-function-pointer-interfaces` (synthesized lexical, confidence 0.25).
- `language/csharp-unsafe-interop` enables `language/cpp-pluggable-allocation` (synthesized lexical, confidence 0.25).
- `language/cxx-atomics` enables `language/c-concurrent-state` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-abi-bridging` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-abi-data-structures` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/c-structs-pointers` (synthesized lexical, confidence 0.25).
- `language/cxx-c-abi` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/annotations-static-state-and-lifecycle` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-c-linkage-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cxx-class-hierarchy` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
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
- `language/cxx-closures` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `domain/isa-portability` (grounded, confidence 0.90).
- `language/cxx-conditional-compilation` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cpp-static-abi-contracts` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cxx-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/cxx-reflection-macros` (synthesized lexical, confidence 0.25).
- `language/cxx-conditional-compilation` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cplusplus-polymorphic-platform-adapters` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-class-inheritance` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-engine-polymorphism` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-inheritance-interfaces` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-native-wrappers` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-polymorphic-ownership` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/cxx-object-model` enables `language/cxx-raii-reference-ownership` (grounded, confidence 0.90).
- `language/cxx-object-model` enables `language/cxx-reflection-macros` (grounded, confidence 0.90).
- `language/cxx-object-model` enables `language/java-android-host-api` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/c-compatible-header-guards` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/c-preprocessor-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/cpp-preprocessor-configuration` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/cpp-single-header-implementation` (synthesized lexical, confidence 0.25).
- `language/cxx-preprocessor-configuration` enables `language/cxx-simd-alignment` (grounded, confidence 0.90).
- `language/cxx-preprocessor-configuration` enables `language/godot-shader-language` (synthesized lexical, confidence 0.25).
- `language/cxx-reflection-macros` enables `domain/serialization` (grounded, confidence 0.90).
- `language/cxx-reflection-macros` enables `language/cxx-stream-serialization` (grounded, confidence 0.90).
- `language/cxx-templates` enables `language/cpp-class-templates` (synthesized lexical, confidence 0.25).
- `language/cxx-templates` enables `language/cpp-templates-and-specialization` (synthesized lexical, confidence 0.25).
- `language/cxx-templates` enables `language/cxx-wide-simd` (grounded, confidence 0.90).
- `language/cxx-templates` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/cxx-wide-simd` enables `language/c-platform-selection` (synthesized lexical, confidence 0.25).
- `language/cxx-wide-simd` enables `language/cpp-simd-intrinsics` (synthesized lexical, confidence 0.25).
- `language/cxx-wide-simd` enables `language/cxx-simd-alignment` (synthesized lexical, confidence 0.25).
- `language/gdscript-declarations` enables `language/gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `language/gdscript-declarations` enables `language/gdscript-signals-callables` (grounded, confidence 0.90).
- `language/gdscript-declarations` enables `language/types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/gdscript-declarations` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/gdscript-plugin-integration` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/gdscript-signals-callables` (synthesized lexical, confidence 0.25).
- `language/gdscript-query-objects` enables `language/types-inference-and-conversions` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/cplusplus-export-callbacks` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/cxx-closures` (synthesized lexical, confidence 0.25).
- `language/gdscript-signals-callables` enables `language/cxx-lambdas-standard-algorithms` (synthesized lexical, confidence 0.25).
- `language/gdscript-typed-collections` enables `domain/gdscript-static-analysis` (asserted, confidence 0.30).
- `language/gdscript-typed-collections` enables `language/cpp-resource-lifetime` (synthesized lexical, confidence 0.25).
- `language/glsl-compute-shaders` enables `domain/gpu-texture-compression` (grounded, confidence 0.90).
- `language/godot-shader-language` enables `language/cpp-const-borrowing` (synthesized lexical, confidence 0.25).
- `language/godot-shader-language` enables `language/cpp-visitor-traversal` (synthesized lexical, confidence 0.25).
- `language/godot-shader-language` enables `language/glsl-compute-shaders` (synthesized lexical, confidence 0.25).
- `language/iteration-protocols` enables `language/cpp-overloading` (synthesized lexical, confidence 0.25).
- `language/javascript-web-runtime` enables `domain/browser-runtime-adapters` (grounded, confidence 0.90).
- `language/javascript-web-runtime` enables `domain/web-javascript-bridge` (grounded, confidence 0.90).
- `language/javascript-web-runtime` enables `language/javascript-browser-ffi` (synthesized lexical, confidence 0.25).
- `language/objective-cpp-apple-adapters` enables `language/objective-cpp-ios-adapters` (synthesized lexical, confidence 0.25).
- `language/properties-and-accessors` enables `language/csharp-partial-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-actions` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-code-generation` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-actions` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-hooks` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-build-scripts` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
- `language/python-scons-build-hooks` enables `language/gdscript-query-objects` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `domain/module-build-configuration` (grounded, confidence 0.90).
- `language/python-scons-configuration` enables `language/python-build-actions` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-code-generation` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-scripts` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-build-source-generation` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-scons-build-hooks` (synthesized lexical, confidence 0.25).
- `language/python-scons-configuration` enables `language/python-scons-module-configuration` (synthesized lexical, confidence 0.25).
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
- Removed cycle edge `domain/undo-redo-history → domain/editor-scene-sessions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/shader-interface-analysis → domain/shader-interface-metadata`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/image-decode-pipeline → domain/brotli-stream-decompression`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/jpeg-decompression-transcoding → domain/image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/jpeg-decompression-transcoding → domain/raster-image-input`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/png-stream-transforms → domain/image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/png-stream-transforms → domain/raster-image-input`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/webp-riff-decoding → domain/image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/webp-riff-decoding → domain/webp-image-io`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/brotli-stream-decompression → domain/prefix-code-decoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/prefix-code-decoding → domain/histograms-and-huffman-codes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/dynamic-invocation-signals → domain/dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/subdivision-surface-evaluation → domain/geometry-families`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/catmull-clark-patch-construction → domain/half-edge-topology`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/gui-controls → domain/gui-control-layout`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/scene-state → domain/packed-scenes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/rendering-assets → domain/geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/block-texture-transcoding → domain/astc-block-coding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/ktx-texture-container → domain/image-format-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/ktx2-container-transcoding → domain/texture-compression-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/dynamic-invocation-signals → domain/signal-awaitability`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/object-identity-lifetime → domain/object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/shader-reflection → domain/shader-interface-metadata`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/shader-reflection → domain/spirv-shader-reflection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/spirv-shader-reflection → domain/spirv-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/c-abi-data-structures → language/c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cplusplus-export-callbacks → language/gdscript-signals-callables`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-closures → language/callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/gdscript-signals-callables → language/callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/csharp-unsafe-interop → language/c-aggregate-callback-modules`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/gdscript-signals-callables → language/csharp-partial-source-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-closures → language/gdscript-signals-callables`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-lambdas-standard-algorithms → language/callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/gdscript-signals-callables → language/signals-and-await`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/crypto-resources → domain/cryptography`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/diagnostic-expectations → domain/device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/dynamic-variant → domain/dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/gdscript-query-objects → language/c-abi-manual-lifetime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/gdscript-query-objects → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/gdscript-declarations → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/gdscript-declarations → language/gdscript-query-objects`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/native-extensions → domain/gdextension-libraries`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/reflection-metadata → domain/reflection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/object-identity-lifetime → domain/resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/upnp-port-mapping → domain/upnp-device-discovery`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/allocation → domain/object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/spatial-indexing → domain/bvh-construction`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/spatial-indexing → domain/bvh-traversal`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/spatial-indexing → domain/motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/skeletal-ragdoll → domain/gltf-node-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/physics-3d-collision-pipeline → domain/convex-collision`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/spatial-indexing → domain/scene-commit`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/themes-and-style-boxes → domain/editor-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/ui-theming → domain/editor-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/themes-and-style-boxes → domain/ui-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/gui-control-layout → domain/ui-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/navigation-regions → domain/navigation-agents`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/scene-tree → domain/gltf-node-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/input-events-actions → domain/input-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/bvh-traversal → domain/ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/geometry-families → domain/geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/motion-blur-geometry → domain/motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/theora-block-video-codec → domain/theora-video-streams`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/primitive-intersection → domain/ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/instancing → domain/geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/project-settings → domain/editor-and-project-settings`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/variant-containers → domain/dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/signed-distance-fields → domain/glyph-rasterization`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/instancing → domain/openxr-runtime-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/vulkan-instance-setup → domain/openxr-runtime-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/motion-blur → domain/ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/runtime-loop-time → domain/motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/runtime-loop-time → domain/ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/concurrency → domain/device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/sdl-thread-abstractions → domain/task-scheduling`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/scene-commit → domain/primitive-references`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/resource-assets → domain/device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/resource-assets → domain/scene-contexts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/scene-commit → domain/resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/scene-commit → domain/scene-contexts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/primitive-references → domain/geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/scene-contexts → domain/resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/reflection → domain/editor-and-project-settings`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-code-generation → language/python-build-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-code-generation → language/python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-scripts → language/python-build-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-scripts → language/python-build-code-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-scripts → language/python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-scons-build-hooks → language/python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-scons-build-hooks → language/python-build-scripts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-scons-build-hooks → language/python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-scons-build-hooks → language/python-scons-module-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-scripts → language/python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-scons-module-configuration → language/python-build-scripts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-scons-module-configuration → language/python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-code-generation → language/python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/alpha-plane-coding → domain/encoder-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/alpha-plane-coding → domain/input-picture-planes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/yuv-rgb-conversion → domain/input-picture-planes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/unicode-normalization → domain/unicode-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-preprocessor-configuration → language/c-preprocessor-platform-selection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-preprocessor-configuration → language/cxx-preprocessor-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-const-borrowing → language/cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/c-preprocessor-platform-selection → language/cxx-preprocessor-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-engine-polymorphism → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-generic-containers → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-templates → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-template-binding → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-templates → language/cpp-template-binding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-templates → language/cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-and-const-access → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-templates → language/cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-and-specialization → language/cpp-class-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-and-specialization → language/cpp-generic-containers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-and-views → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-and-specialization → language/cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-for-binary-data → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-templates → language/cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-traits → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-templates → language/cpp-templates-traits`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-templates → language/cxx-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-preprocessor-configuration → language/c-platform-selection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-preprocessor-configuration → language/c-preprocessor-portability`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-template-binding → language/cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-template-binding → language/cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-template-binding → language/cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-template-binding → language/cpp-templates-traits`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-templates → language/cpp-generic-containers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-templates → language/cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/tls-connection-state → domain/tls-crypto-backend`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/synchronization-primitives → domain/concurrency`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/concurrency → domain/task-scheduling`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/wayland-window-backend → domain/linuxbsd-desktop-integration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/object-model → domain/resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/openxr-runtime-integration → domain/openxr-structure-chains`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-generic-containers → language/cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-and-const-access → language/cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-for-binary-data → language/cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-for-binary-data → language/cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-traits → language/cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-traits → language/cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates-traits → language/cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/objective-cpp-ios-adapters → language/objective-cpp-apple-adapters`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-abstraction-and-polymorphism → language/classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-inheritance-interfaces → language/classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-inheritance-interfaces → language/cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cpp-inheritance-interfaces`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cpp-native-wrappers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-object-model → language/classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-object-model → language/cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-inheritance → language/cxx-object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-abstraction-and-polymorphism → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-abstraction-and-polymorphism → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-inheritance-interfaces → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-inheritance-interfaces → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-inheritance-interfaces → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-inheritance-interfaces → language/cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-object-model → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-object-model → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-inheritance-interfaces → language/cxx-object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-abstraction-and-polymorphism → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-object-hierarchies → language/classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-object-hierarchies → language/cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-object-hierarchies → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-object-hierarchies → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-object-hierarchies → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-class-hierarchy → language/classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-class-hierarchy → language/cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-class-hierarchy → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-class-hierarchy → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-class-hierarchy → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-object-hierarchies → language/cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-object-model → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-object-model → language/cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-templates → language/cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-registration → language/cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-state-and-composition → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-numeric-value-operations → language/cpp-class-state-and-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-copy-on-write-containers → language/cpp-class-state-and-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-copy-on-write-containers → language/cpp-numeric-value-operations`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-classes → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-classes-and-ref-handles → language/cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-classes-and-ref-handles → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/classes-inheritance-and-lookup → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/classes-inheritance-and-lookup → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/classes-inheritance-and-lookup → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-classes-inheritance → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-classes-inheritance → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-classes-inheritance → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-runtime-polymorphism → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-interface-polymorphism → language/cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-interface-polymorphism → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-interface-polymorphism → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/javascript-browser-ffi → language/javascript-web-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-classes-and-ref-handles → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-resource-and-polymorphism → language/cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/c-pointers-arrays-and-strides → language/c-abi-manual-lifetime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cxx-c-abi → language/c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-actions → language/python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-actions → language/python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `domain/localization → domain/text-and-localization`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-class-registration → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/cpp-basis-transcoding → language/cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/c-pointers-arrays-and-strides → language/c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed cycle edge `language/python-build-hooks → language/python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
<!-- rope-ladder:end document -->
