// This file has been auto-generated by Warthog.  Do not update directly as it
// will be re-written.  If you need to change this file, update models or add
// new TypeGraphQL objects
// prettier-ignore
// @ts-ignore
import { DateResolver as Date } from 'graphql-scalars';
// prettier-ignore
// @ts-ignore
import { GraphQLID as ID } from 'graphql';
// prettier-ignore
// @ts-ignore
import { ArgsType, Field as TypeGraphQLField, Float, InputType as TypeGraphQLInputType, Int } from 'type-graphql';
// prettier-ignore
// @ts-ignore
import { registerEnumType, GraphQLISODateTime as DateTime } from "type-graphql";

import * as BN from "bn.js";

// prettier-ignore
// @ts-ignore eslint-disable-next-line @typescript-eslint/no-var-requires
const { GraphQLJSONObject } = require('graphql-type-json');
// prettier-ignore
// @ts-ignore
import { BaseWhereInput, JsonObject, PaginationArgs, DateOnlyString, DateTimeString, BigInt, Bytes } from 'warthog';

// @ts-ignore
import { Entity } from "../src/modules/entity/entity.model";
// @ts-ignore
import { Transfer } from "../src/modules/transfer/transfer.model";

export enum EntityOrderByEnum {
  createdAt_ASC = "createdAt_ASC",
  createdAt_DESC = "createdAt_DESC",

  updatedAt_ASC = "updatedAt_ASC",
  updatedAt_DESC = "updatedAt_DESC",

  deletedAt_ASC = "deletedAt_ASC",
  deletedAt_DESC = "deletedAt_DESC",

  name_ASC = "name_ASC",
  name_DESC = "name_DESC",

  countryId_ASC = "countryId_ASC",
  countryId_DESC = "countryId_DESC",

  cityId_ASC = "cityId_ASC",
  cityId_DESC = "cityId_DESC"
}

registerEnumType(EntityOrderByEnum, {
  name: "EntityOrderByInput"
});

@TypeGraphQLInputType()
export class EntityWhereInput {
  @TypeGraphQLField(() => ID, { nullable: true })
  id_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  id_in?: string[];

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_eq?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_lt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_lte?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_gt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_gte?: Date;

  @TypeGraphQLField(() => ID, { nullable: true })
  createdById_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  createdById_in?: string[];

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_eq?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_lt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_lte?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_gt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_gte?: Date;

  @TypeGraphQLField(() => ID, { nullable: true })
  updatedById_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  updatedById_in?: string[];

  @TypeGraphQLField({ nullable: true })
  deletedAt_all?: Boolean;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_eq?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_lt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_lte?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_gt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_gte?: Date;

  @TypeGraphQLField(() => ID, { nullable: true })
  deletedById_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  deletedById_in?: string[];

  @TypeGraphQLField({ nullable: true })
  name_eq?: string;

  @TypeGraphQLField({ nullable: true })
  name_contains?: string;

  @TypeGraphQLField({ nullable: true })
  name_startsWith?: string;

  @TypeGraphQLField({ nullable: true })
  name_endsWith?: string;

  @TypeGraphQLField(() => [String], { nullable: true })
  name_in?: string[];

  @TypeGraphQLField(() => BigInt, { nullable: true })
  countryId_eq?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  countryId_gt?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  countryId_gte?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  countryId_lt?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  countryId_lte?: BN;

  @TypeGraphQLField(() => [BigInt], { nullable: true })
  countryId_in?: BN[];

  @TypeGraphQLField(() => BigInt, { nullable: true })
  cityId_eq?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  cityId_gt?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  cityId_gte?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  cityId_lt?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  cityId_lte?: BN;

  @TypeGraphQLField(() => [BigInt], { nullable: true })
  cityId_in?: BN[];
}

@TypeGraphQLInputType()
export class EntityWhereUniqueInput {
  @TypeGraphQLField(() => ID)
  id?: string;
}

@TypeGraphQLInputType()
export class EntityCreateInput {
  @TypeGraphQLField({ nullable: true })
  name?: string;

  @TypeGraphQLField(() => BigInt)
  countryId!: BN;

  @TypeGraphQLField(() => BigInt)
  cityId!: BN;
}

@TypeGraphQLInputType()
export class EntityUpdateInput {
  @TypeGraphQLField({ nullable: true })
  name?: string;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  countryId?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  cityId?: BN;
}

@ArgsType()
export class EntityWhereArgs extends PaginationArgs {
  @TypeGraphQLField(() => EntityWhereInput, { nullable: true })
  where?: EntityWhereInput;

  @TypeGraphQLField(() => EntityOrderByEnum, { nullable: true })
  orderBy?: EntityOrderByEnum;
}

@ArgsType()
export class EntityCreateManyArgs {
  @TypeGraphQLField(() => [EntityCreateInput])
  data!: EntityCreateInput[];
}

@ArgsType()
export class EntityUpdateArgs {
  @TypeGraphQLField() data!: EntityUpdateInput;
  @TypeGraphQLField() where!: EntityWhereUniqueInput;
}

export enum TransferOrderByEnum {
  createdAt_ASC = "createdAt_ASC",
  createdAt_DESC = "createdAt_DESC",

