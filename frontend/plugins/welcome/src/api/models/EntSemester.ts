/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntSemesterEdges,
    EntSemesterEdgesFromJSON,
    EntSemesterEdgesFromJSONTyped,
    EntSemesterEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSemester
 */
export interface EntSemester {
    /**
     * Semester holds the value of the "Semester" field.
     * @type {string}
     * @memberof EntSemester
     */
    semester?: string;
    /**
     * 
     * @type {EntSemesterEdges}
     * @memberof EntSemester
     */
    edges?: EntSemesterEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSemester
     */
    id?: number;
}

export function EntSemesterFromJSON(json: any): EntSemester {
    return EntSemesterFromJSONTyped(json, false);
}

export function EntSemesterFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSemester {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'semester': !exists(json, 'Semester') ? undefined : json['Semester'],
        'edges': !exists(json, 'edges') ? undefined : EntSemesterEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntSemesterToJSON(value?: EntSemester | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Semester': value.semester,
        'edges': EntSemesterEdgesToJSON(value.edges),
        'id': value.id,
    };
}


