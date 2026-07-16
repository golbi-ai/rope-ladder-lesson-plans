<!-- rope-ladder:begin document 0c6d5daf6f7136c6e2a66af2c69e5ec7c3e6db62386f76b6e1de4c4504421428 -->
# Survival guide

A terse navigation guide for the ideas most useful when entering this codebase.

## Animation is canvas-state optimization

Animation encoding compares current and prior canvases before choosing a rectangle, blend/disposal strategy, and image encoding candidate.

Start in: thirdparty/libwebp/src/mux/anim_encode.c, thirdparty/libwebp/src/webp/mux.h

## Audio has three distinct composition modes

Interactive transitions, playlists, and synchronization are separate AudioStream resource families.

Start in: modules/interactive_music/audio_stream_interactive.h, modules/interactive_music/audio_stream_playlist.h, modules/interactive_music/audio_stream_synchronized.h

Related topics: [`interactive-music-transitions`](../references/domain-concepts.md#interactive-music-transition-tables)

## Bodies are manager-owned

Use BodyID and body interfaces or locks rather than treating Body pointers as globally stable application handles.

Start in: thirdparty/jolt_physics/Jolt/Physics/Body/BodyInterface.h, thirdparty/jolt_physics/Jolt/Physics/Body/BodyLock.h, thirdparty/jolt_physics/Jolt/Physics/Body/BodyManager.h

Related topics: [`cxx-raii-reference-ownership`](../references/language-concepts.md#raii-non-copyable-resources-and-intrusive-references)

## Build behavior is module-local Python

Check each module's config.py and SCsub before assuming it is compiled for a target.

Start in: modules/godot_physics_2d/config.py, modules/jolt_physics/SCsub, modules/lightmapper_rd/SCsub

Related topics: [`module-build-configuration`](../references/domain-concepts.md#scons-module-build-configuration), [`python-scons-configuration`](../references/language-concepts.md#python-scons-configuration-functions)

## Build composition is Python-driven

SCsub and config.py scripts decide source inclusion, bundled third-party compilation, dependencies, and backend-specific variants.

Start in: modules/openxr/extensions/SCsub, modules/svg/config.py, modules/theora/SCsub

Related topics: [`python-scons-module-configuration`](../references/language-concepts.md#python-scons-module-configuration)

## Build inclusion is controlled by Python and SCsub

Changing a native module requires checking its config.py capability hooks and SCsub source-list conditions.

Start in: modules/mp3/config.py, modules/navigation_3d/SCsub, modules/noise/SCsub

Related topics: [`python-scons-build-hooks`](../references/language-concepts.md#python-scons-build-hooks)

## Check build conditions before assuming availability

SCsub files conditionally include 2D physics, navigation, and 3D source groups.

Start in: scene/resources/2d/SCsub, scene/resources/SCsub, servers/navigation_3d/SCsub, servers/physics_2d/SCsub

Related topics: [`python-build-scripts`](../references/language-concepts.md#python-build-scripts)

## Check feature macros before tracing ThorVG paths

The bundled ThorVG configuration explicitly enables the software rasterizer, SVG loader, PNG loader, and non-web threading.

Start in: thirdparty/thorvg/inc/config.h

Related topics: [`c-preprocessor-configuration`](../references/language-concepts.md#c-preprocessor-feature-and-platform-configuration)

## Choose the codec branch early

WebPConfig.lossless and WebPPicture.use_argb determine whether to follow the VP8 lossy path or the VP8L lossless path.

Start in: thirdparty/libwebp/src/enc/vp8l_enc.c, thirdparty/libwebp/src/enc/webp_enc.c, thirdparty/libwebp/src/webp/encode.h

Related topics: [`encoder-configuration`](../references/domain-concepts.md#encoder-configuration), [`lossless-transform-pipeline`](../references/domain-concepts.md#lossless-pixel-transform-pipeline), [`lossy-macroblock-encoding`](../references/domain-concepts.md#lossy-macroblock-encoding)

## Choose the reflection implementation deliberately

SPIRV-Cross offers compiler-side reflection facilities, while the vendored SPIRV-Reflect C API independently exposes reflection-module queries.

Start in: thirdparty/spirv-cross/spirv_cross.hpp, thirdparty/spirv-reflect/spirv_reflect.c, thirdparty/spirv-reflect/spirv_reflect.h

Related topics: [`shader-reflection`](../references/domain-concepts.md#shader-reflection)

## Concurrency is opt-in through jobs

JobSystem captures dependency state with atomics and offers both immediate single-threaded execution and thread-pool implementations.

Start in: thirdparty/jolt_physics/Jolt/Core/JobSystem.h, thirdparty/jolt_physics/Jolt/Core/JobSystem.inl, thirdparty/jolt_physics/Jolt/Core/JobSystemSingleThreaded.cpp

## Contacts are solver data

The narrow phase produces information that ContactConstraintManager caches and turns into solver constraints.

Start in: thirdparty/jolt_physics/Jolt/Physics/Constraints/ContactConstraintManager.cpp, thirdparty/jolt_physics/Jolt/Physics/Constraints/ContactConstraintManager.h

Related topics: [`constraints`](../references/domain-concepts.md#constraint-solving), [`contact-management`](../references/domain-concepts.md#contact-manifolds-and-warm-starting)

## Diff structures are the VCS contract

Read EditorVCSInterface nested diff types before changing the version-control UI, because they define the data exchanged with VCS backends.

Start in: editor/version_control/editor_vcs_interface.h, editor/version_control/version_control_editor_plugin.h

Related topics: [`version-control-integration`](../references/domain-concepts.md#version-control-integration)

## Do not conflate pathfinding and avoidance

Map regions and links form pathfinding topology, while agents and obstacles supply avoidance data.

Start in: modules/navigation_2d/nav_agent_2d.h, modules/navigation_2d/nav_obstacle_2d.h, modules/navigation_3d/nav_agent_3d.h, modules/navigation_3d/nav_obstacle_3d.h

Related topics: [`navigation-avoidance`](../references/domain-concepts.md#navigation-avoidance), [`navigation-map-composition`](../references/domain-concepts.md#navigation-map-composition)

## Extensions reuse core simulation concepts

Vehicles, characters, soft bodies, and ragdolls build on bodies, shapes, collision queries, and constraints instead of replacing the core pipeline.

Start in: thirdparty/jolt_physics/Jolt/Physics/Character/CharacterVirtual.h, thirdparty/jolt_physics/Jolt/Physics/Ragdoll/Ragdoll.h, thirdparty/jolt_physics/Jolt/Physics/SoftBody/SoftBodyMotionProperties.h, thirdparty/jolt_physics/Jolt/Physics/Vehicle/VehicleConstraint.h

Related topics: [`skeletal-ragdoll`](../references/domain-concepts.md#skeletons-animation-and-ragdolls), [`vehicle-dynamics`](../references/domain-concepts.md#vehicle-dynamics)

## Find Godot deltas before reading generated code

The Vulkan patch series identifies local include routing and the added VMA statistic without requiring a pass through generated headers.

Start in: thirdparty/vulkan/patches/0001-VKEnumStringHelper-godot-vulkan.patch, thirdparty/vulkan/patches/0002-VMA-godot-vulkan.patch, thirdparty/vulkan/patches/0003-VMA-add-vmaCalculateLazilyAllocatedBytes.patch

Related topics: [`godot-thirdparty-adaptation`](../references/domain-concepts.md#godot-specific-third-party-adaptation)

## Find script-facing APIs at registration points

ClassDB binding and GDScript’s Script implementation reveal how native behavior reaches scripts.

Start in: core/object/class_db.h, modules/gdscript/gdscript.h

Related topics: [`reflection`](../references/domain-concepts.md#class-metadata-and-reflection), [`scripting`](../references/domain-concepts.md#script-resources-and-instances)

## Follow Clipper Execute outputs

Clipper64 and ClipperD accept paths first, then Execute materializes either flat paths or a polygon hierarchy.

Start in: thirdparty/clipper2/include/clipper2/clipper.engine.h

Related topics: [`cpp-abstraction-and-polymorphism`](../references/language-concepts.md#c-types-encapsulation-and-inheritance), [`polygon-clipping`](../references/domain-concepts.md#polygon-clipping-and-result-trees)

## Follow ENet state through protocol.c

Host creation establishes peer capacity, while protocol command handling changes each peer's state.

Start in: thirdparty/enet/enet/protocol.h, thirdparty/enet/host.c, thirdparty/enet/protocol.c

Related topics: [`enet-host-peer-transport`](../references/domain-concepts.md#enet-host-and-peer-transport), [`enet-wire-commands`](../references/domain-concepts.md#enet-wire-commands)

## Follow Object through metadata

Object behavior is coupled to GDType and ClassDB rather than isolated in Object alone.

Start in: core/object/class_db.h, core/object/gdtype.h, core/object/object.h

Related topics: [`dynamic-invocation-signals`](../references/domain-concepts.md#dynamic-invocation-and-signals), [`object-identity-lifetime`](../references/domain-concepts.md#object-identity-and-lifetime), [`reflection-metadata`](../references/domain-concepts.md#reflection-metadata)

## Follow PCRE2 JIT into SLJIT only after compilation

Pattern compilation and matching are PCRE2 concerns; the JIT path additionally lowers control flow through SLJIT labels and jumps.

Start in: thirdparty/pcre2/deps/sljit/sljit_src/sljitLir.h, thirdparty/pcre2/src/pcre2_compile.c, thirdparty/pcre2/src/pcre2_jit_compile.c

Related topics: [`regex-compile-match`](../references/domain-concepts.md#regular-expression-compilation-and-matching), [`regex-jit-codegen`](../references/domain-concepts.md#regular-expression-jit-code-generation)

## Follow RID ownership in low-level rendering

RD descriptor objects configure RenderingDevice calls; RIDs identify the live resources those descriptors bind.

Start in: doc/classes/RDTextureFormat.xml, doc/classes/RDUniform.xml, doc/classes/RID.xml, doc/classes/RenderingDevice.xml

Related topics: [`rendering-device-resources`](../references/domain-concepts.md#renderingdevice-descriptors-and-handles), [`resource-identifiers`](../references/domain-concepts.md#resource-and-server-identifiers)

## Follow Resource subclasses for content

Textures, meshes, materials, themes, animations, shapes, and packed scenes all appear as Resource-derived families.

Start in: scene/resources/animation.h, scene/resources/material.h, scene/resources/mesh.h, scene/resources/texture.h

Related topics: [`rendering-resources`](../references/domain-concepts.md#textures-meshes-and-materials), [`resource-assets`](../references/domain-concepts.md#resource-backed-assets)

## Follow SVG into Paint

SvgLoader, SvgNode parsing, and the scene builder are the shortest path from exchanged SVG text to ThorVG drawing objects.

Start in: thirdparty/thorvg/src/loaders/svg/tvgSvgLoader.h, thirdparty/thorvg/src/loaders/svg/tvgSvgLoaderCommon.h, thirdparty/thorvg/src/loaders/svg/tvgSvgSceneBuilder.cpp

## Follow TileMapLayer for cross-cutting tile behavior

TileMapLayer cell data drives rendering, physics, navigation, and debug quadrant work.

Start in: scene/2d/tile_map.h, scene/2d/tile_map_layer.cpp, scene/2d/tile_map_layer.h

Related topics: [`navigation-agents`](../references/domain-concepts.md#navigation-agents-and-regions), [`tilemap-layer-data`](../references/domain-concepts.md#tile-map-layer-data)

## Follow a Graphite result through Segment and Slot

Segment is the run-level result; Slot navigation exposes glyph order and attachment structure.

Start in: thirdparty/graphite/src/gr_segment.cpp, thirdparty/graphite/src/gr_slot.cpp, thirdparty/graphite/src/inc/Segment.h

Related topics: [`graphite-shaping`](../references/domain-concepts.md#graphite-segment-shaping)

## Follow avoidance from neighbors to velocity

For 2D RVO behavior, inspect neighbor computation before the ORCA line construction and linear programs.

Start in: thirdparty/rvo2/rvo2_2d/Agent2d.cpp, thirdparty/rvo2/rvo2_2d/Agent2d.h, thirdparty/rvo2/rvo2_2d/KdTree2d.cpp

## Follow converter dispatch through shared data

UConverter instances delegate encoding-specific behavior through UConverterSharedData and UConverterImpl.

Start in: thirdparty/icu4c/common/ucnv_bld.h, thirdparty/icu4c/common/ucnv_u8.cpp, thirdparty/icu4c/common/ucnvmbcs.cpp

Related topics: [`character-encoding-conversion`](../references/domain-concepts.md#character-encoding-conversion)

## Follow filters at accepted hits

A candidate intersection becomes visible only after geometry and context filter checks accept it.

Start in: thirdparty/embree/kernels/geometry/filter.h, thirdparty/embree/kernels/geometry/intersector_epilog.h

Related topics: [`filter-callbacks`](../references/domain-concepts.md#intersection-and-occlusion-filters), [`hit-results`](../references/domain-concepts.md#intersection-results)

## Follow glTF state first

GLTFState is the shared index table; start there before following document conversion code.

Start in: modules/gltf/gltf_document.cpp, modules/gltf/gltf_state.h

Related topics: [`gltf-accessors`](../references/domain-concepts.md#gltf-binary-accessor-decoding-and-encoding)

## Follow opaque GPU IDs into records

The Vulkan driver header names the resource records; implementation casts from driver IDs reveal where each record is consumed.

Start in: drivers/vulkan/rendering_device_driver_vulkan.cpp, drivers/vulkan/rendering_device_driver_vulkan.h

## Follow pNext deliberately

When a record has `pNext`, inspect the compatible extension records before assuming all configuration lives in the base record.

Start in: thirdparty/vulkan/include/vulkan/vulkan_structs.hpp, thirdparty/vulkan/include/vulkan/vulkan_vi.h

Related topics: [`vulkan-extensible-records`](../references/domain-concepts.md#vulkan-extensible-records)

## Follow preload chains as type dependencies

Preloaded scripts act as values, type names, and possible cycles in these fixtures.

Start in: modules/gdscript/tests/scripts/analyzer/features/preload_cyclic_reference.gd, modules/gdscript/tests/scripts/analyzer/features/preload_cyclic_reference_a.notest.gd, modules/gdscript/tests/scripts/analyzer/features/preload_cyclic_reference_b.notest.gd

Related topics: [`classes-inheritance-and-lookup`](../references/language-concepts.md#classes-inheritance-and-lookup)

## Follow resources through loaders

Resource loading is a registry plus cache and threaded-task system, not a direct file parser.

Start in: core/io/resource.h, core/io/resource_loader.h

Related topics: [`resource-formats`](../references/domain-concepts.md#resource-formats-and-serialization), [`resource-loading`](../references/domain-concepts.md#resource-loading-and-caching)

## Follow server-backed features from scene nodes outward

Physics, navigation, and multiplayer all pair scene-facing objects with lower-level service interfaces.

Start in: scene/main/multiplayer_api.h

Related topics: [`navigation`](../references/domain-concepts.md#navigation-maps-and-path-queries), [`physics-collision`](../references/domain-concepts.md#physics-shapes-objects-and-collision-reports)

## Follow settings by storage boundary

EditorSettings owns editor-side settings and metadata, while ProjectSettingsEditor is a separate project-configuration surface.

Start in: editor/settings/editor_settings.h, editor/settings/project_settings_editor.h

Related topics: [`editor-and-project-settings`](../references/domain-concepts.md#editor-and-project-settings), [`input-action-configuration`](../references/domain-concepts.md#input-action-and-shortcut-configuration)

## Follow shader translation from AST to blocks

Parse-context and intermediate-tree code precede GlslangToSpv lowering, which constructs Module, Function, and Block IR objects.

Start in: thirdparty/glslang/SPIRV/GlslangToSpv.cpp, thirdparty/glslang/SPIRV/spvIR.h, thirdparty/glslang/glslang/MachineIndependent/ParseHelper.h

Related topics: [`shader-language-front-end`](../references/domain-concepts.md#shader-language-front-end), [`spirv-generation`](../references/domain-concepts.md#spir-v-generation)

## Follow shader work end to end

Shader source is built into generated headers, processed by shader tooling, and then bound by RD renderer/effect services.

Start in: servers/rendering/renderer_rd/shader_rd.h, servers/rendering/renderer_rd/shaders/SCsub, servers/rendering/shader_compiler.h, servers/rendering/shader_preprocessor.h

## Follow stream ownership carefully

The Bzip2 wrapper owns decompressor state and its own allocation, but does not close the source stream.

Start in: thirdparty/freetype/src/bzip2/ftbzip2.c

Related topics: [`c-explicit-resource-lifecycles`](../references/language-concepts.md#c-explicit-resource-lifecycles), [`compressed-font-stream-adapters`](../references/domain-concepts.md#compressed-font-stream-adapters), [`font-streams`](../references/domain-concepts.md#font-streams)

## Follow the GDScript pipeline

Read tokenizer and parser first, then analyzer, compiler/bytecode generator, and VM.

Start in: modules/gdscript/gdscript_analyzer.cpp, modules/gdscript/gdscript_compiler.cpp, modules/gdscript/gdscript_parser.h, modules/gdscript/gdscript_tokenizer.h, modules/gdscript/gdscript_vm.cpp

Related topics: [`gdscript-bytecode-runtime`](../references/domain-concepts.md#gdscript-bytecode-compilation-and-execution), [`gdscript-static-analysis`](../references/domain-concepts.md#gdscript-static-analysis)

## Follow the desktop backend split

Linux/BSD services live at the platform layer, while Wayland has its own display server, protocol thread, and embedder.

Start in: platform/linuxbsd/os_linuxbsd.cpp, platform/linuxbsd/wayland/display_server_wayland.cpp, platform/linuxbsd/wayland/wayland_thread.cpp

Related topics: [`linuxbsd-desktop-integration`](../references/domain-concepts.md#linux-bsd-desktop-integration), [`wayland-window-backend`](../references/domain-concepts.md#wayland-window-backend)

## Follow the import seam

Start at ResourceImporterScene, then distinguish source parsing from post-import mutation.

Start in: editor/import/3d/editor_import_collada.cpp, editor/import/3d/resource_importer_scene.h

Related topics: [`collada-import`](../references/domain-concepts.md#collada-scene-interchange), [`post-import-processing`](../references/domain-concepts.md#post-import-processing), [`scene-importing`](../references/domain-concepts.md#scene-importing)

## Follow xatlas from MeshDecl to charts and packing

The public declaration is the intake boundary; internal Mesh, parameterization, and packing classes form the main processing pipeline.

Start in: thirdparty/xatlas/xatlas.cpp, thirdparty/xatlas/xatlas.h

## Generated editor data begins in Python

Documentation, exporter registration, and translations are emitted by build helpers rather than hand-maintained C++ tables.

Start in: editor/editor_builders.py

Related topics: [`build-time-source-generation`](../references/domain-concepts.md#build-time-source-generation), [`python-build-source-generation`](../references/language-concepts.md#python-build-source-generation)

## GridMap work starts with octants

Cell-to-octant mapping drives the component's sparse storage and its rendering, navigation, and physics updates.

Start in: modules/gridmap/grid_map.cpp, modules/gridmap/grid_map.h

## IDE messages have a small stable envelope

Message, MessageContent, MessageDecoder, and Peer define the protocol core; typed JSON requests sit above it.

Start in: modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Message.cs, modules/mono/editor/GodotTools/GodotTools.IdeMessaging/MessageDecoder.cs, modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Peer.cs

Related topics: [`ide-messaging-protocol`](../references/domain-concepts.md#ide-messaging-protocol)

## Inspect generated binding paths when virtual behavior is unclear

The object build script generates virtual-method support, complementing the handwritten object and ClassDB code.

Start in: core/object/SCsub, core/object/make_virtuals.py, core/object/object.h

Related topics: [`python-build-code-generation`](../references/language-concepts.md#python-build-code-generation), [`reflection-metadata`](../references/domain-concepts.md#reflection-metadata)

## JPEG processing is a module graph

Follow pointers from jpeg_decompress_struct into entropy, IDCT, upsample, color-conversion, and quantization modules before changing JPEG behavior.

Start in: thirdparty/libjpeg-turbo/src/jdmerge.h, thirdparty/libjpeg-turbo/src/jdsample.h, thirdparty/libjpeg-turbo/src/jpeglib.h

Related topics: [`c-aggregate-callback-modules`](../references/language-concepts.md#c-aggregate-state-and-callback-modules), [`jpeg-decompression-transcoding`](../references/domain-concepts.md#jpeg-decompression-and-coefficient-transcoding)

## KTX2 levels are offset-indexed

KTX2 private state retains a growable level index whose offsets are relative to image-data start; inspect that state before reasoning about streamed texture data or transcoding.

Start in: thirdparty/libktx/lib/basis_transcode.cpp, thirdparty/libktx/lib/texture2.c, thirdparty/libktx/lib/texture2.h

Related topics: [`cpp-basis-transcoding`](../references/language-concepts.md#c-basis-transcoding-behind-a-c-compatible-api), [`ktx-texture-container`](../references/domain-concepts.md#ktx-texture-containers), [`texture-format-description`](../references/domain-concepts.md#texture-format-descriptions)

## Keep Apple export separate from generic packaging

Apple embedded export is a platform specialization with separate plugin configuration and code-signing structures.

Start in: editor/export/codesign.h, editor/export/editor_export_platform_apple_embedded.h, editor/export/plugin_config_apple_embedded.h

Related topics: [`apple-embedded-packaging`](../references/domain-concepts.md#apple-embedded-packaging-and-signing)

## Keep Apple hosting separate from renderer logic

The Apple embedded target has its own SwiftUI and Objective-C++ host layer while its Vulkan context remains a rendering specialization.

Start in: drivers/apple_embedded/app.swift, drivers/apple_embedded/apple_embedded.mm, drivers/apple_embedded/rendering_context_driver_vulkan_apple_embedded.h

Related topics: [`apple-embedded-hosting`](../references/domain-concepts.md#apple-embedded-hosting)

## Keep Jolt AABB construction isolated

The supplied Jolt code is a focused geometry builder with nodes, triangles, and leaf-size policy.

Start in: thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.cpp, thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.h

Related topics: [`aabb-tree-construction`](../references/domain-concepts.md#aabb-tree-construction)

## Keep frame policy separate from startup

`Main` handles initialization while `MainTimerSync` owns accumulated-time and physics-step calculations.

Start in: main/main.cpp, main/main_timer_sync.cpp, main/main_timer_sync.h

Related topics: [`cpp-frame-scheduling`](../references/language-concepts.md#c-frame-time-arithmetic), [`engine-bootstrap`](../references/domain-concepts.md#engine-bootstrap), [`frame-timing`](../references/domain-concepts.md#frame-timing-and-physics-stepping)

## Keep shader IR separate from output back ends

SPIRV-Cross parses shared typed IR first; GLSL and MSL emitters are specialized compiler layers over it.

Start in: thirdparty/spirv-cross/spirv_common.hpp, thirdparty/spirv-cross/spirv_glsl.hpp, thirdparty/spirv-cross/spirv_msl.hpp, thirdparty/spirv-cross/spirv_parser.hpp

## Keep the Recast build stages in order

The geometry pipeline is heightfield, compact heightfield, contours, polygon mesh, then optional detail mesh.

Start in: thirdparty/recastnavigation/Recast/Include/Recast.h, thirdparty/recastnavigation/Recast/Source/RecastContour.cpp, thirdparty/recastnavigation/Recast/Source/RecastMesh.cpp

Related topics: [`navmesh-contours-polygons`](../references/domain-concepts.md#navigation-contours-and-polygons), [`navmesh-heightfields`](../references/domain-concepts.md#navigation-heightfields)

## Keep vendored subsystem boundaries distinct

AccessKit exposes an accessibility ABI, FSR2 dispatches rendering work, and Basis Universal encodes textures; their source trees are separate integrations.

Start in: thirdparty/accesskit/include/accesskit.h, thirdparty/amd-fsr2/ffx_fsr2.cpp, thirdparty/basis_universal/encoder/basisu_comp.cpp

Related topics: [`temporal-upscaling-dispatch`](../references/domain-concepts.md#temporal-upscaling-dispatch), [`texture-compression-pipeline`](../references/domain-concepts.md#texture-compression-pipeline)

## Managed integration spans native and SDK code

CSharpScript is the native bridge, while the .NET SDK projects test generated C# members and metadata.

Start in: modules/mono/csharp_script.h, modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/ScriptPropertiesGeneratorTests.cs

Related topics: [`csharp-partial-source-generation`](../references/language-concepts.md#c-partial-types-and-source-generator-tests), [`managed-csharp-scripting`](../references/domain-concepts.md#managed-c-script-bridge-and-source-generation)

## Migration preserves comment knowledge

Read SourceLine and the conversion/check method pairs before changing project-upgrade rules.

Start in: editor/project_upgrade/project_converter_3_to_4.h, editor/project_upgrade/renames_map_3_to_4.cpp

Related topics: [`project-upgrade`](../references/domain-concepts.md#project-source-migration)

## Model HarfBuzz shaping as buffer mutation

HarfBuzz applies a font and features to a buffer, replacing Unicode-oriented contents with glyph information and positions.

Start in: thirdparty/harfbuzz/src/hb-buffer.hh, thirdparty/harfbuzz/src/hb-ot-layout-gsubgpos.hh, thirdparty/harfbuzz/src/hb-ot-shape.cc

## Model extensions as resources with lifecycle

GDExtension retains a loader, registered classes, and explicit open, initialize, deinitialize, and close operations.

Start in: core/extension/gdextension.h, core/extension/gdextension_manager.h

Related topics: [`gdextension-libraries`](../references/domain-concepts.md#gdextension-libraries), [`resource-loading`](../references/domain-concepts.md#resource-loading-and-caching)

## Navigation is deliberately paired by dimension

The 2D and 3D modules mirror map, iteration, region, link, agent, obstacle, builder, generator, and query roles.

Start in: modules/navigation_2d/2d/nav_mesh_queries_2d.h, modules/navigation_2d/nav_map_2d.h, modules/navigation_3d/3d/nav_mesh_queries_3d.h, modules/navigation_3d/nav_map_3d.h

Related topics: [`navigation-map-composition`](../references/domain-concepts.md#navigation-map-composition), [`navigation-mesh-construction`](../references/domain-concepts.md#navigation-mesh-construction), [`navigation-path-queries`](../references/domain-concepts.md#navigation-path-queries)

## Ogg is the codec boundary

Vorbis and Theora use Ogg packet types at their codec boundaries, so investigate page and packet framing before debugging encoded-media handoff.

Start in: thirdparty/libogg/ogg/ogg.h, thirdparty/libtheora/theora/theoradec.h, thirdparty/libvorbis/synthesis.c

Related topics: [`ogg-pages-and-packets`](../references/domain-concepts.md#ogg-pages-packets-and-bit-packing), [`theora-block-video-codec`](../references/domain-concepts.md#theora-block-video-coding), [`vorbis-block-synthesis`](../references/domain-concepts.md#vorbis-block-analysis-and-synthesis)

## Optimization paths are configuration-dependent

Architecture-specific routines are selected with macros and initializers, so validate the active build configuration before attributing behavior to a scalar implementation.

Start in: thirdparty/libjpeg-turbo/src/jpeg_nbits.h, thirdparty/libpng/intel/intel_init.c, thirdparty/libtheora/x86/x86enc.c

Related topics: [`c-platform-selection`](../references/language-concepts.md#c-preprocessor-platform-and-precision-selection), [`simd-accelerated-codecs`](../references/domain-concepts.md#optional-simd-codec-paths)

## Platform performance paths are conditional

Vec and matrix operations select scalar or instruction-set-specific implementations through compile-time feature macros.

Start in: thirdparty/jolt_physics/Jolt/Core/Core.h, thirdparty/jolt_physics/Jolt/Math/Vec3.inl, thirdparty/jolt_physics/Jolt/Math/Vec4.inl

Related topics: [`cxx-preprocessor-configuration`](../references/language-concepts.md#preprocessor-selected-platform-configuration), [`cxx-simd-alignment`](../references/language-concepts.md#simd-wrappers-and-alignment)

## Project records are richer than paths

ProjectList::Item is the key record for project-manager behavior, sorting, filtering, compatibility, and launch context.

Start in: editor/project_manager/project_list.cpp, editor/project_manager/project_list.h

Related topics: [`project-catalog`](../references/domain-concepts.md#project-catalog-and-selection), [`run-management`](../references/domain-concepts.md#running-projects-and-instances)

## Read ICU from data structures outward

Generated property tables and tries underpin normalization, properties, locale behavior, and spoof checks.

Start in: thirdparty/icu4c/common/normalizer2impl.h, thirdparty/icu4c/common/uchar_props_data.h, thirdparty/icu4c/common/unicode/ucptrie.h

Related topics: [`unicode-data`](../references/domain-concepts.md#unicode-data-and-properties), [`unicode-normalization`](../references/domain-concepts.md#unicode-normalization)

## Read OpenXR manifests before dispatch forwarding

Runtime selection and generated call forwarding are separate: manifest parsing chooses a runtime before the dispatch table routes calls.

Start in: thirdparty/openxr/src/loader/loader_instance.hpp, thirdparty/openxr/src/loader/manifest_file.cpp, thirdparty/openxr/src/xr_generated_dispatch_table_core.h

Related topics: [`openxr-dispatch`](../references/domain-concepts.md#openxr-dispatch-forwarding), [`openxr-runtime-loading`](../references/domain-concepts.md#openxr-runtime-loading)

## Read VP8L as transforms then symbols

Lossless encoding proceeds from pixel transforms to references, histograms, Huffman trees, and finally VP8L bits.

Start in: thirdparty/libwebp/src/enc/backward_references_enc.c, thirdparty/libwebp/src/enc/histogram_enc.c, thirdparty/libwebp/src/enc/vp8l_enc.c

Related topics: [`histograms-and-huffman-codes`](../references/domain-concepts.md#histograms-and-huffman-codes), [`lossless-transform-pipeline`](../references/domain-concepts.md#lossless-pixel-transform-pipeline)

## Read actions before controller paths

Actions and action sets model application intent; interaction profiles and bindings map that intent to device paths.

Start in: modules/openxr/action_map/openxr_interaction_profile_metadata.h, modules/openxr/openxr_api.h

## Read calls before signals

Callable and Object call paths explain both method dispatch and signal delivery arguments.

Start in: core/object/method_bind.h, core/object/object.h, core/variant/callable.h

Related topics: [`cpp-variadic-binding`](../references/language-concepts.md#c-variadic-binding), [`dynamic-invocation-signals`](../references/domain-concepts.md#dynamic-invocation-and-signals)

## Read compression by public state then algorithm modules

zlib.h and zstd.h identify the public contexts; checksum, entropy, compression, and decompression source files supply the implementation detail.

Start in: thirdparty/zlib/deflate.c, thirdparty/zlib/zlib.h, thirdparty/zstd/compress/zstd_compress.c, thirdparty/zstd/zstd.h

Related topics: [`zlib-stream-codec`](../references/domain-concepts.md#zlib-stream-compression), [`zstandard-stream-codec`](../references/domain-concepts.md#zstandard-stream-compression)

## Read configuration before assuming a path is active

Feature and ISA macros determine which geometry and tasking paths are compiled into a particular build.

Start in: thirdparty/embree/common/sys/sysinfo.h, thirdparty/embree/common/tasking/taskscheduler.h, thirdparty/embree/kernels/config.h

Related topics: [`cxx-conditional-compilation`](../references/language-concepts.md#c-preprocessor-configuration), [`isa-portability`](../references/domain-concepts.md#isa-and-platform-portability)

## Read export from registry to adapter

EditorExport owns presets and targets; EditorExportPlatform performs target behavior; plugins customize the flow.

Start in: editor/export/editor_export.h, editor/export/editor_export_platform.h, editor/export/editor_export_plugin.h, editor/export/editor_export_preset.h

Related topics: [`export-orchestration`](../references/domain-concepts.md#export-orchestration), [`export-presets`](../references/domain-concepts.md#export-presets), [`target-platform-export`](../references/domain-concepts.md#target-platform-export)

## Read lightmapping across CPU and shader files

LightmapperRD's C++ structs and GLSL push-constant and storage layouts form one implementation unit.

Start in: modules/lightmapper_rd/lightmapper_rd.h, modules/lightmapper_rd/lm_common_inc.glsl, modules/lightmapper_rd/lm_compute.glsl

## Read meshoptimizer by pipeline stage

Start with the public structs, then choose clustering, simplification, or codecs rather than reading the subsystem as one algorithm.

Start in: thirdparty/meshoptimizer/clusterizer.cpp, thirdparty/meshoptimizer/indexcodec.cpp, thirdparty/meshoptimizer/meshoptimizer.h, thirdparty/meshoptimizer/simplifier.cpp

## Read method pointer conversion specializations for ABI-facing calls

PtrToArg and GetTypeInfo show how native C++ values are marshalled through runtime method bindings.

Start in: core/variant/binder_common.h, core/variant/method_ptrcall.h, core/variant/type_info.h

Related topics: [`cpp-specialized-marshalling`](../references/language-concepts.md#c-specialized-argument-marshalling), [`reflection-metadata`](../references/domain-concepts.md#reflection-metadata)

## Read motion as a stateful update pipeline

GamepadMotion.hpp combines sample conditioning, orientation/gravity correction, stillness detection, and calibration in one header implementation.

Start in: thirdparty/gamepadmotionhelpers/GamepadMotion.hpp

Related topics: [`cpp-class-state-and-composition`](../references/language-concepts.md#c-class-state-and-composition), [`cpp-numeric-value-operations`](../references/language-concepts.md#c-numeric-value-operations), [`gamepad-motion-fusion`](../references/domain-concepts.md#gamepad-motion-fusion-and-calibration)

## Read platform adapters from their boundaries

DisplayServer classes convert native or browser callbacks into engine input and display behavior.

Start in: platform/linuxbsd/x11/display_server_x11.h, platform/macos/display_server_macos.h, platform/web/display_server_web.cpp, platform/windows/display_server_windows.h

Related topics: [`browser-runtime-adapters`](../references/domain-concepts.md#browser-runtime-adapters), [`platform-display-adaptation`](../references/domain-concepts.md#platform-display-and-window-adaptation)

## Read shader fixtures as input/output pairs

The .glsl files show source conventions while matching .out files show generated C++ embedding.

Start in: tests/python_build/fixtures/glsl/vertex_fragment.glsl, tests/python_build/fixtures/glsl/vertex_fragment.out, tests/python_build/validate_builders.py

## Read shaders in two layers

Shader and ShaderMaterial are high-level resources, whereas RDShaderSource, RDShaderFile, and RDShaderSPIRV belong to RenderingDevice.

Start in: doc/classes/RDShaderSPIRV.xml, doc/classes/RDShaderSource.xml, doc/classes/Shader.xml, doc/classes/ShaderMaterial.xml

Related topics: [`rendering-device-resources`](../references/domain-concepts.md#renderingdevice-descriptors-and-handles), [`shader-programs`](../references/domain-concepts.md#shader-programs-and-material-parameters)

## Read shared ownership from the control block outward

`ReferenceCounter`, `SharedHeader`, and `SharedHandleBase` explain copy, release, parent retention, and final destruction.

Start in: thirdparty/vulkan/include/vulkan/vulkan_shared.hpp

## Read subdivision from basis to grid

Basis evaluation feeds patch evaluation, and patch evaluation feeds grid sampling and tessellation.

Start in: thirdparty/embree/kernels/subdiv/bezier_curve.h, thirdparty/embree/kernels/subdiv/patch_eval_grid.h, thirdparty/embree/kernels/subdiv/tessellation.h

Related topics: [`curve-and-patch-bases`](../references/domain-concepts.md#parametric-curve-bases), [`feature-adaptive-tessellation`](../references/domain-concepts.md#feature-adaptive-tessellation), [`subdivision-surface-evaluation`](../references/domain-concepts.md#subdivision-surface-evaluation)

## Read text controls with text resources

TextEdit and RichTextLabel consume shaped paragraph and glyph data rather than storing only plain strings.

Start in: scene/gui/rich_text_label.h, scene/gui/text_edit.cpp, scene/resources/text_paragraph.h

Related topics: [`text-interface`](../references/domain-concepts.md#text-layout-and-editing)

## Read the C declaration layer first

The repeated C pattern is typed structs plus command-pointer typedefs, with preprocessor guards deciding which declarations appear.

Start in: thirdparty/vulkan/include/vulkan/vulkan_android.h, thirdparty/vulkan/include/vulkan/vulkan_core.h

## Read the SPIR-V module first

SpvReflectShaderModule is the ownership root for entry points and all descriptor, interface, block, and specialization metadata.

Start in: thirdparty/spirv-reflect/spirv_reflect.h

Related topics: [`shader-interface-metadata`](../references/domain-concepts.md#shader-interface-metadata), [`spirv-shader-reflection`](../references/domain-concepts.md#spir-v-shader-reflection)

## Read the paired artifact first

A fixture's `.out` file states the intended observable behavior before implementation details.

Start in: modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd, modules/gdscript/tests/scripts/analyzer/features/lambda_typed.out

Related topics: [`fixture-contracts`](../references/domain-concepts.md#fixture-contracts), [`types-inference-and-conversions`](../references/language-concepts.md#types-inference-and-conversions)

## Read the physics modules as parallel pipelines

The native 2D and 3D modules each combine space ownership, broad phase, narrow phase, constraints, and stepping.

Start in: modules/godot_physics_2d/godot_space_2d.h, modules/godot_physics_3d/godot_space_3d.h

Related topics: [`physics-2d-collision-pipeline`](../references/domain-concepts.md#native-2d-collision-pipeline), [`physics-3d-collision-pipeline`](../references/domain-concepts.md#native-3d-collision-pipeline)

## Read the records as API contracts

Start with creation and presentation records; the supplied files primarily declare external Vulkan data contracts.

Start in: thirdparty/vulkan/include/vulkan/vulkan_structs.hpp, thirdparty/vulkan/include/vulkan/vulkan_vi.h

Related topics: [`vulkan-extensible-records`](../references/domain-concepts.md#vulkan-extensible-records), [`vulkan-instance-setup`](../references/domain-concepts.md#vulkan-instance-setup), [`vulkan-swapchain-presentation`](../references/domain-concepts.md#swapchain-presentation)

## Read warnings as contracts, not commentary

Warning identifiers and exact messages are stored expected outputs.

Start in: modules/gdscript/tests/scripts/analyzer/warnings/unsafe_call_argument.out, modules/gdscript/tests/scripts/parser/warnings/integer_division.out

Related topics: [`diagnostic-expectations`](../references/domain-concepts.md#diagnostic-expectations), [`warnings-and-suppression`](../references/language-concepts.md#warnings-and-selective-suppression)

## Recognize FreeType aggregation units

ftbase.c and autofit.c are deliberate single-object entry units that include their component implementations.

Start in: thirdparty/freetype/src/autofit/autofit.c, thirdparty/freetype/src/base/ftbase.c

Related topics: [`c-preprocessor-composition`](../references/language-concepts.md#c-preprocessor-composition), [`freetype-module-composition`](../references/domain-concepts.md#freetype-module-composition)

## Recognize the Graphite bridge as an adapter

HarfBuzz's Graphite shaper converts between its own face, font, feature, and buffer model and Graphite's C API objects.

Start in: thirdparty/graphite/include/graphite2/Font.h, thirdparty/graphite/include/graphite2/Segment.h, thirdparty/harfbuzz/src/hb-graphite2.cc

Related topics: [`cpp-c-linkage-adapters`](../references/language-concepts.md#c-c-linkage-adapters)

## Recognize the paired text-server designs

Advanced and fallback text servers have parallel shaped-text and font-state structures, so changes should be checked against both implementations.

Start in: modules/text_server_adv/text_server_adv.h, modules/text_server_fb/text_server_fb.h

## Replication policy precedes replication traffic

SceneReplicationConfig is the durable source of spawn, sync, and watch property lists used by spawners and synchronizers.

Start in: modules/multiplayer/multiplayer_spawner.cpp, modules/multiplayer/multiplayer_synchronizer.cpp, modules/multiplayer/scene_replication_config.h

Related topics: [`replicated-property-synchronization`](../references/domain-concepts.md#replicated-property-synchronization), [`replicated-spawning`](../references/domain-concepts.md#replicated-spawning), [`scene-replication-configuration`](../references/domain-concepts.md#scene-replication-configuration)

## Resource modules are self-contained feature slices

Noise, MP3, Ogg, and OpenXR each pair Resource types with module registration and optional editor support.

Start in: modules/mp3/audio_stream_mp3.h, modules/noise/noise.h, modules/ogg/ogg_packet_sequence.h, modules/openxr/action_map/openxr_action_map.h

Related topics: [`mp3-audio-resources`](../references/domain-concepts.md#mp3-audio-resources), [`ogg-packet-sequences`](../references/domain-concepts.md#ogg-packet-sequences), [`openxr-action-configuration`](../references/domain-concepts.md#openxr-action-configuration), [`procedural-noise-textures`](../references/domain-concepts.md#procedural-noise-textures)

## Respect strides in every pixel path

Rows are not assumed contiguous at image width; pointer arithmetic consistently uses the corresponding plane stride.

Start in: thirdparty/libwebp/src/dsp/rescaler.c, thirdparty/libwebp/src/enc/iterator_enc.c, thirdparty/libwebp/src/webp/decode.h

Related topics: [`input-picture-planes`](../references/domain-concepts.md#picture-planes-and-pixel-ownership), [`yuv-rgb-conversion`](../references/domain-concepts.md#yuv-rgb-conversion-and-chroma-processing)

## Script and shader editing are parallel authoring systems

Both have plugin and text-editor layers, but shader editing additionally abstracts editor shader languages and platform baking.

Start in: editor/script/script_editor_plugin.h, editor/shader/editor_shader_language_plugin.h, editor/shader/text_shader_editor.h

Related topics: [`platform-selective-shader-baking`](../references/domain-concepts.md#platform-selective-shader-baking), [`shader-editing-and-language-plugins`](../references/domain-concepts.md#shader-editing-and-language-plugins)

## Script behavior is mostly generated

Before changing reflection or script registration, trace the source generators that produce ScriptPath and member metadata.

Start in: modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPathAttributeGenerator.cs, modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPropertiesGenerator.cs, modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ScriptManagerBridge.cs

Related topics: [`csharp-source-generation`](../references/language-concepts.md#c-partial-types-and-source-generation), [`godot-member-exposure`](../references/domain-concepts.md#godot-member-exposure), [`script-registration-metadata`](../references/domain-concepts.md#managed-script-registration-metadata)

## Separate 2D and 3D physics APIs

The server, direct-state, query-parameter, and motion-result APIs are deliberately paired by dimension.

Start in: doc/classes/PhysicsDirectSpaceState2D.xml, doc/classes/PhysicsDirectSpaceState3D.xml, doc/classes/PhysicsServer2D.xml, doc/classes/PhysicsServer3D.xml

Related topics: [`physics-queries`](../references/domain-concepts.md#physics-queries-and-motion-tests), [`physics-spaces`](../references/domain-concepts.md#physics-spaces-bodies-shapes-and-joints)

## Separate 2D from 3D editing paths

CanvasItemEditor and Node3DEditor are distinct top-level editing implementations; do not assume shared manipulation code.

Start in: editor/scene/3d/node_3d_editor_plugin.h, editor/scene/canvas_item_editor_plugin.h

Related topics: [`scene-authoring`](../references/domain-concepts.md#2d-and-3d-scene-authoring)

## Separate HarfBuzz paint from its output backends

Paint contexts are shared abstractions; raster and vector files provide different realizations.

Start in: thirdparty/harfbuzz/src/hb-paint.hh, thirdparty/harfbuzz/src/hb-raster-paint.hh, thirdparty/harfbuzz/src/hb-vector-paint.hh

Related topics: [`color-font-paint`](../references/domain-concepts.md#color-font-paint-processing), [`raster-font-rendering`](../references/domain-concepts.md#raster-font-rendering), [`vector-font-export`](../references/domain-concepts.md#vector-font-export)

## Separate MSDF generation from its repair pass

MSDF generation produces distance values; error correction is a later, configurable pass over generated bitmap sections.

Start in: thirdparty/msdfgen/core/msdf-error-correction.cpp, thirdparty/msdfgen/core/msdfgen.cpp

Related topics: [`multichannel-distance-fields`](../references/domain-concepts.md#multi-channel-distance-fields)

## Separate PSA policy from TLS protocol flow

Key attributes, slots, and storage live in TF-PSA-Crypto; TLS code owns connection and session progression.

Start in: thirdparty/mbedtls/include/mbedtls/ssl.h, thirdparty/mbedtls/library/ssl_tls.c, thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto.c, thirdparty/mbedtls/tf-psa-crypto/include/psa/crypto_struct.h

Related topics: [`psa-key-lifecycle`](../references/domain-concepts.md#psa-key-lifecycle-and-storage), [`tls-connection-state`](../references/domain-concepts.md#tls-connection-and-session-state)

## Separate SFNT loading from validation

The SFNT component reads tables; otvalid and gxvalid are separate modules that check table structure.

Start in: thirdparty/freetype/src/gxvalid/gxvalid.c, thirdparty/freetype/src/otvalid/otvalid.c, thirdparty/freetype/src/sfnt/sfnt.c

Related topics: [`font-table-validation`](../references/domain-concepts.md#opentype-and-truetypegx-validation), [`sfnt-tables`](../references/domain-concepts.md#sfnt-table-loading)

## Separate Web C++ from browser JavaScript

The Web port intentionally spans C++ adapter classes and JavaScript runtime libraries connected through bridge functions and proxy IDs.

Start in: platform/web/javascript_bridge_singleton.cpp, platform/web/js/libs/library_godot_javascript_singleton.js, platform/web/js/libs/library_godot_runtime.js

Related topics: [`javascript-web-runtime`](../references/language-concepts.md#javascript-browser-runtime-libraries), [`web-javascript-bridge`](../references/domain-concepts.md#web-javascript-bridge)

## Separate WebSocket framing from event dispatch

Wslay keeps frame mechanics and event/message queues in distinct context types and source files.

Start in: thirdparty/wslay/wslay/wslay.h, thirdparty/wslay/wslay_event.c, thirdparty/wslay/wslay_event.h, thirdparty/wslay/wslay_frame.c

Related topics: [`c-function-pointer-apis`](../references/language-concepts.md#c-function-pointer-api-declarations), [`websocket-frame-events`](../references/domain-concepts.md#websocket-framing-and-event-contexts)

## Separate algorithms from packet width

Möller–Trumbore, Plücker, and Woop are algorithm families; the 1, K, M, and specialized forms describe execution shape.

Start in: thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h, thirdparty/embree/kernels/geometry/triangle_intersector_pluecker.h, thirdparty/embree/kernels/geometry/triangle_intersector_woop.h

Related topics: [`simd-ray-packets`](../references/domain-concepts.md#simd-ray-packets), [`triangle-intersection-algorithms`](../references/domain-concepts.md#triangle-intersection-algorithms)

## Separate build from query

PrimRef generation and heuristic builders create BVHs; intersectors traverse the resulting nodes.

Start in: thirdparty/embree/kernels/builders/bvh_builder_sah.h, thirdparty/embree/kernels/builders/primrefgen.h, thirdparty/embree/kernels/bvh/bvh_intersector1.cpp

Related topics: [`bvh-construction`](../references/domain-concepts.md#bvh-construction), [`bvh-traversal`](../references/domain-concepts.md#bvh-traversal), [`primitive-references`](../references/domain-concepts.md#primitive-references)

## Separate byte layout from values

Buffer views select bytes; accessors interpret those bytes as typed data.

Start in: modules/gltf/structures/gltf_accessor.cpp, modules/gltf/structures/gltf_buffer_view.cpp

Related topics: [`cpp-binary-data-codecs`](../references/language-concepts.md#c-binary-data-codecs), [`gltf-accessors`](../references/domain-concepts.md#gltf-binary-accessor-decoding-and-encoding)

## Separate codec direction

Basis code primarily transcodes existing compressed textures, while CVTT computes new compressed blocks from pixel-block input.

Start in: thirdparty/basis_universal/transcoder/basisu_transcoder.cpp, thirdparty/cvtt/ConvectionKernels.h

Related topics: [`basis-transcoding`](../references/domain-concepts.md#basis-texture-transcoding), [`block-texture-encoding`](../references/domain-concepts.md#block-texture-encoding), [`block-texture-transcoding`](../references/domain-concepts.md#gpu-block-texture-conversion)

## Separate geometry from acceleration

Math values and algorithms define the query inputs; BVH and TriangleMesh organize them for repeated spatial queries.

Start in: core/math/aabb.h, core/math/bvh_tree.h, core/math/triangle_mesh.h

Related topics: [`geometry-transforms`](../references/domain-concepts.md#geometry-and-transform-values), [`spatial-indexing`](../references/domain-concepts.md#spatial-indexing-and-ray-queries)

## Separate image coding from container work

Codec paths produce image bytes; mux files arrange those bytes, metadata, and frames into RIFF/WebP chunks.

Start in: thirdparty/libwebp/src/enc/vp8l_enc.c, thirdparty/libwebp/src/enc/webp_enc.c, thirdparty/libwebp/src/mux/muxi.h, thirdparty/libwebp/src/mux/muxinternal.c

Related topics: [`entropy-bitstreams`](../references/domain-concepts.md#entropy-bitstreams)

## Separate imported assets from geometry processing

ufbx imports FBX scene structures, while VHACD consumes explicit mesh points and triangles to produce convex results.

Start in: thirdparty/ufbx/ufbx.h, thirdparty/vhacd/inc/vhacdMesh.h, thirdparty/vhacd/public/VHACD.h

## Separate presentation from widgets

CanvasItem/Viewport/Window provide presentation foundations; Control subclasses provide concrete GUI behavior.

Start in: scene/gui/graph_edit.h, scene/main/canvas_item.h, scene/main/viewport.h

Related topics: [`canvas-and-viewports`](../references/domain-concepts.md#canvas-and-viewport-presentation), [`gui-controls`](../references/domain-concepts.md#gui-controls-and-layout)

## Separate reusable data from placed behavior

Resources carry shared data; Nodes place behavior and relationships into a running scene.

Start in: core/io/resource.h, scene/main/node.h

Related topics: [`resources`](../references/domain-concepts.md#resources-and-asset-lifecycle), [`scene-tree`](../references/domain-concepts.md#scene-graph-and-lifecycle)

## Separate runtime CSG from editor CSG

Runtime shape generation is in csg_shape files; editor gizmo support is isolated under csg/editor and only built for editor builds.

Start in: modules/csg/SCsub, modules/csg/csg_shape.cpp, modules/csg/csg_shape.h, modules/csg/editor/csg_gizmos.h

## Separate the C ABI from C++ conveniences

Use the C header for platform ABI declarations and the C++ headers for typed records, module exports, and diagnostic conversion.

Start in: thirdparty/vulkan/include/vulkan/vulkan_structs.hpp, thirdparty/vulkan/include/vulkan/vulkan_to_string.hpp, thirdparty/vulkan/include/vulkan/vulkan_vi.h, thirdparty/vulkan/include/vulkan/vulkan_video.cppm

Related topics: [`c-compatible-header-guards`](../references/language-concepts.md#c-compatible-guarded-headers), [`cpp-module-interface`](../references/language-concepts.md#c-module-interface-partition), [`cpp-typed-api-records`](../references/language-concepts.md#c-typed-api-records)

## Separate transport APIs from browser adapters

WebRTC and WebSocket C++ abstractions are distinct from their JavaScript browser implementations.

Start in: modules/webrtc/library_godot_webrtc.js, modules/webrtc/webrtc_peer_connection.h, modules/websocket/library_godot_websocket.js, modules/websocket/websocket_peer.h

Related topics: [`webrtc-data-channels`](../references/domain-concepts.md#webrtc-data-channels), [`webrtc-peer-connections`](../references/domain-concepts.md#webrtc-peer-connections), [`websocket-peers`](../references/domain-concepts.md#websocket-peers)

## Separate visibility from drawing

Scene culling builds the visibility and shadow inputs that the forward renderers consume.

Start in: servers/rendering/renderer_rd/forward_clustered/render_forward_clustered.h, servers/rendering/renderer_rd/forward_mobile/render_forward_mobile.h, servers/rendering/renderer_scene_cull.h

Related topics: [`scene-data-and-buffers`](../references/domain-concepts.md#scene-render-data-and-buffers)

## Serialization is explicit and type-aware

Follow RTTI registrations and SaveBinaryState or RestoreBinaryState pairs to understand what data crosses persistence boundaries.

Start in: thirdparty/jolt_physics/Jolt/Core/RTTI.h, thirdparty/jolt_physics/Jolt/Core/StreamUtils.h, thirdparty/jolt_physics/Jolt/ObjectStream/ObjectStream.h

Related topics: [`cxx-reflection-macros`](../references/language-concepts.md#macro-based-rtti-registration), [`cxx-stream-serialization`](../references/language-concepts.md#stream-oriented-binary-serialization), [`serialization`](../references/domain-concepts.md#rtti-based-serialization)

## Shapes drive collision capability

Shape is the common collision contract; concrete shape selection determines the geometry and dispatch path available to a body.

Start in: thirdparty/jolt_physics/Jolt/Physics/Collision/CollisionDispatch.h, thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h

Related topics: [`collision-shapes`](../references/domain-concepts.md#collision-shapes), [`narrow-phase`](../references/domain-concepts.md#narrow-phase-collision-queries)

## Spatial work starts with context and snapshot state

Spatial entities are managed through SpatialContextData, SpatialSnapshotData, component data, and tracker classes.

Start in: modules/openxr/extensions/spatial_entities/openxr_spatial_entities.h, modules/openxr/extensions/spatial_entities/openxr_spatial_entity_extension.h

## Start Vulkan window integration at creation descriptors

The platform header struct shows exactly which native handles must cross into Vulkan.

Start in: thirdparty/vulkan/include/vulkan/vulkan_wayland.h, thirdparty/vulkan/include/vulkan/vulkan_win32.h, thirdparty/vulkan/include/vulkan/vulkan_xcb.h, thirdparty/vulkan/include/vulkan/vulkan_xlib.h

Related topics: [`c-abi-data-structures`](../references/language-concepts.md#c-abi-structures-and-opaque-state)

## Start at Variant

Variant is the interchange value behind dynamic calls, parsing, operators, containers, and many bindings.

Start in: core/variant/variant.h, core/variant/variant_internal.h

Related topics: [`cpp-tagged-storage`](../references/language-concepts.md#c-tagged-storage-and-casts), [`dynamic-variant`](../references/domain-concepts.md#dynamic-runtime-values)

## Start at component assemblers

Files such as sfnt.c and ftcache.c show which implementation files form each FreeType component.

Start in: thirdparty/freetype/src/cache/ftcache.c, thirdparty/freetype/src/sfnt/sfnt.c

Related topics: [`c-preprocessor-composition`](../references/language-concepts.md#c-preprocessor-composition)

## Start at container boundaries

Read the packed headers and descriptor arrays before following codec loops; they determine which bytes each transcoder may interpret.

Start in: thirdparty/basis_universal/transcoder/basisu_file_headers.h, thirdparty/basis_universal/transcoder/basisu_transcoder.h

Related topics: [`basis-container-layout`](../references/domain-concepts.md#basis-file-layout), [`cpp-byte-exact-binary-parsing`](../references/language-concepts.md#c-byte-exact-binary-parsing), [`ktx2-container-transcoding`](../references/domain-concepts.md#ktx2-container-transcoding)

## Start at table routing for HarfBuzz subsetting

The color, layout, ordinary, and variation dispatchers reveal which table families own the actual rewrite logic.

Start in: thirdparty/harfbuzz/src/hb-subset-table-color.cc, thirdparty/harfbuzz/src/hb-subset-table-layout.cc, thirdparty/harfbuzz/src/hb-subset-table-other.cc, thirdparty/harfbuzz/src/hb-subset-table-var.cc

Related topics: [`font-subsetting`](../references/domain-concepts.md#font-subsetting), [`opentype-table-routing`](../references/domain-concepts.md#opentype-table-routing)

## Start at the ABI boundary

NativeFuncs and runtime_interop.cpp must stay ordered and layout-compatible; this is the bridge's most brittle contract.

Start in: modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/NativeFuncs.cs, modules/mono/glue/runtime_interop.cpp

Related topics: [`csharp-unsafe-interop`](../references/language-concepts.md#c-unsafe-interop-and-function-pointers), [`managed-native-interop`](../references/domain-concepts.md#managed-native-interop)

## Start at the OpenXR bridge

OpenXRInterface is the engine-facing adapter; OpenXRAPI contains the central runtime state and internal action/render structures.

Start in: modules/openxr/openxr_api.h, modules/openxr/openxr_interface.h

Related topics: [`openxr-runtime-integration`](../references/domain-concepts.md#openxr-runtime-integration)

## Start at the build boundary

SCsub files reveal which editor subtrees contribute C++ sources and which platform conditions select optional code.

Start in: editor/scene/SCsub, editor/shader/shader_baker/SCsub

## Start at the device graph

GPU work is not executed directly from renderers; it is recorded as typed commands with tracked resource usage and barriers.

Start in: servers/rendering/rendering_device_graph.cpp, servers/rendering/rendering_device_graph.h

Related topics: [`rendering-device-resources`](../references/domain-concepts.md#renderingdevice-descriptors-and-handles)

## Start at the geometry boundary

MeshGLP is the clearest entry point; follow construction into Manifold::Impl, then Halfedge and MeshRelationD.

Start in: thirdparty/manifold/include/manifold/manifold.h, thirdparty/manifold/src/constructors.cpp, thirdparty/manifold/src/impl.h, thirdparty/manifold/src/shared.h

Related topics: [`halfedge-topology`](../references/domain-concepts.md#halfedge-topology), [`manifold-mesh-data`](../references/domain-concepts.md#manifold-mesh-interchange), [`mesh-provenance`](../references/domain-concepts.md#triangle-provenance)

## Start at the platform boundary

Android behavior is split deliberately between JVM host code and C++ native wrappers; inspect both sides of a feature before changing it.

Start in: platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt, platform/android/java_godot_lib_jni.cpp, platform/android/os_android.cpp

Related topics: [`android-jni-interop`](../references/domain-concepts.md#android-jni-interoperation), [`android-platform-host`](../references/domain-concepts.md#android-platform-host)

## Start at the public boundary

Read encode.h and decode.h before internal files; they define input representation, callbacks, initialization, and ownership.

Start in: thirdparty/libwebp/src/webp/decode.h, thirdparty/libwebp/src/webp/encode.h

Related topics: [`input-picture-planes`](../references/domain-concepts.md#picture-planes-and-pixel-ownership)

## Start at the public handles

Trace RTCDevice to RTCScene, RTCGeometry, and RTCBuffer first; these form the external ownership boundary.

Start in: thirdparty/embree/include/embree4/rtcore_buffer.h, thirdparty/embree/include/embree4/rtcore_device.h, thirdparty/embree/include/embree4/rtcore_geometry.h

Related topics: [`buffer-storage`](../references/domain-concepts.md#geometry-buffer-storage), [`device-runtime`](../references/domain-concepts.md#device-runtime), [`scene-commit`](../references/domain-concepts.md#scene-construction-and-commit)

## Start at the simulation façade

PhysicsSystem is the entry point; its getters expose bodies, broad phase, narrow phase, filtering, and settings.

Start in: thirdparty/jolt_physics/Jolt/Physics/PhysicsSystem.cpp, thirdparty/jolt_physics/Jolt/Physics/PhysicsSystem.h

Related topics: [`broad-phase`](../references/domain-concepts.md#broad-phase-collision-detection), [`narrow-phase`](../references/domain-concepts.md#narrow-phase-collision-queries)

## Start at the test entry point

The supplied first-party C++ partition is organized around a shared test runtime and focused API test files.

Start in: tests/test_main.cpp, tests/test_main.h

Related topics: [`cpp-classes-and-ref-handles`](../references/language-concepts.md#c-inheritance-and-engine-reference-handles)

## Start from public shader objects

TShader owns the incoming source configuration; TProgram is the linking and reflection boundary.

Start in: thirdparty/glslang/glslang/MachineIndependent/reflection.h, thirdparty/glslang/glslang/Public/ShaderLang.h

Related topics: [`shader-interface-analysis`](../references/domain-concepts.md#shader-interface-mapping-and-reflection), [`shader-source-compilation`](../references/domain-concepts.md#shader-source-compilation)

## Start from the scene library

scene/SCsub shows which large runtime subsystems are assembled into the scene library.

Start in: scene/2d/node_2d.h, scene/3d/node_3d.h, scene/SCsub

Related topics: [`scene-graph`](../references/domain-concepts.md#scene-graph-nodes), [`three-dimensional-physics`](../references/domain-concepts.md#3d-physics-nodes), [`two-dimensional-physics`](../references/domain-concepts.md#2d-physics-nodes)

## Start with build selection

drivers/SCsub is the shortest route to discovering which driver families are included conditionally.

Start in: drivers/SCsub, drivers/vulkan/SCsub

## Start with configuration

ProjectSettings is the source of project properties, autoloads, feature overrides, and resource-pack setup.

Start in: core/config/project_settings.cpp, core/config/project_settings.h

Related topics: [`packed-resource-files`](../references/domain-concepts.md#packed-resource-files), [`project-settings`](../references/domain-concepts.md#project-settings-and-feature-overrides)

## Start with resource boundaries

Most persisted engine data is a Resource; follow ResourceLoader, ResourceSaver, and ResourceUID before tracing specialized assets.

Start in: doc/classes/Resource.xml, doc/classes/ResourceLoader.xml, doc/classes/ResourceSaver.xml, doc/classes/ResourceUID.xml

Related topics: [`resource-formats`](../references/domain-concepts.md#resource-formats-and-serialization), [`resource-identifiers`](../references/domain-concepts.md#resource-and-server-identifiers)

## Start with runtime ownership

Node and SceneTree establish the runtime hierarchy before inspecting any specialized scene type.

Start in: scene/main/node.cpp, scene/main/node.h, scene/main/scene_tree.h

Related topics: [`scene-tree-execution`](../references/domain-concepts.md#scenetree-execution)

## Start with the C versus C++ split

Platform headers are C ABI adapters; the `vulkan_hpp_*` headers are C++ convenience and lifetime layers.

Start in: thirdparty/vulkan/include/vulkan/vulkan_hpp_macros.hpp, thirdparty/vulkan/include/vulkan/vulkan_ios.h, thirdparty/vulkan/include/vulkan/vulkan_raii.hpp

Related topics: [`cpp-preprocessor-configuration`](../references/language-concepts.md#c-preprocessor-configuration)

## Start with the independent vendor boundaries

Navigation, RVO, SDL, and SPIR-V tooling are separate subsystems in this partition; do not infer direct runtime coupling merely from co-location.

Start in: thirdparty/recastnavigation/Recast/Source/RecastRegion.cpp, thirdparty/rvo2/rvo2_2d/RVOSimulator2d.h, thirdparty/sdl/SDL.c, thirdparty/spirv-cross/spirv_cross.hpp

Related topics: [`navigation-regions`](../references/domain-concepts.md#navigation-region-construction)

## Trace a scene through its editor state

Begin with EditedScene, then follow its history_id and stored editor state before reading UI code.

Start in: editor/editor_data.cpp, editor/editor_data.h, editor/editor_undo_redo_manager.h

Related topics: [`editing-selection-history`](../references/domain-concepts.md#editing-selection-history), [`editor-scene-sessions`](../references/domain-concepts.md#editor-scene-sessions), [`undo-redo-history`](../references/domain-concepts.md#undo-and-redo-history)

## Trace asynchronous work through two queues

MessageQueue handles deferred calls, while WorkerThreadPool represents background task and group execution.

Start in: core/object/message_queue.h, core/object/worker_thread_pool.h, core/os/thread.h

Related topics: [`deferred-execution`](../references/domain-concepts.md#deferred-calls-and-worker-tasks), [`synchronization-primitives`](../references/domain-concepts.md#threads-and-synchronization)

## Trace documentation through its model

Read State.parse_class, TypeName.from_element, and make_rst_class together to follow XML from parsing to output.

Start in: doc/tools/doc_status.py, doc/tools/make_rst.py

Related topics: [`class-reference-generation`](../references/domain-concepts.md#class-reference-generation)

## Trace language-server requests through the workspace

Protocol and text-document layers delegate symbol-oriented work to GDScriptWorkspace and extended parser results.

Start in: modules/gdscript/language_server/gdscript_extend_parser.cpp, modules/gdscript/language_server/gdscript_language_protocol.cpp, modules/gdscript/language_server/gdscript_text_document.cpp, modules/gdscript/language_server/gdscript_workspace.cpp

Related topics: [`gdscript-language-server`](../references/domain-concepts.md#gdscript-language-server-support)

## Trace primitive data first

Start at SubGrid or TriangleM, then follow the matching intersector family to see how geometry becomes a ray-hit candidate.

Start in: thirdparty/embree/kernels/geometry/subgrid.h, thirdparty/embree/kernels/geometry/subgrid_intersector.h, thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h

Related topics: [`ray-primitive-intersection`](../references/domain-concepts.md#ray-primitive-intersection), [`subgrid-intersection`](../references/domain-concepts.md#subgrid-intersection)

## Trace the Metal bridge before platform graphics behavior

Most Metal-facing methods are inline selector dispatchers; begin at the wrapper declaration and then follow the descriptor or encoder type.

Start in: thirdparty/metal-cpp/Metal/MTL4Archive.hpp, thirdparty/metal-cpp/Metal/MTLPrivate.hpp

Related topics: [`cpp-native-wrappers`](../references/language-concepts.md#c-inline-native-wrappers), [`metal-cpp-object-bridge`](../references/domain-concepts.md#metal-cpp-object-bridge)

## Trace the common runtime spine

Start with Object and ClassDB, then follow Node and SceneTree; most public systems attach to that spine.

Start in: core/object/class_db.h, scene/main/node.h, scene/main/scene_tree.h

Related topics: [`object-model`](../references/domain-concepts.md#engine-object-model), [`reflection`](../references/domain-concepts.md#class-metadata-and-reflection), [`scene-tree`](../references/domain-concepts.md#scene-graph-and-lifecycle)

## Treat Array and Dictionary as Variant containers

Their validation and dynamic semantics are implemented in core/variant, while lower-level storage primitives are in core/templates.

Start in: core/templates/cowdata.h, core/variant/array.h, core/variant/dictionary.h

Related topics: [`cpp-copy-on-write-containers`](../references/language-concepts.md#c-copy-on-write-container-storage), [`generic-containers`](../references/domain-concepts.md#generic-container-infrastructure), [`variant-containers`](../references/domain-concepts.md#dynamic-arrays-and-dictionaries)

## Treat CSG as a tree plus a mesh kernel

Read csg_tree.h before boolean3.cpp and boolean_result.cpp; the former defines composition, while the latter performs geometric result assembly.

Start in: thirdparty/manifold/src/boolean3.cpp, thirdparty/manifold/src/boolean_result.cpp, thirdparty/manifold/src/csg_tree.h

Related topics: [`mesh-provenance`](../references/domain-concepts.md#triangle-provenance)

## Treat EditorPlugin as the extension seam

Many user-visible tools are narrow EditorPlugin subclasses, so locate the relevant plugin before following UI implementation details.

Start in: editor/scene/3d/mesh_library_editor_plugin.h, editor/shader/shader_editor_plugin.h

Related topics: [`cpp-plugin-specialization`](../references/language-concepts.md#c-plugin-specialization), [`editor-plugin-lifecycle`](../references/domain-concepts.md#editor-plugin-lifecycle)

## Treat GLSL files as explicit interfaces

Attribute locations, uniforms, inputs, and outputs are declared directly in shader source and compiled by the GLES3 build script.

Start in: drivers/gles3/shaders/SCsub, drivers/gles3/shaders/canvas_occlusion.glsl

## Treat Graphite behavior as font data

SILF rules and glyph data are decoded from the face; the engine executes them rather than hard-coding script behavior.

Start in: thirdparty/graphite/src/Pass.cpp, thirdparty/graphite/src/Silf.cpp, thirdparty/graphite/src/inc/Face.h

Related topics: [`font-table-access`](../references/domain-concepts.md#binary-font-table-access), [`graphite-rule-execution`](../references/domain-concepts.md#graphite-silf-rule-execution), [`graphite-shaping`](../references/domain-concepts.md#graphite-segment-shaping)

## Treat Image as the codec boundary

Compression modules share Image format, mipmap, and byte-buffer operations rather than a common codec implementation.

Start in: modules/astcenc/image_compress_astcenc.cpp, modules/basis_universal/image_compress_basisu.cpp, modules/cvtt/image_compress_cvtt.cpp, modules/etcpak/image_compress_etcpak.cpp

Related topics: [`image-codec-registration`](../references/domain-concepts.md#image-codec-integration)

## Treat Jolt as an alternate implementation

JoltPhysicsServer3D adapts Godot engine concepts to JPH types rather than extending the native GodotPhysicsServer3D implementation.

Start in: modules/jolt_physics/jolt_physics_server_3d.h, modules/jolt_physics/objects/jolt_body_3d.cpp

Related topics: [`physics-space-queries`](../references/domain-concepts.md#3d-physics-query-contracts)

## Treat RTCRayHit as the query contract

Intersection writes hit fields while tfar in the embedded ray tracks the current closest distance.

Start in: thirdparty/embree/include/embree4/rtcore_ray.h, thirdparty/embree/kernels/geometry/intersector_epilog.h

Related topics: [`hit-results`](../references/domain-concepts.md#intersection-results), [`primitive-intersection`](../references/domain-concepts.md#primitive-intersection), [`ray-query`](../references/domain-concepts.md#ray-query-state)

## Treat SDL events as owned queue records

SDL_Event values are wrapped by SDL_EventEntry, which also tracks temporary payload memory and queue linkage.

Start in: thirdparty/sdl/events/SDL_events.c, thirdparty/sdl/events/SDL_eventwatch_c.h, thirdparty/sdl/include/SDL3/SDL_events.h

Related topics: [`sdl-event-queue`](../references/domain-concepts.md#event-queue-and-watches), [`sdl-thread-abstractions`](../references/domain-concepts.md#thread-and-synchronization-abstractions)

## Treat UPnP results as linked parser-owned data

Discovery and port-list parsing return linked records, so inspect allocation and cleanup paths with the list traversal.

Start in: thirdparty/miniupnpc/include/miniupnpc/portlistingparse.h, thirdparty/miniupnpc/include/miniupnpc/upnpdev.h, thirdparty/miniupnpc/src/portlistingparse.c

Related topics: [`upnp-device-discovery`](../references/domain-concepts.md#upnp-device-discovery), [`upnp-port-mapping`](../references/domain-concepts.md#upnp-port-mapping-control)

## Treat Variant transitions as intentional test cases

The suite repeatedly tests behavior when static type information becomes `Variant`.

Start in: modules/gdscript/tests/scripts/analyzer/features/local_inference_is_weak.gd, modules/gdscript/tests/scripts/runtime/errors/invalid_cast.gd

Related topics: [`types-inference-and-conversions`](../references/language-concepts.md#types-inference-and-conversions), [`warnings-and-suppression`](../references/language-concepts.md#warnings-and-selective-suppression)

## Treat VisualShader as persisted graph state

Read VisualShader::Graph, Node, and Connection before investigating individual shader-node classes or editor behavior.

Start in: modules/visual_shader/visual_shader.cpp, modules/visual_shader/visual_shader.h, modules/visual_shader/vs_nodes/visual_shader_nodes.h

Related topics: [`visual-shader-nodes`](../references/domain-concepts.md#visual-shader-nodes), [`visual-shader-vector-operations`](../references/domain-concepts.md#visual-shader-vector-operations)

## Treat Wayland XML as protocol state machines

Requests, events, errors, and destructors define object lifecycle and externally exchanged data.

Start in: thirdparty/wayland-protocols/stable/linux-dmabuf/linux-dmabuf-v1.xml, thirdparty/wayland-protocols/staging/linux-drm-syncobj/linux-drm-syncobj-v1.xml

Related topics: [`explicit-drm-syncobj`](../references/domain-concepts.md#explicit-drm-synchronization-objects)

## Treat WebPPicture as the ownership boundary

The picture can borrow caller planes or own allocations; only free memory allocated through the picture API.

Start in: thirdparty/libwebp/src/enc/picture_enc.c, thirdparty/libwebp/src/webp/encode.h

Related topics: [`input-picture-planes`](../references/domain-concepts.md#picture-planes-and-pixel-ownership)

## Treat actions as event lists

InputMap::Action contains a deadzone and List<Ref<InputEvent>>, while concrete InputEvent subclasses carry device-specific data.

Start in: core/input/input_event.h, core/input/input_map.h

Related topics: [`cpp-engine-polymorphism`](../references/language-concepts.md#c-inheritance-virtual-interfaces-and-ref-ownership), [`input-events-actions`](../references/domain-concepts.md#input-events-and-actions)

## Treat enet_godot.cpp as the integration boundary

The ENet transport core is C code; the UDP and DTLS bridge is expressed as C++ socket classes in the Godot-specific file.

Start in: thirdparty/enet/callbacks.c, thirdparty/enet/enet_godot.cpp

Related topics: [`cpp-classes-inheritance`](../references/language-concepts.md#c-classes-and-inheritance), [`godot-enet-socket-adaptation`](../references/domain-concepts.md#godot-enet-socket-adaptation)

## Treat export as a pipeline

Android export code prepares Gradle and resource inputs; the Gradle project then builds the modules and the editor can sign or verify results.

Start in: platform/android/export/export_plugin.cpp, platform/android/java/editor/src/main/java/org/godotengine/editor/utils/ApkSignerUtil.kt, platform/android/java/settings.gradle

Related topics: [`android-export-pipeline`](../references/domain-concepts.md#android-export-pipeline), [`android-gradle-assembly`](../references/domain-concepts.md#android-gradle-assembly), [`apk-signing`](../references/domain-concepts.md#apk-signing-and-verification)

## Treat extension classes as contracts

Names ending in Extension identify replacement hooks; required virtual methods define the minimum implementation surface.

Start in: doc/classes/PhysicsServer2DExtension.xml, doc/classes/RenderSceneBuffersExtension.xml, doc/classes/TextServerExtension.xml

Related topics: [`extensions`](../references/domain-concepts.md#virtual-implementation-extensions), [`runtime-metadata`](../references/domain-concepts.md#runtime-class-metadata)

## Treat extensions as the feature boundary

Optional OpenXR capabilities are organized as OpenXRExtensionWrapper subclasses, including platform graphics adapters.

Start in: modules/openxr/extensions/SCsub, modules/openxr/extensions/openxr_extension_wrapper.h

Related topics: [`openxr-extension-wrappers`](../references/domain-concepts.md#openxr-extension-wrappers)

## Treat pNext as a structural convention

Many Vulkan inputs and query outputs begin with a type discriminator and a next-structure pointer, including Android external-memory properties.

Start in: thirdparty/vulkan/include/vulkan/vulkan.hpp, thirdparty/vulkan/include/vulkan/vulkan_android.h

## Treat platform create-info structs as external API contracts

Each surface record is passed across the Vulkan boundary and has the recurring `sType` and `pNext` extensibility fields.

Start in: thirdparty/vulkan/include/vulkan/vulkan_metal.h, thirdparty/vulkan/include/vulkan/vulkan_ohos.h, thirdparty/vulkan/include/vulkan/vulkan_screen.h

Related topics: [`extension-structure-chains`](../references/domain-concepts.md#vulkan-extension-structure-chains)

## Treat previews as asynchronous

Do not assume preview generation is immediate; it is queued, threaded, and cached.

Start in: editor/inspector/editor_resource_preview.cpp, editor/inspector/editor_resource_preview.h

Related topics: [`resource-previews`](../references/domain-concepts.md#asynchronous-resource-previews)

## Treat render work as tasks

SVG/WebP loading and software rendering use Task-derived types, so scheduling is part of the data path rather than an isolated utility.

Start in: thirdparty/thorvg/src/loaders/svg/tvgSvgLoader.h, thirdparty/thorvg/src/renderer/sw_engine/tvgSwRenderer.cpp, thirdparty/thorvg/src/renderer/tvgTaskScheduler.h

Related topics: [`task-scheduling`](../references/domain-concepts.md#task-scheduling-and-synchronization)

## Treat scenes as resources

PackedScene delegates stored structure to SceneState node and connection records.

Start in: scene/resources/packed_scene.cpp, scene/resources/packed_scene.h

Related topics: [`resource-assets`](../references/domain-concepts.md#resource-backed-assets)

## Treat server implementations as replaceable

Physics, text, XR, and rendering use derived extension, dummy, or wrapper classes behind stable server interfaces.

Start in: servers/physics_3d/physics_server_3d_dummy.h, servers/physics_3d/physics_server_3d_extension.h, servers/text/text_server_extension.h, servers/xr/xr_interface_extension.h

Related topics: [`cpp-class-inheritance`](../references/language-concepts.md#c-class-inheritance-for-backend-contracts), [`pluggable-server-backends`](../references/domain-concepts.md#pluggable-and-wrapped-server-backends)

## Treat state as caller-owned workspace

KTX2 and Brotli retain operational state across calls, so reset and lifetime boundaries are part of correct use.

Start in: thirdparty/basis_universal/transcoder/basisu_transcoder.h, thirdparty/brotli/dec/state.h

Related topics: [`brotli-stream-decompression`](../references/domain-concepts.md#brotli-stream-decompression), [`c-stateful-streaming-api`](../references/language-concepts.md#c-stateful-streaming-apis), [`ktx2-container-transcoding`](../references/domain-concepts.md#ktx2-container-transcoding)

## Treat subset input and plan as the subsetting spine

Input describes selection; the plan transports derived state into table processors.

Start in: thirdparty/harfbuzz/src/hb-subset-input.hh, thirdparty/harfbuzz/src/hb-subset-plan.hh

Related topics: [`font-subsetting`](../references/domain-concepts.md#font-subsetting)

## Treat the cache as borrowed storage

Small-bitmap and glyph cache nodes are manager-controlled, bounded cache data rather than caller-owned permanent assets.

Start in: thirdparty/freetype/src/cache/ftcmanag.h, thirdparty/freetype/src/cache/ftcmru.h, thirdparty/freetype/src/cache/ftcsbits.h

Related topics: [`glyph-caching`](../references/domain-concepts.md#glyph-and-face-caching)

## Treat the extension manifest as executable test configuration

The compatibility runner and .gdextension file jointly define the external test boundary.

Start in: tests/compatibility_test/godot/compatibility_test.gdextension, tests/compatibility_test/run_compatibility_test.py

## Treat the filesystem as an index

Directory and FileInfo records are the editor's project view; resource dependencies are resolved from their stored metadata.

Start in: editor/file_system/editor_file_system.cpp, editor/file_system/editor_file_system.h

Related topics: [`filesystem-project-index`](../references/domain-concepts.md#project-filesystem-index), [`resource-dependency-resolution`](../references/domain-concepts.md#resource-dependency-resolution)

## Treat this as vendored media code

The supplied partition is organized by upstream libraries, not by Godot engine subsystems; begin with each library's public header and its context object.

Start in: thirdparty/libjpeg-turbo/src/jpeglib.h, thirdparty/libktx/include/ktx.h, thirdparty/libogg/ogg/ogg.h

Related topics: [`image-decode-pipeline`](../references/domain-concepts.md#image-decode-pipelines), [`ktx-texture-container`](../references/domain-concepts.md#ktx-texture-containers), [`ogg-pages-and-packets`](../references/domain-concepts.md#ogg-pages-packets-and-bit-packing)

## Understand translation as domain selection

Translation contains catalogs, while TranslationDomain and TranslationServer decide which catalog and locale apply.

Start in: core/string/translation.h, core/string/translation_domain.h, core/string/translation_server.h

Related topics: [`localization`](../references/domain-concepts.md#translation-and-locale-selection)

## Understand what a scene file contains

Read SceneState before PackedScene: its node and connection tables explain the persisted graph.

Start in: scene/resources/packed_scene.h

Related topics: [`dynamic-values`](../references/domain-concepts.md#dynamic-values-and-dictionaries), [`packed-scenes`](../references/domain-concepts.md#packed-scene-serialization)

## Use LSR to understand locale matching

LocaleMatcher operates on locales plus compact language-script-region values and distance data.

Start in: thirdparty/icu4c/common/localematcher.cpp, thirdparty/icu4c/common/lsr.h, thirdparty/icu4c/common/unicode/localematcher.h

Related topics: [`locale-resolution`](../references/domain-concepts.md#locale-resolution)

## Use StringName for identity-bearing text

Class, property, signal, path, and translation code repeatedly uses StringName instead of raw String.

Start in: core/string/node_path.h, core/string/string_name.h, core/string/translation.h

Related topics: [`localization`](../references/domain-concepts.md#translation-and-locale-selection), [`string-names-paths`](../references/domain-concepts.md#strings-interned-names-and-node-paths)

## Use animation tests to map resource semantics

Track-type coverage is explicit and compact in the Animation tests.

Start in: tests/scene/test_animation.cpp, tests/scene/test_animation_player.cpp

## Use geometry type to find leaf behavior

Mesh and curve source classes provide bounds and primitive data; corresponding geometry intersectors define hit tests.

Start in: thirdparty/embree/kernels/common/scene_curves.h, thirdparty/embree/kernels/common/scene_triangle_mesh.h, thirdparty/embree/kernels/geometry/quad_intersector_moeller.h

Related topics: [`geometry-families`](../references/domain-concepts.md#geometry-families), [`primitive-intersection`](../references/domain-concepts.md#primitive-intersection)

## Use inspector plugins for property UI behavior

Specialized property editing is commonly implemented through EditorInspectorPlugin and EditorProperty rather than the scene editor.

Start in: editor/scene/gradient_editor_plugin.h, editor/scene/gui/control_editor_plugin.h

Related topics: [`inspector-property-editors`](../references/domain-concepts.md#custom-inspector-property-editors), [`resource-specific-editors`](../references/domain-concepts.md#resource-specific-editing)

## Use properties as shared SDL metadata

SDL properties appear beneath hints and I/O streams, so property ownership and lifetime matter when tracing configuration.

Start in: thirdparty/sdl/SDL_hints.c, thirdparty/sdl/SDL_properties.c, thirdparty/sdl/io/SDL_iostream.c

Related topics: [`sdl-iostreams`](../references/domain-concepts.md#abstract-i-o-streams), [`sdl-runtime-properties`](../references/domain-concepts.md#runtime-property-groups-and-hints)

## Use protocol types before UI code

The debug adapter's types and protocol serializer define the externally exchanged debugger data before higher-level editor views consume it.

Start in: editor/debugger/debug_adapter/debug_adapter_protocol.cpp, editor/debugger/debug_adapter/debug_adapter_types.h

Related topics: [`debug-adapter-protocol`](../references/domain-concepts.md#debug-adapter-protocol-integration)

## Use public codec headers as the texture-codec map

The headers distinguish decode entry points from ETC/EAC and BC compression entry points before the large implementation files add optimization details.

Start in: thirdparty/etcpak/DecodeRGB.hpp, thirdparty/etcpak/ProcessDxtc.hpp, thirdparty/etcpak/ProcessRGB.hpp

Related topics: [`codec-simd-specialization`](../references/domain-concepts.md#codec-simd-specialization), [`texture-block-codecs`](../references/domain-concepts.md#texture-block-codecs)

## Use scalar DSP as the algorithmic baseline

Read scalar DSP files first, then compare SSE, NEON, MSA, or MIPS files as equivalent target-specific implementations.

Start in: thirdparty/libwebp/src/dsp/enc.c, thirdparty/libwebp/src/dsp/enc_neon.c, thirdparty/libwebp/src/dsp/enc_sse2.c

Related topics: [`cpu-specialized-dsp`](../references/domain-concepts.md#cpu-specialized-dsp-backends)

## Use scene input when judging completion

Completion results can depend on the referenced `.tscn` hierarchy and attached script types.

Start in: modules/gdscript/tests/scripts/completion/get_node/get_node.tscn, modules/gdscript/tests/scripts/completion/get_node/literal_scene/dollar_class_scene.cfg

Related topics: [`completion-contexts`](../references/domain-concepts.md#contextual-completion-contracts), [`nodepaths-and-indexed-access`](../references/language-concepts.md#node-paths-and-indexed-access), [`scene-contexts`](../references/domain-concepts.md#scene-aware-test-context)

## Use scene nodes for OpenXR presentation

Composition layers and render models have dedicated Node3D-derived scene types rather than being modeled only in the API core.

Start in: modules/openxr/scene/openxr_composition_layer.h, modules/openxr/scene/openxr_render_model.h

Related topics: [`openxr-composition-layers`](../references/domain-concepts.md#openxr-composition-layers), [`openxr-render-models`](../references/domain-concepts.md#openxr-render-models)

## Use server headers as subsystem boundaries

Audio, physics, navigation, and display functionality is grouped behind server interfaces and their query or adapter types.

Start in: servers/audio/audio_server.h, servers/display/display_server.h, servers/navigation_2d/navigation_server_2d.h, servers/physics_2d/physics_server_2d.h

Related topics: [`audio-bus-processing`](../references/domain-concepts.md#audio-buses-streams-and-effects), [`navigation-queries`](../references/domain-concepts.md#2d-and-3d-navigation-queries), [`physics-server-queries`](../references/domain-concepts.md#2d-and-3d-physics-queries), [`platform-presentation`](../references/domain-concepts.md#display-camera-video-and-movie-services)

## Use tests as executable language behavior

The analyzer feature and error scripts define concrete accepted and rejected GDScript cases.

Start in: modules/gdscript/tests/gdscript_test_runner.cpp

Related topics: [`gdscript-typed-collections`](../references/language-concepts.md#gdscript-typed-collections-and-signature-compatibility)

## Use the C++ façade as a separate API layer

Vulkan-Hpp adds templates, dispatch objects, typed flags, ownership wrappers, and standard-library return types over the C declarations.

Start in: thirdparty/vulkan/include/vulkan/vulkan.hpp, thirdparty/vulkan/include/vulkan/vulkan_enums.hpp, thirdparty/vulkan/include/vulkan/vulkan_extension_inspection.hpp

Related topics: [`cpp-templates`](../references/language-concepts.md#c-templates)

## Use the directory as the test boundary

`parser`, `analyzer`, and `runtime` represent distinct validation stages.

Related topics: [`diagnostic-expectations`](../references/domain-concepts.md#diagnostic-expectations)

## Use the instrumented app as the integration contract

The Android test project exercises plugin signals, storage paths, and Java class wrapping together.

Start in: platform/android/java/app/src/androidTestInstrumented/java/com/godot/game/GodotAppTest.kt, platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd, platform/android/java/app/src/instrumented/assets/test/javaclasswrapper/java_class_wrapper_tests.gd

Related topics: [`android-instrumented-tests`](../references/domain-concepts.md#android-instrumented-integration-tests), [`android-plugin-signals`](../references/domain-concepts.md#android-plugins-and-signals), [`android-storage-bridge`](../references/domain-concepts.md#android-storage-bridge)

## Use the plug-in registry

New editor capabilities generally enter through EditorPlugin registration methods rather than direct central switches.

Start in: editor/plugins/editor_plugin.cpp, editor/plugins/editor_plugin.h

Related topics: [`cpp-runtime-polymorphism`](../references/language-concepts.md#c-virtual-interfaces), [`editor-plugin-extension`](../references/domain-concepts.md#editor-plug-in-extension)

## WebP has three execution layers

Separate header parsing and incremental decode from demux and animation composition; they are implemented in different directories and state objects.

Start in: thirdparty/libwebp/src/dec/idec_dec.c, thirdparty/libwebp/src/dec/webp_dec.c, thirdparty/libwebp/src/demux/anim_decode.c

Related topics: [`webp-riff-decoding`](../references/domain-concepts.md#webp-chunk-parsing-incremental-decode-and-animation)
<!-- rope-ladder:end document -->
