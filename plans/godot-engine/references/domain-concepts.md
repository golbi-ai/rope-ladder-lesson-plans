<!-- rope-ladder:begin document f89297358bb5abd31193c228ae0f7ecfa8615ee899ec971419322dc7dd61435d -->
# Domain concepts

Topics are ordered by prerequisite depth. Each repository claim links to source evidence; general facts require external evidence.

## AABB tree construction

This is a geometry-acceleration structure within the supplied Jolt partition.

Jolt's AABBTreeBuilder stores nodes and indexed triangles and accepts a maximum-triangles-per-leaf setting while building its tree.

Code examples: AABBTreeBuilder

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.h` — AABBTreeBuilder
- Code: `thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.cpp` — AABBTreeBuilder

## Remote Android engine service

The service defines numeric message codes and bundle keys, while a fragment binds to and receives remote engine state.

Remote engine service messages convey engine status, errors, surface packages, dimensions, host tokens, and command-line parameters between Android UI and service code.

Code examples: GodotService.EngineStatus

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt` — GodotService
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt` — KEY_ENGINE_STATUS
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/RemoteGodotFragment.kt` — RemoteGodotFragment

## Animation blending and playback

Animation graph resources retain nodes, blend points, transitions, and node connections.

AnimationMixer, AnimationPlayer, AnimationTree, blend spaces, blend trees, and state machines compose animation playback.

Code examples: AnimationMixer, AnimationNodeBlendTree, AnimationNodeStateMachine

Evidence:
- Code: `scene/animation/animation_mixer.h` — AnimationMixer
- Code: `scene/animation/animation_blend_tree.h` — AnimationNodeBlendTree
- Code: `scene/animation/animation_node_state_machine.h` — AnimationNodeStateMachine

## Build-time source generation

These helpers convert build inputs into generated C++ declarations and byte data.

The Python build helpers generate C++ tables for documentation data, exporter registration, compressed documentation, and compressed translations.

Code examples: EditorTranslationList

Evidence:
- Code: `editor/editor_builders.py` — doc_data_class_path_builder
- Code: `editor/editor_builders.py` — register_exporters_builder
- Code: `editor/editor_builders.py` — make_doc_header
- Code: `editor/editor_builders.py` — make_translations

## Color-font paint processing

The paint layer is independent of a particular raster or vector output format.

HarfBuzz supplies paint functions plus bounded and extents contexts to evaluate color-font paint operations.

Code examples: hb_paint_funcs_t, hb_paint_extents_context_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-paint.hh` — hb_paint_funcs_t
- Code: `thirdparty/harfbuzz/src/hb-paint-bounded.hh` — hb_paint_bounded_context_t

## Cryptographic keys, hashing, and TLS support

CryptoKey is a Resource, so key material participates in the engine resource lifecycle.

Crypto, CryptoKey, HashingContext, and HMACContext expose key generation, signatures, encryption, hashing, and message authentication APIs.

Evidence:
- Code: `doc/classes/Crypto.xml` — Crypto
- Code: `doc/classes/CryptoKey.xml` — CryptoKey
- Code: `doc/classes/HMACContext.xml` — HMACContext

## C# translation extraction

The parser builds parse options from project define constants and resolves invocation symbols through a semantic model.

The editor parses C# source into a Roslyn compilation and extracts constant translation-call arguments and nearby comments.

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools/CsTranslationParserPlugin.cs` — CsTranslationParserPlugin
- Code: `modules/mono/editor/GodotTools/GodotTools/CsTranslationParserPlugin.cs` — GetProjectDefineConstants

## Debug Adapter Protocol integration

The protocol layer translates debugger and editor state into a language-agnostic wire format.

The editor's debug adapter serializes protocol content as JSON, manages protocol data types, and starts a local protocol server.

Code examples: DAP::Source, DAP::Breakpoint, DAP::Capabilities