  updatedAt_ASC = "updatedAt_ASC",
  updatedAt_DESC = "updatedAt_DESC",

  deletedAt_ASC = "deletedAt_ASC",
  deletedAt_DESC = "deletedAt_DESC",

  from_ASC = "from_ASC",
  from_DESC = "from_DESC",

  to_ASC = "to_ASC",
  to_DESC = "to_DESC",

  value_ASC = "value_ASC",
  value_DESC = "value_DESC",

  comment_ASC = "comment_ASC",
  comment_DESC = "comment_DESC",

  block_ASC = "block_ASC",
  block_DESC = "block_DESC"
}

registerEnumType(TransferOrderByEnum, {
  name: "TransferOrderByInput"
});

@TypeGraphQLInputType()
export class TransferWhereInput {
  @TypeGraphQLField(() => ID, { nullable: true })
  id_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  id_in?: string[];

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_eq?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_lt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_lte?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_gt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  createdAt_gte?: Date;

  @TypeGraphQLField(() => ID, { nullable: true })
  createdById_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  createdById_in?: string[];

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_eq?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_lt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_lte?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_gt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  updatedAt_gte?: Date;

  @TypeGraphQLField(() => ID, { nullable: true })
  updatedById_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  updatedById_in?: string[];

  @TypeGraphQLField({ nullable: true })
  deletedAt_all?: Boolean;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_eq?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_lt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_lte?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_gt?: Date;

  @TypeGraphQLField(() => DateTime, { nullable: true })
  deletedAt_gte?: Date;

  @TypeGraphQLField(() => ID, { nullable: true })
  deletedById_eq?: string;

  @TypeGraphQLField(() => [ID], { nullable: true })
  deletedById_in?: string[];

  @TypeGraphQLField(() => Bytes, { nullable: true })
  from_eq?: Buffer;

  @TypeGraphQLField(() => [Bytes], { nullable: true })
  from_in?: Buffer[];

  @TypeGraphQLField(() => Bytes, { nullable: true })
  to_eq?: Buffer;

  @TypeGraphQLField(() => [Bytes], { nullable: true })
  to_in?: Buffer[];

  @TypeGraphQLField(() => BigInt, { nullable: true })
  value_eq?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  value_gt?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  value_gte?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  value_lt?: BN;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  value_lte?: BN;

  @TypeGraphQLField(() => [BigInt], { nullable: true })
  value_in?: BN[];

  @TypeGraphQLField({ nullable: true })
  comment_eq?: string;

  @TypeGraphQLField({ nullable: true })
  comment_contains?: string;

  @TypeGraphQLField({ nullable: true })
  comment_startsWith?: string;

  @TypeGraphQLField({ nullable: true })
  comment_endsWith?: string;

  @TypeGraphQLField(() => [String], { nullable: true })
  comment_in?: string[];

  @TypeGraphQLField(() => Int, { nullable: true })
  block_eq?: number;

  @TypeGraphQLField(() => Int, { nullable: true })
  block_gt?: number;

  @TypeGraphQLField(() => Int, { nullable: true })
  block_gte?: number;

  @TypeGraphQLField(() => Int, { nullable: true })
  block_lt?: number;

  @TypeGraphQLField(() => Int, { nullable: true })
  block_lte?: number;

  @TypeGraphQLField(() => [Int], { nullable: true })
  block_in?: number[];
}

@TypeGraphQLInputType()
export class TransferWhereUniqueInput {
  @TypeGraphQLField(() => ID)
  id?: string;
}

@TypeGraphQLInputType()
export class TransferCreateInput {
  @TypeGraphQLField(() => Bytes)
  from!: Buffer;

  @TypeGraphQLField(() => Bytes)
  to!: Buffer;

  @TypeGraphQLField(() => BigInt)
  value!: BN;

  @TypeGraphQLField({ nullable: true })
  comment?: string;

  @TypeGraphQLField()
  block!: number;
}

@TypeGraphQLInputType()
export class TransferUpdateInput {
  @TypeGraphQLField(() => Bytes, { nullable: true })
  from?: Buffer;

  @TypeGraphQLField(() => Bytes, { nullable: true })
  to?: Buffer;

  @TypeGraphQLField(() => BigInt, { nullable: true })
  value?: BN;

  @TypeGraphQLField({ nullable: true })
  comment?: string;

  @TypeGraphQLField({ nullable: true })
  block?: number;
}

@ArgsType()
export class TransferWhereArgs extends PaginationArgs {
  @TypeGraphQLField(() => TransferWhereInput, { nullable: true })
  where?: TransferWhereInput;

  @TypeGraphQLField(() => TransferOrderByEnum, { nullable: true })
  orderBy?: TransferOrderByEnum;
}

@ArgsType()
export class TransferCreateManyArgs {
  @TypeGraphQLField(() => [TransferCreateInput])
  data!: TransferCreateInput[];
}

@ArgsType()
export class TransferUpdateArgs {
  @TypeGraphQLField() data!: TransferUpdateInput;
  @TypeGraphQLField() where!: TransferWhereUniqueInput;
}
