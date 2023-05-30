package api

//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/agent.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/agent_data.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/api_contracts.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/api_faction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/api_register.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/api_ship.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/api_systems.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/chart.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/connected_system.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/contract.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/contract_deliver_good.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/contract_payment.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/contract_terms.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/contract_type.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/cooldown.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/error.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/extraction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/extraction_yield.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/faction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/faction_trait.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/jump_gate.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/market.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/market_trade_good.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/market_transaction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/meta.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/register_faction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/scanned_ship.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/scanned_system.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/scanned_waypoint.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_cargo.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_cargo_item.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_crew.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_engine.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_frame.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_fuel.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_module.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_mount.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_nav.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_nav_flight_mode.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_nav_route.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_nav_route_waypoint.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_reactor.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_registration.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_requirements.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_role.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/ship_type.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/shipyard.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/shipyard_ship.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/shipyard_transaction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/status.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/survey.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/survey_deposit.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/survey_size.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/system.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/system_faction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/system_type.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/system_waypoint.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/trade_good.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/trade_symbol.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/waypoint.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/waypoint_faction.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/waypoint_orbital.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/waypoint_trait.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/waypoint_trait_symbol.proto
//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/waypoint_type.proto