Evidence:
- Code: `editor/debugger/debug_adapter/debug_adapter_protocol.cpp` — Variant(p_params).to_json_string().to_utf8_buffer
- Code: `editor/debugger/debug_adapter/debug_adapter_types.h` — Source, Breakpoint, Capabilities, Message, Scope, StackFrame, Thread, and Variable
- External (primary, unverified (source anchor not found)): [Debug Adapter Protocol Overview](https://microsoft.github.io/debug-adapter-protocol/overview.html), accessed 2026-07-15

## Debugger capture and transport

Remote debugging is modeled as a service with message and capture state rather than as a platform adapter.

The debugger subsystem has engine captures and profilers, local and remote debugger implementations, and a socket-backed remote debugger peer.

Evidence:
- Code: `core/debugger/engine_debugger.h` — EngineDebugger
- Code: `core/debugger/remote_debugger.h` — RemoteDebugger
- Code: `core/debugger/remote_debugger_peer.h` — RemoteDebuggerPeerTCP

## Editor build integration

The implementation includes project-file editing, a build logger that writes an issues CSV, build management, and problem views.

Editor tools generate .NET projects, invoke builds, and surface CSV diagnostics.

Code examples: BuildDiagnostic

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools.ProjectEditor/ProjectGenerator.cs` — ProjectGenerator.GenGameProject
- Code: `modules/mono/editor/GodotTools/GodotTools.BuildLogger/GodotBuildLogger.cs` — GodotBuildLogger.Initialize
- Code: `modules/mono/editor/GodotTools/GodotTools/Build/BuildManager.cs` — BuildManager
- Code: `modules/mono/editor/GodotTools/GodotTools/Build/BuildProblemsView.cs` — BuildProblemsView.ReadDiagnosticsFromFile

## Editor theming

Theme construction is centralized in EditorThemeManager while visual variants reside in separate builders.

Editor theming builds a Theme from ThemeConfiguration and settings, with distinct classic and modern builders plus font, icon, color-map, and scale support.

Evidence:
- Code: `editor/themes/editor_theme_manager.h` — EditorThemeManager::ThemeConfiguration
- Code: `editor/themes/editor_theme_manager.cpp` — text editor color-theme settings and custom theme path handling
- Code: `editor/themes/theme_classic.h` — ThemeClassic
- Code: `editor/themes/theme_modern.h` — ThemeModern
- External (official, unverified (source anchor not found)): [Theme — Godot Engine documentation](https://docs.godotengine.org/en/stable/classes/class_theme.html), accessed 2026-07-15

## FBX scene import through GLTF structures

The implementation maps FBX scene, node, mesh, material, texture, animation, skin, and light data into GLTF-oriented state.

FBXDocument and FBXState specialize GLTF document and state structures while importers expose UFBX and FBX2GLTF editor entry points.

Code examples: FBXDocument, FBXState

Evidence:
- Code: `modules/fbx/fbx_document.cpp` — FBXDocument FBX-to-GLTF conversion paths
- Code: `modules/fbx/editor/editor_scene_importer_ufbx.h` — EditorSceneFormatImporterUFBX

## Files, directories, and configuration

This is direct application data persistence rather than the Resource import and loading path.

FileAccess and DirAccess read or write project and user paths, while ConfigFile stores section and key values.

Code examples: ConfigFile

Evidence:
- Code: `doc/classes/FileAccess.xml` — FileAccess
- Code: `doc/classes/DirAccess.xml` — DirAccess
- Code: `doc/classes/ConfigFile.xml` — ConfigFile

## Font subsetting

The partition contains input, planning, table-specific processing, serialization, and CFF-specific subset code.

HarfBuzz represents caller selections as a subset input and derives an in-memory plan for rewriting a font.

Code examples: hb_subset_input_t, hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-input.hh` — hb_subset_input_t
- Code: `thirdparty/harfbuzz/src/hb-subset-plan.hh` — hb_subset_plan_t

## Binary font-table access

The code directly models table records, offsets, and character-to-glyph lookup structures.

Graphite and HarfBuzz read binary font tables, including cmap mappings, to obtain glyph data and layout behavior.

Code examples: Face, hb_blob_t, hb_face_t

Evidence:
- Code: `thirdparty/graphite/src/inc/Face.h` — Face::Table
- Code: `thirdparty/graphite/src/CmapCache.cpp` — bmp_subtable, smp_subtable, and Cmap
- Code: `thirdparty/harfbuzz/src/hb-open-file.hh` — OpenTypeFontFile and OpenTypeFontFace
- Code: `thirdparty/harfbuzz/src/hb-ot-cmap-table.hh` — OT::cmap and cmap::accelerator_t
- External (official, unverified (source anchor not found)): [cmap — Character To Glyph Index Mapping Table](https://learn.microsoft.com/en-us/typography/opentype/spec/cmap), accessed 2026-07-15

## Gamepad motion fusion and calibration

The implementation uses vector and quaternion operations, smoothing, stillness checks, and calibration confidence settings.

GamepadMotion maintains orientation, gravity, processed acceleration, and automatic gyro-calibration state from gyro and accelerometer samples.

Evidence:
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — GamepadMotion
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — AutoCalibration
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — Motion

## Geometry resources

Geometry is the scene content abstraction that builders and intersectors address by geometry ID.

A Geometry represents typed primitive data, supplies bounds, and provides the common base used by mesh, curve, point, line, grid, user, and instance implementations.

Code examples: RTCGeometry, Geometry

Evidence:
- Code: `thirdparty/embree/kernels/common/geometry.h` — Geometry
- Code: `thirdparty/embree/kernels/common/scene_triangle_mesh.h` — TriangleMesh
- Code: `thirdparty/embree/kernels/common/scene_instance.h` — Instance

## Graphics procedure loading

The source is a vendored generated loader rather than a rendering implementation.

Generated glad loaders resolve OpenGL, EGL, and GLX procedure entry points and track available versions and extensions.

Evidence:
- Code: `thirdparty/glad/gl.c` — _glad_gl_userptr
- Code: `thirdparty/glad/egl.c` — _glad_egl_userptr
- External (official, verified): [Khronos Combined OpenGL Registry](https://registry.khronos.org/OpenGL/), accessed 2026-07-16

## IDE messaging protocol

Peer, decoder, client, server, and CLI forwarding code share Message and MessageContent as the protocol envelope.

IDE clients and the editor exchange framed requests and responses after a handshake, with JSON bodies for typed requests.

Code examples: Message, GodotIdeMetadata

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Message.cs` — Message
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/MessageDecoder.cs` — MessageDecoder
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Peer.cs` — Peer
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Requests/Requests.cs` — Request

## Keyboard and MIDI parsing

Keyboard utilities map key codes and names, while MIDIDriver::Parser interprets status and data bytes into MIDI messages.

Evidence:
- Code: `core/os/keyboard.h` — find_keycode_name
- Code: `core/os/midi_driver.h` — MIDIDriver::Parser
- Code: `core/os/midi_driver.cpp` — MIDIDriver::Parser::parse

## Picture planes and pixel ownership

The selected representation determines which pixel fields the encoder treats as native input.

WebPPicture represents one image as either ARGB pixels or YUV(A) planes with dimensions, strides, optional alpha, and private allocation pointers.

Code examples: WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — struct WebPPicture
- Code: `thirdparty/libwebp/src/enc/iterator_enc.c` — VP8IteratorImport
- External (official, verified): [WebP API Documentation](https://developers.google.com/speed/webp/docs/api), accessed 2026-07-16

## interactive music transition tables

AudioStreamInteractive combines clips with a TransitionKey-indexed transition map so playback can select a rule for a clip-to-clip change.

Code examples: AudioStreamInteractive, AudioStreamInteractive Transition

Evidence:
- Code: `modules/interactive_music/audio_stream_interactive.h` — AudioStreamInteractive::Clip, Transition, TransitionKey, and TransitionKeyHasher
- Code: `modules/interactive_music/audio_stream_interactive.cpp` — transition_map[TransitionKey(keys[i].x, keys[i].y)]

## iOS embedded adapter

The iOS build lists device metrics, display, view, main, and OS Objective-C++ sources in one library.

iOS platform classes adapt the engine’s OS, display, and view responsibilities for an embedded Apple target.

Evidence:
- Code: `platform/ios/SCsub` — ios_lib
- Code: `platform/ios/ios.h` — iOS
- Code: `platform/ios/display_server_ios.h` — DisplayServerIOS
- Code: `platform/ios/os_ios.h` — OS_IOS

## Linux/BSD desktop integration

The platform layer contains both generic Linux/BSD OS code and desktop-protocol integrations.

Linux/BSD platform code integrates OS services, portals, screensaver handling, TTS, X11, and Wayland display servers.

Evidence:
- Code: `platform/linuxbsd/os_linuxbsd.h` — OS_LinuxBSD
- Code: `platform/linuxbsd/freedesktop_portal_desktop.h` — FreeDesktopPortalDesktop
- Code: `platform/linuxbsd/freedesktop_screensaver.h` — FreeDesktopScreenSaver
- Code: `platform/linuxbsd/tts_linux.h` — TTS_Linux

## Locale resolution

The implementation exposes both object-oriented LocaleMatcher APIs and lower-level locale parsing code.

ICU parses locale identifiers and matches desired locales against supported locales using LSR values, likely-subtag data, and locale-distance data.

Code examples: Locale, LSR, LocaleMatcher

Evidence:
- Code: `thirdparty/icu4c/common/unicode/localematcher.h` — LocaleMatcher
- Code: `thirdparty/icu4c/common/lsr.h` — LSR

## Engine allocation helpers

Memory helpers define default allocators, typed allocators, allocation result wrappers, and global nil support.

Evidence:
- Code: `core/os/memory.h` — DefaultAllocator
- Code: `core/os/memory.h` — DefaultTypedAllocator
- Code: `core/os/memory.h` — memnew_result

## mobile XR interface

MobileVRInterface is the module's concrete XRInterface implementation.

Evidence:
- Code: `modules/mobile_vr/mobile_vr_interface.h` — MobileVRInterface : public XRInterface
- Code: `modules/mobile_vr/register_types.cpp` — mobile_vr module registration

## Native desktop services

These services are separate native integrations rather than methods embedded solely in window management.

macOS and Windows platform code includes native-menu and text-to-speech adapters alongside their display servers.

Evidence:
- Code: `platform/macos/native_menu_macos.h` — NativeMenuMacOS
- Code: `platform/macos/tts_macos.h` — TTSUtterance
- Code: `platform/windows/native_menu_windows.h` — NativeMenuWindows
- Code: `platform/windows/tts_windows.h` — TTS_Windows

## OpenXR structure chains

OpenXR input and output structures carry a typed structure identifier and a `next` pointer for extensible structure chains.

Code examples: XrInstanceCreateInfo, XrBaseInStructure

Evidence:
- Code: `thirdparty/openxr/include/openxr/openxr.h` — XrInstanceCreateInfo
- Code: `thirdparty/openxr/src/loader/loader_core.cpp` — reinterpret_cast<const XrBaseInStructure *>(info->next)
- External (official, unverified (source anchor not found)): [The OpenXR™ 1.1.60 Specification (with all registered extensions)](https://registry.khronos.org/OpenXR/specs/1.1/html/xrspec.html), accessed 2026-07-15

## Physics spaces, bodies, shapes, and joints

The implementation provides parallel 2D and 3D API families.

Each physics server models a self-contained space containing bodies, areas, joints, and shapes, and exposes APIs to create and configure them.

Code examples: PhysicsMaterial, RID

Evidence:
- Code: `doc/classes/PhysicsServer2D.xml` — PhysicsServer2D
- Code: `doc/classes/PhysicsServer3D.xml` — PhysicsServer3D

## Polygon clipping and result trees

Integer and scaled floating-point front ends share the ClipperBase execution machinery.

Clipper accepts subject and clip paths, runs a specified clip and fill rule, and builds closed or open paths or a hierarchy.

Code examples: PolyPathD

Evidence:
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h` — Clipper64::Execute
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h` — ClipperD
- Code: `thirdparty/clipper2/include/clipper2/clipper.offset.h` — ClipperOffset

## Portable mathematical fallbacks

The bundled code includes trigonometric, logarithmic, exponential, power, and floating-point classification operations.

SDL dispatches mathematical functions to platform library functions when present or to bundled fdlibm-derived implementations otherwise.

Evidence:
- Code: `thirdparty/sdl/stdlib/SDL_stdlib.c` — SDL_atan
- Code: `thirdparty/sdl/libm/math_libm.h` — SDL_uclibc_atan
- Code: `thirdparty/sdl/libm/e_sqrt.c` — __ieee754_sqrt

## Procedural noise textures

The module supplies an abstract Noise resource, FastNoiseLite implementation, texture resources, editor preview support, and tests.

Noise generators produce values and image data that NoiseTexture2D and NoiseTexture3D turn into textures.

Evidence:
- Code: `modules/noise/noise.h` — Noise
- Code: `modules/noise/fastnoise_lite.h` — FastNoiseLite
- Code: `modules/noise/noise_texture_2d.h` — NoiseTexture2D
- Code: `modules/noise/noise_texture_3d.h` — NoiseTexture3D
- External (official, unverified (source anchor not found)): [Noise — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_noise.html), accessed 2026-07-15

## Ray query state

The API provides scalar rays plus 4-, 8-, 16-, and compile-time-N packet forms.

A ray query carries origin, direction, near and far distances, time, mask, ID, and flags for intersection or occlusion traversal.

Code examples: RTCRay, RTCRayHit

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCRay
- Code: `thirdparty/embree/kernels/common/ray.h` — RayK
- Code: `thirdparty/embree/kernels/common/ray.h` — RayStreamAOS

## Regular-expression compilation and matching

PCRE2 compiles patterns into code objects and matches subjects with traditional and DFA matching engines.

Code examples: pcre2_code, pcre2_match_data

Evidence:
- Code: `thirdparty/pcre2/src/pcre2_compile.c` — pcre2_compile
- Code: `thirdparty/pcre2/src/pcre2_dfa_match.c` — pcre2_dfa_match
- External (official, unverified (source anchor not found)): [pcre2api man page](https://www.pcre.org/current/doc/html/pcre2api.html), accessed 2026-07-15

## Resources and asset lifecycle

Resources are shared data objects used by scenes, scripts, meshes, shapes, materials, and extensions.

A Resource is a RefCounted asset value with a path, name, scene-local option, and loader or saver lifecycle.

Code examples: Resource, Mesh, Material, GDExtension

Evidence:
- Code: `core/io/resource.h` — Resource
- Code: `doc/classes/Resource.xml` — Resource.resource_path

## Shader-language front end

This implementation provides the internal representation consumed by the SPIR-V lowering code.

The vendored glslang front end models shader types, symbols, source locations, parse context, and a typed intermediate tree.

Code examples: spv::Module, spv::Function, spv::Block

Evidence:
- Code: `thirdparty/glslang/glslang/Include/intermediate.h` — TIntermNode
- Code: `thirdparty/glslang/glslang/MachineIndependent/ParseHelper.h` — TParseContext
- External (official, verified): [OpenGL / OpenGL ES Reference Compiler](https://www.khronos.org/opengles/sdk/Reference-Compiler/), accessed 2026-07-16

## Shader source compilation

The public API separates an individual shader object from a program that links shader objects.

TShader and TProgram compile shader source after preprocessed tokens into stage intermediates and program-level diagnostics.

Prerequisites: shader-preprocessing

Code examples: TShader, TProgram

Evidence:
- Code: `thirdparty/glslang/glslang/Public/ShaderLang.h` — glslang::TShader
- Code: `thirdparty/glslang/glslang/Public/ShaderLang.h` — glslang::TProgram
- Code: `thirdparty/glslang/glslang/MachineIndependent/localintermediate.h` — TIntermediate
- External (primary, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.html), accessed 2026-07-15
- External (primary, verified): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-16

## Signal awaitability

SignalAwaiter bridges signal completion back to managed await continuations.

A Signal combines an owner and a signal name and exposes a custom awaiter for asynchronous notification.

Code examples: Variant, ManagedCallbacks

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Signal.cs` — Signal.GetAwaiter
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/SignalAwaiter.cs` — SignalAwaiter
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs` — ManagedCallbacks.SignalAwaiter_SignalCallback

## SPIR-V rewriting and optimization

re-spirv parses a word-addressed shader into instructions, functions, blocks, results, decorations, and specializations before optimization.

Code examples: Shader, Instruction, Optimizer

Evidence:
- Code: `thirdparty/re-spirv/re-spirv.h` — struct Shader
- Code: `thirdparty/re-spirv/re-spirv.cpp` — OptimizerContext

## Strings, interned names, and node paths

The string layer implements Unicode strings, hash-backed interned StringName values, path and subpath storage in NodePath, and fuzzy-search helpers.

Code examples: StringName, NodePath

Evidence:
- Code: `core/string/ustring.h` — String
- Code: `core/string/string_name.h` — StringName
- Code: `core/string/node_path.h` — NodePath
- Code: `core/string/fuzzy_search.h` — FuzzySearch

## Task scheduling and synchronization

The scheduler can be selected between internal, TBB, and PPL implementations through configuration macros.

The internal scheduler represents work as tasks, queues, threads, and thread pools, and provides barriers and atomic helpers for synchronization.

Code examples: TaskScheduler, Task, TaskQueue, ThreadPool, BarrierActive

Evidence:
- Code: `thirdparty/embree/common/tasking/taskscheduler.h` — TASKING_INTERNAL
- Code: `thirdparty/embree/common/tasking/taskschedulerinternal.h` — TaskQueue
- Code: `thirdparty/embree/common/sys/barrier.h` — BarrierActive
- Code: `thirdparty/embree/common/sys/atomic.h` — _atomic_min

## Text backends and translation domains

TextLine and TextParagraph are TextServer abstractions for a single line or paragraph.

TextServerManager registers and selects TextServer implementations, while TranslationServer stores language data in named TranslationDomain collections.

Code examples: TranslationDomain, TextLine

Evidence:
- Code: `doc/classes/TextServerManager.xml` — TextServerManager
- Code: `doc/classes/TranslationServer.xml` — TranslationServer
- Code: `doc/classes/TranslationDomain.xml` — TranslationDomain

## Texture block codecs

etcpak exposes ETC/EAC and BC block compression plus block-decoding entry points that operate on typed source and destination buffers.

Evidence:
- Code: `thirdparty/etcpak/ProcessRGB.hpp` — CompressEtc1Rgb
- Code: `thirdparty/etcpak/ProcessDxtc.hpp` — CompressBc5
- Code: `thirdparty/etcpak/DecodeRGB.hpp` — DecodeRGBBlock

## Texture compression pipeline

The supplied implementation includes ETC-family, UASTC, ASTC, GPU texture, and container-output code.

Vendored Basis Universal code separates frontend block/codebook work, backend encoding, and Basis or KTX2 output creation.

Evidence:
- Code: `thirdparty/basis_universal/encoder/basisu_frontend.h` — basisu_frontend
- Code: `thirdparty/basis_universal/encoder/basisu_backend.h` — basisu_backend and basisu_backend_output
- Code: `thirdparty/basis_universal/encoder/basisu_comp.h` — basis_compressor

## Texture format descriptions

The descriptor is an exchanged binary description of texel layout rather than an opaque GPU-native layout.

Khronos Data Format descriptors represent image-format layout using header-word and sample-word accessors, while helper code creates, interprets, queries, and prints those descriptors.

Code examples: KTX Level Index Entry

Evidence:
- Code: `thirdparty/libktx/include/KHR/khr_df.h` — khr_word_e
- Code: `thirdparty/libktx/external/dfdutils/createdfd.c` — writeHeader
- Code: `thirdparty/libktx/external/dfdutils/interpretdfd.c` — interpretDFD
- External (official, verified): [Khronos Data Format Specification](https://registry.khronos.org/DataFormat/specs/1.4/dataformat.1.4.html), accessed 2026-07-16

## VI native surface creation

This is a platform-specific Vulkan surface extension declaration.

The `VK_NN_vi_surface` header declares input for creating a `VkSurfaceKHR` from a Nintendo VI window handle.

Code examples: VkViSurfaceCreateInfoNN, VkSurfaceKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VkViSurfaceCreateInfoNN
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — vkCreateViSurfaceNN
- External (official, unverified (source anchor not found)): [vkCreateViSurfaceNN(3) :: Vulkan Documentation Project](https://docs.vulkan.org/refpages/latest/refpages/source/vkCreateViSurfaceNN.html), accessed 2026-07-15

## Vulkan video session configuration

This partition declares video API data contracts; it does not show a video-processing implementation.

The generated records model a video session through a video profile and standard-header-version information, with separate records for decode and encode operations.

Code examples: VideoSessionCreateInfoKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct VideoSessionCreateInfoKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct VideoDecodeInfoKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct VideoEncodeInfoKHR

## WebRTC peer connections

The core class has extension and JavaScript-backed implementations.

A WebRTC peer connection exposes connection state, negotiation, ICE candidates, and data-channel creation to Godot.

Code examples: WebRTCPeerConnection

Evidence:
- Code: `modules/webrtc/webrtc_peer_connection.h` — WebRTCPeerConnection
- Code: `modules/webrtc/webrtc_peer_connection_extension.h` — WebRTCPeerConnectionExtension
- External (primary, verified): [WebRTC: Real-Time Communication in Browsers](https://www.w3.org/TR/webrtc/), accessed 2026-07-16

## XR tracking and poses

Trackers classify sources and poses carry the spatial data exposed for those sources.

XRServer coordinates interfaces and tracker types, while positional, body, controller, face, and hand trackers publish XR pose state.

Code examples: XRTracker, XRPose

Evidence:
- Code: `servers/xr/xr_server.h` — XRServer::RenderState
- Code: `servers/xr/xr_tracker.h` — XRTracker
- Code: `servers/xr/xr_pose.h` — XRPose
- Code: `servers/xr/xr_positional_tracker.h` — XRPositionalTracker
- Code: `servers/xr/xr_body_tracker.h` — XRBodyTracker
- Code: `servers/xr/xr_hand_tracker.h` — XRHandTracker

## Alpha-plane coding

Alpha encoding has its own filter choice and can use lossless compression.

The alpha plane is filtered, optionally quantized, and encoded before its bytes are incorporated with the lossy image representation.

Prerequisites: input-picture-planes

Code examples: WebPPicture, WebPConfig

Evidence:
- Code: `thirdparty/libwebp/src/enc/alpha_enc.c` — VP8EncEncodeAlpha
- Code: `thirdparty/libwebp/src/enc/near_lossless_enc.c` — VP8LNearLossless
- Code: `thirdparty/libwebp/src/webp/encode.h` — WebPConfig.alpha_compression, WebPPicture.a

## Thread synchronization

The class documentation specifies lifetime constraints to avoid cleanup races and deadlocks.

Mutex synchronizes access to a critical section between Thread instances and is documented as a reentrant binary semaphore.

Evidence:
- Code: `doc/classes/Mutex.xml` — Mutex

## Device runtime

It is the API-level runtime owner behind scene and buffer creation.

A Device configures Embree execution, exposes device and thread error messages, and limits tasking resources.

Code examples: RTCDevice, Device

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h` — RTCDevice
- Code: `thirdparty/embree/kernels/common/device.h` — Device
- Code: `thirdparty/embree/kernels/common/device.cpp` — Device::getThreadLastErrorMessage

## Geometry families

The checked configuration directly enables triangle geometry, while the indexed source also contains the other family implementations.

The source defines Geometry subclasses for triangle meshes, quad meshes, curves, line segments, points, grids, subdivision meshes, user geometry, and instances.

Prerequisites: geometry-resources

Code examples: TriangleMesh, QuadMesh, CurveGeometry, LineSegments, Points, GridMesh, SubdivMesh, UserGeometry

Evidence:
- Code: `thirdparty/embree/kernels/common/scene_triangle_mesh.h` — TriangleMesh
- Code: `thirdparty/embree/kernels/common/scene_curves.h` — CurveGeometry
- Code: `thirdparty/embree/kernels/common/scene_grid_mesh.h` — GridMesh
- Code: `thirdparty/embree/kernels/config.h` — EMBREE_GEOMETRY_TRIANGLE

## Geometry and transform values

These are the coordinate and bounds values shared by geometric algorithms and dynamic argument conversion.

The math layer defines 2D and 3D vectors, rectangles, boxes, planes, bases, quaternions, projections, and transforms as reusable value types.

Code examples: AABB, Transform3D

Evidence:
- Code: `core/math/aabb.h` — AABB
- Code: `core/math/transform_3d.h` — Transform3D
- Code: `core/math/projection.h` — Projection

## Image codec integration

The supplied modules cover BMP loading, DDS resource loading and saving, and several block-compression implementations.

Image codec modules provide loader, saver, compressor, and decompressor implementations against engine image and resource interfaces.

Code examples: Image, DDSFormatInfo

Evidence:
- Code: `modules/bmp/image_loader_bmp.h` — ImageLoaderBMP
- Code: `modules/dds/image_saver_dds.h` — ImageSaverDDS
- Code: `modules/astcenc/image_compress_astcenc.h` — register_image_compress_astcenc

## Translation and locale selection

Translation catalogs map StringName keys to messages, TranslationDomain selects applicable catalogs, and TranslationServer manages locale, fallback, and plural-rule services.

Prerequisites: string-names-paths

Code examples: Translation, TranslationDomain, StringName

Evidence:
- Code: `core/string/translation.h` — Translation::MessageKey
- Code: `core/string/translation_domain.h` — TranslationDomain
- Code: `core/string/translation_server.h` — TranslationServer
- Code: `core/string/plural_rules.h` — PluralRules

## SCons module build configuration

Python SCons scripts determine whether modules can build and add C++ source files, generated shader headers, or third-party sources to the build environment.

Evidence:
- Code: `modules/godot_physics_2d/config.py` — can_build
- Code: `modules/lightmapper_rd/SCsub` — GLSL_HEADER and add_source_files
- Code: `modules/jolt_physics/SCsub` — thirdparty_sources and add_source_files

## OpenType table routing

Separate dispatchers cover color, layout, ordinary glyph-related, and variation tables.

Font subsetting routes recognized OpenType tags to typed table subsetters; color routing explicitly skips CBDT because CBLC handles it.

Prerequisites: font-subsetting

Code examples: hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-table-color.cc` — _hb_subset_table_color
- Code: `thirdparty/harfbuzz/src/hb-subset-table-layout.cc` — _hb_subset_table_layout

## Plugin and assembly reload

The loader separates shared assemblies from plugin-specific dependency resolution and records a weak reference for load-context tracking.

Editor plug-ins load into AssemblyLoadContext instances that resolve dependencies and may be collectible.

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/PluginLoadContext.cs` — PluginLoadContext
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/Main.cs` — LoadPlugin
- Code: `modules/mono/editor/GodotTools/GodotTools/HotReloadAssemblyWatcher.cs` — HotReloadAssemblyWatcher

## Raster font rendering

The raster implementation includes PNG-related image code as well as scan-conversion-oriented drawing code.

Color paint records are rasterized through image, draw, paint, clipping, edge, and glyph-extents structures.

Prerequisites: color-font-paint

Code examples: hb_raster_image_t, hb_raster_paint_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-raster-image.hh` — hb_raster_image_t
- Code: `thirdparty/harfbuzz/src/hb-raster-paint.hh` — hb_raster_paint_t

## Resource-bundle data

The code has both C-facing UResourceBundle state and C++ ResourceBundle wrappers.

ICU resource bundles expose strings, binary values, integer vectors, keys, names, and locales over loaded resource data.

Code examples: UResourceBundle, ResourceDataValue

Evidence:
- Code: `thirdparty/icu4c/common/uresimp.h` — UResourceBundle
- Code: `thirdparty/icu4c/common/resbund.cpp` — ResourceBundle::getBinary

## Resource loading and caching

Loading is a registry-based service rather than a single file-format implementation.

ResourceLoader selects registered ResourceFormatLoader instances, applies cache modes, reports dependencies, and tracks threaded load tasks.

Code examples: Resource, ResourceLoader::ThreadLoadTask

Evidence:
- Code: `core/io/resource_loader.h` — ResourceLoader
- Code: `core/io/resource_loader.h` — ResourceLoader::ThreadLoadTask
- Code: `core/io/resource.h` — ResourceCache

## Spatial predictive filters

Scalar, SSE2, NEON, MSA, and MIPS implementations cover the same filtering roles.

Spatial filters produce residuals from neighboring pixel-plane values using horizontal, vertical, gradient, and other predictor forms.

Prerequisites: input-picture-planes

Evidence:
- Code: `thirdparty/libwebp/src/dsp/filters.c` — GradientPredictor_C, WebPFiltersInit
- Code: `thirdparty/libwebp/src/utils/filters_utils.c` — WebPEstimateBestFilter
- Code: `thirdparty/libwebp/src/dsp/filters_sse2.c` — GradientPredictor_SSE2

## Vulkan extensible records

The convention appears across the generated binding's externally exchanged configuration records.

Many Vulkan input records carry a `pNext` pointer, while the C VI surface record also carries an `sType` discriminator, forming an extensible record convention.

Code examples: VkViSurfaceCreateInfoNN, GraphicsPipelineCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VkViSurfaceCreateInfoNN
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct GraphicsPipelineCreateInfo

## Editor build composition

Editor source aggregation is expressed as Python-based SCons configuration alongside the C++ implementation.

SCsub build scripts add C++ source files to the editor source set and invoke child SConscript files for nested editor subsystems.

Evidence:
- Code: `editor/import/SCsub` — env.add_source_files
- Code: `editor/scene/2d/SCsub` — SConscript
- Code: `editor/icons/editor_icons_builders.py` — make_editor_icons_action

## Encoder configuration

Initialization and validation are explicit ABI-aware API steps.

WebPConfig selects lossy or lossless encoding and controls quality, effort, segmentation, filtering, alpha, thread, and memory choices for the configured picture.

Prerequisites: input-picture-planes

Code examples: WebPConfig, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — struct WebPConfig, WebPConfigInit, WebPValidateConfig
- Code: `thirdparty/libwebp/src/enc/config_enc.c` — WebPConfigInitInternal
- External (official, verified): [WebP API Documentation](https://developers.google.com/speed/webp/docs/api), accessed 2026-07-16

## HDR, JPEG, and KTX loading

Separate module adapters implement an HDR ImageFormatLoader, a libjpeg-turbo ImageFormatLoader, and a KTX ResourceFormatLoader.

Evidence:
- Code: `modules/hdr/image_loader_hdr.h` — ImageLoaderHDR : public ImageFormatLoader
- Code: `modules/jpg/image_loader_libjpeg_turbo.h` — ImageLoaderLibJPEGTurbo : public ImageFormatLoader
- Code: `modules/ktx/texture_loader_ktx.h` — ResourceFormatKTX : public ResourceFormatLoader

## Mesh geometry algorithms

The implementation uses geometry primitives to build Delaunay triangulations, convex hulls, quick hulls, and polygon triangulations.

Prerequisites: geometry-transforms

Code examples: TriangleMesh, AABB

Evidence:
- Code: `core/math/delaunay_2d.h` — Delaunay2D
- Code: `core/math/delaunay_3d.h` — Delaunay3D
- Code: `core/math/convex_hull.h` — ConvexHullComputer
- Code: `core/math/triangulate.h` — Triangulate

## Triangle provenance

The relation data is used during boolean-result construction and property transfer.

CSG output keeps triangle provenance through TriRef records and MeshRelationD triRef collections so result triangles can be related to source meshes.

Prerequisites: csg-boolean-operations

Code examples: TriRef, MeshRelationD

Evidence:
- Code: `thirdparty/manifold/src/shared.h` — struct TriRef
- Code: `thirdparty/manifold/src/impl.h` — struct Manifold::Impl::MeshRelationD
- Code: `thirdparty/manifold/src/boolean_result.cpp` — outR.meshRelation_.triRef; struct MapTriRef

## OpenXR runtime loading

The embedded loader reads runtime and API-layer manifests, validates `next` chains during instance creation, and creates a loader instance.

Prerequisites: openxr-structure-chains

Code examples: ManifestFile, LoaderInstance, XrInstanceCreateInfo

Evidence:
- Code: `thirdparty/openxr/src/loader/manifest_file.hpp` — RuntimeManifestFile
- Code: `thirdparty/openxr/src/loader/loader_instance.hpp` — LoaderInstance::CreateLoaderInstance

## Packed resource files

PackedData is the virtual-file index behind pack-backed file access.

Packed data records pack paths, offsets, sizes, hashes, encryption flags, bundle flags, and delta flags, then supplies files to the resource-loading service through PackSource.

Prerequisites: resource-loading

Code examples: PackedData::PackedFile

Evidence:
- Code: `core/io/file_access_pack.h` — PackedData
- Code: `core/io/file_access_pack.h` — PackedData::PackedFile
- Code: `core/io/file_access_pack.h` — PackSource

## Running projects and instances

Run UI, process embedding, device selection, and multiple-instance arguments are implemented as related editor services.

The run subsystem starts editor launches, exposes native-device runs, supports configurable multiple instances, and hosts embedded game-view processes.

Evidence:
- Code: `editor/run/editor_run.h` — EditorRun
- Code: `editor/run/editor_run_native.h` — EditorRunNative
- Code: `editor/run/run_instances_dialog.h` — RunInstancesDialog::InstanceData
- Code: `editor/run/embedded_process.h` — EmbeddedProcess

## Threads and synchronization

The OS layer defines Thread, mutexes, condition variables, read-write locks, semaphores, spin locks, and safe binary mutexes.

Evidence:
- Code: `core/os/thread.h` — Thread
- Code: `core/os/mutex.h` — MutexLock
- Code: `core/os/condition_variable.h` — ConditionVariable
- Code: `core/os/rw_lock.h` — RWLock

## Transform interpolation

TransformInterpolator and InterpolatedProperty calculate intermediate geometry transforms and values.

Prerequisites: geometry-transforms

Code examples: Transform3D

Evidence:
- Code: `core/math/transform_interpolator.h` — TransformInterpolator
- Code: `core/templates/interpolated_property.h` — InterpolatedProperty

## Wayland window backend

Wayland state is split between a display server, a protocol thread, and an embedder.

Wayland window handling builds on Linux/BSD platform integration with a dedicated thread, client protocol objects, and window state.

Prerequisites: linuxbsd-desktop-integration

Code examples: WaylandThread.WindowState

Evidence:
- Code: `platform/linuxbsd/wayland/display_server_wayland.h` — DisplayServerWayland
- Code: `platform/linuxbsd/wayland/wayland_thread.h` — WaylandThread
- Code: `platform/linuxbsd/wayland/wayland_thread.h` — WaylandThread::WindowState
- Code: `platform/linuxbsd/wayland/wayland_embedder.h` — WaylandEmbedder

## Codec SIMD specialization

Conditional C preprocessing exposes AVX2, SSE4.1, and NEON codec declarations; C preprocessing is needed to explain which declarations are available.

Evidence:
- Code: `thirdparty/etcpak/Dither.hpp` — DitherAvx2
- Code: `thirdparty/etcpak/Tables.hpp` — g_table_SIMD
- Code: `thirdparty/etcpak/Tables.hpp` — g_table128_NEON

## Editor and project settings

Editor preferences and project configuration are separate concerns in this partition.

EditorSettings persists editor-side setting containers and project metadata, while dedicated dialogs expose editor, project, autoload, action-map, and input-event configuration.

Code examples: EditorSettings

Evidence:
- Code: `editor/settings/editor_settings.h` — EditorSettings, Plugin, and VariantContainer
- Code: `editor/settings/editor_settings.cpp` — EditorSettings::_get_project_metadata_path and metadata load/save code
- Code: `editor/settings/project_settings_editor.h` — ProjectSettingsEditor : public AcceptDialog
- Code: `editor/settings/editor_autoload_settings.h` — EditorAutoloadSettings and AutoloadInfo

## FreeType module composition

The FreeType base and autofit modules combine included implementation units into compiled engine modules; C preprocessing is needed to explain the inclusion-based composition.

Evidence:
- Code: `thirdparty/freetype/src/base/ftbase.c` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/autofit/autofit.c` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/include/freetype/config/ftmodule.h` — FT_USE_MODULE

## Engine object model

This is the base model that the scene, resource, script, and editor APIs build upon.

Object-derived types provide the common runtime identity and behavior base, while RefCounted supplies reference-counted lifetime for resource-like values.

Code examples: Node, Resource, ClassInfo

Evidence:
- Code: `core/object/object.h` — Object
- Code: `core/object/ref_counted.h` — RefCounted

## OpenXR dispatch forwarding

Generated loader entry points forward API calls through a dispatch table belonging to the selected runtime.

Prerequisites: openxr-runtime-loading

Code examples: XrGeneratedDispatchTableCore

Evidence:
- Code: `thirdparty/openxr/src/xr_generated_dispatch_table_core.h` — XrGeneratedDispatchTableCore
- Code: `thirdparty/openxr/src/loader/xr_generated_loader.cpp` — xrWaitFrame

## OpenXR runtime integration

This is the integration boundary; later OpenXR topics specialize its action, extension, spatial, and rendering behavior.

The module places OpenXRInterface over OpenXRAPI so the engine can interact with an OpenXR runtime.

Code examples: OpenXRActionMap, OpenXRFutureResult

Evidence:
- Code: `modules/openxr/openxr_interface.h` — class OpenXRInterface : public XRInterface
- Code: `modules/openxr/openxr_api.h` — class OpenXRAPI
- External (official, unverified (source anchor not found)): [The OpenXR 1.1 Specification](https://registry.khronos.org/OpenXR/specs/1.1/html/xrspec.html), accessed 2026-07-15

## Event queue and watches

Queued entries can carry temporary memory associated with an event payload.

The event runtime stores SDL_Event values in a mutex-protected linked queue with an atomic count and event-watch support.

Code examples: SDL_Event, SDL_EventEntry

Evidence:
- Code: `thirdparty/sdl/events/SDL_events.c` — SDL_EventEntry
- Code: `thirdparty/sdl/events/SDL_events.c` — SDL_EventQ
- Code: `thirdparty/sdl/events/SDL_eventwatch_c.h` — SDL_EventWatchList
- External (official, verified): [SDL3 CategoryEvents](https://wiki.libsdl.org/SDL3/CategoryEvents), accessed 2026-07-16

## Collision shapes

The Shape hierarchy is the common geometry contract for bodies and collision queries.

Collision shapes define the geometric representation used for ray casts, shape casts, point tests, overlap tests, triangle extraction, and collision dispatch.

Code examples: Shape, ShapeSettings

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h` — ShapeSettings, Shape, ShapeFunctions
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionDispatch.h` — CollisionDispatch

## Dynamic values and dictionaries

It is the common value carrier used by reflection, scene serialization, scripting, and several networking APIs.

Variant is the cross-cutting tagged value representation for scalar, math, object, callable, signal, dictionary, array, and packed-array values.

Code examples: Variant, JSON-RPC document

Evidence:
- Code: `core/variant/variant.h` — Variant::Type
- Code: `doc/classes/Dictionary.xml` — Dictionary

## High-level RPC routing

RPC configuration is cached per node and packet processing distinguishes remote-call, spawn, despawn, sync, and system commands.

SceneMultiplayer routes RPC calls, peer state, connection commands, and RPC configuration.

Evidence:
- Code: `modules/multiplayer/scene_multiplayer.h` — SceneMultiplayer
- Code: `modules/multiplayer/scene_rpc_interface.h` — SceneRPCInterface
- Code: `modules/multiplayer/scene_cache_interface.h` — SceneCacheInterface
- External (official, verified): [High-level multiplayer — Godot Engine](https://docs.godotengine.org/en/latest/tutorials/networking/high_level_multiplayer.html), accessed 2026-07-16

## OpenXR action configuration

The action-map subtree uses Resource-derived configuration objects for the OpenXR input graph.

OpenXRActionMap persists action sets, actions, interaction profiles, bindings, modifiers, and haptic feedback configuration.

Code examples: OpenXRActionMap

Evidence:
- Code: `modules/openxr/action_map/openxr_action_map.h` — OpenXRActionMap
- Code: `modules/openxr/action_map/openxr_action.h` — OpenXRAction
- Code: `modules/openxr/action_map/openxr_binding_modifier.h` — OpenXRBindingModifier
- Code: `modules/openxr/action_map/openxr_haptic_feedback.h` — OpenXRHapticBase
- External (official, unverified (source anchor not found)): [OpenXRActionMap — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_openxractionmap.html), accessed 2026-07-15

## Binding modifiers

The implementation has a base modifier plus profile-level and action-binding subclasses, including D-pad and analog-threshold modifiers.

Binding modifiers alter interaction-profile or per-action binding data before it is submitted to OpenXR.

Prerequisites: openxr-action-system, openxr-interaction-profiles

Code examples: OpenXRIPBinding

Evidence:
- Code: `modules/openxr/extensions/openxr_dpad_binding_extension.h` — class OpenXRDpadBindingModifier : public OpenXRIPBindingModifier
- Code: `modules/openxr/extensions/openxr_valve_analog_threshold_extension.h` — class OpenXRAnalogThresholdModifier : public OpenXRActionBindingModifier
- Code: `modules/openxr/doc_classes/OpenXRBindingModifier.xml` — OpenXRBindingModifier::_get_ip_modification

## PSA key lifecycle and storage

The implementation distinguishes volatile identifier ranges and contains dedicated storage and slot-management code.

PSA key records retain configuration-selected algorithms, type, lifetime, identifier, and policy attributes from initialization or import through slot management and optional storage.

Prerequisites: crypto-feature-configuration

Code examples: psa_key_attributes_t

Evidence:
- Code: `thirdparty/mbedtls/tf-psa-crypto/include/psa/crypto_struct.h` — struct psa_key_attributes_s; PSA_KEY_ATTRIBUTES_INIT
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto.c` — psa_import_key
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_slot_management.h` — PSA_KEY_ID_VOLATILE_MIN; PSA_KEY_ID_VOLATILE_MAX
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_storage.c` — persistent key storage implementation
- External (official, verified): [PSA Certified Crypto API 1.1](https://arm-software.github.io/psa-api/crypto/1.1/), accessed 2026-07-16

## Runtime property groups and hints

The implementation uses properties in core services and stream metadata.

SDL runtime property groups associate named values with an ID, while hints retrieve configuration from property and environment sources.

Code examples: SDL_PropertiesID

Evidence:
- Code: `thirdparty/sdl/SDL_properties.c` — SDL_GetStringProperty
- Code: `thirdparty/sdl/SDL_hints.c` — SDL_GetHint
- Code: `thirdparty/sdl/include/SDL3/SDL_properties.h` — SDL_PropertiesID

## Render-pass configuration

The supplied C++ binding also declares the newer `RenderPassCreateInfo2` form.

`RenderPassCreateInfo` combines attachment descriptions, subpass descriptions, and subpass dependencies into one externally exchanged rendering configuration.

Code examples: RenderPassCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct RenderPassCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct RenderPassCreateInfo2

## Worker-based parallelism

Thread behavior is abstracted behind a library interface instead of direct thread calls in encoder modules.

The encoding implementation obtains a WebPWorkerInterface and conditionally splits analysis or lossless work when thread-level configuration permits it.

Prerequisites: encoder-configuration

Code examples: VP8Encoder

Evidence:
- Code: `thirdparty/libwebp/src/utils/thread_utils.h` — WebPWorkerInterface
- Code: `thirdparty/libwebp/src/utils/thread_utils.c` — WebPGetWorkerInterface
- Code: `thirdparty/libwebp/src/enc/analysis_enc.c` — WebPGetWorkerInterface, do_mt
- Code: `thirdparty/libwebp/src/enc/vp8l_enc.c` — WebPGetWorkerInterface

## Collision filtering

The filtering interfaces let a caller restrict which bodies or subshapes may interact.

Collision filtering applies object-layer, broad-phase-layer, group, body, and shape filters before candidate generation or collision queries.

Prerequisites: collision-shapes

Code examples: Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/ObjectLayer.h` — ObjectLayerFilter and ObjectLayerPairFilter
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/ShapeFilter.h` — ShapeFilter and ReversedShapeFilter
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionGroup.h` — CollisionGroup

## Convex support and penetration

ConvexShape exposes support mappings that collision algorithms can query.

Convex collision represents collision shapes with support functions and uses GJK closest-point and EPA penetration calculations.

Prerequisites: collision-shapes

Code examples: Shape

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/ConvexShape.h` — ConvexShape::Support and GetSupportFunction
- Code: `thirdparty/jolt_physics/Jolt/Geometry/GJKClosestPoint.h` — GJKClosestPoint
- Code: `thirdparty/jolt_physics/Jolt/Geometry/EPAPenetrationDepth.h` — EPAPenetrationDepth

## Geometry preprocessing

The geometry utilities are used to prepare collision-friendly representations from source triangles and vertices.

Geometry preprocessing converts triangle lists to indexed geometry, builds convex hulls, and partitions triangles for spatial acceleration structures used by collision shapes.

Prerequisites: collision-shapes

Code examples: Shape

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Geometry/Indexify.h` — Indexify and Deindexify
- Code: `thirdparty/jolt_physics/Jolt/Geometry/ConvexHullBuilder.h` — ConvexHullBuilder
- Code: `thirdparty/jolt_physics/Jolt/TriangleSplitter/TriangleSplitter.h` — TriangleSplitter

## Reflective property inspection

Specialized controls are selected by property shape while inspector plug-ins can supply custom editors.

The inspector reads Object properties into EditorProperty controls and supports custom parsing, revert, copy, paste, keying, pinning, and favorites.

Evidence:
- Code: `editor/inspector/editor_inspector.h` — EditorProperty
- Code: `editor/inspector/editor_inspector.h` — EditorInspectorPlugin
- Code: `editor/inspector/editor_properties.h` — EditorInspectorDefaultPlugin

## Runtime class metadata

This metadata is the bridge between native engine classes and the exposed API.

The runtime records classes, inheritance, methods, properties, signals, and instantiation capability for engine objects.

Code examples: RID, Resource

Evidence:
- Code: `core/object/class_db.h` — ClassDB

## Scene replication configuration

Each configured property stores a spawn flag and a replication mode, and the configuration maintains derived lists for replication phases.

A saved SceneReplicationConfig separates NodePath properties into spawn, sync, and watch lists.

Code examples: SceneReplicationConfig, ReplicationProperty

Evidence:
- Code: `modules/multiplayer/scene_replication_config.h` — SceneReplicationConfig
- Code: `modules/multiplayer/scene_replication_config.cpp` — SceneReplicationConfig::get_spawn_properties
- External (official, unverified (source anchor not found)): [SceneReplicationConfig — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_scenereplicationconfig.html), accessed 2026-07-15

## Abstract I/O streams

The stream implementation also provides fixed-width endian read and write helpers.

SDL I/O streams wrap byte operations for file and memory implementations and expose an SDL property group.

Prerequisites: sdl-runtime-properties

Code examples: SDL_IOStream, SDL_PropertiesID

Evidence:
- Code: `thirdparty/sdl/io/SDL_iostream.c` — SDL_IOStream
- Code: `thirdparty/sdl/io/SDL_iostream.c` — SDL_ReadU32LE
- Code: `thirdparty/sdl/include/SDL3/SDL_iostream.h` — SDL_IOStream

## Text shaping plans

The public shaping entry points accept features and an optional shaper list, while plan creation accepts matching configuration.

A shaping-plan key records segment properties, user features, coordinates, and shaper selection so shaping can use a plan object.

Code examples: hb_shape_plan_t, hb_shape_plan_key_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-shape-plan.hh` — hb_shape_plan_key_t
- Code: `thirdparty/harfbuzz/src/hb-shape.h` — hb_shape

## Class metadata and reflection

Registration also records construction behavior and API exposure for native and extension classes.

ClassDB records Object-derived classes, properties, methods, signals, and constants for runtime lookup.

Prerequisites: object-model

Code examples: ClassInfo, Variant

Evidence:
- Code: `core/object/class_db.h` — ClassDB::ClassInfo
- Code: `doc/classes/ClassDB.xml` — ClassDB

## Replicated spawning

Spawner state records tracked nodes and their spawn data, while the replication interface sends spawn and despawn commands.

The spawner tracks spawnable scenes and sends configured spawn property lists while creating or removing nodes.

Prerequisites: scene-replication-configuration

Code examples: SceneReplicationConfig

Evidence:
- Code: `modules/multiplayer/multiplayer_spawner.h` — MultiplayerSpawner
- Code: `modules/multiplayer/multiplayer_spawner.cpp` — MultiplayerSpawner::get_spawn_argument
- Code: `modules/multiplayer/scene_replication_interface.cpp` — SceneReplicationInterface
- External (official, unverified (source anchor not found)): [MultiplayerSpawner — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_multiplayerspawner.html), accessed 2026-07-15

## Godot member exposure

The source generators inspect compatible C# members and create the metadata queried by ScriptManagerBridge.

Generator output registers exported members, signals, and RPC metadata as engine-visible method and property descriptions.

Code examples: Variant

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPropertiesGenerator.cs` — ScriptPropertiesGenerator
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptMethodsGenerator.cs` — ScriptMethodsGenerator
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptSignalsGenerator.cs` — ScriptSignalsGenerator
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ScriptManagerBridge.cs` — ScriptManagerBridge

## Manifold mesh interchange

This is the external mesh-shaped data carrier at the geometry boundary.

MeshGLP carries indexed triangles, per-vertex properties, merge information, and run metadata into Manifold construction.

Code examples: MeshGLP

Evidence:
- Code: `thirdparty/manifold/include/manifold/manifold.h` — struct MeshGLP
- Code: `thirdparty/manifold/src/constructors.cpp` — Manifold constructors accepting MeshGLP

## Script resources and instances

The native module represents script validity, reload state, members, and compiler/analyzer relationships.

GDScript is a Script Resource whose compiled members and functions supply script instances to Object-derived engine objects.

Prerequisites: reflection, resources

Code examples: GDScript, Node

Evidence:
- Code: `modules/gdscript/gdscript.h` — GDScript
- Code: `modules/gdscript/gdscript.h` — GDScript::MemberInfo

## Project filesystem index

The index separates directory topology from per-file metadata and supports rescanning.

The filesystem service scans project directories into a tree of file records with resource type, UID, import state, dependency list, and script-class metadata.

Code examples: EditorFileSystemDirectory, EditorFileSystemDirectory::FileInfo

Evidence:
- Code: `editor/file_system/editor_file_system.h` — EditorFileSystemDirectory
- Code: `editor/file_system/editor_file_system.h` — EditorFileSystemDirectory::FileInfo
- Code: `editor/file_system/editor_file_system.cpp` — EditorFileSystem::_scan_filesystem

## Fixture contracts

A source fixture provides the case; its paired stored artifact provides the expected observable result.

The suite stores GDScript source alongside expected outcome files and test assets, so behavior is expressed as a checked fixture contract.

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd` — test()
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.out` — expected output

## Managed script registration metadata

The generator emits ScriptPathAttribute declarations and an AssemblyHasScriptsAttribute containing discovered script types.

C# script classes receive generated script-path and assembly script-type metadata for engine discovery.

Code examples: ScriptPathAttribute, AssemblyHasScriptsAttribute

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPathAttributeGenerator.cs` — ScriptPathAttributeGenerator.Execute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/ScriptPathAttribute.cs` — ScriptPathAttribute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/AssemblyHasScriptsAttribute.cs` — AssemblyHasScriptsAttribute

## Unicode data and properties

Generated tables are paired with implementations that expose properties and Unicode processing APIs.

The included code stores generated Unicode script, emoji, normalization, bidi, case, and general-character-property data for lookup-oriented services.

Code examples: UCPTrie, hb_unicode_funcs_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-ucd-table.hh` — _hb_ucd_sc_map
- Code: `thirdparty/icu4c/common/uchar_props_data.h` — propsTrie_index

## Character-encoding conversion

The partition includes UTF-7, UTF-8, UTF-16, UTF-32, ISO-2022, MBCS, and other converter implementations.

ICU models converters with shared static data, an implementation dispatch structure, contexts, and UTF-specific implementations.

Prerequisites: unicode-data

Code examples: UConverter, UConverterSharedData

Evidence:
- Code: `thirdparty/icu4c/common/ucnv_bld.h` — UConverterSharedData
- Code: `thirdparty/icu4c/common/ucnv_u8.cpp` — _UTF8Data

## Contextual completion contracts

The configuration records assertions over displayed or inserted completion text.

Completion fixture contracts pair a cursor-bearing script with configuration rules that include or exclude candidates.

Prerequisites: fixture-contracts

Code examples: Completion Test Configuration, Test Script Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/completion/builtin_enum/builtin_enum_autocomplete.gd` — _ready()
- Code: `modules/gdscript/tests/scripts/completion/builtin_enum/builtin_enum_autocomplete.cfg` — [output]

## Diagnostic expectations

This makes diagnostic behavior part of the tested interface.

Diagnostic fixture contracts preserve parser errors, analyzer warnings, and runtime errors as expected textual results.

Prerequisites: fixture-contracts

Code examples: Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/parser/warnings/integer_division.out` — INTEGER_DIVISION
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/unused_variable.out` — UNUSED_VARIABLE
- Code: `modules/gdscript/tests/scripts/runtime/errors/invalid_cast.out` — expected runtime errors

## Export presets

Preset configuration is target-platform-specific while retaining resource-selection policy.

An export preset selects indexed project files, applies include, exclude, and customized-file rules, and records output path, features, patches, and encryption options.

Prerequisites: filesystem-project-index

Code examples: EditorExportPreset

Evidence:
- Code: `editor/export/editor_export_preset.h` — EditorExportPreset
- Code: `editor/export/editor_export_preset.cpp` — EditorExportPreset

## Identifier spoof checking

The implementation includes confusable data, script sets, check results, and UTF-16 and UTF-8 API paths.

Unicode property checks and configured allowed character or locale sets are held by SpoofImpl and exposed through USpoofChecker APIs.

Prerequisites: unicode-data

Code examples: SpoofImpl, USpoofChecker

Evidence:
- Code: `thirdparty/icu4c/i18n/uspoof_impl.h` — SpoofImpl
- Code: `thirdparty/icu4c/i18n/unicode/uspoof.h` — USpoofChecker

## Language-server semantic fixtures

These inputs are distinct from completion assertions because they model symbols and source ranges.

Language-server fixture contracts provide nested declarations, local scopes, documentation comments, and lambdas for semantic-editor tests.

Prerequisites: fixture-contracts

Code examples: Test Script Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/lsp/class.gd` — Inner三.NestedInInner3
- Code: `modules/gdscript/tests/test_lsp.h` — test_resolve_symbol_at

## Unicode normalization

The code contains NFC, NFKC, FCD, and related Normalizer2 factory paths.

ICU implements Normalizer2 variants and loads normalization data into tries, extra mappings, and composition data.

Prerequisites: unicode-data

Code examples: Normalizer2Impl, UCPTrie

Evidence:
- Code: `thirdparty/icu4c/common/normalizer2impl.h` — Normalizer2Impl
- Code: `thirdparty/icu4c/common/normalizer2.cpp` — Normalizer2Factory::getNFCInstance

## Dynamic runtime values

This is the central interchange value across the object, container, parser, and binding code.

Variant is the tagged runtime-value representation used for construction, conversion, operators, calls, indexing, and member access.

Code examples: Variant

Evidence:
- Code: `core/variant/variant.h` — Variant
- Code: `core/variant/variant_construct.h` — VariantConstructor
- Code: `core/variant/variant_op.h` — CommonEvaluate

## Export orchestration

EditorExport is the registry joining target configuration with exporter extensions.

Export orchestration owns export presets, target platforms, and plugins, and maps a target platform to its runnable preset.

Prerequisites: export-presets

Code examples: EditorExport, EditorExportPreset

Evidence:
- Code: `editor/export/editor_export.h` — EditorExport
- Code: `editor/export/editor_export.cpp` — EditorExport

## GDScript editor services

The services are built only for editor builds.

Editor tooling consumes GDScript parser output to generate documentation, color syntax, extract translations, and provide completion-related behavior.

Prerequisites: gdscript-source-parsing

Code examples: GDScriptParser::Node, GDScriptTokenizer::Token

Evidence:
- Code: `modules/gdscript/editor/gdscript_docgen.cpp` — GDScriptDocGen AST traversal
- Code: `modules/gdscript/editor/gdscript_highlighter.h` — GDScriptSyntaxHighlighter

## Image quality metrics

The encoder exposes aggregate statistics and picture-distortion helpers for reporting and comparison.

PSNR, SSIM, local similarity, and squared-error routines assess encoded or reconstructed image planes.

Prerequisites: input-picture-planes

Code examples: WebPAuxStats, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/dsp/ssim.c` — VP8SSIMGet, VP8SSIM
- Code: `thirdparty/libwebp/src/enc/picture_psnr_enc.c` — WebPPictureDistortion
- Code: `thirdparty/libwebp/src/webp/encode.h` — struct WebPAuxStats, WebPPictureDistortion

## Project source migration

Migration is source-aware instead of a blind string replacement.

ProjectConverter3To4 retains whether each source line is a comment while applying named rename, conversion, and check-only passes to project source.

Code examples: SourceLine

Evidence:
- Code: `editor/project_upgrade/project_converter_3_to_4.h` — SourceLine
- Code: `editor/project_upgrade/project_converter_3_to_4.h` — ProjectConverter3To4
- Code: `editor/project_upgrade/renames_map_3_to_4.h` — RenamesMap3To4

## Scene-aware test context

Completion cases select this context through a configuration `scene` path.

A scene input gives a completion context containing nodes, attached scripts, resources, and unique names.

Prerequisites: completion-contexts

Code examples: Scene Fixture, Completion Test Configuration

Evidence:
- Code: `modules/gdscript/tests/scripts/completion/get_node/get_node.tscn` — [node name="UniqueA" type="Node" parent="UniqueNames"]
- Code: `modules/gdscript/tests/scripts/completion/get_node/literal_scene/dollar_class_scene.cfg` — [input] scene

## YUV/RGB conversion and chroma processing

Format packers also target RGBA, BGRA, RGB565, and RGBA4444 outputs.

Conversion kernels turn RGB or ARGB input into YUV planes, reconstruct RGB outputs, and upsample chroma with neighboring samples.

Prerequisites: input-picture-planes

Code examples: WebPPicture, WebPDecBuffer

Evidence:
- Code: `thirdparty/libwebp/src/dsp/yuv.h` — VP8YUVToR, VP8RGBToY
- Code: `thirdparty/libwebp/src/dsp/yuv.c` — WebPConvertARGBToYUV
- Code: `thirdparty/libwebp/src/dsp/upsampling.c` — WebPUpsamplersInit
- Code: `thirdparty/libwebp/src/enc/picture_csp_enc.c` — WebPPictureARGBToYUVA

## Primitive references

Motion-blur variants carry linear bounds and time-segment information.

A PrimRef supplies bounds plus geometry and primitive identifiers so builders can partition scene primitives independently of their source geometry layout.

Prerequisites: geometry-resources

Code examples: PrimRef, PrimRefMB, SubGridBuildData

Evidence:
- Code: `thirdparty/embree/kernels/builders/primref.h` — PrimRef
- Code: `thirdparty/embree/kernels/builders/primref_mb.h` — PrimRefMB
- Code: `thirdparty/embree/kernels/builders/primrefgen.cpp` — PrimRef construction

## Project settings and feature overrides

This is the central configuration store for project-level state.

ProjectSettings stores named Variant values with persistence metadata, feature overrides, autoload definitions, resource paths, and a change version.

Code examples: ProjectSettings, ProjectSettings::VariantContainer

Evidence:
- Code: `core/config/project_settings.h` — ProjectSettings
- Code: `core/config/project_settings.h` — ProjectSettings::VariantContainer

## Shader editing and language plugins

Language selection is represented by EditorShaderLanguagePlugin instances rather than hard-coded solely in one text editor.

Shader editing uses a ShaderEditorPlugin, a text shader editor, syntax highlighting, shader-language plugins, creation dialogs, and shader-file editing.

Evidence:
- Code: `editor/shader/shader_editor_plugin.h` — ShaderEditorPlugin and EditedShader
- Code: `editor/shader/text_shader_editor.h` — GDShaderSyntaxHighlighter, ShaderTextEditor, and TextShaderEditor
- Code: `editor/shader/editor_shader_language_plugin.h` — EditorShaderLanguagePlugin
- Code: `editor/shader/shader_create_dialog.h` — ShaderCreateDialog and ShaderTypeData
- External (official, unverified (source anchor not found)): [Shading language — Godot Engine documentation](https://docs.godotengine.org/en/stable/tutorials/shaders/shader_reference/shading_language.html), accessed 2026-07-15

## Dynamic arrays and dictionaries

Array and Dictionary store Variant values, while typed validators and typed wrappers constrain their element, key, or value types.

Prerequisites: dynamic-variant

Code examples: Array, Dictionary, Variant

Evidence:
- Code: `core/variant/array.h` — Array
- Code: `core/variant/dictionary.h` — Dictionary
- Code: `core/variant/container_type_validate.h` — ContainerTypeValidate
- Code: `core/variant/typed_dictionary.h` — TypedDictionary

## Text break iteration

The partition includes parsing, table building, runtime iteration, and dictionary or LSTM break-engine code.

ICU builds and executes rule-based break iterators, including dictionary-backed language break engines and compiled rule tables.

Prerequisites: unicode-data

Code examples: RuleBasedBreakIterator, LanguageBreakEngine

Evidence:
- Code: `thirdparty/icu4c/common/unicode/rbbi.h` — RuleBasedBreakIterator
- Code: `thirdparty/icu4c/common/brkeng.h` — LanguageBreakEngine

## Broad-phase collision detection

The code provides brute-force and quadtree broad-phase implementations.

Broad-phase collision detection uses Body world-space AABox bounds and collision filtering to generate candidate body pairs and answer coarse queries.

Prerequisites: body-lifecycle, collision-filtering

Code examples: Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/BroadPhase/BroadPhase.h` — BroadPhase
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/BroadPhase/BroadPhaseQuadTree.h` — BroadPhaseQuadTree
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/BroadPhase/QuadTree.h` — QuadTree

## Browser runtime adapters

The browser implementation is split between C++ platform objects and JavaScript libraries compiled into the Web build.

Web audio, display, input, fetch, MIDI, filesystem, and WebGL adapters call JavaScript libraries and exchange data with C++ Web platform classes.

Code examples: JavaScriptObjectImpl

Evidence:
- Code: `platform/web/audio_driver_web.h` — AudioDriverWorklet
- Code: `platform/web/http_client_web.h` — HTTPClientWeb
- Code: `platform/web/webmidi_driver.h` — MIDIDriverWebMidi
- Code: `platform/web/js/libs/library_godot_audio.js` — GodotAudio

## BVH construction

The generic builder evaluates leaf and split costs, while factory and specialized builders select node and primitive representations.

BVH construction partitions PrimRef records into nodes and leaves through configurable builder callbacks and settings.

Prerequisites: primitive-references

Code examples: BVHN, GeneralBVHBuilder, BVHNodeRecord

Evidence:
- Code: `thirdparty/embree/kernels/builders/bvh_builder_sah.h` — GeneralBVHBuilder
- Code: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build
- Code: `thirdparty/embree/kernels/bvh/bvh_node_ref.h` — BVHNodeRecord

## ENet transport and multiplayer adaptation

The connection implementation also invokes the engine Compression API for payload data.

ENetConnection and ENetPacketPeer wrap connection and peer behavior, while ENetMultiplayerPeer adapts the transport to MultiplayerPeer.

Code examples: ENetConnection, ENetPacketPeer, ENetMultiplayerPeer

Evidence:
- Code: `modules/enet/enet_connection.h` — ENetConnection
- Code: `modules/enet/enet_multiplayer_peer.h` — ENetMultiplayerPeer

## Virtual implementation extensions

The contracts are exposed as subclass APIs such as PhysicsServer2DExtension and TextServerExtension.

Extension classes declare required or optional virtual callbacks so external implementations can replace physics, rendering, text, and scripting behavior.

Prerequisites: runtime-metadata

Code examples: RenderSceneBuffersConfiguration, RID

Evidence:
- Code: `doc/classes/PhysicsServer2DExtension.xml` — PhysicsServer2DExtension
- Code: `doc/classes/TextServerExtension.xml` — TextServerExtension
- Code: `doc/classes/ScriptLanguageExtension.xml` — ScriptLanguageExtension

## GDScript language-server support

LSP-shaped request and response structures are defined in godot_lsp.h and served by JSONRPC-based protocol classes.

The language server reuses parser-derived symbol data to manage documents, resolve symbols, publish diagnostics, provide links, and build workspace edits.

Prerequisites: gdscript-source-parsing

Code examples: LSP::TextDocumentItem, LSP::DocumentSymbol, GDScriptWorkspace

Evidence:
- Code: `modules/gdscript/language_server/godot_lsp.h` — TextDocumentItem, DocumentSymbol, Diagnostic, and WorkspaceEdit
- Code: `modules/gdscript/language_server/gdscript_workspace.cpp` — GDScriptWorkspace::resolve_symbol

## Swapchain presentation

Presentation is represented by configuration and submission records in the generated API layer.

The binding pairs swapchain configuration with a presentation request whose wait semaphores, swapchains, and image indices are array inputs.

Code examples: SwapchainCreateInfoKHR, PresentInfoKHR, SwapchainKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct SwapchainCreateInfoKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct PresentInfoKHR

## BVH traversal

The code includes scalar, packet, hybrid, frustum, point-query, quantized-node, and oriented-node traversal paths.

Traversal tests the committed BVH with a ray query, orders candidate nodes, and calls primitive intersectors for reached leaves.

Prerequisites: bvh-construction, ray-query

Code examples: BVHNIntersector1, BVHNIntersectorKHybrid, TravRay

Evidence:
- Code: `thirdparty/embree/kernels/bvh/bvh_intersector1.h` — BVHNIntersector1
- Code: `thirdparty/embree/kernels/bvh/bvh_intersector_hybrid.h` — BVHNIntersectorKHybrid
- Code: `thirdparty/embree/kernels/bvh/node_intersector1.h` — TravRay

## ENet host and peer transport

ENet creates a host with a bounded peer allocation and manages peer state changes when connecting or disconnecting.

Code examples: ENetHost, ENetPeer

Evidence:
- Code: `thirdparty/enet/host.c` — enet_host_create
- Code: `thirdparty/enet/protocol.c` — enet_protocol_change_state
- Code: `thirdparty/enet/enet/enet.h` — _ENetHost
- Code: `thirdparty/enet/enet/enet.h` — _ENetPeer

## Extension API compatibility baselines

The baseline files record expected API changes such as added arguments, changed types, and removed APIs.

Version-pair directories retain expected extension-API validation differences, and the validator aggregates their relevant diagnostic lines.

Evidence:
- Code: `misc/scripts/validate_extension_api.sh` — get_expected_output
- Code: `misc/extension_api_validation/README.md` — expected output format

## Vulkan extension structure chains

The supplied platform records use the common Vulkan extensible-structure shape.

The structs expose `sType` and `pNext` fields so extension records can be carried through Vulkan creation and query calls.

Code examples: VkMetalSurfaceCreateInfoEXT, VkExternalFormatQNX

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_metal.h` — VkMetalSurfaceCreateInfoEXT
- Code: `thirdparty/vulkan/include/vulkan/vulkan_screen.h` — VkExternalFormatQNX
- External (official, unverified (source anchor not found)): [VkExternalFormatQNX Manual Page](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatQNX.html), accessed 2026-07-15

## Motion blur

The implementation has linear-bound primitive references, motion-blur builders, time-range node forms, and time-segment geometry access.

Motion-blur Geometry and BVH nodes interpolate bounds at ray-query time.

Prerequisites: bvh-construction, geometry-resources, ray-query

Code examples: PrimRefMB, AABBNodeMB_t, AABBNodeMB4D_t

Evidence:
- Code: `thirdparty/embree/kernels/builders/primref_mb.h` — PrimRefMB
- Code: `thirdparty/embree/kernels/bvh/bvh_node_aabb_mb.h` — AABBNodeMB_t
- Code: `thirdparty/embree/kernels/bvh/bvh_node_aabb_mb4d.h` — AABBNodeMB4D_t
- Code: `thirdparty/embree/kernels/common/default.h` — timeSegment

## Narrow-phase collision queries

Results are emitted through specialized collision collectors and dispatch handlers.

Narrow-phase collision queries test collision shapes for rays, points, overlaps, and casts after broad-phase candidates have been selected.

Prerequisites: broad-phase, collision-shapes, convex-collision

Code examples: Shape, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/NarrowPhaseQuery.h` — NarrowPhaseQuery
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionCollector.h` — CollisionCollector
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionDispatch.cpp` — ReversedCollector

## Resource formats and serialization

Format-specific code is separated from the common resource-loading service.

The resource-loading service uses binary, JSON, image, crypto, and extension resource-format loaders and savers.

Prerequisites: resource-loading

Code examples: Resource

Evidence:
- Code: `core/io/resource_format_binary.h` — ResourceFormatLoaderBinary
- Code: `core/io/json_resource_format.h` — ResourceFormatLoaderJSON
- Code: `core/io/image_resource_format.h` — ResourceFormatLoaderImage
- Code: `core/extension/gdextension_resource_format.h` — GDExtensionResourceLoader

## Scene construction and commit

The scene layer also checks whether attached geometry has changed before marking itself modified.

A Scene is created from a Device, retains Geometry instances, and commits a BVH used by query calls.

Prerequisites: bvh-construction, device-runtime, geometry-resources

Code examples: RTCScene, Scene

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h` — RTCScene
- Code: `thirdparty/embree/kernels/common/scene.h` — Scene
- Code: `thirdparty/embree/kernels/common/scene_verify.cpp` — Scene::checkIfModifiedAndSet

## Contact manifolds and warm-starting

The contact manager owns the cache used to retain collision-solver information.

Contact management converts narrow-phase collision results into contact constraints and caches manifolds, body pairs, and contact points between updates.

Prerequisites: narrow-phase

Code examples: CachedManifold, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/ContactConstraintManager.h` — CachedContactPoint, CachedManifold, CachedBodyPair, ManifoldCache
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/ContactConstraintManager.cpp` — ContactConstraintManager::ManifoldCache::Find

## Editor authoring workspaces

These are editor-facing authoring surfaces rather than runtime rendering backends.

The editor implements dock management, filesystem browsing, scene-tree editing, animation editing, asset-library browsing, audio editing, and debugger views as specialized controls and plugins.

Evidence:
- Code: `editor/docks/editor_dock_manager.h` — class EditorDockManager : public Object
- Code: `editor/animation/animation_track_editor.h` — class AnimationTrackEditor : public VBoxContainer
- Code: `editor/docks/scene_tree_dock.h` — class SceneTreeDock : public EditorDock
- Code: `editor/debugger/script_editor_debugger.h` — class ScriptEditorDebugger : public MarginContainer

## ENet wire commands

ENet wire commands select fixed protocol layout sizes and drive peer state transitions; host-peer transport is required to interpret the peer state.

Prerequisites: enet-host-peer-transport

Code examples: ENetPeer

Evidence:
- Code: `thirdparty/enet/enet/protocol.h` — ENetProtocolCommand
- Code: `thirdparty/enet/protocol.c` — commandSizes
- Code: `thirdparty/enet/protocol.c` — enet_protocol_command_size

## Explicit DRM synchronization objects

The schema models synchronization as externally exchanged timeline objects and 64-bit point values.

Wayland protocol objects import DRM synchronization timelines and attach acquire and release timeline points to a surface commit.

Prerequisites: wayland-protocol-objects

Code examples: wp_linux_drm_syncobj_timeline_v1

Evidence:
- Code: `thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml` — interface wp_linux_drm_syncobj_manager_v1
- Code: `thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml` — interface wp_linux_drm_syncobj_timeline_v1
- Code: `thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml` — set_acquire_point and set_release_point
- External (primary, unverified (source anchor not found)): [linux-drm-syncobj-v1 Wayland protocol schema](https://gitlab.freedesktop.org/wayland/wayland-protocols/-/blob/main/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml), accessed 2026-07-15

## GDExtension libraries

The extension runtime exposes a generated interface while retaining library and class lifecycle state in the engine.

GDExtension is a Resource that holds a loader, registered extension classes, initialization state, and library open, initialize, deinitialize, and close operations.

Code examples: GDExtension

Evidence:
- Code: `core/extension/gdextension.h` — GDExtension
- Code: `core/extension/gdextension_interface.cpp` — CallableCustomExtension
- Code: `core/extension/gdextension_library_loader.h` — GDExtensionLibraryLoader

## glTF scene-node hierarchy

GLTFDocument reconstructs parent and child relationships among indexed GLTF node records and attaches meshes, cameras, lights, skins, and skeleton membership through node indexes.

Prerequisites: gltf-asset-model

Code examples: GLTFNode, GLTFState

Evidence:
- Code: `modules/gltf/structures/gltf_node.h` — GLTFNode
- Code: `modules/gltf/gltf_document.cpp` — scene and node parsing paths

## Instancing

The configured API sets RTC_MAX_INSTANCE_LEVEL_COUNT to one.

Instance and InstanceArray are Geometry types attached to a Scene that contribute transformed primitive references and maintain instance-query context state.

Prerequisites: geometry-resources, scene-commit

Code examples: Instance, InstanceArray, InstancePrimitive

Evidence:
- Code: `thirdparty/embree/kernels/common/scene_instance.h` — Instance
- Code: `thirdparty/embree/kernels/common/scene_instance_array.h` — InstanceArray
- Code: `thirdparty/embree/kernels/common/instance_stack.h` — instance stack
- Code: `thirdparty/embree/include/embree4/rtcore_config.h` — RTC_MAX_INSTANCE_LEVEL_COUNT

## OpenXR composition layers

The scene layer has quad, cylinder, and equirect subclasses, while the extension owns swapchain and composition-layer state.

Composition-layer scene nodes and an extension submit specialized layer descriptions to the OpenXR runtime.

Prerequisites: openxr-runtime-integration

Evidence:
- Code: `modules/openxr/scene/openxr_composition_layer.h` — class OpenXRCompositionLayer : public Node3D
- Code: `modules/openxr/scene/openxr_composition_layer_quad.h` — class OpenXRCompositionLayerQuad : public OpenXRCompositionLayer
- Code: `modules/openxr/extensions/openxr_composition_layer_extension.h` — OpenXRCompositionLayerExtension::SwapchainState and CompositionLayer

## OpenXR render models

Render-model data is separate from interaction-profile suggestions, allowing runtime-provided device models to drive presentation.

The render-model extension tracks runtime render models, and scene nodes display models or manage their activation.

Prerequisites: openxr-runtime-integration

Code examples: OpenXRRenderModelData

Evidence:
- Code: `modules/openxr/extensions/openxr_render_model_extension.h` — class OpenXRRenderModelData; class OpenXRRenderModelExtension
- Code: `modules/openxr/scene/openxr_render_model.h` — class OpenXRRenderModel : public Node3D
- Code: `modules/openxr/scene/openxr_render_model_manager.h` — class OpenXRRenderModelManager : public Node3D

## Property-list and scene-property helpers

These helpers bridge reflective property operations with scene state.

PropertyListHelper resolves slash-delimited property names to property records, while PropertyUtils compares properties and walks scene-state pack data.

Prerequisites: scene-serialization

Evidence:
- Code: `scene/property_list_helper.h` — class PropertyListHelper and PropertyListHelper::Property
- Code: `scene/property_list_helper.cpp` — PropertyListHelper::_get_property
- Code: `scene/property_utils.cpp` — PropertyUtils scene-state PackState accesses

## Replicated property synchronization

The synchronizer builds watchers from configured paths and participates in replication start and stop through object configuration calls.

Synchronizers observe configured properties and forward sync property lists to scene replication.

Prerequisites: scene-replication-configuration

Code examples: SceneReplicationConfig

Evidence:
- Code: `modules/multiplayer/multiplayer_synchronizer.h` — MultiplayerSynchronizer
- Code: `modules/multiplayer/multiplayer_synchronizer.cpp` — MultiplayerSynchronizer
- Code: `modules/multiplayer/scene_replication_interface.cpp` — SceneReplicationInterface
- External (official, unverified (source anchor not found)): [MultiplayerSynchronizer — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_multiplayersynchronizer.html), accessed 2026-07-15

## Resource and server identifiers

The two identifiers intentionally serve different lifetimes: project references versus live server objects.

ResourceUID maps project resource identifiers to paths, while an RID is an opaque session-only handle for a low-level server resource.

Prerequisites: resource-formats

Code examples: ResourceUID, RID

Evidence:
- Code: `doc/classes/ResourceUID.xml` — ResourceUID
- Code: `doc/classes/RID.xml` — RID

## Runtime configuration

The main executable and a Godot project file both provide concrete configuration inputs.

Runtime configuration reads project settings such as the main scene and boot-splash options; the application project file also declares its main scene.

Evidence:
- Code: `main/main.cpp` — Main::start
- Code: `platform/android/java/app/src/instrumented/assets/project.godot` — [application]

## Main loop, OS, and time services

OS hosts platform runtime services, MainLoop defines the loop-facing object type, and Time exposes date and time behavior.

Evidence:
- Code: `core/os/os.h` — OS
- Code: `core/os/main_loop.h` — MainLoop
- Code: `core/os/time.h` — Time

## 2D and 3D scene authoring

The partition implements distinct editing surfaces rather than a single dimension-agnostic viewport.

Scene authoring is split between CanvasItemEditor for 2D work and Node3DEditor with Node3DEditorViewport for 3D work.

Evidence:
- Code: `editor/scene/canvas_item_editor_plugin.h` — CanvasItemEditor : public VBoxContainer
- Code: `editor/scene/3d/node_3d_editor_plugin.h` — Node3DEditor : public VBoxContainer
- Code: `editor/scene/3d/node_3d_editor_viewport.h` — Node3DEditorViewport : public Control

## Tile authoring

Tile editing is implemented as several coordinated editor tools rather than a single monolithic panel.

Tile tools edit atlas sources, per-tile data, terrain data, proxies, map layers, and scene-collection sources.

Evidence:
- Code: `editor/scene/2d/tiles/tile_set_atlas_source_editor.h` — TileSetAtlasSourceEditor
- Code: `editor/scene/2d/tiles/tile_data_editors.h` — TileDataEditor
- Code: `editor/scene/2d/tiles/tile_map_layer_editor.h` — TileMapLayerEditor
- Code: `editor/scene/2d/tiles/tile_proxies_manager_dialog.h` — TileProxiesManagerDialog

## ASTC block coding

ASTC logic is used both by the Basis transcoder and its HDR intermediate paths.

ASTC helpers represent physical and logical blocks with endpoint ranges, weight grids, partitions, and bit-level packing.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_astc_helpers.h` — astc_helpers::astc_block
- Code: `thirdparty/basis_universal/transcoder/basisu_astc_helpers.h` — astc_helpers::log_astc_block
- Code: `thirdparty/basis_universal/transcoder/basisu_astc_hdr_core.h` — astc_blk

## CPU-specialized DSP backends

The scalar sources establish algorithms while architecture files express equivalent work with vector intrinsics or target assembly.

DSP function families have scalar, SSE2/SSE4.1/AVX2, NEON, MSA, and MIPS implementations that are compile-time selected for target capabilities.

Evidence:
- Code: `thirdparty/libwebp/src/dsp/enc.c` — VP8EncDspInit
- Code: `thirdparty/libwebp/src/dsp/enc_sse2.c` — SSE2 transform and quantization kernels
- Code: `thirdparty/libwebp/src/dsp/enc_neon.c` — NEON transform and quantization kernels
- Code: `thirdparty/libwebp/src/dsp/enc_msa.c` — MSA transform and prediction kernels

## Engine bootstrap

`Main` is the common startup layer before a selected platform host runs the engine.

Engine bootstrap consumes runtime configuration to initialize execution and select the configured main scene.

Prerequisites: runtime-configuration

Evidence:
- Code: `main/main.cpp` — Main::setup
- Code: `main/main.cpp` — Main::start
- Code: `main/main.h` — Main

## Godot ENet socket adaptation

The Godot ENet socket adapter supplies UDP and DTLS socket classes to the transport; ENet wire commands provide this adapter's transport-facing purpose.

Prerequisites: enet-wire-commands

Evidence:
- Code: `thirdparty/enet/enet_godot.cpp` — ENetGodotSocket
- Code: `thirdparty/enet/enet_godot.cpp` — ENetUDP
- Code: `thirdparty/enet/enet_godot.cpp` — ENetDTLSClient
- Code: `thirdparty/enet/enet_godot.cpp` — ENetDTLSServer

## Intersection results

The intersection distance is maintained through the ray state used by traversal and intersector epilogues.

An RTCRayHit combines a ray query with geometric-normal, barycentric-coordinate, primitive-ID, geometry-ID, and instance-stack hit results.

Prerequisites: ray-query

Code examples: RTCHit, RTCRayHit

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCHit
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCRayHit
- Code: `thirdparty/embree/kernels/geometry/intersector_epilog.h` — Intersect1Epilog1

## Input action and shortcut configuration

The implementation supports keyboard, mouse, and joypad event capture in a shared line-edit control.

Input action configuration depends on editor and project settings because ActionMapEditor, event search, event capture, and input-event configuration edit action-event mappings.

Prerequisites: editor-and-project-settings

Evidence:
- Code: `editor/settings/action_map_editor.h` — ActionMapEditor and ActionInfo
- Code: `editor/settings/input_event_configuration_dialog.h` — InputEventConfigurationDialog
- Code: `editor/settings/event_listener_line_edit.cpp` — EventListenerLineEdit input event handling
- External (official, unverified (source anchor not found)): [Using InputEvent — Godot Engine documentation](https://docs.godotengine.org/en/stable/tutorials/inputs/inputevent.html), accessed 2026-07-15

## Project catalog and selection

The project manager operates on structured records rather than treating projects as paths alone.

ProjectList models known projects with name, path, tags, main scene, version, unsupported features, last-edit time, and favorite state.

Code examples: ProjectCatalog, ProjectCatalogItem

Evidence:
- Code: `editor/project_manager/project_list.h` — ProjectList::Item
- Code: `editor/project_manager/project_list.cpp` — ProjectList

## RenderingDevice descriptors and handles

RD-prefixed descriptor objects carry configuration into the low-level rendering API.

RenderingDevice creates and consumes RID handles for buffers, textures, shaders, uniforms, pipelines, framebuffers, and acceleration structures.

Prerequisites: resource-identifiers

Code examples: RDTextureFormat, RDUniform, RDAccelerationStructureGeometry

Evidence:
- Code: `doc/classes/RenderingDevice.xml` — RenderingDevice
- Code: `doc/classes/RDUniform.xml` — RDUniform
- Code: `doc/classes/RDTextureFormat.xml` — RDTextureFormat

## Resource dependency resolution

Stored dependency values can carry type and UID information that is resolved before presentation.

Dependency resolution uses file-index entries and ResourceUID mappings to expose each indexed resource's current dependency paths.

Prerequisites: filesystem-project-index

Code examples: EditorFileSystemDirectory::FileInfo

Evidence:
- Code: `editor/file_system/editor_file_system.cpp` — EditorFileSystemDirectory::get_file_deps
- Code: `editor/file_system/editor_file_system.cpp` — EditorFileSystem::_scan_for_uid_directory

## Rigid-body motion

MotionProperties is the mutable dynamic state associated with movable rigid bodies.

Rigid-body motion uses Body transform state plus MotionProperties to model static, kinematic, and dynamic behavior, permitted degrees of freedom, mass, inertia, and velocities.

Prerequisites: body-lifecycle

Code examples: Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionType.h` — EMotionType
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionProperties.cpp` — MotionProperties::SetMassProperties
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/AllowedDOFs.h` — EAllowedDOFs

## Scene tree and signal connections

The source treats structural scene navigation and signal wiring as editor services.

Scene authoring is accompanied by a SceneTreeEditor and a ConnectionsDock that inspect nodes, signals, methods, connection arguments, and bound values.

Prerequisites: scene-authoring

Evidence:
- Code: `editor/scene/scene_tree_editor.h` — SceneTreeEditor : public Control
- Code: `editor/scene/connections_dialog.h` — ConnectDialog::ConnectionData and ConnectionsDock
- Code: `editor/scene/connections_dialog.cpp` — signal_name and signal_args metadata handling
- External (official, unverified (source anchor not found)): [Using signals — Godot Engine documentation](https://docs.godotengine.org/en/stable/getting_started/step_by_step/signals.html), accessed 2026-07-15

## Tile libraries, cells, and patterns

Tile sources can expose atlas-backed or scene-backed tiles.

A TileSet Resource owns tile sources, identifies tiles by source, atlas-coordinate, and alternative IDs, and supplies tile data to TileMapLayer cells and TileMapPattern copies.

Prerequisites: resource-formats

Code examples: TileSet, TileData

Evidence:
- Code: `doc/classes/TileSet.xml` — TileSet
- Code: `doc/classes/TileData.xml` — TileData
- Code: `doc/classes/TileMapPattern.xml` — TileMapPattern

## Vulkan instance setup

This is the outer configuration record for the Vulkan API surface represented by the supplied headers.

The binding models instance setup as an `InstanceCreateInfo` record containing optional application metadata plus enabled layer and extension name arrays.

Code examples: InstanceCreateInfo, ApplicationInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct InstanceCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct ApplicationInfo

## Wraparound network time

ENet compares and subtracts time values with an overflow threshold so ordering and differences remain defined across its configured time wraparound.

Evidence:
- Code: `thirdparty/enet/enet/time.h` — ENET_TIME_OVERFLOW
- Code: `thirdparty/enet/enet/time.h` — ENET_TIME_LESS
- Code: `thirdparty/enet/enet/time.h` — ENET_TIME_DIFFERENCE

## Android platform host

The Android partition has both native C++ platform classes and JVM-facing host classes.

Android platform hosting starts and manages the native engine from Android activity and library code, continuing the engine bootstrap on the mobile host.

Prerequisites: engine-bootstrap

Evidence:
- Code: `platform/android/os_android.h` — OS_Android
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt` — Godot
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotActivity.kt` — GodotActivity

## Constraint solving

ConstraintSettings are serializable configuration objects; runtime Constraint objects participate in solver islands.

Constraint solving applies impulses to rigid-body motion and reuses contact manifolds while supporting point, distance, hinge, slider, fixed, gear, pulley, cone, six-degree-of-freedom, path, and swing-twist constraints.

Prerequisites: contact-management, rigid-body-motion

Code examples: ConstraintSettings, TwoBodyConstraint

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/Constraint.h` — ConstraintSettings and Constraint
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/TwoBodyConstraint.cpp` — TwoBodyConstraint::BuildIslands
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/SixDOFConstraint.h` — SixDOFConstraintSettings and SixDOFConstraint

## Frame timing and physics stepping

The timer synchronization layer separates frame-duration handling from main startup.

Frame timing calculates a process delta and bounded physics-step count after engine bootstrap determines the active physics tick configuration.

Prerequisites: engine-bootstrap

Code examples: MainFrameTime

Evidence:
- Code: `main/main_timer_sync.h` — MainFrameTime
- Code: `main/main_timer_sync.h` — MainTimerSync
- Code: `main/main_timer_sync.cpp` — MainTimerSync::advance

## Ray packets and streams

Packet traversal and packet primitive intersectors reuse the same query concepts across multiple lanes.

Packet execution packs several ray queries and their hit results into 4-, 8-, 16-, or N-wide layouts, with array-of-structures and structure-of-arrays stream readers.

Prerequisites: hit-results, ray-query

Code examples: RTCRay4, RTCRayHit16, RTCRayNt, RayStreamSOA

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCRay4
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCRayHitNt
- Code: `thirdparty/embree/kernels/common/ray.h` — RayStreamSOA
- Code: `thirdparty/embree/kernels/bvh/node_intersector_packet.h` — TravRayK

## Performance monitors

The `Performance` object and `MonitorCall` class are in the main runtime partition.

Performance monitors collect engine counters and execute registered monitor calls after engine bootstrap.

Prerequisites: engine-bootstrap

Evidence:
- Code: `main/performance.h` — Performance
- Code: `main/performance.h` — MonitorCall
- Code: `main/performance.cpp` — Performance

## Primitive intersection

The indexed code includes triangle, quad, curve, line, cone, cylinder, disc, grid, object, and instance intersectors.

Primitive intersectors test a ray query against Geometry representations and update RTCRayHit through common intersection or occlusion epilogues.

Prerequisites: geometry-resources, hit-results, ray-query

Code examples: TriangleMIntersectorKMoeller, QuadMIntersector1MoellerTrumbore, ConeCurveIntersector1

Evidence:
- Code: `thirdparty/embree/kernels/geometry/quad_intersector_moeller.h` — QuadMIntersector1MoellerTrumbore
- Code: `thirdparty/embree/kernels/geometry/curve_intersector_virtual.h` — VirtualCurveIntersector
- Code: `thirdparty/embree/kernels/geometry/coneline_intersector.h` — ConeCurveIntersector1
- Code: `thirdparty/embree/kernels/geometry/intersector_epilog.h` — Intersect1Epilog1

## Resource-backed assets

The resource family is the common content representation used throughout the inspected scene partition.

Resource-derived classes represent reusable scene content such as textures, meshes, materials, themes, animations, shapes, and packed scenes.

Code examples: PackedScene, Animation, TileMapPattern

Evidence:
- Code: `scene/resources/packed_scene.h` — class PackedScene : public Resource
- Code: `scene/resources/animation.h` — class Animation : public Resource
- Code: `scene/resources/texture.h` — class Texture : public Resource
- Code: `scene/resources/theme.h` — class Theme : public Resource

## Scene render data and buffers

The renderer receives a compact frame-oriented view rather than owning scene objects directly.

RenderDataRD gathers visible scene-instance pointers and RID lists for lights, probes, decals, lightmaps, and fog volumes, while RenderSceneBuffersRD supplies named GPU scene textures.

Prerequisites: rendering-device-resources

Code examples: RenderDataRD

Evidence:
- Code: `servers/rendering/renderer_rd/storage_rd/render_data_rd.h` — RenderDataRD
- Code: `servers/rendering/renderer_rd/storage_rd/render_scene_buffers_rd.h` — RenderSceneBuffersRD

## Scene graph and lifecycle

Nodes carry parent, owner, children, name, tree, processing, and RPC configuration state.

A scene tree arranges Object-derived Node instances into a parent-child hierarchy that SceneTree owns and updates.

Prerequisites: object-model

Code examples: Node, SceneState

Evidence:
- Code: `scene/main/node.h` — Node::Data
- Code: `scene/main/node.h` — Node::add_child
- Code: `scene/main/scene_tree.h` — SceneTree

## Shader programs and material parameters

The high-level Shader API is distinct from RDShaderFile and RDShaderSPIRV, which are RenderingDevice-facing representations.

A Shader resource supplies custom shader source, ShaderInclude supplies reusable source fragments, and ShaderMaterial binds a Shader with parameter values.

Prerequisites: rendering-device-resources

> Prerequisite evidence for `rendering-device-resources` needs additional evidence: the declared quote did not appear verbatim in this topic description.

Code examples: ShaderMaterial, RDPipelineSpecializationConstant

Evidence:
- Code: `doc/classes/Shader.xml` — Shader
- Code: `doc/classes/ShaderInclude.xml` — ShaderInclude
- Code: `doc/classes/ShaderMaterial.xml` — ShaderMaterial
- Code: `doc/classes/RDShaderSPIRV.xml` — RDShaderSPIRV

## Mbed TLS-backed crypto and transport

The Mbed TLS module supplies Godot Crypto, certificate, TLS context, DTLS server, and TLS peer implementations.

Evidence:
- Code: `modules/mbedtls/crypto_mbedtls.h` — CryptoMbedTLS, CryptoKeyMbedTLS, X509CertificateMbedTLS, and HMACContextMbedTLS
- Code: `modules/mbedtls/dtls_server_mbedtls.h` — DTLSServerMbedTLS
- Code: `modules/mbedtls/stream_peer_mbedtls.h` — StreamPeerMbedTLS

## Viewport render-frame data

These are internal rendering-server objects rather than general script-created scene objects.

RenderData and RenderSceneData exist for a viewport frame, while RenderSceneBuffersConfiguration configures buffers recreated when a viewport changes.

Prerequisites: rendering-device-resources

> Prerequisite evidence for `rendering-device-resources` needs additional evidence: the declared quote did not appear verbatim in this topic description.

Code examples: RenderSceneBuffersConfiguration, RenderData

Evidence:
- Code: `doc/classes/RenderData.xml` — RenderData
- Code: `doc/classes/RenderSceneBuffers.xml` — RenderSceneBuffers
- Code: `doc/classes/RenderSceneBuffersConfiguration.xml` — RenderSceneBuffersConfiguration

## Audio buses, streams, and effects

Effect classes create reference-counted processing instances.

AudioServer owns buses with channels and effects, while Resource-derived streams, effects, and bus layouts define reusable audio configuration.

Prerequisites: resource-assets

Code examples: AudioBusLayout

Evidence:
- Code: `servers/audio/audio_server.h` — AudioServer::Bus, AudioServer::Channel, AudioServer::Effect, and AudioBusLayout
- Code: `servers/audio/audio_stream.h` — class AudioStream : public Resource
- Code: `servers/audio/audio_effect.h` — AudioEffectInstance : public RefCounted and AudioEffect : public Resource

## Canvas and viewport presentation

CanvasItem is the 2D visual branch; Window is a Viewport subclass.

CanvasItem, CanvasLayer, Viewport, and Window extend the Node tree with visual presentation and window-facing state.

Prerequisites: node-hierarchy

Code examples: Node

Evidence:
- Code: `scene/main/canvas_item.h` — class CanvasItem : public Node
- Code: `scene/main/canvas_layer.h` — class CanvasLayer : public Node
- Code: `scene/main/viewport.h` — class Viewport : public Node
- Code: `scene/main/window.h` — class Window : public Viewport

## Expression parsing and evaluation

Expression defines token and expression-node types and evaluates expression nodes against Variant inputs and built-in functions.

Prerequisites: dynamic-variant

Code examples: Variant

Evidence:
- Code: `core/math/expression.h` — Expression
- Code: `core/math/expression.h` — Expression::ENode
- Code: `core/math/expression.cpp` — Expression::execute

## Intersection and occlusion filters

User geometry also exposes externally supplied intersect and occluded function argument structures.

Geometry-installed and query-context filters process candidate hit results before they update an RTCRayHit or report occlusion.

Prerequisites: hit-results

Code examples: RTCFilterFunctionNArguments, RTCIntersectFunctionNArguments, RTCOccludedFunctionNArguments

Evidence:
- Code: `thirdparty/embree/kernels/geometry/filter.h` — runIntersectionFilter1Helper
- Code: `thirdparty/embree/kernels/common/accelset.h` — IntersectFunctionNArguments
- Code: `thirdparty/embree/include/embree4/rtcore_common.h` — RTCFilterFunctionNArguments

## Input events and actions

The model covers physical keyboard, mouse, touch, joystick, gesture, and manually generated action events.

InputMap associates named actions with InputEvent instances, and nodes receive input-event callbacks.

Prerequisites: scene-tree

Code examples: InputEvent

Evidence:
- Code: `core/input/input.h` — Input::parse_input_event
- Code: `doc/classes/InputEvent.xml` — InputEvent
- Code: `doc/classes/InputMap.xml` — InputMap::action_add_event

## MP3 audio resources

AudioStreamMP3 and AudioStreamPlaybackMP3 form the runtime pair, while ResourceImporterMP3 handles editor import.

The MP3 module loads MP3 stream data, supplies resampled playback, and imports MP3 assets.

Evidence:
- Code: `modules/mp3/audio_stream_mp3.h` — AudioStreamMP3
- Code: `modules/mp3/audio_stream_mp3.h` — AudioStreamPlaybackMP3
- Code: `modules/mp3/resource_importer_mp3.h` — ResourceImporterMP3
- External (official, unverified (source anchor not found)): [AudioStreamMP3 — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_audiostreammp3.html), accessed 2026-07-15

## Packets, HTTP, JSON, and JSON-RPC

These APIs separate transport helpers from the high-level multiplayer interface.

Packet peers transmit raw bytes or Variant values, while JSON and JSON-RPC map values into external message documents.

Prerequisites: dynamic-values

Code examples: JSON-RPC document, Variant

Evidence:
- Code: `doc/classes/PacketPeer.xml` — PacketPeer
- Code: `doc/classes/HTTPClient.xml` — HTTPClient.query_string_from_dict
- Code: `doc/classes/JSONRPC.xml` — JSONRPC

## ObjectDB snapshots

The runtime collector transports snapshot data while editor classes turn it into inspectable data and controls.

The debug-only object database profiler collects serialized object snapshots and renders summary, class, node, object, and ref-counted views.

Evidence:
- Code: `modules/objectdb_profiler/snapshot_collector.h` — SnapshotCollector
- Code: `modules/objectdb_profiler/editor/snapshot_data.h` — GameStateSnapshot
- Code: `modules/objectdb_profiler/editor/data_viewers/snapshot_view.h` — SnapshotView
- Code: `modules/objectdb_profiler/editor/objectdb_profiler_panel.h` — ObjectDBProfilerPanel

## Ogg packet sequences

The module defines a Resource plus a ref-counted playback companion.

OggPacketSequence persists packet data, page granule positions, and sample-rate information for sequence playback.

Code examples: OggPacketSequence

Evidence:
- Code: `modules/ogg/ogg_packet_sequence.h` — OggPacketSequence
- Code: `modules/ogg/ogg_packet_sequence.h` — OggPacketSequencePlayback
- External (official, unverified (source anchor not found)): [OggPacketSequence — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_oggpacketsequence.html), accessed 2026-07-15

## Ogg pages, packets, and bit packing

This is the shared container layer used by the supplied Vorbis and Theora implementations.

The Ogg implementation packs variable-sized words into octet streams and maintains stream state that frames packets into page headers and bodies.

Code examples: Ogg Stream State, Ogg Packet

Evidence:
- Code: `thirdparty/libogg/bitwise.c` — oggpack_buffer
- Code: `thirdparty/libogg/ogg/ogg.h` — ogg_stream_state
- External (primary, unverified (source anchor not found)): [Page Multiplexing and Ordering in a Physical Ogg Stream](https://www.xiph.org/ogg/doc/ogg-multiplex.html), accessed 2026-07-15

## Ogg Vorbis audio streams

The module also provides an editor resource importer.

AudioStreamOggVorbis and its playback class provide an audio-stream resource and resampled playback implementation.

Code examples: AudioStreamOggVorbis, AudioStreamPlaybackOggVorbis

Evidence:
- Code: `modules/vorbis/audio_stream_ogg_vorbis.h` — AudioStreamOggVorbis
- Code: `modules/vorbis/audio_stream_ogg_vorbis.h` — AudioStreamPlaybackOggVorbis
- Code: `modules/vorbis/resource_importer_ogg_vorbis.h` — ResourceImporterOggVorbis

## Packed scene serialization

SceneState separates reusable name, Variant, path, node, and connection tables from the PackedScene resource wrapper.

A PackedScene stores a resource-backed SceneState whose node entries, property values, and connection entries reconstruct a node hierarchy.

Prerequisites: resources, scene-tree

Code examples: PackedScene, SceneState

Evidence:
- Code: `scene/resources/packed_scene.h` — SceneState::NodeData
- Code: `scene/resources/packed_scene.h` — SceneState::ConnectionData
- Code: `scene/resources/packed_scene.h` — PackedScene

## Physics shapes, objects, and collision reports

The implementation provides parallel 2D and 3D object, shape, joint, and movement-report APIs.

2D and 3D CollisionObject nodes own Shape resources through shape owners; body movement returns KinematicCollision results.

Prerequisites: resources, scene-tree

Code examples: Shape2D, CollisionShape2D, KinematicCollision2D

Evidence:
- Code: `doc/classes/CollisionObject2D.xml` — CollisionObject2D
- Code: `doc/classes/CollisionShape2D.xml` — CollisionShape2D
- Code: `doc/classes/KinematicCollision2D.xml` — KinematicCollision2D

## Ray–primitive intersection

Embree provides separate kernels for several primitive families while retaining the same hit-candidate flow.

Ray-primitive intersection evaluates rays against triangle, sphere, and rounded-line geometry and forwards valid hit candidates to an epilog.

Code examples: SubGrid, TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h` — MoellerTrumboreIntersector1
- Code: `thirdparty/embree/kernels/geometry/sphere_intersector.h` — SphereIntersector1
- Code: `thirdparty/embree/kernels/geometry/roundline_intersector.h` — RoundLinearCurveIntersector1

## Raycast occlusion culling

Embree-backed lightmap and static raycaster classes supplement the central occlusion-cull implementation.

The raycast module represents occluders, instances, scenarios, and raycast thread data to provide renderer-scene occlusion culling.

Code examples: RaycastOccluder

Evidence:
- Code: `modules/raycast/raycast_occlusion_cull.h` — RaycastOcclusionCull; Occluder, OccluderInstance, Scenario, RaycastThreadData
- Code: `modules/raycast/static_raycaster_embree.h` — class StaticRaycasterEmbree : public StaticRaycaster
- Code: `modules/raycast/lightmap_raycaster_embree.h` — class LightmapRaycasterEmbree : public LightmapRaycaster

## Asynchronous resource previews

Preview generation is a shared inspector-side service rather than a feature embedded in each resource editor.

Resource previews use generators selected by handles, queue source paths or in-memory resources, and cache generated full and small textures with metadata.

Code examples: PreviewRequest, ResourcePreviewCacheEntry

Evidence:
- Code: `editor/inspector/editor_resource_preview.h` — EditorResourcePreviewGenerator
- Code: `editor/inspector/editor_resource_preview.h` — EditorResourcePreview::QueueItem
- Code: `editor/inspector/editor_resource_preview.cpp` — EditorResourcePreview::_thread

## Scene graph nodes

2D nodes inherit from CanvasItem through Node2D, while 3D nodes inherit from Node through Node3D.

Node2D and Node3D are scene graph node bases for specialized runtime behavior.

Code examples: Node2D, Node3D

Evidence:
- Code: `scene/2d/node_2d.h` — Node2D
- Code: `scene/3d/node_3d.h` — Node3D

## Scene importing

This is the central pipeline for bringing 3D scene formats into the editor.

Scene importers expose extensions, options, flags, and an import operation that produces a Node-based scene.

Code examples: ColladaParseState

Evidence:
- Code: `editor/import/3d/resource_importer_scene.h` — EditorSceneFormatImporter
- Code: `editor/import/3d/resource_importer_scene.h` — ResourceImporterScene

## SceneTree execution

This is the runtime coordinator above individual nodes.

A SceneTree schedules the Node hierarchy and maintains scene-level timer and tween processing collections.

Prerequisites: node-hierarchy

Code examples: Node

Evidence:
- Code: `scene/main/scene_tree.h` — class SceneTree
- Code: `scene/main/scene_tree.cpp` — SceneTree timer and tween processing blocks

## Screen-space and environment effects

These classes are effect-oriented services layered over scene buffers and shader pipelines.

RD effects and environment services allocate and process scene buffers for luminance, temporal anti-aliasing, reflections, ambient effects, fog, global illumination, sky, tone mapping, and scaling.

Prerequisites: scene-data-and-buffers, shader-source-toolchain

Evidence:
- Code: `servers/rendering/renderer_rd/effects/luminance.h` — Luminance, LuminanceBuffers
- Code: `servers/rendering/renderer_rd/effects/taa.h` — TAA
- Code: `servers/rendering/renderer_rd/effects/ss_effects.h` — SSEffects
- Code: `servers/rendering/renderer_rd/environment/fog.h` — Fog, VolumetricFog
- Code: `servers/rendering/renderer_rd/environment/gi.h` — GI, SDFGI
- Code: `servers/rendering/renderer_rd/environment/sky.h` — SkyRD

## Skeletons, animation, and ragdolls

SkeletonPose, SkeletalAnimation, SkeletonMapper, and Ragdoll form the animation-to-physics integration surface.

Skeletons retain named joints and parent relationships, animations produce joint states, and ragdolls connect those joints to Body parts and two-body constraints.

Prerequisites: body-lifecycle, constraints

Code examples: Skeleton, Joint, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Skeleton/Skeleton.cpp` — Skeleton::CalculateParentJointIndices and AreJointsCorrectlyOrdered
- Code: `thirdparty/jolt_physics/Jolt/Skeleton/SkeletalAnimation.h` — SkeletalAnimation
- Code: `thirdparty/jolt_physics/Jolt/Physics/Ragdoll/Ragdoll.h` — RagdollSettings and Ragdoll

## SPIR-V generation

The lowering code, builder, and IR classes together form the shader binary-generation path.

A traverser lowers the front end's typed intermediate tree to SPIR-V instructions and a module/function/block graph.

Prerequisites: shader-language-front-end

Code examples: spv::Module, spv::Function, spv::Block

Evidence:
- Code: `thirdparty/glslang/SPIRV/GlslangToSpv.cpp` — TGlslangToSpvTraverser
- Code: `thirdparty/glslang/SPIRV/SpvBuilder.cpp` — spv::Builder
- Code: `thirdparty/glslang/SPIRV/spvIR.h` — spv::Module
- External (official, verified): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-16

## Temporal upscaling dispatch

The implementation contains temporal resource selection and named shader-pass sources.

Vendored FSR2 dispatch code selects alternating frame resources, handles accumulation reset, and issues compute work based on render and display dimensions.

Evidence:
- Code: `thirdparty/amd-fsr2/ffx_fsr2.cpp` — fsr2Dispatch
- Code: `thirdparty/amd-fsr2/shaders/ffx_fsr2_accumulate_pass.glsl` — accumulation compute pass

## Theora video streams

The stream, loader, playback, and writer types separate runtime playback from editor export behavior.

The Theora module defines a video-stream resource, playback implementation, resource loader, and OGV movie writer.

Code examples: VideoStreamTheora

Evidence:
- Code: `modules/theora/video_stream_theora.h` — VideoStreamPlaybackTheora, VideoStreamTheora, ResourceFormatLoaderTheora
- Code: `modules/theora/editor/movie_writer_ogv.h` — class MovieWriterOGV : public MovieWriter

## TLS connection and session state

Client, TLS 1.2, TLS 1.3, message, and common TLS units implement different portions of the protocol path.

The TLS state machine performs cryptographic operations while carrying configuration, negotiated session, handshake, record, and protocol state in mbedtls_ssl_context.

Prerequisites: cryptographic-operation-dispatch

Code examples: mbedtls_ssl_session

Evidence:
- Code: `thirdparty/mbedtls/include/mbedtls/ssl.h` — struct mbedtls_ssl_session; struct mbedtls_ssl_config; struct mbedtls_ssl_context
- Code: `thirdparty/mbedtls/library/ssl_tls.c` — mbedtls_ssl_session_save; mbedtls_ssl_session_load
- Code: `thirdparty/mbedtls/library/ssl_client.h` — mbedtls_ssl_write_client_hello
- External (official, verified): [Mbed TLS documentation hub](https://mbed-tls.readthedocs.io/en/latest/), accessed 2026-07-16

## Control-tree user interfaces

The same tree model supports widgets, dialogs, menus, graph editors, and editor controls.

Control-derived nodes compose user interfaces, and Container nodes automatically arrange their Control children.

Prerequisites: scene-tree

Code examples: Node

Evidence:
- Code: `doc/classes/Control.xml` — Control
- Code: `doc/classes/Container.xml` — Container
- Code: `doc/classes/GraphEdit.xml` — GraphEdit

## Variant text parsing and writing

VariantParser and VariantWriter serialize Variant values as String text through stream, token, tag, and resource-parser abstractions.

Prerequisites: dynamic-variant, string-names-paths

Code examples: Variant, Array, Dictionary

Evidence:
- Code: `core/variant/variant_parser.h` — VariantParser
- Code: `core/variant/variant_parser.h` — VariantWriter
- Code: `core/variant/variant_parser.cpp` — VariantParser::parse

## Vehicle dynamics

Wheeled, motorcycle, and tracked controllers are separate implementations over the vehicle constraint.

Vehicle dynamics couples a VehicleConstraint with a rigid Body, wheel collision tests, suspension, controllers, engine, transmission, differentials, tracks, and anti-roll bars.

Prerequisites: constraints, narrow-phase, rigid-body-motion

Code examples: Body, VehicleTransmissionSettings

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleConstraint.h` — VehicleConstraint
- Code: `thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleCollisionTester.h` — VehicleCollisionTester, VehicleCollisionTesterRay, VehicleCollisionTesterCastSphere, VehicleCollisionTesterCastCylinder
- Code: `thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleTransmission.cpp` — VehicleTransmissionSettings::SaveBinaryState

## Visual Shader source partition

This inventory establishes the module's source partition without relying on unlisted implementation bodies.

The Visual Shader build script compiles the module's C++ sources, its node sources, and editor sources when building the editor.

Evidence:
- Code: `modules/visual_shader/SCsub` — env_visual_shader.add_source_files

## Visual shader nodes

The module defines input, output, parameter, expression, varying, math, texture, and particle node families.

Node resources provide reusable shader operations within typed graphs of node records and connection records.

Prerequisites: visual-shader-graph

Code examples: VisualShaderNode

Evidence:
- Code: `modules/visual_shader/visual_shader.h` — VisualShaderNode
- Code: `modules/visual_shader/vs_nodes/visual_shader_nodes.h` — VisualShaderNodeVectorOp
- Code: `modules/visual_shader/vs_nodes/visual_shader_particle_nodes.h` — VisualShaderNodeParticleEmitter

## WebSocket peers

The browser build includes a JavaScript socket bridge, while native builds include a WSLay-based peer.

WebSocketPeer defines packet-oriented socket behavior and has native and browser-backed implementations.

Code examples: WebSocketPeer

Evidence:
- Code: `modules/websocket/websocket_peer.h` — WebSocketPeer
- Code: `modules/websocket/wsl_peer.h` — WSLPeer
- Code: `modules/websocket/emws_peer.h` — EMWSPeer
- External (primary, verified): [RFC 6455: The WebSocket Protocol](https://www.rfc-editor.org/rfc/rfc6455), accessed 2026-07-16

## Zstandard stream compression

The implementation is divided into common, compress, and decompress subtrees.

The bundled code provides Zstandard stream compression contexts, decompression contexts, dictionaries, entropy tables, hash-based match finders, and optional worker pools.

Evidence:
- Code: `thirdparty/zstd/zstd.h` — ZSTD_CCtx, ZSTD_DCtx, ZSTD_CDict, ZSTD_DDict
- Code: `thirdparty/zstd/compress/zstd_compress_internal.h` — ZSTD_CCtx_s and ZSTD_MatchState_t
- Code: `thirdparty/zstd/decompress/zstd_decompress_internal.h` — ZSTD_DCtx_s
- Code: `thirdparty/zstd/common/pool.h` — POOL_create and POOL_joinJobs
- External (primary, unverified (source anchor not found)): [RFC 8878: Zstandard Compression and the application/zstd Media Type](https://www.rfc-editor.org/rfc/rfc8878.html), accessed 2026-07-15

## 3D gizmo authoring

The editor distributes type-specific 3D visualization behavior across dedicated gizmo plug-ins.

Node3D editor gizmo plug-ins visualize and manipulate cameras, lights, meshes, occluders, particle emitters, physics objects, skeleton tools, and visibility volumes.

Evidence:
- Code: `editor/scene/3d/gizmos/camera_3d_gizmo_plugin.h` — Camera3DGizmoPlugin
- Code: `editor/scene/3d/gizmos/light_3d_gizmo_plugin.h` — Light3DGizmoPlugin
- Code: `editor/scene/3d/gizmos/mesh_instance_3d_gizmo_plugin.h` — MeshInstance3DGizmoPlugin
- Code: `editor/scene/3d/gizmos/physics/collision_shape_3d_gizmo_plugin.h` — CollisionShape3DGizmoPlugin

## Cryptographic resources and contexts

The crypto area combines persisted key or certificate resources with operational contexts.

CryptoKey and X509Certificate are Resource types, while AES, hashing, HMAC, TLS options, and crypto operations are exposed through dedicated contexts and services.

Code examples: CryptoKey, X509Certificate

Evidence:
- Code: `core/crypto/crypto.h` — CryptoKey
- Code: `core/crypto/crypto.h` — X509Certificate
- Code: `core/crypto/hashing_context.h` — HashingContext

## Editor plugins and customization hooks

Editor-facing APIs are distinct from exported-project APIs and specialize controls, property editing, importing, and export behavior.

EditorPlugin subclasses register inspector, import, export, debugger, and resource-preview hooks in the editor.

Code examples: Resource

Evidence:
- Code: `doc/classes/EditorPlugin.xml` — EditorPlugin
- Code: `doc/classes/EditorInspectorPlugin.xml` — EditorInspectorPlugin
- Code: `doc/classes/EditorImportPlugin.xml` — EditorImportPlugin

## Font streams

This is the byte-access boundary used by the supplied compression and font-reading code.

Font input is represented by an FT_Stream that can deliver bytes through an in-memory base or a read callback.

Code examples: FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c` — ft_bzip2_file_fill_input
- External (official, verified): [BZIP2 Streams - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-bzip2.html), accessed 2026-07-16

## GUI controls and layout

The implementation includes buttons, text and code editing, dialogs, file selection, color picking, graph editing, and flow layout.

Control-derived widgets and Container-derived layouts implement the scene GUI layer.

Prerequisites: scene-graph

Code examples: Control, Container, CodeEdit

Evidence:
- Code: `scene/gui/control.h` — Control
- Code: `scene/gui/container.h` — Container
- Code: `scene/gui/code_edit.h` — CodeEdit

## Input events and actions

The implementation separates concrete device events from named action bindings.

InputEvent resource subclasses represent key, mouse, joypad, touch, gesture, MIDI, shortcut, and action input; InputMap associates named actions with InputEvent lists and deadzones.

Code examples: InputEvent, InputMap::Action

Evidence:
- Code: `core/input/input_event.h` — InputEvent
- Code: `core/input/input_event.h` — InputEventKey
- Code: `core/input/input_map.h` — InputMap::Action

## Localization and translation-template generation

Message context, singular text, plural text, comments, and source locations are explicitly handled during template generation.

Localization tooling parses translation inputs, exposes locale selection and preview, edits localization settings, and generates message maps and POT-style template output.

Code examples: Translation Message

Evidence:
- Code: `editor/translations/editor_translation_parser.h` — EditorTranslationParserPlugin and EditorTranslationParser
- Code: `editor/translations/template_generator.cpp` — TranslationTemplateGenerator message context, plural, comment, location, MessageKey, and MessageMap handling
- Code: `editor/translations/editor_translation_preview_menu.h` — EditorTranslationPreviewMenu : public PopupMenu
- External (official, unverified (source anchor not found)): [Translation — Godot Engine documentation](https://docs.godotengine.org/en/stable/classes/class_translation.html), accessed 2026-07-15

## managed C# script bridge and source generation

The Mono module represents managed scripts through CSharpScript and CSharpInstance, while its SDK tests use partial C# types and source generators for properties, methods, signals, serialization, and script paths.

Code examples: CSharpScript

Evidence:
- Code: `modules/mono/csharp_script.h` — CSharpScript and CSharpInstance
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/ScriptPropertiesGeneratorTests.cs` — ScriptPropertiesGeneratorTests

## Navigation agents and regions

Separate 2D and 3D implementations retain paths, path metadata, links, regions, and obstacles.

Navigation agents consume navigation-result paths, while regions, links, and obstacles describe navigation-world participation.

Prerequisites: scene-graph

Code examples: NavigationAgent2D, NavigationAgent3D, NavigationRegion3D

Evidence:
- Code: `scene/2d/navigation/navigation_agent_2d.h` — NavigationAgent2D
- Code: `scene/3d/navigation/navigation_agent_3d.h` — NavigationAgent3D
- Code: `scene/3d/navigation/navigation_region_3d.h` — NavigationRegion3D

## Networking and transport I/O

Protocol-facing abstractions are organized in core I/O, with Web-specific implementations for browser constraints.

The core exposes HTTP clients, stream peers, packet peers, TCP and UDP servers, DTLS and TLS peers, IP resolution, and Unix-domain socket support.

Code examples: IPAddress

Evidence:
- Code: `core/io/http_client.h` — HTTPClient
- Code: `core/io/packet_peer.h` — PacketPeer
- Code: `core/io/stream_peer.h` — StreamPeer
- Code: `core/io/ip_address.h` — IPAddress

## CPU and GPU particle systems

The module also includes 3D GPU collision and attractor node families.

2D and 3D particle nodes provide separate CPU and GPU particle-system implementations.

Prerequisites: scene-graph

Code examples: CPUParticles2D, GPUParticles3D

Evidence:
- Code: `scene/2d/cpu_particles_2d.h` — CPUParticles2D
- Code: `scene/3d/gpu_particles_3d.h` — GPUParticles3D
- Code: `scene/3d/gpu_particles_collision_3d.h` — GPUParticlesCollision3D

## Scene state inspection

It is an inspection representation rather than a mutable scene graph.

SceneState exposes a packed scene's resources, nodes, exported properties, overrides, and built-in scripts without instantiating the scene.

Prerequisites: resource-formats

Code examples: SceneState, Resource

Evidence:
- Code: `doc/classes/SceneState.xml` — SceneState

## Skeletal modifiers and inverse kinematics

The repository supplies general IK bases plus FABRIK and Jacobian solver variants.

IK skeleton modifiers process bone chains, targets, and joint-limitation resources to modify skeletal poses.

Prerequisites: resources

Code examples: Resource

Evidence:
- Code: `doc/classes/IKModifier3D.xml` — IKModifier3D
- Code: `doc/classes/IterateIK3D.xml` — IterateIK3D
- Code: `doc/classes/JointLimitation3D.xml` — JointLimitation3D

## Skeleton modification and physical-bone simulation

Physical-bone adapters transfer simulated transforms back to visual bones.

2D modifier stacks order Resource-based bone modifications, while SkeletonModifier3D nodes run after AnimationMixer playback and can implement IK, constraints, and skeleton physics.

Code examples: SkeletonModificationStack2D, SkeletonProfile

Evidence:
- Code: `doc/classes/SkeletonModificationStack2D.xml` — SkeletonModificationStack2D
- Code: `doc/classes/SkeletonModifier3D.xml` — SkeletonModifier3D
- Code: `doc/classes/SkeletonModification2DPhysicalBones.xml` — SkeletonModification2DPhysicalBones

## SPIR-V shader reflection

The C header also provides a C++ ShaderModule wrapper when compiled as C++.

The reflection API creates a shader-module metadata view from SPIR-V code and exposes entry points, descriptors, interfaces, push constants, and specialization constants.

Code examples: SpvReflectShaderModule, SpvReflectEntryPoint

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectShaderModule
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectCreateShaderModule, spvReflectDestroyShaderModule
- External (official, unverified (source anchor not found)): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-15

## Byte streams and socket servers

StreamPeerBuffer provides the same stream cursor model over an in-memory byte array.

StreamPeer exposes typed and raw byte-stream I/O, specialized peers implement TCP, TLS, compression, and Unix-domain sockets, and socket servers accept peer connections.

Code examples: StreamPeerBuffer, RID

Evidence:
- Code: `doc/classes/StreamPeer.xml` — StreamPeer
- Code: `doc/classes/StreamPeerBuffer.xml` — StreamPeerBuffer
- Code: `doc/classes/TCPServer.xml` — TCPServer

## Triangle intersection algorithms

Triangle intersection uses Möller–Trumbore, Plücker, and Woop kernels; ray-primitive intersection is required to interpret their hit candidates.

Prerequisites: ray-primitive-intersection

Code examples: TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h` — MoellerTrumboreIntersector1
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_pluecker.h` — PlueckerIntersector1
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_woop.h` — WoopIntersector1

## 2D shapes, tiles, and paths

The 2D resource partition is conditionally compiled for physics and navigation features.

Resource-owned Shape2D subclasses, TileSet data, NavigationPolygon data, and PolygonPathFinder provide 2D geometry and tile-oriented content.

Prerequisites: resource-assets

Code examples: TileMapPattern

Evidence:
- Code: `scene/resources/2d/shape_2d.h` — class Shape2D : public Resource
- Code: `scene/resources/2d/tile_set.h` — class TileMapPattern and class TileSet
- Code: `scene/resources/2d/navigation_polygon.h` — class NavigationPolygon : public Resource
- Code: `scene/resources/2d/polygon_path_finder.h` — class PolygonPathFinder : public Resource

## Vorbis block analysis and synthesis

The implementation has separate analysis and synthesis entry points over the same block-oriented codec model.

Vorbis analysis and synthesis operate on codec blocks and Ogg packets, using mapping, floor, residue, codebook, window, and MDCT modules.

Prerequisites: ogg-pages-and-packets

Code examples: Vorbis Block, Ogg Packet

Evidence:
- Code: `thirdparty/libvorbis/analysis.c` — vorbis_analysis
- Code: `thirdparty/libvorbis/synthesis.c` — vorbis_synthesis
- Code: `thirdparty/libvorbis/mdct.c` — mdct_backward
- External (primary, unverified (source anchor not found)): [Vorbis I Specification](https://xiph.org/vorbis/doc/Vorbis_I_spec.html), accessed 2026-07-15

## X.509 certificate processing

The implementation is split into certificate, request, revocation-list, OID, and writing translation units.

TLS credential processing parses, writes, and verifies certificates, certificate requests, revocation lists, names, and related ASN.1 data.

Prerequisites: tls-connection-state

Evidence:
- Code: `thirdparty/mbedtls/include/mbedtls/x509_crt.h` — mbedtls_x509_crt API
- Code: `thirdparty/mbedtls/library/x509_crt.c` — mbedtls_x509_crt_verify_with_profile
- Code: `thirdparty/mbedtls/library/x509_csr.c` — X.509 CSR parsing implementation
- Code: `thirdparty/mbedtls/library/x509_oid.c` — X.509 OID lookup implementation
- External (official, verified): [Mbed TLS documentation hub](https://mbed-tls.readthedocs.io/en/latest/), accessed 2026-07-16

## Compact heightfield representation

Region algorithms index cells and spans through the heightfield's width and height.

A compact heightfield stores grid dimensions, compact cells, spans, and per-span area data for navigation processing.

Code examples: rcCompactHeightfield

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp` — rcBuildRegions

## Compressed font stream adapters

The Bzip2 adapter checks the header, initializes a decompressor, resets on backward seeks, and installs read and close callbacks.

Optional Bzip2 and LZW components wrap a compatible source stream and expose decompressed bytes through the same FT_Stream callbacks.

Prerequisites: font-streams

Code examples: FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c` — FT_Stream_OpenBzip2
- Code: `thirdparty/freetype/src/lzw/ftlzw.c` — FT_Stream_OpenLZW
- External (official, verified): [BZIP2 Streams - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-bzip2.html), accessed 2026-07-16

## Parametric curve bases

These basis implementations are reused by the subdivision-side curve and patch evaluators.

Bezier, B-spline, and Catmull-Rom basis evaluators generate positions and derivatives from four control points.

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/bezier_curve.h` — BezierBasis
- Code: `thirdparty/embree/kernels/subdiv/bspline_curve.h` — BSplineBasis
- Code: `thirdparty/embree/kernels/subdiv/catmullrom_curve.h` — CatmullRomBasis

## Editor plug-in extension

EditorPlugin is the integration surface used by the editor’s built-in tool plug-ins as well as external extensions.

Plugins attach behavior through polymorphic C++ hook interfaces and may register import, scene-format, post-import, inspector, gizmo, debugger, and resource-conversion plugins.

Evidence:
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin::add_scene_format_importer_plugin
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin::add_inspector_plugin

## GUI controls and layout

The partition includes both leaf controls and container controls.

Control-derived GUI classes implement selection, input, scrolling, split layout, menus, and graph editing within the CanvasItem tree.

Prerequisites: canvas-and-viewports

Code examples: GraphEditConnection

Evidence:
- Code: `scene/gui/scroll_bar.h` — class ScrollBar : public Range
- Code: `scene/gui/split_container.h` — class SplitContainer : public Container
- Code: `scene/gui/tree.h` — class Tree : public Control
- Code: `scene/gui/graph_edit.h` — class GraphEdit : public Control

## Metal-cpp object bridge

The wrapper layer covers Metal, Metal 4, MetalFX, and QuartzCore headers.

The Metal portion models framework objects as C++ wrapper classes whose inline methods send selectors to the underlying object.

Code examples: MTL4::Archive

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTL4Archive.hpp` — MTL4::Archive::newBinaryFunction
- External (official, verified): [Get started with Metal-cpp](https://developer.apple.com/metal/cpp/), accessed 2026-07-16

## Motion-blur geometry

Motion-blur triangle code computes vertices at ray time before triangle intersection, so triangle intersection is required to interpret its candidates.

Prerequisites: triangle-intersection-algorithms

Code examples: TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/trianglev_mb_intersector.h` — TriangleMvMBIntersector1Moeller
- Code: `thirdparty/embree/kernels/geometry/trianglev_mb_intersector.h` — TriangleMvMBIntersectorKPluecker
- Code: `thirdparty/embree/kernels/geometry/subgrid_mb_intersector.h` — SubGridMBIntersectorKPluecker

## Platform display and window adaptation

The same engine-facing display role is implemented separately for X11, macOS, Windows, and Web.

DisplayServer implementations adapt platform windows and events; the Web adapter constructs InputEvent instances and forwards them to Input.

Prerequisites: input-events-actions

Code examples: InputEvent

Evidence:
- Code: `platform/web/display_server_web.cpp` — DisplayServerWeb::_mouse_button_callback
- Code: `platform/linuxbsd/x11/display_server_x11.h` — DisplayServerX11
- Code: `platform/macos/display_server_macos.h` — DisplayServerMacOS
- Code: `platform/windows/display_server_windows.h` — DisplayServerWindows

## Regular-expression JIT code generation

PCRE2 JIT compilation turns compiled pattern control flow into architecture-specific generated code using SLJIT labels and jumps.

Prerequisites: regex-compile-match

Code examples: sljit_compiler, sljit_jump, sljit_label

Evidence:
- Code: `thirdparty/pcre2/src/pcre2_jit_compile.c` — pcre2_jit_compile
- Code: `thirdparty/pcre2/deps/sljit/sljit_src/sljitLir.h` — struct sljit_jump

## Renderer-owned resource storage

Storage isolates resource lifetime and backend-specific data from scene traversal and draw submission.

The RD storage services own backend representations of lights, materials, meshes, particles, textures, and GPU resources addressed by rendering IDs.

Prerequisites: rendering-device-resources

Evidence:
- Code: `servers/rendering/renderer_rd/storage_rd/light_storage.h` — LightStorage
- Code: `servers/rendering/renderer_rd/storage_rd/material_storage.h` — MaterialStorage
- Code: `servers/rendering/renderer_rd/storage_rd/mesh_storage.h` — MeshStorage
- Code: `servers/rendering/renderer_rd/storage_rd/texture_storage.h` — TextureStorage

## Meshes, materials, textures, and instancing

The resource layer holds drawable data while scene nodes select where and how it is rendered.

Mesh resources provide surfaces; geometry nodes instance them, Materials control shading, and MultiMesh batches many instances.

Prerequisites: resources, scene-tree

Code examples: Mesh, Material

Evidence:
- Code: `scene/resources/mesh.h` — Mesh::get_surface_count
- Code: `scene/resources/mesh.h` — Mesh::surface_set_material
- Code: `doc/classes/MultiMesh.xml` — MultiMesh

## SFNT table loading

The component is shared infrastructure for OpenType and TrueType-related font handling.

SFNT services read structured tables from font streams, including cmap, metrics, embedded bitmaps, color layers, SVG documents, and WOFF/WOFF2 data.

Prerequisites: font-streams

Code examples: FT_Stream, PFR_FaceRec

Evidence:
- Code: `thirdparty/freetype/src/sfnt/ttload.h` — tt_face_load_font_dir
- Code: `thirdparty/freetype/src/sfnt/ttcmap.c` — tt_cmap_classes
- Code: `thirdparty/freetype/src/sfnt/ttsvg.h` — tt_face_load_svg_doc
- External (official, verified): [TrueType Tables - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html), accessed 2026-07-16

## Shader interface metadata

The module and entry-point structures use counts plus arrays or pointers to express the reflected relationships.

Reflection exposes descriptor sets, interface variables, push-constant blocks, and specialization constants from one reflected shader module; it depends on SPIR-V shader reflection because it is contained in the reflected shader module.

Prerequisites: spirv-shader-reflection

Code examples: SpvReflectDescriptorSet, SpvReflectDescriptorBinding, SpvReflectInterfaceVariable, SpvReflectBlockVariable, SpvReflectSpecializationConstant

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectDescriptorSet, SpvReflectDescriptorBinding
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectInterfaceVariable, SpvReflectBlockVariable, SpvReflectSpecializationConstant
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectEntryPoint

## Skeletal modifiers and inverse kinematics

The module contains multiple IK algorithms and skeletal runtime modifiers.

SkeletonModifier3D subclasses apply bone constraints, IK solvers, retargeting, spring-bone simulation, and XR tracker data.

Prerequisites: scene-graph

Code examples: SkeletonModifier3D, ChainIK3D, SpringBoneSimulator3D

Evidence:
- Code: `scene/3d/skeleton_modifier_3d.h` — SkeletonModifier3D
- Code: `scene/3d/chain_ik_3d.h` — ChainIK3D
- Code: `scene/3d/spring_bone_simulator_3d.h` — SpringBoneSimulator3D

## 2D skeleton modification stacks

The supplied stack includes CCDIK, FABRIK, LookAt, TwoBoneIK, jiggle, and physical-bone variants.

SkeletonModification2D resources define individual 2D skeleton operations, and SkeletonModificationStack2D holds modifications for a Skeleton2D.

Prerequisites: resource-assets

Evidence:
- Code: `scene/resources/2d/skeleton/skeleton_modification_2d.h` — class SkeletonModification2D : public Resource
- Code: `scene/resources/2d/skeleton/skeleton_modification_stack_2d.h` — class SkeletonModificationStack2D : public Resource
- Code: `scene/resources/2d/skeleton/SCsub` — registered skeleton modification source files

## Subgrid intersection

Subgrid intersection uses grid identifiers to load a quad neighborhood and applies triangle intersection algorithms to its triangles; triangle intersection algorithms are needed to interpret those tests.

Prerequisites: triangle-intersection-algorithms

Code examples: SubGrid, GridMesh

Evidence:
- Code: `thirdparty/embree/kernels/geometry/subgrid.h` — SubGrid
- Code: `thirdparty/embree/kernels/geometry/subgrid_intersector.h` — SubGridIntersector1Moeller
- Code: `thirdparty/embree/kernels/geometry/subgrid_intersector.h` — SubGridIntersector1Pluecker

## Tile map layer data

A TileMap node coordinates TileMapLayer instances.

TileMapLayer stores cell data keyed by grid coordinates and derives rendering, physics, navigation, and debug quadrants from that data.

Prerequisites: scene-graph

Code examples: TileMapLayer, TileMapLayerCellData

Evidence:
- Code: `scene/2d/tile_map_layer.h` — TileMapLayer
- Code: `scene/2d/tile_map_layer.h` — CellData
- Code: `scene/2d/tile_map.h` — TileMap

## UI themes and style boxes

Panel, PanelContainer, PopupPanel, separators, and controls consume these style definitions.

Theme resources apply reusable colors, constants, fonts, icons, and StyleBoxes across Control and Window branches, while StyleBox defines a visual box treatment.

Code examples: Theme, StyleBox

Evidence:
- Code: `doc/classes/Theme.xml` — Theme
- Code: `doc/classes/StyleBox.xml` — StyleBox
- Code: `doc/classes/PanelContainer.xml` — PanelContainer

## Version-control integration

The interface separates VCS-provided data structures from the editor-facing plugin.

Version-control integration exchanges structured status, commit, file-diff, hunk, and line data through EditorVCSInterface and presents it through VersionControlEditorPlugin.

Code examples: VCS Diff File, VCS Diff Hunk, VCS Diff Line

Evidence:
- Code: `editor/version_control/editor_vcs_interface.h` — EditorVCSInterface::DiffLine, DiffHunk, DiffFile, Commit, and StatusFile
- Code: `editor/version_control/version_control_editor_plugin.h` — VersionControlEditorPlugin : public EditorPlugin
- External (official, unverified (source anchor not found)): [Version control systems — Godot Engine documentation](https://docs.godotengine.org/en/4.4/tutorials/best_practices/version_control_systems.html), accessed 2026-07-15

## WebRTC data channels

The module supplies a base channel plus extension and JavaScript implementations.

A WebRTC data channel carries PacketPeer data through a WebRTC peer connection.

Prerequisites: webrtc-peer-connections

Code examples: WebRTCDataChannel

Evidence:
- Code: `modules/webrtc/webrtc_data_channel.h` — WebRTCDataChannel
- Code: `modules/webrtc/webrtc_data_channel_js.h` — WebRTCDataChannelJS
- External (primary, verified): [WebRTC: Real-Time Communication in Browsers](https://www.w3.org/TR/webrtc/), accessed 2026-07-16

## WebSocket framing and event contexts

The implementation is organized around incremental frame reads and writes plus event callbacks.

Wslay separates frame-level I/O from event-level message handling through frame contexts, event contexts, callback tables, message sources, and send/control queues.

Code examples: wslay_event_context

Evidence:
- Code: `thirdparty/wslay/wslay/wslay.h` — wslay_frame_context, wslay_event_context, wslay_event_callbacks
- Code: `thirdparty/wslay/wslay_event.h` — struct wslay_event_context
- Code: `thirdparty/wslay/wslay_frame.c` — wslay_frame_recv and wslay_frame_send
- External (primary, unverified (source anchor not found)): [RFC 6455: The WebSocket Protocol](https://www.rfc-editor.org/rfc/rfc6455.html), accessed 2026-07-15

## COLLADA scene interchange

COLLADA is represented as an intermediate parser state before editor scene construction.

The COLLADA parser supplies a scene importer with image, material, effect, mesh, skin, morph, node, visual-scene, and animation data collected in Collada::State.

Prerequisites: scene-importing

Code examples: ColladaParseState

Evidence:
- Code: `editor/import/3d/collada.h` — Collada::State
- Code: `editor/import/3d/editor_import_collada.cpp` — ColladaImport

## Editor plugin lifecycle

The implementation favors many focused plugin subclasses over one monolithic editor tool.

C++ editor-plugin specializations package feature-specific editor behavior behind EditorPlugin-derived classes, including scene, script, and shader tools.

Evidence:
- Code: `editor/scene/3d/mesh_library_editor_plugin.h` — MeshLibraryEditorPlugin : public EditorPlugin
- Code: `editor/shader/shader_editor_plugin.h` — ShaderEditorPlugin : public EditorPlugin
- External (official, unverified (source anchor not found)): [EditorPlugin — Godot Engine documentation](https://docs.godotengine.org/en/latest/classes/class_editorplugin.html), accessed 2026-07-15

## OpenType and TrueTypeGX validation

The OpenType module covers BASE, GDEF, GPOS, GSUB, JSTF, and MATH; the GX module contains validators for multiple AAT table families.

OpenType and TrueTypeGX validators check structured tables after SFNT parsing so higher layers can consume checked offsets and indices.

Prerequisites: sfnt-tables

Evidence:
- Code: `thirdparty/freetype/src/otvalid/otvalid.c` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/gxvalid/gxvalid.c` — FT_MAKE_OPTION_SINGLE_OBJECT
- External (official, unverified (source anchor not found)): [OpenType Validation - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-ot_validation.html), accessed 2026-07-15
- External (official, verified): [TrueTypeGX/AAT Validation - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-gx_validation.html), accessed 2026-07-16

## glTF material and texture conversion

GLTFDocument maps glTF material dictionaries, PBR metallic-roughness values, textures, samplers, UV coordinates, color transforms, alpha mode, and texture-transform extensions to Godot material resources.

Prerequisites: gltf-asset-model

Code examples: GLTFState

Evidence:
- Code: `modules/gltf/gltf_document.cpp` — material parsing and KHR_texture_transform handling
- Code: `modules/gltf/structures/gltf_texture.h` — GLTFTexture

## GPU resources and pipelines

Descriptors configure GPU resources and pipeline state before a device produces allocations, functions, or pipeline state objects.

Prerequisites: metal-cpp-object-bridge

Code examples: MTL::Buffer, MTL::Texture, MTL4::PipelineDescriptor

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTLDevice.hpp` — MTL::Device
- Code: `thirdparty/metal-cpp/Metal/MTL4PipelineState.hpp` — MTL4::PipelineDescriptor

## Halfedge topology

Topology is separate from vertex positions and property vertices.

The Manifold mesh kernel represents every triangle boundary with directed Halfedge records whose pairedHalfedge links encode adjacency.

Prerequisites: manifold-mesh-data

Code examples: Halfedge, MeshGLP

Evidence:
- Code: `thirdparty/manifold/src/shared.h` — struct Halfedge
- Code: `thirdparty/manifold/src/impl.cpp` — PrepHalfedges; Manifold::Impl halfedge_ construction

## Managed-native interop

The bridge has a size-checked unmanaged callback table and native runtime interop table with matching ordering requirements.

Unsafe C# connects managed values and callbacks across a fixed native/managed ABI using variant conversion, GC handles, and function-pointer callbacks.

Code examples: ManagedCallbacks, Variant

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/NativeFuncs.cs` — NativeFuncs.Initialize
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/Marshaling.cs` — Marshaling
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs` — ManagedCallbacks
- Code: `modules/mono/glue/runtime_interop.cpp` — godotsharp::get_runtime_interop_funcs

## Navigation mesh construction

2D construction uses polygon geometry and 3D construction uses Recast-backed mesh generation when configured.

The 2D and 3D generators turn source geometry into navigation mesh data and connect region or link geometry into map iterations.

Evidence:
- Code: `modules/navigation_2d/2d/nav_mesh_generator_2d.h` — NavMeshGenerator2D
- Code: `modules/navigation_2d/2d/nav_region_builder_2d.h` — NavRegionBuilder2D
- Code: `modules/navigation_3d/3d/nav_mesh_generator_3d.h` — NavMeshGenerator3D
- Code: `modules/navigation_3d/3d/nav_region_builder_3d.h` — NavRegionBuilder3D

## Navigation region construction

The implementation contains watershed, monotone, and layer-oriented region-building paths.

Region construction labels connected compact heightfield spans into navigation regions.

Prerequisites: compact-heightfield

Code examples: rcCompactHeightfield

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp` — rcBuildRegions
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp` — rcBuildRegionsMonotone
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp` — rcBuildLayerRegions

## Navigation heightfields

Recast rasterizes input triangles into a heightfield and converts its spans into a compact heightfield with adjacency.

Code examples: rcHeightfield, rcCompactHeightfield, rcCompactSpan

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Include/Recast.h` — rcHeightfield
- Code: `thirdparty/recastnavigation/Recast/Source/Recast.cpp` — rcBuildCompactHeightfield
- External (primary, verified): [recastnavigation/recastnavigation](https://github.com/recastnavigation/recastnavigation), accessed 2026-07-16

## OpenXR extension wrappers

The supplied extensions cover platform graphics bindings, controller profiles, composition layers, hand tracking, render models, futures, and spatial capabilities.

Extension wrappers isolate optional OpenXR runtime features behind a common base interface.

Prerequisites: openxr-runtime-integration

Code examples: OpenXRFutureResult, OpenXRRenderModelData

Evidence:
- Code: `modules/openxr/extensions/openxr_extension_wrapper.h` — class OpenXRExtensionWrapper : public Object
- Code: `modules/openxr/openxr_api.cpp` — OpenXRAPI::get_registered_extension_wrappers
- Code: `modules/openxr/extensions/SCsub` — extension and platform source selection

## Display, camera, video, and movie services

A headless DisplayServer implementation is also present.

DisplayServer abstracts display functionality, CameraServer manages CameraFeed objects, VideoStreamPlayer presents video in a Control, and MovieWriter emits movie output.

Prerequisites: canvas-and-viewports

Evidence:
- Code: `servers/display/display_server.h` — class DisplayServer
- Code: `servers/display/display_server_headless.h` — class DisplayServerHeadless : public DisplayServer
- Code: `servers/camera/camera_server.h` — class CameraServer
- Code: `scene/gui/video_stream_player.h` — class VideoStreamPlayer : public Control
- Code: `servers/movie_writer/movie_writer.h` — class MovieWriter

## Pluggable and wrapped server backends

The same server-facing abstractions support fallback, extension-provided, and wrapped implementations.

C++ inheritance supplies extension, dummy, and multithread-wrapped implementations behind the physics, text, XR, and rendering server interfaces.

Evidence:
- Code: `servers/physics_3d/physics_server_3d_extension.h` — PhysicsServer3DExtension
- Code: `servers/physics_3d/physics_server_3d_dummy.h` — PhysicsServer3DDummy
- Code: `servers/text/text_server_extension.h` — TextServerExtension
- Code: `servers/xr/xr_interface_extension.h` — XRInterfaceExtension

## Post-import processing

The implementation separates parsing a source format from editor-controlled transformations of its imported result.

Post-import processing operates on the imported scene after format conversion and accepts options for node, mesh, material, animation, and skeleton categories.

Prerequisites: scene-importing

Evidence:
- Code: `editor/import/3d/resource_importer_scene.h` — EditorScenePostImport
- Code: `editor/import/3d/resource_importer_scene.h` — EditorScenePostImportPlugin::InternalImportCategory
- Code: `editor/import/3d/post_import_plugin_skeleton_rest_fixer.h` — PostImportPluginSkeletonRestFixer

## Resource-specific editing

The implementation chooses type-specific editing experiences for these resource families.

Dedicated editor controls and docks edit gradients, curves, materials, sprite frames, mesh libraries, textures, packed scenes, and resource preloaders.

Evidence:
- Code: `editor/scene/curve_editor_plugin.h` — CurveEdit and CurveEditor
- Code: `editor/scene/sprite_frames_editor_plugin.h` — SpriteFramesEditor : public EditorDock
- Code: `editor/scene/texture/texture_editor_plugin.h` — TexturePreview and TextureEditorPlugin
- Code: `editor/scene/resource_preloader_editor_plugin.h` — ResourcePreloaderEditor : public EditorDock

## Shader reflection

It is implemented independently from SPIRV-Cross under a C API.

The reflection API reads SPIR-V IR and returns module, entry-point, descriptor-binding, interface-variable, and push-constant metadata.

Prerequisites: spirv-intermediate-representation

Code examples: SpvReflectShaderModule, SpvReflectDescriptorBinding

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SpvReflectShaderModule
- Code: `thirdparty/spirv-reflect/spirv_reflect.c` — spvReflectGetDescriptorBinding
- Code: `thirdparty/spirv-reflect/spirv_reflect.c` — spvReflectGetEntryPointPushConstantBlock
- External (official, unverified (source anchor not found)): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-15

## Spatial indexing and ray queries

Spatial trees use geometry bounds to maintain items, cull them, refit nodes, and accelerate triangle-mesh and static-ray queries.

Prerequisites: geometry-transforms

Code examples: TriangleMesh, AABB

Evidence:
- Code: `core/math/bvh_tree.h` — BVH_Tree
- Code: `core/math/dynamic_bvh.h` — DynamicBVH
- Code: `core/math/triangle_mesh.h` — TriangleMesh
- Code: `core/math/static_raycaster.h` — StaticRaycaster

## Subdivision surface evaluation

Subdivision evaluation combines curve and patch bases with grid sampling; curve and patch bases are needed to explain the evaluator.

Prerequisites: curve-and-patch-bases

Code examples: SubGrid, GridMesh

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/patch.h` — EvalPatch
- Code: `thirdparty/embree/kernels/subdiv/patch_eval.h` — PatchEval
- Code: `thirdparty/embree/kernels/subdiv/patch_eval_grid.h` — PatchEvalGrid

## Text layout and editing

TextEdit, LineEdit, Label, and RichTextLabel each consume shaped-text resources in this partition.

GUI text controls hold shaped line and paragraph data, use text-server glyph and selection queries, and track IME text and selection state.

Prerequisites: gui-controls

Evidence:
- Code: `scene/gui/text_edit.h` — TextEdit::Text and TextEdit::Line
- Code: `scene/gui/text_edit.cpp` — TextEdit IME and shaped_text_get_selection calls
- Code: `scene/gui/rich_text_label.h` — class RichTextLabel and RichTextLabel::Line
- Code: `scene/resources/text_paragraph.h` — class TextParagraph : public RefCounted

## Themes and style boxes

Default-theme construction also appears in the inspected theme source.

Resource-backed Theme data is resolved through ThemeDB, ThemeContext, and ThemeOwner, while StyleBox subclasses supply control styling.

Prerequisites: resource-assets

Evidence:
- Code: `scene/resources/theme.h` — class Theme : public Resource
- Code: `scene/resources/style_box.h` — class StyleBox : public Resource
- Code: `scene/theme/theme_db.h` — class ThemeDB and class ThemeContext
- Code: `scene/theme/default_theme.cpp` — default control colors and StyleBoxFlat construction

## 3D shapes, worlds, and skins

The 3D resource partition is selected only when 3D is enabled by the build configuration.

Resource-owned Shape3D subclasses, World3D, Skin, and mesh import data represent 3D collision, world, skeletal, and mesh content.

Prerequisites: resource-assets

Evidence:
- Code: `scene/resources/3d/shape_3d.h` — class Shape3D : public Resource
- Code: `scene/resources/3d/world_3d.h` — class World3D : public Resource
- Code: `scene/resources/3d/skin.h` — class Skin : public Resource and Skin::Bind
- Code: `scene/resources/3d/importer_mesh.h` — class ImporterMesh : public Resource

## WebRTC multiplayer mesh

ConnectedPeer records support the multiplayer abstraction's peer tracking.

A WebRTCMultiplayerPeer keeps a mesh of WebRTC peer connections and their data channels.

Prerequisites: webrtc-data-channels, webrtc-peer-connections

Code examples: WebRTCMultiplayerPeer

Evidence:
- Code: `modules/webrtc/webrtc_multiplayer_peer.h` — WebRTCMultiplayerPeer
- Code: `modules/webrtc/webrtc_multiplayer_peer.h` — ConnectedPeer

## WebSocket multiplayer

Its packet and pending-peer records are the cross-cutting state used by the multiplayer transport.

WebSocketMultiplayerPeer tracks WebSocket peers plus pending peers and packets.

Prerequisites: websocket-peers

Code examples: WebSocketMultiplayerPeer, WebSocketMultiplayerPacket

Evidence:
- Code: `modules/websocket/websocket_multiplayer_peer.h` — WebSocketMultiplayerPeer
- Code: `modules/websocket/websocket_multiplayer_peer.h` — WebSocketMultiplayerPeer::Packet
- Code: `modules/websocket/websocket_multiplayer_peer.h` — WebSocketMultiplayerPeer::PendingPeer

## WebXR session bridge

The JavaScript bridge manages sessions, reference spaces, frame data, views, and input-source data for the C++ implementation.

WebXR interface state receives browser session callbacks and input sources through web-platform bindings.

Code examples: WebXRInterfaceJS, WebXRInputSource

Evidence:
- Code: `modules/webxr/webxr_interface.h` — WebXRInterface
- Code: `modules/webxr/webxr_interface_js.h` — WebXRInterfaceJS::InputSource
- Code: `modules/webxr/native/library_godot_webxr.js` — GodotWebXR
- External (primary, verified): [WebXR Device API](https://www.w3.org/TR/webxr/), accessed 2026-07-16

## zlib stream compression

Checksum, compression, and decompression code are separate source modules behind zlib.h.

The bundled zlib sources implement Adler-32 and CRC-32 checksums, DEFLATE compression trees, and inflate state for a public streaming compression interface.

Code examples: z_stream

Evidence:
- Code: `thirdparty/zlib/zlib.h` — z_stream and struct internal_state
- Code: `thirdparty/zlib/adler32.c` — adler32
- Code: `thirdparty/zlib/deflate.h` — dyn_ltree, dyn_dtree, bl_tree
- Code: `thirdparty/zlib/inflate.h` — struct inflate_state
- External (primary, unverified (source anchor not found)): [RFC 1950: ZLIB Compressed Data Format Specification](https://www.rfc-editor.org/rfc/rfc1950), accessed 2026-07-15

## Allocation and reference ownership

Builders receive allocator callbacks, while shared runtime objects inherit or use RefCount-based ownership.

Reference-counted objects, FastAllocator blocks, cached allocation, aligned allocation, and monitored allocation support resource and BVH memory management.

Code examples: RefCount, Ref, FastAllocator, CachedAllocator

Evidence:
- Code: `thirdparty/embree/common/sys/ref.h` — RefCount
- Code: `thirdparty/embree/common/sys/ref.h` — Ref
- Code: `thirdparty/embree/kernels/common/alloc.h` — FastAllocator
- Code: `thirdparty/embree/kernels/common/alloc.h` — CachedAllocator

## Android export pipeline

`EditorExportPlatformAndroid` is the editor-facing exporter and its helpers hold export-specific data structures.

Android export derives Gradle project files, manifest content, localized project names, icons, ABIs, and plugin configuration from project and export inputs.

Code examples: APKExportData, PluginConfigAndroid, LauncherIcon, CustomExportData

Evidence:
- Code: `platform/android/export/export_plugin.h` — EditorExportPlatformAndroid
- Code: `platform/android/export/export_plugin.h` — APKExportData
- Code: `platform/android/export/godot_plugin_config.h` — PluginConfigAndroid
- Code: `platform/android/export/gradle_export_util.h` — CustomExportData

## Basis texture transcoding

The implementation exposes a container-level transcoder plus specialized ETC1S, UASTC, and ASTC paths.

The Basis Universal subsystem decodes ETC1S and UASTC texture slices into selected texture formats through high-level and low-level transcoders.

Code examples: BasisFileHeader, BasisSliceDescriptor

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — basisu_transcoder
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — basisu_lowlevel_etc1s_transcoder

## BVH split heuristics

The supplied builders calculate costs from bounds, primitive counts, block size, traversal cost, and intersection cost.

SAH, Morton, spatial, temporal, strand, open-merge, and motion-blur heuristics choose how PrimRef records are divided while constructing a BVH.

Prerequisites: bvh-construction, primitive-references

Code examples: BinMapping, MortonCodeMapping, SpatialBinMapping, HeuristicMBlurTemporalSplit

Evidence:
- Code: `thirdparty/embree/kernels/builders/heuristic_binning.h` — BinMapping
- Code: `thirdparty/embree/kernels/builders/bvh_builder_morton.h` — MortonCodeMapping
- Code: `thirdparty/embree/kernels/builders/heuristic_spatial.h` — SpatialBinMapping
- Code: `thirdparty/embree/kernels/builders/heuristic_timesplit_array.h` — HeuristicMBlurTemporalSplit

## CFF font subsetting

CFF-specific code is distinct from generic table dispatch because it carries charstring and subroutine structures.

The font subsetting pipeline contains CFF1 and CFF2 plans, accelerators, serializers, and subroutine-subsetting helpers.

Prerequisites: font-subsetting

Code examples: hb_subset_plan_t, hb_subset_accelerator_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-cff-common.hh` — cff_subset_accelerator_t
- Code: `thirdparty/harfbuzz/src/hb-subset-cff2.cc` — cff2_subset_plan

## Editor scene sessions

EditedScene is the durable editor-side record for an open scene rather than the scene tree itself.

EditorData represents each open editor scene with its root, project path, plugin state, selection, undo history ID, and timing/version metadata.

Code examples: EditedScene

Evidence:
- Code: `editor/editor_data.h` — EditorData::EditedScene
- Code: `editor/editor_data.cpp` — EditorData::load_editor_plugin_states_from_config

## Feature-adaptive tessellation

Feature-adaptive tessellation recursively partitions parameter ranges for patch evaluation; subdivision evaluation is needed to explain the patch evaluation.

Prerequisites: subdivision-surface-evaluation

Code examples: SubGrid

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/feature_adaptive_eval.h` — FeatureAdaptiveEval
- Code: `thirdparty/embree/kernels/subdiv/feature_adaptive_eval_grid.h` — FeatureAdaptiveEvalGrid
- Code: `thirdparty/embree/kernels/subdiv/tessellation.h` — tessellation

## GPU command encoding

Command buffers create render, compute, blit, and resource-state encoders that record work against configured GPU resources.

Prerequisites: gpu-resources-pipelines

Code examples: MTL::CommandBuffer, MTL::RenderCommandEncoder, MTL::ComputeCommandEncoder

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTLCommandBuffer.hpp` — MTL::CommandBuffer
- Code: `thirdparty/metal-cpp/Metal/MTLComputeCommandEncoder.hpp` — MTL::ComputeCommandEncoder

## Half-edge topology

Subdivision connectivity is represented by HalfEdge navigation and by Catmull–Clark rings that inspect neighboring edges, faces, and crease information.

Code examples: HalfEdge

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/half_edge.h` — HalfEdge
- Code: `thirdparty/embree/kernels/subdiv/catmullclark_ring.h` — CatmullClark1RingT

## Custom inspector property editors

This path is used for gradients, materials, textures, particle-process ranges, GUI controls, fonts, and style boxes.

The implementation uses editor-plugin lifecycle integration to replace or extend property editing with specialized EditorInspectorPlugin and EditorProperty classes.

Prerequisites: editor-plugin-lifecycle

Evidence:
- Code: `editor/scene/gradient_editor_plugin.h` — EditorInspectorPluginGradient : public EditorInspectorPlugin
- Code: `editor/scene/particle_process_material_editor_plugin.h` — ParticleProcessMaterialMinMaxPropertyEditor : public EditorProperty
- Code: `editor/scene/gui/font_config_plugin.h` — EditorInspectorPluginFontVariation and EditorPropertyOTVariation
- External (official, unverified (source anchor not found)): [Inspector plugins — Godot Engine documentation](https://docs.godotengine.org/en/stable/tutorials/plugins/editor/inspector_plugins.html), accessed 2026-07-15

## MetalFX scaling and interpolation

MetalFX descriptors create spatial, temporal, temporal-denoised, and frame-interpolation effect instances from a Metal device.

Prerequisites: gpu-resources-pipelines

Code examples: MTLFX::SpatialScaler, MTLFX::TemporalScaler, MTLFX::FrameInterpolator

Evidence:
- Code: `thirdparty/metal-cpp/MetalFX/MTLFXSpatialScaler.hpp` — MTLFX::SpatialScalerDescriptor::newSpatialScaler
- Code: `thirdparty/metal-cpp/MetalFX/MTLFXFrameInterpolator.hpp` — MTLFX::FrameInterpolatorDescriptor::newFrameInterpolator

## Navigation maps and path queries

Both 2D and 3D APIs expose equivalent navigation query objects and server services.

Navigation agents use server maps, regions, path-query parameters, and path-query results to follow a target position.

Code examples: NavigationPathQueryParameters2D, NavigationPathQueryResult2D

Evidence:
- Code: `servers/navigation_server_2d.h` — NavigationServer2D
- Code: `doc/classes/NavigationPathQueryParameters2D.xml` — NavigationPathQueryParameters2D
- Code: `doc/classes/NavigationPathQueryResult2D.xml` — NavigationPathQueryResult2D

## Navigation map composition

Parallel 2D and 3D implementations expose map-owned region, link, agent, and obstacle collections and iteration snapshots.

Navigation maps collect regions, links, agents, and obstacles and build per-frame read iterations.

Evidence:
- Code: `modules/navigation_2d/nav_map_2d.h` — NavMap2D
- Code: `modules/navigation_2d/2d/nav_map_iteration_2d.h` — NavMapIteration2D
- Code: `modules/navigation_3d/nav_map_3d.h` — NavMap3D
- Code: `modules/navigation_3d/3d/nav_map_iteration_3d.h` — NavMapIteration3D
- External (official, unverified (source anchor not found)): [2D navigation overview — Godot Engine](https://docs.godotengine.org/en/latest/tutorials/navigation/navigation_introduction_2d.html), accessed 2026-07-15

## Navigation contours and polygons

Recast extracts contours from a compact heightfield, then builds polygon and detail meshes from those contours.

Prerequisites: navmesh-heightfields

Code examples: rcContourSet, rcPolyMesh, rcPolyMeshDetail

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Source/RecastContour.cpp` — rcBuildContours
- Code: `thirdparty/recastnavigation/Recast/Source/RecastMesh.cpp` — rcBuildPolyMesh

## OpenEXR image decoding

The supplied partition demonstrates integration of a single-header image decoder rather than a ThorVG loader implementation.

TinyEXR is compiled as a header implementation with zlib included first, and its public header defines EXR-oriented image API data and functions.

Code examples: EXRImage

Evidence:
- Code: `thirdparty/tinyexr/tinyexr.cc` — TINYEXR_IMPLEMENTATION
- Code: `thirdparty/tinyexr/tinyexr.h` — EXRImage, EXRHeader

## native 2D collision pipeline

Physics spaces hold collision objects; the 2D BVH broad phase finds candidate collision objects and the solver dispatches shape-pair routines, including separating-axis tests.

Code examples: GodotSpace2D, GodotCollisionObject2D

Evidence:
- Code: `modules/godot_physics_2d/godot_broad_phase_2d_bvh.h` — GodotBroadPhase2DBVH
- Code: `modules/godot_physics_2d/godot_collision_solver_2d_sat.cpp` — SeparatorAxisTest2D

## Platform-selective shader baking

The platform split occurs in build manifests and is reflected in three platform adapter classes.

Platform-selective shader baking depends on shader editing and compiles distinct Vulkan, D3D12, or Metal export-plugin sources when matching SCons environment flags are enabled.

Prerequisites: shader-editing-and-language-plugins

Evidence:
- Code: `editor/shader/shader_baker/SCsub` — vulkan, d3d12, and metal conditional source selection
- Code: `editor/shader/shader_baker/shader_baker_export_plugin_platform_d3d12.h` — ShaderBakerExportPluginPlatformD3D12
- Code: `editor/shader/shader_baker/shader_baker_export_plugin_platform_metal.h` — ShaderBakerExportPluginPlatformMetal
- Code: `editor/shader/shader_baker/shader_baker_export_plugin_platform_vulkan.h` — ShaderBakerExportPluginPlatformVulkan

## Raster image input

The Basis encoder dependency includes separate PNG and JPEG decoding implementations.

Raster image input adapters decode PNG and JPEG bytes into image buffers for texture-processing callers.

Code examples: PngInfo

Evidence:
- Code: `thirdparty/basis_universal/encoder/pvpngreader.h` — pv_png::get_png_info
- Code: `thirdparty/basis_universal/encoder/pvpngreader.h` — pv_png::load_png
- Code: `thirdparty/basis_universal/encoder/jpgd.h` — jpeg_decoder

## Textures, meshes, and materials

The inspected source includes array meshes, primitive meshes, compressed and image textures, and several material families.

Resource-owned Mesh, Material, Texture, Environment, and Sky classes carry rendering-facing content and configuration.

Prerequisites: resource-assets

Evidence:
- Code: `scene/resources/mesh.h` — class Mesh and class ArrayMesh
- Code: `scene/resources/material.h` — class Material, ShaderMaterial, and BaseMaterial3D
- Code: `scene/resources/texture.h` — Texture, Texture2D, TextureLayered, and Texture3D
- Code: `scene/resources/environment.h` — class Environment : public Resource

## Standalone resource importing

This is separate from scene import and produces editor resource assets from individual files.

Resource-importer classes handle images, textures, texture atlases, SVG, fonts, WAV audio, CSV translations, shaders, and bitmaps.

Evidence:
- Code: `editor/import/resource_importer_image.h` — ResourceImporterImage
- Code: `editor/import/resource_importer_texture.h` — ResourceImporterTexture
- Code: `editor/import/resource_importer_texture_atlas.h` — ResourceImporterTextureAtlas
- Code: `editor/import/resource_importer_svg.h` — ResourceImporterSVG
- Code: `editor/import/resource_importer_dynamic_font.h` — ResourceImporterDynamicFont
- Code: `editor/import/resource_importer_wav.h` — ResourceImporterWAV
- Code: `editor/import/resource_importer_csv_translation.h` — ResourceImporterCSVTranslation
- Code: `editor/import/resource_importer_shader_file.h` — ResourceImporterShaderFile
- Code: `editor/import/resource_importer_bitmask.h` — ResourceImporterBitMap

## Texture resources and fallback placeholders

Layered and 3D textures require consistently sized image data across their layers.

Texture resources represent 2D, 3D, layered, and RenderingDevice-backed image data; placeholder texture classes preserve limited shape information when source texture data is unavailable or omitted.

Prerequisites: resource-formats

Code examples: RDTextureFormat, RID

Evidence:
- Code: `doc/classes/Texture2D.xml` — Texture2D
- Code: `doc/classes/TextureLayered.xml` — TextureLayered
- Code: `doc/classes/Texture2DRD.xml` — Texture2DRD
- Code: `doc/classes/PlaceholderTexture2D.xml` — PlaceholderTexture2D

## Variable-font table subsetting

Variation handling is controlled by table tags and user-axis locations in the subset plan.

The font subsetting pipeline dispatches HVAR, VVAR, gvar, fvar, avar, cvar, and mvar tables, with fvar and avar passed through when user axis locations are empty.

Prerequisites: font-subsetting

Code examples: hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-table-var.cc` — _hb_subset_table_var
- Code: `thirdparty/harfbuzz/src/hb-subset-plan-var.cc` — hb_subset_plan_var

## Android Gradle assembly

The Android build root explicitly includes all five module names.

Gradle settings and build logic assemble app, library, editor, native-source-index, and install-time asset-pack modules for the Android export pipeline.

Prerequisites: android-export-pipeline

Evidence:
- Code: `platform/android/java/settings.gradle` — include ':app'
- Code: `platform/android/java/settings.gradle` — include ':lib'
- Code: `platform/android/java/build.gradle` — generateBuildTasks
- External (official, verified): [Gradle DSL Reference](https://docs.gradle.org/current/dsl/), accessed 2026-07-16

## APK signing and verification

The Android editor has a Kotlin utility that imports `ApkSigner` and `ApkVerifier`.

APK signing and verification operate on Android export pipeline artifacts through the bundled apksig implementation.

Prerequisites: android-export-pipeline

Evidence:
- Code: `platform/android/java/editor/src/main/java/org/godotengine/editor/utils/ApkSignerUtil.kt` — signApk
- Code: `platform/android/java/editor/src/main/java/com/android/apksig/ApkSigner.java` — ApkSigner
- Code: `platform/android/java/editor/src/main/java/com/android/apksig/ApkVerifier.java` — ApkVerifier

## Basis file layout

The transcoder reads a header followed by an array of slice descriptions addressed through header offsets.

A .basis file header and slice descriptors locate compressed slices and identify their texture form before Basis Universal decodes them.

Prerequisites: basis-transcoding

Code examples: BasisFileHeader, BasisSliceDescriptor

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_file_headers.h` — basis_file_header
- Code: `thirdparty/basis_universal/transcoder/basisu_file_headers.h` — basis_slice_desc
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_transcoder::get_file_info

## GPU block texture conversion

The transcoders support a range of destination block formats and uncompressed pixel formats.

Texture conversion code maps ETC1S, UASTC, and ASTC blocks to GPU block or uncompressed target formats, calculating output block counts and storage.

Prerequisites: basis-transcoding

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_lowlevel_etc1s_transcoder::transcode_slice
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_lowlevel_uastc_ldr_4x4_transcoder::transcode_slice

## Geometry buffer storage

The public API also defines opaque RTCBuffer handles and buffer-type slots for indices, vertices, attributes, normals, transforms, and related data.

A Buffer is a reference-counted storage abstraction, while BufferView and RawBufferView expose views used to bind geometry data.

Code examples: RTCBuffer, Buffer

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h` — RTCBufferType
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h` — RTCBuffer
- Code: `thirdparty/embree/kernels/common/buffer.h` — BufferView

## Catmull–Clark patch construction

Catmull–Clark patch construction reads a half-edge topology and converts its limit data to patch geometry; half-edge topology is needed to explain this construction.

Prerequisites: half-edge-topology

Code examples: HalfEdge

Evidence:
- Code: `thirdparty/embree/kernels/subdiv/catmullclark_patch.h` — CatmullClarkPatchT
- Code: `thirdparty/embree/kernels/subdiv/catmullclark_patch.h` — GeneralCatmullClarkPatchT
- Code: `thirdparty/embree/kernels/subdiv/gregory_patch.h` — GregoryPatchT

## Editing selection history

The history stores object paths and the currently selected point in that path history.

Selection history gives an editor scene session back-and-forward navigation across selected objects and nested-resource properties.

Prerequisites: editor-scene-sessions

Code examples: EditedScene

Evidence:
- Code: `editor/editor_data.h` — EditorSelectionHistory
- Code: `editor/editor_data.cpp` — EditorSelectionHistory::add_object

## Editor plugin state and custom types

EditorData selects a main editor plugin for an object and retains plugin-provided type metadata.

Editor plugins can handle objects, register custom types, and serialize plugin state into the active scene's editor state.

Prerequisites: editor-scene-sessions

Code examples: EditedScene

Evidence:
- Code: `editor/editor_data.h` — EditorData::CustomType
- Code: `editor/editor_data.cpp` — EditorData::get_handling_main_editor
- Code: `editor/editor_data.cpp` — EditorData::get_scene_editor_states_with_selection

## Font driver modules

Each format family has a driver class and generally separates driver, object, loader, and glyph-loading concerns.

Format-specific driver classes expose CFF, CID, PCF, PFR, Type 1, Type 42, Windows FNT, and TrueType implementations to the FreeType module layer.

Code examples: PFR_FaceRec

Evidence:
- Code: `thirdparty/freetype/src/cff/cffdrivr.h` — cff_driver_class
- Code: `thirdparty/freetype/src/truetype/ttdriver.h` — tt_driver_class
- Code: `thirdparty/freetype/src/winfonts/winfnt.c` — winfnt_driver_class
- External (official, verified): [Module Management - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-module_management.html), accessed 2026-07-16

## Godot-specific third-party adaptation

These files document local deltas from upstream vendored code.

The repository applies VMA allocator integration patches, redirects Vulkan includes to Godot's Vulkan header, adds lazy-allocation statistics, and configures portability macros for bundled dependencies.

Evidence:
- Code: `thirdparty/vulkan/patches/0001-VKEnumStringHelper-godot-vulkan.patch` — drivers/vulkan/godot_vulkan.h include
- Code: `thirdparty/vulkan/patches/0002-VMA-godot-vulkan.patch` — drivers/vulkan/godot_vulkan.h include
- Code: `thirdparty/vulkan/patches/0003-VMA-add-vmaCalculateLazilyAllocatedBytes.patch` — vmaCalculateLazilyAllocatedBytes
- Code: `thirdparty/wslay/config.h` — WORDS_BIGENDIAN
- Code: `thirdparty/wslay/patches/0001-msvc-build-fix.patch` — SSIZE_T ssize_t

## D3D12 memory allocation

The subsystem includes block metadata, pools, budgets, defragmentation, and virtual allocations.

D3D12MA's Allocator, Pool, and VirtualBlock APIs manage resource and virtual allocations using allocation callbacks.

Evidence:
- Code: `thirdparty/d3d12ma/D3D12MemAlloc.h` — D3D12MA::Allocator
- Code: `thirdparty/d3d12ma/D3D12MemAlloc.h` — D3D12MA::Pool
- Code: `thirdparty/d3d12ma/D3D12MemAlloc.h` — D3D12MA::VirtualBlock

## Histograms and Huffman codes

The lossless encoder can use multiple histogram groups over image tiles.

Histograms built from backward-reference symbols are clustered and converted into Huffman code lengths, codes, and coded tree headers.

Prerequisites: backward-references-and-color-cache

Code examples: VP8LBackwardRefs

Evidence:
- Code: `thirdparty/libwebp/src/enc/histogram_enc.c` — VP8LHistogram, VP8LGetHistoImageSymbols
- Code: `thirdparty/libwebp/src/utils/huffman_encode_utils.h` — HuffmanTreeToken, HuffmanTreeCode
- Code: `thirdparty/libwebp/src/enc/vp8l_enc.c` — StoreHuffmanTree

## KTX2 container transcoding

KTX2 support is an optional container path alongside direct .basis transcoding.

The KTX2 transcoder retains the source data, parses the KTX2 header and per-level index, then transcodes its texture levels.

Prerequisites: basis-transcoding

Code examples: Ktx2Header, Ktx2LevelIndex, Ktx2TranscoderState

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — ktx2_transcoder
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — ktx2_header
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — ktx2_level_index

## Lossless pixel transform pipeline

Transform choice is analyzed from the input and recorded as lossless coding features.

The lossless encoder transforms ARGB pixels through predictor, subtract-green, cross-color, and color-indexing stages before entropy coding.

Prerequisites: input-picture-planes

Code examples: WebPPicture, VP8Encoder

Evidence:
- Code: `thirdparty/libwebp/src/enc/vp8l_enc.c` — VP8LEncodeImage, ApplySubtractGreenTransform
- Code: `thirdparty/libwebp/src/dsp/lossless.c` — VP8LSubtractGreenFromBlueAndRed, VP8LTransformColorInverse
- Code: `thirdparty/libwebp/src/enc/predictor_enc.c` — VP8LResidualImage

## Lossy macroblock encoding

Analysis, iteration, quantization, filtering, and syntax emission are separate source modules in the lossy path.

The lossy encoder walks the configured picture planes by macroblock, choosing prediction modes and producing residual tokens for the VP8 stream.

Prerequisites: encoder-configuration, input-picture-planes

Code examples: VP8Encoder, WebPPicture, WebPConfig

Evidence:
- Code: `thirdparty/libwebp/src/enc/analysis_enc.c` — VP8EncAnalyze
- Code: `thirdparty/libwebp/src/enc/iterator_enc.c` — VP8IteratorInit
- Code: `thirdparty/libwebp/src/enc/frame_enc.c` — VP8EncTokenLoop
- Code: `thirdparty/libwebp/src/enc/vp8i_enc.h` — struct VP8Encoder

## Metal drawable presentation

QuartzCore Metal layers produce drawables that implement the Metal drawable interface for presentation.

Prerequisites: gpu-command-encoding

Code examples: CA::MetalLayer, CA::MetalDrawable

Evidence:
- Code: `thirdparty/metal-cpp/QuartzCore/CAMetalLayer.hpp` — CA::MetalLayer::nextDrawable
- Code: `thirdparty/metal-cpp/QuartzCore/CAMetalDrawable.hpp` — CA::MetalDrawable

## Navigation avoidance

Both dimensions model agent positions and velocities plus obstacle geometry, and their maps expose agent and obstacle collections.

Agents and obstacles are map members that provide avoidance data alongside pathfinding data.

Prerequisites: navigation-map-composition

Evidence:
- Code: `modules/navigation_2d/nav_agent_2d.h` — NavAgent2D
- Code: `modules/navigation_2d/nav_obstacle_2d.h` — NavObstacle2D
- Code: `modules/navigation_3d/nav_agent_3d.h` — NavAgent3D
- Code: `modules/navigation_3d/nav_obstacle_3d.h` — NavObstacle3D

## Navigation path queries

The query implementations maintain begin and end polygons, search data, path points, and path-return constraints.

Path queries consume a map read iteration, select polygons, and construct a path corridor in 2D or 3D.

Prerequisites: navigation-map-composition

Evidence:
- Code: `modules/navigation_2d/2d/nav_mesh_queries_2d.h` — NavMeshQueries2D
- Code: `modules/navigation_3d/3d/nav_mesh_queries_3d.h` — NavMeshQueries3D
- Code: `modules/navigation_2d/nav_utils_2d.h` — Nav2D::NavigationPoly
- Code: `modules/navigation_3d/nav_utils_3d.h` — Nav3D::NavigationPoly

## HTTP and multiplayer

The two APIs appear as separate scene-main facilities.

HTTPRequest is a Node API for HTTP work, while MultiplayerAPI and MultiplayerPeer define reference-counted multiplayer and packet-peer interfaces.

Prerequisites: node-hierarchy

Evidence:
- Code: `scene/main/http_request.h` — class HTTPRequest : public Node
- Code: `scene/main/multiplayer_api.h` — class MultiplayerAPI : public RefCounted
- Code: `scene/main/multiplayer_peer.h` — class MultiplayerPeer : public PacketPeer

## Object identity and lifetime

Object supplies engine object identity and ObjectDB support, while RefCounted and Ref provide reference-managed object use.

Code examples: Object, ObjectID

Evidence:
- Code: `core/object/object.h` — Object
- Code: `core/object/object_id.h` — ObjectID
- Code: `core/object/ref_counted.h` — RefCounted

## native 3D collision pipeline

The native 3D server manages spaces, bodies, shapes, and joints; its collision code includes GJK/EPA support and SAT shape-pair routines.

Code examples: GodotSpace3D, GodotCollisionObject3D

Evidence:
- Code: `modules/godot_physics_3d/godot_physics_server_3d.h` — GodotPhysicsServer3D
- Code: `modules/godot_physics_3d/gjk_epa.cpp` — GJK and EPA implementation
- Code: `modules/godot_physics_3d/godot_collision_solver_3d_sat.cpp` — SeparatorAxisTest

## Pseudo-random generation

RandomPCG supplies pseudo-random generation and RandomNumberGenerator exposes a reference-counted runtime wrapper.

Evidence:
- Code: `core/math/random_pcg.h` — RandomPCG
- Code: `core/math/random_number_generator.h` — RandomNumberGenerator

## 2D physics nodes

The hierarchy includes static, rigid, character, animatable, and joint node types.

2D physics scene graph nodes provide collision objects, bodies, joints, casts, and collision results.

Prerequisites: scene-graph

Code examples: CollisionObject2D, RigidBody2D, KinematicCollision2D

Evidence:
- Code: `scene/2d/physics/collision_object_2d.h` — CollisionObject2D
- Code: `scene/2d/physics/rigid_body_2d.h` — RigidBody2D
- Code: `scene/2d/physics/kinematic_collision_2d.h` — KinematicCollision2D

## Undo and redo history

The manager separates histories by integer ID and retains reference-counted objects required by undo and redo stacks.

Undo and redo assigns actions and saved versions to a scene session, allowing edits to identify a history and report unsaved state.

Prerequisites: editor-scene-sessions

Code examples: EditorUndoRedoManager::History, EditedScene

Evidence:
- Code: `editor/editor_undo_redo_manager.h` — EditorUndoRedoManager::History
- Code: `editor/editor_data.cpp` — EditorData::is_scene_changed

## TrueType GX variation data

The variation loader is compiled with the TrueType driver and models avar correspondence data.

TrueType GX variation loading maps a face's variation tables and coordinates into variable-font state.

Prerequisites: sfnt-tables

Evidence:
- Code: `thirdparty/freetype/src/truetype/ttgxvar.c` — TT_CONFIG_OPTION_GX_VAR_SUPPORT
- Code: `thirdparty/freetype/src/truetype/ttgxvar.h` — GX_AVarCorrespondenceRec_
- External (official, verified): [OpenType Font Variations, TrueType GX, and Adobe MM Fonts - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html), accessed 2026-07-16

## ZIP archive I/O

The module includes tests for compressed and uncompressed archives.

ZIPReader reads archives and ZIPPacker creates archives through reference-counted API objects.

Code examples: ZIPReader, ZIPPacker

Evidence:
- Code: `modules/zip/zip_reader.h` — ZIPReader
- Code: `modules/zip/zip_packer.h` — ZIPPacker
- Code: `modules/zip/tests/test_zip.h` — ZipTest

## Block texture encoding

This is a separate encoder subsystem from Basis transcoding.

CVTT compression input uses pixel blocks, options, encoding plans, and format-specific BC6H, BC7, ETC, and S3TC computation.

Evidence:
- Code: `thirdparty/cvtt/ConvectionKernels.h` — cvtt::PixelBlockU8
- Code: `thirdparty/cvtt/ConvectionKernels.h` — cvtt::Options
- Code: `thirdparty/cvtt/ConvectionKernels_BC67.cpp` — BC6HData::ModeDescriptor
- Code: `thirdparty/cvtt/ConvectionKernels_ETC.cpp` — cvtt::Internal::ETCComputer

## Dynamic invocation and signals

The object and callable layers call Object methods and signal handlers with Variant argument arrays, including bound and unbound callable forms.

Prerequisites: dynamic-variant, object-identity-lifetime

Code examples: Object, Variant

Evidence:
- Code: `core/object/object.h` — Object::callp
- Code: `core/variant/callable.h` — Callable
- Code: `core/variant/callable_bind.h` — CallableCustomBind

## Entropy bitstreams

The same utility area contains both writing and parsing primitives.

VP8 and VP8L writers serialize codec decisions into range-coded or raw bit-level byte streams while readers maintain cursor, range, and value state.

Evidence:
- Code: `thirdparty/libwebp/src/utils/bit_writer_utils.h` — struct VP8BitWriter, VP8LBitWriter
- Code: `thirdparty/libwebp/src/utils/bit_writer_utils.c` — VP8PutBit, VP8LBitWriterAppend
- Code: `thirdparty/libwebp/src/utils/bit_reader_utils.h` — struct VP8BitReader
- Code: `thirdparty/libwebp/src/utils/bit_reader_inl_utils.h` — VP8GetBit, VP8GetSignedValue

## glTF binary accessor decoding and encoding

Mesh attributes, indices, animation samples, and other numeric glTF values pass through this layer.

GLTFAccessor converts typed values to and from GLTFState’s indexed buffer views, including sparse data when the sparse representation is smaller than dense data.

Prerequisites: gltf-asset-model

Code examples: GLTFAccessor, GLTFBufferView, GLTFState

Evidence:
- Code: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor::_decode_as_numbers and encode_new_sparse_accessor_from_vec3s
- Code: `modules/gltf/structures/gltf_buffer_view.cpp` — GLTFBufferView::load_buffer_view_data
- External (primary, unverified (source anchor not found)): [glTF 2.0 Specification](https://registry.khronos.org/glTF/specs/2.0/glTF-2.0.html), accessed 2026-07-15

## Glyph and face caching

Cache entries are organized around manager-owned caches, cache nodes, glyph families, and MRU ordering.

The cache manager holds faces made available by format drivers and caches character maps, glyph images, and small-bitmap nodes under resource limits.

Prerequisites: font-driver-modules

Code examples: FTC_SNodeRec

Evidence:
- Code: `thirdparty/freetype/src/cache/ftcmanag.h` — FTC_Manager
- Code: `thirdparty/freetype/src/cache/ftcsbits.h` — FTC_SNodeRec_
- Code: `thirdparty/freetype/src/cache/ftcmru.h` — FTC_MruList
- External (official, verified): [Cache Sub-System - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-cache_subsystem.html), accessed 2026-07-16

## Glyph rasterization

The raster and smooth components expose raster-function objects, while other components provide specialized renderers.

Renderer modules turn driver-loaded glyph outlines into monochrome, gray, SDF, or SVG-backed bitmap representations.

Prerequisites: font-driver-modules

Code examples: SVG_RendererRec

Evidence:
- Code: `thirdparty/freetype/src/raster/ftraster.c` — ft_standard_raster
- Code: `thirdparty/freetype/src/smooth/ftgrays.c` — ft_grays_raster
- Code: `thirdparty/freetype/src/svg/svgtypes.h` — SVG_RendererRec_
- External (official, verified): [Scanline Converter - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-raster.html), accessed 2026-07-16

## KTX texture containers

KTX1 and KTX2 are represented by related texture objects with internal constructors and type-specific vtables.

KTX texture containers carry a texture format description, per-level indexing, stream or memory-backed data, virtual dispatch tables, and KTX2 supercompression state.

Prerequisites: texture-format-description

Code examples: KTX2 Texture, KTX2 Private Texture State, KTX Level Index Entry

Evidence:
- Code: `thirdparty/libktx/include/ktx.h` — ktxTexture
- Code: `thirdparty/libktx/lib/texture2.h` — ktxTexture2_private
- Code: `thirdparty/libktx/lib/basis_transcode.cpp` — ktxTexture2_transcodeUastc
- External (official, unverified (source anchor not found)): [KTX - GPU Texture Container Format](https://www.khronos.org/ktx/), accessed 2026-07-15

## Native support algorithms

The support subtree implements deterministic PCG random state evolution alongside compression, noise, color, audio, and packing algorithms.

Code examples: pcg32_random_t

Evidence:
- Code: `thirdparty/misc/pcg.cpp` — pcg32_srandom_r
- Code: `thirdparty/misc/FastNoiseLite.h` — FastNoiseLite

## 2D and 3D navigation queries

2D uses Vector2 paths and 3D uses Vector3 paths.

RefCounted navigation query parameter and result objects exchange path points, path types, and path-owner identifiers with the 2D and 3D navigation server APIs.

Code examples: NavigationPathQueryResult2D

Evidence:
- Code: `servers/navigation_2d/navigation_path_query_parameters_2d.h` — class NavigationPathQueryParameters2D : public RefCounted
- Code: `servers/navigation_2d/navigation_path_query_result_2d.h` — NavigationPathQueryResult2D path getters
- Code: `servers/navigation_3d/navigation_path_query_parameters_3d.h` — class NavigationPathQueryParameters3D : public RefCounted
- Code: `servers/navigation_3d/navigation_path_query_result_3d.h` — NavigationPathQueryResult3D path getters

## 2D and 3D physics queries

The 2D and 3D partitions mirror this server-plus-query arrangement.

Physics servers expose direct body and space state APIs, while RefCounted query parameter and result objects package ray, point, shape, and motion requests.

Evidence:
- Code: `servers/physics_2d/physics_server_2d.h` — class PhysicsServer2D
- Code: `servers/physics_2d/direct_states/physics_direct_space_state_2d.h` — class PhysicsDirectSpaceState2D
- Code: `servers/physics_2d/queries/physics_ray_query_parameters_2d.h` — class PhysicsRayQueryParameters2D : public RefCounted
- Code: `servers/physics_3d/physics_server_3d.h` — class PhysicsServer3D

## 3D physics query contracts

The public query objects retain typed parameters while server-side direct-space operations consume them.

Physics-server types define ray, point, shape, and motion parameter and result records; RefCounted query objects expose those contracts to callers.

Code examples: PhysicsRayQueryParameters3D

Evidence:
- Code: `servers/physics_3d/physics_server_3d_types.h` — PS3DT::RayParameters, RayResult, PointParameters, ShapeParameters, MotionParameters, MotionResult
- Code: `servers/physics_3d/queries/physics_ray_query_parameters_3d.h` — PhysicsRayQueryParameters3D
- Code: `servers/physics_3d/queries/physics_testmotion_query_result_3d.h` — PhysicsTestMotionResult3D

## PostScript font services

The supplied sources include Type 1 and Type 2 decoding, CFF decoding, parser tables, and a separate hinter module.

PSAux supplies parsing, decoding, character maps, and hint services shared by CFF, CID, and Type 1 drivers.

Prerequisites: font-driver-modules

Evidence:
- Code: `thirdparty/freetype/src/psaux/psauxmod.c` — psaux_interface
- Code: `thirdparty/freetype/src/pshinter/pshrec.h` — PS_HintsRec
- External (official, unverified (source anchor not found)): [The Type 1 and CID Drivers - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-t1_cid_driver.html), accessed 2026-07-15

## Bitstream and Huffman decoding

This is the shared compression idiom visible in Basis internal decoding and Brotli decoder code.

Both Basis and Brotli implement bit readers and Huffman decoding tables to recover symbols from compressed bit streams.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder_internal.h` — huffman_decoding_table
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder_internal.h` — bitwise_decoder
- Code: `thirdparty/brotli/dec/huffman.c` — BrotliBuildHuffmanTable

## Reflection metadata

ClassDB, GDType, MethodInfo, and PropertyInfo describe Object classes, inheritance, methods, properties, constants, enums, and signals.

Prerequisites: object-identity-lifetime

Code examples: Object, GDType, MethodInfo, PropertyInfo

Evidence:
- Code: `core/object/class_db.h` — ClassDB
- Code: `core/object/gdtype.h` — GDType
- Code: `core/object/method_info.h` — MethodInfo
- Code: `core/object/property_info.h` — PropertyInfo

## Regular-expression matching

RegExMatch is the cross-call result object for search operations rather than a local temporary.

The regex module exposes compiled regular expressions and RefCounted match results with ranges.

Code examples: RegExMatch

Evidence:
- Code: `modules/regex/regex.h` — class RegExMatch; struct Range; class RegEx
- Code: `modules/regex/regex.cpp` — RegEx::_sub

## Script languages and instances

The scripting layer attaches Object-facing script instances, manages registered script languages, and provides extension-backed script implementations.

Prerequisites: object-identity-lifetime

Code examples: Object, Variant

Evidence:
- Code: `core/object/script_language.h` — ScriptServer
- Code: `core/object/script_language.h` — ScriptLanguage
- Code: `core/object/script_instance.h` — ScriptInstance
- Code: `core/object/script_language_extension.h` — ScriptInstanceExtension

## Target-platform export

The base platform type supplies common packaging facilities while derived platform types supply target behavior.

A target platform implementation supplies target-specific validation, option, run, and project, pack, or ZIP export behavior to export orchestration.

Prerequisites: export-orchestration

Code examples: EditorExportPlatform, EditorExportPlatform::ExportMessage

Evidence:
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::save_pack
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::save_zip

## 3D physics nodes

The hierarchy includes static, rigid, character, physical-bone, soft-body, and vehicle implementations.

3D physics scene graph nodes provide collision objects, bodies, joints, casts, soft bodies, and vehicles.

Prerequisites: scene-graph

Code examples: CollisionObject3D, RigidBody3D, SoftBody3D

Evidence:
- Code: `scene/3d/physics/collision_object_3d.h` — CollisionObject3D
- Code: `scene/3d/physics/rigid_body_3d.h` — RigidBody3D
- Code: `scene/3d/physics/soft_body_3d.h` — SoftBody3D

## Transform, quantization, and rate-distortion search

The code includes both direct quantization and trellis-based rate-distortion choices.

The macroblock path transforms residuals, quantizes them with VP8Matrix parameters, evaluates rate-distortion costs, and reconstructs blocks.

Prerequisites: lossy-macroblock-encoding

Code examples: VP8Encoder

Evidence:
- Code: `thirdparty/libwebp/src/enc/quant_enc.c` — TrellisQuantizeBlock, VP8SetSegmentParams
- Code: `thirdparty/libwebp/src/enc/cost_enc.h` — struct VP8Residual
- Code: `thirdparty/libwebp/src/dsp/enc.c` — VP8FTransform, VP8EncQuantizeBlock

## UPnP device control

UPNPDevice is the public device object; UPNPMiniUPNP and UPNPDeviceMiniUPNP provide the third-party-backed forms.

The UPnP module represents generic devices and MiniUPnP-backed specializations behind RefCounted APIs.

Code examples: UPNPDevice

Evidence:
- Code: `modules/upnp/upnp.h` — class UPNP : public RefCounted
- Code: `modules/upnp/upnp_device.h` — class UPNPDevice : public RefCounted
- Code: `modules/upnp/upnp_miniupnp.h` — class UPNPMiniUPNP : public UPNP

## Web JavaScript bridge

This is the explicit engine-to-browser object boundary, distinct from browser input and audio adapters.

The JavaScript bridge exposes evaluation, interface lookup, callback creation, object creation, buffer conversion, downloads, PWA operations, and JavaScript object proxies.

Code examples: JavaScriptObjectImpl

Evidence:
- Code: `platform/web/api/javascript_bridge_singleton.h` — JavaScriptBridge
- Code: `platform/web/javascript_bridge_singleton.cpp` — JavaScriptObjectImpl
- Code: `platform/web/js/libs/library_godot_javascript_singleton.js` — GodotJSWrapper

## WebP image I/O

It conditionally compiles bundled libwebp sources through its module build script.

The WebP module implements image loading, saving, and shared buffer helpers.

Code examples: ImageLoaderWebP, ResourceSaverWebP

Evidence:
- Code: `modules/webp/image_loader_webp.h` — ImageLoaderWebP
- Code: `modules/webp/resource_saver_webp.h` — ResourceSaverWebP
- Code: `modules/webp/webp_common.h` — webp_common

## ZIP64 archive I/O

MiniZip reads and writes ZIP archives while its I/O API abstracts file opening, seeking, telling, and allocation callbacks.

Code examples: zlib_filefunc64_32_def

Evidence:
- Code: `thirdparty/minizip/ioapi.h` — zlib_filefunc64_32_def
- Code: `thirdparty/minizip/zip.c` — zipOpenNewFileInZip4_64

## Brotli stream decompression

Its public decoding path is stateful and can return output incrementally.

The Brotli decoder consumes a guarded byte reader, decodes Huffman and prefix-coded streams, and exposes output from a decoder state.

Prerequisites: prefix-code-decoding

Code examples: BrotliDecoderState

Evidence:
- Code: `thirdparty/brotli/dec/bit_reader.h` — BrotliBitReader
- Code: `thirdparty/brotli/dec/decode.c` — BrotliDecoderTakeOutput
- Code: `thirdparty/brotli/dec/state.h` — BrotliDecoderStateStruct

## Class-reference generation

The Python tool owns the intermediate documentation model and output formatting.

Documentation generation parses class-reference XML into class, type, property, parameter, signal, method, and constant definitions before producing reStructuredText.

Code examples: ClassDef, TypeName, ClassStatusProgress

Evidence:
- Code: `doc/tools/make_rst.py` — State.parse_class
- Code: `doc/tools/make_rst.py` — TypeName.from_element
- Code: `doc/tools/make_rst.py` — make_rst_class

## Endian-safe binary I/O

The implementation uses conditional endianness macros instead of assuming a host byte order.

Endian macros and memory conversion helpers normalize integer fields while container and bitstream code serializes binary bytes.

Code examples: WebPData

Evidence:
- Code: `thirdparty/libwebp/src/utils/endian_inl_utils.h` — HToLE32, HToLE16, BSwap32
- Code: `thirdparty/libwebp/src/webp/format_constants.h` — MKFOURCC
- Code: `thirdparty/libwebp/src/mux/muxinternal.c` — MuxAssemble

## Native extension loading

This lets native libraries add engine-facing types without compiling them into the core binary.

A GDExtension is a Resource loaded through a loader; its registered classes become ClassDB-visible extension classes.

Prerequisites: reflection, resources

Code examples: GDExtension, ClassInfo

Evidence:
- Code: `core/extension/gdextension.h` — GDExtension
- Code: `core/object/class_db.h` — ClassDB::register_extension_class
- Code: `doc/classes/GDExtensionManager.xml` — GDExtensionManager

## Physics queries and motion tests

Direct-space-state calls provide ad hoc queries, while cast nodes keep query behavior in the scene tree.

Physics queries use physics spaces to submit configured point, ray, shape, and motion tests and receive collision data.

Prerequisites: physics-spaces

Code examples: PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D

Evidence:
- Code: `doc/classes/PhysicsDirectSpaceState2D.xml` — PhysicsDirectSpaceState2D
- Code: `doc/classes/PhysicsShapeQueryParameters2D.xml` — PhysicsShapeQueryParameters2D
- Code: `doc/classes/PhysicsTestMotionResult2D.xml` — PhysicsTestMotionResult2D

## Platform export packaging

Export code is platform-specific but follows EditorExportPlatform-derived services.

Platform export plug-ins implement macOS, Web, Windows, and visionOS export behavior; the Windows exporter also includes PE template modification structures.

Evidence:
- Code: `platform/macos/export/export_plugin.h` — EditorExportPlatformMacOS
- Code: `platform/web/export/export_plugin.h` — EditorExportPlatformWeb
- Code: `platform/windows/export/template_modifier.h` — TemplateModifier
- Code: `platform/visionos/export/export_plugin.h` — EditorExportPlatformVisionOS

## Signed-distance-field glyph rendering

Common SDF support includes distance normalization and shared renderer properties.

The SDF and BSDF renderers use rasterization to populate signed-distance output bitmaps with configurable parameters.

Prerequisites: glyph-rasterization

Evidence:
- Code: `thirdparty/freetype/src/sdf/sdf.c` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/sdf/ftsdfcommon.c` — ft_sdf_format
- External (official, verified): [Scanline Converter - FreeType API Reference](https://freetype.org/freetype2/docs/reference/ft2-raster.html), accessed 2026-07-16

## Theora block video coding

Encoding and decoding are implemented separately but share transform, quantization, fragment, and state machinery.

The Theora codec consumes Ogg packets while unpacking quantization parameters, DCT token data, motion-compensated frame state, and optional accelerated transform routines.

Prerequisites: ogg-pages-and-packets

Code examples: Theora Stream Information, Ogg Packet

Evidence:
- Code: `thirdparty/libtheora/dequant.c` — oc_quant_params_unpack
- Code: `thirdparty/libtheora/decode.c` — th_decode_packetin
- Code: `thirdparty/libtheora/mcenc.c` — oc_mcenc_search
- External (primary, unverified (source anchor not found)): [Theora Specification](https://www.theora.org/doc/Theora.pdf), accessed 2026-07-15

## Undo and redo actions

UndoRedo records Object/Callable operations as actions so they can be replayed in undo or redo direction.

Prerequisites: dynamic-invocation-signals

Code examples: Object, Variant

Evidence:
- Code: `core/object/undo_redo.h` — UndoRedo
- Code: `core/object/undo_redo.h` — UndoRedo::Operation
- Code: `core/object/undo_redo.h` — UndoRedo::Action

## UPnP device discovery

MiniUPnPc discovers UPnP devices, represents them as a linked device list, and parses an IGD description into URLs and service data.

Code examples: UPNPDev, UPNPUrls, IGDdatas

Evidence:
- Code: `thirdparty/miniupnpc/src/minissdpc.h` — getDevicesFromMiniSSDPD
- Code: `thirdparty/miniupnpc/src/miniupnpc.c` — UPNP_GetValidIGD

## Graphics pipeline configuration

The header represents pipeline setup as a composed configuration record rather than a local rendering algorithm.

`GraphicsPipelineCreateInfo` groups shader stages with separate vertex-input, assembly, viewport, rasterization, multisample, depth-stencil, color-blend, and dynamic-state records.

Code examples: GraphicsPipelineCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct GraphicsPipelineCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct PipelineShaderStageCreateInfo

## Image decode pipelines

This is the common lifecycle behind the JPEG, PNG, and WebP implementations in this partition.

The vendored image libraries parse encoded input, keep decoder state, and emit raster rows or planes; JPEG has scanline and upsampling modules, PNG applies row transforms, and WebP routes decoded output through VP8 I/O.

Code examples: JPEG Decompression Context, PNG Information State, WebP Decoder State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jdsample.c` — jinit_upsampler
- Code: `thirdparty/libpng/pngtrans.c` — png_set_bgr
- Code: `thirdparty/libwebp/src/dec/io_dec.c` — VP8InitIoInternal
- External (official, unverified (source anchor not found)): [Portable Network Graphics (PNG) Specification (Third Edition)](https://www.w3.org/TR/png-3/), accessed 2026-07-15
- External (official, unverified (source anchor not found)): [WebP Container Specification](https://developers.google.com/speed/webp/docs/riff_container), accessed 2026-07-15

## UPnP port-mapping control

SOAP command helpers add, delete, query, and list port mappings on a discovered device.

Prerequisites: upnp-device-discovery

Code examples: PortMapping, PortMappingParserData

Evidence:
- Code: `thirdparty/miniupnpc/include/miniupnpc/upnpcommands.h` — UPNP_AddPortMapping
- Code: `thirdparty/miniupnpc/src/portlistingparse.c` — ParsePortListing

## JPEG decompression and coefficient transcoding

The normal decompression path is modular; the coefficient path is explicitly documented as a transcoding mode.

The JPEG code implements an image decode pipeline that selects decompression modules, can merge chroma upsampling with YCC-to-RGB conversion, and can expose raw DCT coefficient arrays for transcoding.

Prerequisites: image-decode-pipeline

Code examples: JPEG Decompression Context, JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jdmerge.c` — jinit_merged_upsampler
- Code: `thirdparty/libjpeg-turbo/src/jdtrans.c` — jpeg_read_coefficients
- External (official, unverified (HTTP 406)): [libjpeg-turbo Documentation](https://libjpeg-turbo.org/Documentation/Documentation), accessed 2026-07-15

## PNG streaming I/O and row transforms

I/O customization and pixel transformations are separate from information-structure mutation.

The PNG code implements an image decode pipeline with replaceable read and write callbacks, push-mode input states, metadata validity flags, and row-level transformations such as BGR mapping and 16-bit byte swapping.

Prerequisites: image-decode-pipeline

Code examples: PNG Information State

Evidence:
- Code: `thirdparty/libpng/pngrio.c` — png_default_read_data
- Code: `thirdparty/libpng/pngpread.c` — PNG_READ_SIG_MODE
- Code: `thirdparty/libpng/pngtrans.c` — png_set_swap
- External (official, unverified (source anchor not found)): [Portable Network Graphics (PNG) Specification (Third Edition)](https://www.w3.org/TR/png-3/), accessed 2026-07-15

## WebP chunk parsing, incremental decode, and animation

The decoder, demuxer, and animation decoder are distinct modules over shared WebP byte data.

The WebP code implements an image decode pipeline that recognizes VP8 and VP8L payloads, incrementally retains input, demuxes RIFF chunks into frames, and composites animation frames.

Prerequisites: image-decode-pipeline

Code examples: WebP Decoder State

Evidence:
- Code: `thirdparty/libwebp/src/dec/webp_dec.c` — ParseOptionalChunks
- Code: `thirdparty/libwebp/src/dec/idec_dec.c` — WebPIDecoder
- Code: `thirdparty/libwebp/src/demux/demux.c` — WebPDemuxer
- External (official, unverified (source anchor not found)): [WebP Container Specification](https://developers.google.com/speed/webp/docs/riff_container), accessed 2026-07-15

## PNG image codec

PNG support is isolated as a driver with loader and saver classes.

The PNG driver supplies image loading and resource saving implementations and can build bundled libpng sources.

Prerequisites: driver-build-composition

Evidence:
- Code: `drivers/png/image_loader_png.h` — class ImageLoaderPNG : public ImageFormatLoader
- Code: `drivers/png/resource_saver_png.h` — class ResourceSaverPNG : public ResourceFormatSaver
- Code: `drivers/png/SCsub` — builtin_libpng source selection

## Optional SIMD codec paths

Selection occurs through feature macros, runtime checks where implemented, and function-pointer assignment.

The codec libraries use C conditional compilation to select architecture-specific NEON, SSE2, VSX, LSX, MMX, and other optimized routines, while retaining scalar fallbacks.

Code examples: JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jsimd.h` — WITH_SIMD
- Code: `thirdparty/libpng/intel/intel_init.c` — png_init_filter_functions_sse2
- Code: `thirdparty/libtheora/x86/x86enc.c` — oc_enc_accel_init_x86

## RTTI-based serialization

The same registration pattern is used by settings, materials, constraints, skeletons, and curves.

RTTI-based serialization registers attributes for object types and writes or restores binary state through StreamIn and StreamOut.

Code examples: ConstraintSettings, PhysicsMaterialSimple, Skeleton

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/ObjectStream/SerializableAttribute.h` — SerializableAttribute
- Code: `thirdparty/jolt_physics/Jolt/Core/RTTI.h` — RTTI::GetAttribute
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/PhysicsMaterialSimple.cpp` — PhysicsMaterialSimple::SaveBinaryState and RestoreBinaryState

## SIMD ray packets

SIMD ray packets evaluate width-parameterized ray and triangle lanes, so triangle intersection is required to interpret lane-validity masks.

Prerequisites: triangle-intersection-algorithms

Code examples: TriangleMi

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h` — MoellerTrumboreIntersectorK
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_pluecker.h` — PlueckerIntersectorK
- Code: `thirdparty/embree/kernels/geometry/sphere_intersector.h` — SphereIntersectorK

## ISA and platform portability

The source contains SSE2, AVX, AVX2, AVX-512, ARM, and WebAssembly-oriented paths.

Preprocessor-selected ISA configuration chooses SIMD namespaces and headers, while platform shims adapt unavailable WebAssembly control-register operations.

Code examples: vint4, vint8, vuint16

Evidence:
- Code: `thirdparty/embree/common/sys/sysinfo.h` — ISA selection macros
- Code: `thirdparty/embree/common/simd/vint4_sse2.h` — vint<4>
- Code: `thirdparty/embree/common/simd/vint8_avx2.h` — vint<8>
- Code: `thirdparty/embree/common/simd/wasm/emulation.h` — _MM_SET_EXCEPTION_MASK

## Compile-time platform backends

The source tree contains platform-specific input, loading, timing, synchronization, and device paths.

SDL uses compile-time conditions to select Linux, Windows, macOS, dummy, and other platform backend implementations.

Evidence:
- Code: `thirdparty/sdl/include/build_config/SDL_build_config.h` — SDL_build_config_h_
- Code: `thirdparty/sdl/thread/pthread/SDL_systhread_c.h` — SYS_ThreadHandle
- Code: `thirdparty/sdl/thread/windows/SDL_systhread_c.h` — SYS_ThreadHandle

## State recording and validation

StateRecorderImpl supports rewinding, clearing, reading, writing, and validation diagnostics.

State recording saves and validates Body and Constraint state through a stream that can compare replayed bytes with current values.

Prerequisites: body-lifecycle, constraints, serialization

Code examples: Body, TwoBodyConstraint

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/StateRecorder.h` — StateRecorder
- Code: `thirdparty/jolt_physics/Jolt/Physics/StateRecorderImpl.cpp` — StateRecorderImpl::ReadBytes, WriteBytes, Rewind, Clear
- Code: `thirdparty/jolt_physics/Jolt/Physics/DeterminismLog.h` — DeterminismLog

## HID device enumeration and backends

Linux, macOS, Windows, libusb, and SDL joystick HIDAPI code are present in this partition.

SDL enumerates HID devices into linked device-information records and routes opened devices through platform and driver backends.

Code examples: SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/hidapi/SDL_hidapi.c` — SDL_hid_enumerate
- Code: `thirdparty/sdl/hidapi/hidapi/hidapi.h` — hid_device_info
- Code: `thirdparty/sdl/joystick/hidapi/SDL_hidapijoystick_c.h` — SDL_HIDAPI_Device

## Profiling name interning

The profiling implementation interns StringName metadata and source-location data for tracing backends.

Prerequisites: string-names-paths

Code examples: StringName

Evidence:
- Code: `core/profiling/profiling.cpp` — TracyInternTable
- Code: `core/profiling/profiling.cpp` — intern_source_location

## Rendering backend abstraction

The source partition contains separate backend implementations rather than one graphics API implementation.

Rendering backends specialize common context and device-driver abstractions for Vulkan, GLES3, Direct3D 12, and Metal.

Prerequisites: driver-build-composition

Code examples: RenderingDeviceDriverVulkan::BufferInfo, RenderingDeviceDriverVulkan::CommandBufferInfo

Evidence:
- Code: `drivers/vulkan/rendering_device_driver_vulkan.h` — class RenderingDeviceDriverVulkan : public RenderingDeviceDriver
- Code: `drivers/d3d12/rendering_device_driver_d3d12.h` — class RenderingDeviceDriverD3D12 : public RenderingDeviceDriver
- Code: `drivers/metal/rendering_device_driver_metal.h` — class RenderingDeviceDriverMetal : public RenderingDeviceDriver
- Code: `drivers/gles3/rasterizer_gles3.h` — class RasterizerGLES3 : public RendererCompositor

## Thread and synchronization abstractions

The event queue and device code can use these facilities rather than directly naming every platform primitive.

SDL implements mutex, semaphore, condition, read/write lock, thread, and TLS abstractions with generic, pthread, and Windows backends.

Evidence:
- Code: `thirdparty/sdl/include/SDL3/SDL_mutex.h` — SDL_Mutex
- Code: `thirdparty/sdl/thread/pthread/SDL_sysmutex.c` — SDL_LockMutex
- Code: `thirdparty/sdl/thread/windows/SDL_sysmutex.c` — SDL_CreateMutex

## Vector font export

The vector path, draw, paint, PDF, and SVG files form a separate output family from raster paint.

Color paint records are emitted through vector drawing and paint backends that include PDF and SVG implementations.

Prerequisites: color-font-paint

Code examples: hb_vector_paint_t, hb_vector_draw_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-vector-paint.hh` — hb_vector_paint_t
- Code: `thirdparty/harfbuzz/src/hb-vector-paint-svg.cc` — hb_vector_svg_sweep_ctx_t

## Android JNI interoperation

Native wrappers, Java bridge classes, and callable adapters form the implemented boundary.

JNI interoperation converts values and directs callbacks across the Android platform host’s Java/native boundary using C++ Variant marshalling.

Prerequisites: android-platform-host

Code examples: Callable

Evidence:
- Code: `platform/android/java_class_wrapper.h` — JavaClassWrapper
- Code: `platform/android/java_class_wrapper.cpp` — JavaClassWrapper::_jobject_to_variant
- Code: `platform/android/variant/callable_jni.cpp` — callable_jni
- External (official, unverified (source anchor not found)): [JNI tips](https://developer.android.com/ndk/guides/jni-tips), accessed 2026-07-15

## Android rendering and input

The implementation includes separate GL and Vulkan render-view classes plus a shared Android input handler.

Android rendering views and native rendering drivers carry surface and input work from the Android platform host.

Prerequisites: android-platform-host

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotGLRenderView.java` — GodotGLRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotVulkanRenderView.java` — GodotVulkanRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/input/GodotInputHandler.java` — GodotInputHandler
- Code: `platform/android/rendering_context_driver_vulkan_android.h` — RenderingContextDriverVulkanAndroid

## Deferred calls and worker tasks

MessageQueue records Object/Callable work for later execution, while WorkerThreadPool represents task and group work for worker threads.

Prerequisites: dynamic-invocation-signals

Code examples: Object, Variant

Evidence:
- Code: `core/object/message_queue.h` — MessageQueue
- Code: `core/object/message_queue.h` — MessageQueue::Message
- Code: `core/object/worker_thread_pool.h` — WorkerThreadPool::Task
- Code: `core/templates/command_queue_mt.h` — CommandQueueMT

## Gamepad mapping and classification

A built-in mapping database and controller-type rules are included.

Gamepad mapping classifies HID enumeration results and maps device controls to SDL gamepad axes and buttons.

Prerequisites: hid-device-enumeration

Code examples: SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/joystick/SDL_gamepad.c` — SDL_Gamepad
- Code: `thirdparty/sdl/joystick/SDL_gamepad_db.h` — s_GamepadMappings
- Code: `thirdparty/sdl/joystick/controller_type.c` — SDL_GetGamepadTypeFromVIDPID

## Generic container infrastructure

The template layer supplies the C++ storage and lookup structures that underpin Array and Dictionary, including CowData, Vector, maps, sets, lists, spans, and RID owners.

Prerequisites: variant-containers

Code examples: Array, Dictionary

Evidence:
- Code: `core/templates/cowdata.h` — CowData
- Code: `core/templates/vector.h` — Vector
- Code: `core/templates/hash_map.h` — HashMap
- Code: `core/templates/rid_owner.h` — RID_Owner

## Multi-channel distance fields

MSDFgen represents vector shapes, colors their edges, and generates distance-field pixels for one, three, or four channels.

Code examples: msdfgen::Shape, msdfgen::BitmapSection<float, 3>

Evidence:
- Code: `thirdparty/msdfgen/msdfgen.h` — msdfgen public interface
- Code: `thirdparty/msdfgen/core/msdfgen.cpp` — DistancePixelConversion<MultiDistance>
- External (primary, verified): [Chlumsky/msdfgen](https://github.com/Chlumsky/msdfgen), accessed 2026-07-16

## Native platform adapters

Unix and Windows have parallel implementations for several services, while Apple has additional embedded implementations.

Platform adapters implement filesystem, sockets, IP resolution, operating-system services, and threads behind engine interfaces selected by the driver build.

Prerequisites: driver-build-composition

Evidence:
- Code: `drivers/unix/file_access_unix.h` — class FileAccessUnix : public FileAccess
- Code: `drivers/unix/net_socket_unix.h` — class NetSocketUnix : public NetSocket
- Code: `drivers/windows/net_socket_winsock.h` — class NetSocketWinSock : public NetSocket

## Visual shader vector operations

The node set includes operator, function, distance, refraction, compose, and decompose variants.

Vector-operation nodes apply arithmetic and vector functions as operations within a visual shader graph.

Prerequisites: visual-shader-nodes

Code examples: VisualShaderNodeVectorOp

Evidence:
- Code: `modules/visual_shader/vs_nodes/visual_shader_nodes.h` — VisualShaderNodeVectorOp
- Code: `modules/visual_shader/vs_nodes/visual_shader_nodes.h` — VisualShaderNodeVectorRefract

## Android plugins and signals

The test app registers plugin classes in its manifest and GDScript tests retrieve and connect to plugin signals.

Android plugins are discovered from application metadata and expose annotated methods and signals through JNI interoperation.

Prerequisites: android-jni-interop

Code examples: SignalInfo

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/plugin/GodotPluginRegistry.java` — GodotPluginRegistry
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/plugin/SignalInfo.java` — SignalInfo
- Code: `platform/android/java/app/src/instrumented/AndroidManifest.xml` — org.godotengine.plugin.v2.SignalTestPlugin
- Code: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd` — test_signal_connection

## Android storage bridge

The Java/Kotlin implementation has separate access types rather than a single Android storage path.

Android file and directory access handlers implement asset, filesystem, MediaStore, and SAF paths through JNI interoperation.

Prerequisites: android-jni-interop

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/file/FileAccessHandler.kt` — FileAccessHandler
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/file/MediaStoreData.kt` — MediaStoreData
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/file/SAFData.kt` — SAFData
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/io/directory/DirectoryAccessHandler.kt` — DirectoryAccessHandler

## Apple embedded hosting

This layer packages the engine for Apple embedded targets.

Apple embedded hosting uses SwiftUI and Objective-C++ alongside display, keyboard, view-controller, text-to-speech, and Vulkan-context platform adapters.

Prerequisites: native-platform-adapters

Evidence:
- Code: `drivers/apple_embedded/app.swift` — GodotSwiftUIViewController and SwiftUIApp
- Code: `drivers/apple_embedded/display_server_apple_embedded.h` — class DisplayServerAppleEmbedded : public DisplayServer
- Code: `drivers/apple_embedded/rendering_context_driver_vulkan_apple_embedded.h` — class RenderingContextDriverVulkanAppleEmbedded : public RenderingContextDriverVulkan

## Audio backend adapters

The implementations derive from the common AudioDriver interface.

Audio output is provided by platform adapters selected with the driver build for ALSA, PulseAudio, CoreAudio, WASAPI, and XAudio2.

Prerequisites: native-platform-adapters

Evidence:
- Code: `drivers/alsa/audio_driver_alsa.h` — class AudioDriverALSA : public AudioDriver
- Code: `drivers/pulseaudio/audio_driver_pulseaudio.h` — class AudioDriverPulseAudio : public AudioDriver
- Code: `drivers/coreaudio/audio_driver_coreaudio.h` — class AudioDriverCoreAudio : public AudioDriver
- Code: `drivers/wasapi/audio_driver_wasapi.h` — class AudioDriverWASAPI : public AudioDriver

## GDScript static analysis

The supplied analyzer tests exercise typed arrays and dictionaries, enum types, override signatures, casts, return types, and invalid assignments.

The analyzer consumes parsed script trees to resolve inheritance, infer and check data types, validate annotations, and diagnose invalid expressions.

Prerequisites: gdscript-source-parsing

Code examples: GDScriptParser::Node, GDScript

Evidence:
- Code: `modules/gdscript/gdscript_analyzer.h` — GDScriptAnalyzer
- Code: `modules/gdscript/tests/scripts/analyzer/errors/typed_dictionary.out` — typed dictionary analyzer expectations

## GPU block-compression dispatch

The module packages BC1, BC4, BC6H, alpha-stitching, and RGB-to-RGBA shader sources.

The Betsy compressor relies on GLSL compute shaders that fetch source texels or blocks and write compressed results to storage images.

Code examples: Image

Evidence:
- Code: `modules/betsy/image_compress_betsy.h` — BetsyCompressor
- Code: `modules/betsy/alpha_stitch.glsl` — main

## MIDI input adapters

Each named implementation derives from the common MIDIDriver interface.

MIDI input is supplied by platform adapters for ALSA MIDI, CoreMIDI, and WinMM.

Prerequisites: native-platform-adapters

Evidence:
- Code: `drivers/alsamidi/midi_driver_alsamidi.h` — class MIDIDriverALSAMidi : public MIDIDriver
- Code: `drivers/coremidi/midi_driver_coremidi.h` — class MIDIDriverCoreMidi : public MIDIDriver
- Code: `drivers/winmidi/midi_driver_winmidi.h` — class MIDIDriverWinMidi : public MIDIDriver

## Shader interface mapping and reflection

Program-facing reflection records are derived from shader declarations and linker objects.

IO mapping and reflection use intermediate-tree traversal to expose a compiled shader interface as uniform, buffer, atomic-counter, and pipeline input/output entries.

Prerequisites: shader-source-compilation

Code examples: TProgram, TObjectReflection

Evidence:
- Code: `thirdparty/glslang/glslang/MachineIndependent/iomapper.h` — TDefaultIoResolverBase and TGlslIoMapper
- Code: `thirdparty/glslang/glslang/MachineIndependent/reflection.h` — TReflection
- Code: `thirdparty/glslang/glslang/MachineIndependent/reflection.cpp` — TReflectionTraverser
- External (primary, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.html), accessed 2026-07-15
- External (primary, unverified (source anchor not found)): [SPIR-V Specification](https://registry.khronos.org/SPIR-V/specs/unified1/SPIRV.html), accessed 2026-07-15

## Android instrumented integration tests

The Android test source includes Kotlin test code, Java/Kotlin plugin fixtures, and a Godot project with GDScript tests.

Android instrumentation tests launch an app project that exercises plugins, signals, storage, and JavaClassWrapper conversions.

Prerequisites: android-plugin-signals

Evidence:
- Code: `platform/android/java/app/src/androidTestInstrumented/java/com/godot/game/GodotAppTest.kt` — GodotAppTest
- Code: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd` — test_signal_emission
- Code: `platform/android/java/app/src/instrumented/assets/test/file_access/file_access_tests.gd` — FileAccessTests
- Code: `platform/android/java/app/src/instrumented/assets/test/javaclasswrapper/java_class_wrapper_tests.gd` — JavaClassWrapperTests

## Apple embedded packaging and signing

The Apple exporter extends the common platform exporter and uses dedicated signing and plugin-configuration types.

Apple embedded packaging adds Xcode project data, assets, Apple plugin configuration, and code-signing structures to a target-platform export.

Prerequisites: target-platform-export

Code examples: EditorExportPlatform, EditorExportPlatform::ExportMessage

Evidence:
- Code: `editor/export/editor_export_platform_apple_embedded.h` — EditorExportPlatformAppleEmbedded
- Code: `editor/export/plugin_config_apple_embedded.h` — PluginConfigAppleEmbedded
- Code: `editor/export/codesign.h` — CodeSign

## GDScript bytecode compilation and execution

GDScriptFunction stores execution metadata, while GDScriptInstance binds a script to an engine object.

The compiler lowers analyzed script trees through code-generation classes, and the VM executes functions using validated calls, getters, setters, and container operations.

Prerequisites: gdscript-static-analysis

Code examples: GDScript, GDScriptInstance, GDScriptFunction

Evidence:
- Code: `modules/gdscript/gdscript_byte_codegen.h` — GDScriptByteCodeGenerator
- Code: `modules/gdscript/gdscript_vm.cpp` — GDScriptFunction VM execution

## Graphite SILF rule execution

Rules are loaded from Graphite font data, matched during passes, and evaluated by the machine implementation.

Graphite executes decoded SILF rule constraints and actions through a finite-state matcher and bytecode machine.

Prerequisites: font-table-access

Code examples: Silf, Segment, Slot

Evidence:
- Code: `thirdparty/graphite/src/Silf.cpp` — graphite2::Silf
- Code: `thirdparty/graphite/src/Pass.cpp` — graphite2::Pass
- Code: `thirdparty/graphite/src/Code.cpp` — Machine::Code::decoder
- Code: `thirdparty/graphite/src/inc/Rule.h` — Rule, RuleEntry, State, and FiniteStateMachine
- External (official, unverified (HTTP 403)): [Graphite technical overview](https://graphite.sil.org/graphite_techAbout.html), accessed 2026-07-15

## Graphite segment shaping

A face selects the applicable SILF behavior, a font supplies scaled metrics, and a segment exposes shaped slots.

Graphite uses decoded font tables and rule passes to turn Unicode into a Segment whose linked Slots carry glyph and placement state.

Prerequisites: font-table-access, graphite-rule-execution

Code examples: Face, Font, Segment, Slot, FeatureVal

Evidence:
- Code: `thirdparty/graphite/src/inc/Segment.h` — graphite2::Segment
- Code: `thirdparty/graphite/src/Segment.cpp` — graphite2::Segment construction
- Code: `thirdparty/graphite/src/gr_segment.cpp` — gr_seg_cinfo, gr_seg_first_slot, and gr_seg_last_slot
- Code: `thirdparty/graphite/src/gr_slot.cpp` — gr_slot_next_in_segment and attachment traversal
- External (official, unverified (HTTP 403)): [Graphite technical overview](https://graphite.sil.org/graphite_techAbout.html), accessed 2026-07-15

## Cycle resolutions

- Removed `undo-redo-history → editor-scene-sessions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `shader-interface-analysis → shader-interface-metadata`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `image-decode-pipeline → brotli-stream-decompression`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `jpeg-decompression-transcoding → image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `jpeg-decompression-transcoding → raster-image-input`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `png-stream-transforms → image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `png-stream-transforms → raster-image-input`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `webp-riff-decoding → image-decode-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `webp-riff-decoding → webp-image-io`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `brotli-stream-decompression → prefix-code-decoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `prefix-code-decoding → histograms-and-huffman-codes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `dynamic-invocation-signals → dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `subdivision-surface-evaluation → geometry-families`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `catmull-clark-patch-construction → half-edge-topology`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gui-controls → gui-control-layout`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `scene-state → packed-scenes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `rendering-assets → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `block-texture-transcoding → astc-block-coding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `ktx-texture-container → image-format-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `ktx2-container-transcoding → texture-compression-pipeline`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `dynamic-invocation-signals → signal-awaitability`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `object-identity-lifetime → object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `shader-reflection → shader-interface-metadata`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `shader-reflection → spirv-shader-reflection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `spirv-shader-reflection → spirv-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `crypto-resources → cryptography`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `diagnostic-expectations → device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `dynamic-variant → dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `native-extensions → gdextension-libraries`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `reflection-metadata → reflection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `object-identity-lifetime → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `upnp-port-mapping → upnp-device-discovery`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `allocation → object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `spatial-indexing → bvh-construction`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `spatial-indexing → bvh-traversal`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `spatial-indexing → motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `skeletal-ragdoll → gltf-node-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `physics-3d-collision-pipeline → convex-collision`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `spatial-indexing → scene-commit`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `themes-and-style-boxes → editor-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `ui-theming → editor-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `themes-and-style-boxes → ui-theming`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gui-control-layout → ui-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `navigation-regions → navigation-agents`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `scene-tree → gltf-node-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `input-events-actions → input-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `bvh-traversal → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `geometry-families → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `motion-blur-geometry → motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `theora-block-video-codec → theora-video-streams`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `primitive-intersection → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `instancing → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `project-settings → editor-and-project-settings`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `variant-containers → dynamic-values`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `signed-distance-fields → glyph-rasterization`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `instancing → openxr-runtime-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `vulkan-instance-setup → openxr-runtime-loading`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `motion-blur → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `runtime-loop-time → motion-blur`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `runtime-loop-time → ray-query`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `concurrency → device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `sdl-thread-abstractions → task-scheduling`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `scene-commit → primitive-references`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `resource-assets → device-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `resource-assets → scene-contexts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `scene-commit → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `scene-commit → scene-contexts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `primitive-references → geometry-resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `scene-contexts → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `reflection → editor-and-project-settings`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `alpha-plane-coding → encoder-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `alpha-plane-coding → input-picture-planes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `yuv-rgb-conversion → input-picture-planes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `unicode-normalization → unicode-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `tls-connection-state → tls-crypto-backend`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `synchronization-primitives → concurrency`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `concurrency → task-scheduling`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `wayland-window-backend → linuxbsd-desktop-integration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `object-model → resources`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `openxr-runtime-integration → openxr-structure-chains`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `localization → text-and-localization`: removed to preserve the stronger inferred prerequisite hierarchy.
<!-- rope-ladder:end document -->
